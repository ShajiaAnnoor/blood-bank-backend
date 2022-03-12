package cache

type ConnectionStatus interface {
	Has(id1, id2 string) (bool, error)
	Set(id1, id2 string, status bool) error
}
