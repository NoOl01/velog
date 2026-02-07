package format

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nool01/velog/internal/logger_config"
	"github.com/nool01/velog/internal/tokens"
	"github.com/nool01/velog/pkg/velog/velog_config"
	"github.com/nool01/velog/pkg/velog/velog_errs"
)

var LevelToBytes = map[velog_config.LogLevel][]byte{
	velog_config.INFO:    []byte("INFO"),
	velog_config.WARNING: []byte("WARNING"),
	velog_config.ERROR:   []byte("ERROR"),
	velog_config.DEBUG:   []byte("DEBUG"),
	velog_config.PANIC:   []byte("PANIC"),
}

var L = map[string][]byte{}

type LogFormat []tokens.TokenType

var LogFormatTokens *LogFormat

func Format(format, separator string) {
	LogFormatTokens = &LogFormat{}

	if separator == "" {
		separator = " | "
	}

	L["s"] = []byte(separator)

	if err := ParseFormat(format, LogFormatTokens); err != nil {
		fmt.Println("ParseFormat(format) failed:", err)
	}
}

func ParseFormat(format string, logFormat *LogFormat) error {
	var tokenRegex = regexp.MustCompile(`\$\{([a-zA-Z0-9_]+)}`)
	formatTokens := tokenRegex.FindAllStringSubmatch(format, -1)

	for _, token := range formatTokens {
		token[1] = strings.TrimSpace(token[1])
		if token[1] == "" {
			continue
		}

		switch token[1] {
		case "name":
			*logFormat = append(*logFormat, tokens.TokenName)
		case "content":
			*logFormat = append(*logFormat, tokens.TokenContent)
		case "level":
			*logFormat = append(*logFormat, tokens.TokenLevel)
		case "timestamp":
			*logFormat = append(*logFormat, tokens.TokenTimestamp)
		case "caller":
			*logFormat = append(*logFormat, tokens.TokenCaller)
		case "s":
			*logFormat = append(*logFormat, tokens.TokenSeparator)
		default:
			return velog_errs.UnknownFormat
		}
	}

	return nil
}

func AutoDisable() {
	for _, token := range *LogFormatTokens {
		switch token {
		case tokens.TokenTimestamp:
			logger_config.LoggerFuncConfig.Timestamp = true
		case tokens.TokenCaller:
			logger_config.LoggerFuncConfig.Caller = true
		default:
			continue
		}
	}
}
