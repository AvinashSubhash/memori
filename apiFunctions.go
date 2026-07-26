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

func GetTopicDetails(id uint) models.Topic {
	fmt.Println("Fetching Topic Details for Topic:", id)
	var TopicDetails models.Topic
	result := DB.First(&TopicDetails, id)
	if result.Error != nil {
		fmt.Println("Error Fetching details for Topic:", id, result.Error)
		return models.Topic{}
	}
	return TopicDetails
}

func incrementTopic(id uint) {
	fmt.Println("Incrementing Topic: ", id)
	// implement generic topic fetching function
	topicDetails := GetTopicDetails(id)

	newRevisionDate, newInterval := calculateNextRevisionDate(topicDetails.NextRevisionDate, topicDetails.CurrentInterval)
	result := DB.Model(&models.Topic{}).Where("id = ?", id).Updates(models.Topic{CurrentInterval: newInterval, NextRevisionDate: newRevisionDate})
	if result.Error != nil {
		fmt.Println("Error updating data: ", result.Error)
		return
	}
	fmt.Println("Data updated successfully.")
}

func calculateNextInterval(currInterval int) int {
	interval_map := map[int]int{
		1:  4,
		4:  7,
		7:  14,
		14: 28,
	}

	return interval_map[currInterval]
}

func calculateNextRevisionDate(currDate time.Time, currInterval int) (time.Time, int) {
	nextInterval := calculateNextInterval(currInterval)
	nextRevisionDate := currDate.AddDate(0, 0, nextInterval)
	fmt.Println("Next Revision Date: ", nextRevisionDate)
	return nextRevisionDate, nextInterval
}
