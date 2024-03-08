package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// Пример реализации функции saveToFile
func saveToFile(fileNames []string, outFile string) error {
	file, err := os.Create(outFile)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, fileName := range fileNames {
		_, err = file.WriteString(fileName + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// Реализовать функцию получения всех файлов из директории
func getFiles(dir string) ([]string, error) {
	const readAllFiles = 0

	directory, err := os.Open(dir)
	if err != nil {
		return nil, err
	}

	defer directory.Close()

	files, err := directory.Readdir(readAllFiles)
	if err != nil {
		return nil, err
	}

	var fileNames []string

	for _, file := range files {
		if file.IsDir() {
			val, err := getFiles(dir + "/" + file.Name())
			if err != nil {
				return nil, err
			}

			fileNames = append(fileNames, val...)

			continue
		}

		fileNames = append(fileNames, dir+"/"+file.Name())
	}

	return fileNames, nil
}

// Реализовать функцию фильтрации переданных названий файлов
func filterFiles(files []string, filter string) ([]string, error) {
	if filter == "" {
		return files, nil
	}

	fileNames := make([]string, 0)

	if strings.HasPrefix(filter, "*") {
		for _, file := range files {
			if strings.HasSuffix(file, filter[1:]) {
				fileNames = append(fileNames, file)
			}
		}

		return fileNames, nil
	}

	for _, file := range files {
		if strings.Contains(file, filter) {
			fileNames = append(fileNames, file)
		}
	}

	return fileNames, nil
}

func main() {
	files, err := getFiles("./cmd")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(files)
}
