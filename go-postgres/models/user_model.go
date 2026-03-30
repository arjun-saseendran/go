package models

import (
	"database/sql"
	"fmt"
	"workout/db"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"required"`
	Place string `json:"place" binding:"required"`
}

var users = []User{}

// save
func (u *User) Save() error {
	query := `INSERT INTO users(name, age, place) VALUES($1, $2, $3) RETURNING id`
	err := db.DB.QueryRow(query, u.Name, u.Age, u.Place).Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

// users data
func GetAllUsers() ([]User, error) {
	// create empty slice
	var users = []User{}

	// query for all user data
	query := `SELECT id, name, age, place FROM users`

	// execute query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	// close db connection after the process
	defer rows.Close()

	// loop the rows
	for rows.Next() {
		// create temp user
		var u User

		// scan current row save to temp user
		err := rows.Scan(&u.ID, &u.Name, &u.Age, &u.Place)
		if err != nil {
			return nil, err
		}

		// add user data to slice
		users = append(users, u)
	}
	// check the error when loop
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// return the full user data
	return users, nil
}

// user data based on id
func GetUserByID(id int) (User, error) {
	var u User
	query := `SELECT id, name, age, place FROM users WHERE id = $1`
	err := db.DB.QueryRow(query, id).Scan(&u.ID,&u.Name, &u.Age, &u.Place)
	if err != nil {
		if err == sql.ErrNoRows {
			return u, fmt.Errorf("user with ID %d not found!", id)
		}
		return u, err
	}
	return u, nil
}

// update user data
func (u *User) Update() error {
	// query for update data
	query := `UPDATE users SET name = $1, age = $2, place = $3 WHERE id = $4`

	// exeute query
	result, err := db.DB.Exec(query, u.Name, u.Age, u.Place, u.ID)
	if err != nil {
		return err
	}
	// check already updated
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return fmt.Errorf("user with id %d not found", u.ID)
	}

	return nil
}

// delete user
func Delete(id int) error {
	// delete query
	query := `DELETE FROM users WHERE id = $1`

	// execute query
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	// check already deleted
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return fmt.Errorf("User with id %d not found!", id)
	}
	return nil
}
