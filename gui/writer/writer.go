package writer

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

type LogWriter struct{ io.Writer }

func NewLogWriter() *LogWriter {
	return &LogWriter{log.Writer()}
}

func (w *LogWriter) Enable() *LogWriter {
	w.Writer = os.Stdout

	return w
}
func (w *LogWriter) Disable() *LogWriter {
	w.Writer = ioutil.Discard

	return w
}
