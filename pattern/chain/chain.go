package main

import (
	"fmt"
)

type Service interface {
	Execute(*Data)
	SetNext(Service)
}

type Data struct {
	GetSource    bool
	UpdateSource bool
}

type Device struct {
	Name string
	Next Service
}

func (d *Device) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("Data from device [%v] has already been gotten\n", d.Name)
		d.Next.Execute(data)
		return
	}
	fmt.Printf("Data from device [%v] havs been gotten\n", d.Name)
	data.GetSource = true
	d.Next.Execute(data)
}

func (d *Device) SetNext(service Service) {
	d.Next = service
}

type UpdateDataService struct {
	Name string
	Next Service
}

func (d *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("Data from device [%v] has already been updated\n", d.Name)
		d.Next.Execute(data)
		return
	}
	fmt.Printf("Data from device [%v] has been updated\n", d.Name)
	data.GetSource = true
	d.Next.Execute(data)
}

func (d *UpdateDataService) SetNext(service Service) {
	d.Next = service
}

type DataService struct {
	Next Service
}

func (d *DataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Println("Data has not been updated\n")
		return
	}
	fmt.Println("Data has been saved")
}

func (d *DataService) SetNext(service Service) {
	d.Next = service
}

func main() {
	device := &Device{Name: "Device 1"}
	updateService := &UpdateDataService{Name: "Update 1"}
	dataService := &DataService{}
	device.SetNext(updateService)
	updateService.SetNext(dataService)
	data := &Data{}
	device.Execute(data)
}
