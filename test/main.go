package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sync"
	"time"
)

func main() {
	k := 3
	networkName := "test_kademlia"
	imageName := "test"
	workdir := "/app"

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}
	volume := fmt.Sprintf("%s:/app", pwd)

	var wg sync.WaitGroup

	for i := 1; i <= k; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			containerName := fmt.Sprintf("client%d", i)
			dbVolume := fmt.Sprintf("db_%s:/app/data", containerName)
			cmd := exec.Command(
				"docker", "run",
				"-d",
				"--network", networkName,
				"--network-alias", containerName,
				// "--hostname", "0.0.0.0",
				"-v", volume,
				"-v", dbVolume, // Volumen único para la DB
				"-w", workdir,
				"--name", containerName,
				imageName,
				"sh",
			)
			err := cmd.Start()
			if err != nil {
				log.Printf("Failed to run container %s: %v", containerName, err)
			} else {
				log.Printf("Container %s started successfully", containerName)
			}
		}(i)
		<-time.After(1 * time.Minute)
	}

	wg.Wait()
	log.Println("All containers have been started")
}
