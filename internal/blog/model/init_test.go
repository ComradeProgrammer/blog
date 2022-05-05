package model

import (
	"testing"
)

func TestMain(m *testing.M) {
	ConnectDatabase("database_test.sqlite")
	ClearDatabase()
	m.Run()
}
