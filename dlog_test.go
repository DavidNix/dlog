package dlog

import (
	"bytes"
	"errors"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetPrefix(t *testing.T) {
	buf := bytes.NewBufferString("")
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)
	defer log.SetPrefix("")

	var prefixTests = []struct {
		prefix   string
		expected string
	}{
		{"", "Test\n"},
		{"Prefix", "[Prefix] Test\n"},
		{"Error Happened", "[Error Happened] Test\n"},
	}

	for _, test := range prefixTests {
		SetPrefix(test.prefix)
		Print("Test")

		assert.Equal(t, buf.String(), test.expected)

		buf.Reset()
	}
}

func Test_PanicIf_doesNotPanicIfErrAbsent(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Error("panic ocurred")
		}
	}()
	PanicIf(nil)
}

func Test_PanicIf_PanicsIfErrPresent(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("did not panic")
		}
	}()
	PanicIf(errors.New("an error"))
}

func Test_Debug_logsIffDebugIsTrue(t *testing.T) {
	buf := bytes.NewBufferString("")
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)

	SetDebug(false)
	Debug("I should not log")

	if buf.String() != "" {
		t.Fatalf("expected no logs, got %q", buf.String())
	}

	SetDebug(true)
	Debug("I should log")

	assert.Equal(t, buf.String(), "[DEBUG] I should log\n")
}

func Test_Debugf_logsIffDebugIsTrue(t *testing.T) {
	buf := bytes.NewBufferString("")
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)

	SetDebug(false)
	Debugf("I should not log")
	if buf.String() != "" {
		t.Fatalf("expected no logs, got %q", buf.String())
	}

	SetDebug(true)
	Debugf("I should %s", "log")

	assert.Equal(t, buf.String(), "[DEBUG] I should log\n")
}

func Test_Debugln_logsIffDebugIsTrue(t *testing.T) {
	buf := bytes.NewBufferString("")
	log.SetOutput(buf)
	defer log.SetOutput(os.Stderr)

	SetDebug(false)
	Debugln("I should not log")
	if buf.String() != "" {
		t.Fatalf("expected no logs, got %q", buf.String())
	}

	SetDebug(true)
	Debugln("I should", "log")

	assert.Equal(t, buf.String(), "[DEBUG] I should log\n")
}
