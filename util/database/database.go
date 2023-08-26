package database

import (
	"context"
	"fmt"

	"time"

	mysqlDSN "github.com/go-sql-driver/mysql"
	"github.com/mandarinkb/go-example-lib/util/config"
	"github.com/mandarinkb/go-example-lib/util/datetime"
	"github.com/mandarinkb/go-example-lib/util/logg"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var Conn *gorm.DB

// Override logger interface from gorm logger.
type SqlLogger struct {
	logger.Interface
}

// InitDatabase is Initail database connection.
func InitDatabase() bool {

	gorm_db, err := initDatabaseGORM()

	if err != nil {
		logg.Printlogger("***** Database Connection failed ***** :: | ", err.Error())
		return false
	}
	Conn = gorm_db

	return true

}

// initDatabaseGORM is connect database with gorm lib.
func initDatabaseGORM() (db *gorm.DB, err error) {

	locAt := datetime.GetCurrentLocationTimeZoneAsiaBangkok()
	dsn_read_write := mysqlDSN.Config{
		User:                 config.Env.DB_READ_WRITE_USERNAME,
		Passwd:               config.Env.DB_READ_WRITE_PASSWORD,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", config.Env.DB_READ_WRITE_HOST, config.Env.DB_READ_WRITE_PORT),
		DBName:               config.Env.DB_READ_WRITE_NAME,
		ParseTime:            true,
		Loc:                  locAt,
		AllowNativePasswords: true,
	}
	dsn_read_only := mysqlDSN.Config{
		User:                 config.Env.DB_READ_ONLY_USERNAME,
		Passwd:               config.Env.DB_READ_ONLY_PASSWORD,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%v:%v", config.Env.DB_READ_ONLY_HOST, config.Env.DB_READ_ONLY_PORT),
		DBName:               config.Env.DB_READ_ONLY_NAME,
		ParseTime:            true,
		Loc:                  locAt,
		AllowNativePasswords: true,
	}

	gorm_db, err := gorm.Open(
		mysql.Open(
			dsn_read_write.FormatDSN(),
		),
		&gorm.Config{
			// Logger: &SqlLogger{},
			// Logger: logger.Default.LogMode(logger.Info),
			// DryRun: true,
		},
	)

	if err != nil {
		logg.Printlogger("***** Connect failed (dsn_read_write) ***** :: | ", err.Error())
		return nil, err
	}

	err = gorm_db.Use(
		dbresolver.Register(
			dbresolver.Config{
				Replicas: []gorm.Dialector{mysql.Open(dsn_read_only.FormatDSN())},
			},
		).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)
	if err != nil {
		logg.Printlogger("***** Connect failed (dsn_read_only) ***** :: | ", err.Error())
		return nil, err
	}

	return gorm_db, nil
}

// DatabaseClose is close the connection database.
func DatabaseClose() {
	dbInstance, _ := Conn.DB()
	_ = dbInstance.Close()
}

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

// Trace is Print SQL statement
func (sqlLog *SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sqlStatement, rowsAffected := fc()
	logg.Printlogger(fmt.Sprintf("%v***** SQL Statement ***** | ", Green), fmt.Sprintf("%v[ Row Affected : %v%v%v ] -> %v%v \n%v", BlueBold, YellowBold, rowsAffected, BlueBold, YellowBold, sqlStatement, Reset))
}
