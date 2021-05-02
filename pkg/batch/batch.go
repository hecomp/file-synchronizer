package batch

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/radovskyb/watcher"

	"github.com/hecomp/file-synchronizer/internal/constants"
	"github.com/hecomp/file-synchronizer/internal/utils"
	"github.com/hecomp/file-synchronizer/pkg/db/repo"
	"github.com/hecomp/file-synchronizer/pkg/destination"
	"github.com/hecomp/file-synchronizer/pkg/source"
)

type IBatch interface {
	Start()
}

type Batch struct {
	Src source.ISource
	Dest destination.IDestination
	Paths *source.Paths
	logger log.Logger
}

func NewBatch(pathMap *source.Paths, batchRepo repo.IBatchRepository, logger log.Logger) IBatch {
	dest := destination.NewDestination(batchRepo)
	source := source.NewSource()
	return &Batch{
		Dest: dest,
		Paths: pathMap,
		Src: source,
		logger: logger,
	}
}

func (b *Batch) Start() {

	b.logger.Log("Start Watching first data bin file")
	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Rename, watcher.Move, watcher.Create)
	r := regexp.MustCompile("^FIRST_DATA_BIN$")
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				b.logger.Log("Received First Data Bin File: %s", event.Path)
				err := b.RunFirstDataBinProcess()
				if err != nil {
					b.logger.Log("Error processing First Data bin file: %s", err)
				}
			case err := <-w.Error:
				b.logger.Log("Received unknown error %s", err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add(b.Paths.BinFiles[constants.FirstDataBin].CurrentPath); err != nil {
		b.logger.Log("Error occurred: %s", err)
	}

	for path, f := range w.WatchedFiles() {
		b.logger.Log("Watching path %s:\n", path, f.Name())
	}

	fmt.Println()

	if err := w.Start(time.Millisecond * 100); err != nil {
		b.logger.Log("Error occurred: %s", err)
	}
}

func (b *Batch) RunFirstDataBinProcess() error {
	var err error
	b.logger.Log("RunFirstDataBinProcess")
	fileName := b.Paths.BinFiles[constants.FirstDataBin].FileName
	b.Src, err = b.Src.LocateSource(b.Paths.BinFiles[constants.FirstDataBin].CurrentPath, fileName)
	if err != nil {
		b.logger.Log("Error locating first_data_bin file", err)
		return err
	}
	fileBINData, err := b.Src.ReadFirstDataBinFile()
	if err != nil {
		b.logger.Log("Error reading first_data_bin file", err)
		return err
	}
	if err := b.Dest.UploadFirstDataBinFile(fileBINData); err != nil {
		b.logger.Log("Error uploading first_data_bin file", err)
		return err
	}
	b.Src.GetFile().Close()
	currentPath := b.Paths.BinFiles[constants.FirstDataBin].CurrentPath
	// Copy File processed in file to archive
	if err := utils.ArchiveFile(b.Paths.BinFiles[constants.FirstDataBin].ArchivePath, currentPath, fileName); err != nil {
		b.logger.Log("Error archiving first_data_bin file", err)
		return err
	}
	// Delete processed
	fullFilePath := fmt.Sprintf("%s%s", currentPath, fileName)
	if err := utils.DeleteFile(fullFilePath); err != nil {
		b.logger.Log("Error deleting first_data_bin file", err)
		return err
	}
	return nil
}