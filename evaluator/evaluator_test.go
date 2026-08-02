package evaluator

import (
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"testing"
)

func TestEvalBasics(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"5 + 5 + 5 - 10", int64(5)},
		{"2 * (5 + 10)", int64(30)},
		{"true == false", false},
		{"if (1 < 2) { 10 } else { 20 }", int64(10)},
		{"let a = 5; let b = a * 2; b;", int64(10)},
		{"let add = fn(a, b) { a + b; }; add(2, 3);", int64(5)},
		{"let newAdder = fn(x) { fn(y) { x + y; } }; let addTwo = newAdder(2); addTwo(3);", int64(5)},
		{`"hello" + " " + "world"`, "hello world"},
		{`"hello" == "hello"`, true},
		{"len([1, 2, 3])", int64(3)},
		{"[1, 2, 3][1]", int64(2)},
		{`{"one": 1, "two": 2}["two"]`, int64(2)},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testExpectedObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testExpectedObject(t *testing.T, obj object.Object, expected interface{}) {
	t.Helper()

	switch expected := expected.(type) {
	case int64:
		result, ok := obj.(*object.Integer)
		if !ok {
			t.Fatalf("object is not Integer. got=%T (%+v)", obj, obj)
		}
		if result.Value != expected {
			t.Fatalf("object has wrong value. got=%d, want=%d", result.Value, expected)
		}
	case bool:
		result, ok := obj.(*object.Boolean)
		if !ok {
			t.Fatalf("object is not Boolean. got=%T (%+v)", obj, obj)
		}
		if result.Value != expected {
			t.Fatalf("object has wrong value. got=%t, want=%t", result.Value, expected)
		}
	case string:
		result, ok := obj.(*object.String)
		if !ok {
			t.Fatalf("object is not String. got=%T (%+v)", obj, obj)
		}
		if result.Value != expected {
			t.Fatalf("object has wrong value. got=%q, want=%q", result.Value, expected)
		}
	}
}
