package logview

import (
	"os/exec"
)

func RunCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = NewWriter()
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
