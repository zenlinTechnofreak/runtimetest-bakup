// +build v0.1.1

// Copyright 2015 Huawei Inc. All Rights Reserved.
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

package linuxnamespace

import (
	"errors"
	"github.com/huawei-openlab/oct/tools/runtimeValidator/adaptor"
	"github.com/huawei-openlab/oct/tools/runtimeValidator/manager"
	"github.com/huawei-openlab/oct/tools/runtimeValidator/utils/configconvert"
	"github.com/huawei-openlab/oct/tools/runtimeValidator/utils/specsinit"
	"github.com/opencontainers/specs"
	"log"
	"os/exec"
	"strings"
)

/**
*Need mount proc and set mnt namespace when get namespace from container
*and the specs.Process.Terminal must be false when call runc in programe.
 */
var linuxSpec specs.LinuxSpec = specsinit.SetLinuxspecMinimum()
var linuxRuntimeSpec specs.LinuxRuntimeSpec = specsinit.SetLinuxruntimeMinimum()

var TestSuiteNP manager.TestSuite = manager.TestSuite{Name: "LinuxSpec.Linux.Namespaces"}

func init() {
	TestSuiteNP.AddTestCase("TestPidPathEmpty", TestPidPathEmpty)
	TestSuiteNP.AddTestCase("TestPidPathUnempty", TestPidPathUnempty)

	TestSuiteNP.AddTestCase("TestIpcPathEmpty", TestIpcPathEmpty)
	TestSuiteNP.AddTestCase("TestIpcPathUnempty", TestIpcPathUnempty)

	TestSuiteNP.AddTestCase("TestNetPathEmpty", TestNetPathEmpty)
	TestSuiteNP.AddTestCase("TestNetPathUnempty", TestNetPathUnempty)

	TestSuiteNP.AddTestCase("TestUtsPathEmpty", TestUtsPathEmpty)
	TestSuiteNP.AddTestCase("TestUtsPathUnempty", TestUtsPathUnempty)

	TestSuiteNP.AddTestCase("TestMntPathEmpty", TestMntPathEmpty)
	TestSuiteNP.AddTestCase("TestMntPathUnempty", TestMntPathUnempty)
	manager.Manager.AddTestSuite(TestSuiteNP)
}

/**
*Get host namespace
 */
func getHostNs(bashCommand string) (string, error) {
	var cmd *exec.Cmd
	cmd = exec.Command("/bin/sh", "-c", bashCommand)
	out, err := cmd.Output()
	if err != nil {
		return "", errors.New(string(out) + err.Error())
	}

	str := strings.TrimSuffix(string(out), "\n")
	str = strings.TrimSpace(str)
	if str == "" {
		return "", errors.New("can not find namespace on host.")
	}
	return str, nil
}

/**
*container unreused namespace of host
 */
func TestPathEmpty(linuxSpec *specs.LinuxSpec, linuxRuntimeSpec *specs.LinuxRuntimeSpec, hostNamespacePath string) (string, error) {
	//1. output json file for runc
	configFile := "./config.json"
	runtimeFile := "./runtime.json"
	err := configconvert.LinuxSpecToConfig(configFile, linuxSpec)
	err = configconvert.LinuxRuntimeToConfig(runtimeFile, linuxRuntimeSpec)
	if err != nil {
		log.Fatalf("write config error, %v", err)
	}

	//2. get container's pid namespace after executing  runc
	out, err := adaptor.StartRunc(configFile, runtimeFile)
	if err != nil {
		return manager.UNSPPORTED, errors.New(string(out) + err.Error())
	}
	containerNs := strings.TrimSuffix(string(out), "\n")
	containerNs = strings.TrimSpace(containerNs)
	if containerNs == "" {
		log.Fatalf("can not find namespace in container.")
	}

	//3. get host's all pid namespace
	cmd := "readlink " + hostNamespacePath + "|sort -u"
	hostNs, err := getHostNs(cmd)
	if err != nil {
		log.Fatalf("get host namespace error,%v\n", err)
	}

	//4. juge if the container's pid namespace is not in host namespaces
	var result string
	if strings.Contains(hostNs, containerNs) {
		result = manager.FAILED
	} else {
		result = manager.PASSED
	}

	return result, nil
}

/**
*container reused namespace of host
 */
func TestPathUnEmpty(linuxSpec *specs.LinuxSpec, linuxRuntimeSpec *specs.LinuxRuntimeSpec, hostNamespacePath string) (string, error) {
	//1. output json file for runc
	configFile := "./config.json"
	runtimeFile := "./runtime.json"
	err := configconvert.LinuxSpecToConfig(configFile, linuxSpec)
	err = configconvert.LinuxRuntimeToConfig(runtimeFile, linuxRuntimeSpec)
	if err != nil {
		log.Fatalf("write config error, %v", err)
	}

	//2. get container's pid namespace after executing  runc
	out, err := adaptor.StartRunc(configFile, runtimeFile)
	if err != nil {
		return manager.UNSPPORTED, errors.New(string(out) + err.Error())
		//log.Fatalf("write config error, %v\n", errors.New(string(out)+err.Error()))
	}
	containerNs := strings.TrimSuffix(string(out), "\n")
	containerNs = strings.TrimSpace(containerNs)
	if containerNs == "" {
		log.Fatalf("can not find namespace in container.")
	}

	//3. get host's pid namespace
	hostNs, err := getHostNs("readlink " + hostNamespacePath)
	if err != nil {
		log.Fatalf("get host namespace error,%v\n", err)
	}

	//4. juge if the container's pid namespace is in host namespaces
	var result string
	if strings.Contains(hostNs, containerNs) {
		result = manager.PASSED
	} else {
		result = manager.FAILED
	}

	return result, nil
}
