package main

import (
	"gopkg.in/urfave/cli.v2"
	"gopkg.in/urfave/cli.v2/altsrc"
	"git.hubteam.com/zklapow/singularity-cli/commands"
	"git.hubteam.com/zklapow/singularity-cli/client"
	"os"
	"fmt"
)

func main() {
	conf := Config{}

	flags := []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name: "base-uri",
			Destination: &conf.BaseUri,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name: "user",
			Destination: &conf.User,
		}),
	}

	app := &cli.App{
		EnableBashCompletion: true,
		Before: altsrc.InitInputSourceWithContext(flags, func(context *cli.Context) (altsrc.InputSourceContext, error) {
			source, err := altsrc.NewTomlSourceFromFile("/Users/zklapow/.sng/config.toml")
			if err != nil {
				fmt.Printf("Failed to load config from file %#v", err)
			}

			return source, nil
		}),

		Flags: flags,

		Commands: []*cli.Command{
			{
				Category: "requests",
				Name: "list",
				Aliases: []string{"l"},
				Usage: "list all requests",
				ArgsUsage: "[request]",
				Action: func(c *cli.Context) error {
					if c.Args().Get(0) != "" {
						commands.FindRequestsMatching(conf.getClient(), c.Args().Get(0))
						return nil
					}

					commands.ListAllRequests(conf.getClient())
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}

type Config struct {
	BaseUri string
	User string
}

func (c *Config) getClient() (*client.SingularityClient)  {
	return client.NewSingularityClient(c.BaseUri, map[string]string{"X-HubSpot-User": c.User})
}
