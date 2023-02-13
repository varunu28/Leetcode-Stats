package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	readmeFileName = "README.md"
	repositoryUrl  = "https://github.com/varunu28/LeetCode-Java-Solutions/tree/master"
)

func Compute(directoryPath string) {
	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		completePath := directoryPath + "/" + e.Name()
		isDir, _ := isDirectory(completePath)
		if isDir && !strings.HasSuffix(e.Name(), ".git") {
			removeMarkdownFile(completePath)
			count := updateMarkdownAndGetStats(completePath, e.Name())
			fmt.Println(e.Name() + ": " + strconv.Itoa(count))
		}
	}
}

func updateMarkdownAndGetStats(path, category string) int {
	f, err := os.Create(path + "/" + readmeFileName)
	if err != nil {
		return 0
	}
	defer f.Close()
	writeHeader(f, category)
	entries, _ := os.ReadDir(path)
	count := 1
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		f.Write([]byte(buildFileRecord(e.Name(), category, count)))
		count += 1
	}
	return count
}

func buildFileRecord(name, category string, count int) string {
	fileName := name
	idx := strings.Index(fileName, ".java")
	fileNameForUrl := strings.Replace(fileName, " ", "%20", -1)
	return fmt.Sprintf("%d | [%s](%s/%s/%s)\n", count, fileName[0:idx], repositoryUrl, category, fileNameForUrl)
}

func writeHeader(f *os.File, category string) {
	f.Write([]byte("# " + category + " LeetCode Java Solutions \n"))
	f.Write([]byte("S.no | Coding Problem \n"))
	f.Write([]byte("--- | --- \n"))
}

func removeMarkdownFile(path string) {
	_, err := os.Stat(path + "/" + readmeFileName)
	if os.IsNotExist(err) {
		return
	}
	os.Remove(path + "/" + readmeFileName)
}

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), nil
}
