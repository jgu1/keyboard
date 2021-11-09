// Package keyboard by jgu1
package keyboard

import (
    "bufio"
    "os"
    "strconv"
    "strings"
    "fmt"
)

// GetFloat by jgu1
func GetFloat() (float64,error) {
    fmt.Println("welcome to use jgu1's package!")
    reader := bufio.NewReader(os.Stdin)
    input,err := reader.ReadString('\n')
    if err != nil {
        return 0, err
    }
    
    input = strings.TrimSpace(input)
    number, err := strconv.ParseFloat(input, 64)
    if err != nil {
        return 0, err
    }
    return number*10, nil
}
