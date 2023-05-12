package rovers

import (
	"errors"
	"fmt"
	"strconv"
)


type RoverPosition struct {
	X int 
	y int 
	direction byte
}

func move(instructions []string) ([]RoverPosition, error) {
	xMax, yMax, err := getGridMax(instructions[0])
	if err != nil {
		return []RoverPosition{}, err
	}

	rps := make([]RoverPosition, 0, 0)

	for i := 1; i < len(instructions); {
		rp, err := getRoverPosition(instructions[i]) 
		if err != nil {
			return []RoverPosition{}, err
		}

		newRoverPosition, err := moveRover(rp, xMax, yMax, instructions[i + 1])
		if err != nil {
			return []RoverPosition{}, err
		}

		rps = append(rps, newRoverPosition)
		i += 2 
	}

	return rps, nil
}

func moveRover(r RoverPosition, xMax int, yMax int, instructions string) (RoverPosition, error) {
	for _, v := range instructions {
		switch {
		case v == 'L' && r.direction == 'N':
			r.direction = 'W'
		case v == 'L' && r.direction == 'W':
			r.direction = 'S'
		case v == 'L' && r.direction == 'S':
			r.direction = 'E'
		case v == 'L' && r.direction == 'E':
			r.direction = 'N'


		case v == 'R' && r.direction == 'N':
			r.direction = 'E'
		case v == 'R' && r.direction == 'E':
			r.direction = 'S'
		case v == 'R' && r.direction == 'S':
			r.direction = 'W'
		case v == 'R' && r.direction == 'W':
			r.direction = 'N'


		case v == 'M' && r.direction == 'N':
			r.y += 1
			if r.y > yMax {
				return RoverPosition{}, errors.New("rover out of Northern bounds")
			}
		case v == 'M' && r.direction == 'W':
			r.X -= 1
			if r.X < 0 {
				return RoverPosition{}, errors.New("rover out of Western bounds")
			}
		case v == 'M' && r.direction == 'S':
			r.y -= 1
			if r.y < 0 {
				return RoverPosition{}, errors.New("rover out of Southern bounds")
			}
		case v == 'M' && r.direction == 'E':
			r.X += 1
			if r.X > xMax {
				return RoverPosition{}, errors.New("rover out of Eastern bounds")
			}
		default:
			return RoverPosition{}, errors.New("invalid character")
		}
	}

	return r, nil
}

func getRoverPosition(position string) (RoverPosition, error) {
	var r RoverPosition
	storeK := 0
	spaceCount := 0
	var err error

	for k, v := range position {
		if v == ' ' {
			if spaceCount == 0 {
				spaceCount++
				storeK = k + 1
				tempX := position[:k]
				r.X, err = strconv.Atoi(tempX)
				if err != nil {
					return RoverPosition{}, fmt.Errorf("something wrong with strconv.Atoi(tempX)")
				}
				continue
			}
			if spaceCount == 1 {
				spaceCount++
				tempY := position[storeK:k]
				r.y, err = strconv.Atoi(tempY)
				if err != nil {
					return RoverPosition{}, fmt.Errorf("something wrong with strconv.Atoi(tempY)")
				}
				r.direction = position[len(position) - 1]

				return r, nil
			}
		}
	}

	fmt.Println("no spaces?")
	return RoverPosition{}, fmt.Errorf("something wrong")
}

func getGridMax(max string) (int, int, error) {
	i, j := 0, 0
	var err error
	for k, v := range max {
		fmt.Println(v)
		if v == ' ' {
			temp := max[:k]
			i, err = strconv.Atoi(temp)
			if err != nil {
				return 0, 0, fmt.Errorf("something wrong with strconv.Atoi(temp)/i")
			}
			temp = max[k+1:]
			j, err = strconv.Atoi(temp)
			if err != nil {
				return 0, 0, fmt.Errorf("something wrong with strconv.Atoi(temp)/j")
			}
			return i, j, nil
		}
	}

	return 0, 0, fmt.Errorf("something wrong with input data")
}