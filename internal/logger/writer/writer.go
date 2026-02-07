package writer

import (
	"bytes"
	"sync"
	"time"

	"github.com/nool01/velog/internal/buffer"
	"github.com/nool01/velog/internal/logger/console_logger"
	"github.com/nool01/velog/internal/logger/file_logger"
	"github.com/nool01/velog/internal/logger_config"
)

var (
	logWg            sync.WaitGroup
	LogBufferChannel = make(chan *bytes.Buffer, 5000)
	logQuit          = make(chan struct{})
	once             sync.Once
)

func StartWriter() {
	logWg.Add(1)
	go writer()
}

func writer() {
	defer logWg.Done()

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case buf := <-LogBufferChannel:
			writeToLoggers(buf)
		case <-ticker.C:
			invokeFlush()
		case <-logQuit:
			for {
				select {
				case buf := <-LogBufferChannel:
					writeToLoggers(buf)
				default:
					invokeFlush()
					return
				}
			}
		}
	}
}

func writeToLoggers(buf *bytes.Buffer) {
	data := buf.Bytes()
	if logger_config.ApiConfig.Console.Enabled {
		console_logger.Write(data)
	}
	if logger_config.ApiConfig.File.Enabled {
		file_logger.Write(data)
	}
	buffer.PutBuffer(buf)
}

func invokeFlush() {
	if logger_config.ApiConfig.Console.Enabled {
		console_logger.InvokeFlush()
	}
	if logger_config.ApiConfig.File.Enabled {
		file_logger.InvokeFlush()
	}
}

func StopWriter() {
	once.Do(func() {
		close(logQuit)
	})
	invokeFlush()
	logWg.Wait()
}
