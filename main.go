// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This program is a CLI tool to connect to a Postgres instance and create a database from a SQL file.
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

// main is the entry point of the program.
func main() {
	host := flag.String("host", "localhost", "The database host.")
	port := flag.Int("port", 5432, "The database port.")
	user := flag.String("user", "postgres", "The database user.")
	password := flag.String("password", "", "The database password.")
	dbname := flag.String("dbname", "postgres", "The name of the database to connect to.")
	file := flag.String("file", "", "The path to the SQL file to execute.")

	flag.Parse()

	if *file == "" {
		fmt.Println("Please provide the path to the SQL file to execute using the -file flag.")
		os.Exit(1)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		*host, *port, *user, *password, *dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database!")

	sqlFile, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully executed the SQL file!")
}
