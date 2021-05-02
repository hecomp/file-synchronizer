package main

import (
	"gorm.io/gorm"
	"os"

	"github.com/go-kit/kit/log"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/hecomp/file-synchronizer/internal/constants"
	"github.com/hecomp/file-synchronizer/internal/utils"
	batch2 "github.com/hecomp/file-synchronizer/pkg/batch"
	db2 "github.com/hecomp/file-synchronizer/pkg/db"
	"github.com/hecomp/file-synchronizer/pkg/db/repo"
	"github.com/hecomp/file-synchronizer/pkg/source"
)

var (
	app = kingpin.New("file-synchronizer", "Golang implementation for file processing").Author("file-synchronizer")
	env = app.Flag("env", "Defineds the environment").Envar("ENV").Default("locl").String()
	logLevel = app.Flag("log-level", "Defines log level").Envar("LOG_LEVEL").Default("DEBUG").Enum("DEBUG", "INFO", "ERROR")

	firstDataArchivePath = app.Flag("firstDataArchivePath", "first bin file archive path").Envar("FIRST_DATA_ARCHIVE_PATH").Default("data/first_data_bin/archive/").String()
	firstDataCurrentPath = app.Flag("firstDataArchivePath", "first bin file current path").Envar("FIRST_DATA_CURRENT_PATH").Default("data/first_data_bin/current/").String()

	archive = "first_data_bin/archive/"
	current = "first_data_bin/current/"
)

func main()  {

	// Create a single logger, which we'll use and give to other components.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	var configs *utils.Configurations
	{
		configs = utils.NewConfigurations(logger)
	}

	var db *gorm.DB
	var err error
	{
		db, err = db2.NewConnection(configs, logger)
		if err != nil {
			logger.Log("unable to connect to db", "error", err)
			panic(err)
		}
		//defer db.DB()
	}

	var (
		batchRepo = repo.NewBatchRepository(db, logger)
	)

	var binFiles []*source.BinFile
	{
		binFiles = source.NewBinFiles(constants.FirstDataBin, current, archive, constants.FirstDataBinFileName, binFiles)
	}

	pathMaps := source.NewPaths(binFiles)
	batch := batch2.NewBatch(pathMaps, batchRepo, logger)
	batch.Start()
}

