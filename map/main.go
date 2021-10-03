package main

import (
    "fmt"
)

type A struct {
    val string
}

func(a A) Print() {
    fmt.Printf("A.Print: %s\n", a.val)
}

func forEach[T any](l []T, f func(T)) {
    for _, n := range l {
        f(n)
    }
}

func mapFunc[T1, T2 any](l []T1, f func(T1) T2) []T2 {
    ret := make([]T2, 0)
    for _, n := range l {
        ret = append(ret, f(n))
    }
    return ret

}

func main() {
    vi := []int{0, 1, 2, 3, 4, 5}
    vs := []string{"a", "hoge", "huga"}
    va := []A{A{"a"}, A{"b"}}

    viMap := mapFunc(vi, func(i int) int {return i * 2 })
    vsMap := mapFunc(vs, func(s string) string {return s + "_suff"})
    vaMap := mapFunc(va, func(a A) A {return A{a.val + "A"}})

    fmt.Printf("int: %v, type: %T\n", viMap, viMap)
    fmt.Printf("string: %v, type: %T\n", vsMap, vsMap)
    fmt.Printf("A: %v, type: %T\n", vaMap, vaMap)

    forEach(va, func(a A) {a.Print()})
}
