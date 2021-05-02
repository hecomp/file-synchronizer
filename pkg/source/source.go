package source

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"os"

	"github.com/ahmedalhulaibi/flatfile"

	"github.com/hecomp/file-synchronizer/internal/files"
)

type ISource interface {
	ReadFirstDataBinFile() ([]*files.FirstDataBINFile, error)
	LocateSource(currentLocation, fileName string) (ISource, error)
	GetFile() *os.File
}

type Source struct {
	File *os.File
}

func NewSource() ISource  {
	return &Source{}
}

// New creates a new FileLocation based on specified file.
func (s *Source) LocateSource(currentLocation, fileName string) (ISource, error) {
	path := fmt.Sprint(currentLocation, fileName)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &Source{
		File: file,
	}, nil
}

func (s *Source) ReadFirstDataBinFile() ([]*files.FirstDataBINFile, error) {
	var binFileRecords []*files.FirstDataBINFile
	var err error
	data := bufio.NewScanner(charmap.ISO8859_1.NewDecoder().Reader(s.File))
	defer s.File.Close()

	for data.Scan() {
		if string(data.Text()[0]) == "D"{
			binFileRecords, err = buildFirstDataBinFileRecords(data, binFileRecords...)
			if err != nil {
				return nil, err
			}
		}
	}
	err = data.Err()
	if err != nil {
		return nil, err
	}
	return binFileRecords, nil
}

func (s *Source) GetFile() *os.File {
	return s.File
}

func buildFirstDataBinFileRecords(data *bufio.Scanner, binFileRecords ...*files.FirstDataBINFile) ([]*files.FirstDataBINFile, error) {
	binFileData := &files.FirstDataBINFile{}
	err := flatfile.Unmarshal([]byte(data.Text()), binFileData, 0, 0, false)
	if err != nil {
		return nil, err
	}
	binFileRecords = append(binFileRecords, binFileData)
	return binFileRecords, nil
}
