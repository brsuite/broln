package kvdb

import (
	"github.com/btcsuite/btclog"
	"github.com/brsuite/broln/kvdb/postgres"
)

// log is a logger that is initialized as disabled.  This means the package will
// not perform any logging by default until a logger is set.
var log = btclog.Disabled

// UseLogger uses a specified Logger to output package logging info.
func UseLogger(logger btclog.Logger) {
	log = logger

	postgres.UseLogger(log)
}
