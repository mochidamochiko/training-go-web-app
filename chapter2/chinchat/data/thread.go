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

// スレッドのリプライ数を返す
func (thread *Thread) NumReplies() (count int) {
	log.Printf("Pg query: Get replies count")
	rows, err := Db.Query(
		"SELECT count(*) from posts where thread_id = $1",
		thread.Id,
	)
	if err != nil {
		return
	}

	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}

	rows.Close()

	return
}

// スレッドを開始したユーザ名を取得
func (thread *Thread) User() (user User) {
	log.Printf("Pg query: Get user name")

	user = User{}

	err := Db.QueryRow(
		"SELECT id, uuid, name, email, created_at FROM users WHERE id = $1",
		thread.UserId,
	).Scan(
		&user.Id,
		&user.Uuid,
		&user.Name,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		log.Print(err)
	}

	return
}

// CreatedAtの表示形式を変換
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("2006/01/02 15:04:05")
}
