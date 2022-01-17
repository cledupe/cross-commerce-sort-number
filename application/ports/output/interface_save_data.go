package output

//Interface para saida de dados no ETL
type InterfaceData interface {
	Save(data []float64)
	Load() ([]float64, error)
}
