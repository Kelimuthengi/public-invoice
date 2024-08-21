package common

// import (
// 	"fmt"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// 	"your_project/models" // replace with the actual import path
// )

// func main() {
// 	// Initialize GORM with an SQLite database for demonstration
// 	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	// Migrate the schema
// 	db.AutoMigrate(&models.User{}, &models.Parent{}, &models.Student{})

// 	// Example usage
// 	user := models.User{
// 		Username: "parentuser",
// 		Email:    "parent@example.com",
// 		Password: "securepassword",
// 	}

// 	if err := db.Create(&user).Error; err != nil {
// 		fmt.Println("Error creating user:", err)
// 		return
// 	}

// 	parent := models.Parent{
// 		Username:    "johnparent",
// 		Email:       "johnparent@example.com",
// 		Address:     "123 Parent Street",
// 		PhoneNumber: "123-456-7890",
// 		UserID:      user.ID,
// 	}

// 	if err := db.Create(&parent).Error; err != nil {
// 		fmt.Println("Error creating parent:", err)
// 		return
// 	}

// 	student1 := models.Student{
// 		Username:        "student1",
// 		AdmissionNumber: "A001",
// 		Stream:          "Science",
// 		BoardingStatus:  true,
// 		HostelName:      "Sunrise Hostel",
// 		ParentID:        parent.ID,
// 	}

// 	student2 := models.Student{
// 		Username:        "student2",
// 		AdmissionNumber: "A002",
// 		Stream:          "Arts",
// 		BoardingStatus:  false,
// 		ParentID:        parent.ID,
// 	}

// 	if err := db.Create(&student1).Error; err != nil {
// 		fmt.Println("Error creating student1:", err)
// 	} else {
// 		fmt.Println("Student1 created successfully:", student1)
// 	}

// 	if err := db.Create(&student2).Error; err != nil {
// 		fmt.Println("Error creating student2:", err)
// 	} else {
// 		fmt.Println("Student2 created successfully:", student2)
// 	}
// }
