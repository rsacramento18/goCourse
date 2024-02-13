package main

import (
	"fmt"
	"log"
	"os"

	"frontendmasters.com/go/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	c, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")
	if err != nil {
		log.Fatalf("Error while reading file %v", err)
	}
	fmt.Println(c)
	newName := "Ana\nCarolina\nCoelho\ne\nSilva\nAntunes\nVicente"
	newContent := fmt.Sprintf("Original: %v\nNew Content: %v\n", c, newName)
	fileutils.WriteTextFile(rootPath+"/data/output.txt", newContent)

	newC, err2 := fileutils.ReadTextFile(rootPath + "/data/output.txt")
	if err2 != nil {
		log.Fatalf("Error while reading file %v", err2)
	}
	fmt.Println(newC)
}
