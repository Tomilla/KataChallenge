package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBitCounting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BitCounting Suite")
}
