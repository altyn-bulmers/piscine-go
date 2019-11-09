package main

import (
 "fmt"
)

type person struct {
  age       int
  weight    int
  firstName string
  next      *person
}

func main() {
  mike := &person{33, 200, "mike", nil}
  personList := mike  

  greg := &person{34, 443, "greg", nil}
  james := &person{45, 554, "james", nil}
  marshall := &person{56, 123, "marshall", nil}  
  

  personList = addNodeEnd(greg, personList)
  personList = addNodeEnd(james, personList)
  personList = addNodeEnd(marshall, personList)
  printList(personList)

}

func printList(personList *person) {
  for p := personList; p != nil; p = p.next {
   fmt.Println(p)
  }
}

func addNodeEnd(newPerson, personList *person) *person {
  if personList == nil {
    return personList
  }
  for p := personList; p != nil; p = p.next {
    if p.next == nil {
      p.next = newPerson
      return personList
    }
   }
  return personList
}