package queue

type Consumer interface {
	Consume() error
}
