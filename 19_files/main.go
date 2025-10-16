package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("File Name:", fileInfo.Name())
	// fmt.Println("File Size:", fileInfo.Size(), "bytes")
	// fmt.Println("Is Directory:", fileInfo.IsDir())

	// defer f.Close()

	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// defer f.Close()

	// buffer := make([]byte, 30)
	// readFil, readingFileErr := f.Read(buffer)
	// if readingFileErr != nil {
	// 	fmt.Println("Error:", readingFileErr)
	// 	return
	// }
	// fmt.Println(buffer)
	// for i := 0; i < len(buffer); i++ {
	// 	fmt.Print(string(buffer[i]) + " ")
	// }
	// fmt.Println("Bytes read:", readFil)
	// fmt.Println("Content:", string(buffer[:readFil]))

	// f, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(string(f))

	// dir, err := os.Open("../")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(0)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// for _, file := range fileInfo {
	// 	fmt.Println("Name:", file.Name(), "IsDir:", file.IsDir())
	// }

	// f, err := os.Create("newfile.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// defer f.Close()

	// f.WriteString("Hello World")

	// byteString := []byte("\nthis is byte str adsh")

	// f.Write(byteString)

	// read from example.txt and write to newfile.txt
	sourceFile, err := os.Open("example.txt")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create("newfile.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer destFile.Close()

	// buffer := make([]byte, 30)

	// for i := 0; i < 5; i++ {
	// 	readBytes, err := sourceFile.Read(buffer)
	// 	if err != nil {
	// 		fmt.Println("Error:", err)
	// 		return
	// 	}
	// 	destFile.Write(buffer[:readBytes])
	// }

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error:", err)
			}
			break
		}
		writer.WriteByte(b)
	}

	writer.Flush()

}
