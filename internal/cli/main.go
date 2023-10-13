package cli

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/moabukar/cia/internal/scanner"
	"github.com/urfave/cli"
)

func NewApp(version string, start time.Time) *cli.App {
	app := cli.NewApp()
	app.Name = "cia"
	app.Version = version
	app.Usage = " CIA is your go-to CLI tool for analyzing container images. It can pull images, scan for vulnerabilities, and output reports in multiple formats."
	app.EnableBashCompletion = true

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "format",
			Usage: "Report format (json, xml)",
			Value: "json",
		},
		&cli.BoolFlag{
			Name:  "skip-pull",
			Usage: "Skip pulling image before scanning",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "scan",
			Usage: "Scans the given container image",
			Action: func(c *cli.Context) error {
				imageName := c.Args().Get(0)
				if imageName == "" {
					log.Fatalf("Image name is required")
					return fmt.Errorf("image name is required")
				}

				options := scanner.ScanOptions{
					Format:   c.String("format"),
					SkipPull: c.Bool("skip-pull"),
				}

				err := scanner.ScanImage(imageName, options)
				if err != nil {
					log.Fatalf("Error scanning image: %v", err)
					return fmt.Errorf("error scanning image: %v", err)
				}

				return nil
			},
		},
		{
			Name:  "report",
			Usage: "Generate a report of the last scan",
			Action: func(c *cli.Context) error {
				// Your logic for generating a report goes here.
				return nil
			},
		},
		{
			Name:  "version",
			Usage: "Show the version of CIA tool",
			Action: func(c *cli.Context) error {
				fmt.Println("Version:", app.Version)
				return nil
			},
		},
	}

	return app
}

// Run handles the instanciation of the CLI application
func Execute(version string, args []string) {
	err := NewApp(version, time.Now()).Run(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
