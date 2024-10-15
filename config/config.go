package config

import (
	"fmt"
	"os"
)

var (
	DbName     string
	UserName   string
	DbPort     string
	DbPassword string
	ListenPort string
)

func GetDbName() string {
	DbName = os.Getenv("POSTGRES_DB")
	if DbName == "" {
		fmt.Println("Can't get the DbName!")
	}
	return DbName
}

func GetUserName() string {
	UserName = os.Getenv("POSTGRES_USER")
	if UserName == "" {
		fmt.Println("Can't get the DbUserName!")
	}
	return DbName
}

func GetDbPort() string {
	DbPort = os.Getenv("PORT")
	if DbPort == "" {
		fmt.Println("Can't get the DbPort!")
	}
	return DbPort
}

func GetDbPassword() string {
	DbPassword = os.Getenv("POSTGRES_PASSWORD")
	if DbPassword == "" {
		fmt.Println("Can't get the DbPassword!")
	}
	return DbPassword
}

func GetListenPort() string {
	ListenPort = os.Getenv("LISTEN_PORT")

	if ListenPort == "" {
		fmt.Println("Null Value ListenPort")
		return ""
	}
	return ListenPort
}
