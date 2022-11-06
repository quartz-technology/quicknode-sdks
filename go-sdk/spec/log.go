package spec

type LogType string

const (
	TRANSFER LogType = "TRANSFER"
	ORDER    LogType = "ORDER"
	MINT     LogType = "MINT"
)
