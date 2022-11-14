package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

// NOTE: `go help work` := A great leap forward/breakthrough/saltation in Golang for packages
// and directories management!

type GenericType interface {
	~int | ~uint | ~[]uint | string

	// NOTE: If we re-declare those association methods as follows, all of the
	// underlying types (includes: int64, uint64, etc) must be implemented those methods as well.
	//
	// CreateFile(string) (*os.File, error)
	// ExportToFile(*os.File) (int, error)
}

// NOTE: This type must be declared in the `GenericType` interface to be passed
// as a type-parameter for `GenericStruct[T]` (exp: `GenericStruct[OtherInteger]`).
// Reason: Because of mismatching underlying-type constraints.
type OtherInteger uint64

type GenericStruct[T GenericType] struct {
	Data     T
	Ctx      context.Context
	TestFunc func(output chan T, date time.Time) (string, error)
	// NOTE: Both function-typed properties have identical semantics.
	// TestFunc func(T, time.Time) (string, error)
}

func (gs *GenericStruct[T]) CreateFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	return file, err
}

// NOTE: Pointer receiver's generic type-parameter cannot consummed slices as it type.
func (gs *GenericStruct[T]) ExportToFile(file *os.File) (int, error) {
	fileContent := ConvertToPrimitive(gs.Data)
	numBytes, err := file.Write([]byte(fileContent))

	defer file.Close()

	return numBytes, err
}

func ConvertToPrimitive(data interface{}) string {
	if strings.Contains(reflect.TypeOf(data).String(), "map[string]interface{}") {
		destructure := data.(map[string]interface{})
		if genericData, ok := destructure["Data"].(string); ok {
			return genericData
		}
	}

	if strings.Contains(reflect.TypeOf(data).String(), "string") {
		return data.(string)
	}
	return *new(string)
}

// Unused structure:
type EmbeddedStruct struct {
	// NOTE: From line 25-28 so this was failed.
	// _ *GenericStruct[OtherInteger] `bson:",inline"`
	_ *GenericStruct[string] `bson:",inline"`
}

func main() {
	gs := GenericStruct[string]{
		Data: "I want something just like this! ~ One Republic",
		Ctx:  *new(context.Context),
		TestFunc: func(output chan string, date time.Time) (string, error) {
			curDate := time.Now().String()
			// Channel-typed string declaration syntax: `curDateStr := make(chan string)`.
			go func() {
				output <- curDate
			}()
			return <-output, nil
		},
	}

	fileName := "test.txt"
	newFile, _ := gs.CreateFile(fileName)
	fmt.Printf("=====>\t%+v\n", newFile)

	n, _ := gs.ExportToFile(newFile)
	fmt.Printf("=====>\t%+v\n", n)
}
