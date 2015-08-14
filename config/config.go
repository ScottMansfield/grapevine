/**
 * Config model
 */
package config

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"
import "time"

type Config struct {
    Seeds          []string
    Port           int
    GossipInterval time.Duration
    SyncInterval   time.Duration
}

func Parse(fileName string) (Config, error) {
    fmt.Println("Parsing", fileName)
    var retval Config
    
    file, err := os.Open(fileName)
    if err != nil { return retval, err }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        
        if strings.HasPrefix(line, "#") || line == "" {
            continue
        }
        
        parts := strings.SplitN(line, "=", 2)
        
        switch strings.TrimSpace(parts[0]) {
            case "seeds":
                retval.Seeds = strings.Split(strings.TrimSpace(parts[1]), ",")
                
            case "port":
                val, err := strconv.Atoi(strings.TrimSpace(parts[1]))
                if err != nil { return retval, err }
                retval.Port = val
                
            case "gossipInterval":
                val, err := strconv.Atoi(strings.TrimSpace(parts[1]))
                if err != nil { return retval, err }
                retval.GossipInterval = time.Duration(val) * time.Millisecond
                
            case "syncInterval":
                val, err := strconv.Atoi(strings.TrimSpace(parts[1]))
                if err != nil { return retval, err }
                retval.SyncInterval = time.Duration(val) * time.Millisecond
        }
    }
    
    return retval, nil
}
