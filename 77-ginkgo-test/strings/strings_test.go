package strings_test

import (
	"testing"

	mystrings "demo/strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Strings package suite")
}

var _ = Describe("Strings", func() {
	Context("String Functions", func() {
		It("has to reverse the string", func() {
			input := "Victoria Secrets & Co"
			// Ω(mystrings.Reverse(&input)).Should(Equal("oC & sterceS airotciV"))
			Expect(mystrings.Reverse(&input)).Should(Equal("oC & sterceS airotciV"))
		})
		It("has to reverse nothing when nil is passed", func() {
			Expect(mystrings.Reverse(nil)).ShouldNot(Equal(nil))
		})
	})
})

// First download ginkgo
// Create your own package and use . notation to use Functions or identifiers in that package
// Describe
// Context
// It
// Specify
