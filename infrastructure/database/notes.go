package database

import (
	"database/sql"
	"errors"
	"rest-api-endpoints/domain/entity"
)

type NoteRepo struct {
	db *sql.DB
}

func NewNotesRepository(dbc *sql.DB) *NoteRepo {
	if dbc == nil {
		return nil
	}

	return &NoteRepo{dbc}
}

//Create part
func (nt *NoteRepo) Create(note *entity.CreateNote) error {

	var lastInsertId int64
	sSql := "INSERT INTO notes(accountid, nmid, note) VALUES($1, $2, $3) ON CONFLICT DO NOTHING RETURNING noteid;"
	err := nt.db.QueryRow(sSql, note.AccountID, note.ProductID, note.Text).Scan(&lastInsertId)

	if err != nil {
		return err
	}

	if lastInsertId == 0 {
		return errors.New("failed to create note")
	}

	return nil
}

// Fetch Read part
func (nt *NoteRepo) Fetch(accountID int64, productId int64) (entity.Notes, error) {

	result := entity.Notes{}

	var rows *sql.Rows
	var err error

	if productId == 0 {
		rows, err = nt.db.Query("SELECT accountid, nmid, note, created, updated FROM notes WHERE accountid=$1", accountID)
	} else {
		rows, err = nt.db.Query("SELECT accountid, nmid, note, created, updated FROM notes WHERE accountid=$1 AND nmid=$2", accountID, productId)
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		note := entity.Note{}

		err = rows.Scan(&note.AccountID, &note.ProductID, &note.Text, &note.Created, &note.Updated)
		if err != nil {
			panic(err.Error())
		}

		result = append(result, note)
	}

	if len(result) == 0 {
		return nil, errors.New("user not found")
	}
	return result, nil
}

// Update part
func (nt *NoteRepo) Update(note *entity.CreateNote) error {

	sSql := "UPDATE notes SET note=$3, updated=NOW() WHERE accountid=$1 AND nmid=$2"
	res, err := nt.db.Exec(sSql, note.AccountID, note.ProductID, note.Text)

	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// Delete part
func (nt *NoteRepo) Delete(note *entity.CreateNote) error {

	sSql := "DELETE FROM notes WHERE accountid=$1 AND nmid=$2"
	_, err := nt.db.Exec(sSql, note.AccountID, note.ProductID)
	if err != nil {
		return err
	}

	return nil
}
