package pkg

import (
	"os"
	"path/filepath"
)

func SortDownloads() error {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return err
	}

	downloadDir := homeDir + "/Downloads"

	files, err := os.ReadDir(downloadDir)

	if err != nil {
		return err
	}

	for _, file := range files {

		fileExt := filepath.Ext(file.Name())
		fileDir := downloadDir + "/" + file.Name()

		switch fileExt {
		case ".mkv", ".mp4":
			videoDir := homeDir + "/Videos/" + file.Name()
			moveFileToNewDir(fileDir, videoDir)
		case ".png", ".jpeg", ".svg", ".jpg":
			imageDir := homeDir + "/Pictures/" + file.Name()
			moveFileToNewDir(fileDir, imageDir)
		default:
			documentsDir := homeDir + "/Documents/" + file.Name()
			moveFileToNewDir(fileDir, documentsDir)
		}
	}

	return nil
}
