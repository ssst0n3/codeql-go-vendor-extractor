package codeql_go_vendor_extractor

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadPackage(t *testing.T) {
	assert.NoError(t, os.Setenv("GO111MODULE", "off"))
	gopath := os.Getenv("GOPATH")
	assert.True(t, len(gopath) > 0)
	targetDir := gopath + "/src/github.com/docker/docker"
	buildFlags := []string{
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
	pkgs, err := LoadPackage(targetDir, buildFlags, []string{"github.com/docker/docker/cmd/dockerd"})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(pkgs))
	assert.Equal(t, "github.com/docker/docker/cmd/dockerd", pkgs[0].PkgPath)
}
