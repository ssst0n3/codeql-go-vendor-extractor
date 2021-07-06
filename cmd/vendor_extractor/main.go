package main

import (
	"github.com/ssst0n3/awesome_libs/log"
	codeql_go_vendor_extractor "github.com/ssst0n3/codeql-go-vendor-extractor"
	"github.com/urfave/cli/v2"
	"os"
)

const usage = `A codeql extractor for go project using pure vendor mode.

`

func main() {
	app := &cli.App{
		Name:  "vendor_extractor",
		Usage: usage,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "package",
				Value: ".",
				Usage: "package to be extract",
			},
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
		Action: func(context *cli.Context) (err error) {
			pkg := context.String("package")
			pkgs, err := codeql_go_vendor_extractor.LoadPackage(
				"", nil,
				[]string{pkg},
			)
			if err != nil {
				return
			}
			if len(pkgs) == 0 {
				log.Logger.Info("No packages found.")
			}
			log.Logger.Info("Extracting universe scope.")
			codeql_go_vendor_extractor.ExtractUniverseScope()
			log.Logger.Info("Done extracting universe scope.")
			//codeql_go_vendor_extractor.CollectPkgPath(pkgs)
			codeql_go_vendor_extractor.ExtractType(pkgs)
			log.Logger.Info("Done processing dependencies.")
			codeql_go_vendor_extractor.ExtractPackages(pkgs)
			return
		},
		Before: func(context *cli.Context) (err error) {
			//debug := context.Bool("debug")
			//if !debug {
			//	log.Logger.SetOutput(ioutil.Discard)
			//} else {
			//	log.Logger.Info("debug mode on")
			//}
			return
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Logger.Fatal(err)
	}
}
