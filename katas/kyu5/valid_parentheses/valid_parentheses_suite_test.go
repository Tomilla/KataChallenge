package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestValidParentheses(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ValidParentheses Suite")
}
