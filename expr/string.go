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

package expr

import (
	"fmt"
	"github.com/katydid/katydid/expr/types"
	"regexp"
	"strconv"
	"strings"
)

func (this *Expr) String() string {
	space := this.Comma.String()
	if this.Terminal != nil {
		return space + this.GetTerminal().String()
	}
	if this.List != nil {
		return space + this.GetList().String()
	}
	if this.Function != nil {
		return space + this.GetFunction().String()
	}
	return this.BuiltIn.String()
}

func (this *NameExpr) String() string {
	v := this.GetValue()
	return v.(interface {
		String() string
	}).String()
}

// _decimal_digit : '0' - '9' ;
// _upcase : 'A'-'Z' ;
// _lowcase : 'a'-'z' ;
// _id_char : _upcase | _lowcase | '_' | _decimal_digit ;
// _id : (_upcase | _lowcase | '_' ) {_id_char} ;
// id : _id ;
var idRegexp = regexp.MustCompile("^[a-zA-Z_][_a-zA-Z0-9]*$")

func isId(s string) bool {
	return idRegexp.MatchString(s)
}

func (this *Name) String() string {
	if this.DoubleValue != nil {
		return this.Before.String() + strconv.FormatFloat(this.GetDoubleValue(), 'f', -1, 10)
	}
	if this.IntValue != nil {
		return this.Before.String() + strconv.FormatInt(this.GetIntValue(), 10)
	}
	if this.UintValue != nil {
		return this.Before.String() + "uint(" + strconv.FormatUint(this.GetUintValue(), 10) + ")"
	}
	if this.BoolValue != nil {
		return this.Before.String() + strconv.FormatBool(this.GetBoolValue())
	}
	if this.StringValue != nil {
		if isId(this.GetStringValue()) {
			return this.Before.String() + this.GetStringValue()
		}
		return this.Before.String() + strconv.Quote(this.GetStringValue())
	}
	if this.BytesValue != nil {
		return this.Before.String() + fmt.Sprintf("%#v", this.GetBytesValue())
	}
	panic("unreachable")
}

func (this *AnyName) String() string {
	return this.Underscore.String()
}

func (this *AnyNameExcept) String() string {
	return this.Exclamation.String() + this.OpenParen.String() +
		this.Except.String() + this.CloseParen.String()
}

func (this *NameChoice) String() string {
	return this.OpenParen.String() + this.Left.String() +
		this.Pipe.String() + this.Right.String() +
		this.CloseParen.String()
}

func (this *List) String() string {
	es := make([]string, len(this.GetElems()))
	for i, v := range this.GetElems() {
		es[i] = v.String()
	}
	return this.Before.String() + listTypeToString(this.Type) + this.OpenCurly.String() + strings.Join(es, "") + this.CloseCurly.String()
}

func listTypeToString(typ types.Type) string {
	switch typ {
	case types.LIST_DOUBLE:
		return "[]double"
	case types.LIST_INT:
		return "[]int"
	case types.LIST_UINT:
		return "[]uint"
	case types.LIST_BOOL:
		return "[]bool"
	case types.LIST_STRING:
		return "[]string"
	case types.LIST_BYTES:
		return "[][]byte"
	}
	panic("unreachable")
}

func (this *Function) String() string {
	ps := make([]string, len(this.GetParams()))
	for i, v := range this.GetParams() {
		ps[i] = v.String()
	}
	return this.Before.String() + this.GetName() + this.OpenParen.String() + strings.Join(ps, "") + this.CloseParen.String()
}

func (this *BuiltIn) String() string {
	return this.Symbol.String() + this.Expr.String()
}

func (this *Terminal) String() string {
	if this.DoubleValue != nil {
		return this.Before.String() + strconv.FormatFloat(this.GetDoubleValue(), 'f', -1, 64)
	}
	if this.IntValue != nil {
		return this.Before.String() + strconv.FormatInt(this.GetIntValue(), 10)
	}
	if this.UintValue != nil {
		return this.Before.String() + "uint(" + strconv.FormatUint(this.GetUintValue(), 10) + ")"
	}
	if this.BoolValue != nil {
		return this.Before.String() + strconv.FormatBool(this.GetBoolValue())
	}
	if this.StringValue != nil {
		return this.Before.String() + strconv.Quote(this.GetStringValue())
	}
	if this.BytesValue != nil {
		return this.Before.String() + fmt.Sprintf("%#v", this.GetBytesValue())
	}
	if this.Variable != nil {
		return this.Before.String() + this.Variable.String()
	}
	panic("unreachable")
}

func (this *Variable) String() string {
	typ := this.GetType()
	if types.IsList(typ) {
		types.ListToSingle(typ)
	}
	switch typ {
	case types.SINGLE_DOUBLE:
		return "$double"
	case types.SINGLE_INT:
		return "$int"
	case types.SINGLE_UINT:
		return "$uint"
	case types.SINGLE_BOOL:
		return "$bool"
	case types.SINGLE_STRING:
		return "$string"
	case types.SINGLE_BYTES:
		return "$[]byte"
	}
	panic(fmt.Errorf("unknown type %s", this.GetType()))
}

func (this *Keyword) String() string {
	if this == nil {
		return ""
	}
	return this.Before.String() + this.Value
}

func (this *Space) String() string {
	if this == nil {
		return ""
	}
	return strings.Join(this.Space, "")
}