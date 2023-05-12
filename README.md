
Mars Rover Challenge

Setup


Assumptions





Examples

Example curl to test "input" path

curl -X POST http://127.0.0.1:8095/input -H "Content-Type: application/json"  -d "{\"max-grid\": \"20 20\", \"rovers\": [{\"rover-start\": \"4 5 N\", \"movement\": \"LMMMRM\"}, {\"rover-start\": \"8 8 N\", \"movement\": \"LMMMRM\"}]}"


Example curl to test "file" path

curl -X POST http://127.0.0.1:8095/file  -d "file=squad-input"