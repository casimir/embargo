// Package eval provides a reflective evaluator for structs.
package eval

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

// NopeStr is the value return if there is an invalid methdod.
const NopeStr = "<UNKNOWN>"

// Evaluator is implemented by any value that as an Eval method. The Eval
// method is used to evaluate a list of tokens corresponding to
// attributs/method of the value.
type Evaluator interface {
	Eval(args ...string) string
}

func evalAttr(obj interface{}, name string) string {
	if len(name) <= 0 {
		return NopeStr
	}
	val := reflect.Indirect(reflect.ValueOf(obj))
	field := val.FieldByName(toExportCase(name))
	if !field.IsValid() {
		log.Println("Invalid field:", name)
		return NopeStr
	}
	return fmt.Sprint(field.Interface())
}

func evalMethod(obj interface{}, name string, args ...string) string {
	if len(name) <= 0 {
		return NopeStr
	}
	val := reflect.Indirect(reflect.ValueOf(obj))
	method := val.MethodByName(toExportCase(name))
	if !method.IsValid() {
		log.Println("Invalid method:", name)
		return NopeStr
	}
	if err := checkSignature(method.Type(), args); err != nil {
		log.Printf("Signature mismatch for method: %s → %s", name, err)
		return NopeStr
	}

	in := toValues(args...)
	out := method.Call(in)
	return fmt.Sprint(out[0])
}

func checkSignature(tmethod reflect.Type, params []string) error {
	if tmethod.NumIn() != len(params) || tmethod.NumOut() < 1 {
		return errors.New(fmt.Sprintf("Count mismatch: {in: %d, out: %d} ≠ {in: %d, out: %d}",
			tmethod.NumIn(), len(params), tmethod.NumOut(), 1))
	}
	if tmethod.Out(0).Kind() != reflect.String {
		return errors.New(fmt.Sprintf("Type mismatch (return): %s ≠ string",
			tmethod.Out(0).Kind()))
	}
	for i, it := range params {
		if tmethod.In(i) != reflect.TypeOf(it) {
			return errors.New(fmt.Sprintf("Type mismatch (%d): %s ≠ %s",
				i+1, tmethod.In(i).Kind(), reflect.TypeOf(it).Kind()))
		}
	}
	return nil
}

func toExportCase(name string) string {
	return strings.ToUpper(name[:1]) + name[1:]
}

func toValues(args ...string) []reflect.Value {
	values := []reflect.Value{}
	for _, it := range args {
		values = append(values, reflect.ValueOf(it))
	}
	return values
}

// Struct is a generic implementation of Evaluator. It can be used as a wrapper
// for any type.
type Struct struct{ obj interface{} }

// Eval return the attribute value or the method return corresponding to args.
func (e *Struct) Eval(args ...string) string {
	if len(args) == 1 {
		return evalAttr(e.obj, args[0])
	} else {
		return evalMethod(e.obj, args[0], args[1:]...)
	}
}

// New create a new Struct.
func New(obj interface{}) Evaluator {
	return &Struct{obj: obj}
}

var modules map[string]Evaluator = map[string]Evaluator{}

// Register adds a module to the list used by Eval().
// Modules are recognized by its name. Adding a module with the same name will
// override the previous one.
func Register(name string, module Evaluator) { modules[name] = module }

// Eval tries to evaluate input with every registered modules.
func Eval(input string) string {
	vars := Parse(input)
	out := input

	for k, v := range vars {
		moduleName := v.Module
		e, ok := modules[moduleName]
		if !ok {
			log.Println("Unknown module: " + moduleName)
			continue
		}
		out = strings.Replace(out, k, e.Eval(v.Tokens...), -1)
	}
	return out
}
