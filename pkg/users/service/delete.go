package service

func (s service) DeleteByEmail(email string) error {
	err := s.repo.DeleteByEmail(email)
	if err != nil {
		return err
	}
	return nil
}
