package handlers

// Handler is implemented by any handler.
// The Handle method is used to process event
type Handler interface {
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(oldObj, newObj interface{})
}

// Map maps each event handler function to a name for easily lookup
var Map = map[string]interface{}{
	"default": &Default{},
}

// Default handler implements Handler interface,
// print each event with JSON format
type Default struct {
}

func newDefaultHandler() *Default {
	return &Default{}
}

func (d *Default) ObjectCreated(obj interface{}) {

}

func (d *Default) ObjectDeleted(obj interface{}) {

}

func (d *Default) ObjectUpdated(oldObj, newObj interface{}) {

}
