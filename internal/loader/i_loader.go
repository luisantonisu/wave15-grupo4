package loader

type Loader interface {
	Load() (*DB, error)
}
