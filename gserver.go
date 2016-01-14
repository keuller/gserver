package main

import (
    "fmt"
    "log"
    "github.com/gorilla/mux"
    "github.com/gorilla/websocket"
    "io/ioutil"
    "net/http"
    s "strings"
    "os"
    "path/filepath"
)

func index(res http.ResponseWriter, req *http.Request) {
    var b string = `<html><title>gserver</title><body style='font-family:Helvetica,Arial,Tahoma'><h2>Simple Go Server works!</h2>` +
                   `<p>Put your JSON files inside <strong>data</strong> folder.</p>` +
                   `<p>For example using the pattern <strong>'api_v1_todos.json'</strong> the  URL <strong>'/api/v1/todos'</strong> will be generated automatically, providing the content of JSON file.</p></body></html>`

    res.Header().Set("Access-Control-Allow-Origin", "*")
    res.Header().Set("Content-Type", "text/html;charset=utf-8")
    res.WriteHeader(http.StatusOK)
    fmt.Fprintf(res, b)
}

func webSocket(res http.ResponseWriter, req *http.Request) {
    var upgrader = websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
        CheckOrigin: func(r *http.Request) bool { return true },
    }

    res.Header().Set("Access-Control-Allow-Origin", "*")

    conn, err := upgrader.Upgrade(res, req, nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            return
        }
        if err = conn.WriteMessage(messageType, p); err != nil {
            return
        }
    }
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

    entries := make(map[string]string, 0)

    // checks if directory exists
    _, err := os.Stat(dir)
    if err != nil || os.IsNotExist(err) {
      return entries
    }

    // reads directory 'data'
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        log.Println("Local 'data' dir couldn't be read.")
        return entries
    }

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

func get_addr(val ...string) string {
    for _, str := range val {
        if s.HasPrefix(str, "--addr") {
            return str[7:len(str)]
        }
    }
    return ""
}

func get_port(val ...string) string {
    for _, str := range val {
        if s.HasPrefix(str, "--port") {
            return str[7:len(str)]
        }
    }
    return ""
}

func main() {
    var addr string = "0.0.0.0"
    var port string = "9000"

    log.SetPrefix("[gserver] ")
    entries := readFiles()
    router := mux.NewRouter()

    count := len(os.Args)
    if count == 3 {
      addr = get_addr(string(os.Args[1]), string(os.Args[2]))
      port = get_port(string(os.Args[1]), string(os.Args[2]))
    }

    if count == 2 {
      addr = get_addr(string(os.Args[1]))
      if addr == "" {
        addr = "0.0.0.0"
      }
      port = get_port(string(os.Args[1]))
      if port == "" {
        port = "9000"
      }
    }

    // gets the current path
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        return
    }

    fmt.Println("Simple Go Server version 1.1.1")
    fmt.Println("Server is running at http://"+ addr +":" + port)
    fmt.Println("")

    publicDir := string(dir)
    fmt.Println("Current directory:", publicDir)

    // simple doc page
    fmt.Println("Adding handler for /doc")
    router.HandleFunc("/doc", index)

    // websockets
    fmt.Println("Adding handler for /echo")
    router.HandleFunc("/echo", webSocket)

    // register all 'simulated' endpoints
    for path, file := range entries {
        fmt.Println("Adding handler for", path)
        router.HandleFunc(path, createHandler(file))
    }

    // provide static files from current directory
    router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(publicDir))))

    http.ListenAndServe(addr + ":" + port, router)
}
