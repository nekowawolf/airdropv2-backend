package test

import (
	"fmt"
	"testing"

	"github.com/nekowawolf/airdropv2/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInsertNotes(t *testing.T) {
	title := "Test"
	content := "unit testing."

	result, err := module.InsertNotes(title, content)
	if err != nil {
		t.Errorf("Failed to insert notes: %v", err)
		return
	}

	fmt.Printf("Inserted Notes ID: %v\n", result)
}

func TestGetAllNotes(t *testing.T) {
	data, err := module.GetAllNotes()
	if err != nil {
		t.Errorf("Failed to retrieve notes: %v", err)
	} else if len(data) == 0 {
		t.Errorf("No notes found")
	} else {
		fmt.Printf("Retrieved notes: %v\n", data)
	}
}

func TestGetNotesByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("674253f0e2a1fdaf1c25c333") 
	if err != nil {
		t.Fatalf("Invalid ObjectID: %v", err)
	}

	notes, err := module.GetNotesByID(id)
	if err != nil {
		t.Errorf("Failed to get notes: %v", err)
	} else {
		t.Logf("Retrieved Notes: %+v", notes)
	}
}

func TestUpdateNotesByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("674253f0e2a1fdaf1c25c333")
	if err != nil {
		t.Fatalf("Invalid ObjectID: %v", err)
	}

	newTitle := "Test1"
	newContent := "Testing"

	updatedNote, err := module.UpdateNotesByID(id, newTitle, newContent)
	if err != nil {
		t.Errorf("Failed to update note: %v", err)
		return
	}

	t.Logf("Successfully updated note: %+v", updatedNote)
}

func TestDeleteNotesByID(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("67439328c0edc2854b3a1828")
	if err != nil {
		t.Errorf("Invalid ID format: %v", err)
		return
	}

	err = module.DeleteNotesByID(id)
	if err != nil {
		t.Errorf("Failed to delete note by ID: %v", err)
		return
	}

	t.Logf("Note with ID %s deleted successfully", id.Hex())
}