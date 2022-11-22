package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Resources:
// + https://www.golinuxcloud.com/golang-context/
// + https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go
// + https://btree.dev/golang-context

// Reasons to use context.Context:
// + It simplifies the implementation for deadlines and cancellation across your processes or API.
// + It prepares your code for scaling, for example, using Context will make your code clean and easy to manipulate in the future by chaining all your process in a child parent relationship, you can tie/join any process together.
// + It’s easy to use.
// + Goroutine safe, i.e you can run the same context on different goroutines without leaks.

// Deadline: Returns the time when the context should be cancelled, together with a Boolean that is false when there is no deadline.
// Done: Returns a receive-only channel of empty structs, which signals when the context should be cancelled.
// Err: Returns nil while the done channel is open; otherwise it returns the cause of the context cancellation.
// Value: Returns a value associated with a key for the current context, or nil if there's no value for the key.

// Best Practice when using context.Context:
// + Whenever you need to use context.Context, make sure it always the first argument.
// + Always use “ctx” it will work perfectly if you use another variable name but just follow the majority, you don’t need to be unique with things like this.
// + Make sure the cancel function is called.
// + Don’t use a struct to add a context in a method, always add it to the argument i.e context.Context.
// + Don’t overuse context.WithValue.

func revealCtxBackground(ctx context.Context) {
	printWithPattern("=", "ctx.Err()", ctx.Err())
	printWithPattern("=", "ctx.Done()", ctx.Done())
	printWithPattern("=", "ctx.Value(\"key\")", ctx.Value("key"))

	time, ok := ctx.Deadline()
	res := fmt.Sprintf("%+v, %+v", time, ok)
	printWithPattern("=", "ctx.Deadline()", res)
}

func revealCtxWithTicker(ctx context.Context) {
	done := ctx.Done()
	// If we do not set the upper bound for this loop, this context of looping through
	// all existing integers might continue infinitely because the context is never completed.
	for sec := 0; sec < 8; sec++ {
		select {
		case <-done:
			return
		case <-time.After(time.Second):
			printWithPattern("=", "Tick (s)", sec)
		}
	}
}

const SHORT_DURATION = 100 * time.Millisecond

func revealCtxWithDeadline(ctx context.Context) {
	duration := time.Now().Add(SHORT_DURATION)
	ctxWithDeadline, cancel := context.WithDeadline(ctx, duration)

	// Even though ctxWithDeadline will expire in the near future,
	// it is good practice to call its cancellation function in any case.
	// If this couldn't be satisfied (not done), the context and its parents would live longer than necessary.
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		printWithPattern("=", "Context's duration", "Overslept!")
	case <-ctxWithDeadline.Done():
		// Error: exceeded/violated the dealine constraint.
		printWithPattern("=", "Error", ctxWithDeadline.Err())
	}
}

func revealCtxWithTimeout(ctx context.Context) {
	// Pass a ctxWithTimeout with a timeout duration to tell a blocking function that it should abandon
	// its work after the timeout elapses.
	ctxWithTimeout, cancel := context.WithTimeout(ctx, SHORT_DURATION)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		printWithPattern("=", "Context's duration", "Overslept!")
	case <-ctxWithTimeout.Done():
		// Error: exceeded/violated the dealine constraint.
		printWithPattern("=", "Error: ", ctxWithTimeout.Err())
	}
}

// NOTE:
// Return type's convention:
//	`func(context.Context) chan int` || `func(ctx context.Context) chan int`
// equals with:
//	`func(context.Context) <-chan int` || `func(ctx context.Context) <-chan int`

func revealCtxWithCancelSig(ctx context.Context) {
	// Anonymous function genInts generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of genInts need to cancel the context once they are done consuming
	// generated integers so as not to leak the internal goroutine started by genInts.
	genInts := func(ctxGenRoutine context.Context) <-chan int {
		dest := make(chan int)
		counter := 1
		go func() {
			for {
				select {
				case <-ctxGenRoutine.Done():
					return // Returning to avoid leaking the goroutine.

				// Increases an integer counter across multiple goroutine and sends it to the returned channel.
				// To use the gen method properly, you need to cancel the context immediately when they are consuming the integers,
				// so we can avoid any leakage in the goroutines internally.
				case dest <- counter:
					counter++
				}
			}
		}()
		return dest
	}

	// Cancel the current context whenever we are finished consuming integers
	// in each goroutine.
	ctxWithCancelSig, cancel := context.WithCancel(ctx)

	defer cancel()

	for num := range genInts(ctxWithCancelSig) {
		printWithPattern("=", "Generated number", num)
		if num == 5 {
			break
		}
	}
}

type KeyStore string

const SESSION_ID KeyStore = "12345678"

func revealCtxWithValue(ctx context.Context) {
	// This context's specialize method can be used to modify the underlying value of a constant.
	// After creating the context.WithValue, passed in the parent context.Background, and added a SESSION_ID key with the value of itself,
	// by joining this context to all our request, we can easily grab the SESSION_ID
	ctxWithValue := context.WithValue(ctx, SESSION_ID, "SESSION_ID's value now ~ 123456789")
	sessionID := ctxWithValue.Value(SESSION_ID)
	strSessionID, retrievable := sessionID.(string)
	if !retrievable {
		log.Fatalln("Error: sessionID's underlying type was not string")
	}
	printWithPattern("=", "Value", strSessionID)
}
