package entropy

import (
	"testing"
)

func TestPrintEntropy(t *testing.T) {
	passphrase := "vaNquished68eMperor9squirrelGetawaY/rustproof284564+"
	passphrase2 := "lipstickunflavoredgleefuljeepfondue"
	passphrase3 := "escalatorought"
	PrintEntropy(passphrase)
	PrintEntropy(passphrase2)
	PrintEntropy(passphrase3)
}
