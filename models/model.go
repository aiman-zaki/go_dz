package models

type Model interface {
	Create() error
	Read() error
	Update() error
	Delete() error
}
