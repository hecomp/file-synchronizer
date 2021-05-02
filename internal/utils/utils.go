package utils

import (
	"fmt"
	"io"
	"os"
	"time"
)

func ArchiveFile(archivePath, currentPath, fileName string) error  {
	const DateTimeFormat string = "010220061504"

	fullFilePath := fmt.Sprintf("%s%s", currentPath, fileName)
	currentFile, err := os.Open(fullFilePath)
	if err != nil {
		return err
	}
	defer currentFile.Close()

	currentTime := time.Now()
	newFile := fmt.Sprint(archivePath, fileName, "_", currentTime.Format(DateTimeFormat))
	destinationFile, err := os.Create(newFile)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, currentFile)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFile(path string) error  {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}