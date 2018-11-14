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

package backend

import (
	"time"

	"github.com/beaukode/gohound/app"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// MongoDb Use mongo database to get hounds
type MongoDb struct {
	session    *mgo.Session
	collection *mgo.Collection
}

// NewMongoDb Use mongo database to get hounds
func NewMongoDb(url string, collection string) (*MongoDb, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &MongoDb{session: session, collection: session.DB("").C(collection)}, nil
}

// GetNextHounds Obtain next hounds to process
func (mdb *MongoDb) GetNextHounds(count int) ([]app.HoundInfo, error) {
	result := []app.HoundInfo{}
	err := mdb.collection.Find(bson.M{"nexttime": bson.M{"$lt": time.Now()}}).Limit(count).All(&result)

	return result, err
}

// Close Cleanup & Close
func (mdb *MongoDb) Close() {
	mdb.session.Close()
}
