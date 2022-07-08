package datagenerator_test

type GenericDataGenerator interface {
	Create(low, high int) GenericDataGenerator
	GetRequest() int
}
