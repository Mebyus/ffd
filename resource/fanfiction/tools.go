package fanfiction

const Hostname = "www.fanfiction.net"

type ffTools struct{}

func NewTools() *ffTools {
	return &ffTools{}
}
