# Tableau Go Client
Simple Tableau client library for Go. This client library used by our internal project to automate users onboarding.

## Features
* All [authentication](https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_authentication.htm) methods implemented.
* All [groups and site users](https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref_users_and_groups.htm) methods implemented.

## Installation
```
go get github.com/tiketdatarisal/tableau
```

## Usage
The following samples will assist you to use this library.
```go
import "github.com/tiketdatarisal/tableau"
```
First you have to configure Tableau client library, fill in `Host`, `Username`, `Password`, and `ContentUrl`. By default `Version` will used version `3.10`.
```go
cfg := tableau.Config{
	Host:       "https://your-tableau-server.com/",
	Version:    "3.12",
	Username:   "your-user-name",
	Password:   "your-password",
	ContentUrl: "your-content-url",
}
```
Create a new Tableau client instance.
```go
client, err := tableau.NewClient(cfg)
if err != nil {
	panic(err)
}
```
You will need to authenticate user before using client library. **NOTE:** This library will automatically called this method before calling other methods.
```go
err = client.Authentication.SignIn()
if err != nil {
    panic(err)
}
```
After that you can call other methods that you need. For example to get list of users on current site you can call following method.
```go
users, err := client.UsersGroups.GetUsersOnSite()
if err != nil {
    panic(err)
}
```
