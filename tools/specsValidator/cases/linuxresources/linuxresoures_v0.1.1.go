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

package linuxresources

import (
	"errors"
	"fmt"
	"github.com/huawei-openlab/oct/tools/specsValidator/adaptor"
	"github.com/huawei-openlab/oct/tools/specsValidator/manager"
	"github.com/huawei-openlab/oct/tools/specsValidator/utils/configconvert"
	"github.com/huawei-openlab/oct/tools/specsValidator/utils/specsinit"
	"github.com/opencontainers/specs"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var TestSuiteLinuxResources manager.TestSuite = manager.TestSuite{Name: "LinuxSpec.Linux.Resources"}

func init() {
	TestSuiteLinuxResources.AddTestCase("TestMemoryLimit", TestMemoryLimit)
	manager.Manager.AddTestSuite(TestSuiteLinuxResources)
}

func setResources(resources specs.Resources) (specs.LinuxSpec, specs.LinuxRuntimeSpec) {
	linuxSpec := specsinit.SetLinuxspecMinimum()
	linuxRuntimeSpec := specsinit.SetLinuxruntimeMinimum()
	linuxRuntimeSpec.Linux.Resources = &resources
	return linuxSpec, linuxRuntimeSpec
}

func testResources(linuxSpec *specs.LinuxSpec, linuxRuntimeSpec *specs.LinuxRuntimeSpec) (string, error) {
	configFile := "./config.json"
	runtimeFile := "./runtime.json"
	linuxSpec.Spec.Process.Args = []string{"/bin/bash", "-c", "sleep 3s"}
	err := configconvert.LinuxSpecToConfig(configFile, linuxSpec)
	err = configconvert.LinuxRuntimeToConfig(runtimeFile, linuxRuntimeSpec)
	out, err := adaptor.StartRunc(configFile, runtimeFile)
	if err != nil {
		return manager.UNSPPORTED, errors.New("StartRunc error :" + out + "," + err.Error())
	} else {
		return manager.PASSED, nil
	}
}

func checkConfigurationFromHost(filename string, configvalue string, failinfo string) (string, error) {
	cmd := exec.Command("bash", "-c", "cat  /sys/fs/cgroup/*/*/*/*/specsValidator/"+filename)
	cmdouput, err := cmd.Output()
	if err != nil {
		log.Fatalf("[specsValidator] linux resources test : read the "+filename+" error, %v", err)
		return manager.UNKNOWNERR, err
	} else {
		if strings.EqualFold(strings.TrimSpace(string(cmdouput)), configvalue) {
			return manager.PASSED, nil
		} else {
			return manager.FAILED, errors.New("test failed because" + failinfo)
		}
	}
}

func cleanCgroup() {
	// cmd := exec.Command("bash", "-c", "rmdir /sys/fs/cgroup/memory/user/1002.user/c2.session/specsValidator")
	// outPut, err := cmd.Output()
	var cmd *exec.Cmd
	time.Sleep(time.Second * 15)
	cmd = exec.Command("rmdir", "/sys/fs/cgroup/*/user/*/*/specsValidator")
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	outPut, err := cmd.Output()
	if err != nil {
		fmt.Println(string(outPut))
		log.Fatalf("[specsValidator] linux resources test : clean cgroup error , %v", err)
	}
	fmt.Println("clean cgroup sucess, ")
}