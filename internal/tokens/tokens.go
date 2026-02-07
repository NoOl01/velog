package tokens

type TokenType int
type JsonToken int

const (
	TokenName TokenType = iota
	TokenContent
	TokenLevel
	TokenTimestamp
	TokenCaller
	TokenSeparator
)
