package codeql_go_vendor_extractor

import (
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectPkgPathPost(t *testing.T) {
	pkgs, err := LoadPackage(targetDir, buildFlags, pkgPattern)
	assert.NoError(t, err)
	pkgDirs := make(map[string]string)
	wantedRoots := make(map[string]bool)
	CollectPkgPathPost(pkgs[0], pkgDirs, wantedRoots)
	log.Logger.Info(pkgDirs)
}

func TestCollectPkgPath(t *testing.T) {
	pkgs, err := LoadPackage(targetDir, buildFlags, pkgPattern)
	assert.NoError(t, err)
	pkgDirs, wantedRoots := CollectPkgPath(pkgs)
	log.Logger.Info(pkgDirs)
	log.Logger.Info(wantedRoots)
}

