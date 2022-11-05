package main

import (
	"bufio"
	"fmt"
	"instance"

	// "fmt"
	
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	// "github.com/mdp/qrterminal/v3"

	"go.mau.fi/whatsmeow"
	waBinary "go.mau.fi/whatsmeow/binary"
	waLog "go.mau.fi/whatsmeow/util/log"
	// waLog "go.mau.fi/whatsmeow/util/log"
)


var clis []*whatsmeow.Client

var cli *whatsmeow.Client

var log waLog.Logger

func main() {
	waBinary.IndentXML = true
	

	fmt.Println("Oláaaaa")

	c := make(chan os.Signal)
	input := make(chan string)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		defer close(input)
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			line := strings.TrimSpace(scan.Text())
			if len(line) > 0 {
				input <- line
			}
		}
	}()
	for {
		select {
		case <-c:
			log.Infof("Interrupt received, exiting")
			// cli.Disconnect()
			return
		case cmd := <-input:

			if len(cmd) == 0 {
				log.Infof("Stdin closed, exiting")
				//cli.Disconnect()
				return
			}
			
			args := strings.Fields(cmd)
			cmd = args[0]
			args = args[1:]

			if cmd == "start" {
				fmt.Println("Passsou aqui")
				//go instance()
				cli = instance.Instance(args[0])
				if cli == nil {

					clis = append(clis, cli) 
				}
			}

			// go handleCmd(strings.ToLower(cmd), args)
		}
	}
}