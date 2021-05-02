package destination

import (
	"github.com/hecomp/file-synchronizer/internal/files"
	"github.com/hecomp/file-synchronizer/pkg/db"
	"github.com/hecomp/file-synchronizer/pkg/db/repo"
)

const DefaultBulkInsertSize = 100

type IDestination interface {
	UploadFirstDataBinFile(records []*files.FirstDataBINFile) error
}

type Destination struct {
	Repo repo.IBatchRepository
}

func NewDestination(repo repo.IBatchRepository) IDestination  {
	return &Destination{
		Repo: repo,
	}
}

func (d Destination) UploadFirstDataBinFile(records []*files.FirstDataBINFile) error {
	var err error

	binFileRecords, err := BuildFirstDataBinFileRange(records)
	if err != nil {
		return err
	}
	for _, binFileRecordsSet := range SplitFirstDataRecords(binFileRecords, DefaultBulkInsertSize) {
		err = d.Repo.UploadFirstDataBinRange(binFileRecordsSet)
		if err != nil {
			return err
		}
	}
	return nil
}

func SplitFirstDataRecords(records []*db.GlobalBinRange, size int) [][]*db.GlobalBinRange {
	var chunkSet [][]*db.GlobalBinRange
	var chunk []*db.GlobalBinRange

	for len(records) > size {
		chunk, records = records[:size], records[:size]
		chunkSet = append(chunkSet, chunk)
	}
	if len(records) > 0 {
		chunkSet = append(chunkSet, records[:])
	}
	return chunkSet
}

func BuildFirstDataBinFileRange(records []*files.FirstDataBINFile) ([]*db.GlobalBinRange, error)  {
	var binFileRangeRecords []*db.GlobalBinRange
	for _, binData := range records {
		newBinRange := db.NewFirstDataBinRange(binData)
		binBinRange := newBinRange.(*db.GlobalBinRange)
		binFileRangeRecords = append(binFileRangeRecords, binBinRange)
	}
	return binFileRangeRecords, nil
}