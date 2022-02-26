// Package symbol keeps a correspondence between symbolic labels and numeric addresses.
package symbol

type Symbol struct{}

// New instantiates a new symbol table.
func New() *Symbol {
	return &Symbol{}
}

// AddEntry adds the pair symbol, address to the table.
func (s *Symbol) AddEntry(symbol string, address int) {
	return
}

// Contains asserts if the table contains the given symbol.
func (s *Symbol) Contains(symbol string) bool {
	return false
}

// GetAddress returns the address associated with the symbol.
func (s *Symbol) GetAddress(symbol string) int {
	return 0
}
