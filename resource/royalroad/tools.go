package royalroad

type rrTools struct {
	ficName    string
	chapters   int
	saveSource bool
}

func NewTools() *rrTools {
	return &rrTools{}
}
