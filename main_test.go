// thanks https://www.pullrequest.com/blog/unit-testing-in-go/ 
package main_test

import (
	"fmt"
	"testing"
	"https://github.com/MakeMeSenpai/MakeUtility/blob/master/main.go"
)

// written out
func TestHello(t *testing) {
    actual := main.HelloWorld()
    if actual != "hello world" {
        t.Errorf("expected 'hello world', got '%s'", actual)
    }
}

// shorthand
assert.Equal(t, "hello world", main.HelloWorld())

// returns coverage of sequential tests
func Coverage()
