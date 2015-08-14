/**
 * Holds basic types for holding state
 */
package state

type State int
const (
    UP State = iota
    DOWN
    DEAD
    UNKNOWN
    SUSPICIOUS
)

type HostData struct {
    DNSName string
    Port int
    State State
}

func (hd *HostData) Equals(other HostData) bool {
    return hd.DNSName == other.DNSName &&
           hd.Port    == other.Port    &&
           hd.State   == other.State
}
