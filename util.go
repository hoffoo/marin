// Misc utilities

package marin

import (
    "fmt"
)

// panic if err is not nil
func PANIC_ON_ERR(err error) {
    if err != nil {
        panic(err)
    }
}

// a debug function to print varidic printf aruments
func MSG(f string, v ...interface{}) {
    if len(v) == 0 {
        fmt.Println(f)
    } else {
        fmt.Printf(f+"\n", v)
    }
}

// a debug function for prototyping: make go not complain about
// unused variables
func USE(i ...interface{}) {
    return
}
