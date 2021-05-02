package db

import (
	//"database/sql/driver"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	//"gorm.io/gorm/logger"
	//"gorm.io/gorm/schema"
	//"gorm.io/plugin/dbresolver"

	"github.com/go-kit/kit/log"
	"github.com/hecomp/file-synchronizer/internal/utils"
	//"github.com/jmoiron/sqlx"
)

//type Connector interface {
//	Connect(mDB *sql.DB)
//	Close() error
//	Ping() error
//	Get() *gorm.DB
//}
//
//type connector struct {
//
//}

// NewConnection creates the connection to the database
func NewConnection(config *utils.Configurations, logger log.Logger) (*gorm.DB, error) {

	var conn string

	if config.DBConn != "" {
		conn = config.DBConn
	} else {
		host := config.DBHost
		port := config.DBPort
		user := config.DBUser
		dbName := config.DBName
		password := config.DBPass
		conn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	}
	logger.Log("connection string", conn)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})

	//db, err := sqlx.Connect("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
