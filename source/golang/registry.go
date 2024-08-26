package golang

import (
	"fmt"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/golang-migrate/migrate/v4/source"
)

type Registry struct {
	migrations   *source.Migrations
	migrationsMu sync.Mutex
}

var defaultRegistry = NewRegistry()

func (m *Registry) Register(f source.Func, filename string) {
	migration, err := source.DefaultParse(filepath.Base(filename))
	if err != nil {
		panic(fmt.Sprintf("can't parse filename %q: %s", filename, err))
	}
	migration.Func = f

	m.migrationsMu.Lock()

	m.migrations.Append(migration)
	m.migrationsMu.Unlock()
}

func NewRegistry() *Registry {
	return &Registry{
		migrations:   source.NewMigrations(),
		migrationsMu: sync.Mutex{},
	}
}

// Register globally registers a registry.
func Register(f source.Func) {
	_, filename, _, _ := runtime.Caller(1)

	defaultRegistry.Register(f, filename)
}
