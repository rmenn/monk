package node

type NodeInterface interface {
	Start(url string)
	Stop()
	Status() (string, error)
}

type NodeType int

const (
	MONK_MASTER NodeType = iota
	MONK_PUPIL
)

type Node struct {
	URL  string
	Type NodeType
}
