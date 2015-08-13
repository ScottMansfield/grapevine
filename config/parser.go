/**
 * Parser for config file
 */
package config;

import "bufio"
import "fmt"
import "os"
import "strings"

func Parse(fileName string) (Config, error) {
    fmt.Println("Parsing", fileName)
    var retval Config
    
    file, err := os.Open(fileName)
    if err != nil { return retval, err }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    
    for scanner.Scan() {
        line := scanner.Text()
        
        parts := strings.SplitN(line, "=", 2)
        
        switch parts[0] {
            case "seeds":
                retval.Seeds = strings.Split(parts[1], ",")
        }
    }
    
    return retval, nil
}