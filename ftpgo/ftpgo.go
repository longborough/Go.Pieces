package main

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
    "os"
)

func logopen(sys string, user string, pw string) sftp.Client {
	addr := sys + ":22"
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pw),
		},
		//Ciphers: []string{"3des-cbc", "aes256-cbc", "aes192-cbc", "aes128-cbc"},
	}
	conn, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	client, err := sftp.NewClient(conn)
	if err != nil {
		panic("Failed to create client: " + err.Error())
	}
	return *client
}

func main() {
     var this os.FileInfo
	logs := logopen("tksftp.thy.com", "alcslog", "pL4fTy9%tG")
	dir, err := logs.ReadDir("./ALCSRO")
    println(err.Error())
	for _, this = range dir { 
    println(this.Name())
}
    
//	logs.Close()
}
