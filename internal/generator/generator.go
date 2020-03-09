package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"sync"
)

const packageName = "openapi"

var pool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

type Generator struct {
	name    string
	imports imports
	body    *bytes.Buffer

	ShowSource bool
}

type import_ struct {
	name string
	path string
}

type imports struct {
	stdpkg     []import_
	thirdparty []import_
}

func New(name string) *Generator {
	return &Generator{
		name: name,
		body: new(bytes.Buffer),
	}
}

func (g *Generator) Printf(format string, args ...interface{}) {
	fmt.Fprintf(g.body, format, args...)
}

func (g *Generator) Import(name, path string) {
	if strings.ContainsRune(strings.Split(path, "/")[0], '.') {
		g.imports.thirdparty = append(g.imports.thirdparty, import_{
			name: name,
			path: path,
		})
		return
	}
	g.imports.stdpkg = append(g.imports.stdpkg, import_{
		name: name,
		path: path,
	})
}

var WriteFile = ioutil.WriteFile

func (g *Generator) Save(filepath string) error {
	buf := pool.Get().(*bytes.Buffer)
	defer pool.Put(buf)
	buf.Reset()

	fmt.Fprintf(buf, "// Code generated by %s. DO NOT EDIT.", g.name)
	fmt.Fprintf(buf, "\n\npackage %s", packageName)
	fmt.Fprintf(buf, "\n%s", g.imports.String())
	fmt.Fprintf(buf, "\n%s", g.body.String())

	src, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("error on formatting source code: %v", err)
		return err
	}
	if g.ShowSource {
		printSource(src)
	}
	return WriteFile(filepath, src, 0644)
}

func (imports imports) String() string {
	if len(imports.stdpkg) == 0 && len(imports.thirdparty) == 0 {
		return ""
	}
	buf := pool.Get().(*bytes.Buffer)
	defer pool.Put(buf)
	buf.Reset()

	outf := func(format string, args ...interface{}) {
		fmt.Fprintf(buf, format, args...)
	}

	fmt.Fprintf(buf, "\nimport (")
	for i := range imports.stdpkg {
		fmt.Fprintf(buf, "\n")
		if imports.stdpkg[i].name != "" {
			fmt.Fprintf(buf, "%s ", imports.stdpkg[i].name)
		}
		fmt.Fprintf(buf, "%s", strconv.Quote(imports.stdpkg[i].path))
	}
	if len(imports.stdpkg) != 0 {
		fmt.Fprintf(buf, "\n")
	}
	for i := range imports.thirdparty {
		outf("\n")
		if imports.thirdparty[i].name != "" {
			outf("%s ", imports.thirdparty[i].name)
		}
		outf("%s", strconv.Quote(imports.thirdparty[i].path))
	}
	outf("\n)")
	return buf.String()
}

func printSource(src []byte) {
	scanner := bufio.NewScanner(bytes.NewReader(src))
	log.Print("=== SOURCE ===")
	for scanner.Scan() {
		log.Print(scanner.Text())
	}
	log.Print("=== SOURCE ===")
}
