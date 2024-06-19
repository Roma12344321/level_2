package main

import "fmt"

type Command interface {
	Execute()
}

type DataBase struct{}

func (d *DataBase) Insert() {
	fmt.Println("Inserting record...")
}

func (d *DataBase) Update() {
	fmt.Println("Updating record...")
}

func (d *DataBase) Select() {
	fmt.Println("Reading record...")
}

func (d *DataBase) Delete() {
	fmt.Println("Deleting record...")
}

type InsertCommand struct {
	*DataBase
}

func (c *InsertCommand) Execute() {
	c.DataBase.Insert()
}

type UpdateCommand struct {
	*DataBase
}

func (c *UpdateCommand) Execute() {
	c.DataBase.Update()
}

type SelectCommand struct {
	*DataBase
}

func (c *SelectCommand) Execute() {
	c.DataBase.Select()
}

type DeleteCommand struct {
	*DataBase
}

func (c *DeleteCommand) Execute() {
	c.DataBase.Delete()
}

type Developer struct {
	Insert Command
	Update Command
	Select Command
	Delete Command
}

func (d *Developer) InsertRecord() {
	d.Insert.Execute()
}

func (d *Developer) UpdateRecord() {
	d.Update.Execute()
}

func (d *Developer) SelectRecord() {
	d.Select.Execute()
}

func (d *Developer) DeleteRecord() {
	d.Delete.Execute()
}

func main() {
	db := &DataBase{}
	dev := &Developer{
		Insert: &InsertCommand{db},
		Update: &UpdateCommand{db},
		Select: &SelectCommand{db},
		Delete: &DeleteCommand{db},
	}
	dev.InsertRecord()
	dev.UpdateRecord()
	dev.SelectRecord()
	dev.DeleteRecord()
}
