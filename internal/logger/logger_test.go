package logger

import (
	"bytes"
	"strings"
	"testing"

	"github.com/nool01/velog/internal/buffer"
	"github.com/nool01/velog/internal/format"
	"github.com/nool01/velog/internal/logger/writer"
	"github.com/nool01/velog/internal/logger_config"
	"github.com/nool01/velog/pkg/velog/velog_config"
)

func TestLogger_Log(t *testing.T) {
	format.LogFormatTokens = &format.LogFormat{0, 5, 1, 5, 2, 5}
	logger_config.LoggerFuncConfig = &logger_config.LoggerFunc{
		Timestamp: false,
		Caller:    false,
	}

	format.L["s"] = []byte(" | ")

	ch := make(chan *bytes.Buffer, 1)
	writer.LogBufferChannel = ch

	l := &Logger{}

	l.Info("TestLogger_Log", "Hello World")

	select {
	case buf := <-ch:
		got := buf.String()
		want := "TestLogger_Log | Hello World | INFO | \n"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		buffer.PutBuffer(buf)
	default:
		t.Fatalf("log buffer channel is empty")
	}
}

func TestLogger_Debug(t *testing.T) {
	format.LogFormatTokens = &format.LogFormat{0, 5, 1, 5, 2, 5}
	logger_config.LoggerFuncConfig = &logger_config.LoggerFunc{
		Timestamp: false,
		Caller:    false,
	}

	format.L["s"] = []byte(" | ")
	l := &Logger{}

	writer.LogBufferChannel = make(chan *bytes.Buffer, 1)
	logger_config.ApiConfig = &velog_config.Config{
		Debug: false,
	}

	l.Debug("TestLogger_Debug", "Hello World")
	select {
	case <-writer.LogBufferChannel:
		t.Fatalf("buffer should be empty when debug is off")
	default:
	}

	logger_config.ApiConfig = &velog_config.Config{
		Debug: true,
	}
	l.Debug("TestLogger_Debug", "Hello World")
	select {
	case buf := <-writer.LogBufferChannel:
		if !strings.Contains(buf.String(), "TestLogger_Debug | Hello World | DEBUG") {
			t.Errorf("unexpected log buffer %s", buf.String())
		}
		buffer.PutBuffer(buf)
	default:
		t.Fatalf("log buffer channel is empty")
	}
}
