package product

import "ecommerce/domain"


type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service{
	return &service{
		productRepo: productRepo,
	}
} 

func (svc *service) Create(p domain.Product) (*domain.Product, error){
	prd, err := svc.productRepo.Create(p)
	if err != nil{
		return nil, err
	}
	if prd == nil{
		return nil, nil
	}
	return prd, nil
}
func (svc *service) Get(productID int) (*domain.Product, error){
	prd, err := svc.productRepo.Get(productID)
	if err != nil{
		return nil, err
	}
	if prd == nil{
		return nil, nil
	}
	return prd, nil
}
func (svc *service) List() ([]*domain.Product, error){
	prd, err := svc.productRepo.List()
	if err != nil{
		return nil, err
	}
	if prd == nil{
		return nil, nil
	}
	return prd, nil
}
func (svc *service) Update(p domain.Product) (*domain.Product, error){
	prd, err := svc.productRepo.Update(p)
	if err != nil{
		return nil, err
	}
	if prd == nil{
		return nil, nil
	}
	return prd, nil
}
func (svc *service) Delete(ProductID int) error{
	err := svc.productRepo.Delete(ProductID)
	if err != nil{
		return err
	}
	return nil
}
	
