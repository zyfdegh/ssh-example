package main

import (
	"bytes"
	"fmt"
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
	out := execute("ls")
	fmt.Println(out)
}

func execute(cmd string) (stdout string) {
	// private key
	// signer, err := ssh.ParsePrivateKey([]byte(key))
	// if err != nil {
	// 	log.Printf("parse private key error: %v", err)
	// 	return
	// }
	// config := &ssh.ClientConfig{
	// 	User: user,
	// 	Auth: []ssh.AuthMethod{
	// 		ssh.PublicKeys(signer),
	// 	},
	// }
	// password
	config := &ssh.ClientConfig{
		User: "root",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
	}
	client, err := ssh.Dial("tcp", "192.168.59.129:22", config)
	if err != nil {
		log.Printf("ssh dial error: %v", err)
		return
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Printf("new session error: %v", err)
		return
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	err = session.Run(cmd)
	if err != nil {
		log.Printf("run cmd error: %v", err)
		return
	}
	return b.String()
}
