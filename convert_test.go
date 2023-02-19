package go_time_converter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	actual := Convert(15)
	expected := 15
	assert.Equal(t, expected, actual, fmt.Sprintf("Expected to be %v", expected))
}
