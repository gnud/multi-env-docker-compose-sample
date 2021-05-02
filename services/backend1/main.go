package main

import (
	"fmt"
	"log"
    "os"
	"net/http"
)

var defaultPort = "8001"
var choicenPort = defaultPort

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w, `
            ##         .
        ## ## ##        ==
    ## ## ## ## ##    ===
    /"""""""""""""""""\___/ ===
    {                       /  ===-
    \______ O           __/
   ====\    \         __/
        \____\_______/
	
Hello from Docker(%s)!
`, choicenPort)
    fmt.Println(r.URL)
    fmt.Println()
}

func main() {
    
    var envPort = os.Getenv("BACKEND1_PORT")
    
    if len(envPort) > 0 {
        choicenPort = envPort
    }
    
    fmt.Printf("Port to use %s.", choicenPort)
    var listenPort = fmt.Sprintf(":%s", choicenPort)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(listenPort, nil))
}
