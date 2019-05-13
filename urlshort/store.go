package urlshort

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Store interface {
	Add(redirect *Redirect) error
	Get(src string) (*Redirect, error)
}

type rdbStore struct {
	db *sql.DB
}

func MakeRDBStore(connString string) (Store, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}
	return rdbStore{db: db}, nil
}

func (store rdbStore) Add(redirect *Redirect) error {
	_, err := store.db.Query("INSERT INTO redirects(src, dest) VALUES ($1,$2)", redirect.Src, redirect.Dest)
	return err
}

func (store rdbStore) Get(src string) (*Redirect, error) {
	row := store.db.QueryRow("SELECT src, dest FROM redirects WHERE src=$1", src)

	redirect := &Redirect{}
	switch err := row.Scan(&redirect.Src, &redirect.Dest); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return redirect, nil
	default:
		fmt.Printf("Error: %v", err)
		panic(err)
	}
}
