package db

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/patrickishaf/lema/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	usersSeedFile = "./seeds/users.csv"
	addrsSeedFile = "./seeds/addresses.csv"
	postsSeedFile = "./seeds/posts.csv"
)

var dbName string = os.Getenv("DB_NAME")

func seedUsersFromFile(filename string) error {
	destDB, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening destination database:", err)
		return err
	}

	if err := destDB.AutoMigrate(&models.User{}, &models.Address{}, &models.Post{}); err != nil {
		fmt.Println("Error migrating destination database:", err)
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return err
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		var user models.User
		user.ID = row[0]
		user.Name = row[1]
		user.Email = row[2]
		user.Username = row[3]

		if result := destDB.Create(&user); result.Error != nil {
			fmt.Println("Error inserting record into destination database:", result.Error)
			return err
		}
	}

	fmt.Println("Data copied successfully!")
	return nil
}

func seedAddressesFromFile(filename string) error {
	destDB, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening destination database:", err)
		return err
	}

	if err := destDB.AutoMigrate(&models.User{}, &models.Address{}, &models.Post{}); err != nil {
		fmt.Println("Error migrating destination database:", err)
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return err
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		var address models.Address
		address.ID = row[0]
		address.UserID = row[1]
		address.Street = row[2]
		address.State = row[3]
		address.City = row[4]
		address.Zipcode = row[5]

		if result := destDB.Create(&address); result.Error != nil {
			fmt.Println("Error inserting post into destination database:", result.Error)
			return err
		}
	}

	fmt.Println("Data copied successfully!")
	return nil
}

func seedPostsFromFile(filename string) error {
	destDB, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		fmt.Println("Error opening destination database:", err)
		return err
	}

	if err := destDB.AutoMigrate(&models.User{}, &models.Address{}, &models.Post{}); err != nil {
		fmt.Println("Error migrating destination database:", err)
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return err
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}

		var post models.Post
		post.ID = row[0]
		post.UserID = row[1]
		post.Title = row[2]
		post.Body = row[3]
		post.CreatedAt = row[4]

		if result := destDB.Create(&post); result.Error != nil {
			fmt.Println("Error inserting record into destination database:", result.Error)
			return err
		}
	}

	fmt.Println("Data copied successfully!")
	return nil
}

func SeedDatabase() error {
	err := seedUsersFromFile(usersSeedFile)
	if err != nil {
		return err
	}
	err = seedAddressesFromFile(addrsSeedFile)
	if err != nil {
		return err
	}
	err = seedPostsFromFile(postsSeedFile)
	if err != nil {
		return err
	}
	return nil
}
