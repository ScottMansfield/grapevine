/**
 * Contains the structures used for inter-process communication
 */
package gossip

// TODO: make this more sane later
import "../clock"
import "../state"

type FullData struct {
    Hosts []state.HostData
    Clock clock.Clock
}

type Delta struct {
    Host state.HostData
    OldState state.State
    NewState state.State
}
