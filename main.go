package main

import (
	"errors"
	"fmt"
	"git.hubteam.com/zklapow/singularity-cli/client"
	"git.hubteam.com/zklapow/singularity-cli/commands"
	"gopkg.in/urfave/cli.v2"
	"gopkg.in/urfave/cli.v2/altsrc"
	"os"
	"strconv"
)

func main() {
	conf := Config{}

	flags := []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "base-uri",
			Destination: &conf.BaseUri,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:        "user",
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
				Category:  "requests",
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "list all requests",
				ArgsUsage: "[request]",
				Action: func(c *cli.Context) error {
					if c.Args().Get(0) != "" {
						commands.FindRequestsMatching(conf.getClient(), c.Args().Get(0))
						return nil
					}

					commands.ListAllRequests(conf.getClient())
					return nil
				},
				BashComplete: completeFromCachedRequestList(&conf),
			},
			{
				Category:  "requests",
				Name:      "show",
				Usage:     "show details of a request",
				ArgsUsage: "[request]",
				Action: func(c *cli.Context) error {
					commands.ShowRequestDetails(conf.getClient(), c.Args().Get(0))
					return nil
				},
				BashComplete: completeFromCachedRequestList(&conf),
			},
			{
				Category:  "requests",
				Name:      "pause",
				Usage:     "pause a request",
				ArgsUsage: "request",
				Before: func(c *cli.Context) error {
					if c.Args().Get(0) == "" {
						return errors.New("Error: Must specify a request to pause")
					}

					return nil
				},
				Action: func(c *cli.Context) error {
					commands.PauseRequest(conf.getClient(), c.Args().Get(0))
					return nil
				},
				BashComplete: completeFromCachedRequestList(&conf),
			},
			{
				Category:  "requests",
				Name:      "unpause",
				Usage:     "unpause a request",
				ArgsUsage: "request",
				Before: func(c *cli.Context) error {
					if c.Args().Get(0) == "" {
						return errors.New("Error: Must specify a request to un-pause")
					}

					return nil
				},
				Action: func(c *cli.Context) error {
					commands.UnPauseRequest(conf.getClient(), c.Args().Get(0))
					return nil
				},
				BashComplete: completeFromCachedRequestList(&conf),
			},
			{
				Category:  "requests",
				Name:      "scale",
				Usage:     "scale a request",
				ArgsUsage: "request instanceCount",
				Before: func(c *cli.Context) error {
					if c.Args().Get(0) == "" || c.Args().Get(1) == "" {
						return errors.New("Error: Please specify a request to scale and instance count to scale to")
					}

					if _, err := strconv.Atoi(c.Args().Get(1)); err != nil {
						return errors.New("Error: Instance count must be a number")
					}

					return nil
				},
				Action: func(c *cli.Context) error {
					instanceCount, _ := strconv.Atoi(c.Args().Get(1))

					commands.ScaleRequest(conf.getClient(), c.Args().Get(0), instanceCount)
					return nil
				},
				BashComplete: completeFromCachedRequestList(&conf),
			},
		},
	}

	app.Run(os.Args)
}

type Config struct {
	BaseUri string
	User    string
}

func (c *Config) getClient() *client.SingularityClient {
	return client.NewSingularityClient(c.BaseUri, map[string]string{"X-HubSpot-User": c.User})
}

func completeFromCachedRequestList(conf *Config) cli.BashCompleteFunc {
	return func(c *cli.Context) {
		requests, err := conf.getClient().GetCachedRequestList()
		if err != nil {
			return
		}

		for _, req := range requests {
			fmt.Println(req)
		}

		return
	}
}
