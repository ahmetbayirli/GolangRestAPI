package models

import (
	"errors"

	_ "github.com/pelletier/go-toml/query"
	_ "go.starlark.net/resolve"
	"scroll2top.com/golang-rest-api/db"
	"scroll2top.com/golang-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// var events = []Event{}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password)
    VALUES (?,?)
    
    `
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

    hashedPassword, err :=    utils.HashPassword(u.Password)

    if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email,  hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.ID = id

	return err
}

func (u *User) ValidateCredenentials() error{
    query := `SELECT id, password FROM users where email = ?`
    row := db.DB.QueryRow(query, u.Email)

    var retrievedPassword string
    err := row.Scan(&u.ID, &retrievedPassword)
    if err != nil {
        return err
    }

    passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

    if !passwordIsValid {
        return errors.New("credentials invalid")
    }

    return nil
}
// func GetAllEvents() ([]Event, error) {

// 	query := `SELECT * FROM events`

// 	rows, err := db.DB.Query(query)

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var events []Event

// 	for rows.Next() {
// 		var event Event
// 		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
// 		if err != nil {
// 			return nil, err
// 		}
// 		events = append(events, event)

// 	}

// 	return events, nil
// }

// func GetEventById(id int64) (*Event, error) {
// 	query := `SELECT * FROM events WHERE id = ?`

// 	row := db.DB.QueryRow(query, id)

// 	var event Event
// 	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &event, nil
// }

// func (e *Event) Update() error {
// 	query := `UPDATE events
//     SET name = ? , description = ?, location = ?, dateTime = ?, userId = ?
//     WHERE id = ?
//     `
// 	stmt, err := db.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)

// 	if err != nil {
// 		return err
// 	}

// 	return err

// 	// events = append(events, *e)

// }

// func (e *Event) Delete() error {
// 	query := `DELETE FROM events WHERE id = ?`

// 	stmt, err := db.DB.Prepare(query)

// 	defer stmt.Close()

// 	_, err = stmt.Exec(e.ID)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
