package database

type UserDBRepository struct {
}

func (udbr UserDBRepository) GetByEmail(email string) {
	// SELECT email, password_hash, created_at FROM users WHERE email = email
	panic("implement me")
}