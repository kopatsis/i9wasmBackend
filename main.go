package main

import (
	"encoding/json"
	"syscall/js"
)

type InputA struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type InputB struct {
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}

type Output struct {
	CombinedName    string `json:"combined_name"`
	TotalValue      int    `json:"total_value"`
	CombinedDetails string `json:"combined_details"`
}

func processJSON(this js.Value, args []js.Value) interface{} {
	var inputA InputA
	json.Unmarshal([]byte(args[0].String()), &inputA)

	var inputB InputB
	json.Unmarshal([]byte(args[1].String()), &inputB)

	output := Output{
		CombinedName:    inputA.Name + " & " + inputB.Description,
		TotalValue:      inputA.Value + inputB.Amount,
		CombinedDetails: inputA.Name + ": " + inputB.Description,
	}

	outputJSON, _ := json.Marshal(output)

	return js.ValueOf(string(outputJSON))
}

func main() {
	js.Global().Set("processJSON", js.FuncOf(processJSON))

	select {}
}
