package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Host Lookup CLI",
		Usage: "Query IPs, CNAMEs, MX Records and Name Servers",
	}

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "google.com",
			Usage: "Specify the host to look up",
		},
	}

	app.Commands = []*cli.Command{
		{
			Name:  "ns",
			Usage: "Look up host name servers",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for _, n := range ns {
					fmt.Println(n.Host)
				}
				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Look up host IP addresses",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for _, i := range ip {
					fmt.Println(i)
				}
				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Look up host CNAME",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Look up host MX records",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					fmt.Println(err)
					return err
				}
				for _, m := range mx {
					fmt.Println(m.Host, m.Pref)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}