package main

import (
	"errors"
	"os"
	"os/exec"
)

func fillEnv(env Environment) error {
	for envName, value := range env {
		if !value.NeedRemove {
			if err := os.Setenv(envName, value.Value); err != nil {
				return err
			}
		} else {
			if err := os.Unsetenv(envName); err != nil {
				return err
			}
		}
	}
	return nil
}

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := exec.Command(cmd[0], cmd[1:]...) // #nosec G204
	if err := fillEnv(env); err != nil {
		return 1
	}

	command.Env = os.Environ()
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		var exitError *exec.ExitError
		if errors.As(err, &exitError) {
			returnCode = exitError.ExitCode()
		} else {
			returnCode = 1
		}
		return
	}

	return
}
