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

// Add new draft into DB
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

// Return one draft
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

// Return latest 10 drafts added
func (m *DraftModel) Latest() ([]*Draft, error) {
	stmt := `SELECT id, title, content, created, changed
	FROM draft
	ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	drafts := []*Draft{}

	for rows.Next() {
		d := &Draft{}
		err = rows.Scan(&d.id, &d.title, &d.content, &d.created, &d.changed)
		if err != nil {
			return nil, err
		}
		drafts = append(drafts, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return drafts, nil
}
