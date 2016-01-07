package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "io/ioutil"
    "net/http"
    s "strings"
    "os"
    "path/filepath"
)

func index(res http.ResponseWriter, req *http.Request) {
    var b string = `<html><title>JSON Server</title><body style='font-family:Helvetica,Arial,Tahoma'><h2>Simple Server is works!</h2>` +
                   `<p>Put your JSON files inside <strong>data</strong> folder.</p>` +
                   `<p>For example using the pattern <strong>'api_v1_todos.json'</strong> the  URL <strong>'/api/v1/todos'</strong> will be generated automatically, providing the content of JSON file.</p></body></html>`

    res.Header().Set("Access-Control-Allow-Origin", "*")
    res.Header().Set("Content-Type", "text/html;charset=utf-8")
    res.WriteHeader(http.StatusOK)
    fmt.Fprintf(res, b)
}

func createHandler(file string) func(res http.ResponseWriter, req *http.Request) {
    return func(res http.ResponseWriter, req *http.Request) {
        content, err := ioutil.ReadFile(file)
        if err != nil {
            return
        }

        res.Header().Set("Access-Control-Allow-Origin", "*")
        res.Header().Set("Content-Type", "application/json;charset=utf-8")
        res.WriteHeader(http.StatusOK)
        fmt.Fprintln(res, string(content))
    }
}

func readFiles() map[string]string {
    sep := string(os.PathSeparator)
    dir := "."+ sep + "data"

    // checks if directory exists
    _, err := os.Stat(dir)
    if err != nil || os.IsNotExist(err) {
        os.Mkdir(dir, 0766)
    }

    // reads directory 'data'
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        panic(err)
        return nil
    }

    entries := make(map[string]string, 0)
    var fileName string
    var canonical string
    var path string

    for i := 0; i < len(files); i++ {
        fileName = files[i].Name()
        if s.HasSuffix(fileName, ".json") && !files[i].IsDir() {
            canonical = fileName[0:len(fileName)-5]
            path = "/" + s.Replace(canonical, "_", "/", -1)
            entries[path] = "."+ sep + "data" + sep + fileName
        }
    }

    return entries
}

func main() {
    entries := readFiles()
    router := mux.NewRouter()

    // gets the current path
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        return
    }

    publicDir := string(dir)
    fmt.Println("Static directory file", publicDir)

    for path, file := range entries {
        fmt.Println("creating handler for", path)
        router.HandleFunc(path, createHandler(file))
    }

    fmt.Println("creating handler for /doc")
    router.HandleFunc("/doc", index)

    // provide static files from current directory
    router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(publicDir))))

    fmt.Println("Server is running at http://0.0.0.0:9000")
    http.ListenAndServe(":9000", router)
}
