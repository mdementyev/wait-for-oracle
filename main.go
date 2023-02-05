package main

import (
	"fmt"
	"github.com/jnovack/flag"
	ora "github.com/sijms/go-ora/v2"
	"os"
	"time"
)

func connect(dsn string) (bool, error) {
	DB, err := ora.NewConnection(dsn)
	if err != nil {
		return false, err
	}

	err = DB.Open()
	if err != nil {
		return false, err
	}

	defer DB.Close()

	stmt := ora.NewStmt("SELECT 'X' from DUAL", DB)

	defer stmt.Close()

	rows, err := stmt.Query(nil)
	if err != nil {
		return false, err
	}

	defer rows.Close()

	return true, nil
}

func tryConnectIndefinitely(dsn string) {
	for {
		_, err := connect(dsn)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		break
	}
}

func main() {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "ORACLE", 0)

	user := fs.String("user", "", "Oracle database user")
	password := fs.String("password", "", "Oracle database password")
	host := fs.String("host", "", "Oracle database host")
	port := fs.String("port", "1521", "Oracle database port")
	sid := fs.String("sid", "", "Oracle database SID")
	timeout := fs.Int("timeout", 0, "Timeout in seconds, zero means no timeout")
	fs.Parse(os.Args[1:])

	if *user == "" || *password == "" || *host == "" || *port == "" || *sid == "" {
		fs.Usage()
		println("Arguments can be set via environment variables (e.g. ORACLE_USER)")
		return
	}

	dsn := "oracle://" + *user + ":" + *password + "@" + *host + ":" + *port + "/" + *sid

	if (*timeout) > 0 {
		c1 := make(chan bool, 1)
		go func() {
			tryConnectIndefinitely(dsn)

			c1 <- true
		}()

		select {
		case <-c1:
			fmt.Println("Connected")
			os.Exit(0)
		case <-time.After(time.Duration(*timeout) * time.Second):
			fmt.Println("Timeout")
			os.Exit(1)
		}
	} else {
		tryConnectIndefinitely(dsn)

		fmt.Println("Connected")
		os.Exit(0)
	}
}
