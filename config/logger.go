package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer // o que vai escrever dentro do logger
}

func NewLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)

	return &Logger{
		debug: log.New(writer, "🐛DEBUG: ", logger.Flags()),
		info: log.New(writer, "📃INFO: ", logger.Flags()),
		warning: log.New(writer, "⚠️ WARNING: ", logger.Flags()),
		err: log.New(writer, "⛔ERROR: ", logger.Flags()),
		writer: writer,
	}
}

//Create non-formatted Logs
//...interface{} => recebe qualquer coisa e qualqer quantidade de coisas
//(l *Logger) => significa quase que um metodo dentro de uma classe publica que
//pode ser acessado de outros lugares
func (l *Logger) Debug(values ...interface{})  {
	l.debug.Print(values...)
}
func (l *Logger) Info(values ...interface{})  {
	l.info.Print(values...)
}
func (l *Logger) Warning(values ...interface{})  {
	l.warning.Print(values...)
}
func (l *Logger) Error(values ...interface{})  {
	l.err.Print(values...)
}

// Create format Enabled Logs
func (l *Logger) Debugf(format string, values ...interface{})  {
	l.debug.Printf(format, values...)
}
func (l *Logger) Infof(format string, values ...interface{})  {
	l.info.Printf(format, values...)
}
func (l *Logger) Warningf(format string, values ...interface{})  {
	l.warning.Printf(format, values...)
}
func (l *Logger) Errfor(format string, values ...interface{})  {
	l.err.Printf(format, values...)
}