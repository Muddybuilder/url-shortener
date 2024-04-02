package main

import "strconv"

type DB struct {
	data     map[string]string
	_counter int
}

func newDB() *DB {
	var db DB
	db.data = make(map[string]string)
	db._counter = 0
	return &db
}

type URL interface {
	insert(k string)
	retrieve(k string) string
	doesExist(k string) bool
}

func (db *DB) insert(k string) string {
	v := strconv.Itoa(db._counter)
	db._counter++
	db.data[k] = v
	return v
}

func (db *DB) retrieve(k string) string {
	return db.data[k]

}

func (db *DB) doesExist(k string) bool {
	_, ok := db.data[k]
	return ok
}
