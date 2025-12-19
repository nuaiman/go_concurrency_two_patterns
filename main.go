package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(resultChannel chan string, jobChannel chan string) {
	for job := range jobChannel {
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Image processed: %s \n", job)

		resultChannel <- job
	}
}

func main() {
	images := []string{
		"image_1.png",
		"image_2.png",
		"image_3.png",
		"image_4.png",
	}

	var wg sync.WaitGroup
	resultChannel := make(chan string, 2)
	jobChannel := make(chan string, len(images))

	totalWorkers := 2
	startTime := time.Now()

	for i:=1;i<=totalWorkers; i++{
		wg.Go(func() {worker(resultChannel,jobChannel)})
	}
	
	go func ()  {
		wg.Wait()
		close(resultChannel)
	} ()

	for i:=range images {
		jobChannel <- images[i]
	}

	close(jobChannel)

	for result := range resultChannel {
		fmt.Printf("recieved %s \n", result)
	}

	fmt.Printf("Total Time Taken: %s \n", time.Since(startTime))
}