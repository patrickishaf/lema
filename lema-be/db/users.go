package db

import "github.com/patrickishaf/lema-be/models"

func FindUsers(limit int, offset int) []models.User {
	var users []models.User
	db.Find(&users).Limit(limit).Offset(offset)
	return users
}

func FindUserCount() int64 {
	var count int64
	getDB().Model(&models.User{}).Count(&count)
	return count
}

func FindUserById(id uint) models.User {
	var user models.User
	db.Where(&models.User{ID: id}).First(&user)
	return user
}

func insertDummyUsers() error {
	users := []models.User{
		{
			Name:     "James Sunderland",
			Username: "jsunderland",
			Email:    "james.sunderland@acme.corp",
			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
		},
		{
			Name:     "Heather Mayson",
			Username: "hmayson",
			Email:    "h.mayson@acme.corp",
			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
		},
		{
			Name:     "Henry Townsend",
			Username: "htownsend",
			Email:    "henry_townsend@acme.corp",
			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
		},
		{
			Name:     "Walter Sullivan",
			Username: "wsullivan",
			Email:    "walter.s@acme.corp",
			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
		},
	}

	result := getDB().Create(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
