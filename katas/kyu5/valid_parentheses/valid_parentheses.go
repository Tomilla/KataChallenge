package kata

import (
	"fmt"
)

// A basic stack (LIFO) data structure
type Stack struct {
	stack []rune
}

func (s *Stack) push(c rune) {
	s.stack = append(s.stack, c)
}

func (s *Stack) pop() rune {
	if len(s.stack) == 0 {
		return '\x00'
	}

	n := len(s.stack) - 1
	defer func() {
		s.stack = s.stack[:n]
	}()
	return s.stack[n]
}

func (s Stack) top() rune {
	if len(s.stack) == 0 {
		return '\x00'
	}
	n := len(s.stack) - 1
	return s.stack[n]
}
func (s Stack) len() int {
	return len(s.stack)
}

func ValidParentheses(parens string) bool {
	stack := Stack{}
	fmt.Printf("Check: %v\n", parens)

	// check empty
	// target := strings.TrimSpace(parens)
	// if len(target) == 0 {
	//     return false
	// }

	for _, char := range parens {
		if char == '(' {
			// fmt.Println("\tfound: (")
			stack.push(char)
		} else if char == ')' {
			// fmt.Println("\tfound: )")
			elem := stack.pop()
			if elem != '(' {
				return false
			}
		}
	}
	if stack.len() > 0 {
		return false
	}
	return true
}

// Wrong Solution
// Your condition is wrong. Because you cannot rely on the counts of character '(' or ')'.
// Let me give you an example:
// ")()(" your solution will return true for this input string, however its not a valid parentheses.
// The right parentheses not only the counts of '(' or ')' should be equal, but the order also need to be matched.
/*
func ValidParentheses2(parens string) bool {
    substringA := ")"
    substringB := "("
    countA := strings.Count(parens, substringA)
    countB := strings.Count(parens, substringB)
    if countA/2 == 0 && countB/2 == 0 {
        return true
    } else {
        return false
    }

}
*/
