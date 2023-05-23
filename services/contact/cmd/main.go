package main

import (
	"fmt"

	"Lyrond.assignment_3/pkg/store/postgres"
	"Lyrond.assignment_3/services/contact/internal/domain"
)

func main() {
	dcp := &postgres.DbConnParams{
		Host:     "localhost",
		Port:     5432,
		User:     "user",
		Password: "User1234!",
		DbName:   "ass_3",
	}

	db, err := postgres.OpenDB(dcp)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	akhmet := domain.NewContact("Akhmetov", "Akhmet", "Akhmetovich")
	alina := domain.NewContact("Gumileva", "Alina", "Kalkenovna")
	group_1 := domain.NewGroup("Group 1")

	fmt.Println(akhmet, alina, group_1)
}
