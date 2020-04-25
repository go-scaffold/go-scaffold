package app

import (
	"os"
	"os/exec"
)

func runInitScript(path string, workDir string, errHandler func(v ...interface{})) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return
	}
	cmd := &exec.Cmd{
		Dir:    workDir,
		Path:   path,
		Stderr: os.Stderr,
		Stdout: os.Stdout,
	}
	err := cmd.Run()
	if err != nil {
		errHandler("Error while executing init script. ", err)
		return
	}
}
