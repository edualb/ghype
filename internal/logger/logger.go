package logger

import (
	"log"
	"os"
	"sync"
)

type Logger struct {
	stderr *os.File
	stdout *os.File

	mu *sync.Mutex
}

func New() *Logger {
	return &Logger{
		stderr: os.Stderr,
		stdout: os.Stdout,

		mu: &sync.Mutex{},
	}
}

func (l *Logger) Errof(msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.stderr == nil {
		l.stderr = os.Stderr
	}
	logger := log.New(l.stderr, "", log.LstdFlags)
	w := logger.Writer()
	w.Write([]byte(msg))
}

func (l *Logger) Infof(msg string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.stdout == nil {
		l.stdout = os.Stdout
	}
	logger := log.New(l.stdout, "", log.LstdFlags)
	w := logger.Writer()
	w.Write([]byte(msg))
}
