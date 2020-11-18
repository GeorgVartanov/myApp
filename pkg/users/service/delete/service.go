package delete

type RepositoryUserDelete interface {
	Delete(email string) error
}

type ServiceUserDelete interface {
	DeleteUser(email string) error
}

//service Create ...
type service struct {
	repo RepositoryUserDelete
}

//NewCreateUserService Init
func NewCreateUserService(r RepositoryUserDelete) ServiceUserDelete {
	return &service{repo: r}
}

func (s service) DeleteUser(email string) error {
	err := s.repo.Delete(email)
	if err != nil {
		return err
	}
	return nil
}
