package main

import "github.com/hookenz/gokv"

func main() {
	store := gokv.NewStore()

	_, err := store.Open("sqlite3", "./kv.db")
	if err != nil {
		panic(err)
	}
	defer store.Close()

	bucket, err := store.CreateBucket("version")
	bucket.Set("version", "1.25")
}
