package main

import (
	"errors"
	"fmt"
)

type House struct {
	Address     string
	Debt        int64
	ManagedBy   *ManagementCompany
	IsConnected bool
}

type ManagementCompany struct {
	Name   string
	Houses map[string]*House
}

const MaxDebt = 100000

func (mc *ManagementCompany) ConnectHouse(house *House) error {
	if house.ManagedBy != nil {
		return errors.New("house already managed")
	}

	if mc.Houses == nil {
		mc.Houses = make(map[string]*House)
	}

	house.ManagedBy = mc
	house.IsConnected = true
	mc.Houses[house.Address] = house
	return nil
}

func (mc *ManagementCompany) DisconnectHouse(house *House, forced bool) error {
	if house.ManagedBy != mc {
		return errors.New("house is not managed by this company")
	}

	if house.Debt > 0 && !forced {
		return errors.New("cannot disconnect with debt")
	}

	delete(mc.Houses, house.Address)
	house.ManagedBy = nil
	house.IsConnected = false
	return nil
}

func (mc *ManagementCompany) SendBills() {
	for _, h := range mc.Houses {
		fmt.Printf("Sending bills to %s\n", h.Address)
	}
}

func (mc *ManagementCompany) HandleProblem(house *House) error {
	if !house.IsConnected || house.ManagedBy != mc {
		return errors.New("not connected to management company or serviced other company")
	}
	fmt.Println("Your request has been received. We are working on a solution.")
	return nil
}

func (house *House) ChangeCompany(newCompany *ManagementCompany) error {
	if house.Debt > 0 {
		return fmt.Errorf("cannot change company, debt: %d", house.Debt)
	}

	if house.ManagedBy != nil {
		delete(house.ManagedBy.Houses, house.Address)
	}

	house.ManagedBy = newCompany
	house.IsConnected = true
	if newCompany.Houses == nil {
		newCompany.Houses = make(map[string]*House)
	}
	newCompany.Houses[house.Address] = house
	return nil
}
