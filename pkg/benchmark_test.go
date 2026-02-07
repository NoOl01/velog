package pkg

import (
	"testing"

	"github.com/nool01/velog/pkg/velog"
	"github.com/nool01/velog/pkg/velog/velog_config"
)

func BenchmarkLogToConsole(b *testing.B) {
	config := &velog_config.Config{
		Console: velog_config.ConsoleConfig{
			Enabled: true,
		},
		Format:    "${name} ${s} ${content} ${s} ${level} ${s} ${timestamp} ${s} ${caller}",
		Separator: " || ",
	}
	log := velog.Start(config)

	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("test", "hello world")
		}
	})

	velog.Stop()
}

func BenchmarkLogToConsoleWithoutCaller(b *testing.B) {
	config := &velog_config.Config{
		Console: velog_config.ConsoleConfig{
			Enabled: true,
		},
		Format:    "${name} ${s} ${content} ${s} ${level} ${s} ${timestamp}",
		Separator: " || ",
	}
	log := velog.Start(config)

	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("test", "hello world")
		}
	})

	velog.Stop()
}

func BenchmarkLogToConsoleWithoutCallerAndTimestamp(b *testing.B) {
	config := &velog_config.Config{
		Console: velog_config.ConsoleConfig{
			Enabled: true,
		},
		Format:    "${name} ${s} ${content} ${s} ${level}",
		Separator: " || ",
	}
	log := velog.Start(config)

	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("test", "hello world")
		}
	})

	velog.Stop()
}

func BenchmarkLogToFileWithoutCallerAndTimestamp(b *testing.B) {
	config := &velog_config.Config{
		File: velog_config.FileConfig{
			Enabled:       true,
			Directory:     "./test/",
			FileName:      "test",
			FileExtension: "txt",
		},
		Format:    "${name} ${s} ${content} ${s} ${level}",
		Separator: " || ",
	}
	log := velog.Start(config)

	b.ReportAllocs()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			log.Info("test", "hello world")
		}
	})

	velog.Stop()
}

func BenchmarkRealisticUsage(b *testing.B) {
	config := &velog_config.Config{
		File: velog_config.FileConfig{
			Enabled:       true,
			Directory:     "./test/",
			FileName:      "realistic",
			FileExtension: "txt",
		},
		Format:    "${name} ${s} ${content} ${s} ${level}",
		Separator: " || ",
	}

	log := velog.Start(config)
	defer velog.Stop()

	b.ReportAllocs()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		log.Info("service", "processing request")
	}
}
