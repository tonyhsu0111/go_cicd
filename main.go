package main

import (
	"fmt"
	"go_cicd/command"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Version(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Version: v0.001") //這個寫入到 w 的是輸出到客戶端的
}
func Ssh_Command_Function() {
	r := mux.NewRouter()
	r.HandleFunc("/version", Version).Methods("GET")
	r.HandleFunc("/linuxcmd", command.SshCmd).Methods("POST")
	r.HandleFunc("/windowscmd", command.SshCmd).Methods("POST")
	log.Println("service starting...")
	log.Fatal(http.ListenAndServe(":8000", r)) //設定監聽的埠
}

func main() {
	go Ssh_Command_Function()
	r := mux.NewRouter()
	r.HandleFunc("/version", Version).Methods("GET")
	r.HandleFunc("/linuxcmd", command.SshCmd).Methods("POST")
	r.HandleFunc("/windowscmd", command.SshCmd).Methods("POST")
	log.Println("service starting...")
	log.Fatal(http.ListenAndServe(":8001", r)) //設定監聽的埠

}
