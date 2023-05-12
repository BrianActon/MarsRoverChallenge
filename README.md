
Mars Rover Challenge

Setup

1. Do a git pull into your local repository
``` 
git pull https://github.com/BrianActon/MarsRoverChallenge 
```

2. cd into this directory and run 'go test ./...'

3. From the command line within in the mars-rover-challenge folder, you can run 'go run main.go'

4. Open a second terminal, and from the command line you can run the curl examples listed below.


Examples

Example curl to test "input" path

curl -X POST http://127.0.0.1:8095/input -H "Content-Type: application/json"  -d "{\"max-grid\": \"20 20\", \"rovers\": [{\"rover-start\": \"4 5 N\", \"movement\": \"LMMMRM\"}, {\"rover-start\": \"8 8 N\", \"movement\": \"LMMMRM\"}]}"


Example curl to test "file" path

curl -X POST http://127.0.0.1:8095/file  -d "file=squad-input"



Assumptions

- Each rover is identified by its positioning
- If a rover attempts to move beyond the boundries of the current grid, then its position is not updated
- Sending instructions to rovers on another planet could have a few unknown difficulties, so I built in 2 methods of getting the data processed. Via JSON input and via a file as input. 
- The entity in control of data going into the app


Customizing rover starting positions

Simply edit files/squad-input, and then run "curl -X POST http://127.0.0.1:8095/file  -d "file=squad-input""


