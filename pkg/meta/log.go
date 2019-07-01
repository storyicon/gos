package meta

import (
	log "github.com/sirupsen/logrus"
)

// FmtFormatter is used to print clean output
type FmtFormatter struct{}

// Format implements the interface of logrus Formater
func (f *FmtFormatter) Format(entry *log.Entry) ([]byte, error) {
	return []byte("gos: " + entry.Message + "\r\n"), nil
}

func init() {
	log.SetFormatter(&FmtFormatter{})
}
