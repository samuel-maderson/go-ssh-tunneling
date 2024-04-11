package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"

	"go-ssh-tunneling/src/types"
)

var (
	rhost         string
	rport         int
	lhost         string
	lport         int
	bastionhost   string
	key           string
	user          string
	custom_config = types.Config{}
)

func init() {

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(data, &custom_config)
	rhost = custom_config.Remote.Host
	rport = custom_config.Remote.Port
	lhost = custom_config.Local.Host
	lport = custom_config.Local.Port
	bastionhost = custom_config.Bastion.Host
	key = custom_config.Key
	user = custom_config.User
}

func StartTunneling(rport int, rhost string, lhost string, lport int, key string) {

	cmd := exec.Command("ssh", "-f", "-N", "-L", strconv.Itoa(lport)+":"+rhost+":"+strconv.Itoa(rport), "-i", key, user+"@"+bastionhost)

	output, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(output))
}

func main() {

	log.Println("\033[1;32m[+]\033[0m Starting SSH Tunneling")
	StartTunneling(rport, rhost, lhost, lport, key)
}
