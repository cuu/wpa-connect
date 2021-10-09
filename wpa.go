package wpaconnect

import "github.com/cuu/wpa-connect/internal/log"

func SetSilentMode() {
	log.SetSilentMode()
}

func SetInfoMode() {
	log.SetInfoMode()
}

func SetVerboseMode() {
	log.SetVerboseMode()
}

func SetDebugMode() {
	log.SetDebugMode()
}
