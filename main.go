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

package main

import (
	"fmt"

	"github.com/beaukode/gohound/backend"
)

// MongoURL Url de connexion au cluster
const MongoURL = "mongodb://gomon:dZExedYpYy5ATCVA@mdbcluster-shard-00-00-nlhts.gcp.mongodb.net:27017,mdbcluster-shard-00-01-nlhts.gcp.mongodb.net:27017,mdbcluster-shard-00-02-nlhts.gcp.mongodb.net:27017/pingdog?ssl=true&replicaSet=mdbcluster-shard-0&authSource=admin"

func main() {
	backend, err := backend.NewMongoDb(MongoURL, "hounds")
	defer backend.Close()
	if err != nil {
		fmt.Println(err)
	}
	hounds, err := backend.GetNextHounds(10)
	fmt.Printf("Got %d hounds to do\n", len(hounds))
	hounds, err = backend.GetNextHounds(10)
	fmt.Printf("Got %d hounds to do\n", len(hounds))
	fmt.Println("The End.")
}
