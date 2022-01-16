package input

type InterfaceData interface {
	GetData(page int, numbers chan<- []float64)
}
