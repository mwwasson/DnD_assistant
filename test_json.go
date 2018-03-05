package main

import (
    "os"
    "fmt"
)

func main () {
    
 //   fmt.Println(len(os.Args), os.Args)

    for i, file_name := range os.Args {
        if i == 0 {
            continue
        }
        fmt.Printf("Opening %s\n",file_name)
        var data Data
        if data.Load_json_file(file_name) < 0 {
            fmt.Println("Could not open file.\n")
            continue
        } 
        data.Print_rules()
        data.Print_gear()
        data.Print_armor()
        data.Print_weapons()
        data.Print_spells()
    }    
}
