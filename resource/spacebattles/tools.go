package spacebattles

type sbTools struct {
	ficName    string
	chapters   int
	saveSource bool
}

func NewTools() *sbTools {
	return &sbTools{}
}
