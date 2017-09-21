package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	// **incorrect answer
	// buffer := make([]byte, 1024)
	// rand.Reader.Read(buffer)
	// fmt.Println(buffer)

	// randFile, err := os.Create("rand.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer randFile.Close()

	// size, err := randFile.Write(buffer)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("write size : %d\n", size)

	// **correct answer
	file, err := os.Create("rand2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.CopyN(file, rand.Reader, 1024)
}
