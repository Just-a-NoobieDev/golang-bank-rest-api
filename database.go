package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)
	CreateTransfer(*Transfer) error
	GetTransfers() ([]*Transfer, error)
	GetTransferByFrom(int) (*Transfer, error)
	GetTransferByTo(int) (*Transfer, error)
}

type PosgresDatabase struct {
	db *sql.DB
}

func NewPostgresDatabase() (*PosgresDatabase, error)  {
	connStr := "user=docker_user password=docker_user dbname=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PosgresDatabase{db: db}, nil

}

func (p *PosgresDatabase) Init() error {
	 err := p.CreateAccountTable()

	 if err != nil {
		 return err
	 }
	 

	 erre := p.CreateTransferTable()

	 if erre != nil {
		 return erre
	 }

	 return nil
}

func (p *PosgresDatabase) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS accounts (
		id SERIAL PRIMARY KEY, 
		first_name VARCHAR(50), 
		last_name VARCHAR(50), 
		account_number SERIAL UNIQUE, 
		balance SERIAL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`

	_, err := p.db.Exec(query)
	return err
}

func (p *PosgresDatabase) CreateTransferTable() error {
	query := `CREATE TABLE IF NOT EXISTS transfers (
		id SERIAL PRIMARY KEY NOT NULL, 
		from_account SERIAL REFERENCES accounts(account_number) ON DELETE CASCADE, 
		to_account SERIAL REFERENCES accounts(account_number) ON DELETE CASCADE, 
		amount SERIAL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);`

	_, err := p.db.Exec(query)
	return err
}


func (p *PosgresDatabase) CreateAccount(a *Account) error {
	query := `INSERT INTO accounts (first_name, last_name, account_number, balance, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	
	stmt, err := p.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(a.FirstName, a.LastName, a.Number, a.Balance, a.CreatedAt).Scan(&a.ID)
	return err

}

func (p *PosgresDatabase) DeleteAccount(id int) error {

	query := `DELETE FROM accounts WHERE id = $1;`

	stmt, err := p.db.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}


func (p *PosgresDatabase) GetAccounts() ([]*Account, error) {
	query := `SELECT * FROM accounts;`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := []*Account{}

	for rows.Next() {
		account := &Account{}
		err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (p *PosgresDatabase) GetAccountById(id int) (*Account, error) {
	
	query := `SELECT * FROM accounts WHERE id = $1;`

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	account := &Account{}

	err = stmt.QueryRow(id).Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return account, nil

}


func (p *PosgresDatabase) CreateTransfer(t *Transfer) error {
	query := `INSERT INTO transfers (from_account, to_account, amount, created_at) VALUES ($1, $2, $3, $4) RETURNING id;`

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	err = stmt.QueryRow(t.From, t.To, t.Amount, t.CreatedAt).Scan(&t.ID)
	return err
}


func (p *PosgresDatabase) GetTransfers() ([]*Transfer, error) {
	query := `SELECT * FROM transfers;`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	transfers := []*Transfer{}

	for rows.Next() {
		transfer := &Transfer{}
		err := rows.Scan(
			&transfer.ID,
			&transfer.From,
			&transfer.To,
			&transfer.Amount,
			&transfer.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		transfers = append(transfers, transfer)
	}

	return transfers, nil

}

func (p *PosgresDatabase) GetTransferByFrom(id int) (*Transfer, error) {
	query := `SELECT * FROM transfers WHERE from_account = $1;`

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	transfer := &Transfer{}

	err = stmt.QueryRow(id).Scan(
		&transfer.ID,
		&transfer.From,
		&transfer.To,
		&transfer.Amount,
		&transfer.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return transfer, nil
}

func (p *PosgresDatabase) GetTransferByTo(id int) (*Transfer, error) {
	query := `SELECT * FROM transfers WHERE to_account = $1;`

	stmt, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	transfer := &Transfer{}

	err = stmt.QueryRow(id).Scan(
		&transfer.ID,
		&transfer.From,
		&transfer.To,
		&transfer.Amount,
		&transfer.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return transfer, nil
}
