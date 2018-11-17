// Copyright 2018 Jérémie COLOMBO
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package backend_test

import (
	"github.com/beaukode/gohound/backend"
	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
//func Test(t *testing.T) { TestingT(t) }

type FileSuite struct {
}

var _ = Suite(&FileSuite{})

func (s *FileSuite) SetUpSuite(c *C) {

}

func (s *FileSuite) SetUpTest(c *C) {

}

func (s *FileSuite) TestNewFile(c *C) {
	fbe, err := backend.NewFile("./file_test.yaml")
	c.Assert(err, IsNil)
	c.Assert(fbe, FitsTypeOf, &backend.File{})
}

func (s *FileSuite) TestNewFileMissingFile(c *C) {
	fbe, err := backend.NewFile("./i_dont_exists.yaml")
	c.Assert(err, ErrorMatches, "*no such file or directory")
	c.Assert(fbe, IsNil)
}

func (s *FileSuite) TestGetNextTodo(c *C) {
	todo, err := s.getBackend().GetNextTodo(10)
	c.Assert(todo, HasLen, 4)
	c.Assert(err, IsNil)
	c.Assert(todo[0].Probetype, Equals, "tcp-connect")
	c.Assert(todo[1].Probetype, Equals, "http-response")
	c.Assert(todo[2].Probetype, Equals, "tcp-connect")
	c.Assert(todo[3].Probetype, Equals, "tcp-connect")
}

func (s *FileSuite) TestGetNextTodoUseLimit(c *C) {
	todo, err := s.getBackend().GetNextTodo(1)
	c.Assert(todo, HasLen, 1)
	c.Assert(err, IsNil)
}

func (s *FileSuite) TestGetNextPreventConcurrentAccess(c *C) {
	fbe := s.getBackend()
	// Consume 1st probe
	todo, err := fbe.GetNextTodo(1)
	c.Assert(todo, HasLen, 1)
	c.Assert(err, IsNil)
	c.Assert(todo[0].Probetype, Equals, "tcp-connect")

	// Consume next 3 probes
	todo, err = fbe.GetNextTodo(3)
	c.Assert(todo, HasLen, 3)
	c.Assert(err, IsNil)
	c.Assert(todo[0].Probetype, Equals, "http-response")

	// No more probe
	todo, err = fbe.GetNextTodo(10)
	c.Assert(todo, HasLen, 0)
	c.Assert(err, IsNil)
}

func (s *FileSuite) getBackend() *backend.File {
	fbe, _ := backend.NewFile("./file_test.yaml")
	return fbe
}
