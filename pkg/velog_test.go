package pkg

import (
	"bytes"
	"os"
	"testing"
	"time"

	"github.com/nool01/velog/pkg/velog"
	"github.com/nool01/velog/pkg/velog/velog_config"
)

func TestLoggerIntegration(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() { os.Stdout = oldStdout }()

	config := &velog_config.Config{
		Format:    "${name} ${s} ${content} ${s} ${level} ${s}",
		Separator: " | ",
		Debug:     true,
		Console: velog_config.ConsoleConfig{
			Enabled: true,
		},
	}

	log := velog.Start(config)
	defer velog.Stop()

	log.Info("TestLogger", "Hello World")
	time.Sleep(1 * time.Second)
	w.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	got := buf.String()
	want := "TestLogger | Hello World | INFO | \n"

	if got != want {
		t.Errorf("integration test failed: got %q, want %q", got, want)
	}
}
