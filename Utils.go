package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func randFloat(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func firstRun() bool {
	filename := "General_specs.txt"

	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			os.Create(filename)
			return true
		}
	} else {
		return false
	}
	return false
}

func getSystemLanguage() (string, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("defaults", "read", "-g", "AppleLanguages")
	case "linux":
		cmd = exec.Command("bash", "-c", "echo $LANG")
	case "windows":
		cmd = exec.Command("cmd", "/c", "echo %LANG%")
	default:
		return "", fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	language := strings.TrimSpace(string(output))
	if runtime.GOOS == "darwin" {
		language = strings.ReplaceAll(language, "([", "")
		language = strings.ReplaceAll(language, "])", "")
		language = strings.Split(language, ", ")[0]
	}

	return language, nil
}
