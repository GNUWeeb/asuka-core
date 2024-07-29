package utils

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/docker/docker/pkg/archive"
)

func TarWithOpt(src string) (io.ReadCloser, error) {

	tar, err := archive.TarWithOptions(src, &archive.TarOptions{})
	if err != nil {
		return nil, err
	}
	return tar, nil
}

func DebugPrint(rd io.Reader) error {
	var lastLine string

	scanner := bufio.NewScanner(rd)
	for scanner.Scan() {
		lastLine = scanner.Text()
		fmt.Println(scanner.Text())
	}

	errLine := &ErrorLine{}
	json.Unmarshal([]byte(lastLine), errLine)
	if errLine.Error != "" {
		return errors.New(errLine.Error)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
