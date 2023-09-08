package template

import (
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
)

const (
	url     = "https://github.com/samaelkorn/generator-web"
	tempDir = "temp"
)

func DownloadTemplate() error {
	workDir, err := os.Getwd()
	if err != nil {
		return err
	}

	workPath := filepath.Join(workDir, "generator")
	tempPath := filepath.Join(workPath, tempDir)
	distPath := filepath.Join(workPath, "dist")

	err = os.RemoveAll(distPath)

	if err != nil {
		return err
	}

	err = cloneGit(tempPath)

	if err != nil {
		return err
	}

	err = os.Rename(filepath.Join(tempPath, "dist"), distPath)

	if err != nil {
		return err
	}

	err = os.RemoveAll(tempPath)

	return err
}

func cloneGit(tempPath string) error {
	_, err := git.PlainClone(tempPath, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	return err
}
