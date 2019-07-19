package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"

	"github.com/urfave/cli"
)

const (
	Separator     = os.PathSeparator
	ListSeparator = os.PathListSeparator
)

func main() {
	var (
		err       error
		folder    string
		inputUrl  string
		gpath     string
		clonePath string
	)

	app := cli.NewApp()
	app.Name = "goclone"
	app.Usage = "Create folder path based on url"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "url, u",
			Usage:       "Github url",
			Destination: &inputUrl,
		},
		cli.StringFlag{
			Name:        "folder, f",
			Usage:       "Destination location",
			Value:       "",
			Destination: &folder,
		},
	}

	app.Action = func(c *cli.Context) error {
		u, err := url.Parse(inputUrl)
		if err != nil {
			panic(err)
		}
		if folder == "go" {
			gpath = os.Getenv("GOPATH")
			clonePath = fmt.Sprintf("%s/src/%s%s", gpath, u.Host, path.Dir(u.Path))
			err := os.MkdirAll(clonePath, os.ModePerm)
			if err != nil {
				fmt.Printf("Error: %s", err)
			}

			doClone := fmt.Sprintf("cd %s && git clone %s", clonePath, inputUrl)
			cmdCD := exec.Command("/bin/sh", "-c", doClone)
			out, err := cmdCD.Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", out)

		}
		if folder == "git" {
			gpath = os.Getenv("GITPATH")
			clonePath = fmt.Sprintf("%s/", gpath)
			doClone := fmt.Sprintf("cd %s && git clone %s", clonePath, inputUrl)
			cmdCD := exec.Command("/bin/sh", "-c", doClone)
			out, err := cmdCD.Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", out)

		}
		return nil
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
