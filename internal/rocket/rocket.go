package rocket

//Rocket - should of countain the definition on our
//rocket
type Rocket struct {
	ID      string
	Name    string
	Type    string
	Flights int
}

// Service - our rocket service , responsible for
// updating the rocket inventory
type Service struct{}

func New() Service {
	return Service{}
}
