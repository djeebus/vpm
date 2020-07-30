package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/vroomy/plugins"
)

func handleError(err error) {
	out.Error(err.Error())
	os.Exit(1)
}

func keyHasSuffixInAny(key string, pluginNames ...string) bool {
	_, key = plugins.ParseKey(key)
	for _, name := range pluginNames {
		_, name = plugins.ParseKey(name)
		if name == key {
			return true
		}
	}

	return false
}

func removeBranchHash(gitURL string) (out string) {
	out = strings.Split(gitURL, "#")[0]
	out = strings.Split(out, "@")[0]
	return
}

func getGoPath() (string, error) {
	goEnvCmd := exec.Command("go", "env", "GOPATH")
	output, err := goEnvCmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	goPath := strings.TrimSpace(string(output))
	return goPath, nil
}
