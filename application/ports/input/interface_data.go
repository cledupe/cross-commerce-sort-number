package input

//Interface para entrada de dados no ETL
type InterfaceData interface {
	GetData(page int, numbers chan<- []float64)
}
