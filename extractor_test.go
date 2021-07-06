package codeql_go_vendor_extractor

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	awesome_error.CheckFatal(os.Setenv("GO111MODULE", "off"))
	awesome_error.CheckFatal(os.Setenv("CODEQL_EXTRACTOR_GO_TRAP_DIR", "/tmp/ql/trap/go"))
	awesome_error.CheckFatal(os.Setenv("CODEQL_EXTRACTOR_GO_SOURCE_ARCHIVE_DIR", "/tmp/ql/src"))
}

var (
	gopath     = os.Getenv("GOPATH")
	targetDir  = gopath + "/src/github.com/docker/docker"
	buildFlags = []string{
		"-tags", "netgo osusergo static_build apparmor seccomp journald",
		"-installsuffix", "netgo",
		"-buildmode=pie",
		`-ldflags`, `
-w
-X "github.com/docker/docker/dockerversion.Version=dev" 
-X "github.com/docker/docker/dockerversion.GitCommit=8891c58a43"
-X "github.com/docker/docker/dockerversion.BuildTime=2021-07-05T02:59:10.000000000+00:00"
-X "github.com/docker/docker/dockerversion.IAmStatic=true"
-X "github.com/docker/docker/dockerversion.PlatformName="
-X "github.com/docker/docker/dockerversion.ProductName="
-X "github.com/docker/docker/dockerversion.DefaultProductLicense="
-X "github.com/docker/docker/dockerversion.InitCommitID=de40ad007797e0dcd8b7126f27bb87401d224240"
-extldflags "-static"
`}
	pkgName    = "github.com/docker/docker/cmd/dockerd"
	pkgPattern = []string{pkgName}
)

func TestLoadPackage(t *testing.T) {
	assert.True(t, len(gopath) > 0)
	pkgs, err := LoadPackage(targetDir, buildFlags, pkgPattern)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(pkgs))
	assert.Equal(t, pkgName, pkgs[0].PkgPath)
}

func TestExtractPackage(t *testing.T) {
	assert.True(t, len(gopath) > 0)
	pkgs, err := LoadPackage(targetDir, buildFlags, pkgPattern)
	assert.NoError(t, err)
	ExtractPackage(pkgs[0])
}
