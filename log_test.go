package log

import (
	"log"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	log.SetOutput(&OutWriter{})
	m.Run()
}

var testOutput string

func TestInitFromString(t *testing.T) {
	if err := InitFromString(""); err.Error() != "unable to init logger: unknown level ''" {
		t.Error("expected error when initialising with empty string but got none")
	}

	if err := InitFromString("unknown"); err.Error() != "unable to init logger: unknown level 'unknown'" {
		t.Errorf("expected error when initialising with 'unknown' but got: %s", err.Error())
	}

	if InitFromString("error"); level != LevelError {
		t.Error("failed to initialise to level 'error'")
	}

	if InitFromString("info"); level != LevelInfo {
		t.Error("failed to initialise to level 'info'")
	}

	if InitFromString("warn"); level != LevelWarn {
		t.Error("failed to initialise to level 'warn'")
	}

	if InitFromString("debug"); level != LevelDebug {
		t.Error("failed to initialise to level 'debug'")
	}

	if InitFromString("trace"); level != LevelTrace {
		t.Error("failed to initialise to level 'trace'")
	}

	if InitFromString("none"); level != LevelNone {
		t.Error("failed to initialise to level 'none'")
	}
}

func TestInit(t *testing.T) {
	for _, l := range []int{-1, 6} {
		if err := Init(l); err == nil {
			t.Errorf("expected error when level was '%d' but got none", l)
		}
	}
}

func TestTrace(t *testing.T) {
	testCases := map[int][]string{
		LevelNone:  []string{"Trace me", ""},
		LevelError: []string{"Trace me", ""},
		LevelInfo:  []string{"Trace me", ""},
		LevelWarn:  []string{"Trace me", ""},
		LevelDebug: []string{"Trace me", ""},
		LevelTrace: []string{"Trace me", "Trace me"},
	}
	testLogMethod(t, func(s string) { Trace(s) }, testCases)
}

func TestDebug(t *testing.T) {
	testCases := map[int][]string{
		LevelNone:  []string{"Debug me", ""},
		LevelError: []string{"Debug me", ""},
		LevelInfo:  []string{"Debug me", ""},
		LevelWarn:  []string{"Debug me", ""},
		LevelDebug: []string{"Debug me", "Debug me"},
		LevelTrace: []string{"Debug me", "Debug me"},
	}
	testLogMethod(t, func(s string) { Debug(s) }, testCases)
}

func TestWarn(t *testing.T) {
	testCases := map[int][]string{
		LevelNone:  []string{"Warn me", ""},
		LevelError: []string{"Warn me", ""},
		LevelInfo:  []string{"Warn me", ""},
		LevelWarn:  []string{"Warn me", "Warn me"},
		LevelDebug: []string{"Warn me", "Warn me"},
		LevelTrace: []string{"Warn me", "Warn me"},
	}
	testLogMethod(t, func(s string) { Warn(s) }, testCases)
}

func TestInfo(t *testing.T) {
	testCases := map[int][]string{
		LevelNone:  []string{"Info me", ""},
		LevelError: []string{"Info me", ""},
		LevelInfo:  []string{"Info me", "Info me"},
		LevelWarn:  []string{"Info me", "Info me"},
		LevelDebug: []string{"Info me", "Info me"},
		LevelTrace: []string{"Info me", "Info me"},
	}
	testLogMethod(t, func(s string) { Info(s) }, testCases)
}

func TestError(t *testing.T) {
	testCases := map[int][]string{
		LevelNone:  []string{"Error me", ""},
		LevelError: []string{"Error me", "Error me"},
		LevelInfo:  []string{"Error me", "Error me"},
		LevelWarn:  []string{"Error me", "Error me"},
		LevelDebug: []string{"Error me", "Error me"},
		LevelTrace: []string{"Error me", "Error me"},
	}
	testLogMethod(t, func(s string) { Error(s) }, testCases)
}

func testLogMethod(t *testing.T, f func(string), cases map[int][]string) {
	for lev, v := range cases {
		clear()
		Init(lev)
		f(v[0])
		if !strings.Contains(testOutput, v[1]) {
			t.Errorf("at level '%d', expected output '%s' but got '%s'", lev, v[1], testOutput)
		}
	}
}

// OutWriter implements the io.Writer interface. Used to capture test output
type OutWriter struct{}

func (o *OutWriter) Write(p []byte) (int, error) {
	testOutput = string(p)
	return 0, nil
}

func clear() {
	testOutput = ""
}
