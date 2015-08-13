/**
 * Main grapevine entry point. Bootstraps and establishes communication.
 */
 
package main;

import "fmt"

import "./clock"
import "./config"

func main() {
    conf, err := config.Parse("config.cfg")
    if err != nil { panic(err) }
    
    fmt.Printf("%v\n", conf)
    
    clock := &clock.Clock{ }
    fmt.Println(clock.Size())
}