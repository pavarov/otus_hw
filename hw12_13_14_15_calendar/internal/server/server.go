package server

type Interface interface {
	Run() error
	Stop() error
}
