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
	"time"

	"github.com/beaukode/gohound/app"
	"gopkg.in/yaml.v2"
)

// File Use a yaml file to get probes
type File struct {
	config config
	probes []app.ProbeInfo
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

	probes := make([]app.ProbeInfo, len(config.Probes), len(config.Probes))
	for i, v := range config.Probes {
		probes[i].Nexttime = time.Now()
		probes[i].Probetype = v.Type
	}

	return &File{config: config, probes: probes}, nil
}

// GetNextTodo Obtain next things to do
func (f *File) GetNextTodo(count int) ([]app.ProbeInfo, error) {
	var result []app.ProbeInfo

	now := time.Now()
	for _, v := range f.probes {
		if v.Nexttime.Before(now) {
			result = append(result, v)
			if len(result) == count {
				break
			}
		}
	}

	return result, nil
}

// Close Cleanup & Close
func (f *File) Close() {

}
