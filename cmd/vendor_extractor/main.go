package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	codeql_go_vendor_extractor "github.com/ssst0n3/codeql-go-vendor-extractor"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
)

const usage = `A codeql extractor for go project using pure vendor mode.

`

func main() {
	app := &cli.App{
		Name:  "vendor_extractor",
		Usage: usage,
		Commands: []*cli.Command{

		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the using",
			},
			&cli.BoolFlag{
				Name:  "debug",
				Value: false,
				Usage: "Output information for helping debugging vendor_extractor",
			},
		},
		Before: func(context *cli.Context) (err error) {
			debug := context.Bool("debug")
			if !debug {
				log.Logger.SetOutput(ioutil.Discard)
			} else {
				log.Logger.Info("debug mode on")
			}
			pkgs, err := codeql_go_vendor_extractor.LoadPackage("", nil, []string{"github.com/docker/docker/cmd/dockerd"})
			awesome_error.CheckFatal(err)
			spew.Dump(pkgs)
			codeql_go_vendor_extractor.ExtractUniverseScope()
			return
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Logger.Fatal(err)
	}
}
