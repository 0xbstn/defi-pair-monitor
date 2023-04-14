package monitor

type Factory struct {
	Name           string
	FactoryAddress [32]byte
}

type Monitor struct {
	Name        string
	ChainID     int
	RPC         string
	FactoryList []Factory
}
