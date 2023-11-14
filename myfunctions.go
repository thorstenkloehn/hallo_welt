// layouts/myfunctions.go
package main

import "fmt"

func SayHello(name string) string {
    return fmt.Sprintf("Hallo, %s!", name)
}