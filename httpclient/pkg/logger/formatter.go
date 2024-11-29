package logger

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

type MyLoggerFormatter struct {
}

func (f *MyLoggerFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields)
	for k, v := range entry.Data {
		data[k] = v
	}
	data["time"] = entry.Time.Format("2006-01-02 15:04:05")
	data["level"] = entry.Level.String()
	data["msg"] = entry.Message
	body, _ := json.MarshalIndent(data, "", "    ")
	content := fmt.Sprintf("%s\n", string(body))
	return []byte(content), nil
}
