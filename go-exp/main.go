package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unsafe"
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
	extraData, _ := gs.TestFunc(make(chan T), time.Now())
	numBytes, err := file.Write([]byte(fileContent + extraData))

	// NOTE:
	// - The `defer` statement is used to prevent the execution of a function until all other functions are executed.
	// - When multiple `defer` are invoked in a program, the order of execution of the `defer` statements will be LIFO (Last In First Out).
	defer file.Close()

	return numBytes, err
}

func ConvertToPrimitive(data interface{}) string {
	if strings.Contains(reflect.TypeOf(data).String(), "map[string]interface{}") {
		// NOTE: Go Type Assertions.
		destructure := data.(map[string]interface{})
		if genericData, ok := destructure["Data"].(string); ok {
			return genericData
		}
	}

	if strings.Contains(reflect.TypeOf(data).String(), "string") {
		// NOTE: Go Type Assertions.
		return data.(string)
	}
	return *new(string)
}

// Unused structure:
type EmbeddedStruct struct {
	// NOTE: From line 30-33 so this below field declaration was failed.
	// _ *GenericStruct[OtherInteger] `bson:",inline"`
	_ *GenericStruct[string] `bson:",inline"`
}

func GenericTernary[T any](condition bool, vTrue, vFalse T) T {
	if condition {
		return vTrue
	}
	return vFalse
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

func (gt *GradeTuition) CalculateFeeEachGrade(gradeNum int) *GradeTuition {
	if gradeNum == 0 || gt.IsKeyExisted(0) {
		return nil
	}

	for g := 1; g <= gradeNum; g++ {
		(*gt)[g] = ""
	}
	fmt.Printf("%+v\n", *gt)

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

// 11172022:
// Concurrency, channel, and blocked-channel in Golang demystification/proclamation by example:

func printChanVal(val chan int) int {
	val <- 10
	return <-val
}

func insertMultiChans(left, right chan int) {
	// NOTE: In 10e4 channels, 10e4 routines block each other in parallel.
	// --> So all routines exist in concurrency at the same time.
	left <- 1 + <-right // Equals to: `1 + right`.
}

// The core functions begin here:

func handlePanic() {
	r := recover()
	if r != nil {
		fmt.Println("INFO: Process recovery: ", r)
	}
}

func printWithPattern(pattern, desc string, data interface{}) {
	// NOTE: Execute the handlePanic even after panic occurs.
	defer handlePanic()

	if strings.Compare("+", pattern) == 0 && data == nil && desc == "" {
		today := time.Now().String()
		fmt.Printf("%s %s\n", pattern, today)
	}

	if strings.Compare("=", pattern) == 0 && data != nil && desc != "" {
		dataWithDesc := fmt.Sprintf("%s: %v", desc, data)
		fmt.Printf("%s>\t%+v\n", strings.Repeat(pattern, 5), dataWithDesc)
	}

	if strings.Compare("", pattern) == 0 {
		// NOTE: The `panic` statement is used to stop/terminate our program immediately.
		panic("Wrong pattern indication!")
	}
}

func main() {
	// 11152022:
	printWithPattern("+", "", nil)
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
	printWithPattern("=", "Newfile's address", newFile)

	n, _ := gs.ExportToFile(newFile)
	printWithPattern("=", "File's content length", n)

	// Link: https://golangdocs.com/converting-string-to-integer-in-golang
	// Or we can use:
	// testConv, _ = strconv.Atoi("100")
	testConv, _ := strconv.ParseInt("100", 16, 64)
	testTernaryOperator := GenericTernary(n > 1, int64(n)+testConv, 1)
	printWithPattern("=", "Ternary function execution output", testTernaryOperator)

	// 11162022:
	printWithPattern("+", "", nil)
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
	printWithPattern("=", "Key existed", testKeyExisted)

	// NOTE: Cannot use `gradeMap := new/*new(GradeTuition)` -> because we couldn't assign values to nil map.
	gradeMap := make(GradeTuition)
	printWithPattern("=", "Calculate with grade zero", gradeMap.CalculateFeeEachGrade(0))
	printWithPattern("=", "Calculate with positive grade number", gradeMap.CalculateFeeEachGrade(10))

	// 11172022:
	printWithPattern("+", "", nil)
	testChanVal := make(chan int)

	// NOTE: If we are not using keyword `go` to invoke new go-thread/goroutine,
	// this error will be appeared: `fatal error: all goroutines are asleep - deadlock!`.
	// Exp: `insertChanneling(testChanVal, nil)`.
	go printChanVal(testChanVal)                                                        // This segment didn't return anything to stdout (?).
	printWithPattern("=", "Total goroutines", runtime.Stack(make([]byte, 0, 10), true)) // Return: 0.

	// Another example:
	leftMost := make(chan int) // The first goroutine in the list.
	left := leftMost           // The left one.
	right := leftMost          // The right one.

	// NOTE: In 10e4 channels, 10e4 routines block each other in parallel.
	// --> So all routines exist in concurrency at the same time.
	const routines = 10e4
	for i := 0; i < routines; i++ {
		go insertMultiChans(left, right)
		left = right
	}

	// Release the last or most recently invoked routine (the "most right" one, horizontally).
	go func(val chan int) {
		val <- 1
	}(right)

	// NOTE: After the latest routine has been released, emitted or unveiled,
	// all of the remaining ones will be executed as normal.
	printWithPattern("=", "Channel's address", leftMost)
	printWithPattern("=", "Channel's final value", <-leftMost)

	// 11222022:
	printWithPattern("+", "", nil)
	var ctx = context.Background()

	revealCtxBackground(ctx)
	revealCtxWithTicker(ctx)
	revealCtxWithDeadline(ctx)
	revealCtxWithTimeout(ctx)
	revealCtxWithCancelSig(ctx)
	revealCtxWithValue(ctx)

	// 12012022:
	printWithPattern("+", "", nil)

	var tmpArr0 = [...]int{1, 2, 3}
	var tmpArr1 = []int{4, 5, 6}
	printWithPattern("=", "Print array with its length", fmt.Sprintf("%#v", tmpArr0))
	printWithPattern("=", "Print array with its length", fmt.Sprintf("%#v", tmpArr1))
}
