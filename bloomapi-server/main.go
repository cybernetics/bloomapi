package main

import (
	"os"
	"fmt"
	"github.com/gocodo/bloomapi"
	"github.com/spf13/viper"
)

func showUsage() {
	fmt.Println("Usage: bloomapi-npi <command>")
	fmt.Println("=============================\n")
	fmt.Println("Avaialable commands:")
	fmt.Println("bloomapi-npi server    # run BloomAPI server")
	fmt.Println("bloomapi-npi bootstrap # setup BloomAPI shared schema")
}

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("Invalid command usage\n")
		showUsage()
		os.Exit(1)
	}

	arg := os.Args[1]

	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch arg {
	case "server":
		bloomapi.Server()
	case "bootstrap":
		bloomapi.Bootstrap()
	default:
		fmt.Println("Invalid command:", arg)
		showUsage()
		os.Exit(1)
	}
}