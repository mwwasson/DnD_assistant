package main

import (
    "os"
    "fmt"
    "github.com/mwwasson/DnD_5e"
)

func main () {
    
    fmt.Println("Iterating through all JSON files")

    files := [2]string{"./example/equipment_test.json", "./example/spell_test.json"}

    for _, element := range files {
        var data DnD_5e.Data

        fmt.Printf ("Opening %s\n",element)
        if data.Load_json_file(element) < 0 {
            os.Exit(1)
        }

        data.Print_rules()
        data.Print_gear()
        data.Print_armor()
        data.Print_weapons()
        data.Print_spells()
    }
    
}
