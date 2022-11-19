package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	var b bytes.Buffer
	file := "log.txt"
	if err := ioutil.WriteFile(file, b.Bytes(), 0666); err != nil {
		panic(err)
	}
	fii, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fii.Close()
	fi, err := fii.Stat()

	if err != nil {
		fmt.Println("Ошибка открытие файла", err)

	}

	fmt.Printf("File name: %v\n", fi.Name())
	fmt.Printf("Is Directory: %t\n", fi.IsDir())
	fmt.Printf("Size: %d\n", fi.Size())
	fmt.Printf("Mode: %v\n", fi.Mode())

}
