package account

import (
	"demo/password/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("data.json")
	if err != nil {
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		color.Red("Не удалось разобрать файл data.json")
		return &Vault{
			Accounts:  []Account{},
			UpdatedAt: time.Now(),
		}
	}

	return &vault
}

func (vault *Vault) AddAccount(acc Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()

	data, err := vault.ToBytes()
	if err != nil {
		color.Red("Не удалось преобразовать файл data.json")
	}

	files.WriteFile(data, "data.json")
}

// Метод для приобразования файла в Byte массив
func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (vault *Vault) FindAccountByUrl(url string) []Account {
	var accounts []Account

	for _, acc := range vault.Accounts {
		isMatched := strings.Contains(acc.Url, url)
		if isMatched {
			accounts = append(accounts, acc)
		}
	}

	return accounts
}
