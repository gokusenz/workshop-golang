package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		filename := fmt.Sprintf("%v.test.txt", i)
		fmt.Println("writing: " + filename)

		f, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		for i := 0; i < 10; i++ {
			fmt.Println(filename, ":", i)
			c := fmt.Sprintf("%v ", i)
			io.WriteString(f, c)
			time.Sleep(time.Millisecond * 100)
		}
	}
	var input string
	fmt.Scanln(&input)
}
