package rovers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMove(t *testing.T) {
	fmt.Println("TestMove")
	instructions := []string{"10 14", "5 8 N", "RMMMLMM"}
	want := []RoverPosition{{X: 8, y: 10, direction:'N'}}
	got, err := move(instructions)

	fmt.Println("want: ", want, " got: ", got, "got[0].direction:", got[0].direction)

	assert.Nil(t, err)
	assert.EqualValues(t, got, want)
}


func TestMoveRover(t *testing.T) {
	fmt.Println("TestMoveRover")
	test := []struct {
		testName string
		rp RoverPosition
		xMax int
		yMax int
		instructions string
		want RoverPosition
		wantErr string
	}{ 
		{
			testName: "pass",
			rp: RoverPosition{5, 8, 'N'},
			xMax: 20, 
			yMax: 40, 
			instructions: "MMMMRMMMMLMMM",
			want: RoverPosition{9, 15, 'N'},
			wantErr: "",
		},
		{
			testName: "invalid character",
			rp: RoverPosition{5, 8, 'N'},
			xMax: 20, 
			yMax: 40, 
			instructions: "MMMPMRMMMMLMMM",
			want: RoverPosition{},
			wantErr: "invalid character",
		},
		{
			testName: "North bound error",
			rp: RoverPosition{5, 8, 'N'},
			xMax: 20, 
			yMax: 9, 
			instructions: "MMMMMM",
			want: RoverPosition{},
			wantErr: "rover out of Northern bounds",
		},
		{
			testName: "South bound error",
			rp: RoverPosition{5, 8, 'N'},
			xMax: 20, 
			yMax: 9, 
			instructions: "LLMMMMMMMMM",
			want: RoverPosition{},
			wantErr: "rover out of Southern bounds",
		},
		{
			testName: "East bound error",
			rp: RoverPosition{5, 8, 'N'},
			xMax: 6, 
			yMax: 9, 
			instructions: "RMMMMMM",
			want: RoverPosition{},
			wantErr: "rover out of Eastern bounds",
		},
		{
			testName: "West bound error",
			rp: RoverPosition{5, 8, 'N'},
			xMax: 6, 
			yMax: 9, 
			instructions: "LMMMMMM",
			want: RoverPosition{},
			wantErr: "rover out of Western bounds",
		},
	}
	
	for _, v := range test {
		got, err := moveRover(v.rp, v.xMax, v.yMax, v.instructions)
	
		switch v.testName {
		case "pass":
			assert.Nil(t, err)
		default:
			if err.Error() != v.wantErr {
				t.Errorf("Error got = %v, and Expected = %v.", err, v.wantErr)
			}
		}
		assert.EqualValues(t, got, v.want)
	}
}

func TestGetRoverPosition(t *testing.T) {
	fmt.Println("TestGetRoverPosition")
	test := []struct {
		testName string
		instruction string
		want RoverPosition
		wantErr string
	}{ 
		{
			testName: "pass",
			instruction: "3 5 N",
			want: RoverPosition{3, 5, 'N'},
			wantErr: "",
		},
		{
			testName: "fail",
			instruction: "35 N",
			want: RoverPosition{},
			wantErr: "",
		},
	}
	
	for _, v := range test {
		got, err := getRoverPosition(v.instruction)
	
		if v.testName == "pass"{
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
		
		assert.EqualValues(t, got, v.want)
	}
}

func TestGetGridMax(t *testing.T) {
	fmt.Println("TestGetGridMax")
	test := []struct {
		testName string
		max string
		wantX int
		wantY int
		wantErr string
	}{ 
		{
			testName: "pass",
			max: "30 30",
			wantX: 30,
			wantY: 30,
			wantErr: "",
		},
		{
			testName: "fail",
			max: "35 N",
			wantX: 0,
			wantY: 0,
			wantErr: "",
		},
	}
	
	for _, v := range test {
		gotX, gotY, err := getGridMax(v.max)
	
		if v.testName == "pass"{
			assert.Nil(t, err)
		} else {
			assert.NotNil(t, err)
		}
		
		assert.EqualValues(t, gotX, v.wantX)
		assert.EqualValues(t, gotY, v.wantY)
	}
}
