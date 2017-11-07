package lengthOfLastWord

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestLengthOfLastWord(t *testing.T) {
	g := Goblin(t)
	g.Describe("Length of last word", func() {
		g.It("Should be 5", func() {
			g.Assert(lengthOfLastWord("hello world")).Equal(5)
		})
		g.It("Should be 1", func() {
			g.Assert(lengthOfLastWord("### $$@ 2")).Equal(1)
		})
		g.It("Should be 0", func() {
			g.Assert(lengthOfLastWord("")).Equal(0)
		})
	})
}
