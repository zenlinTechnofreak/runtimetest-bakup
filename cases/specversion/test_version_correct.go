// +build predraft

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

package specversion

import (
	"github.com/huawei-openlab/oct/tools/runtimeValidator/manager"
)

func TestVersionCorrect() string {

	linuxspec := setVersion(testValueCorrect)
	version := linuxspec.Spec.Version
	result, err := testVersion(&linuxspec)
	var testResult manager.TestResult
	testResult.Set("TestVersionCorrect", version, err, result)
	return testResult.Marshal()
}
