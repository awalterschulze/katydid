//  Copyright 2013 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

//Command funcs-gen generates some of the code in the funcs package.
package main

import (
	"strings"

	"github.com/katydid/katydid/gen"
)

const compareStr = `
type {{.Type}}{{.CName}} struct {
	V1 {{.CType}}
	V2 {{.CType}}
}

func (this *{{.Type}}{{.CName}}) Eval() (bool, error) {
	v1, err := this.V1.Eval()
	if err != nil {
		return {{.Error}}, nil
	}
	v2, err := this.V2.Eval()
	if err != nil {
		return {{.Error}}, nil
	}
	{{if .Eval}}{{.Eval}}{{else}}return v1 {{.Operator}} v2, nil{{end}}
}

func init() {
	Register("{{.Name}}", new({{.Type}}{{.CName}}))
}

// {{.CType}}{{.CName}} returns a new {{.Comment}} function.
func {{.CType}}{{.CName}}(a, b {{.CType}}) Bool {
	return &{{.Type}}{{.CName}}{V1: a, V2: b}
}
`

type compare struct {
	Name     string
	Operator string
	Type     string
	Eval     string
	CType    string
	Error    string
	Comment  string
}

func (this *compare) CName() string {
	if this.Name == "ge" || this.Name == "le" {
		return strings.ToUpper(this.Name)
	}
	return gen.CapFirst(this.Name)
}

const newFuncStr = `
//New{{.}}Func dynamically creates and asserts the returning function is of type {{.}}.
//This function is used by the compose library to compile functions together.
func New{{.}}Func(uniq string, values ...interface{}) ({{.}}, error) {
	f, err := newFunc(uniq, values...)
	if err != nil {
		return nil, err
	}
	return f.({{.}}), nil
}
`

const constStr = `
type Const{{.CType}} interface {
	{{.CType}}
}

var typConst{{.CType}} reflect.Type = reflect.TypeOf((*Const{{.CType}})(nil)).Elem()

type const{{.CType}} struct {
	v {{.GoType}}
}

//{{.CType}}Const returns a new constant function of type {{.CType}}
func {{.CType}}Const(v {{.GoType}}) Const{{.CType}} {
	return &const{{.CType}}{v}
}

func (this *const{{.CType}}) IsConst() {}

func (this *const{{.CType}}) Eval() ({{.GoType}}, error) {
	return this.v, nil
}

func (this *const{{.CType}}) String() string {
	{{if .ListType}}ss := make([]string, len(this.v))
	for i := range this.v {
		ss[i] = fmt.Sprintf("{{.String}}", this.v[i])
	}
	return "[]{{.ListType}}{" + strings.Join(ss, ",") + "}"{{else}}return fmt.Sprintf("{{.String}}", this.v){{end}}
}
`

type conster struct {
	CType    string
	GoType   string
	String   string
	ListType string
}

const listStr = `
type listOf{{.FuncType}} struct {
	List []{{.FuncType}}
}

//NewListOf{{.FuncType}} returns a new function that when evaluated returns a list of type {{.FuncType}}
func NewListOf{{.FuncType}}(v []{{.FuncType}}) {{.CType}} {
	return &listOf{{.FuncType}}{v}
}

func (this *listOf{{.FuncType}}) Eval() ([]{{.GoType}}, error) {
	res := make([]{{.GoType}}, len(this.List))
	var err error
	for i, e := range this.List {
		res[i], err = e.Eval()
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (this *listOf{{.FuncType}}) String() string {
	ss := make([]string, len(this.List))
	for i := range this.List {
		ss[i] = sprint(this.List[i])
	}
	return "[]{{.Type}}{" + strings.Join(ss, ",") + "}"
}

func (this *listOf{{.FuncType}}) IsListOf() {}
`

type list struct {
	Type     string
	CType    string
	FuncType string
	GoType   string
}

const printStr = `
type print{{.Name}} struct {
	E {{.Name}}
}

func (this *print{{.Name}}) Eval() ({{.GoType}}, error) {
	v, err := this.E.Eval()
	if err != nil {
		fmt.Printf("error: %#v\n", v)
	} else {
		fmt.Printf("value: %#v\n", v)
	}
	return v, err
}

func (this *print{{.Name}}) IsVariable() {}

func init() {
	Register("print", new(print{{.Name}}))
}

//Print{{.Name}} returns a function that prints out the value of the argument function and returns its value.
func Print{{.Name}}(e {{.Name}}) {{.Name}} {
	return &print{{.Name}}{E: e}
}
`

type printer struct {
	Name   string
	GoType string
}

const lengthStr = `
type len{{.}} struct {
	E {{.}}
}

func (this *len{{.}}) Eval() (int64, error) {
	e, err := this.E.Eval()
	if err != nil {
		return 0, err
	}
	return int64(len(e)), nil
}

func init() {
	Register("length", new(len{{.}}))
}

//Len{{.}} returns a function that returns the length of a list of type {{.}}
func Len{{.}}(e {{.}}) Int {
	return &len{{.}}{E: e}
}
`

const elemStr = `
type elem{{.ListType}} struct {
	List {{.ListType}}
	Index Int
}

func (this *elem{{.ListType}}) Eval() ({{.ReturnType}}, error) {
	list, err := this.List.Eval()
	if err != nil {
		return {{.Default}}, err
	}
	index64, err := this.Index.Eval()
	if err != nil {
		return {{.Default}}, err
	}
	index := int(index64)
	if len(list) == 0 {
		return {{.Default}}, NewRangeCheckErr(index, len(list))
	}
	if index < 0 {
		index = index % len(list)	
	}
	if len(list) <= index {
		return {{.Default}}, NewRangeCheckErr(index, len(list))
	}
	return list[index], nil
}

func init() {
	Register("elem", new(elem{{.ListType}}))
}

//Elem{{.ListType}} returns a function that returns the n'th element of the list.
func Elem{{.ListType}}(list {{.ListType}}, n Int) {{.ThrowType}} {
	return &elem{{.ListType}}{
		List: list,
		Index: n,
	}
}
`

type elemer struct {
	ListType   string
	ReturnType string
	ThrowType  string
	Default    string
}

const rangeStr = `
type range{{.ListType}} struct {
	List {{.ListType}}
	First Int
	Last Int
}

func (this *range{{.ListType}}) Eval() ({{.ReturnType}}, error) {
	list, err := this.List.Eval()
	if err != nil {
		return nil, err
	}
	first64, err := this.First.Eval()
	if err != nil {
		return nil, err
	}
	first := int(first64)
	if len(list) == 0 {
		return nil, NewRangeCheckErr(first, len(list))
	}
	if first < 0 {
		first = first % len(list)	
	}
	if first > len(list) {
		return nil, NewRangeCheckErr(first, len(list))
	}
	last64, err := this.Last.Eval()
	if err != nil {
		return nil, err
	}
	last := int(last64)
	if last < 0 {
		last = last % len(list)	
	}
	if last > len(list) {
		return nil, NewRangeCheckErr(last, len(list))
	}
	if first > last {
		return nil, NewRangeErr(first, last)
	}
	return list[first:last], nil
}

func init() {
	Register("range", new(range{{.ListType}}))
}

//Range{{.ListType}} returns a function that returns a range of elements from a list.
func Range{{.ListType}}(list {{.ListType}}, from, to Int) {{.ListType}} {
	return &range{{.ListType}}{
		List: list,
		First: from,
		Last: to,
	}
}
`

type ranger struct {
	ListType   string
	ReturnType string
}

const variableStr = `
type var{{.Name}} struct {
	Value parser.Value
}

var _ Setter = &var{{.Name}}{}
var _ Variable = &var{{.Name}}{}

func (this *var{{.Name}}) Eval() ({{.GoType}}, error) {
	v, err := this.Value.{{.Name}}()
	if err != nil {
		return {{.Default}}, err
	}
	return v, nil
}

func (this *var{{.Name}}) IsVariable() {}

func (this *var{{.Name}}) SetValue(v parser.Value) {
	this.Value = v
}

func (this *var{{.Name}}) String() string {
	return "${{.Decode}}"
}

//{{.Name}}Var returns a variable of type {{.Name}}
func {{.Name}}Var() *var{{.Name}} {
	return &var{{.Name}}{}
}

`

type varer struct {
	Name    string
	Decode  string
	GoType  string
	Default string
}

const typStr = `
type typ{{.Name}} struct {
	E {{.Name}}
}

func (this *typ{{.Name}}) Eval() (bool, error) {
	_, err := this.E.Eval()
	return (err == nil), nil
}

func init() {
	Register("type", new(typ{{.Name}}))
}

//Type{{.Name}} returns a function that returns true if the error returned by the argument function is nil.
func Type{{.Name}}(v {{.Name}}) Bool {
	return &typ{{.Name}}{E: v}
}
`

type typer struct {
	Name string
}

const inSetStr = `
type inSet{{.Name}} struct {
	Elem {{.Name}}
	List {{.ConstListType}}
	set map[{{.Type}}]struct{}
}

func (this *inSet{{.Name}}) Init() error {
	if this.set != nil {
		return nil
	}
	l, err := this.List.Eval()
	if err != nil {
		return err
	}
	this.set = make(map[{{.Type}}]struct{})
	for i := range l {
		this.set[l[i]] = struct{}{}
	}
	return nil
}

func (this *inSet{{.Name}}) Eval() (bool, error) {
	v, err := this.Elem.Eval()
	if err != nil {
		return false, err
	}
	_, ok := this.set[v]
	return ok, nil
}

func init() {
	Register("contains", new(inSet{{.Name}}))
}

//Contains{{.Name}} returns a function that checks whether the element is contained in the list.
func Contains{{.Name}}(element {{.Name}}, list {{.ConstListType}}) Bool {
	return &inSet{{.Name}}{element, list, nil}
}
`

type inSeter struct {
	Name          string
	ConstListType string
	Type          string
}

func main() {
	gen := gen.NewPackage("funcs")
	gen(compareStr, "compare.gen.go", []interface{}{
		&compare{"ge", ">=", "double", "", "Double", "false", "greater than or equal"},
		&compare{"ge", ">=", "int", "", "Int", "false", "greater than or equal"},
		&compare{"ge", ">=", "uint", "", "Uint", "false", "greater than or equal"},
		&compare{"ge", "", "bytes", "return bytes.Compare(v1, v2) >= 0, nil", "Bytes", "false", "greater than or equal"},
		&compare{"gt", ">", "double", "", "Double", "false", "greater than"},
		&compare{"gt", ">", "int", "", "Int", "false", "greater than"},
		&compare{"gt", ">", "uint", "", "Uint", "false", "greater than"},
		&compare{"gt", "", "bytes", "return bytes.Compare(v1, v2) > 0, nil", "Bytes", "false", "greater than"},
		&compare{"le", "<=", "double", "", "Double", "false", "less than or equal"},
		&compare{"le", "<=", "int", "", "Int", "false", "less than or equal"},
		&compare{"le", "<=", "uint", "", "Uint", "false", "less than or equal"},
		&compare{"le", "", "bytes", "return bytes.Compare(v1, v2) <= 0, nil", "Bytes", "false", "less than or equal"},
		&compare{"lt", "<", "double", "", "Double", "false", "less than"},
		&compare{"lt", "<", "int", "", "Int", "false", "less than"},
		&compare{"lt", "<", "uint", "", "Uint", "false", "less than"},
		&compare{"lt", "", "bytes", "return bytes.Compare(v1, v2) < 0, nil", "Bytes", "false", "less than"},
		&compare{"eq", "==", "double", "", "Double", "false", "equal"},
		&compare{"eq", "==", "int", "", "Int", "false", "equal"},
		&compare{"eq", "==", "uint", "", "Uint", "false", "equal"},
		&compare{"eq", "==", "bool", "", "Bool", "false", "equal"},
		&compare{"eq", "==", "string", "", "String", "false", "equal"},
		&compare{"eq", "", "bytes", "return bytes.Equal(v1, v2), nil", "Bytes", "false", "equal"},
		&compare{"ne", "!=", "double", "", "Double", "false", "not equal"},
		&compare{"ne", "!=", "int", "", "Int", "false", "not equal"},
		&compare{"ne", "!=", "uint", "", "Uint", "false", "not equal"},
		&compare{"ne", "!=", "bool", "", "Bool", "false", "not equal"},
		&compare{"ne", "!=", "string", "", "String", "false", "not equal"},
		&compare{"ne", "", "bytes", "return !bytes.Equal(v1, v2), nil", "Bytes", "false", "not equal"},
	}, `"bytes"`)
	gen(newFuncStr, "newfunc.gen.go", []interface{}{
		"Double",
		"Int",
		"Uint",
		"Bool",
		"String",
		"Bytes",
		"Doubles",
		"Ints",
		"Uints",
		"Bools",
		"Strings",
		"ListOfBytes",
	})
	gen(constStr, "const.gen.go", []interface{}{
		&conster{"Double", "float64", "double(%f)", ""},
		&conster{"Int", "int64", "int(%d)", ""},
		&conster{"Uint", "uint64", "uint(%d)", ""},
		&conster{"Bool", "bool", "%v", ""},
		&conster{"String", "string", "`%s`", ""},
		&conster{"Bytes", "[]byte", "%#v", ""},
		&conster{"Doubles", "[]float64", "double(%f)", "double"},
		&conster{"Ints", "[]int64", "int(%d)", "int"},
		&conster{"Uints", "[]uint64", "uint(%d)", "uint"},
		&conster{"Bools", "[]bool", "%v", "bool"},
		&conster{"Strings", "[]string", "`%s`", "string"},
		&conster{"ListOfBytes", "[][]byte", "%#v", "[]byte"},
	}, `"fmt"`, `"strings"`, `"reflect"`)
	gen(listStr, "list.gen.go", []interface{}{
		&list{"double", "Doubles", "Double", "float64"},
		&list{"int", "Ints", "Int", "int64"},
		&list{"uint", "Uints", "Uint", "uint64"},
		&list{"bool", "Bools", "Bool", "bool"},
		&list{"string", "Strings", "String", "string"},
		&list{"[]byte", "ListOfBytes", "Bytes", "[]byte"},
	}, `"strings"`)
	gen(printStr, "print.gen.go", []interface{}{
		&printer{"Double", "float64"},
		&printer{"Int", "int64"},
		&printer{"Uint", "uint64"},
		&printer{"Bool", "bool"},
		&printer{"String", "string"},
		&printer{"Bytes", "[]byte"},
		&printer{"Doubles", "[]float64"},
		&printer{"Ints", "[]int64"},
		&printer{"Uints", "[]uint64"},
		&printer{"Bools", "[]bool"},
		&printer{"Strings", "[]string"},
		&printer{"ListOfBytes", "[][]byte"},
	}, `"fmt"`)
	gen(lengthStr, "length.gen.go", []interface{}{
		"Doubles",
		"Ints",
		"Uints",
		"Bools",
		"Strings",
		"ListOfBytes",
		"String",
		"Bytes",
	})
	gen(elemStr, "elem.gen.go", []interface{}{
		&elemer{"Doubles", "float64", "Double", "0"},
		&elemer{"Ints", "int64", "Int", "0"},
		&elemer{"Uints", "uint64", "Uint", "0"},
		&elemer{"Bools", "bool", "Bool", "false"},
		&elemer{"Strings", "string", "String", `""`},
		&elemer{"ListOfBytes", "[]byte", "Bytes", "nil"},
	})
	gen(rangeStr, "range.gen.go", []interface{}{
		&ranger{"Doubles", "[]float64"},
		&ranger{"Ints", "[]int64"},
		&ranger{"Uints", "[]uint64"},
		&ranger{"Bools", "[]bool"},
		&ranger{"Strings", "[]string"},
		&ranger{"ListOfBytes", "[][]byte"},
	})
	gen(variableStr, "variable.gen.go", []interface{}{
		&varer{"Double", "double", "float64", "0"},
		&varer{"Int", "int", "int64", "0"},
		&varer{"Uint", "uint", "uint64", "0"},
		&varer{"Bool", "bool", "bool", "false"},
		&varer{"String", "string", "string", `""`},
		&varer{"Bytes", "[]byte", "[]byte", "nil"},
	}, `"github.com/katydid/katydid/parser"`)
	gen(typStr, "type.gen.go", []interface{}{
		&typer{"Double"},
		&typer{"Int"},
		&typer{"Uint"},
		&typer{"Bool"},
		&typer{"String"},
		&typer{"Bytes"},
	})
	gen(inSetStr, "inset.gen.go", []interface{}{
		&inSeter{"Int", "ConstInts", "int64"},
		&inSeter{"Uint", "ConstUints", "uint64"},
		&inSeter{"String", "ConstStrings", "string"},
	})
}
