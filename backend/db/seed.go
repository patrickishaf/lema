package db

// func insertSeedUsers() error {
// 	users := []models.User{
// 		{
// 			Name:     "James Sunderland",
// 			Username: "jsunderland",
// 			Email:    "james.sunderland@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Heather Mayson",
// 			Username: "hmayson",
// 			Email:    "h.mayson@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Henry Townsend",
// 			Username: "htownsend",
// 			Email:    "henry_townsend@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Sullivan",
// 			Username: "wsullivan",
// 			Email:    "walter.s@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "John Sunderland",
// 			Username: "jnsunderland",
// 			Email:    "john.sunderland@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Heather Grayson",
// 			Username: "hgrayson",
// 			Email:    "h.grayson@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Andy Townsend",
// 			Username: "atownsend",
// 			Email:    "andy_townsend@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Bullivan",
// 			Username: "wbullivan",
// 			Email:    "walter.b@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "Jaime Sunderland",
// 			Username: "jmesunderland",
// 			Email:    "jme.sunderland@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Weather Mayson",
// 			Username: "wmayson",
// 			Email:    "w.mayson@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Fenrir Townsend",
// 			Username: "ftownsend",
// 			Email:    "fenrir_townsend@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Boulevard",
// 			Username: "wboulevard",
// 			Email:    "walter.blvd@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "Zend Southampton",
// 			Username: "zsouthampton",
// 			Email:    "z.souhampton@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Leather Mayson",
// 			Username: "lmayson",
// 			Email:    "l.mayson@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Henry Brownsend",
// 			Username: "hbrownsend",
// 			Email:    "henry_brownsend@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Salter Mullivan",
// 			Username: "smullivan1",
// 			Email:    "salter.m1@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "James Sunderland",
// 			Username: "jsunderland1",
// 			Email:    "james.sunderland1@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Heather Mayson",
// 			Username: "hmayson1",
// 			Email:    "h.mayson1@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Henry Townsend",
// 			Username: "htownsend1",
// 			Email:    "henry_townsend1@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Sullivan",
// 			Username: "wsullivan1",
// 			Email:    "walter.s1@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "John Sunderland",
// 			Username: "jnsunderland1",
// 			Email:    "john.sunderland1@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Heather Grayson",
// 			Username: "hgrayson1",
// 			Email:    "h.grayson1@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Andy Townsend",
// 			Username: "atownsend1",
// 			Email:    "andy_townsend1@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Bullivan",
// 			Username: "wbullivan1",
// 			Email:    "walter.b1@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "Jaime Sunderland",
// 			Username: "jmesunderland1",
// 			Email:    "jme.sunderland1@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Weather Mayson",
// 			Username: "wmayson1",
// 			Email:    "w.mayson1@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Fenrir Townsend",
// 			Username: "ftownsend1",
// 			Email:    "fenrir_townsend1@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Boulevard",
// 			Username: "wboulevard1",
// 			Email:    "walter.blvd@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "Zend Southampton",
// 			Username: "zsouthampton1",
// 			Email:    "z.souhampton1@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Leather Mayson",
// 			Username: "lmayson1",
// 			Email:    "l.mayson1@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Henry Brownsend",
// 			Username: "hbrownsend1",
// 			Email:    "henry_brownsend1@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Salter Mullivan",
// 			Username: "smullivan1",
// 			Email:    "salter.m1@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "James Sunderland",
// 			Username: "jsunderland2",
// 			Email:    "james.sunderland2@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Heather Mayson",
// 			Username: "hmayson2",
// 			Email:    "h.mayson2@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Henry Townsend",
// 			Username: "htownsend2",
// 			Email:    "henry_townsend2@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Sullivan",
// 			Username: "wsullivan2",
// 			Email:    "walter.s2@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "John Sunderland",
// 			Username: "jnsunderland2",
// 			Email:    "john.sunderland2@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Heather Grayson",
// 			Username: "hgrayson2",
// 			Email:    "h.grayson2@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Andy Townsend",
// 			Username: "atownsend2",
// 			Email:    "andy_townsend2@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Bullivan",
// 			Username: "wbullivan2",
// 			Email:    "walter.b2@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "Jaime Sunderland",
// 			Username: "jmesunderland2",
// 			Email:    "jme.sunderland2@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Weather Mayson",
// 			Username: "wmayson2",
// 			Email:    "w.mayson2@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Fenrir Townsend",
// 			Username: "ftownsend2",
// 			Email:    "fenrir_townsend2@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Boulevard",
// 			Username: "wboulevard2",
// 			Email:    "walter.blvd2@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "Zend Southampton",
// 			Username: "zsouthampton2",
// 			Email:    "z.souhampton2@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Leather Mayson",
// 			Username: "lmayson2",
// 			Email:    "l.mayson2@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Henry Brownsend",
// 			Username: "hbrownsend2",
// 			Email:    "henry_brownsend2@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Salter Mullivan",
// 			Username: "smullivan2",
// 			Email:    "salter.m2@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "James Sunderland",
// 			Username: "jsunderland",
// 			Email:    "james.sunderland@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Heather Mayson",
// 			Username: "hmayson",
// 			Email:    "h.mayson@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Henry Townsend",
// 			Username: "htownsend",
// 			Email:    "henry_townsend@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Sullivan",
// 			Username: "wsullivan",
// 			Email:    "walter.s@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "John Sunderland",
// 			Username: "jnsunderland",
// 			Email:    "john.sunderland@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Heather Grayson",
// 			Username: "hgrayson",
// 			Email:    "h.grayson@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Andy Townsend",
// 			Username: "atownsend",
// 			Email:    "andy_townsend@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Bullivan",
// 			Username: "wbullivan",
// 			Email:    "walter.b@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "Jaime Sunderland",
// 			Username: "jmesunderland",
// 			Email:    "jme.sunderland@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Weather Mayson",
// 			Username: "wmayson",
// 			Email:    "w.mayson@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Fenrir Townsend",
// 			Username: "ftownsend",
// 			Email:    "fenrir_townsend@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Boulevard",
// 			Username: "wboulevard",
// 			Email:    "walter.blvd@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "Zend Southampton",
// 			Username: "zsouthampton",
// 			Email:    "z.souhampton@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Leather Mayson",
// 			Username: "lmayson",
// 			Email:    "l.mayson@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Henry Brownsend",
// 			Username: "hbrownsend",
// 			Email:    "henry_brownsend@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Salter Mullivan",
// 			Username: "smullivan1",
// 			Email:    "salter.m1@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "James Sunderland",
// 			Username: "jsunderland1",
// 			Email:    "james.sunderland1@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Heather Mayson",
// 			Username: "hmayson1",
// 			Email:    "h.mayson1@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Henry Townsend",
// 			Username: "htownsend1",
// 			Email:    "henry_townsend1@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Sullivan",
// 			Username: "wsullivan1",
// 			Email:    "walter.s1@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "John Sunderland",
// 			Username: "jnsunderland1",
// 			Email:    "john.sunderland1@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Heather Grayson",
// 			Username: "hgrayson1",
// 			Email:    "h.grayson1@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Andy Townsend",
// 			Username: "atownsend1",
// 			Email:    "andy_townsend1@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Bullivan",
// 			Username: "wbullivan1",
// 			Email:    "walter.b1@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "Jaime Sunderland",
// 			Username: "jmesunderland1",
// 			Email:    "jme.sunderland1@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Weather Mayson",
// 			Username: "wmayson1",
// 			Email:    "w.mayson1@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Fenrir Townsend",
// 			Username: "ftownsend1",
// 			Email:    "fenrir_townsend1@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Boulevard",
// 			Username: "wboulevard1",
// 			Email:    "walter.blvd@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "Zend Southampton",
// 			Username: "zsouthampton1",
// 			Email:    "z.souhampton1@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Leather Mayson",
// 			Username: "lmayson1",
// 			Email:    "l.mayson1@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Henry Brownsend",
// 			Username: "hbrownsend1",
// 			Email:    "henry_brownsend1@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Salter Mullivan",
// 			Username: "smullivan1",
// 			Email:    "salter.m1@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "James Sunderland",
// 			Username: "jsunderland2",
// 			Email:    "james.sunderland2@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Heather Mayson",
// 			Username: "hmayson2",
// 			Email:    "h.mayson2@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Henry Townsend",
// 			Username: "htownsend2",
// 			Email:    "henry_townsend2@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Sullivan",
// 			Username: "wsullivan2",
// 			Email:    "walter.s2@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "John Sunderland",
// 			Username: "jnsunderland2",
// 			Email:    "john.sunderland2@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Heather Grayson",
// 			Username: "hgrayson2",
// 			Email:    "h.grayson2@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Andy Townsend",
// 			Username: "atownsend2",
// 			Email:    "andy_townsend2@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Bullivan",
// 			Username: "wbullivan2",
// 			Email:    "walter.b2@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "Jaime Sunderland",
// 			Username: "jmesunderland2",
// 			Email:    "jme.sunderland2@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Weather Mayson",
// 			Username: "wmayson2",
// 			Email:    "w.mayson2@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Fenrir Townsend",
// 			Username: "ftownsend2",
// 			Email:    "fenrir_townsend2@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Walter Boulevard",
// 			Username: "wboulevard2",
// 			Email:    "walter.blvd2@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 		{
// 			Name:     "Zend Southampton",
// 			Username: "zsouthampton2",
// 			Email:    "z.souhampton2@acme.corp",
// 			Address:  "11 Katz St., Pennsylvania, Centralia, M4A2T6",
// 		},
// 		{
// 			Name:     "Leather Mayson",
// 			Username: "lmayson2",
// 			Email:    "l.mayson2@acme.corp",
// 			Address:  "24 Lindsey St., British Columbia, Vancouver, N9M2",
// 		},
// 		{
// 			Name:     "Henry Brownsend",
// 			Username: "hbrownsend2",
// 			Email:    "henry_brownsend2@acme.corp",
// 			Address:  "10 Rendell St., Ontario, Toronto, M2K3B8",
// 		},
// 		{
// 			Name:     "Salter Mullivan",
// 			Username: "smullivan2",
// 			Email:    "salter.m2@acme.corp",
// 			Address:  "9 Wiltse Road, Alberta, Canmore, N9W4H9 ",
// 		},
// 	}

// 	result := getDB().Create(&users)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func insertDummyPosts() error {
// 	anchorText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dol..."
// 	posts := []models.Post{
// 		{
// 			UserID: 1,
// 			Title:  "I got a letter",
// 			Body:   anchorText,
// 		},
// 		{
// 			UserID: 1,
// 			Title:  "What a nice town",
// 			Body:   anchorText,
// 		},
// 		{
// 			UserID: 1,
// 			Title:  "I love my wife, Mary",
// 			Body:   anchorText,
// 		},
// 		{
// 			UserID: 1,
// 			Title:  "How can anyone eat pizza at a time like this?",
// 			Body:   anchorText,
// 		},
// 		{
// 			UserID: 2,
// 			Title:  "I love my husband, Brad",
// 			Body:   anchorText,
// 		},
// 		{
// 			UserID: 2,
// 			Title:  "I can definitely eat pizza at a time like this",
// 			Body:   anchorText,
// 		},
// 	}
// 	result := getDB().Create(&posts)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }

// func seedDatabase() {
// 	if err := insertSeedUsers(); err != nil {
// 		log.Println("failed to insert seed users", err)
// 	}
// 	if err := insertDummyPosts(); err != nil {
// 		log.Println("failed to insert seed posts", err)
// 	}
// }
