package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("logi.txt")
	if err := os.Chmod("logi.txt", 0444); err != nil {
		fmt.Println("Ошибка", err)
		return
	}
	defer file.Close()

	f, err := os.Open("logi.txt")
	if err != nil {
		panic(err)
	}

	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("File name: %v\n", fi.Name())
	fmt.Printf("Is Directory: %t\n", fi.IsDir())
	fmt.Printf("Size: %d\n", fi.Size())
	fmt.Printf("Mode: %v\n", fi.Mode())

}
