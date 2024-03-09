package event

type handler struct {
	handle any
	once   bool
}
