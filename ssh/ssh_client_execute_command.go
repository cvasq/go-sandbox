// Establishes an SSH connection and runs a "df -h" command

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/crypto/ssh"
)

var username = "ubuntu"
var host = "192.168.1.10:22"
var port = "22"
var privateKeyFilePath = "/home/toor/.ssh/private_key.pem"

func getKeySigner(privateKeyFile string) ssh.Signer {
	privateKeyData, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatal("Error loading private key file: ", err)
	}

	privateKey, err := ssh.ParsePrivateKey(privateKeyData)
	if err != nil {
		log.Fatal("Error parsing private key file: ", err)
	}

	return privateKey
}

func main() {

	privateKey := getKeySigner(privateKeyFilePath)
	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(privateKey),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		log.Fatal("Error dialing ssh server: ", err)
	}

	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("df -h"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())

}
