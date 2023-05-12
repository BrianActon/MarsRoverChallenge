package rovers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Rover struct {
	RoverStart string `json:"rover-start"`
	Movement string   `json:"movement"`
}

type RoverSquadInput struct {
    MaxGrid string  `json:"max-grid"`
    Rovers  []Rover `json:"rovers"`
}

func Input(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Input")

    var rsi RoverSquadInput

    fmt.Println("r.Body", r.Body)
    err := json.NewDecoder(r.Body).Decode(&rsi)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if rsi.MaxGrid == "" {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Println("rsi = ", rsi)

    instructions := buildInstructions(rsi)

    fmt.Println("instructions = ", instructions)

    newPositions, err := move(instructions)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "New positions: %+v", newPositions)
}


func buildInstructions(rsi RoverSquadInput) []string {
	arr := make([]string, 0)

	arr = append(arr, rsi.MaxGrid)

	for _, v := range rsi.Rovers {
		fmt.Println("v = ", v)
		arr = append(arr, v.RoverStart)
		arr = append(arr, v.Movement)
	}

	fmt.Println("arr = ", arr)
	return arr
}