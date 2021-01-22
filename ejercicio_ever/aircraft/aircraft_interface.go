package aircraft

// Aircraft interface
type Aircraft interface {
	GetType() string
	Flight() error
	Attack() (string, error)
}
