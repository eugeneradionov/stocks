package password

type service struct{}

func New() Service {
	return &service{}
}
