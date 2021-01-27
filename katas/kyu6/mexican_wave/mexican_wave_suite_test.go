package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMexicanWave(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MexicanWave Suite")
}
