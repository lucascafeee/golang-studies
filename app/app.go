package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Single page aplication"
	app.Usage = "Search Ips and server names on internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPS de endereços na internet",
			Flags:  flags,
			Action: fetchIPS,
		},
		{
			Name:   "servers",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: fetchSERVERS,
		},
	}

	return app
}

func fetchIPS(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
func fetchSERVERS(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host) // nameserver
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
