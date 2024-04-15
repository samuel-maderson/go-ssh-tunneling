package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"strconv"

	"go-ssh-tunneling/src/types"
)

var (
	custom_config = types.HostJSON{}
)

func init() {

	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(data, &custom_config)
}

func StartTunneling(rport int, rhost string, lhost string, lport int, key string, user string, bastionhost string) {

	cmd := exec.Command("ssh", "-f", "-N", "-L", strconv.Itoa(lport)+":"+rhost+":"+strconv.Itoa(rport), "-i", key, user+"@"+bastionhost)

	err := cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}

}

func main() {

	for _, host := range custom_config.Hosts {
		log.Print("\033[1;32m[+]\033[0m Starting Tunneling for host: ", host.Remote.Host, ":", host.Remote.Port)
		log.Print("\033[1;32m[+]\033[0m Listening on host: ", host.Local.Host, ":", host.Local.Port)
		log.Print()
		StartTunneling(host.Remote.Port, host.Remote.Host, host.Local.Host, host.Local.Port, host.Key, host.User, host.Bastion.Host)
	}
}
