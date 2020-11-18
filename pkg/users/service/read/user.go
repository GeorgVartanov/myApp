package read

//User from read
type User struct {
	Email         string `json:"email" db:"email"`
	Password      string `json:"password" db:"password"`
	DisplayName   string `json:"displayName" db:"display_name"`
}

//User from read
type UserWithOutPassword struct {
	Email         string `json:"email" db:"email"`
	DisplayName   string `json:"displayName" db:"display_name"`
}


