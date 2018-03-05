package DnD5e

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type Rule_instance struct {
    Title   string    `json:"title"`
    Desc    string    `json:"description"`
}

type Spell_instance struct {
    Title           string  `json:"title"`
    Level           int     `json:"level"`
    School          string  `json:"school"`
    Casting_time    string  `json:"casting_time"`
    Range           string  `json:"range"`
    Components      string  `json:"components"`
    Duration        string  `json:"duration"`
    Desc            string  `json:"description"`
    Higher_levels   string  `json:"higher_levels"`
}

type Armor_instance struct {
    Type        string  `json:"type"`
    Name        string  `json:"name"`
    Cost        float32  `json:"cost_gp"`
    AC          string  `json:"ac"`
    Strength    int     `json:"str"`
    Stealth     string  `json:"stealth"`
    Weight      float32 `json:"weight"` 

}

type Weapon_instance struct {
    Type        string  `json:"type"`
    Name        string  `json:"name"`
    Cost        float32 `json:"cost_gp"`
    Damage      string  `json:"damage"`
    Weight      float32 `json:"weight"`
    Prop        string  `json:"properties"`
    Desc        string  `json:"description"`
     

}

type Gear_instance struct {
    Item    string  `json:"item"`
    Cost    float32 `json:"cost_gp"`
    Weight  float32 `json:"weight"`
    Desc    string  `json:"description"` 

}

type Data struct {
    Rule    []Rule_instance     `json:"rule"`   
    Spell   []Spell_instance    `json:"spell"`
    Armor   []Armor_instance    `json:"armor"`
    Weapon  []Weapon_instance   `json:"weapon"`
    Gear    []Gear_instance     `json:"adventuring_gear"`
}

func (container *Data) Load_json_file(file string) int {
    raw, err := ioutil.ReadFile(file)
    
    if err != nil {
        fmt.Println(err.Error())
        return -1
    }

    json.Unmarshal(raw, &container)
    return 0
}
