package main

import (
  "fmt"
)

type Item struct {
  Description string
  FLink func()  // activated func on use
  Usable bool
}


func (i *Item) init(item_name, desc string, flink func(), usable bool) {
  i.Description = desc
  i.FLink = flink
  i.Usable = usable
  ItemsMap[item_name] = *i
}

func (c *Character) addItem(item_name string, quantity int) {
  if _, has_item := c.Inventory[item_name]; !has_item {
    c.Inventory[item_name] += quantity
  } else {
    c.Inventory[item_name] = quantity
  }
}

func SelectItem(itemName interface{}) {
  item_name := itemName.(string)
  fmt.Print(item_name, ":\n")
  fmt.Print(ItemsMap[item_name].Description, "\n")
  choice_str := "0 - Back | "
  if ItemsMap[item_name].Usable {
    choice_str += "1 - Use"
  }
  fmt.Print(choice_str, "\n")
  input := TakeIntInput()
  if !ItemsMap[item_name].Usable {
    input = 0
  }
  switch intut {
  case 0:
    Player.displayInventory()
  case 1:
    ItemsMap[item_name].FLink()
  default:
    Player.displayInventory()
  }
}
