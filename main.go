package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tomassantunes/port-scanner/port"
	"github.com/urfave/cli"
)

func main() {
	var ip string
	var pc string // protocol
	var p int     // port

	app := &cli.App{
		Name:   "Port Scanner",
		Usage:  "Scan for open ports",
		Author: "Tomás Antunes - @tomassantunes",
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
			Value:       "tcp, udp",
			Usage:       "Protocol for IP(s) scan, e.g --pc tcp",
			Destination: &pc,
		},
		cli.IntFlag{
			Name:        "port, p",
			Usage:       "Port to scan, e.g --p 200",
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

	if len(pc) > 0 && p > 0 {
		open := port.ScanPort(pc, ip, p)
		fmt.Printf("Port %s/%d: %s\n", pc, p, open.State)
	}

	results := port.InitialScan("localhost")
	fmt.Println(results)
}
