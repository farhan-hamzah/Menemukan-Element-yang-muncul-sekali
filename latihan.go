package main
import "fmt"

func main() {
    const NMAX int = 100
    var A [NMAX]int
    var i, n, j int
    var unik int
    var found bool

    fmt.Scan(&n)
    for i = 0; i < n; i++ {
        fmt.Scan(&A[i])
    }

    for i = 0; i < n; i++ {
        found = true
        for j = 0; j < n; j++ {
            if i != j && A[i] == A[j] {
                found = false
                break
            }
        }
        if found {
            unik = A[i]
            break
        }
    }

    fmt.Print(unik)
}