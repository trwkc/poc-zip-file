package main

import (
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	compress()
	zipper()
}

func compress() {
	inputFile, err := os.Open("test.xlsx")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// --- compress file gzip ---
	gzipWriter, err := os.Create("test_1.xlsx.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer gzipWriter.Close()

	gzWriter := gzip.NewWriter(gzipWriter)
	defer gzWriter.Close()

	_, err = io.Copy(gzWriter, inputFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("compress file created successfully")
	// --- compress file gzip ---
}

func zipper() {
	inputFile, err := os.Open("test.xlsx")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()
	// --- zip file ---
	archive, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	zWriter := zip.NewWriter(archive)
	defer zWriter.Close()

	wz, err := zWriter.Create("test_2.xlsx")
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(wz, inputFile); err != nil {
		panic(err)
	}
	zWriter.Close()
	fmt.Println("archive file created successfully")
	// --- zip file ---
}
