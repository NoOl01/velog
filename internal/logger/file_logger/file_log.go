package file_logger

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var (
	wg     sync.WaitGroup
	writer *bufio.Writer
	file   *os.File
	mutex  sync.Mutex
	count  int
)

const maxBatchSize = 262144

func InitWriter() {
	swapFile()
	writer = bufio.NewWriter(file)
	startFileSwapper()
	count = 0
}

func Write(buf []byte) {
	mutex.Lock()
	defer mutex.Unlock()

	n, err := writer.Write(buf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	count += n

	if count >= maxBatchSize {
		writer.Flush()
		count = 0
	}
}

func InvokeFlush() {
	mutex.Lock()
	defer mutex.Unlock()

	writer.Flush()
	count = 0
}
