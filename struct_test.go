package sutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type DeepStruct struct {
	C *string `json:"c"`
	F string  `json:"f"`
	D int64   `json:"d"`
}

type InnerStruct struct {
	B *int64      `json:"b"`
	D *DeepStruct `json:"d"`
}

type MyStruct struct {
	A string       `json:"a"`
	I *InnerStruct `json:"i"`
}

type Int64 *int64
type String *string

func NewInt64(value int64) Int64 {
	return &value
}
func NewString(value string) String {
	return &value
}

func TestSet_Copy_Value(t *testing.T) {

	v := MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2), D: &DeepStruct{C: NewString("a")}}}

	r, err := New(v).WithValue("test").WithPath("A", false).Set()

	assert.NotNil(t, err)
	assert.False(t, r)
	assert.Equal(t, v.A, "Foo")
}

func TestSet_Address_Value(t *testing.T) {

	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2), D: &DeepStruct{C: NewString("a")}}}

	r, err := New(v).WithValue("test").WithPath("A", false).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Equal(t, v.A, "test")
}

func TestSetIsNil(t *testing.T) {

	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2), D: &DeepStruct{C: NewString("a")}}}

	r, err := New(v).WithValue(nil).WithPath("I.D.C", false).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Nil(t, v.I.D.C)
}

func TestSetIsNil_json(t *testing.T) {

	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2), D: &DeepStruct{C: NewString("a")}}}

	r, err := New(v).WithValue(nil).WithPath("i.d.c", true).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Nil(t, v.I.D.C)
}

func TestSet(t *testing.T) {

	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2), D: &DeepStruct{C: NewString("A")}}}

	r, err := New(v).WithValue(NewString("test")).WithPath("I.D.C", false).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Equal(t, *v.I.D.C, *NewString("test"))
}

func TestSet_json(t *testing.T) {

	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2), D: &DeepStruct{C: NewString("A")}}}

	r, err := New(v).WithValue(NewString("test")).WithPath("i.d.c", true).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Equal(t, *v.I.D.C, *NewString("test"))
}

func TestSet_Without_Addr(t *testing.T) {

	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2), D: &DeepStruct{C: NewString("A"), D: 1}}}

	r, err := New(v).WithValue(NewInt64(2)).WithPath("I.D.D", false).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Equal(t, v.I.D.D, int64(2))
}

func TestSet_Without_Addr_json(t *testing.T) {

	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2), D: &DeepStruct{C: NewString("A"), D: 1}}}

	r, err := New(v).WithValue(NewInt64(2)).WithPath("i.d.d", true).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Equal(t, v.I.D.D, int64(2))
}

func TestSet_Struct(t *testing.T) {

	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2)}}

	r, err := New(v).WithValue(&InnerStruct{B: NewInt64(3)}).WithPath("I", false).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Equal(t, *v.I.B, int64(3))
}

func TestSet_Struct_json(t *testing.T) {

	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2)}}

	r, err := New(v).WithValue(&InnerStruct{B: NewInt64(3)}).WithPath("i", true).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Equal(t, *v.I.B, int64(3))
}

func TestSet_High_Level(t *testing.T) {

	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2)}}

	r, err := New(v).WithValue("test").WithPath("A", false).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Equal(t, v.A, "test")
}

func TestSet_High_Level_json(t *testing.T) {

	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2)}}

	r, err := New(v).WithValue("test").WithPath("a", true).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Equal(t, v.A, "test")
}

func TestSet_Low_Level_String(t *testing.T) {
	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2), D: &DeepStruct{C: NewString("A"), D: 1, F: "Foo"}}}

	r, err := New(v).WithValue("test").WithPath("I.D.F", false).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Equal(t, v.I.D.F, "test")
}

func TestSet_Low_Level_String_Json(t *testing.T) {
	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2), D: &DeepStruct{C: NewString("A"), D: 1, F: "Foo"}}}

	r, err := New(v).WithValue("test").WithPath("i.d.f", true).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Equal(t, v.I.D.F, "test")
}

func TestSet_WithoutStruct(t *testing.T) {
	v := NewInt64(1)

	r, err := New(v).WithValue(NewInt64(2)).WithPath("", false).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Equal(t, *v, int64(2))
}

func TestSet_With_Nil(t *testing.T) {
	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2), D: &DeepStruct{C: NewString("A"), D: 1, F: "Foo"}}}

	r, err := New(v).WithValue(nil).WithPath("I.B", false).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Nil(t, v.I.B)
}

func TestSet_With_Nil_Json(t *testing.T) {
	v := &MyStruct{A: "Foo", I: &InnerStruct{B: NewInt64(2), D: &DeepStruct{C: NewString("A"), D: 1, F: "Foo"}}}

	r, err := New(v).WithValue(nil).WithPath("i.b", true).Set()

	assert.Nil(t, err)
	assert.True(t, r)
	assert.Nil(t, v.I.B)
}
