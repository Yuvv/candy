package lang

type AutoCloseable interface {
	// Close function closes this resource
	Close()
}
