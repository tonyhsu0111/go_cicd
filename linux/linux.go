package linux

import (
	"fmt"
	sshclient "go_cicd/component"
	"io"
	"log"
	"net/http"
	"os/exec"
	"time"
)

// type sshConfig interface {
// 	sshClient() string
// 	sshHost() string
// 	sshUser() string
// 	sshPasswrod() string
// 	sshType() string    // password或者key
// 	sshKeyPath() string // ssh id_rsa.id路徑
// 	sshPort() int
// }

func ssh_cmd01() {
	var svnConfig sshclient.SSHConfig

	svnConfig.SSHHost = "192.168.150.94"
	svnConfig.SSHUser = "root"
	svnConfig.SSHPasswrod = "1qasde32"
	svnConfig.SSHType = "password"    // password或者key 這裡是key型別登陸遠端伺服器
	svnConfig.SSHKeyPath = "./id_rsa" // ssh id_rsa.id路徑
	svnConfig.SSHPort = 22
	testCmd := "ls -al"
	sshclient.SSHRemoteExcute(svnConfig, testCmd)
}

func CICD_CMD02(w http.ResponseWriter, r *http.Request) {
	ssh_cmd01()
	time.Sleep(1 * time.Second)
	io.WriteString(w, "<h1>Hi io.WriteString</h1>")

}

func Linux_cmd01() {
	cmd := exec.Command("ls", "-lah")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))

}

func CICD_CMD01(w http.ResponseWriter, r *http.Request) {
	Linux_cmd01()
	time.Sleep(1 * time.Second)
	// readText, err := ioutil.ReadFile("deploy.txt")
	// fmt.Println(string(readText))
	// if err != nil {
	// 	log.Fatal(err) // if err exists log fetal and exit
	// }
	// fmt.Fprintf(w, "%s", string(readText))
	io.WriteString(w, "<h1>Hi io.WriteString</h1>")

}
