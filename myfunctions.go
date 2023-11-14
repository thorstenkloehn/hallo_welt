// layouts/myfunctions.go
package github.com/thorstenkloehn/hallo_welt

import "fmt"

func SayHello(name string) string {
    return fmt.Sprintf("Hallo, %s!", name)
}
