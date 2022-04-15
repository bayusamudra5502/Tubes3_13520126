package test

import (
	"testing"

	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/lib"
	"github.com/stretchr/testify/assert"
)

func TestBoyerMoore(t *testing.T) {
	t.Run("Pattern in text", func(t *testing.T) {
		assert.Equal(t, 0, lib.BoyerMoore("AAAAAAAAA", "AAAA"))
		assert.Equal(t, 0, lib.BoyerMoore("ACTGA", "ACTGA"))

		assert.Equal(t, 3, lib.BoyerMoore("ACTACGCATG", "ACG"))
		assert.Equal(t, 10, lib.BoyerMoore("ACTGAACTGAACTGT", "ACTGT"))

		assert.Equal(t, 12, lib.BoyerMoore("ACTACTACTACTACTGTAAA", "ACTGT"))
		assert.Equal(t, 7, lib.BoyerMoore("AACTGATAACTGAA", "AACTGAA"))
	})

	t.Run("Pattern not in text", func(t *testing.T) {
		assert.Equal(t, -1, lib.BoyerMoore("ACTGAACGATGA", "ACTGG"))
		assert.Equal(t, -1, lib.BoyerMoore("ACTGA", "ATTT"))
		assert.Equal(t, -1, lib.BoyerMoore("TTTTTTT", "CCCCCC"))
		assert.Equal(t, -1, lib.BoyerMoore("ACT", "ACGCA"))
		assert.Equal(t, -1, lib.BoyerMoore("Aku Anak Sehat", "xyz"))
	})

}
