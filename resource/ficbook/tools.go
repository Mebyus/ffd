package ficbook

const Hostname = "ficbook.net"

type fbTools struct{}

func NewTools() *fbTools {
	return &fbTools{}
}