package envdump

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
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

	path := filepath.Clean(file.Name())
	if err := notifyLocal(path); err != nil {
		return "", err
	}

	return path, nil
}

func notifyLocal(path string) error {
	values := url.Values{}
	values.Set("message", path)
	endpoint := "http://localhost:8080/message?" + values.Encode()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(endpoint)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("local notify failed: %s", resp.Status)
	}

	return nil
}
