package models

import (
	"time"

	"example.com/events/db"
)

type Events struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time
	UserID      int64
}

var events = []Events{}

func (e *Events) Save() error {
	query := `
    INSERT INTO events(name, description, location, dateTime, user_id)
    VALUES (?, ?, ?, ?, ?)
    `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetallEvents() ([]Events, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Events

	for rows.Next() {
		var e Events
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEventById(id int64) (*Events, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var e Events
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

	if err != nil {
		return nil, err
	}

	return &e, nil
}

func (e Events) UpdateEvent() error {
	query := `
		UPDATE events
		SET name = ?, description = ?, location = ?,dateTime = ?
		WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return err
}

func (e Events) DeleteEvent() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err
}

func (e Events) Register (user_id int64) error {
	query := "INSERT INTO registrations(event_id , user_id) VALUES(?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}

	defer stmt.Close()
	_,err = stmt.Exec(e.ID,user_id)
	return err 
}

func (e Events) CancelRegisteration(user_id int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil{
		return err
	}

	defer stmt.Close()
	_,err = stmt.Exec(e.ID,user_id)
	return err 

}