package main

import (
	"go/ast"
	"go/build"
	"go/doc"
	"go/parser"
	"go/token"
	"log"
	"os"
)

type typeOption struct {
	pkgPath string
}

func analysis(option typeOption) *doc.Package {
	pkg, err := build.Import(option.pkgPath, "", build.ImportComment)
	if err != nil {
		log.Fatalf("Import Dir error:%s", err.Error())
	}

	astPkg := parsePackage(pkg)

	docPkg := doc.New(astPkg, pkg.ImportPath, doc.AllDecls)

	return docPkg
}

func parsePackage(pkg *build.Package) *ast.Package {
	fs := token.NewFileSet()
	include := func(info os.FileInfo) bool {
		for _, name := range pkg.GoFiles {
			if name == info.Name() {
				return true
			}
		}
		for _, name := range pkg.CgoFiles {
			if name == info.Name() {
				return true
			}
		}
		return false
	}
	pkgs, err := parser.ParseDir(fs, pkg.Dir, include, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure they are all in one package.
	if len(pkgs) != 1 {
		log.Fatalf("multiple packages in directory %s", pkg.Dir)
	}
	astPkg := pkgs[pkg.Name]
	return astPkg
}
