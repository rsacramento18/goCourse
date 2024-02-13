package fileutils

import (
	"os"
)

func WriteTextFile(filename, content string) error {
	err := os.WriteFile(filename, []byte(content), 0o644)
	return err
}
