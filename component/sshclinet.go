package component

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/crypto/ssh"
)

// ssh配置型別
type SshConfig struct {
	SshHost     string
	SshUser     string
	SshPassword string
	SshType     string // password或者key
	SshKeyPath  string // ssh id_rsa.id路徑
	SshPort     int
}

func PublicKeyAuthFunc(kPath string) ssh.AuthMethod {
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

func SSHRemoteExcute(sshCfg SshConfig, cmd string) string {

	// 建立ssh登入配置
	config := &ssh.ClientConfig{
		Timeout:         time.Second, // ssh連線time out時間一秒鐘,如果ssh驗證錯誤會在一秒鐘返回
		User:            sshCfg.SshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 這個可以,但是不夠安全
	}
	if sshCfg.SshType == "password" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshCfg.SshPassword)}
	} else {
		config.Auth = []ssh.AuthMethod{PublicKeyAuthFunc(sshCfg.SshKeyPath)}
		// return
	}

	// dial 獲取ssh client
	addr := fmt.Sprintf("%s:%d", sshCfg.SshHost, sshCfg.SshPort)
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
	return fmt.Sprintf("cmd output:%s", string(combo))
}
