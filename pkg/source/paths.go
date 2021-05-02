package source

type Paths struct {
	BinFiles map[string]*BinFile
}

func NewPaths(binFiles []*BinFile) *Paths {
	pathsMap := make(map[string]*BinFile)

	for _, binValue := range binFiles {
		pathsMap[binValue.FileType] = binValue
	}
	return &Paths{
		pathsMap,
	}
}
