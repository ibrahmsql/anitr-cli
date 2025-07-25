package rofi

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"

	"github.com/xeyossr/anitr-cli/internal"
)

func isRofiExist() error {
	_, err := exec.LookPath("rofi")
	return err
}

func SelectionList(params internal.UiParams) (string, error) {
	err := isRofiExist()
	if err != nil {
		return "", errors.New("-rofi modunun çalışması için rofi'nin sisteminize yüklü olması gerekmektedir")
	}

	args := []string{"-dmenu", "-p", params.Label}

	if params.RofiFlags != nil {
		flags := strings.Split(*params.RofiFlags, " ")
		args = append(args, flags...)
	}

	input := bytes.NewBufferString("")
	for _, opt := range *params.List {
		input.WriteString(opt + "\n")
	}

	cmd := exec.Command("rofi", args...)
	cmd.Stdin = input

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	selection := strings.TrimSpace(string(out))
	return selection, nil
}

func InputFromUser(params internal.UiParams) (string, error) {
	err := isRofiExist()
	if err != nil {
		return "", errors.New("-rofi modunun çalışması için rofi'nin sisteminize yüklü olması gerekmektedir")
	}

	args := []string{"-dmenu", "-p", params.Label}

	if params.RofiFlags != nil {
		flags := strings.Split(*params.RofiFlags, " ")
		args = append(args, flags...)
	}

	cmd := exec.Command("rofi", args...)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	resp := strings.TrimSpace(string(out))
	return resp, nil
}
