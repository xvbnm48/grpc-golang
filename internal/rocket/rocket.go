//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/xvbnm48/grpc-golang/internal/rocket Store

package rocket

import "context"

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
type Service struct {
	Store Store
}

//New return a new instance of our rocket service
func New(store Store) Service {
	return Service{
		Store: store,
	}
}

//get rocket by id- retrives a rocket based on the id the store
func (s Service) GetRocketByID(ctx context.Context, id string) (Rocket, error) {
	rkt, err := s.Store.GetRocketByID(id)
	if err != nil {
		return Rocket{}, err
	}
	return rkt, nil

}

//inser rocket - insert a new rocket to store
func (s Service) InsertRocket(ctx context.Context, rkt Rocket) (Rocket, error) {
	rkt, err := s.Store.InsertRocket(rkt)

	if err != nil {
		return Rocket{}, err
	}

	return rkt, nil
}

//delete rocket - deletes a rocket from our inventory
func (s Service) DeleteRocket(id string) error {
	err := s.Store.DeleteRocket(id)
	if err != nil {
		return err
	}

	return nil
}
