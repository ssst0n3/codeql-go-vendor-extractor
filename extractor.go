package codeql_go_vendor_extractor

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"golang.org/x/tools/go/packages"
)

func LoadPackage(targetDir string, buildFlags []string, patterns []string) (pkgs []*packages.Package, err error) {
	log.Logger.Info("Running packages.Load.")
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedDeps |
			packages.NeedTypes | packages.NeedTypesSizes |
			packages.NeedTypesInfo | packages.NeedSyntax,
		BuildFlags: buildFlags,
		Dir:        targetDir,
		// TODO: env
		//Env: []string{},
	}
	pkgs, err = packages.Load(cfg, patterns...)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	log.Logger.Info("Done running packages.Load.")
	return
}
