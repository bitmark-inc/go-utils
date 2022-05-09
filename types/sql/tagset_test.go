// SPDX-License-Identifier: BSD-2-Clause
// Copyright (c) 2022-2022 Bitmark Inc.
// Use of this source code is governed by an BSD 2 Clause
// license that can be found in the LICENSE file.

// test
package sqlTypes_test

import (
	"testing"

	sqlTypes "github.com/bitmark-inc/go-utils/types/sql"
	"github.com/stretchr/testify/suite"
)

type GithubTestSuite struct {
	suite.Suite
}

func (s *GithubTestSuite) SetupSuite() {
	s.T().Logf("setup suite\n")
}

func (s *GithubTestSuite) TearDownSuite() {
	s.T().Logf("tear down suite\n")
}

func (s *GithubTestSuite) SetupTest() {
	s.T().Logf("setup test\n")
}

func (s *GithubTestSuite) TearDownTest() {
	s.T().Logf("tear down test\n")
}

// tests

func (s *GithubTestSuite) TestTagSet() {

	ts := sqlTypes.TagSet{}

	ts.Add("bug")
	ts.Add("bug")
	ts.Add("ui")
	ts.Add("bug")
	ts.Add("apple")

	s.True(ts.Has("bug"))
	s.True(ts.Has("ui"))

	s.False(ts.Has("notexist"))

	ordered := []string{"apple", "bug", "ui"}
	s.Equal(ordered, ts.All())
}

func (s *GithubTestSuite) TestNewTagSet() {

	ts := sqlTypes.NewTagSet([]string{"bug", "spider", "snake"})

	s.False(ts.Has("elephant"))
	s.True(ts.Has("bug"))
	s.True(ts.Has("snake"))

	ordered := []string{"bug", "snake", "spider"}
	s.Equal(ordered, ts.All())

}

func (s *GithubTestSuite) TestOrdering() {

	ts := sqlTypes.NewTagSet([]string{"zebra", "mouse", "elephant", "kitten"})

	ordered := []string{"elephant", "kitten", "mouse", "zebra"}
	s.Equal(ordered, ts.All())

}

func TestGithub(t *testing.T) {
	suite.Run(t, &GithubTestSuite{})
}
