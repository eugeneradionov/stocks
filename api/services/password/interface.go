package password

type Service interface {
	EncryptPassword(password, salt string) (string, error)
	CompareHashAndPassword(password, salt, hash string) bool
	GenerateSalt() (string, error)
}
