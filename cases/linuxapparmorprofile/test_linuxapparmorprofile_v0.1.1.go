// +build v0.1.1

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
//

package linuxapparmorprofile

import (
	"github.com/huawei-openlab/oct/tools/runtimeValidator/adaptor"
	"github.com/huawei-openlab/oct/tools/runtimeValidator/manager"
)

func TestLinuxApparmorProfile() string {
	apparmorfile := "testapporprofile"
	linuxspec, linuxruntimespec := setApparmorProfile(apparmorfile)
	result, err := testApparmorProfile(&linuxspec, &linuxruntimespec)
	adaptor.DeleteRun()
	var testResult manager.TestResult
	testResult.Set("TestLinuxApparmorProfile", linuxRuntimeSpec.Linux.ApparmorProfile, err, result)
	return testResult.Marshal()
}
