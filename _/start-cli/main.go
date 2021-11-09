package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.EnableBashCompletion = true
	// app.UseShortOptionHandling = true

	app.Version = "v0.1.0"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "lang",
			Aliases: []string{"l"},
			Value:   "english",
			Usage:   "language for the greeting",
			EnvVars: []string{"APP_LANG"},
			// Required: true,
		},
		&cli.TimestampFlag{
			Name:        "meeting",
			Usage:       "timestamp for the meeting",
			Layout:      "2006-01-02T15:04:05",
			DefaultText: "now",
		},
	}

	tasks := []string{"cook", "clean", "laundry", "eat", "sleep", "code"}

	app.Commands = []*cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			// Category: "category",
			Action: func(c *cli.Context) error {
				fmt.Println("added task:", c.Args().First())
				return nil
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task:", c.Args().First())
				return nil
			},
			BashComplete: func(c *cli.Context) {
				// This will complete if no args are passed
				if c.NArg() > 0 {
					return
				}
				for _, t := range tasks {
					fmt.Println(t)
				}
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			Subcommands: []*cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template:", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template:", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "Nefertiti"
		if c.NArg() > 0 {
			name = c.Args().Get(0)
		}
		if c.String("lang") == "spanish" {
			fmt.Println("Hola", name)
		} else {
			fmt.Println("Hello", name)
		}
		timestamp := c.Timestamp("meeting")
		if timestamp != nil {
			fmt.Println(timestamp.String())
		}
		return nil
	}

	// sort.Sort(cli.FlagsByName(app.Flags))
	// sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
