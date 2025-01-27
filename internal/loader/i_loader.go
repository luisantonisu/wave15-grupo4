package loader

type Loader interface {
	Load() map[int]any
}