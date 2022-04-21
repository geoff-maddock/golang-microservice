package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "geoff",
		Price: 1.00,
		SKU:   "abc-def-xyz",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
