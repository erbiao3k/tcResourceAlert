package config

import (
	"os"
	"strings"
)

var (
	SecretId  = SecretIdKey[0]
	SecretKey = SecretIdKey[1]
	SecretIdKey = strings.Split(os.Getenv("SecretIdKey"), ",")
	WecomRobot = strings.Split(os.Getenv("WecomRobot"), ",")
)
