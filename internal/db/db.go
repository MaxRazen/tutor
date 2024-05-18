package db

import (
	"database/sql"
	"log"

	"github.com/MaxRazen/tutor/internal/config"
	"github.com/MaxRazen/tutor/internal/utils"
	_ "github.com/go-sql-driver/mysql"
)

var connection *DB

type DB struct {
	dbName string
	conn   *sql.DB
}

type QueryArgs struct {
	Args []any
}

func (db *DB) IsTableExist(table string) bool {
	var count int = 0

	sql := `SELECT COUNT(*) as count
	FROM information_schema.tables
	WHERE table_schema = ?
	  AND table_name = ?
	LIMIT 1`
	db.conn.QueryRow(sql, db.dbName, table).Scan(&count)

	return count > 0
}

func (db *DB) First(sql string, args ...any) *sql.Row {
	return db.conn.QueryRow(sql, args...)
}

func (db *DB) Exec(sql string, args ...any) error {
	_, err := db.conn.Exec(sql, args...)
	return err
}

func Connect() {
	dsn := config.GetEnv(config.DB_DSN, "")
	conn, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalln(err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatalln(err)
	}

	connection = &DB{
		dbName: utils.GetLastSegment(dsn, "/"),
		conn:   conn,
	}
}

func GetConnection() *DB {
	return connection
}
