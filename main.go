package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mattn/go-pipeline"
	"github.com/urfave/cli"
)

func main() {
	var all bool
	app := cli.NewApp()
	app.Name = "docker-selector"
	app.Usage = "docker container selector."

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "all, a",
			Usage:       "Show all containers",
			Destination: &all,
		},
	}

	app.Action = func(c *cli.Context) error {
		var ps string
		if all {
			ps = execDockerPecoAll()
		} else {
			ps = execDockerPeco()
		}
		id := extractID(ps)
		fmt.Print(id)
		return nil
	}

	app.Run(os.Args)
}

func execDockerPeco() string {
	out, err := pipeline.Output(
		[]string{"docker", "ps", "--format", "{{.ID}}\t{{.Image}}\t{{.Status}}\t{{.Command}}\t{{.RunningFor}}"},
		[]string{"peco"},
	)
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func execDockerPecoAll() string {
	out, err := pipeline.Output(
		[]string{"docker", "ps", "-a", "--format", "{{.ID}}\t{{.Image}}\t{{.Status}}\t{{.Command}}\t{{.RunningFor}}"},
		[]string{"peco"},
	)
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func extractID(processes string) string {
	lines := strings.Split(processes, "\n")
	ids := make([]string, len(lines))
	for i, l := range lines {
		pos := strings.Index(l, "\t")
		if pos > 0 {
			ids[i] = l[0:pos]
		}
	}
	return strings.Join(ids, " ")
}
