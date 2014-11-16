package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    v := strings.Fields(s);
    ret := make(map[string]int);
    for i := range v {
        ret[v[i]]++;
    }
    return ret;
}

func main() {
    wc.Test(WordCount)
}