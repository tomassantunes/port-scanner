package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/tomassantunes/port-scanner/port"
	"github.com/urfave/cli"
)

func main() {
	var ip string
	var pc string // protocol
	var p string  // port

	app := &cli.App{
		Name:   "Port Scanner",
		Usage:  "Scan for open ports",
		Author: "Tomás Antunes - github.com/tomassantunes",
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "ip",
			Usage:       "e.g --ip 127.0.0.1/12, --ip localhost",
			Required:    true,
			Destination: &ip,
		},
		cli.StringFlag{
			Name:        "protocol, pc",
			Value:       "tcp",
			Usage:       "Protocol for IP(s) scan, e.g --pc tcp",
			Destination: &pc,
		},
		cli.StringFlag{
			Name:        "port, p",
			Usage:       "Ports to scan, e.g --p 80,120,200",
			Destination: &p,
		},
		cli.StringFlag{
			Name:  "timeout, t",
			Value: "2000",
			Usage: "Timeout in milliseconds, e.g --t 3000",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	if len(pc) > 0 && len(p) > 0 {
		ports := strings.Split(p, ",")

		for i := 0; i < len(ports); i++ {
			portNum, errP := strconv.Atoi(ports[i])

			if errP != nil {
				break
			}

			open := port.ScanPort(pc, ip, portNum)
			fmt.Printf("Port %s/%d: %s\n", pc, portNum, open.State)
		}
	} else {
		results := port.InitialScan("localhost")
		fmt.Println(results)
	}

}
