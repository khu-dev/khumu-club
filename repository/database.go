package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/khu-dev/khumu-club/config"
	"github.com/khu-dev/khumu-club/ent"
	"github.com/khu-dev/khumu-club/ent/migrate"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
	"time"
)

func Connect() *ent.Client {
	// parseTime=true가 없을 시
	// Error: unsupported Scan, storing driver.Value type []uint8 into type *time.Time
	// ref: https://stackoverflow.com/questions/45040319/unsupported-scan-storing-driver-value-type-uint8-into-type-time-time
	drv, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
			config.Config.DB.MySQL.User,
			config.Config.DB.MySQL.Password,
			config.Config.DB.MySQL.Host,
			config.Config.DB.MySQL.DatabaseName,
		))
	if err != nil {
		log.Panic(err)
	}
	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	conn, err := db.Conn(context.TODO())
	if err != nil {
		log.Panic(err)
	}
	conn.Close()

	db.SetConnMaxLifetime(time.Hour)
	ent.Debug()
	ent.Log(func(i ...interface{}) {
		log.Warn(i...)
	})

	client := ent.NewClient(ent.Driver(drv))
	err = client.Schema.Create(
		context.TODO(),
		migrate.WithDropIndex(true),
		migrate.WithForeignKeys(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func ConnectForTest() *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
