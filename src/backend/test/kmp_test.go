package test

import (
	"testing"

	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/lib"
	"github.com/stretchr/testify/assert"
)

func TestKmp(t *testing.T) {
	t.Run("Pattern in text", func(t *testing.T) {
		assert.Equal(t, 0, lib.Kmp("AAAAAAAAA", "AAAA"))
		assert.Equal(t, 0, lib.Kmp("ACTGA", "ACTGA"))

		assert.Equal(t, 3, lib.Kmp("ACTACGCATG", "ACG"))
		assert.Equal(t, 10, lib.Kmp("ACTGAACTGAACTGT", "ACTGT"))

		assert.Equal(t, 12, lib.Kmp("ACTACTACTACTACTGTAAA", "ACTGT"))
		assert.Equal(t, 7, lib.Kmp("AACTGATAACTGAA", "AACTGAA"))
	})

	t.Run("Pattern not in text", func(t *testing.T) {
		assert.Equal(t, -1, lib.Kmp("ACTGAACGATGA", "ACTGG"))
		assert.Equal(t, -1, lib.Kmp("ACTGA", "ATTT"))
		assert.Equal(t, -1, lib.Kmp("TTTTTTT", "CCCCCC"))
		assert.Equal(t, -1, lib.Kmp("ACT", "ACGCA"))
		assert.Equal(t, -1, lib.Kmp("Aku Anak Sehat", "xyz"))
	})
}
