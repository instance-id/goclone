package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	var (
		err       error
		folder    string
		inputUrl  string
		gpath     string
		clonePath string
		openApp   string
		debug     bool
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
		cli.StringFlag{
			Name:        "open, o",
			Usage:       "Open folder in application",
			Value:       "none",
			Destination: &openApp,
		},
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Debug flag to print data from input path",
			Destination: &debug,
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
			if debug {
				fullPath := strings.Replace(path.Base(u.Path), ".git", "", -1)
				fmt.Printf("%s\n", fullPath)
				fmt.Printf("%s/%s\n", clonePath, fullPath)
			} else {
				err := os.MkdirAll(clonePath, os.ModePerm)
				if err != nil {
					fmt.Printf("Error: %s", err)
				}

				var doClone string
				if openApp == "none" {
					doClone = fmt.Sprintf("cd %s && git clone %s", clonePath, inputUrl)
				} else {
					fileName := strings.Replace(path.Base(u.Path), ".git", "", -1)
					fullPath := fmt.Sprintf("%s/%s\n", clonePath, fileName)
					doClone = fmt.Sprintf("cd %s && git clone %s && %s %s", clonePath, inputUrl, openApp, fullPath)
				}

				cmdCD := exec.Command("/bin/sh", "-c", doClone)
				out, err := cmdCD.Output()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s", out)
			}
		}
		if folder == "git" {
			gpath = os.Getenv("GITPATH")
			clonePath = fmt.Sprintf("%s/", gpath)
			if debug {
				fmt.Printf(clonePath)
			} else {
				var doClone string
				if openApp == "none" {
					doClone = fmt.Sprintf("cd %s && git clone %s", clonePath, inputUrl)
				} else {
					fileName := strings.Replace(path.Base(u.Path), ".git", "", -1)
					fullPath := fmt.Sprintf("%s/%s\n", clonePath, fileName)
					doClone = fmt.Sprintf("cd %s && git clone %s && %s %s", clonePath, inputUrl, openApp, fullPath)
				}

				cmdCD := exec.Command("/bin/sh", "-c", doClone)
				out, err := cmdCD.Output()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%s", out)
			}
		}
		return nil
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
