package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("log.txt")

	if err != nil {
		fmt.Println("Ошибка чтение файла", err)
		return
	}
	defer file.Close()
	fi, err := file.Stat()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1)
	if _, err := io.ReadFull(file, buf); err != nil {
		fmt.Println("Ошибка последовательности байтов", err)
		return
	}

	fmt.Printf("%s\n", buf)
	fmt.Printf("File name: %v\n", fi.Name())
	fmt.Printf("Is Directory: %t\n", fi.IsDir())
	fmt.Printf("Size: %d\n", fi.Size())
	fmt.Printf("Mode: %v\n", fi.Mode())

}
