package rocket

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)

	defer mockCtrl.Finish()
	t.Run("test get rocket by id", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			GetRocketByID(id).
			Return(Rocket{
				ID: id,
			}, nil)
		rocketService := New(rocketStoreMock)
		rkt, err := rocketService.GetRocketByID(
			context.Background(),
			id,
		)
		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rkt.ID)
	})
}
