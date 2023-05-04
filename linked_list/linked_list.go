/*
input: string
output: corrected word

if input is a word in the trie, return the word
return the word whose prefix matches the input

wordFoundSoFar = ""
currentNode = trie root

iterate input, for each char
  if currentNode[char]
	  wordFoundSoFar += char

		currentNode = currentNode[char]
		next

	closestChar = keys(currentNode)[0]
	wordFoundSofar += closestChar

	break

return wordFoundSoFar
--
*/
