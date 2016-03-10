package main

import (
	"github.com/Telanj/dextroyer/command"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "dextroyer"
	app.Version = "0.0.1"
	app.Author = "Julio Telan"
	app.Email = "telanjulio@gmail.com"
	app.Usage = "A simple command line client for issuing http calls for dex."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "address, A",
			Value: "192.168.200.90",
			Usage:  "the remote endpoint for the cluster",
			EnvVar: "CONSULORETCD_KV_ADDRESS",
		},
		cli.IntFlag{
			Name:   "port, p",
			Value: 8500,
			Usage:  "the port of the remote endpoint for the cluster",
			EnvVar: "CONSULORETCD_KV_PORT",
		},
		cli.StringFlag{
			Name:   "ttl, t",
			Value: "",
			Usage:  "the time to live value of that is expected",
			EnvVar: "CONSULORETCD_KV_TTL",
		},
	}
	app.Commands = []cli.Command{
		command.NewGetCommand(),
		command.NewPutCommand(),
		command.NewDeleteCommand(),
		command.NewCASCommand(),
	}

	app.Run(os.Args)
}