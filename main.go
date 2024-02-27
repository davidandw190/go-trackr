package gotrackr

import (
	"log"

	"github.com/davidandw190/go-trackr/config"
	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage, err := NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}

	store := NewStore(db)

	api := NewAPIServer(":3000", store)
	api.Serve()
}
