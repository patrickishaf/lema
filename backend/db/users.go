package db

import (
	"log"

	"github.com/patrickishaf/lema/models"
)

func FindUsers(limit int, offset int) ([]models.User, error) {
	var users []models.User
	err := db.Limit(limit).Offset(offset).Order("id asc, email").Find(&users).Error
	if err != nil {
		log.Printf("failed to find users. error: %v", err)
		return nil, err
	}
	return users, nil
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
		{
			Name:     "John Sunderland",
			Username: "jnsunderland",
			Email:    "john.sunderland@acme.corp",
			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
		},
		{
			Name:     "Heather Grayson",
			Username: "hgrayson",
			Email:    "h.grayson@acme.corp",
			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
		},
		{
			Name:     "Andy Townsend",
			Username: "atownsend",
			Email:    "andy_townsend@acme.corp",
			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
		},
		{
			Name:     "Walter Bullivan",
			Username: "wbullivan",
			Email:    "walter.b@acme.corp",
			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
		},
		{
			Name:     "Jaime Sunderland",
			Username: "jmesunderland",
			Email:    "jme.sunderland@acme.corp",
			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
		},
		{
			Name:     "Weather Mayson",
			Username: "wmayson",
			Email:    "w.mayson@acme.corp",
			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
		},
		{
			Name:     "Fenrir Townsend",
			Username: "ftownsend",
			Email:    "fenrir_townsend@acme.corp",
			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
		},
		{
			Name:     "Walter Boulevard",
			Username: "wboulevard",
			Email:    "walter.blvd@acme.corp",
			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
		},
		{
			Name:     "Zend Southampton",
			Username: "zsouthampton",
			Email:    "z.souhampton@acme.corp",
			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
		},
		{
			Name:     "Leather Mayson",
			Username: "lmayson",
			Email:    "l.mayson@acme.corp",
			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
		},
		{
			Name:     "Henry Brownsend",
			Username: "hbrownsend",
			Email:    "henry_brownsend@acme.corp",
			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
		},
		{
			Name:     "Salter Mullivan",
			Username: "smullivan",
			Email:    "salter.m@acme.corp",
			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
		},
	}

	result := getDB().Create(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
