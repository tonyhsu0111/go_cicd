package windows

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
)

func winLaunch() {
	cmd := exec.Command("cmd", "dir")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(out))
}

// ssh配置型別
type sshConfig struct {
	sshHost     string
	sshUser     string
	sshPasswrod string
	sshType     string // password或者key
	sshKeyPath  string // ssh id_rsa.id路徑
	sshPort     int
}

func publicKeyAuthFunc(kPath string) ssh.AuthMethod {
	keyPath, err := homedir.Expand(kPath)
	if err != nil {
		log.Fatal("find key's home dir failed", err)
	}

	key, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal("ssh key file read failed", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatal("ssh key signer failed", err)
	}
	return ssh.PublicKeys(signer)
}

func sshRemoteExcute(sshCfg sshConfig, cmd string) {

	// 建立ssh登入配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second, // ssh連線time out時間一秒鐘,如果ssh驗證錯誤會在一秒鐘返回
		User:            sshCfg.sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 這個可以,但是不夠安全
	}
	if sshCfg.sshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshCfg.sshPasswrod)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshCfg.sshKeyPath)}
		// return
	}

	// dial 獲取ssh client
	addr := fmt.Sprintf("%s:%d", sshCfg.sshHost, sshCfg.sshPort)
	sshClient, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatal("genarate ssh client fail", err)
	}
	defer sshClient.Close()

	// 建立ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("genarate ssh client fail", err)
	}

	defer session.Close()

	// 執行遠端命令
	combo, err := session.CombinedOutput(cmd)
	if err != nil {
		log.Fatal("remote excute cmd fail", err)
	}
	log.Println("cmd output:", string(combo))
}

func ssh_wincmd01() {
	var svnConfig sshConfig
	svnConfig.sshHost = "192.168.150.90"
	svnConfig.sshUser = "administrator"
	svnConfig.sshPasswrod = "1qasde32W"
	svnConfig.sshType = "password"    // password或者key 這裡是key型別登陸遠端伺服器
	svnConfig.sshKeyPath = "./id_rsa" // ssh id_rsa.id路徑
	svnConfig.sshPort = 22
	testCmd := "D: && cd d:\\PublishCenter\\sea_batch\\ && dir"
	sshRemoteExcute(svnConfig, testCmd)

}
func WinCICD_CMD01(w http.ResponseWriter, r *http.Request) {
	ssh_wincmd01()
	time.Sleep(1 * time.Second)
	io.WriteString(w, "<h1>Hi io.WriteString</h1>")

}
