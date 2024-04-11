package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"go-ssh-tunneling/src/types"

	"github.com/alexflint/go-arg"
)

var (
	rhost       string
	rport       int
	lhost       string
	lport       int
	bastionhost string
	key         string
	user        string
	args        = &types.Args{}
)

func init() {

	arg.MustParse(args)
	rport = 3306
	rhost = args.Rhost
	lhost = "localhost"
	lport = 53306
	key = args.Key
	user = args.User
	bastionhost = "172.31.75.252"

	if args.Key == "" && args.User == "" && args.Rhost == "" {
		log.Fatalln("Key or User not provided")
	}

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
