package main

import (
	"github.com/sonjafeitsch/workshop-registration/pkg/client"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Workshop Registration"

	var baseUrl string
	var token string

	app.Flags = []cli.Flag{
		&cli.StringFlag{Name: "baseUrl", Required: true, Destination: &baseUrl},
		&cli.StringFlag{Name: "token", Required: true, Destination: &token},
	}

	app.Action = func(context *cli.Context) error {
		newClient := client.NewClient(baseUrl, token)
		_, err := newClient.GetCustomer("test@example.com")
		return err
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
