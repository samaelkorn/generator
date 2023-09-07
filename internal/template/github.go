package template

import (
	"os"

	"github.com/go-git/go-git/v5"
)

const url = "https://github.com/samaelkorn/generator-web"

func DownloadTemplate() error {

	pwd, errPath := os.Getwd()
	if errPath != nil {
		return errPath
	}

	_, err := git.PlainClone(pwd, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	return err
}
