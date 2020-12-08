package main

import (
	"database/sql"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/scc/go-common/log"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	migrate "github.com/rubenv/sql-migrate"

	"order-manager/dbschema"
	"order-manager/pkg/api/v3/delivery/broker"
	"order-manager/pkg/api/v3/delivery/rest"
	"order-manager/pkg/cfg"
)

func RunAPI(api rest.API) error {
	return api.Run()
}

func RunBroker(brk broker.Broker, ll []broker.Listener) error {
	brk.AddListeners(ll...)
	return brk.Run()
}

const (
	migrateUp   string = "up"
	migrateDown string = "down"
)

var directions = map[string]migrate.MigrationDirection{
	migrateUp:   migrate.Up,
	migrateDown: migrate.Down,
}

func Migrate(db *sql.DB, c *cfg.Cfg) error {
	if !c.Storage.Postgres.AutoMigrate {
		return nil
	}

	if _, ok := directions[c.Storage.Postgres.MigrationDirection]; !ok {
		return fmt.Errorf("invalid direction: want [%s|%s], get `%s`", migrateUp, migrateDown, c.Storage.Postgres.MigrationDirection)
	}

	s := migrate.AssetMigrationSource{
		Asset:    dbschema.Asset,
		AssetDir: dbschema.AssetDir,
		Dir:      "migrations",
	}

	if _, err := migrate.Exec(db, "postgres", s, directions[c.Storage.Postgres.MigrationDirection]); err != nil {
		return errors.Wrap(err, "unable to migrate")
	}

	return nil
}

func Shutdown(l log.Logger, db *sql.DB, brk broker.Broker) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTERM)
	sig := <-ch
	l.Infof("received signal %q", sig)

	db.Close()
	brk.Close()
	os.Exit(0)
}
