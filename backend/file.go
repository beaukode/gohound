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
	"os"

	"github.com/beaukode/gohound/app"
	"gopkg.in/yaml.v2"
)

// File Use a yaml file to get probes
type File struct {
	config config
}

type probe struct {
	Type     string `yaml:"type"`
	Interval int    `yaml:"interval"`
}

type config struct {
	Probes []probe `yaml:"probes"`
}

// NewFile Use a yaml file to get probes
func NewFile(path string) (*File, error) {
	reader, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	dec := yaml.NewDecoder(reader)
	config := config{}
	err = dec.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &File{config: config}, nil
}

// GetNextHounds Obtain next hounds to process
func (f *File) GetNextHounds(count int) ([]app.HoundInfo, error) {
	result := []app.HoundInfo{}

	// TODO

	return result, nil
}

// Close Cleanup & Close
func (f *File) Close() {

}
