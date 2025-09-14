package repositories

type Repository interface {
	Create(entity interface{}) error
	GetByID(id string) (interface{}, error)
	Update(entity interface{}) error
	Delete(id string) error
	List(limit, offset int) ([]interface{}, error)
}

type BaseRepository struct {
}

func NewBaseRepository() *BaseRepository {
	return &BaseRepository{}
}
