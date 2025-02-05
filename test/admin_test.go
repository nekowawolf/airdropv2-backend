package test

import (
	"fmt"
	"testing"

	"github.com/nekowawolf/airdropv2/module"
)

func TestInsertAdmin(t *testing.T) {
	username := "admin"
	password := ""

	result, err := module.InsertAdmin(username, password)
	if err != nil {
		t.Errorf("Failed to admin notes: %v", err)
		return
	}

	fmt.Printf("Inserted Admin ID: %v\n", result)
}

func TestLoginAdmin(t *testing.T) {
	username := "admin"
	password := "admin123"

	success, err := module.LoginAdmin(username, password)
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


