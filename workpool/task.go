package workpool

type Task interface {
	UUID() string
	Run()
}
