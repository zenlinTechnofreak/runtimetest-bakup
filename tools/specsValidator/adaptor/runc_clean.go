package adaptor

import (
	"log"
	"os/exec"
)

func CleanRunc() {
	KillRunc()
	DeleteRun()
}

func KillRunc() {
	cmd := exec.Command("/bin/bash", "-c", "kill -9 `ps -ef|grep runc|grep -v grep|awk '{print $2}'`")
	_, err := cmd.Output()
	if err != nil {
		log.Printf("[clean runc] kill process error , %v", err)
	}
}

func DeleteRun() {
	_, err := exec.Command("/bin/bash", "-c", "ls /run/oci/specsValidator").Output()
	if err == nil {
		cmd := exec.Command("/bin/bash", "-c", "rm -r /run/oci/specsValidator")
		_, err = cmd.Output()
		if err != nil {
			log.Printf("[clean runc] delete folder error , %v", err)
		}
	}

}