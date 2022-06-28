package main

import (
	"github.com/jackc/pgx/v4"
	"context"
	"log"
	"os"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// connect to database
	connectionUrl := os.Getenv("DATABASE_URL") // "postgres://postgres:postgres@172.18.0.2:5432/postgres"
	connection, err := pgx.Connect(ctx, connectionUrl)
	if err != nil {
		log.Fatal("DB connection error", err)
	}
	defer connection.Close(ctx)

	// ping database connection
	err = connection.Ping(ctx)
	if err != nil {
		log.Fatal("DB ping error", err)
	}

	// generate posts
	for z := 0; z < 20; z ++ {
		text := fmt.Sprintf("Post #%d", z)

		_, err := connection.Exec(ctx, "INSERT INTO posts (text, time) VALUES ($1, $2);", text, time.Now())
		if err != nil {
			log.Print("ERR:", err)
		}
	}

	// get posts
	rows, err := connection.Query(ctx, "SELECT * FROM posts")
	if err != nil {
		log.Print("ERR:", err)
	}

	for rows.Next() {
		var id uint64
		var text string
		var timestamp time.Time

		err := rows.Scan(&id, &text, &timestamp)
		if err != nil {
			log.Print("ERR:", err)
			continue
		}

		fmt.Printf("%d. %s (%s)\n", id, text, timestamp.Format(time.RFC3339))
	}

	fmt.Println("done!")
}
