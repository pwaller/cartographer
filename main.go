package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"log"
	"os"

	"golang.org/x/tools/go/loader"
)

func main() {
	var conf loader.Config
	_, err := conf.FromArgs(os.Args[1:], false)
	if err != nil {
		log.Fatal(err)
	}
	program, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}
	for _, createdPackage := range program.Imported {
		walkPackage(program, createdPackage)
	}
	for _, createdPackage := range program.Created {
		walkPackage(program, createdPackage)
	}
}

func walkPackage(program *loader.Program, info *loader.PackageInfo) {
	for _, file := range info.Files {
		v := &visitor{
			program,
			info,
		}
		ast.Walk(v, file)
	}
}

type visitor struct {
	program *loader.Program
	pkg     *loader.PackageInfo
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	if stmt, isRange := node.(*ast.RangeStmt); isRange {
		if _, isMap := v.pkg.TypeOf(stmt.X).(*types.Map); isMap {
			fmt.Println(v.program.Fset.Position(node.Pos()))
		}
	}
	return v
}
