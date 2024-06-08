package room

import (
	"database/sql"
	"log"
	"time"

	"github.com/MaxRazen/tutor/internal/auth"
	"github.com/MaxRazen/tutor/internal/db"
	"github.com/MaxRazen/tutor/internal/utils"
)

type CreationData struct {
	Mode string `json:"mode"`
}

type Room struct {
	ID        int
	UserId    int
	Mode      string
	CreatedAt time.Time
	IsClosed  bool
}

type RoomStat struct {
	ID       int
	RoomID   int
	Requests int
	TalkTime float32
}

func CreateRoom(data CreationData, u *auth.User) (*Room, error) {
	var roomId int

	err := db.GetConnection().Transaction(func(tx *sql.Tx) error {
		sql := `INSERT INTO rooms (user_id, mode) VALUES (?, ?)`
		r, err := tx.Exec(sql, u.ID, data.Mode)

		if err != nil {
			return err
		}

		insertedId, err := r.LastInsertId()

		if err != nil {
			return err
		}

		roomId = int(insertedId)

		sql = `INSERT INTO room_stats (room_id) VALUES (?)`
		_, err = tx.Exec(sql, roomId)

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("room:createRoom: %v", err)
		return nil, err
	}

	log.Printf("room: room #%v has been created by user #%v", roomId, u.ID)

	room, err := FindRoom(roomId, u.ID)

	if err != nil {
		log.Println(err)
	}

	return room, nil
}

func FindRoom(roomId, relatedUserId int) (*Room, error) {
	var id int
	var userId int
	var mode string
	var createdAt string
	var isClosed bool

	sql := `SELECT id, user_id, mode, is_closed, created_at FROM rooms WHERE id = ? AND user_id = ?`
	err := db.GetConnection().First(sql, roomId, relatedUserId).Scan(&id, &userId, &mode, &isClosed, &createdAt)

	if err != nil {
		return nil, err
	}

	room := Room{
		ID:        id,
		UserId:    userId,
		Mode:      mode,
		CreatedAt: utils.ParseDatetime(createdAt),
		IsClosed:  isClosed,
	}

	return &room, nil
}
