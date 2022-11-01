package command

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	sshclient "go_cicd/component"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

type DeployJson struct {
	Host string `json:"Host,omitempty"`
	// User     string `json:"User,omitempty"`
	// Password string `json:"Password,omitempty"`
	Port    int    `json:"Port,omitempty"`
	Command string `json:"command,omitempty"`
}

const Port = 22

func Ssh_Agent(Host, User, Password, Cmd string) string {
	var agentConfig sshclient.SshConfig
	agentConfig.SshHost = Host
	agentConfig.SshUser = User
	agentConfig.SshPassword = Password
	agentConfig.SshType = "password"    // password或者key 這裡是key型別登陸遠端伺服器
	agentConfig.SshKeyPath = "./id_rsa" // ssh id_rsa.id路徑
	agentConfig.SshPort = Port
	CommandLine := Cmd
	return sshclient.SSHRemoteExcute(agentConfig, CommandLine)

}

func ParseBasicAuth(auth string) []string {
	const prefix = "Basic "
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return []string{}
	}
	cs := string(c)
	// fmt.Println(cs)
	str1 := strings.Split(cs, ":")
	return str1
}

func SshCmd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	deployjson := new(DeployJson)
	dataBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	err = json.Unmarshal(dataBytes, deployjson)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	// 走BasicAuthName 方式
	auth := r.Header.Get("Authorization")

	authlist := ParseBasicAuth(auth)
	User := authlist[0]     // array[0] name
	Password := authlist[1] // array[0] name
	// fmt.Println("Username:"+User, "Password:"+Password)
	// fmt.Println(fmt.Sprintf("%s", deployjson))

	Host := deployjson.Host
	// 傳Json 方式
	// User := deployjson.User
	// Password := deployjson.Password
	CommandLine := deployjson.Command
	w.Write([]byte(Ssh_Agent(Host, User, Password, CommandLine)))
	time.Sleep(1 * time.Second)

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
