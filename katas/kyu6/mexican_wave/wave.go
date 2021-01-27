package kata

/*
Introduction
The wave (known as the Mexican wave in the English-speaking world outside North America) is an example of metachronal rhythm achieved in a packed stadium when successive groups of spectators briefly stand, yell, and raise their arms. Immediately upon stretching to full height, the spectator returns to the usual seated position.

The result is a wave of standing spectators that travels through the crowd, even though individual spectators never move away from their seats. In many large arenas the crowd is seated in a contiguous circuit all the way around the sport field, and so the wave is able to travel continuously around the arena; in discontiguous seating arrangements, the wave can instead reflect back and forth through the crowd. When the gap in seating is narrow, the wave can sometimes pass through it. Usually only one wave crest will be present at any given time in an arena, although simultaneous, counter-rotating waves have been produced. (Source Wikipedia)
Task
In this simple Kata your task is to create a function that turns a string into a Mexican Wave. You will be passed a string and you must return that string in an array where an uppercase letter is a person standing up.
Rules
1.  The input string will always be lower case but maybe empty.

2.  If the character in the string is whitespace then pass over it as if it was an empty seat
Example
wave("hello") => []string{"Hello", "hEllo", "heLlo", "helLo", "hellO"}
Good luck and enjoy!
*/
import (
	"strings"
	"unicode"
)

func Wave(word string) []string {
	var result []string = []string{}

	var lenWord = len(word)
	if lenWord == 0 {
		return result
	}
	var gapLowerToUpper rune = 32

	for i := 0; i < lenWord; i++ {
		if word[i] == 32 {
			continue
		}
		var charList []rune
		for j, item := range word {
			if j == i {
				item -= gapLowerToUpper
			}
			charList = append(charList, item)
		}
		result = append(result, string(charList))
	}
	return result
}

func Wave2(words string) (wave []string) {
	wave = []string{} // leaving the array uninitialised would be better practice
	for i, c := range words {
		if unicode.IsSpace(c) {
			continue
		}
		upperC := string(c - 'a' + 'A')
		wave = append(wave, words[:i]+upperC+words[i+1:])
	}
	return
}

func Wave3(words string) []string {
	var result []string
	result = make([]string, 0, len(words))
	for i := 0; i < len(words); i++ {
		if !(unicode.IsSpace(rune(words[i]))) {
			// one int index ==> single ==> rune
			// curStr := strings.ToUpper(string(words[i]))
			// two int index ==> range ==> string
			curStr := strings.ToUpper(words[i : i+1])
			result = append(result, words[:i]+curStr+words[i+1:])
		}
	}

	return result
}

func Wave4(s string) []string {
	resArr := []string{}
	for i, c := range s {
		// POINT: from string to rune
		tmpArr := []rune(s)
		if c != 32 {
			tmpArr[i] = c - 32
			resArr = append(resArr, string(tmpArr))
		}
	}
	return resArr

}
