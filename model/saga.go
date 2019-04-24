package model

import "time"

// EndChapter ends the
func EndChapter(tx Queryer, id int64) error {
	_, err := tx.Exec(`
		UPDATE
			chapter
		SET
			end_date = $1
		WHERE
			id = $2
`, time.Now(), id)
	if err != nil {
		return err
	}
	return nil

}
