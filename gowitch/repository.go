package gowitch

type RepositoryInterface interface {
	GetAll() []interface{}
	GetById() interface{}
	Delete() interface{}
	Update() interface{}
	Create() interface{}
}

type Repository struct {
	Model interface{}
}
