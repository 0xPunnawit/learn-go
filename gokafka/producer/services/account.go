package services

import (
	"errors"
	"events"
	"log"
	"producer/commands"

	"github.com/google/uuid"
)

type AccountService interface {
	OpenAccount(commands commands.OpenAccountCommand) (id string, err error)
	DepositFund(commands commands.DepositFundCommand) error
	WithdrawFund(commands commands.WithdrawFundCommand) error
	CloseAccount(commands commands.CloseAccountCommand) error
}

type accountService struct {
	eventProducer EventProducer
}

func NewAccountService(eventProducer EventProducer) AccountService {
	return accountService{eventProducer}
}

func (obj accountService) OpenAccount(commands commands.OpenAccountCommand) (id string, err error) {
	if commands.AccountHolder == "" || commands.AccountType == 0 || commands.OpeningBalance == 0 {
		return "", errors.New("bad request")
	}

	event := events.OpenAccountEvent{
		ID:             uuid.NewString(),
		AccountHolder:  commands.AccountHolder,
		AccountType:    commands.AccountType,
		OpeningBalance: commands.OpeningBalance,
	}
	log.Printf("%#v", event)

	return event.ID, obj.eventProducer.Produce(event)
}

func (obj accountService) DepositFund(commands commands.DepositFundCommand) error {
	if commands.ID == "" || commands.Amount == 0 {
		return errors.New("bad request")
	}

	event := events.DepositFundEvent{
		ID:     commands.ID,
		Amount: commands.Amount,
	}
	return obj.eventProducer.Produce(event)

}

func (obj accountService) WithdrawFund(commands commands.WithdrawFundCommand) error {
	if commands.ID == "" || commands.Amount == 0 {
		return errors.New("bad request")
	}

	event := events.WithdrawFundEvent{
		ID:     commands.ID,
		Amount: commands.Amount,
	}
	return obj.eventProducer.Produce(event)
}

func (obj accountService) CloseAccount(commands commands.CloseAccountCommand) error {
	if commands.ID == "" {
		return errors.New("bad request")
	}

	event := events.CloseAccountEvent{
		ID: commands.ID,
	}
	return obj.eventProducer.Produce(event)
}
