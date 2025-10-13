package services

type Service interface {
	Create(data interface{}) (interface{}, error)
	GetByID(id string) (interface{}, error)
	Update(id string, data interface{}) (interface{}, error)
	Delete(id string) error
	List(limit, offset int) ([]interface{}, error)
}

type BaseService struct {
}

func NewBaseService() *BaseService {
	return &BaseService{}
}
