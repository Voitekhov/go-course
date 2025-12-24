package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func newCompany(name string) *ManagementCompany {
	return &ManagementCompany{
		Name:            name,
		HousesByAddress: make(map[string]*House),
	}
}

func newHouse(address string, debt int64) *House {
	return &House{
		Address: address,
		Debt:    debt,
	}
}

func TestConnectHouse(t *testing.T) {
	mc := newCompany("MC-1")
	house := newHouse("Pushkin st 10", 0)

	err := mc.ConnectHouse(house)

	require.NoError(t, err)
	require.True(t, house.IsConnected)
	require.Equal(t, mc, house.ManagedBy)
	require.Contains(t, mc.HousesByAddress, "Pushkin st 10")
}

func TestConnectHouse_AlreadyConnected(t *testing.T) {
	mc := newCompany("MC-1")
	house := newHouse("Pushkin st 10", 0)

	require.NoError(t, mc.ConnectHouse(house))

	err := mc.ConnectHouse(house)
	require.Error(t, err)
}

func TestDisconnectHouse(t *testing.T) {
	mc := newCompany("MC-1")
	house := newHouse("Pushkin st 10", 0)

	require.NoError(t, mc.ConnectHouse(house))

	err := mc.DisconnectHouse(house, false)

	require.NoError(t, err)
	require.False(t, house.IsConnected)
	require.Nil(t, house.ManagedBy)
	require.NotContains(t, mc.HousesByAddress, house.Address)
}

func TestDisconnectHouse_FailedByDebt(t *testing.T) {
	mc := newCompany("MC-1")
	house := newHouse("Pushkin st 10", 50000)

	require.NoError(t, mc.ConnectHouse(house))

	err := mc.DisconnectHouse(house, false)
	require.Error(t, err)
	require.True(t, house.IsConnected)
}

func TestDisconnectHouse_ForcedByDebt(t *testing.T) {
	mc := newCompany("MC-1")
	house := newHouse("Pushkin st 10", 150000)

	require.NoError(t, mc.ConnectHouse(house))
	err := mc.DisconnectHouse(house, true)
	require.NoError(t, err)
	require.False(t, house.IsConnected)
	require.Nil(t, house.ManagedBy)
}

func TestHandleProblem_work(t *testing.T) {
	mc := newCompany("MC-1")
	house := newHouse("Pushkin st 10", 0)

	require.NoError(t, mc.ConnectHouse(house))

	err := mc.HandleProblem(house)
	require.NoError(t, err)
}

func TestHandleProblem_AnotherCompany(t *testing.T) {
	mc1 := newCompany("MC-1")
	mc2 := newCompany("MC-2")
	house := newHouse("Pushkin st 10", 0)

	require.NoError(t, mc1.ConnectHouse(house))

	err := mc2.HandleProblem(house)
	require.Error(t, err)
}

func TestChangeCompany_work(t *testing.T) {
	mc1 := newCompany("MC-1")
	mc2 := newCompany("MC-2")
	house := newHouse("Pushkin st 10", 0)

	require.NoError(t, mc1.ConnectHouse(house))

	err := house.ChangeCompany(mc2)

	require.NoError(t, err)
	require.Equal(t, mc2, house.ManagedBy)
	require.Contains(t, mc2.HousesByAddress, house.Address)
	require.NotContains(t, mc1.HousesByAddress, house.Address)
}

func TestChangeCompany_FailedByDebt(t *testing.T) {
	mc1 := newCompany("MC-1")
	mc2 := newCompany("MC-2")
	house := newHouse("Pushkin st 10", 30000)

	require.NoError(t, mc1.ConnectHouse(house))

	err := house.ChangeCompany(mc2)

	require.Error(t, err)
	require.Equal(t, mc1, house.ManagedBy)
	require.Contains(t, mc1.HousesByAddress, house.Address)
	require.NotContains(t, mc2.HousesByAddress, house.Address)
}
