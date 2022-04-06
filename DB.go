package main

type Database struct {
	DB map[string]string
}

func NewDatabase() *Database {
	store := &Database{DB: map[string]string{}}
	return store
}
func (DB Database) Get(key string) string {
	return DB.DB[key]
}
func (DB Database) Put(key string, value string) {
	DB.DB[key] = value
}
