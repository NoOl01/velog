package log_data

type LogData struct {
	Level     []byte
	Timestamp []byte
	Caller    []byte
	Separator []byte
	Name      []byte
	Content   []byte
}
