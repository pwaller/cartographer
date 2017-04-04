package main

import (
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

func isMap(t types.Type) bool {
	_, ok := t.(*types.Map)
	return ok
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	switch stmt := node.(type) {
	case *ast.RangeStmt:
		if isMap(v.pkg.TypeOf(stmt.X)) {
			log.Println("Map at", v.program.Fset.Position(node.Pos()))
		}
	default:
	}
	return v
}
