package migration

import (
	"database/sql"

	"github.com/deluan/navidrome/consts"
	"github.com/google/uuid"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upAddDefaultTranscodings, downAddDefaultTranscodings)
}

func upAddDefaultTranscodings(tx *sql.Tx) error {
	row := tx.QueryRow("SELECT COUNT(*) FROM transcoding")
	var count int
	err := row.Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	stmt, err := tx.Prepare("insert into transcoding (id, name, target_format, default_bit_rate, command) values (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	for _, t := range consts.DefaultTranscodings {
		r, _ := uuid.NewRandom()
		_, err := stmt.Exec(r.String(), t["name"], t["targetFormat"], t["defaultBitRate"], t["command"])
		if err != nil {
			return err
		}
	}
	return nil
}

func downAddDefaultTranscodings(tx *sql.Tx) error {
	return nil
}
