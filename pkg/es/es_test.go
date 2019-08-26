package es

import (
	"testing"
)

func TestCleanLog(t *testing.T) {
	err := CleanByDay(
		[]string{"http://n1626.ops.gaoshou.me:9200"},
		"python-logging-*",
		"python-logging-2006.01.02",
		7,
		true,
	)
	if err != nil {
		t.Error(err)
	}
}
