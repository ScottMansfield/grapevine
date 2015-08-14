/**
 * Performs gossip conversation between two nodes
 */

package gossip

import "time"

// TODO: make this more sane later
import "../config"
import "../state"

type gossipCmd struct {
    hostData []state.HostData
}

type syncCmd struct {
    hostData []state.HostData
}

var timeToGossip <-chan time.Time
var gossip chan gossipCmd

var timeToSync <-chan time.Time
var sync chan syncCmd

func Initialize(conf config.Config) {
    // start the tickers
    ticker := time.NewTicker(conf.GossipInterval)
    timeToGossip = ticker.C
    gossip = make(chan gossipCmd)
    
    ticker = time.NewTicker(conf.SyncInterval)
    timeToSync = ticker.C
    sync = make(chan syncCmd)
    
    go gossipSelect()
    go performGossip()
    
    go syncSelect()
    go performSync()
}

// reads from timeToGossip, applies logic to select
// oher nodes to talk to, and forwards to gossip
func gossipSelect() {
    for {
        
    }
}

// reads from the gossip queue and gossips with all
// the specified neighbors
func performGossip() {
    for {
        
    }
}

// reads from timeToSync, applies logic to select
// another node to talk to, and forwards to sync
func syncSelect() {
    for {
        
    }
}

// reads from the sync channel and does a full sync
// with the specified neighbor
func performSync() {
    for {
        
    }
}
