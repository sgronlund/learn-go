package main

import "github.com/gocql/gocql"

func connect() (*gocql.Session, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "my_db"
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4

	return cluster.CreateSession()
}

func main() {
	session, _ := connect()
	err := session.Query("CREATE KEYSPACE my_db WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 3}")
	err = session.Query("CREATE TABLE customers (customer_id text PRIMARY KEY, name text, balance int, country text)")
	if err != nil {

	}
	println("här")
	defer session.Close()
}
