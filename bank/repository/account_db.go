package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepository(db *sqlx.DB) AccountRepository {
	return accountRepositoryDB{db: db}
}

func (r accountRepositoryDB) Create(acc Account) (*Account, error) {
	query := `
		INSERT INTO accounts (customer_id, opening_date, account_type, amount, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING account_id, customer_id, opening_date, account_type, amount, status;
	`

	// ใช้ QueryRowx().StructScan() เพื่อ map ตรงเข้ากับ struct
	var newAcc Account
	err := r.db.QueryRowx(
		query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	).StructScan(&newAcc)

	if err != nil {
		return nil, err
	}

	return &newAcc, nil
}

func (r accountRepositoryDB) GetAll(customerID int) ([]Account, error) {
	query := `
		SELECT account_id, customer_id, opening_date, account_type, amount, status
		FROM accounts
		WHERE customer_id = $1
	`
	var accounts []Account
	err := r.db.Select(&accounts, query, customerID)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
