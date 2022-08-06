package main

import (
	"fmt"
	"github.com/tiketdatarisal/tableau"
)

func main() {
	cfg := tableau.Config{
		Host:       "https://your-tableau-server.com/",
		Version:    "3.12",
		Username:   "your-user-name",
		Password:   "your-password",
		ContentUrl: "your-content-url",
	}

	client, err := tableau.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	err = client.Authentication.SignIn()
	if err != nil {
		panic(err)
	}

	users, err := client.UsersGroups.GetUsersOnSite()
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Printf("ID: %s, Name: %s, Site Role: %s\n", *user.ID, *user.Name, *user.SiteRole)
	}
}
