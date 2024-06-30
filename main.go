package main
import (
	"fmt"
	"net/http"
)
func homeage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, User")
}
func main() {
http.HandleFunc("/", homePage)
http.ListenAndServe(":4000", nil)
}

//what do i need to put in terminal for go.mod 
//how to fix this error on lne 10
//what to expect from these doing	