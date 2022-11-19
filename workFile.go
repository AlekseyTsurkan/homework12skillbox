package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	const exit = "exit"

	file, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("Не смогли создать файл ", err)
	}
	defer file.Close()

	d := time.Now()
	hour, min, sec := d.Clock()

	year, month, day := d.Date()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите текст ")

		scanner.Scan()
		stick := scanner.Text()
		file.WriteString(fmt.Sprint(year, ":", month, ":", day, ":", "time", hour, ":", min, ":", sec, ":", stick, "\n"))
		if stick == exit {
			break
		}

	}

}
