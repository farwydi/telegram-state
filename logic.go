package main

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
)

type State int

const (
	initial State = iota
	start
	gettingUserAge
	gettingUserEmail
	questionN1
	questionN2
)

type machine struct {
	currentState State

	lock             sync.RWMutex
	transactionTable map[State]map[string]State
}

var m = machine{
	currentState: initial,

	transactionTable: map[State]map[string]State{
		initial: {
			"command_start": start,
		},

		start: {
			"callback_collectingUserInformation": gettingUserAge,
			"callback_repeat":                    questionN1,
		},

		gettingUserAge: {
			"text": gettingUserEmail,
		},

		gettingUserEmail: {
			"text": questionN1,
		},

		questionN1: {
			"callback_answerN1": questionN2,
			"callback_answerN2": questionN2,
			"callback_answerN3": questionN2,
		},

		questionN2: {
			"callback_answerN1": start,
			"callback_answerN2": start,
			"callback_answerN3": start,
			"callback_back":     questionN1,
		},
	},
}

func main() {
	fmt.Println(m.transactionTable)

	m.Tx(context.TODO(), tgbotapi.Update{})
}

func (m *machine) dest(action string) (State, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	table := m.transactionTable[m.currentState]
	nextState, found := table[action]
	return nextState, found
}

func (m *machine) Tx(ctx context.Context, update tgbotapi.Update) {
	action := ""
	if update.CallbackQuery != nil {
		action = update.CallbackQuery.Data
	}
	if update.Message != nil {
		action = update.Message.Text
	}

	nextState, found := m.dest(action)
	if !found {
		log.Println("No tx")
		return
	}

	X(ctx, nextState)
}

func X(ctx context.Context, nextState State) {

}
