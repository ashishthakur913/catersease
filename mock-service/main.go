package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "mydb"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 1000000; i++ {
		user := generateUser()
		stmt, err := db.Prepare("INSERT INTO users (city, country, email, google_url, name, notes, phone_number, postal_code, state, street, unit) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
		if err != nil {
			log.Fatal(err)
		}
		_, err = stmt.Exec(user["city"], user["country"], user["email"], user["google_url"], user["name"], user["notes"], user["phone_number"], user["postal_code"], user["state"], user["street"], user["unit"])
		if err != nil {
			log.Fatal(err)
		}
		stmt.Close()

		if i%100 == 0 {
			fmt.Printf("Inserted %d users\n", i)
		}
	}

	insertOrdersForUsers(db)
}

func generateUser() map[string]string {
	rand.Seed(time.Now().UnixNano())

	city := randomString(10)
	country := randomString(10)
	email := randomString(10) + "@example.com"
	googleURL := "https://google.com/" + randomString(10)
	name := randomString(10)
	notes := randomString(10)
	phoneNumber := randomDigits(10)
	postalCode := randomDigits(6)
	state := randomString(10)
	street := randomString(10)
	unit := randomString(5)

	return map[string]string{
		"city":         city,
		"country":      country,
		"email":        email,
		"google_url":   googleURL,
		"name":         name,
		"notes":        getStringOrNil(&notes),
		"phone_number": phoneNumber,
		"postal_code":  postalCode,
		"state":        state,
		"street":       street,
		"unit":         getStringOrNil(&unit),
	}
}

func getStringOrNil(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func randomString(length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteByte(randUpperChar())
	}
	return builder.String()
}

func randUpperChar() byte {
	return byte(randInt(65, 90))
}

func randomDigits(length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		builder.WriteByte(byte(randInt(48, 57)))
	}
	return builder.String()
}

func randInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func insertOrdersForUsers(db *sql.DB) error {
	batchSize := 100
	var lastID int64 = 0

	for {
		rows, err := db.Query("SELECT id, name FROM users WHERE id > $1 ORDER BY id LIMIT $2", lastID, batchSize)
		if err != nil {
			return err
		}
		defer rows.Close()

		if !rows.Next() {
			break
		}

		var orders []string
		var users []struct {
			ID   int64
			Name string
		}

		for rows.Next() {
			var user struct {
				ID   int64
				Name string
			}
			if err := rows.Scan(&user.ID, &user.Name); err != nil {
				return err
			}
			users = append(users, user)

			driverID := rand.Intn(5) + 1
			order := fmt.Sprintf("INSERT INTO orders (user_id, driver_id, status, notes) VALUES (%d, %d, 'pending', 'Order for user %s')", user.ID, driverID, user.Name)
			orders = append(orders, order)
		}

		lastID = users[len(users)-1].ID

		tx, err := db.Begin()
		if err != nil {
			return err
		}

		for _, order := range orders {
			_, err := tx.Exec(order)
			if err != nil {
				tx.Rollback()
				return err
			}
		}

		err = tx.Commit()
		if err != nil {
			return err
		}

		fmt.Printf("Inserted orders for users up to ID %d\n", lastID)
	}

	return nil
}
