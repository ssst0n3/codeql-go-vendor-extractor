package codeql_go_vendor_extractor

import (
	_ "github.com/github/codeql-go/extractor"
	"github.com/github/codeql-go/extractor/trap"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	_ "unsafe"
)

/*
borrowed from https://github.com/github/codeql-go/blob/v1.27.0/extractor/extractor.go#L213
*/

//go:linkname ExtractUniverseScope github.com/github/codeql-go/extractor.extractUniverseScope
func ExtractUniverseScope()

//go:linkname extractPackageScope github.com/github/codeql-go/extractor.extractPackageScope
func extractPackageScope(tw *trap.Writer, pkg *packages.Package) trap.Label

//go:linkname extractObjectType github.com/github/codeql-go/extractor.extractObjectType
func extractObjectType(tw *trap.Writer, obj types.Object, lbl trap.Label)

//go:linkname extractError github.com/github/codeql-go/extractor.extractError
func extractError(tw *trap.Writer, err packages.Error, pkglbl trap.Label, idx int)

//go:linkname semaphore github.com/github/codeql-go/extractor.semaphore
type semaphore struct {}

//go:linkname newSemaphore github.com/github/codeql-go/extractor.newSemaphore
func newSemaphore(max int) *semaphore

//go:linkname extractFile github.com/github/codeql-go/extractor.extractFile
func extractFile(ast *ast.File, pkg *packages.Package, fdSem *semaphore) error