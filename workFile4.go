package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	var b bytes.Buffer

	const exit = "exit"

	d := time.Now()
	hour, min, sec := d.Clock()

	year, month, day := d.Date()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите текст ")

		scanner.Scan()
		stick := scanner.Text()
		b.WriteString(fmt.Sprint(year, ":", month, ":", day, ":", "time", hour, ":", min, ":", sec, ":", stick, "\n"))
		if stick == exit {
			break
		}

	}

	fileName := "log.txt"
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	defer file.Close()
	resultButes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println("Сохраненый лог:")

	fmt.Println(string(resultButes))
}
