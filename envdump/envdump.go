package envdump

import (
	"bufio"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// DumpAllToRandomLog writes all current environment variables to a random log file
// in the user's home directory and returns the file path.
func DumpAllToRandomLog() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	pattern := "envdump-" + time.Now().Format("20060102-150405") + "-*.log"
	file, err := os.CreateTemp(homeDir, pattern)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if err := file.Chmod(0o600); err != nil {
		return "", err
	}

	env := os.Environ()
	sort.Strings(env)

	writer := bufio.NewWriter(file)
	for _, line := range env {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			return "", err
		}
	}
	if err := writer.Flush(); err != nil {
		return "", err
	}

	return filepath.Clean(file.Name()), nil
}
