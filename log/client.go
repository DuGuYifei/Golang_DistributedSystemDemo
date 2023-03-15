package log

import (
	"bytes"
	"distributed/registry"
	"fmt"
	stlog "log"
	"net/http"
)

type clientLogger struct {
	url string
}

func (cl clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer([]byte(data))
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to send log message. Service responded with %d - %s", res.StatusCode, res.Status)
	}
	return len(data), nil
}

func SetClientLogger(serviceURL string, clientService registry.ServiceName) clientLogger {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stlog.SetFlags(0)
	logger := &clientLogger{url: serviceURL}
	stlog.SetOutput(logger)
	return *logger
}
