package model

import (
	"testing"
)

func TestMain(m *testing.M) {
	InitDatabase("database_test.sqlite")
	ClearDatabase()
	m.Run()
}
