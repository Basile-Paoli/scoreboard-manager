package output

import (
	"fmt"
	"os"
	"strings"
)

func WriteFile(filename string, content string) error {
	path := strings.Split(filename, "/")
	if len(path) > 1 {
		dir := strings.Join(path[:len(path)-1], "/")
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	err := os.WriteFile(filename, []byte(content), os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}
	return nil
}
