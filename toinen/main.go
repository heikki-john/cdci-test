package main

import "fmt"

func main() {
    fmt.Println("Moikka")
    msg := sayHello("Alice")
    fmt.Println(msg)
}

func sayHello(name string) string {
    return fmt.Sprintf("Hello %s", name)
}
