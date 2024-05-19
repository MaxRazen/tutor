package db

import (
	"log"
	"time"
)

type migration struct {
	name string
	up   string
	down string
}

func MigrateDB() {
	log.Println("Migrating DB...")

	migrationTable := migration{
		name: "create_migration_table",
		up: `CREATE TABLE migrations (
			id int unsigned PRIMARY KEY AUTO_INCREMENT,
			name varchar(255) UNIQUE NOT NULL,
			created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL
		) ENGINE=InnoDB`,
		down: `DROP TABLE IF EXISTS migrations`,
	}

	if !GetConnection().IsTableExist("migrations") {
		err := GetConnection().Exec(migrationTable.up)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Table `migrations` is created")
	}

	var migrations []migration = []migration{
		{
			name: "create_users_table",
			up: `CREATE TABLE users (
				id int unsigned PRIMARY KEY AUTO_INCREMENT,
				name varchar(128) NOT NULL,
				email varchar(128) UNIQUE NOT NULL,
				social_id varchar(128) NOT NULL,
				avatar varchar(255) NOT NULL,
				last_logged_at timestamp DEFAULT CURRENT_TIMESTAMP,
				created_at timestamp DEFAULT CURRENT_TIMESTAMP
			 ) ENGINE=InnoDB`,
			down: `DROP TABLE IF EXISTS users`,
		},
		{
			name: "create_rooms_table",
			up: `CREATE TABLE rooms (
				id int unsigned PRIMARY KEY AUTO_INCREMENT,
				user_id int unsigned NOT NULL,
				mode enum('call', 'chat') NOT NULL,
				is_closed bool DEFAULT false NOT NULL ,
				created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
				CONSTRAINT FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
			) ENGINE=InnoDB`,
			down: `DROP TABLE IF EXISTS rooms`,
		},
		{
			name: "create_room_stats_table",
			up: `CREATE TABLE room_stats (
				id int unsigned PRIMARY KEY AUTO_INCREMENT,
				room_id int unsigned NOT NULL,
				requests int unsigned DEFAULT 0 NOT NULL,
				talk_time float DEFAULT 0,
				CONSTRAINT FOREIGN KEY (room_id) REFERENCES rooms (id) ON DELETE CASCADE
			) ENGINE=InnoDB`,
			down: `DROP TABLE IF EXISTS room_stats`,
		},
	}

	for _, m := range migrations {
		if isMigrationAppied(m.name) {
			continue
		}

		if err := applyMigration(m); err != nil {
			log.Fatalln(err)
		}

		log.Printf("Migration `%s` is applied", m.name)
	}
}

func isMigrationAppied(name string) bool {
	var count int = 0
	q := `SELECT COUNT(*) FROM migrations WHERE name = ? LIMIT 1`

	err := GetConnection().First(q, name).Scan(&count)

	if err != nil {
		log.Fatalln(err)
	}

	return count > 0
}

func applyMigration(m migration) error {
	err := GetConnection().Exec(m.up)

	if err != nil {
		return err
	}

	sql := `INSERT INTO migrations (name, created_at) VALUES (?, ?)`
	return GetConnection().Exec(sql, m.name, time.Now())
}
