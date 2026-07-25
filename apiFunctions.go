package main

import (
	"fmt"
	"memori/server/models"
	"time"
)

func insertTopic(name string, description string) {
	fmt.Println("Inserting new topic into database.")
	insertData := models.Topic{Name: name, Description: description, NextRevisionDate: time.Now(), CurrentInterval: 1}
	result := DB.Create(&insertData)
	if result.Error != nil {
		fmt.Println("Error creating topic:", result.Error)
		return
	}
	fmt.Println("New topic inserted.")
}
