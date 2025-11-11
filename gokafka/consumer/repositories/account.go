package repositories

import "gorm.io/gorm"

type BankAccount struct {
	ID            string
	AccountHolder string
	AccountType   int
	Balance       float64
}

func (BankAccount) TableName() string {
	return "ing_banks"
}

type AccountRepository interface {
	Save(bankAccount BankAccount) error
	Delete(id string) error
	FindAll() (bankAccount []BankAccount, err error)
	FindById(id string) (bankAccount BankAccount, err error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	db.Table("ing_banks").AutoMigrate(&BankAccount{})
	return accountRepository{db}
}

func (obj accountRepository) Save(bankAccount BankAccount) error {
	return obj.db.Save(bankAccount).Error
}

func (obj accountRepository) Delete(id string) error {
	return obj.db.Table("ing_banks").Where("id = ?", id).Delete(&BankAccount{}).Error
}

func (obj accountRepository) FindAll() (bankAccounts []BankAccount, err error) {
	err = obj.db.Table("ing_banks").Find(&bankAccounts).Error
	return bankAccounts, err
}

func (obj accountRepository) FindById(id string) (bankAccount BankAccount, err error) {
	err = obj.db.Table("ing_banks").Where("id = ?", id).Find(&bankAccount).Error
	return bankAccount, err

}
