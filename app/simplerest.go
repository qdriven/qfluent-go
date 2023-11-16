package main

import (
	"github.com/qdriven/qfluent-go/app/simplerest"
	"log/slog"
)

func main() {
	slog.Info("start Auth Server")
	simplerest.RunRestGoAuthSqliteServer(9090)
}
