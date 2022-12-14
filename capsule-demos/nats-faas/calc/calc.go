package main

import (
	/* string to json */
	"github.com/tidwall/gjson"
	/* create json string */
	"github.com/tidwall/sjson"

	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.OnNatsMessage(Handle)
}

func Handle(params []string) {
	natsMessage := params[0]
	hf.Log("π subject: " + hf.NatsGetSubject() + ", arguments: " + natsMessage)

	operation := gjson.Get(natsMessage, "operation").String()
	operand1 := gjson.Get(natsMessage, "operand1").Float() 
	operand2 := gjson.Get(natsMessage, "operand2").Float() 

	var res float64

	switch operation {
	case "+":
		res = operand1 + operand2
	case "-":
		res = operand1 - operand2
	case "*":
		res = operand1 * operand2
	case "/":
		res = operand1 / operand2
	default:
		res = 0.0
	}

	result := `{"result": ""}`
	result, _ = sjson.Set(result, "result", res)

	hf.Log("#οΈβ£ result: " + result)

	hf.NatsReply(result, 10)

}

//export OnLoad
func OnLoad() {
	hf.Log("π I'm the calc function")
	hf.Log("πListening on: " + hf.NatsGetSubject())
	hf.Log("π NATS server: " + hf.NatsGetServer())

}

//export OnExit
func OnExit() {
	hf.Log("ππ€ Bye! Have a nice day")
}
