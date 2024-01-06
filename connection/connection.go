package connection

import (
	"sync"

	"github.com/aditya3232/tes_backend_developer_golang_bank_ina_muhammad_aditya/config"
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
// - dan untk elastic di log nya
func init() {
	initOnce.Do(func() {
		db, err := connectDatabaseMysql()
		if err != nil {
			panic(err)
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
