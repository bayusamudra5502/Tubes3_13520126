package test

import (
	"testing"

	"github.com/bayusamudra5502/Tubes3_13520126/src/backend/lib"
	"github.com/stretchr/testify/assert"
)

func TestSimilarity(t *testing.T){
	assert.Equal(t, lib.Similarity("Aku Anak Sehat", "xyz"), 0.0)
}