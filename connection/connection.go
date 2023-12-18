package connection

import (
	"sync"

	"github.com/aditya3232/latihan-go-ms-one.git/config"
	log "github.com/aditya3232/latihan-go-ms-one.git/log"
	"gorm.io/gorm"
)

type Connection struct {
	db *gorm.DB
}

var (
	debug      int = config.CONFIG.DEBUG
	connection Connection
	initOnce   sync.Once
)

// untuk matikan koneksi ke database
// - dari init nya
// - dan dari repository nya
func init() {
	initOnce.Do(func() {
		db, err := connectDatabaseMysql()
		if err != nil {
			log.Panic(err)
		}

		connection = Connection{
			db: db,
		}
	})
}

func Close() {
	if connection.db != nil {
		sqlDB, _ := connection.db.DB()
		sqlDB.Close()
		connection.db = nil
	}
}
