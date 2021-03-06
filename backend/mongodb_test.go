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
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/beaukode/gohound/backend"
	"github.com/globalsign/mgo/bson"

	"github.com/globalsign/mgo"

	. "gopkg.in/check.v1"
)

type MongodbSuite struct {
	url        string
	collection *mgo.Collection
}

var _ = Suite(&MongodbSuite{})

var mongodb = flag.Bool("mongodb", false, "Include mongodb tests")

func (s *MongodbSuite) SetUpSuite(c *C) {
	if !*mongodb {
		c.Skip("-mongodb not provided")
	}
	url := "localhost:47017/gohound_test"
	session, err := mgo.Dial(url)
	if err != nil {
		c.Fatalf("Unable to dial  %s : is mongodb server ready ?", url)
		return
	}
	s.url = url
	rand.Seed(time.Now().UnixNano())
	name := fmt.Sprintf("hounds_test%d", rand.Intn(1000))
	name = "test"
	s.collection = session.DB("").C(name)
}

func (s *MongodbSuite) SetUpTest(c *C) {
	s.collection.DropCollection()
	s.collection.Insert(bson.M{"probetype": "tcp-connect", "nexttime": time.Date(2017, 11, 9, 12, 0, 0, 0, time.UTC), "interval": 1})
	s.collection.Insert(bson.M{"probetype": "http-response", "nexttime": time.Date(2017, 11, 9, 12, 0, 0, 0, time.UTC), "interval": 30})
	s.collection.Insert(bson.M{"probetype": "tcp-connect", "nexttime": time.Date(2016, 11, 9, 12, 0, 0, 0, time.UTC), "interval": 30})
	s.collection.Insert(bson.M{"probetype": "tcp-connect", "nexttime": time.Date(2017, 11, 9, 12, 0, 0, 0, time.UTC), "interval": 30})
}

func (s *MongodbSuite) TestNewMongoDb(c *C) {
	mdb, err := backend.NewMongoDb(s.url, s.collection.Name)
	c.Assert(err, IsNil)
	c.Assert(mdb, FitsTypeOf, &backend.MongoDb{})
}

func (s *MongodbSuite) TestGetNextTodo(c *C) {
	base := interfaceTests{backend: s.getBackend()}
	base.TestGetNextTodo(c)
}

func (s *MongodbSuite) TestGetNextTodoUseLimit(c *C) {
	base := interfaceTests{backend: s.getBackend()}
	base.TestGetNextTodoUseLimit(c)
}

func (s *MongodbSuite) TestGetNextPreventConcurrentAccess(c *C) {
	base := interfaceTests{backend: s.getBackend()}
	base.TestGetNextPreventConcurrentAccess(c)
}

func (s *MongodbSuite) TestUpdate(c *C) {
	base := interfaceTests{backend: s.getBackend()}
	base.TestUpdate(c)
}

func (s *MongodbSuite) getBackend() *backend.MongoDb {
	fbe, _ := backend.NewMongoDb(s.url, s.collection.Name)
	return fbe
}
