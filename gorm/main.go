package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connect to PostgreSQL using GORM
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Example user input (simulating JSON input)
	userInput := map[string]interface{}{
		"filter": map[string]interface{}{
			"age": map[string]interface{}{"$gte": 18},
		},
		"sort":   map[string]interface{}{"name": 1},
		"limit":  10,
		"offset": 0,
	}

	// Fetch data securely
	var users []User
	query := buildQuery(db, userInput)
	query.Find(&users)

	// Print results
	for _, user := range users {
		fmt.Printf("User: %+v\n", user)
	}
}

func buildQuery(db *gorm.DB, userInput map[string]interface{}) *gorm.DB {
	query := db.Model(&User{})

	// Handle filter conditions
	if filters, ok := userInput["filter"].(map[string]interface{}); ok {
		for key, value := range filters {
			switch v := value.(type) {
			case map[string]interface{}:
				for op, val := range v {
					switch op {
					case "$gte":
						query = query.Where(fmt.Sprintf("%s >= ?", key), val)
					case "$lte":
						query = query.Where(fmt.Sprintf("%s <= ?", key), val)
					case "$eq":
						query = query.Where(fmt.Sprintf("%s = ?", key), val)
					}
				}
			default:
				query = query.Where(fmt.Sprintf("%s = ?", key), value)
			}
		}
	}

	// Handle sorting
	if sort, ok := userInput["sort"].(map[string]interface{}); ok {
		for key, order := range sort {
			if order.(int) == 1 {
				query = query.Order(key + " ASC")
			} else {
				query = query.Order(key + " DESC")
			}
		}
	}

	// Handle limit & offset
	if limit, ok := userInput["limit"].(int); ok {
		query = query.Limit(limit)
	}
	if offset, ok := userInput["offset"].(int); ok {
		query = query.Offset(offset)
	}

	return query
}

type User struct{}