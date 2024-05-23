package main

import (
	"fmt"
	imageprocessing "goroutines_pipeline/image_processing"
	"image"
	"os"
	"strings"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input path create a job and add it to
		// the out channel
		for _, p := range paths {
			job := Job{InputPath: p,
				OutPath: strings.Replace(p, "images/", "images/output/", 1)}
			job.Image = imageprocessing.ReadImage(p)
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		// For each input job, create a new job after resize and add it to
		// the out channel
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Resize(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input { // Read from the channel
			job.Image = imageprocessing.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input { // Read from the channel
			imageprocessing.WriteImage(job.OutPath, job.Image)
			_, err := os.Stat(job.OutPath)
			if err != nil {
				out <- false
				fmt.Println("Failed to save image: ", job.OutPath)
			} else {
				out <- true
			}
		}
		close(out)
	}()
	return out
}

func main() {

	var verifiedPaths []string

	imagePaths := []string{"images/baby_fox.jpg",
		"images/fox_mountain.jpg",
		"images/hawksbill_sea_turtle.jpg",
		"images/pangolin.jpg",
	}

	for _, image := range imagePaths {
		_, error := os.Stat(image)
		if error != nil {
			fmt.Println("Image not found: ", image, "removing from processing list.")
		} else {
			verifiedPaths = append(verifiedPaths, image)
		}

	}

	channel1 := loadImage(verifiedPaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	for success := range writeResults {
		if success {
			fmt.Println("Success!")
		} else {
			fmt.Println("Failed!")
		}
	}
}
