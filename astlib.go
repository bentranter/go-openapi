// +build ignore

package main

import (
	"go/ast"
	"log"
	"reflect"
	"strings"
)

func ast2type(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr: // pointer
		return "*" + t.X.(*ast.Ident).Name
	case *ast.ArrayType:
		return "[]" + ast2type(t.Elt)
	case *ast.MapType:
		return "map[" + ast2type(t.Key) + "]" + ast2type(t.Value)
	case *ast.InterfaceType:
		return "interface{}"
	default:
		log.Fatalf("unknown type: %s", reflect.TypeOf(t))
	}
	return ""
}

type tags map[string][]string

func parseTags(f *ast.Field) tags {
	if f.Tag == nil {
		return nil
	}
	s := strings.Trim(f.Tag.Value, "`")
	if s == "" {
		return nil
	}
	t := tags{}
	for _, tt := range strings.Fields(s) {
		kv := strings.Split(tt, ":")
		t[kv[0]] = append(t[kv[0]], strings.Split(strings.Trim(kv[1], `"`), ",")...)
	}
	return t
}

func (t tags) get(key string) string {
	if vs, ok := t[key]; ok {
		return vs[0]
	}
	return ""
}