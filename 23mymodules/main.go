// 1. go mod init github.com/sejoonkim/mymodules
// 2. if using module -> go mod tidy
// 3. go mod verify : check how many modules are verified
// 4. go list : go list all : go list -m all : go list -m -versions github.com/gorilla/mux
// 5. go mod why github.com/gorilla/ mux
// 6. go mod graph
// 7. go mod edit -go 1.16 : go mod edit -module
// 8. go mod vendor = node modules
// 9. go run -mod=vendor main.go : first looks into vendor folder for customized codes

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}
