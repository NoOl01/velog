package console_logger

import (
	"bufio"
	"fmt"
	"os"
)

var (
	writer *bufio.Writer
	count  int
)

const maxBatchSize = 1024

func InitWriter() {
	writer = bufio.NewWriter(os.Stdout)
	count = 0
}

func Write(buf []byte) {
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
	if count > 0 {
		err := writer.Flush()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		count = 0
	}
}
