package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
      return true
    }

    anchor, runner := 0, 0

    for ; runner <= len(t) - 1; runner++ {
        if s[anchor] == t[runner] {
            if anchor == len(s) - 1 {
                return true
            }

            anchor++
        }
    }

    return false
}

func main() {
	fmt.Println(isSubsequence("axc", "ahbgdc"));
}
