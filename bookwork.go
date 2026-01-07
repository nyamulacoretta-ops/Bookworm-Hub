package main

import (
	"encoding/json"
	"os"
)

//loadBookworms reads the file available and returns the list of bookworms,
//their loved books and all

func loadBookworm(filePath string) ([]bookworm, error) {
	//both slice of bookworms and error return zero
	f, err := os.Open(filePath) //cof, err := os.Open(filePath)  cof.Close()
	if err != nil {             //f is just a variable
		return nil, err
	}
	defer f.Close() //the close function will execute at the end
	//of the function

	var BookWorms []bookworm

	err = json.NewDecoder(f).Decode(&BookWorms)
	if err != nil {
		return nil, err
	}
	return BookWorms, nil //outputting the bookworms decoded
}

// We are translating the jsons into structs by using tags
type bookworm struct {
	Name string  `json:"name"`
	Book []Books `json:"Books"` //slice
}
type Books struct {
	Author  string `json:"author"`
	Title   string `json:"title"`
	Year    int    `json:"year of publication"`
	Edition int    `json:"edition"`
}
