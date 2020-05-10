# gophrase
CLI passphrase generator written in Go

### About Diceware Passwords
[Diceware](https://en.wikipedia.org/wiki/Diceware) is a method that allows anyone to use ordinary dice to generate an easy to remember passphrase
from a list of words using only dice. The original creator was Arnold Reinhold. Passwords generated in this
way possess about 12.9 bits of entropy per word if the attacker knows Diceware was the method used, the list, 
and the word length. Less information can yield higher passphrase entropy. 

### About Entropy
Entropy is essentially a measure of how unpredictable a password is. 

The below chart is a good metric of password strength based on entropy.
>< 28 bits = Very Weak; might keep out family members  
 28 - 35 bits = Weak; should keep out most people, often good for desktop login passwords  
 36 - 59 bits = Reasonable; fairly secure passwords for network and company   
 60 - 127 bits = Strong; can be good for guarding financial information  
 128+ bits = Very Strong; often overkill ([Pleacher](https://www.pleacher.com/mp/mlessons/algebra/entropy.html))

### Gophrase Usage 
Call the bin file and pass a number that represents the number of words you would like your phrase to be. 
Gophrase defaults to a three word passphrase with no capitalization, special characters or numbers.

Ex: `gophrase gen` will yield a default three word passphrase  
Ex: `gophrase gen 5` will yield a five word passphrase

The [current recommendation](http://world.std.com/~reinhold/dicewarefaq.html#howlong) for secure passwords is a passphrase of at least six words.

If you want to switch the word list simply pass the letter a, b, c, or d after the number of words.
Ex: `gophrase gen 3 d`  
Note: If changing the default passphrase length, wordlist must be the second argument. 

**Options**

a = EFF short word list 2  
b = EFF short word list 1  
c = EFF large word list  
d = Reinhold word list  

### Flags
`-c` Adds capital letters to the passphrase  
`-n` Adds numbers to the passphrase   
`-s` Adds special characters to the passphrase  

Ex: `gophrase gen -n -s -c` will generate a three word passphrase with numbers, special characters, and capitals. 

The order of the flags does not matter. 

### Corpus 

This passphrase defaults to the Electronic Frontier Foundation's short word list that uses only words that have unique three-character prefixes. If you're interested I would recommend reading the [EFF's post](https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) on their word lists. 
There are 4 lists to choose from, if you opt for the Reinhold list it's possible you will have numbers in the generated password. 
If you don't want numbers or special characters, the other lists should work for you.

### Further Reading
https://theworld.com/~reinhold/diceware.html
https://www.pleacher.com/mp/mlessons/algebra/entropy.html

### Usage in Production Environments
This is the use at your own risk warning. As this program uses the math/rand package and not the crypto/rand package. As such I would avoid production 
environments because of the possibility of being subject to a time attack. That being said security is all about 
threat models, and most users who just need a quick simple password aren't being actively targeted on their machines. 

Replacing the math/rand with crypto/rand isn't out of the picture, it's just not a priority at the moment.
 
### Roadmap

1. CLI validation
2. Defaults for wordlist and word count
