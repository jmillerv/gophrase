# gophrase
CLI passphrase generator written in Go

### About Diceware Passwords

### Usage 
Call the bin file and pass a number that represents the number of words you would like your phrase to be. 
Ex: `gophrase 5`

### Corpus 

This passphrase generator currently relies on the Electronic Frontier Foundation's shortword list that uses only words that have unique three-character prefixes. If you're
interested I woudl recommend reading the [EFF's post](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) on their word lists. 

### Roadmap

1. Different corpus options
2. Special Characters
3. Common character substitution (a=>@, i=1) 
4. Random capitalization
5. Numbers

