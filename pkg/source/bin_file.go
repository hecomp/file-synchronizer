package source

type IBinFile interface {}

type BinFile struct {
	CurrentPath string
	ArchivePath string
	FileName string
	FileType string
}

func NewBinFiles(fileType, current, archive, fileName string, binFiles []*BinFile) []*BinFile {
	binFile := &BinFile{
		FileType: fileType,
		CurrentPath: current,
		ArchivePath: archive,
		FileName: fileName,
	}
	binFiles = append(binFiles, binFile)
	return binFiles
}