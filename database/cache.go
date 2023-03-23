package database

import "github.com/Calgorr/URL_Shortener/model"

func AddLink(link *model.Link) error {
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()
	sqlstt := `INSERT INTO link (address,hash,usedtimes) VALUES ($1,$2,$3)`
	_, err = db.Exec(sqlstt, link.Address, link.Hash, link.UsedTimes)
	if err != nil {
		return err
	}
	return nil
}
