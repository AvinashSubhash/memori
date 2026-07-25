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

func updateTopic(id uint) {
	fmt.Println("Updating topic.")
	result := DB.Model(&models.Topic{}).Where("id = ?", id).Updates(models.Topic{CurrentInterval: 0, NextRevisionDate: time.Now()})
	if result.Error != nil {
		fmt.Println("Error updating data: ", result.Error)
		return
	}
	fmt.Println("Data updated successfully.")
}
