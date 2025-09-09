package main

import (
	"fmt"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList []Bin

func createBin(id, name string, isPrivate bool) Bin {
	newBin := Bin{
		Id:        id,
		Private:   isPrivate,
		CreatedAt: time.Now(),
		Name:      name,
	}

	return newBin
}

func main() {
	newBin := createBin("Bin001", "New bin", false)
	fmt.Println(newBin)
}
