package services

import "log"

// @todo very basic type should be refactored

var Mode string

func Initialize(mode string) {
	Mode = mode
}

func LogDeb(v ...interface{}) {
	switch Mode {
	case "debug":
		log.Printf("%+v\n", v...)
	default:
	}
}

func LogDev(v ...interface{}) {
	switch Mode {
	case "debug":
		log.Printf("%+v\n", v...)
	case "dev":
		log.Printf("%+v\n", v...)
	default:
	}
}
