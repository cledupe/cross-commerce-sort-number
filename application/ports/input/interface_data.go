package input

type InterfaceData interface {
	GetData(page int) ([]float32, error)
}
