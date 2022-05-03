package data

import (
	"log"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	CreatedAt time.Time
}

// 全スレッドの一覧を返す
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query(
		"SELECT id, created_at FROM threads ORDER BY created_ad DESC",
	)
	if err != nil {
		log.Print("errrrrrrrrorrr")
		// log.Print(err)
		return
	}
	// fmt.Printf("in Threads no error")
	// log.Print(rows)
	rows.Close()

	return
}
