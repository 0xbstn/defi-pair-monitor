package monitor

type Factory struct {
	Name           string
	FactoryAddress [20]byte
}

type Monitor struct {
	Name        string
	ChainID     int
	RPC         string
	FactoryList []Factory
}
