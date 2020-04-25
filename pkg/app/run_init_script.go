package app

import (
	"os"
	"os/exec"
)

func runInitScript(path string, workDir string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}
	cmd := &exec.Cmd{
		Dir:    workDir,
		Path:   path,
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
	return cmd.Run()
}
