// Misc utilities

package marin

import (
    "fmt"
)

func PANIC_ON_ERR(err error) {
    if err != nil {
        panic(err)
    }
}

func MSG(f string, v ...interface{}) {
    if len(v) == 0 {
        fmt.Println(v)
    } else {
        fmt.Printf(f+"\n", v)
    }
}
