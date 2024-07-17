package app

import (
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	got, err := NewConfig("../../example.config.yaml")
	if err != nil {
		t.Fatal(err)
	}

	want := Config{
		Port:       5050,
		Postgres:   "postgres://postgres:dev@localhost:5151/postgres?sslmode=disable",
		SessionAge: time.Hour,
	}

	if *got != want {
		t.Fatalf("want: %v\ngot: %v", want, got)
	}

}
