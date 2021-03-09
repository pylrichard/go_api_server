package util

import (
	"testing"
)

func TestGenShortID(t *testing.T)  {
	shortID, err := GenShortID()
	if shortID == "" || err != nil {
		t.Error("GenShortID failed")
	}

	t.Log("GenShortID test pass")
}

func BenchmarkGenShortID(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}

func BenchmarkGenShortIDTimeConsuming(b *testing.B)  {
	b.StopTimer()

	shortID, err := GenShortID()
	if shortID == "" || err != nil {
		b.Error(err)
	}

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}