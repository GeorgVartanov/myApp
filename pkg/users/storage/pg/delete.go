package pg

//	Delete(email string) error
func (u *UserPostgresStorage) Delete(email string) error {
	//var user read.User

	//tx := u.MustBegin()
	//
	//fmt.Printf("SELECT ONE USER %v \n", user)
	//
	//// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	//err := tx.QueryRowx(`SELECT email, password, display_name FROM app_user WHERE email=$1`, email).StructScan(&user)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = tx.Commit()

	_, err := u.Exec("DELETE FROM app_user WHERE email=$1", email)
	if err != nil {
		return err
	}

	return nil
}
