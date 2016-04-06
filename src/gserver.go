package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sort"
	s "strings"
)

// App build information
var (
	Version string
	Build   string
)

type indexHandler func(http.ResponseWriter, *http.Request)

// Commandline flags
var verbose bool
var echoWebsocket bool
var addr string
var port string
var dataDir string

func init() {
	flag.BoolVar(&verbose, "v", false, "Verbose output")
	flag.BoolVar(&echoWebsocket, "websocket", false, "Open a websocket on /echo")
	flag.StringVar(&addr, "addr", "0.0.0.0", "Address to serve on")
	flag.StringVar(&port, "port", "9000", "Port to listen on")
	flag.StringVar(&dataDir, "data", "data", "json file names will be converted to rest paths")
}

func getIndex(entries map[string]string) indexHandler {
	if len(entries) == 0 {
		return func(res http.ResponseWriter, req *http.Request) {
			const NoData string = `<html><title>gserver</title><body>
				<h2>No files to serve</h2>
				<p>Put some JSON files to be served inside a folder. The default name is <strong>data</strong> in the startup folder.</p>
				<p>The filenames will define the REST API. For example, using the name <strong>'api_v1_todos.json'</strong> will result in the URL endpoint
				<strong>'/api/v1/todos'</strong> providing the content of JSON file. <strong>'api_v1_todos_{id:[0-9]+}.json'</strong> will result in the URL endpoint
				<strong>'/api/v1/todos/<number>'</strong> where the regex must match.</p>
				</body></html>`

			res.Header().Set("Access-Control-Allow-Origin", "*")
			res.Header().Set("Content-Type", "text/html;charset=utf-8")
			res.WriteHeader(http.StatusOK)
			fmt.Fprintf(res, NoData)
		}
	}
	// We have a rest API
	return func(res http.ResponseWriter, req *http.Request) {
		const APIDataStart string = "<html><title>gserver</title><body><h2>The following endpoints are available:</h2>"
		const APIDataEnd string = "</body></html>"
		var urls []string
		for key := range entries {
			urls = append(urls, "\n<li><a href=\""+key+"\">"+key+"</a></li>")
		}
		sort.Strings(urls)
		restUrls := "<ul style=\"list-style-type:none\">" + s.Join(urls, "") + "\n</ul>"
		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Content-Type", "text/html;charset=utf-8")
		res.WriteHeader(http.StatusOK)
		fmt.Fprintf(res, APIDataStart +restUrls+ APIDataEnd)
	}

}

func webSocket(res http.ResponseWriter, req *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	res.Header().Set("Access-Control-Allow-Origin", "*")

	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		log.Printf("Not a web socket connection: %s \n", err)
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

func createHandler(path string, file string) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			return
		}

		res.Header().Set("Access-Control-Allow-Origin", "*")
		res.Header().Set("Content-Type", "application/json;charset=utf-8")
		res.WriteHeader(http.StatusOK)
		fmt.Fprintln(res, string(content))
		if verbose {
			log.Println("GET " + path)
		}
	}
}

func isDir(pth string) (bool, error) {
	fi, err := os.Stat(pth)
	if err != nil {
		return false, err
	}
	return fi.IsDir(), nil
}

// Validate data directory
func readFiles() map[string]string {

	sep := string(os.PathSeparator)
	entries := make(map[string]string, 0)

	// Check if directory exists
	isD, err := isDir(dataDir)
	if !isD || err != nil {
		log.Println("No data directory found: " + dataDir)
		return entries
	}

	// Get files from the data directory
	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		log.Println("Local dir '" + dataDir + "' couldn't be read.")
		return entries
	}

	if len(files) == 0 {
		log.Println("No files found in directory " + dataDir)
	}

	var fileName string
	var canonical string
	var path string

	for i := 0; i < len(files); i++ {
		fileName = files[i].Name()
		if s.HasSuffix(fileName, ".json") && !files[i].IsDir() {
			canonical = fileName[0 : len(fileName)-5]
			path = "/" + s.Replace(canonical, "_", "/", -1)
			entries[path] = "." + sep + dataDir + sep + fileName
		}
	}

	return entries
}

func main() {
	flag.Parse()

	log.SetPrefix("[gserver] ")

	entries := readFiles()

	router := mux.NewRouter()

	if verbose {
		log.Println("Simple Go Server version " + Version)
		log.Println("(build " + Build + ")")
	}

	// Index page
	index := getIndex(entries)
	router.HandleFunc("/", index)

	// Websocket
	if echoWebsocket {
		log.Println("Adding handler for /echo")
		router.HandleFunc("/echo", webSocket)
	}

	// Register all 'simulated' endpoints
	for path, file := range entries {
		if verbose {
			log.Println("Adding handler for", path)
		}
		router.HandleFunc(path, createHandler(path, file))
	}

	addrPort := addr + ":" + port
	listener, err := net.Listen("tcp4", addrPort)
	if err != nil {
		log.Fatal("Unable to ", err)
	}

	log.Println("Server is running at http://" + addrPort)
	log.Fatal(http.Serve(listener, router))
}
