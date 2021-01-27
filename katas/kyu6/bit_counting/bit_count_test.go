package kata_test

import (
	"log"
	"reflect"
	"runtime"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "katas/kyu6/bit_counting"
)

var _ = Describe("BitCount", func() {

})

var _ = Describe("CountBits()", func() {
	It("basic tests", func() {
		for _, testFunc := range []func(uint) int{CountBits, CountBits2} {
			log.Printf("Test Func: %s, %p\n", runtime.FuncForPC(
				reflect.ValueOf(testFunc).Pointer()).Name(), testFunc)
			Expect(testFunc(0)).To(Equal(0))
			Expect(testFunc(4)).To(Equal(1))
			Expect(testFunc(7)).To(Equal(3))
			Expect(testFunc(9)).To(Equal(2))
			Expect(testFunc(10)).To(Equal(2))
			Expect(testFunc(15)).To(Equal(4))
		}
	})
})
