package product

type Service interface {
	GetProductById(param *getProductByIDrequest) (*Product, error)
	GetProducts(params *getProductsRequest) (*ProductList, error)
	InsertProduct(params *getAddProductRequest) (int64, error)
}

type service struct{ repo Repository }

func NewService(repo Repository) Service {
	return &service{repo: repo}

}

func (s *service) GetProductById(param *getProductByIDrequest) (*Product, error) {
	return s.repo.GetProductById(param.ProductID)

}

func (s *service) GetProducts(params *getProductsRequest) (*ProductList, error) {
	products, err := s.repo.GetProducts(params)
	if err != nil {
		panic(err)
	}

	totalproduct, err := s.repo.GetTotalProducts()
	if err != nil {
		panic(err)
	}
	return &ProductList{Data: products, TotalRecords: totalproduct}, nil
}

func (s *service) InsertProduct(params *getAddProductRequest) (int64, error) {
	return s.repo.InsertProduct(params)
}
