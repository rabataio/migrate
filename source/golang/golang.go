package golang

import (
	"io"
	"os"

	"github.com/golang-migrate/migrate/v4/source"
)

func init() {
	source.Register("golang", &Golang{})
}

func WithInstance(registry *Registry) source.Driver {
	return &Golang{
		registry: registry,
	}
}

// Golang is a source.Driver that reads migrations from a map of Migrations.
type Golang struct {
	registry *Registry
}

// Open implements source.Driver.
func (g *Golang) Open(_ string) (source.Driver, error) {
	g.registry = defaultRegistry
	return g, nil
}

func (g *Golang) Close() error {
	return nil
}

func (g *Golang) First() (version uint, err error) {
	v, ok := g.registry.migrations.First()
	if !ok {
		return 0, os.ErrNotExist
	}
	return v, nil
}

func (g *Golang) Prev(version uint) (prevVersion uint, err error) {
	v, ok := g.registry.migrations.Prev(version)
	if !ok {
		return 0, os.ErrNotExist
	}
	return v, nil
}

func (g *Golang) Next(version uint) (nextVersion uint, err error) {
	v, ok := g.registry.migrations.Next(version)
	if !ok {
		return 0, os.ErrNotExist
	}
	return v, nil
}

func (g *Golang) ReadUp(version uint) (r io.ReadCloser, f source.Func, identifier string, err error) {
	if m, ok := g.registry.migrations.Up(version); ok {
		return nil, m.Func, m.Identifier, nil
	}
	return nil, nil, "", os.ErrNotExist
}

func (g *Golang) ReadDown(version uint) (r io.ReadCloser, f source.Func, identifier string, err error) {
	if m, ok := g.registry.migrations.Down(version); ok {
		return nil, m.Func, m.Identifier, nil
	}
	return nil, nil, "", os.ErrNotExist
}
