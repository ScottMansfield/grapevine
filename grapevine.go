/**
 * Main grapevine entry point. Bootstraps and establishes communication.
 */
 
package main

import "fmt"

// TODO: make this more sane later
import "./clock"
import "./config"
//import "./gossip"
//import "./state"

func main() {
    conf, err := config.Parse("config.cfg")
    if err != nil { panic(err) }
    
    fmt.Printf("%v\n", conf)
    
    clock := &clock.Clock{ }
    fmt.Printf("Size: %v | Epoch: %v", clock.Size, clock.Epoch)
}