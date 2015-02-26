package main 

import (
    "flag"
    "fmt"
)

const APP_VERSION = "0.1"
var siteFlag = flag.String("site", "", "must be test.com www.test.com")
// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
    
//    site := *siteFlag
    
    
    
    
    
    
    
    
}

