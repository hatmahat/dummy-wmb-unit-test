package main

import (
	"enigmacamp.com/golatihanlagi/delivery"
	_ "github.com/lib/pq"
)

func main() {
	delivery.NewServer().Run()
}

/*
go get github.com/stretchr/testify

set MONGO_HOST=127.0.0.1
set MONGO_PORT=27017
set MONGO_DB=enigma
set MONGO_USER=mahatmawisesa
set MONGO_PASSWORD=password
set API_HOST=localhost
set API_PORT=8888
*/
