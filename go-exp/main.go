package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

// Anouncement: This is the place for daily coding improvement by solving some interesting challenges (almost everyday!).

// 11152022:
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

// NOTE: Pointer receiver's generic type-parameter cannot consume slices as it type.
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

// 11162022:
// Resource: Interface to string conversion (https://www.golinuxcloud.com/golang-interface-to-string/)

// Problem: Whenever a student levels up to an odd-grade, the tuition fee must be upgraded as follows:
// Mapping: `[1, 3, 5, 7, etc] -> [100, 100*3, 100*5, 100*7, etc]`, but the adjacent even-grades' tuition fee stays the same.

type GradeTuition map[int]interface{}

func isSlicesContains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func (gt GradeTuition) IsKeyExisted(key int) bool {
	totalGrades := len(gt)
	keySlices := make([]int, 0, totalGrades)
	for k := range gt {
		keySlices = append(keySlices, k)
	}

	return isSlicesContains(keySlices, key)
}

func (gt *GradeTuition) CalculateTotalFee(gradeNum int) *GradeTuition {
	if gradeNum == 0 || gt.IsKeyExisted(0) {
		return nil
	}

	var baseTuition = 100
	for grade := range *gt {
		if grade == 1 {
			(*gt)[grade] = baseTuition
		}
		if grade%2 == 1 {
			(*gt)[grade] = baseTuition * grade
			(*gt)[grade+1] = (*gt)[grade]
		}
	}

	return gt
}

// Core functions:
func printWithPattern(pattern string, data interface{}) {
	if strings.Compare("+", pattern) == 0 && data == nil {
		today := time.Now().String()
		fmt.Printf("%s %s\n", pattern, today)
	}

	if strings.Compare("=", pattern) == 0 && data != nil {
		fmt.Printf("%s>\t%+v\n", strings.Repeat(pattern, 5), data)
	}
}

func main() {
	// 11152022:
	printWithPattern("+", nil)
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
	printWithPattern("=", newFile)

	n, _ := gs.ExportToFile(newFile)
	printWithPattern("=", n)

	// 11162022:
	printWithPattern("+", nil)
	expGt := GradeTuition{
		1: "100",
		2: "100",
		3: "300",
		4: "300",
		5: "500",
		6: "500",
		7: "700",
	}
	testKeyExisted := expGt.IsKeyExisted(0)
	printWithPattern("=", testKeyExisted)

	gradeMap := GradeTuition{
		1: "",
		2: "",
		3: "",
		4: "",
		5: "",
	}
	printWithPattern("=", gradeMap.CalculateTotalFee(0))
	printWithPattern("=", gradeMap.CalculateTotalFee(3))
}
