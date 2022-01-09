package framework

type UserIterator struct {
	Index int
	Users []*User
}

func (u *UserIterator) HasNext() bool {
	if u.Index < len(u.Users) {
		return true
	}
	return false

}
func (u *UserIterator) GetNext() *User {
	if u.HasNext() {
		user := u.Users[u.Index]
		u.Index++
		return user
	}
	return nil
}