package rovers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildInstructions(t *testing.T) {
	fmt.Println("TestBuildInstructions")

    rsi := RoverSquadInput{
    	MaxGrid: "30 30",
    	Rovers: []Rover{
    		{
				RoverStart: "4 6 N",
				Movement: "M",
    		},
    	},
    }


	want := []string{"30 30", "4 6 N", "M"}
	got := buildInstructions(rsi)

	assert.EqualValues(t, got, want)
}