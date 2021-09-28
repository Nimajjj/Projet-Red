package main

type Equipment struct {
  Name string
  Slot string
  Effect int
}

var Empty Equipment

var AdventurerHat Equipment
var AdventurerCloak Equipment
var DarkestSword Equipment
var Crown Equipment

var EquipmentList = map[string]Equipment{}


func (e *Equipment) initStuff(name, slot string, effect int) {
  e.Name = name
  e.Slot = slot
  e.Effect = effect
}


func InitEquipments() {
  Empty = Equipment{}
  AdventurerHat.initStuff("Adventurer Hat", "Head", 10)
  AdventurerCloak.initStuff("Adventurer Cloak", "Body", 15)
  DarkestSword.initStuff("Darkest Sword", "Weapon", 0)
  Crown.initStuff("Crown", "Head", 5)

  EquipmentList["Adventurer Hat"] = AdventurerHat
  EquipmentList["Adventurer Cloak"] = AdventurerCloak
  EquipmentList["Darkest Sword"] = DarkestSword
  EquipmentList["Crown"] = Crown
}
