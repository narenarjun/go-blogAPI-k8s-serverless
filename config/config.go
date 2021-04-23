package config

import (
	"log"
	"os"
	"strconv"
)

var (
	PORT     = 4000
	DBURL    = ""
	DBDRIVER = ""
	TOKENKEY []byte
)

func Load() {

	var err error
	var er1 bool
	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		// log.Println(err)
		PORT = 4004
	}

	DBDRIVER, er1 = os.LookupEnv("DB_DRIVER")

	if !er1 {
		log.Fatalln("DB_DRIVER environment variable is missing.")
	}

	DBURL, er1 = os.LookupEnv("DATABASE_URL")

	if !er1 {
		log.Fatalln("DATABASE_URL environment variable is missing.")
	}

	key, er2 := os.LookupEnv("JWT_SECRET")

	if !er2 {
		log.Fatalln("JWT_SECRET enviroment variable is missing")
	}

	TOKENKEY = []byte(key)

}
