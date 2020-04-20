# gophrase
CLI passphrase generator written in Go

### About Diceware Passwords

### Usage 
Call the bin file and pass a number that represents the number of words you would like your phrase to be. 
Ex: `gophrase 5`

### Corpus 

This passphrase defaults to the Electronic Frontier Foundation's short wordlist that uses only words that have unique three-character prefixes. If you're interested I would recommend reading the [EFF's post](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) on their word lists. 
There are 4 lists to choose from, if you opt for the Reinhold list it's possible you will have numbers in the generated password. 
If you don't want numbers or special characters, the other lists should work for you.

### Roadmap

1. Different corpus options
2. Special Characters
3. Random capitalization
4. Numbers

