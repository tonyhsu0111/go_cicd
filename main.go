package main

import (
	"fmt"
	"go_cicd/linux"
	"go_cicd/windows"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func sayHello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello CICD Test !") //這個寫入到 w 的是輸出到客戶端的
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", sayHello).Methods("POST")
	r.HandleFunc("/test", linux.CICD_CMD01).Methods("POST")
	r.HandleFunc("/linuxtest02", linux.CICD_CMD02).Methods("POST")
	r.HandleFunc("/wintest01", windows.WinCICD_CMD01).Methods("POST")
	log.Println("service starting...")
	log.Fatal(http.ListenAndServe(":8400", r)) //設定監聽的埠

	fmt.Println(fmt.Sprintf("%0.f", 1.01/3.002))

}
