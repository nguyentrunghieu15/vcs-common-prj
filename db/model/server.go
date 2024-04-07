package model

type ServerStatus string

const (
	On  ServerStatus = "on"
	Off ServerStatus = "off"
)

type Server struct {
	BaseModel
	Name   string
	Status ServerStatus
	Ipv4   string
}
