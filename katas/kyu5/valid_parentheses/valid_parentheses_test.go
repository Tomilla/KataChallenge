package kata_test

import (
	"math/rand"
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "katas/kyu5/valid_parentheses"
)

var _ = Describe("Example Tests", func() {
	for _, testFunc := range []func(string) bool{ValidParentheses} {
		It("passes example tests", func() {
			Expect(testFunc("()")).To(Equal(true))
			Expect(testFunc("(")).To(Equal(false))
			Expect(testFunc(")")).To(Equal(false))
			Expect(testFunc(")(()))")).To(Equal(false))
			Expect(testFunc("(())((()())())")).To(Equal(true))
			Expect(testFunc("(())(()())())")).To(Equal(false))
			Expect(testFunc("((()()))((()())())")).To(Equal(true))
			Expect(testFunc("(()()())")).To(Equal(true))
			Expect(testFunc("(((()()()())))(")).To(Equal(false))
		})
		It("empty tests", func() {
			Expect(testFunc("")).To(Equal(true))
			Expect(testFunc(" ")).To(Equal(true))
			Expect(testFunc("( ")).To(Equal(false))
			Expect(testFunc(")")).To(Equal(false))
		})

	}
})

var _ = Describe("Base Tests", func() {
	It("should detect valid parentheses", func() {
		Expect(ValidParentheses("()")).To(Equal(true))
		Expect(ValidParentheses("()()")).To(Equal(true))
		Expect(ValidParentheses("(())")).To(Equal(true))
		Expect(ValidParentheses("((((()))))")).To(Equal(true))
		Expect(ValidParentheses("(()()()())(())")).To(Equal(true))
		Expect(ValidParentheses("(())((()((()))))")).To(Equal(true))
	})

	It("should detect invalid parentheses", func() {
		Expect(ValidParentheses(")")).To(Equal(false))
		Expect(ValidParentheses("")).To(Equal(true))
		Expect(ValidParentheses("())")).To(Equal(false))
		Expect(ValidParentheses("()()()())")).To(Equal(false))
		Expect(ValidParentheses("((((((((")).To(Equal(false))
		Expect(ValidParentheses("())(")).To(Equal(false))
		Expect(ValidParentheses(")()()()(")).To(Equal(false))
		Expect(ValidParentheses("(()()))(")).To(Equal(false))
		Expect(ValidParentheses(")()(")).To(Equal(false))
		Expect(ValidParentheses("())(()")).To(Equal(false))
	})
})

var _ = Describe("Random Tests", func() {
	It("should work for random cases", func() {
		rand.Seed(time.Now().UTC().UnixNano())

		for i := 0; i < 10; i++ {
			Expect(ValidParentheses(randomValid())).To(Equal(true))
			Expect(ValidParentheses(randomInvalid())).To(Equal(false))
		}
	})
})

func randomValid() string {
	n := randInt(1, 5)
	m := randInt(1, 5)

	return strings.Repeat("(", m) + strings.Repeat("()", n) + strings.Repeat(")", m)
}

func randomInvalid() string {
	valid := randomValid()

	if randInt(0, 1) == 1 {
		return valid + ")"
	}

	return valid + "("
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}
