package main

import (
	"database/sql"
	"fmt"

	"github.com/lrnholanda/go-intensivo-jul/entity/infra/database"
	"github.com/lrnholanda/go-intensivo-jul/usercase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := database.NewOrderRepository(db)
	uc := usercase.NewCalculateFinalPrice(orderRepository)

	input := usercase.OrderInput{
		ID:    "1234",
		Price: 10.0,
		Tax:   1.0,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
