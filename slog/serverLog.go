package utils

import (
	slog "github.com/cihub/seelog"
)

var Slogger slog.LoggerInterface

func init() {
	var err error
	Slogger, err = slog.LoggerFromConfigAsFile("slog.xml")
	if err != nil {
		panic(err)
	}
}
