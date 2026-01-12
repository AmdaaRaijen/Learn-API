package products

type Service interface {
	GetListOfProducts() ([]string, error)
}

type service struct {}

func NewService() *service {
	return  &service{}
}

func (s *service) GetListOfProducts() ([]string, error) {
	data := []string{"product1", "product2", "product3"}
	
	return  data, nil
}