package file_logger

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/nool01/velog/internal/logger_config"
)

func startFileSwapper() {
	wg.Add(1)
	go swapper()
}

func swapper() {
	wg.Done()

	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	time.Sleep(next.Sub(now))

	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for {
		mutex.Lock()
		writer.Flush()
		if file != nil {
			err := file.Close()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
		swapFile()
		writer = bufio.NewWriter(file)
		mutex.Unlock()
		<-ticker.C
	}
}

func swapFile() {
	directory := strings.TrimSuffix(logger_config.ApiConfig.File.Directory, "/")

	if err := os.MkdirAll(directory, 0755); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	now := time.Now().Format("2006-01-02")
	filePath := directory + "/" + logger_config.ApiConfig.File.FileName + "(" + now + ")" + "." + logger_config.ApiConfig.File.FileExtension

	newFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	file = newFile
}
