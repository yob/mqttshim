package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/yob/mqttshim"
)

func logf(msg string, args ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, msg+"\n", args...)

	if err != nil {
		panic("failed to write to stderr: " + err.Error())
	}
}

func main() {
	app := cli.NewApp()

	app.Name = "mqttshim"
	app.Usage = "log mqtt messages to stdout"
	app.Version = "dev"

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "start the mqttshim server",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "bind, b",
					Usage: "IP address to listen on",
					Value: "0.0.0.0",
				},
				cli.IntFlag{
					Name:  "port, p",
					Usage: "the port to listen on",
					Value: 1883,
				},
			},
			Action: func(ctx *cli.Context) error {
				server := mqttshim.NewMqttShim()
				return server.StartServer(ctx.String("bind"), ctx.Int("port"))
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logf("%+v", err)
		os.Exit(1)
	}
}
