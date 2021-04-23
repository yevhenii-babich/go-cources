package some

type main struct {
}

func (m main) Get() {
	println("some.one.go")
}

func NewMain() Main {
	return &main{}
}

type Main interface {
	Get()
}
