package stream_manager_test

import (
	"fmt"
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/stream_manager"
	"github.com/stretchr/testify/assert"
)

func Test_StreamManager(t *testing.T) {
	manager := stream_manager.New()

	type type1 struct {
		Field1 string
	}

	fn := func(actual any, err error) error {
		fmt.Printf("received %v", actual)
		assert.Equal(t, `555-555-5555`, actual)
		return nil
	}

	err := manager.Register(type1{}, fn)
	assert.Nil(t, err)

	input := `{
        "name": "github.com/boundedinfinity/schema/idiomatic/stream_manager_test/type1",
        "instance": {
                "Field1": "some stuff"
            }
        }`

	err = manager.Handle([]byte(input))
	assert.Nil(t, err)
}
