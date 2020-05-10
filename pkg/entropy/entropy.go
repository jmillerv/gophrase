package entropy

import (
	"fmt"
	"github.com/nbutton23/zxcvbn-go"
)

//func checkEntropyRating(entropy float64) string {
//	return ""
//}

func PrintEntropy(passphrase string ) {
	entropy := zxcvbn.PasswordStrength(passphrase,nil)
	fmt.Print(entropy)
}