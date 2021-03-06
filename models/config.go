package models

import (
	"time"
)

type Config struct {
	WriteWait                 time.Duration
	ReadWait                  time.Duration
	PingPeriod                time.Duration
	MaxMessageSize            int64
	BroadcastMessageQueueSize int64
	DBPath                    string
	HostName                  string
	FullyQualifiedHost        string
}

var GlobalConfig *Config

func InitializeConfig() {
	GlobalConfig = &Config{
		DBPath:             "localhost:3306",
		HostName:           "localhost:8080",
		FullyQualifiedHost: "http://52.23.176.27:8080/",
	}
}
