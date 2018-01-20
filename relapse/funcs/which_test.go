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

package funcs

import (
	"testing"

	"github.com/katydid/katydid/relapse/types"
)

func whichtest(t *testing.T, name string, params ...types.Type) {
	f, err := funcsMap.which(name, params...)
	if err != nil {
		t.Fatal(err)
	}
	if f.Name != name {
		t.Fatalf("name: want %s got %s", name, f.Name)
	}
}

func TestWhichStringEq(t *testing.T) {
	whichtest(t, "eq", types.SINGLE_STRING, types.SINGLE_STRING)
}

func TestWhichInt64Eq(t *testing.T) {
	whichtest(t, "eq", types.SINGLE_INT, types.SINGLE_INT)
}

func TestWhichInt64Ge(t *testing.T) {
	whichtest(t, "ge", types.SINGLE_INT, types.SINGLE_INT)
}

func TestWhichElem(t *testing.T) {
	whichtest(t, "elem", types.LIST_UINT, types.SINGLE_INT)
}
