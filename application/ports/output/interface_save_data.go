package output

type InterfaceData interface {
	Save(data []float64)
	Load() ([]float64, error)
}
