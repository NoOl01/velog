package console_logger

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var (
	writer *bufio.Writer
	mutex  sync.Mutex
	count  int
)

const maxBatchSize = 1024

func InitWriter() {
	writer = bufio.NewWriter(os.Stdout)
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
		InvokeFlush()
	}
}

func InvokeFlush() {
	mutex.Lock()
	defer mutex.Unlock()
	flushLocked()
}

func flushLocked() {
	if count > 0 {
		if err := writer.Flush(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		count = 0
	}
}
