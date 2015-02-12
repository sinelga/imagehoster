package main 

import (
    "flag"
    "fmt"
    "startones"
    "contents_feeder/find_all_characters"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
    flag.Parse() // Scan the arguments list 

    if *versionFlag {
        fmt.Println("Version:", APP_VERSION)
    }
    
    
    golog,config := startones.Start()
    
    characters := find_all_characters.FindAll(golog,config)
    
    
    for _,character := range characters {
    	
    	
    	fmt.Println(character.Name);
    	
    	
    } 
    
    
    
    
}

