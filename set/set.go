package set

type Interface interface {
	Add(v interface{}) bool
	Remove(v interface{}) bool
	IsElementOf(v interface{}) bool
	Size() int
	Values() []interface{}
}
type Emptier interface {
	Empty()
}
