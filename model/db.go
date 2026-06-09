package model

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/test/model"
)

// ObjectBox initialization
func InitBedRock() (*objectbox.ObjectBox, error) {
	objectBox, err := objectbox.NewBuilder().Model(model.ObjectBoxModel()).Build()
	if err != nil {
		fmt.Println("Error initializing objectbox: ", err.Error())
		return nil, err
	}
	return objectBox, nil
}

// Seed objectbox with the Narrative Files. Skips seeding if the Narrative data already exists in objectbox, to avoid duplicates. This will be replaced with a more robust seeding process once objectbox is fully implemented.
func SeedBedRock(narratives string, objectBox *objectbox.ObjectBox) {
	// For each Narrative file, check if the corresponding Narrative data already exists in objectbox. If not, parse the Narrative file and write the data to objectbox.

	filepath.Walk(narratives, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Error walking through narrative directory: ", err.Error())
			return err
		}
		fmt.Println(path, info.Name())
		return nil
	})
}

// func GetObjectBox() (*objectbox.ObjectBox, error) {

// }
//
