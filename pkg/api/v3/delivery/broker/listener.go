package broker

type Listener interface {
	Listen() error
}
