package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	router := mux.NewRouter()
	// наше приложение будет слушать запросы на localhost:8080/api/hello
	router.HandleFunc("/api/hello", HelloHandler).Methods("GET")
	http.ListenAndServe(":8080", router)
}

//package main
//
//import (
//	"fmt"
//	"net/http"
//	"strconv"
//)
//
//var counter int
//
//func GetHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodGet {
//		fmt.Fprintln(w, "Counter равен", strconv.Itoa(counter))
//	} else {
//		fmt.Fprintln(w, "Поддерживается только метод GET")
//	}
//}
//
//func PostHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodGet {
//		counter++
//		fmt.Fprintln(w, "Counter увеличен на 1")
//	} else {
//		fmt.Fprintln(w, "Поддерживается только метод POST")
//
//	}
//}
//
//func main() {
//	http.HandleFunc("/get", GetHandler)
//	http.HandleFunc("/post", PostHandler)
//	http.ListenAndServe("localhost:8080", nil)
//}
