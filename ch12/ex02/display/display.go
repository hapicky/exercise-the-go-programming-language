package main

import (
	"fmt"
	"reflect"
	"strconv"

	"gopl.io/ch7/eval"
)

func main() {
	e, _ := eval.Parse("sqrt(A / pi)")
	Display("e", e)

	type point struct {
		x int
		y int
	}
	pm := make(map[point]bool)
	pm[point{x: 100, y: 50}] = true
	Display("pm", pm)

	am := make(map[[2]int]bool)
	k := [2]int{100, 50}
	am[k] = true
	Display("am", am)

	type node struct {
		link *node
	}

	node1 := node{}
	node2 := node{link: &node1}
	node1.link = &node2
	Display("node1", node1)
}

var counter int
const step_limit = 100

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	counter = 0
	display(name, reflect.ValueOf(x))
}

func display(path string, v reflect.Value) {
	counter++
	if counter > step_limit {
		panic("Display: exceed step limit.")
	}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(+%s", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default:
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		s := v.Type().Name() + "{"
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				s += ", "
			}
			s += fmt.Sprintf("%s:%s", v.Type().Field(i).Name, formatAtom(v.Field(i)))
		}
		return s + "}"
	case reflect.Array:
		s := "["
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				s += ", "
			}
			s += formatAtom(v.Index(i))
		}
		return s + "]"
	default:
		return v.Type().String() + " value"
	}
}
