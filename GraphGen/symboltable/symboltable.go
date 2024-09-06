package symboltable

type Type int32

const (
	PERSON Type = iota
	GROUP
	RELATIONSHIP
)

type SymbolTable struct {
	Table map[string]Type
}

func New() *SymbolTable {
	return &SymbolTable{
		Table: make(map[string]Type),
	}
}

func (s *SymbolTable) AddSymbol(name string, symbolType Type) {
	s.Table[name] = symbolType
}

func (s *SymbolTable) Exists(name string) bool {
	_, ok := s.Table[name]

	return ok
}

func (s *SymbolTable) GetType(name string) Type {
	return s.Table[name]
}
