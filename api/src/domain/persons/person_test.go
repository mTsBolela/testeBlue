package persons_test

import (
	"contablue/src/domain/persons"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/google/uuid"
)

func TestPerson(t *testing.T) {
	const errorDefult = "Expected response is [%v], but responsed value was [%v]"

	t.Run("when all parameters are valid, return a valid person struct", func(t *testing.T) {
		person := persons.NewPerson()
		person.ID = uuid.New().String()
		person.Name = "John Doe"
		person.Customer = "true"
		person.Provider = "false"
		person.Status = "active"
		require.Nil(t, person.Prepare())
	})

	t.Run("when param name is empty, return a error", func(t *testing.T) {
		person := persons.Person{
			ID:   uuid.New().String(),
			Name: "",
		}
		require.NotNil(t, person.Prepare())
		require.Error(t, person.Prepare())
	})

	t.Run("when the customer parameter is invalid, returns an error", func(t *testing.T) {
		person := persons.Person{
			ID:       uuid.New().String(),
			Name:     "John doe",
			Customer: "",
		}
		require.NotNil(t, person.Prepare())
		require.Error(t, person.Prepare())
		require.Equal(t, "Customer is required", person.Prepare().Error())
	})

	t.Run("when the customer parameter is invalid, returns an error", func(t *testing.T) {
		person := persons.Person{
			ID:       uuid.New().String(),
			Name:     "John doe",
			Customer: "true",
			Provider: "f",
		}
		require.NotNil(t, person.Prepare())
		require.Error(t, person.Prepare())
		require.Equal(t, "Provider is required", person.Prepare().Error())
	})

	t.Run("when the status parameter is invalid, returns an error", func(t *testing.T) {
		person := persons.Person{
			ID:       uuid.New().String(),
			Name:     "John doe",
			Customer: "true",
			Provider: "false",
			Status:   "a",
		}
		require.NotNil(t, person.Prepare())
		require.Error(t, person.Prepare())
		require.Equal(t, "Invalid person status", person.Prepare().Error())
	})
}
