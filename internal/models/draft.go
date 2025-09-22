package models

import (
	"database/sql"
	"errors"
	"time"
)

type Draft struct {
	id      int
	title   string
	content string
	created time.Time
	changed time.Time
}

type DraftModel struct {
	DB *sql.DB
}

func (m *DraftModel) Insert(title string, content string) (int, error) {
	stmt := `INSERT INTO draft(title, content, created, changed)
	VALUES(?, ?, NOW(), NOW())`
	result, err := m.DB.Exec(stmt, title, content)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *DraftModel) Get(id int) (*Draft, error) {
	stmt := `SELECT id, title, content, created, changed
	FROM draft
	WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	d := &Draft{}

	err := row.Scan(&d.id, &d.title, &d.content, &d.created, &d.changed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return d, nil
}

func (m *DraftModel) Latest() ([]*Draft, error) {
	return nil, nil
}
