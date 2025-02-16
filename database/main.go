package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	fmt.Printf("Current directory: %s\n", currentDir)

	// Listar todos los archivos en la ruta actual
	files, err := os.ReadDir(currentDir)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	fmt.Println("Files in current directory:")
	for _, file := range files {
		fmt.Println(file.Name())
	}

	sqlBytes, err := os.ReadFile("database/distnetdb.sql")
	if err != nil {
		log.Fatalf("Failed to load sql script: %v", err)
	}
	fmt.Printf("sqlBytes: %v\n", string(sqlBytes))
}
