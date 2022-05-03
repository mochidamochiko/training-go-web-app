package data

import (
	"log"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

// 全スレッドの一覧を返す
func Threads() (threads []Thread, err error) {
	log.Printf("Pg query: Get all thread")
	rows, err := Db.Query(
		"SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC",
	)
	if err != nil {
		log.Printf("Error in data.Threads(): %s", err)
		return
	}

	for rows.Next() {
		th := Thread{}

		if err = rows.Scan(
			&th.Id,
			&th.Uuid,
			&th.Topic,
			&th.UserId,
			&th.CreatedAt,
		); err != nil {
			return
		}
		threads = append(threads, th)
	}

	rows.Close()

	return
}
