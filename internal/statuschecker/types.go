package statuschecker

const limiter int = 5

type checker struct{}

type Result struct {
	Url string
	Up  bool
}

type StatusChecker interface {
	Check(concurrency int, src string, urls ...string) (<-chan *Result, <-chan struct{}, error)
}
