package DnD_5e

import (
    "fmt"
)

func (container *Data) Print_rules() int {
    if len(container.Rule) <= 0 {
        return -1
    }
    fmt.Printf("Rules:\n")
    for _,element := range container.Rule{
        fmt.Printf("Title: %s, Description: %s\n", element.Title, element.Desc)
        fmt.Printf("\n")
    }
    fmt.Printf("----------------------------------------\n\n")
    return 0
}

func (container *Data) Print_gear() int {
    if len(container.Gear) <= 0 {
        return -1
    }
    fmt.Printf("Adventuring Gear:\n")
    for _,element := range container.Gear{
        fmt.Printf("Item: %s\n", element.Item)
        fmt.Printf("Cost: %.2f gp\n", element.Cost)
        fmt.Printf("Weight: %.2f lbs\n", element.Weight)
        if element.Desc != "" {
            fmt.Printf("Description: \n%s\n", element.Desc)
        }
        fmt.Printf("\n")
    }
    fmt.Printf("----------------------------------------\n\n")
    return 0
}

func (container *Data) Print_armor() int {
    if len(container.Armor) <= 0 {
        return -1
    }
    fmt.Printf("Armor:\n")
    for _,element := range container.Armor{
        fmt.Printf("Type: %s\n", element.Type)
        fmt.Printf("Name: %s\n", element.Name)
        fmt.Printf("Cost: %.2f gp\n", element.Cost)
        fmt.Printf("AC %s\n", element.AC)
        fmt.Printf("Str: ")
        if element.Strength > 0 {
            fmt.Printf("%d\n", element.Strength)
        } else {
            fmt.Printf("-\n")
        }

        fmt.Printf("Stealth: %s\n", element.Stealth)
        fmt.Printf("Weight: %.1f lbs\n", element.Weight)
        fmt.Printf("\n")
    }
    fmt.Printf("----------------------------------------\n\n")
    return 0
}

func (container *Data) Print_weapons() int {
    if len(container.Weapon) <= 0 {
        return -1
    }
    fmt.Printf("Weapons:\n")
    for _,element := range container.Weapon{
        fmt.Printf("Type: %s\n", element.Type)
        fmt.Printf("Name: %s\n", element.Name)
        fmt.Printf("Cost: %.2f gp\n", element.Cost)
        fmt.Printf("Damage: %s\n", element.Damage)
        fmt.Printf("Weight: %.2f lbs\n", element.Weight)
        fmt.Printf("Properties: %s\n", element.Prop)
        if element.Desc != "" {
            fmt.Printf("Description: \n%s\n", element.Desc)
        }
        fmt.Printf("\n")
    }
    fmt.Printf("----------------------------------------\n\n")
    return 0
}


func (container *Data) Print_spells() int {
    if len(container.Spell) <= 0 {
        return -1
    }
    fmt.Printf("Spells:\n")
    for _,element := range container.Spell{
        fmt.Printf("Title: %s\n", element.Title)
        fmt.Printf("Level: %d %s\n", element.Level, element.School)
        fmt.Printf("Casting Time: %s\n", element.Casting_time)
        fmt.Printf("Range: %s\n", element.Range)
        fmt.Printf("Components: %s\n", element.Components)
        fmt.Printf("Duration: %s\n", element.Duration)
        fmt.Printf("Description: %s\n", element.Desc)
        fmt.Printf("At Higher Levels: %s\n", element.Higher_levels) 
        fmt.Printf("\n")
    }
    fmt.Printf("----------------------------------------\n\n")
    return 0
}
