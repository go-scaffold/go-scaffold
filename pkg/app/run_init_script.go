package app

import (
	"os/exec"
)

func runInitScript(path string, workDir string) error {
	cmd := &exec.Cmd{
		Dir:  workDir,
		Path: path,
	}
	return cmd.Run()
}
