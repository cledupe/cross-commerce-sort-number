package domain

type StNumber struct {
	value float32
}

func (n *StNumber) Set(number float32) {
	n.value = number
}

func (n StNumber) Get() float32 {
	return n.value
}
