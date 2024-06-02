package room

import (
	"log"
	"time"

	"github.com/MaxRazen/tutor/internal/db"
	"github.com/MaxRazen/tutor/internal/utils"
)

type RoomHistory struct {
	ID            int
	RoomId        int
	Authorship    string
	Transcription string
	Recording     string
	CreatedAt     time.Time
}

func createRoomHistoryRecord(rh RoomHistory) error {
	sql := `INSERT INTO room_history (room_id, authorship, transcription, recording) VALUES (?, ?, ?, ?)`

	return db.GetConnection().Exec(sql, rh.RoomId, rh.Authorship, rh.Transcription, rh.Recording)
}

func LoadRoomHistory(roomId int) ([]RoomHistory, error) {
	var records []RoomHistory
	sql := `SELECT id, role, transcription, recording, created_at
	FROM room_history
	WHERE room_id = ?
	ORDER BY id ASC
	`

	rows, err := db.GetConnection().Get(sql, roomId)

	if err != nil {
		log.Printf("room/history: fetching error %v", err)
		return records, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			id            int
			authorship    string
			transcription string
			recording     string
			createdAt     string
		)

		err := rows.Scan(&id, &authorship, &transcription, &recording, &createdAt)

		if err != nil {
			log.Printf("room/history: row parsing error %v", err)
			return records, err
		}

		records = append(records, RoomHistory{
			ID:            id,
			RoomId:        roomId,
			Authorship:    authorship,
			Transcription: transcription,
			Recording:     recording,
			CreatedAt:     utils.ParseTimestamp(createdAt),
		})
	}

	return records, nil
}
