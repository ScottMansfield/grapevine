/**
 * Vector clock implementation
 */
package clock

import "math"

// Model and instance
type Clock struct {
    // TODO: wtf do I put here
    Size int
    Epoch int64
}

var clock *Clock

// Operation queues and message structs
type UpdateMsg struct {
    Respond chan Clock
}
var Update chan UpdateMsg

type SetMsg struct {
    Epoch int64
    Respond chan Clock
}
var Set chan SetMsg

type GetMsg struct {
    Respond chan Clock
}
var Get chan GetMsg

func init() {
    // init channels
    Update = make(chan UpdateMsg)
    Set = make(chan SetMsg)
    Get = make(chan GetMsg)
    
    // Create clock manager
    go clockManager()
}

// Serializes all interaction with the clock
// I'm sure there's some kind of insidious bug hidden
// in this code somewhere...
func clockManager() {
    var clock Clock
    
    for {
        select {
            case updateReq := <-Update:
                clock.Epoch += 1
                updateReq.Respond <- clock
            case setReq := <-Set:
                clock.Epoch = int64(math.Max(float64(clock.Epoch), float64(setReq.Epoch)))
                setReq.Respond <- clock
            case getReq := <-Get:
                getReq.Respond <- clock
        }
    }
}
