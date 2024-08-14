package utils

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/Knetic/govaluate"
)

type URLParams struct {
	UserID   int
	Age      int
	Greeting string
}

// urlTemplate := "https://example.com/api/user/{{.UserID}}/details?age={{.Age}}&greeting={{.Greeting}}"
func evaluateExpression(expression string, params URLParams) string {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		panic(err)
	}

	parameters := make(map[string]interface{}, 8)
	parameters["UserID"] = params.UserID
	parameters["Age"] = params.Age
	parameters["Greeting"] = params.Greeting

	result, err := expr.Evaluate(parameters)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%v", result)
}
func EvalTest(expression string) (*string, error) {

	params := URLParams{
		UserID:   12345,
		Age:      30,
		Greeting: evaluateExpression(`"Hello " + "World!"`, URLParams{}),
	}

	tmpl, err := template.New("url").Parse(expression)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer

	err = tmpl.Execute(&b, params)
	if err != nil {
		return nil, err
	}
	s := b.String()
	return &s, nil
}
