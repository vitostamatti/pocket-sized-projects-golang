package main

import (
	"os"
	"time"

	"github.com/vitostamatti/pocketlog/pocketlog"
)

func main() {
	lgr := pocketlog.New(pocketlog.LevelDebug, pocketlog.WithOutput(os.Stdout))
	lgr.Infof("A little copying is better than little depndency.")
	lgr.Errorf("Errors are values. Documentation is for %s.", "users")
	lgr.Infof("Hallo, %d %v", 2022, time.Now())
}
