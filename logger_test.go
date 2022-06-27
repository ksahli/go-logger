package logger_test

import (
	"bytes"
	"testing"

	"github.com/ksahli/go-logger"
)

type Recorder struct {
	bytes []byte
}

func (r *Recorder) Write(p []byte) (int, error) {
	r.bytes = p
	return len(p), nil
}

func TestDebug(t *testing.T) {

	for _, level := range []byte{logger.Debug} {
		stdout, stderr := new(Recorder), new(Recorder)
		sut := logger.New(level, stdout, stderr, "component")
		sut.Debug("this is a test message %s", "hello")

		got := stdout.bytes
		want := []string{"DEBUG", "component", "this is a test message", "hello"}
		for _, w := range want {
			if !bytes.Contains(got, []byte(w)) {
				msg := "want '%s' to contain '%s', but it was not the case"
				t.Fatalf(msg, got, w)
			}
		}

		if len(stderr.bytes) > 0 {
			msg := "want nothing to be written to stderr, got '%s'"
			t.Fatalf(msg, stderr.bytes)
		}
	}

	for _, level := range []byte{logger.Trace, logger.Error} {
		stdout, stderr := new(Recorder), new(Recorder)
		sut := logger.New(level, stdout, stderr, "component")
		sut.Debug("this is a test message %s", "hello")

		if len(stdout.bytes) > 0 {
			msg := "want nothing to be written to stdout, got '%s'"
			t.Fatalf(msg, stderr.bytes)
		}

		if len(stderr.bytes) > 0 {
			msg := "want nothing to be written to stderr, got '%s'"
			t.Fatalf(msg, stderr.bytes)
		}
	}
}

func TestTrace(t *testing.T) {

	for _, level := range []byte{logger.Debug, logger.Trace} {
		stdout, stderr := new(Recorder), new(Recorder)
		sut := logger.New(level, stdout, stderr, "component")
		sut.Trace("this is a test message %s", "hello")

		got := stdout.bytes
		want := []string{"TRACE", "component", "this is a test message", "hello"}
		for _, w := range want {
			if !bytes.Contains(got, []byte(w)) {
				msg := "want '%s' to contain '%s', but it was not the case"
				t.Fatalf(msg, got, w)
			}
		}

		if len(stderr.bytes) > 0 {
			msg := "want nothing to be written to stderr, got '%s'"
			t.Fatalf(msg, stderr.bytes)
		}
	}

	for _, level := range []byte{logger.Error} {
		stdout, stderr := new(Recorder), new(Recorder)
		sut := logger.New(level, stdout, stderr, "component")
		sut.Debug("this is a test message %s", "hello")

		if len(stdout.bytes) > 0 {
			msg := "want nothing to be written to stdout, got '%s'"
			t.Fatalf(msg, stderr.bytes)
		}

		if len(stderr.bytes) > 0 {
			msg := "want nothing to be written to stderr, got '%s'"
			t.Fatalf(msg, stderr.bytes)
		}
	}
}

func TestError(t *testing.T) {

	for _, level := range []byte{logger.Debug, logger.Trace, logger.Error} {
		stdout, stderr := new(Recorder), new(Recorder)
		sut := logger.New(level, stdout, stderr, "component")
		sut.Error("this is a test message %s", "hello")

		got := stderr.bytes
		want := []string{"ERROR", "component", "this is a test message", "hello"}
		for _, w := range want {
			if !bytes.Contains(got, []byte(w)) {
				msg := "want '%s' to contain '%s', but it was not the case"
				t.Fatalf(msg, got, w)
			}
		}

		if len(stdout.bytes) > 0 {
			msg := "want nothing to be written to stdout, got '%s'"
			t.Fatalf(msg, stderr.bytes)
		}
	}

}
