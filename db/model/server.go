package model

type ServerStatus int

const (
	On ServerStatus = iota
	Off
)

var status = []string{"on", "off"}

func (c *ServerStatus) String() string {
	return status[*c]
}

type Server struct {
	BaseModel
	name   string
	status ServerStatus
	ipv4   string
}
