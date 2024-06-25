package main

import (
	"strings"
)

func mergeAlternately(word1 string, word2 string) string {
    letters_word1 := strings.Split(word1,"")
    letters_word2 := strings.Split(word2,"")
    var list []string 
	i:=0 
	for (i < len(letters_word1) && i < len(letters_word2)) {
		list = append(list, letters_word1[i])
        list = append(list, letters_word2[i])
		i++;
	}
	if (len(letters_word1) > len(letters_word2)){
		for i:= len(letters_word2) ; i < len(letters_word1) ; i++ {
            list = append(list, letters_word1[i])
    }} else {
		for i:= len(letters_word1) ; i < len(letters_word2) ; i++ {
            list = append(list, letters_word2[i])
    }}
	return (strings.Join(list,""))
}

func main() {
	word1 := "abc"
    word2 := "pqr"
    result := mergeAlternately(word1, word2)
    print(result)
}
