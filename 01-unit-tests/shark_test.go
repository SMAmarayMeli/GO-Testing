package hunt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	// arrange
	s := Shark{
		hungry: true,
		tired:  false,
		speed:  5,
	}
	p := Prey{
		name:  "Perry",
		speed: 3,
	}
	// act
	err := s.Hunt(&p)
	// assert
	assert.NoError(t, err)

}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	// arrange
	s := Shark{
		hungry: true,
		tired:  true,
		speed:  5,
	}
	p := Prey{
		name:  "Perry",
		speed: 3,
	}
	expected := fmt.Errorf("cannot hunt, i am really tired")
	// act
	err := s.Hunt(&p)
	// assert
	assert.Equal(t, expected, err)
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	// arrange
	s := Shark{
		hungry: false,
		tired:  false,
		speed:  5,
	}
	p := Prey{
		name:  "Perry",
		speed: 3,
	}
	expected := fmt.Errorf("cannot hunt, i am not hungry")
	// act
	err := s.Hunt(&p)
	// assert
	assert.Equal(t, expected, err)
}

func TestSharkCannotReachThePrey(t *testing.T) {
	// arrange
	s := Shark{
		hungry: true,
		tired:  false,
		speed:  5,
	}
	p := Prey{
		name:  "Perry",
		speed: 10,
	}
	expected := fmt.Errorf("could not catch it")
	// act
	err := s.Hunt(&p)
	// assert
	assert.Equal(t, expected, err)
}

func TestSharkHuntNilPrey(t *testing.T) {
	// arrange
	s := Shark{
		hungry: true,
		tired:  false,
		speed:  5,
	}
	expected := fmt.Errorf("pls not nil")
	// act
	err := s.Hunt(nil)
	// assert
	t.Log(err)
	assert.Equal(t, expected, err)
}
