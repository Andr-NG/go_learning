package report

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func Getwd() {

	cwd := os.Getenv("ROOT")
	if cwd == "" {
		cwd = "./.."
	}

	path := filepath.Join(cwd, "report", time.Now().Format("2006-01-02_15:04:05"))
	fmt.Println(path)

	err := os.MkdirAll(path, 0o755)
	if err != nil {
		fmt.Printf("path not created: %s\n", err)
	} else {
		fmt.Println("path created")
	}

	os.Create(path + "/text.txt")
	fmt.Println(cwd)
}

func FindGoMod() (string, error) {

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(cwd + "/go.mod")); err == nil {
			return cwd, nil
		}
		parent := filepath.Dir(cwd)
		cwd = parent
	}

}
