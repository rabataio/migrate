package golang_test

import (
	"github.com/golang-migrate/migrate/v4/source/golang"
	st "github.com/golang-migrate/migrate/v4/source/testing"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/source/golang/testdata/migrations"
)

func Test(t *testing.T) {
	g := &golang.Golang{}
	d, err := g.Open("")
	if err != nil {
		// reuse the embed.FS set in example_test.go
		t.Fatal(err)
	}

	st.Test(t, d)
}
