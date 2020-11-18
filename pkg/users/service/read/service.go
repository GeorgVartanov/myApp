package read

type RepositoryUserSelect interface {
	SelectAll() ([]UserWithOutPassword, error)
	SelectOne(email string) (*User, error)
	SelectOneWithOutPassword(email string) (*UserWithOutPassword, error)
}

type ServiceUserRead interface {
	GetAll() ([]UserWithOutPassword, error)
	GetOne(email string) (*User, error)
	GetOneWithOutPassword(email string) (*UserWithOutPassword, error)
}

//service Create ...
type service struct {
	repo RepositoryUserSelect
}

//NewCreateUserService Init
func NewCreateUserService(r RepositoryUserSelect) ServiceUserRead {
	return &service{repo: r}
}

func (s service) GetAll() ([]UserWithOutPassword, error) {
	userDB, err := s.repo.SelectAll()
	if err != nil {
		return nil, err
	}
	return userDB, nil
}

func (s service) GetOne(email string) (*User, error) {
	userDB, err := s.repo.SelectOne(email)
	if err != nil {
		return nil, err
	}
	return userDB, nil
}

func (s service) GetOneWithOutPassword(email string) (*UserWithOutPassword, error) {
	userDB, err := s.repo.SelectOneWithOutPassword(email)
	if err != nil {
		return nil, err
	}
	return userDB, nil
}
