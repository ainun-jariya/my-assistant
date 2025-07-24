package sandbox

import (
	"bytes"
	"os/exec"
)

func RunPythonCode(code string) (string, error) {
	cmd := exec.Command("docker", "run", "--rm", "-i", "python:3.10-slim", "python3", "-c", code)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stderr.String(), err
	}
	return out.String(), nil
}
