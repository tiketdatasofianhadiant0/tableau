package main

import (
	"encoding/base64"
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

	image, err := client.WorkbooksViews.QueryViewImage("your-view-id")
	if err != nil {
		panic(err)
	}

	image64 := `data:image/png;base64,` + base64.StdEncoding.EncodeToString(image)
	fmt.Println(image64)
}
