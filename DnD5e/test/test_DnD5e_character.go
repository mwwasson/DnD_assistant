package main

import (
    "fmt"
    "strings"
    "github.com/mwwasson/DnD_assistant/DnD5e"
)

func main() {
	var character DnD5e.Character
	character.Initialize()
	character.SetProficiencyBonus(2)
	
	character.SetAttributeAll ("str", 14, 2, 0, true)
	character.SetAttributeAll ("dex", 12, 0, 1, false)
	character.SetAttributeAll ("con", 10, 1, 0, true)
	character.SetAttributeAll ("int", 9, 0, 0, false)
	character.SetAttributeAll ("wis", 11, 0, 0, false)
	character.SetAttributeAll ("cha", 18, 0, 0, false)
	
	attributes := character.GetAttributeList()
	
	
	fmt.Printf("Proficiency Bonus: %d\n", character.GetProficiencyBonus())
	for _,k := range attributes{
		attr, mod, _ := character.GetAttribute(k)
		base,race,misc,_,_ := character.GetAttributeAll(k)
		fmt.Printf("%s: %d (%d): %d + %d + %d\n", strings.Title(k), attr, mod, base, race, misc)
	}
	
	fmt.Printf("Saving Throws\n")
	for _,k := range attributes{
		sav,_ := character.GetSavingThrow (k)
		fmt.Printf("%s: %d\n", strings.Title(k), sav)
	}
	/* Skills */
	
	fmt.Printf("Skills\n")
	
	keys := character.GetSkillList()
	for _,k := range keys{
		fmt.Printf("%s\n", strings.Title(k))
	}
}
