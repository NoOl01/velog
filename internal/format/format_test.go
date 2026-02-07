package format

import (
	"errors"
	"reflect"
	"testing"

	"github.com/nool01/velog/internal/logger_config"
	"github.com/nool01/velog/pkg/velog/velog_errs"
)

func TestParseFormat(t *testing.T) {
	tests := []struct {
		format string
		want   LogFormat
	}{
		{"${name} ${s} ${content} ${s} ${level} ${s} ${timestamp} ${s} ${caller}", LogFormat{0, 5, 1, 5, 2, 5, 3, 5, 4}},
		{"${name} ${s} ${level} ${s} ${timestamp} ${s}", LogFormat{0, 5, 2, 5, 3, 5}},
	}

	for _, test := range tests {
		t.Run(test.format, func(t *testing.T) {
			got := LogFormat{}
			if err := ParseFormat(test.format, &got); err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got %+v, want %+v", got, test.want)
			}
		})
	}
}

func TestParseFormat_UnknownFormat(t *testing.T) {
	format := "${name} ${s} ${unknown} ${s}"

	logFormat := LogFormat{}
	err := ParseFormat(format, &logFormat)
	if !errors.Is(err, velog_errs.UnknownFormat) {
		t.Fatal("expected UnknownFormat error")
	}
}

func TestAutoDisable(t *testing.T) {
	tests := []struct {
		format string
		want   []bool
	}{
		{"${name} ${s} ${content} ${s} ${level} ${s} ${timestamp} ${s} ${caller}", []bool{true, true}},
		{"${name} ${s} ${level} ${s} ${timestamp} ${s}", []bool{true, false}},
		{"${name} ${s} ${level} ${s} ${caller} ${s}", []bool{false, true}},
	}

	for _, test := range tests {
		t.Run(test.format, func(t *testing.T) {
			LogFormatTokens = &LogFormat{}
			logger_config.LoggerFuncConfig = &logger_config.LoggerFunc{
				Timestamp: false,
				Caller:    false,
			}

			if err := ParseFormat(test.format, LogFormatTokens); err != nil {
				t.Fatal(err)
			}

			AutoDisable()

			if logger_config.LoggerFuncConfig.Timestamp != test.want[0] || logger_config.LoggerFuncConfig.Caller != test.want[1] {
				t.Errorf("got %+v, want %+v", logger_config.LoggerFuncConfig.Timestamp, test.want[0])
			}
		})
	}
}
