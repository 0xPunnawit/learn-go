package repository

import (
	"fmt"
	"goorm/database" // 1. import database เพื่อใช้ database.DB
	"goorm/models"   // 2. import models เพื่อใช้ models.Gender
)

// ทำให้ฟังก์ชัน return error ออกมาด้วย จะดีมากครับ
func CreateGenders(name string) error {
	gender := models.Gender{Name: name}
	result := database.DB.Create(&gender)

	if result.Error != nil {
		return result.Error
	}
	fmt.Println(gender)

	return result.Error
}

func GetGenders() {
	genders := []models.Gender{}
	tx := database.DB.Order("id").Find(&genders)
	if tx.Error != nil {
		fmt.Println("❌ Error fetching genders:", tx.Error)
		return
	}

	fmt.Println("✅ Genders found:", genders)
}

func GetGender(id uint) {
	gender := models.Gender{}
	tx := database.DB.Order("id").First(&gender, id)
	if tx.Error != nil {
		fmt.Println("❌ Error fetching gender:", tx.Error)
		return
	}

	fmt.Println("✅ Gender found:", gender)
}

func GetGenderByName(name string) {
	genders := []models.Gender{}
	tx := database.DB.Where("name = ?", name).Find(&genders)
	if tx.Error != nil {
		fmt.Println("❌ Error fetching gender:", tx.Error)
		return
	}

	fmt.Println("✅ Gender found:", genders)
}

func UpdateGender(id uint, name string) error {
	gender := models.Gender{}
	tx := database.DB.First(&gender, id)
	if tx.Error != nil {
		return tx.Error
	}

	gender.Name = name
	result := database.DB.Save(&gender)

	if result.Error != nil {
		return result.Error
	}
	fmt.Println(gender)

	return result.Error
}

func DeleteGender(id uint) error {
	var gender models.Gender

	// ลองค้นก่อนว่ามีจริงไหม เพื่อให้รู้ว่าลบได้หรือไม่
	if err := database.DB.First(&gender, id).Error; err != nil {
		return fmt.Errorf("gender with ID %d not found: %w", id, err)
	}

	// ถ้ามีข้อมูลจริง ค่อยลบ
	result := database.DB.Delete(&gender)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("✅ Gender deleted:", gender)
	return nil
}

func CreateCustomer(name string, genderID uint) {
	customer := models.Customer{
		Name:     name,
		GenderId: genderID,
	}
	tx := database.DB.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(customer)
}

func GetCustomers() {
	customers := []models.Customer{}
	tx := database.DB.Preload("Gender").Find(&customers)
	if tx.Error != nil {
		fmt.Println("❌ Error fetching customers:", tx.Error)
		return
	}

	for _, customer := range customers {
		fmt.Printf("%v | %v | %v\n", customer.ID, customer.Name, customer.Gender.Name)
	}

	fmt.Println("✅ Customers found:", customers)
}
