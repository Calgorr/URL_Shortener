package database

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Calgorr/URL_Shortener/model"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "calgor"
	password = "ami1r3ali"
	dbname   = "calhounio_demo"
)

var db *sql.DB

func Connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	return db, err
}

func AddLink(link *model.Link) error {
	db, err := Connect()
	if err != nil {
		return errors.New("Internal Server Error")
	}
	defer db.Close()
	sqlstt := `INSERT INTO links (address,hash,used_times) VALUES ($1,$2,$3)`
	_, err = db.Exec(sqlstt, link.Address, link.Hash, link.UsedTimes)
	if err != nil {
		return errors.New("Internal Server Error")
	}
	return nil
}

func GetLink(hash string) (*model.Link, error) {
	db, err := Connect()
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}
	defer db.Close()
	sqlstt := `SELECT * FROM links WHERE hash=$1`
	var link model.Link
	err = db.QueryRow(sqlstt, hash).Scan(&link.Address, &link.Hash, &link.UsedTimes)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}
	return &link, nil
}
