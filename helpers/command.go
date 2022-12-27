package helpers

import (
	"os/exec"
	"strings"
)

func ExecuteCommand(cmd string) (string, error) {
	s := strings.Split(cmd, " ")
	c, err := exec.Command(s[0], s[1:]...).Output()
	if err != nil {
		return "", err
	}

	cString := string(c)
	cString = strings.ReplaceAll(cString, " ", "&nbsp;")
	cString = strings.ReplaceAll(cString, "\r", "")
	cString = strings.ReplaceAll(cString, "\n", "<br/>")

	return cString, nil

}
