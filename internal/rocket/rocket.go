package rocket

//Rocket - should of countain the definition on our
//rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// store - defines the interface we expect
//our database the rocket inventory
type Store interface {
	GetRocketByID(id string) (Rocket, error)
	InsertRocket(rkt Rocket) (Rocket, error)
	DeleteRocket(id string) error
}

// Service - our rocket service , responsible for
// updating the rocket inventory
type Service struct{}

//New return a new instance of our rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}
