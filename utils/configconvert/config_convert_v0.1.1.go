// +build v0.1.1

// Copyright 2014 Google Inc. All Rights Reserved.
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

package configconvert

import (
	"encoding/json"
	"github.com/opencontainers/specs"
	"io/ioutil"
)

//read config.json to specs.LinuxSpec
func ConfigToLinuxRuntime(filePath string) (*specs.LinuxRuntimeSpec, error) {
	out, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var linuxruntime specs.LinuxRuntimeSpec
	err = json.Unmarshal(out, &linuxruntime)
	if err != nil {
		return nil, err
	}

	return &linuxruntime, nil
}

//write specs.LinuxRuntimeSpec to runtime.json
func LinuxRuntimeToConfig(filePath string, linuxRuntime *specs.LinuxRuntimeSpec) error {
	stream, err := json.Marshal(linuxRuntime)
	if err != nil {
		return err
	}
	objToJson(stream, filePath)
	return err
}
