package main

import (
	"database/sql"
	config "delivery/configs"
	"delivery/internal/server"
	myLog "delivery/logs"
	"fmt"
	"os"
	"time"
)

func main() {
	l := myLog.LogInit()
	cfg := config.NewConfig(l)
	conn, err := sql.Open(
		cfg.Driver,
		cfg.DataSourceName,
	)

	defer conn.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = conn.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn.SetMaxIdleConns(2)
	conn.SetMaxIdleConns(10)
	conn.SetConnMaxLifetime(10 * time.Second)

	s := server.NewServer(cfg, conn)

	err = s.Start()
	if err != nil {
		cfg.Logger.Error("Can't start server with error", err)
		os.Exit(1)
	}
}
