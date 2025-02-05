package test

import (
	"testing"
	"fmt"

	"github.com/nekowawolf/airdropv2/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertAirdropFree(t *testing.T) {
	name := "CARV"
	task := "DAILY"
	link := "https://protocol.carv.io/airdrop"
	level := "easy"
	status := "active"
	backed := "HashKey Capital, ConsenSys"
	funds := "53.37M"
	supply := "N/A"
	marketCap := "N/A"
	vesting := "N/A"
	linkClaim := "N/A"
	price := 0.0
	usdIncome := 0

	hasil, err := module.InsertAirdropFree(name, task, link, level, status, backed, funds, supply, marketCap, vesting, linkClaim, price, usdIncome)
	if err != nil {
		t.Errorf("Failed to insert AirdropFree: %v", err)
	} else {
		fmt.Printf("Inserted AirdropFree ID: %v\n", hasil)
	}
}

func TestInsertAirdropPaid(t *testing.T) {
	name := "Owlto"
	task := "RETRO"
	link := "https://owlto.finance/"
	level := "medium"
	status := "active"
	backed := "Bixin Ventures, GSR"
	funds := "8.00M"
	supply := "N/A"
	marketCap := "N/A"
	vesting := "N/A"
	linkClaim := "N/A"
	price := 0.0
	usdIncome := 0

	hasil, err := module.InsertAirdropPaid(name, task, link, level, status, backed, funds, supply, marketCap, vesting, linkClaim, price, usdIncome)
	if err != nil {
		t.Errorf("Failed to insert AirdropPaid: %v", err)
	} else {
		fmt.Printf("Inserted AirdropPaid ID: %v\n", hasil)
	}
}

func TestGetAllAirdropFree(t *testing.T) {
	data, err := module.GetAllAirdropFree()
	if err != nil {
		t.Errorf("Failed to retrieve AirdropFree data: %v", err)
	} else if len(data) == 0 {
		t.Errorf("No AirdropFree data found")
	} else {
		fmt.Printf("Retrieved AirdropFree data: %v\n", data)
	}
}

func TestGetAllAirdrop(t *testing.T) {
	allAirdrops, err := module.GetAllAirdrop()
	if err != nil {
		t.Errorf("Failed to retrieve all airdrops: %v", err)
		return
	}

	if len(allAirdrops) == 0 {
		t.Errorf("No airdrops found.")
		return
	}

	fmt.Printf("Retrieved %d airdrops:\n", len(allAirdrops))
	for _, airdrop := range allAirdrops {
		fmt.Printf("%v\n", airdrop)
	}

	freeAirdrops, err := module.GetAllAirdropFree()
	if err != nil {
		t.Errorf("Failed to retrieve AirdropFree data: %v", err)
		return
	}

	paidAirdrops, err := module.GetAllAirdropPaid()
	if err != nil {
		t.Errorf("Failed to retrieve AirdropPaid data: %v", err)
		return
	}

	expectedCount := len(freeAirdrops) + len(paidAirdrops)
	if len(allAirdrops) != expectedCount {
		t.Errorf("Expected %d airdrops, but got %d", expectedCount, len(allAirdrops))
	}
}

func TestGetAllAirdropPaid(t *testing.T) {
	data, err := module.GetAllAirdropPaid()
	if err != nil {
		t.Errorf("Failed to retrieve AirdropPaid data: %v", err)
	} else if len(data) == 0 {
		t.Errorf("No AirdropPaid data found")
	} else {
		fmt.Printf("Retrieved AirdropPaid data: %v\n", data)
	}
}

func TestGetAirdropFreeByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66cfa14dd39e7e3b0c85b295")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	airdrop, err := module.GetAirdropFreeByID(id)
	if err != nil {
		t.Errorf("Failed to retrieve AirdropFree by ID: %v", err)
	} else {
		fmt.Printf("Retrieved AirdropFree by ID: %v\n", airdrop)
	}
}

func TestGetAirdropPaidByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66cfa14dd39e7e3b0c85b297")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	airdrop, err := module.GetAirdropPaidByID(id)
	if err != nil {
		t.Errorf("Failed to retrieve AirdropPaid by ID: %v", err)
	} else {
		fmt.Printf("Retrieved AirdropPaid by ID: %v\n", airdrop)
	}
}

func TestGetAirdropFreeByName(t *testing.T) {
	name := "Initia"
	airdrop, err := module.GetAirdropFreeByName(name)
	if err != nil {
		t.Fatalf("Error calling GetAirdropFreeByName: %v", err)
	}

	fmt.Println("Free Airdrops found:", airdrop)
}

func TestGetAirdropPaidByName(t *testing.T) {
	name := "Dflow"
	airdrop, err := module.GetAirdropPaidByName(name)
	if err != nil {
		t.Fatalf("Error calling GetAirdropPaidByName: %v", err)
	}

	fmt.Println("Paid Airdrops found:", airdrop)
}

func TestGetAllAirdropByName(t *testing.T) {
	name := "Sender" 

	airdrops, err := module.GetAllAirdropByName(name)
	if err != nil {
		t.Fatalf("Error calling GetAirdropByName: %v", err)
	}

	if len(airdrops) == 0 {
		t.Errorf("No airdrops found with name '%s'", name)
	} else {
		fmt.Println("Airdrops found:", airdrops)
	}
}

func TestUpdateAirdropFreeByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66f68eb6d7123a87a5065b63")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	newName := "CARV"
	newTask := "DAILY"
	newLink := "https://protocol.carv.io/airdrop"
	newLevel := "easy"
	newStatus := "ended"
	newBacked := "HashKey Capital, ConsenSys"
	newFunds := "53.37M"
	newSupply := "100M" 
	newMarketCap := "632.23M"
	newVesting := "no"
	newLinkClaim := "https://protocol.carv.io/claim"
	newPrice := 0.2
	newUSDIncome := 1200

	err = module.UpdateAirdropFreeByID(id, newName, newTask, newLink, newLevel, newStatus, newBacked, newFunds, newSupply, newMarketCap, newVesting, newLinkClaim, newPrice, newUSDIncome)
	if err != nil {
		t.Errorf("Failed to update AirdropFree by ID: %v", err)
		return
	}

	airdrop, err := module.GetAirdropFreeByID(id)
	if err != nil {
		t.Errorf("Failed to retrieve AirdropFree by ID after update: %v", err)
		return
	}

	if airdrop.Name != newName || airdrop.Task != newTask || airdrop.Link != newLink || airdrop.Level != newLevel ||
		airdrop.Status != newStatus || airdrop.Backed != newBacked || airdrop.Funds != newFunds || airdrop.Supply != newSupply ||
		airdrop.MarketCap != newMarketCap || airdrop.Vesting != newVesting || airdrop.LinkClaim != newLinkClaim ||
		airdrop.Price != newPrice || airdrop.USDIncome != newUSDIncome {
		t.Errorf("AirdropFree not updated correctly. Got: %+v", airdrop)
	} else {
		fmt.Printf("AirdropFree updated successfully: %+v\n", airdrop)
	}
}

func TestUpdateAirdropPaidByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66f6905e14b6a0632a67bffd")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	newName := "Owlto"
	newTask := "RETRO"
	newLink := "https://owlto.finance"
	newLevel := "medium"
	newStatus := "ended"
	newBacked := "Bixin Ventures, GSR"
	newFunds := "8.00M"
	newSupply := "50M" 
	newMarketCap := "788.75M"
	newVesting := "yes"
	newLinkClaim := "https://owlto.finance/airdrop"
	newPrice := 0.5
	newUSDIncome := 230

	err = module.UpdateAirdropPaidByID(id, newName, newTask, newLink, newLevel, newStatus, newBacked, newFunds, newSupply, newMarketCap, newVesting, newLinkClaim, newPrice, newUSDIncome)
	if err != nil {
		t.Errorf("Failed to update AirdropPaid by ID: %v", err)
		return
	}

	airdrop, err := module.GetAirdropPaidByID(id)
	if err != nil {
		t.Errorf("Failed to retrieve AirdropPaid by ID after update: %v", err)
		return
	}

	if airdrop.Name != newName || airdrop.Task != newTask || airdrop.Link != newLink || airdrop.Level != newLevel ||
		airdrop.Status != newStatus || airdrop.Backed != newBacked || airdrop.Funds != newFunds || airdrop.Supply != newSupply ||
		airdrop.MarketCap != newMarketCap || airdrop.Vesting != newVesting || airdrop.LinkClaim != newLinkClaim ||
		airdrop.Price != newPrice || airdrop.USDIncome != newUSDIncome {
		t.Errorf("AirdropPaid not updated correctly. Got: %+v", airdrop)
	} else {
		fmt.Printf("AirdropPaid updated successfully: %+v\n", airdrop)
	}
}

func TestDeleteAirdropFreeByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66ea30e2017690fa6e447744")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	err = module.DeleteAirdropFreeByID(id)
	if err != nil {
		t.Errorf("Failed to delete AirdropFree by ID: %v", err)
		return
	}

	airdrop, err := module.GetAirdropFreeByID(id)
	if err == nil {
		t.Errorf("Expected no document, but found AirdropFree: %+v", airdrop)
	} else {
		fmt.Printf("AirdropFree deleted successfully, no document found with ID: %s\n", id.Hex())
	}
}

func TestDeleteAirdropPaidByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("66f68c6ca4684f29ed12c4ec")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	err = module.DeleteAirdropPaidByID(id)
	if err != nil {
		t.Errorf("Failed to delete AirdropPaid by ID: %v", err)
		return
	}

	airdrop, err := module.GetAirdropPaidByID(id)
	if err == nil {
		t.Errorf("Expected no document, but found AirdropPaid: %+v", airdrop)
	} else {
		fmt.Printf("AirdropPaid deleted successfully, no document found with ID: %s\n", id.Hex())
	}
}