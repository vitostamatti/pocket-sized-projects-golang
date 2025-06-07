package pocketlog_test

import (
	"testing"

	"github.com/vitostamatti/pocketlog/pocketlog"
)

type testWriter struct {
	contents string
}

func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}

const (
	debugMessage = "Why write I still all one, ever the same,"
	infoMessage  = "And keep invention in a note weed,"
	errorMessage = "That every word doth almost tell my name,"
)

func TestLogger(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]struct {
		level    pocketlog.Level
		expected string
	}{
		"debug": {
			level: pocketlog.LevelDebug,
			expected: `{"level":"[DEBUG]","message":"` + debugMessage + "\"}\n" +
				`{"level":"[INFO]","message":"` + infoMessage + "\"}\n" +
				`{"level":"[ERROR]","message":"` + errorMessage + "\"}\n",
		},
		"info": {
			level: pocketlog.LevelInfo,
			expected: `{"level":"[INFO]","message":"` + infoMessage + "\"}\n" +
				`{"level":"[ERROR]","message":"` + errorMessage + "\"}\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: `{"level":"[ERROR]","message":"` + errorMessage + "\"}\n",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}
			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))
			testedLogger.Debugf(debugMessage)
			testedLogger.Infof(infoMessage)
			testedLogger.Errorf(errorMessage)
			if tw.contents != tc.expected {
				t.Errorf("expected %q, got %q", tc.expected, tw.contents)
			}
		})
	}
}
