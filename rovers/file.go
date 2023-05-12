package rovers

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func File(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File")

	if err := r.ParseForm(); err != nil {
            fmt.Println("parsing failed")
            return
        }

	newFile := r.FormValue("file")

	instructions, err := fetchInstructions(newFile)
    if err != nil || instructions == nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Println("instructions = ", instructions)

    newPositions, err := move(instructions)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "New positions: %+v", newPositions)
}

func fetchInstructions(newFile string) ([]string, error) {
	fmt.Println("fetchInstructions")
	var err error
	var f *os.File

	if newFile == "test-squad-input" {
		f, err = os.Open(fmt.Sprintf("../files/%s", newFile))
		if err != nil {
			fmt.Println("not open here either", err)
		} else {
			fmt.Println("opened", f)
		}
	} else {
		f, err = os.Open(fmt.Sprintf("files/%s", newFile))
		if err != nil {
			return nil, err
		}
	}
	defer f.Close()

	fmt.Println("f = ", f)

	scanFile := bufio.NewScanner(f)
    scanFile.Split(bufio.ScanLines)

	fmt.Println("scanFile = ", scanFile)

    var instructions []string

    for scanFile.Scan() {
        instructions = append(instructions, scanFile.Text())
    }

    return instructions, nil
}