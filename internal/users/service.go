package users

type Service interface {
}

type service struct {
	r Repository
}

func NewService(r Repository) service {
	return service{r: r}
}
