// SPDX-License-Identifier: BSD-2-Clause
// Copyright (c) 2022-2022 Bitmark Inc.
// Use of this source code is governed by an BSD 2 Clause
// license that can be found in the LICENSE file.

package dict_test

import (
	"testing"

	"github.com/bitmark-inc/go-utils/types/dict"
)

var data = dict.M{
	"a": "a string 1",
	"b": 99,
	"c": dict.A{
		dict.M{
			"a": "a string 2",
			"b": 88,
		},
		dict.M{
			"a": "a string 3",
			"b": 77,
		},
	},
	"d": dict.A{"one", "two", "three"},
}
var d = dict.New(data)

type strTest struct {
	idx []interface{}
	val string
}
type intTest struct {
	idx []interface{}
	val int
}
type int64Test struct {
	idx []interface{}
	val int64
}

func TestString(t *testing.T) {
	tests := []strTest{
		{
			idx: []interface{}{"a"},
			val: "a string 1",
		},
		{
			idx: []interface{}{"c", 0, "a"},
			val: "a string 2",
		},
		{
			idx: []interface{}{"c", 1, "a"},
			val: "a string 3",
		},
	}

	for i, test := range tests {
		actual := d.String(test.idx...)
		if actual != test.val {
			t.Errorf("%d: actual: %q  expected: %q", i, actual, test.val)
		}
	}
}

func TestStringSub(t *testing.T) {
	tests := []strTest{
		{
			idx: []interface{}{0, "a"},
			val: "a string 2",
		},
		{
			idx: []interface{}{1, "a"},
			val: "a string 3",
		},
	}

	d := d.Object("c")

	for i, test := range tests {
		actual := d.String(test.idx...)
		if actual != test.val {
			t.Errorf("%d: actual: %q  expected: %q", i, actual, test.val)
		}
	}
}

func TestInt(t *testing.T) {
	tests := []intTest{
		{
			idx: []interface{}{"a"},
			val: 0,
		},
		{
			idx: []interface{}{"c", 0, "b"},
			val: 88,
		},
		{
			idx: []interface{}{"c", 1, "b"},
			val: 77,
		},
	}

	for i, test := range tests {
		actual := d.Int(test.idx...)
		if actual != test.val {
			t.Errorf("%d: actual: %q  expected: %q", i, actual, test.val)
		}
	}
}

func TestInt64(t *testing.T) {
	tests := []int64Test{
		{
			idx: []interface{}{"a"},
			val: 0,
		},
		{
			idx: []interface{}{"c", 0, "b"},
			val: 88,
		},
		{
			idx: []interface{}{"c", 1, "b"},
			val: 77,
		},
	}

	for i, test := range tests {
		actual := d.Int64(test.idx...)
		if actual != test.val {
			t.Errorf("%d: actual: %q  expected: %q", i, actual, test.val)
		}
	}
}

func TestLength(t *testing.T) {
	tests := []intTest{
		{
			idx: []interface{}{"c"},
			val: 2,
		},
		{
			idx: []interface{}{"d"},
			val: 3,
		},
	}

	for i, test := range tests {

		actual := d.Length(test.idx...)
		if actual != test.val {
			t.Errorf("%d: actual: %q  expected: %q", i, actual, test.val)
		}

	}
}

/*
	al := 2
	if a.Length() != al {
		t.Errorf("%d: actual: %q  expected: %q", i, al, test.val)

	}
	t.Errorf("object: %q", a.String(0))

	t.Errorf("object: %q", a.String(1))

}
*/
