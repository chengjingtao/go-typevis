package main

import (
	"go/doc"
	"log"
	"os"
	"text/template"
)

var graphTemplate = `digraph G {
	fontname = "Bitstream Vera Sans"
	fontsize = 8

	node [
			fontname = "Bitstream Vera Sans"
			fontsize = 8
			shape = "record"
	]

	edge [
			fontname = "Bitstream Vera Sans"
			fontsize = 8
	]

	subgraph cluster_{{.Name}}{
		label = "Package {{.Name}}"
		{{range $i, $entity := .Types}}

		{{$entity.Name}} [
		label = "{
			{{$entity.Name}} |
			{{- range $j, $f := $entity.Vars}}
				{{- range $k, $n := $f.Names}}
			{{$n.Name}}\l
				{{- end}}
			{{- end}}|
			{{- range $j, $f := $entity.Funcs}}
			{{$f.Name}}()\l
			{{- end}}|
			{{range $j, $f := $entity.Methods}}
			{{$f.Name}}()\l
			{{- end}}
		}"
		]

		{{- end}}
	}

}
`

func render(pkg *doc.Package) {

	tmpl, err := template.New("gotype-vis").Parse(graphTemplate)
	if err != nil {
		log.Fatal("Parse template error:%s", err.Error())
	}

	// pkg.Types[0].Methods[0].Orig
	// pkg.Types[0].Vars[0].Names
	tmpl.Execute(os.Stdout, pkg)
}
