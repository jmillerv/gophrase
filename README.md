# gophrase
CLI passphrase generator written in Go

### About Diceware Passwords
[Diceware](https://en.wikipedia.org/wiki/Diceware) is a method that allows anyone to use ordinary dice to generate an easy to remember passphrase
from a list of words using only dice. The original creator was Arnold Reinhold. Passwords generated in this
way possess about 12.9 bits of entropy per word if the attacker knows Diceware was the method used, the list, 
and the word length. Less information can yield higher passphrase entropy. 

### Usage 
Call the bin file and pass a number that represents the number of words you would like your phrase to be. 
Gophrase defaults to a three word passphrase with no capitalization, special characters or numbers.

Ex: `gophrase gen` will yield a default three word passphrase  
Ex: `gophrase gen 5` will yield a five word passphrase

The [current recommendation](http://world.std.com/~reinhold/dicewarefaq.html#howlong) for secure passwords is a passphrase of at least six words.

If you want to switch the word list simply pass the letter a, b, c, or d after the number of words.
Ex: `gophrase gen 3 d`  
a = EFF short word list 2  
b = EFF short word list 1  
c = EFF large word list  
d = Reinhold word list  

### Flags
`-c` Adds capital letters to the passphrase  
`-n` Adds numbers to the passphrase   
`-s` Adds special characters to the passphrase  

Ex: `gophrase gen -n -s -c`

### Corpus 

This passphrase defaults to the Electronic Frontier Foundation's short word list that uses only words that have unique three-character prefixes. If you're interested I would recommend reading the [EFF's post](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) on their word lists. 
There are 4 lists to choose from, if you opt for the Reinhold list it's possible you will have numbers in the generated password. 
If you don't want numbers or special characters, the other lists should work for you.

### Further Reading
https://theworld.com/~reinhold/diceware.html

### Roadmap

1. CLI validation
2. Print list options
3. Entropy output