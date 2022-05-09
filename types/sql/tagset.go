// SPDX-License-Identifier: BSD-2-Clause
// Copyright (c) 2022-2022 Bitmark Inc.
// Use of this source code is governed by an BSD 2 Clause
// license that can be found in the LICENSE file.

package sqlTypes

import (
	"database/sql/driver"
	"fmt"
	"sort"
	"strings"
	"text/scanner"
)

// TagSet is a set of unique tag entries
type TagSet map[string]struct{}

// Add make a value present
func (a TagSet) Add(tag string) {
	a[tag] = struct{}{}
}

// AddAll make a list of values present
func (a TagSet) AddAll(tags []string) {
	for _, tag := range tags {
		a[tag] = struct{}{}
	}
}

// NewTagSet creates a fresh set from a list of strings
func NewTagSet(tags []string) TagSet {
	t := TagSet{}
	t.AddAll(tags)
	return t
}

// Has check if a value is present
func (a TagSet) Has(tag string) bool {
	_, ok := a[tag]
	return ok

}

// All the values sorted in ascending order
func (a TagSet) All() []string {
	list := make([]string, 0, len(a))
	for k := range a {
		list = append(list, k)
	}
	sort.Strings(list)
	return list
}

// Scan Postgres array into a TagSet, implements sql.Scanner interface
// parses a PostgreSQL array value like {abcd,"a,b",1234}
func (a *TagSet) Scan(value interface{}) error {

	src, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to unmarshal TagSet value: %v", value)
	}

	var s scanner.Scanner
	s.Init(strings.NewReader(src))
	s.Filename = "TagSet"
	s.Mode ^= scanner.SkipComments | scanner.ScanInts | scanner.ScanFloats

	m := TagSet{}
	t := ""
	start := false
	finish := false
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()
		switch txt {
		case "{":
			start = true
		case "}":
			finish = true
			fallthrough
		case ",":
			if len(t) > 0 {
				if t[0] == '"' {
					t = strings.ReplaceAll(t[1:len(t)-1], `\"`, `"`)
				}
				m[t] = struct{}{}
				t = ""
			}
		default:
			t = t + txt
		}
	}

	if !start || !finish {
		return fmt.Errorf("improper TagSet value: %q", src)
	}
	*a = m
	return nil
}

// Value returns Postgres Array value, implement driver.Value interface
// output is a PostgreSQL array of strings, all quoted: {"abcd","a,b","1234","he said: \"ok\""}
func (a TagSet) Value() (driver.Value, error) {
	s := "{" // array start
	for k := range a {
		s = s + fmt.Sprintf("%q,", k)
	}
	if len(s) > 1 {
		s = s[:len(s)-1] // strip extra ','
	}
	s = s + "}" // array finish
	return s, nil
}
