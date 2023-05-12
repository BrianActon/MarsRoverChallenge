package rovers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchInstructions(t *testing.T) {
	fmt.Println("TestFetchInstructions")

	fileName := "test-squad-input"

	want := []string{"5 5", "1 1 N", "L"}
	got, err := fetchInstructions(fileName)

	assert.Nil(t, err)
	assert.EqualValues(t, got, want)
}
