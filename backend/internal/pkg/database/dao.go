package database

import (
	"database/sql"
	"fmt"
	"test/internal/pkg/database/entity"
)

type Dao struct {
	db *sql.DB
}

func CreateNewDao(db *sql.DB) *Dao {
	return &Dao{db: db}
}

func (dao *Dao) GetKey(id string) (*entity.Key, error) {
	k := new(entity.Key)
	s := `SELECT * FROM keys WHERE id=$1`
	row := dao.db.QueryRow(s, id)
	err := row.Scan(&k.Id, &k.Payload)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return k, nil
	default:
		panic(err)
	}
}

func (dao *Dao) SaveKey(k entity.Key) error {
	s := `INSERT INTO keys (id, payload) VALUES ($1, $2) RETURNING id`
	var id int
	err := dao.db.QueryRow(s, k.Id, k.Payload).Scan(&id)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted public key %v", id)
	return nil
}

func (dao *Dao) GetFile(id string) (*entity.File, error) {
	f := new(entity.File)
	s := `SELECT * FROM files WHERE id=$1`
	row := dao.db.QueryRow(s, id)
	err := row.Scan(f.Id, f.Content)
	switch err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return f, nil
	default:
		panic(err)
	}
}

func (dao *Dao) GetAllFiles() ([]entity.File, error) {
	var files []entity.File
	s := `SELECT * FROM files`
	r, err := dao.db.Query(s)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	for r.Next() {
		var f entity.File
		err := r.Scan(&f.Id, &f.Content)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
	}
	return files, nil
}

func (dao *Dao) SaveFile(f entity.File) error {
	s := `INSERT INTO files (id, content) VALUES ($1, $2) RETURNING id`
	var id int
	err := dao.db.QueryRow(s, f.Id, f.Content).Scan(&id)
	if err != nil {
		return err
	}
	fmt.Printf("Inserted public key %v", id)
	return nil
}
