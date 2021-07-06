package codeql_go_vendor_extractor

import (
	"github.com/github/codeql-go/extractor/dbscheme"
	"github.com/github/codeql-go/extractor/trap"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"golang.org/x/tools/go/packages"
)

func ExtractTypePre(pkg *packages.Package) bool {
	return true
}

func ExtractTypePost(pkg *packages.Package) {
	log.Logger.Infof("Extracting types for package %s.", pkg.PkgPath)

	tw, err := trap.NewWriter(pkg.PkgPath, pkg)
	awesome_error.CheckFatal(err)
	defer tw.Close()

	scope := extractPackageScope(tw, pkg)
	tw.ForEachObject(extractObjectType)
	lbl := tw.Labeler.GlobalID(pkg.PkgPath + ";pkg")
	dbscheme.PackagesTable.Emit(tw, lbl, pkg.Name, pkg.PkgPath, scope)

	if len(pkg.Errors) != 0 {
		log.Logger.Warningf("encountered errors extracting package `%s`:", pkg.PkgPath)
		for i, err := range pkg.Errors {
			log.Logger.Warnf("  %s", err.Error())
			extractError(tw, err, lbl, i)
		}
	}
	log.Logger.Infof("Done extracting types for package %s.", pkg.PkgPath)
}

func ExtractType(pkgs []*packages.Package) {
	packages.Visit(pkgs, ExtractTypePre, ExtractTypePost)
}

func extractPackagePre(pkg *packages.Package) bool {
	return true
}

func extractPackagePost(pkg *packages.Package) {
	fdSem := newSemaphore(100)
	for _, astFile := range pkg.Syntax {
		err := extractFile(astFile, pkg, fdSem)
		awesome_error.CheckFatal(err)
	}
}

func ExtractPackages(pkgs []*packages.Package) {
	log.Logger.Info("Starting to extract packages.")
	packages.Visit(pkgs, extractPackagePre, extractPackagePost)
	log.Logger.Info("Done extracting packages.")
}
