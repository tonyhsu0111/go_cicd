package command

import (
	"log"
	"net/http"
	"os/exec"
	"time"

	sshclient "go_cicd/component"
)

func WinLaunch() {
	cmd := exec.Command("cmd", "dir")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(out))
}

func Ssh_wincmd01() string {
	var svnConfig sshclient.SshConfig
	svnConfig.SshHost = "192.168.150.91"
	svnConfig.SshUser = "administrator"
	svnConfig.SshPassword = "1qasde32W"
	svnConfig.SshType = "password"    // password或者key 這裡是key型別登陸遠端伺服器
	svnConfig.SshKeyPath = "./id_rsa" // ssh id_rsa.id路徑
	svnConfig.SshPort = 22
	testCmd := "D: && cd d:\\PublishCenter\\sea_batch\\ && dir"
	return sshclient.SSHRemoteExcute(svnConfig, testCmd)

}
func WinCICD_CMD01(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(Ssh_wincmd01()))
	time.Sleep(1 * time.Second)

}
