package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgconn"
)

func main() {
	pgConn, err := pgconn.Connect(context.Background(), os.Getenv("PGX_TEST_CONN_STRING"))
	if err != nil {
		log.Fatalln(err)
	}
	defer pgConn.Close(context.Background())

	result := pgConn.ExecParams(context.Background(), "select generate_series(1,3)", nil, nil, nil, nil).Read()
	if result.Err != nil {
		log.Fatalln(result.Err)
	}

	for _, row := range result.Rows {
		log.Println(string(row[0]))
	}

	log.Println(result.CommandTag)
	// Output:
	// 1
	// 2
	// 3
	// SELECT 3
}
