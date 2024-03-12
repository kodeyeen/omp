package event

type listener struct {
	handler any
	once    bool
}
