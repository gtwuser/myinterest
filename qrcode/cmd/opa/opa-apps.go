package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/open-policy-agent/opa/rego"
	"io/ioutil"
)

func main() {
	ctx := context.Background()
	r := rego.New(
		rego.Query(`s = data.example.public_server`),
		rego.Load([]string{`example.rego`}, nil),
	)
	query, _ := r.PrepareForEval(ctx)

	var input interface{}
	bs, _ := ioutil.ReadFile(`v1-data-input.json`)
	_ = json.Unmarshal(bs, &input)

	rs, _ := query.Eval(ctx, rego.EvalInput(input))
	fmt.Println(rs[0].Expressions)
	fmt.Println(rs[0].Bindings[`s`])
}

func logErrorF(msg string, args ...interface{}) {
	fmt.Printf(msg+`\n`, args)
}
