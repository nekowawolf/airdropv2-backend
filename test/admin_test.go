package test

import (
	"github.com/nekowawolf/airdropv2/features/admin"
	"fmt"
	"testing"

)

func TestInsertAdmin(t *testing.T) {
	username := "admin"
	password := ""

	result, err := admin.InsertAdmin(username, password)
	if err != nil {
		t.Errorf("Failed to admin notes: %v", err)
		return
	}

	fmt.Printf("Inserted Admin ID: %v\n", result)
}

func TestLoginAdmin(t *testing.T) {
	username := "admin"
	password := "admin123"

	success, err := admin.LoginAdmin(username, password)
	if err != nil {
		t.Errorf("Login failed: %v", err)
		return
	}

	if !success {
		t.Errorf("Login should be successful with correct credentials")
		return
	}

	fmt.Println("Login successful")
}


