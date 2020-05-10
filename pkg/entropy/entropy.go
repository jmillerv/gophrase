package entropy

import (
	"fmt"
	"github.com/nbutton23/zxcvbn-go"
)

// Initially I wrote my own entropy calculator but found this package that's based on a Dropbox
// it's more conservative in its estimation than my original and those I found online. So, it's included here.

func PrintEntropy(passphrase string ) {
	entropy := zxcvbn.PasswordStrength(passphrase,nil)
	fmt.Println("Password: " + entropy.Password)
	fmt.Println("Entropy: ", entropy.Entropy)
	fmt.Println("Crack Time " + entropy.CrackTimeDisplay)
	fmt.Println("Score: ", entropy.Score, "/ 4")
}