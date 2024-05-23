package main

import (
	imageprocessing "goroutines_pipeline/image_processing"
	"testing"
)

func BenchmarkLoadImage(b *testing.B) {
	imagePaths := []string{"images/baby_fox.jpg",
		"images/fox_mountain.jpg",
		"images/hawksbill_sea_turtle.jpg",
		"images/pangolin.jpg",
	}

	for i := 0; i < b.N; i++ {
		loadImage(imagePaths)
	}
}

func BenchmarkResize(b *testing.B) {
	imagePaths := []string{"images/baby_fox.jpg",
		"images/fox_mountain.jpg",
		"images/hawksbill_sea_turtle.jpg",
		"images/pangolin.jpg",
	}

	jobs := loadImage(imagePaths)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		resize(jobs)
	}
}

func BenchmarkConvertToGrayscale(b *testing.B) {
	imagePaths := []string{"images/baby_fox.jpg",
		"images/fox_mountain.jpg",
		"images/hawksbill_sea_turtle.jpg",
		"images/pangolin.jpg",
	}

	jobs := loadImage(imagePaths)
	jobs = resize(jobs)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		convertToGrayscale(jobs)
	}
}

func BenchmarkSaveImage(b *testing.B) {
	imagePaths := []string{"images/baby_fox.jpg",
		"images/fox_mountain.jpg",
		"images/hawksbill_sea_turtle.jpg",
		"images/pangolin.jpg",
	}

	jobs := loadImage(imagePaths)
	jobs = resize(jobs)
	jobs = convertToGrayscale(jobs)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		saveImage(jobs)
	}
}

// Regular Testing functions
// Test images are found within the test_images folder

func TestLoadImage(t *testing.T) {
	testImages := []string{
		"test_images/baby_fox.jpg",
		"test_images/fox_mountain.jpg",
		"test_images/hawksbill_sea_turtle.jpg",
		"test_images/pangolin.jpg",
	}
	ch := loadImage(testImages)

	for job := range ch {
		if job.Image == nil {
			t.Errorf("Image was not loaded for %s", job.InputPath)
		}
	}
}

// Ensure image is correctly resized to 500x500
func TestResizeSpecificImage(t *testing.T) {
	originalImg := imageprocessing.ReadImage("test_images/baby_fox.jpg")
	if originalImg == nil {
		t.Fatal("Failed to load the image")
	}

	resizedImg := imageprocessing.Resize(originalImg)
	bounds := resizedImg.Bounds()
	if width, height := bounds.Dx(), bounds.Dy(); width != 500 || height != 500 {
		t.Errorf("Image was not resized correctly: got dimensions %dx%d, want 500x500", width, height)
	}
}
