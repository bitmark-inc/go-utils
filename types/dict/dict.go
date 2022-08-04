// SPDX-License-Identifier: BSD-2-Clause
// Copyright (c) 2022-2022 Bitmark Inc.
// Use of this source code is governed by an BSD 2 Clause
// license that can be found in the LICENSE file.

package dict

import (
	"fmt"
	"strconv"
)

// M is used as an abbriviation for generic map
type M map[string]interface{}

// A is used as an abbreviation for a generic array
type A []interface{}

// T is the generic object type
type T struct {
	d interface{}
}

// New converts a generic map or array to a generic object
func New(d interface{}) T {
	return T{
		d: d,
	}
}

// Object extracts a sub-object from an object
func (w T) Object(keys ...interface{}) T {
	d := w.d
	for _, key := range keys {
		i, k := idx(key)
		switch t := d.(type) {
		case M:
			d = t[k]
		case map[string]interface{}:
			d = t[k]
		case A:
			d = t[i]
		case []interface{}:
			d = t[i]
		default:
			return New(nil)
		}
	}
	return New(d)
}

// String extract a field as a string
func (w T) String(keys ...interface{}) string {
	switch t := w.Object(keys...).d.(type) {
	case string:
		return t
	case float64:
		return fmt.Sprintf("%f", t)
	case int64:
		return fmt.Sprintf("%d", t)
	case int:
		return fmt.Sprintf("%d", t)
	case nil:
		return ""
	default:
		return ""
	}
}

// Length gets the length of an array or zero if not
func (w T) Length(keys ...interface{}) int {
	switch t := w.Object(keys...).d.(type) {
	case M:
		return len(t)
	case map[string]interface{}:
		return len(t)
	case A:
		return len(t)
	case []interface{}:
		return len(t)
	default:
	}
	return 0
}

// Int returns the int value of a number or string
func (w T) Int(keys ...interface{}) int {
	switch t := w.Object(keys...).d.(type) {
	case string:
		n, _ := strconv.ParseInt(t, 10, 32)
		return int(n)
	case float64:
		return int(t)
	case int64:
		return int(t)
	case int:
		return t
	case nil:
		return 0
	default:
		return 0
	}
}

// Int returns the int64 value of a number or string
func (w T) Int64(keys ...interface{}) int64 {
	switch t := w.Object(keys...).d.(type) {
	case string:
		n, _ := strconv.ParseInt(t, 10, 64)
		return n
	case float64:
		return int64(t)
	case int64:
		return t
	case int:
		return int64(t)
	case nil:
		return 0
	default:
		return 0
	}
}

// internal functions

func idx(key interface{}) (int, string) {
	switch t := key.(type) {
	case string:
		i, _ := strconv.Atoi(t)
		return i, t
	case int:
		return t, strconv.Itoa(t)
	}
	return 0, ""
}
