package test

import (
	"testing"

	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/lib"
	"github.com/stretchr/testify/assert"
)

func TestSimilarity(t *testing.T) {
	t.Run("Exactly same", func(t *testing.T) {
		assert.Equal(t, 1.0, lib.Similarity("ACTGAACGATGA", "ACTGA"))
		assert.Equal(t, 1.0, lib.Similarity("ACTGA", "ACTGA"))
		
		assert.Equal(t, 1.0, lib.Similarity("ATTAACTGAATTA", "ACTGA"))
		assert.Equal(t, 1.0, lib.Similarity("ACTGACTGTCTGA", "ACTGT"))
		
		assert.Equal(t, 1.0, lib.Similarity("ACTGAATACTGT", "ACTGT"))
		assert.Equal(t, 1.0, lib.Similarity("AACTGATAACTGAA", "AACTGAA"))
	})

	t.Run("Patially Same", func(t *testing.T) {
		assert.Equal(t, 0.8, lib.Similarity("ACTGAACGATGA", "ACTGG"))
		assert.Equal(t, 0.5, lib.Similarity("ACTGA", "ATTT"))
		assert.Equal(t, 0.0, lib.Similarity("TTTTTTT", "CCCCCC"))
		assert.Equal(t, 0.4, lib.Similarity("ACT", "ACGCA"))
		assert.Equal(t, 0.0, lib.Similarity("Aku Anak Sehat", "xyz"))
	})

}