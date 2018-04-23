package DnD5e

import (
	"math"
    "strings"
    "sort"
)


type Character struct {
	proficiency int
    attribute  	attribute_m
    skill   	map[string]*skillInstance_s
}

func (a *Character) Initialize() {
	a.proficiency = 0
	a.attribute.attr = make(map[string]*attribute_s)
	a.initializeAttributes()
	a.skill = make(map[string]*skillInstance_s)
	a.initializeSkills()
}

func (a Character) GetProficiencyBonus () int {
	return a.proficiency
}

func (a *Character) SetProficiencyBonus (value int) {
	a.proficiency = value
}

/* Attributes */

/* Public Attributes */

func (a Character) GetAttribute (attr string) (int, int, bool){

	base,race,misc,_,valid := a.attribute.get(strings.ToLower(attr))
	
	sum := base + race + misc
	mod := int(math.Floor(float64(sum/2)) - 5)

    return sum, mod, valid
}

func (a Character) GetAttributeBase(attr string) (int, bool) {
	base,_,_,_,valid := a.attribute.get(strings.ToLower(attr))
	return base, valid
}
func (a Character) GetAttributeRacial(attr string) (int, bool) {
	_,race,_,_,valid := a.attribute.get(strings.ToLower(attr))
	return race, valid
}
func (a Character) GetAttributeMisc(attr string) (int, bool) {
	_,_,misc,_,valid := a.attribute.get(strings.ToLower(attr))
	return misc, valid
}
func (a Character) GetAttributeProf(attr string) (bool, bool) {
	_,_,_,prof,valid := a.attribute.get(strings.ToLower(attr))
	return prof, valid
}
func (a Character) GetAttributeAll(attr string) (int, int, int, bool, bool){
	return a.attribute.get(strings.ToLower(attr))
}

func (a Character) GetAttributeList() []string{
	var keys []string
	for key,_ := range a.attribute.attr{
		keys = append(keys, key)
	}
	//sort.Strings(keys)
	return keys
}

func (a Character) GetSavingThrow (attr string) (int, bool) {
	_,_,_,prof,valid := a.attribute.get(strings.ToLower(attr))
	if valid {
		_,save,_ := a.GetAttribute(attr)
		if prof {
			save += a.proficiency
		}
		return save, true
	} else {
		return 0,false
	}
}

func (a *Character) SetAttributeBase (attr string, base int) bool {
	_,race,misc,prof,valid := a.attribute.get(strings.ToLower(attr))
	if valid {
		return a.attribute.set (attr, base, race, misc,prof)
	} else {
		return false
	}
}
func (a *Character) SetAttributeRacial(attr string, race int) bool {
	base, _, misc,prof,valid := a.attribute.get(strings.ToLower(attr))
	if valid {
		return a.attribute.set (attr, base, race, misc,prof)
	} else {
		return false
	}
}
func (a *Character) SetAttributeMisc(attr string, misc int) bool {
	base, race, _,prof,valid := a.attribute.get(strings.ToLower(attr))
	if valid {
		return a.attribute.set (attr, base, race, misc,prof)
	} else {
		return false
	}
}
func (a *Character) SetAttributeProf (attr string, prof bool) bool {
	base, race,misc,_,valid := a.attribute.get(strings.ToLower(attr))
	if valid {
		return a.attribute.set (attr, base, race, misc,prof)
	} else {
		return false
	}
}
func (a *Character) SetAttributeAll (attr string, base int, race int, misc int, prof bool) (bool){
	return a.attribute.set (attr, base, race, misc,prof)
}

func (a *Character) ClearAttribute(attr string) bool {
	return a.attribute.set (attr, 10, 0, 0,false)
}
/* Private  Attributes*/
type attribute_m struct {
	attr map[string]*attribute_s
}
type attribute_s struct {
    base        int
    race_adj    int
    misc_adj    int
    prof 		bool
}

func (a *attribute_m) add (attr_name string) bool {
	if _, exists := a.attr[attr_name]; exists {
		return false
	} else {
		a.attr[attr_name] = &attribute_s{0,0,0,false}
		return true
	}
}
func (a attribute_m) get (attr_name string) (int, int, int, bool, bool){
	if _, exists := a.attr[attr_name]; exists {
		base := a.attr[attr_name].base
		race_adj := a.attr[attr_name].race_adj
		misc_adj := a.attr[attr_name].misc_adj
		proficient := a.attr[attr_name].prof
		return base, race_adj, misc_adj, proficient, true
	} else {
		return 0,0,0,false,false
	}
}
func (a attribute_m) set (attr string, base int, race int, misc int, proficient bool) (bool){
	if _, exists := a.attr[attr]; exists {
		a.attr[attr].base = base
		a.attr[attr].race_adj = race
		a.attr[attr].misc_adj = misc
		a.attr[attr].prof = proficient
		return true
	} else {
		return false
	}
}

func (a *Character) initializeAttributes(){
	a.attribute.add("str")
	a.SetAttributeAll("str", 10, 0, 0, false)
	
	a.attribute.add("dex")
	a.SetAttributeAll("dex", 10, 0, 0, false)
	
	a.attribute.add("con")
	a.SetAttributeAll("con", 10, 0, 0, false)
	
	a.attribute.add("int")
	a.SetAttributeAll("int", 10, 0, 0, false)
	
	a.attribute.add("wis")
	a.SetAttributeAll("wis", 10, 0, 0, false)
	
	a.attribute.add("cha")
	a.SetAttributeAll("cha", 10, 0, 0, false)
}


/* Skills */
/* Public Skills */
func (a Character) GetSkillList() []string{
	
	var keys []string
	for key,_ := range a.skill{
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
func (a *Character) SetSkillTrained (skill string, trained bool) bool {
	if _, exists := a.skill[skill]; exists {
		a.skill[skill].trained = trained
		return true
	} else {
		return false
	}
}
func (a *Character) SetSkillMultiplier (skill string, multi int) bool {
	if _, exists := a.skill[skill]; exists {
		a.skill[skill].multiplier = multi
		return true
	} else {
		return false
	}
}

/* Private Skills */
type skillInstance_s struct {
	ability string
	trained	bool
	multiplier int /* multiplier of the proficiency bonus */
}
func (a *Character) initializeSkills(){
	
	a.setSkill("acrobatics", "dex", false, 1)
	a.setSkill("animal handling", "wis", false, 1)
	a.setSkill("arcana", "int", false, 1)
	a.setSkill("athletics", "str", false, 1)
	a.setSkill("deception", "cha", false, 1)
	a.setSkill("history", "int", false, 1)
	a.setSkill("insight", "dex", false, 1)
	a.setSkill("intimidation", "cha", false, 1)
	a.setSkill("investigation", "int", false, 1)
	a.setSkill("medicine", "wis", false, 1)
	a.setSkill("nature", "int", false, 1)
	a.setSkill("perception", "wis", false, 1)
	a.setSkill("performance", "cha", false, 1)
	a.setSkill("persuasion", "cha", false, 1)
	a.setSkill("religion", "int", false, 1)
	a.setSkill("sleight of hand", "dex", false, 1)
	a.setSkill("stealth", "dex", false, 1)
	a.setSkill("survival", "wis", false, 1)
	
}
func (a *Character) setSkill(s string, abil string, train bool, mult int){
	if _, exists := a.skill[s]; exists {
		a.skill[s].ability = abil
		a.skill[s].trained = train
		a.skill[s].multiplier = mult
	} else {
		a.skill[s] = &skillInstance_s{abil, train, mult}
	}
}





type Armor_class struct {
}

type Initiative struct {
}

type Speed struct {
}

type Hit_points struct {
}

type Hit_dice struct {
}

type Class struct {
}

type Death_save struct {
}
