// Code generated by funcs-gen.
// DO NOT EDIT!

package funcs

import (
	"strings"
)

type listOfFloat64 struct {
	List []Float64
}

func NewListOfFloat64(v []Float64) Float64s {
	return &listOfFloat64{v}
}

func (this *listOfFloat64) Eval() []float64 {
	res := make([]float64, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval()
	}
	return res
}

func (this *listOfFloat64) String() string {
	ss := make([]string, len(this.List))
	for i := range this.List {
		ss[i] = Sprint(this.List[i])
	}
	return "[]float64{" + strings.Join(ss, ",") + "}"
}

type listOfFloat32 struct {
	List []Float32
}

func NewListOfFloat32(v []Float32) Float32s {
	return &listOfFloat32{v}
}

func (this *listOfFloat32) Eval() []float32 {
	res := make([]float32, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval()
	}
	return res
}

func (this *listOfFloat32) String() string {
	ss := make([]string, len(this.List))
	for i := range this.List {
		ss[i] = Sprint(this.List[i])
	}
	return "[]float32{" + strings.Join(ss, ",") + "}"
}

type listOfInt64 struct {
	List []Int64
}

func NewListOfInt64(v []Int64) Int64s {
	return &listOfInt64{v}
}

func (this *listOfInt64) Eval() []int64 {
	res := make([]int64, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval()
	}
	return res
}

func (this *listOfInt64) String() string {
	ss := make([]string, len(this.List))
	for i := range this.List {
		ss[i] = Sprint(this.List[i])
	}
	return "[]int64{" + strings.Join(ss, ",") + "}"
}

type listOfUint64 struct {
	List []Uint64
}

func NewListOfUint64(v []Uint64) Uint64s {
	return &listOfUint64{v}
}

func (this *listOfUint64) Eval() []uint64 {
	res := make([]uint64, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval()
	}
	return res
}

func (this *listOfUint64) String() string {
	ss := make([]string, len(this.List))
	for i := range this.List {
		ss[i] = Sprint(this.List[i])
	}
	return "[]uint64{" + strings.Join(ss, ",") + "}"
}

type listOfInt32 struct {
	List []Int32
}

func NewListOfInt32(v []Int32) Int32s {
	return &listOfInt32{v}
}

func (this *listOfInt32) Eval() []int32 {
	res := make([]int32, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval()
	}
	return res
}

func (this *listOfInt32) String() string {
	ss := make([]string, len(this.List))
	for i := range this.List {
		ss[i] = Sprint(this.List[i])
	}
	return "[]int32{" + strings.Join(ss, ",") + "}"
}

type listOfBool struct {
	List []Bool
}

func NewListOfBool(v []Bool) Bools {
	return &listOfBool{v}
}

func (this *listOfBool) Eval() []bool {
	res := make([]bool, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval()
	}
	return res
}

func (this *listOfBool) String() string {
	ss := make([]string, len(this.List))
	for i := range this.List {
		ss[i] = Sprint(this.List[i])
	}
	return "[]bool{" + strings.Join(ss, ",") + "}"
}

type listOfString struct {
	List []String
}

func NewListOfString(v []String) Strings {
	return &listOfString{v}
}

func (this *listOfString) Eval() []string {
	res := make([]string, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval()
	}
	return res
}

func (this *listOfString) String() string {
	ss := make([]string, len(this.List))
	for i := range this.List {
		ss[i] = Sprint(this.List[i])
	}
	return "[]string{" + strings.Join(ss, ",") + "}"
}

type listOfBytes struct {
	List []Bytes
}

func NewListOfBytes(v []Bytes) ListOfBytes {
	return &listOfBytes{v}
}

func (this *listOfBytes) Eval() [][]byte {
	res := make([][]byte, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval()
	}
	return res
}

func (this *listOfBytes) String() string {
	ss := make([]string, len(this.List))
	for i := range this.List {
		ss[i] = Sprint(this.List[i])
	}
	return "[][]byte{" + strings.Join(ss, ",") + "}"
}

type listOfUint32 struct {
	List []Uint32
}

func NewListOfUint32(v []Uint32) Uint32s {
	return &listOfUint32{v}
}

func (this *listOfUint32) Eval() []uint32 {
	res := make([]uint32, len(this.List))
	for i, e := range this.List {
		res[i] = e.Eval()
	}
	return res
}

func (this *listOfUint32) String() string {
	ss := make([]string, len(this.List))
	for i := range this.List {
		ss[i] = Sprint(this.List[i])
	}
	return "[]uint32{" + strings.Join(ss, ",") + "}"
}