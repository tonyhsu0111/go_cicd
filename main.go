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

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/version", Version).Methods("POST")
	r.HandleFunc("/linuxcmd", command.SshCmd).Methods("POST")
	r.HandleFunc("/windowscmd", command.SshCmd).Methods("POST")
	r.HandleFunc("/test", command.CICD_CMD01).Methods("POST")
	log.Println("service starting...")
	log.Fatal(http.ListenAndServe(":8400", r)) //設定監聽的埠

	fmt.Println(fmt.Sprintf("%0.f", 1.01/3.002))

}
