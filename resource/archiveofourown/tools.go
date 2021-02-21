package archiveofourown

const Hostname = "archiveofourown.org"

type ao3Tools struct{}

func NewTools() *ao3Tools {
	return &ao3Tools{}
}
