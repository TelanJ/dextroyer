package command

import (
	"fmt"
	"os"
	"net/http"
	"github.com/codegangsta/cli"
)

func NewGetCommand() cli.Command {
	return cli.Command{
		Name:  "get",
		Usage: "retrieve the value of a key",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:   "consul, c",
				Usage:  "use consul",
				EnvVar: "",
			},
			cli.BoolFlag{
				Name:   "etcd, e",
				Usage:  "use etcd",
				EnvVar: "",
			},
		},
		Action: func(c *cli.Context) {
			client, _ := makeClient(c)
			if len(c.Args()) == 0 {
				fmt.Fprintln(os.Stderr, "Error: Key required")
				os.Exit(1)
			}

			key := c.Args()[0]

			client.GetKey(key,consuloretcd.KeyOptions{})
			fmt.Println(client.GetKey(key,consuloretcd.KeyOptions{}))
		},
	}
}