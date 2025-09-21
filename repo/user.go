package repo

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface{
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
	// Get(userID int) (*User, error)
	// List() ([]*User, error)
	// Update(user User) (*User, error)
	// Delete(userID int) error
}

type userRepo struct{
	users []User
}

func NewUserRepo() UserRepo{
	return &userRepo{}
}
func (u userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return  &user, nil
	}
	user.ID = len(u.users) + 1
	u.users = append(u.users, user)
	return &user, nil
}


func (u *userRepo) Find(email, pass string)  (*User, error){
 for _, user := range u.users{
	 if user.Email == email && user.Password == pass{
		 return &user, nil
	 }
 }
 return nil, nil
}