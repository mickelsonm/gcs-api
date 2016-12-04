package client

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //this is needed for the mysql driver
	"github.com/mickelsonm/gcs-api/helpers/database"
)

//Clients is a group of clients
type Clients []Client

//Client is a client structure
type Client struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
}

var (
	getAllClients = `SELECT id, first_name, last_name, city, state, postal_code, email, phone FROM Clients`
	getClientByID = `SELECT id, first_name, last_name, city, state, postal_code, email, phone FROM Clients WHERE id = ?`
)

//GetAllClients gets a listing of all the clients
func GetAllClients() (Clients, error) {
	var clients Clients
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return clients, err
	}
	defer db.Close()

	stmt, err := db.Prepare(getAllClients)
	if err != nil {
		return clients, err
	}
	defer stmt.Close()
	res, err := stmt.Query()
	for res.Next() {
		var c Client
		err = res.Scan(&c.ID, &c.FirstName, &c.LastName, &c.City, &c.State, &c.PostalCode, &c.Email, &c.Phone)
		if err == nil {
			clients = append(clients, c)
		}
	}
	defer res.Close()
	return clients, nil
}

//GetClientByID gets a client by its id
func GetClientByID(id string) (Client, error) {
	var c Client
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return c, err
	}
	defer db.Close()

	stmt, err := db.Prepare(getClientByID)
	if err != nil {
		return c, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&c.ID, &c.FirstName, &c.LastName, &c.City, &c.State, &c.PostalCode,
		&c.Email, &c.Phone)

	return c, err
}
