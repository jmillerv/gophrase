package entropy

import (
	"fmt"
	"github.com/nbutton23/zxcvbn-go"
	"github.com/nbutton23/zxcvbn-go/scoring"
)

type Metadata struct {
	Entropy   float64 `json:"entropy,omitempty"`
	CrackTime string  `json:"crack_time,omitempty"`
	Score     int     `json:"score,omitempty"`
}

// PrintEntropy returns the password and metadata about the strength of a given passphrase to STD out
func PrintEntropy(passphrase string) {
	entropy := getEntropy(passphrase)
	fmt.Println("Password: " + entropy.Password)
	fmt.Println("Entropy: ", entropy.Entropy, "bits")
	fmt.Println("Crack Time: " + entropy.CrackTimeDisplay)
	fmt.Println("Score: ", entropy.Score, "/ 4")
}

// GetEntropy returns the Metadata about the password
func GetEntropy(passphrase string) *Metadata {
	entropy := getEntropy(passphrase)
	return &Metadata{
		Entropy:   entropy.Entropy,
		CrackTime: entropy.CrackTimeDisplay,
		Score:     entropy.Score,
	}
}

func getEntropy(passphrase string) scoring.MinEntropyMatch {
	return zxcvbn.PasswordStrength(passphrase, nil)
}
