package config

import (
	"github.com/gophrase/internal"
	"testing"
)

func TestPrintConfigDefaults(t *testing.T) {
	var defaults = Default{
		WordCount: 3,
		WordList:  internal.EFF_SHORT_2,
	}
	PrintConfigDefaults(defaults)
}
