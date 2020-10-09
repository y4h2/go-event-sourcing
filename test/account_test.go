package account_test

import (
	"event-sourcing/enum"
	"event-sourcing/eventstore"
	"event-sourcing/service"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AccountTestSuite struct {
	suite.Suite

	eventStore    *eventstore.EventStore
	accountEntity *service.AccountEntity
}

func (suite *AccountTestSuite) SetupTest() {
	suite.eventStore = eventstore.NewEventStore()
	suite.accountEntity = service.NewAccountEntity(suite.eventStore)
}

func (suite *AccountTestSuite) TeardownTest() {
	suite.eventStore.CleanUp()
}

func (suite *AccountTestSuite) TestOpen() {

	account1 := "account1"
	err := suite.accountEntity.Open(account1, 100)
	suite.NoError(err)

	events := suite.eventStore.ListAll()

	suite.Equal(events[0].ID, 0)
	suite.Equal(events[0].Type, enum.EventType.Open)
	suite.Equal(events[0].Entity, enum.EventEntity.Account)
	suite.Equal(events[0].EntityID, account1)
}

func (suite *AccountTestSuite) Transfer() {

}
func (suite *AccountTestSuite) Store() {

}

func (suite *AccountTestSuite) Withdraw() {

}

func TestAccountSuite(t *testing.T) {
	suite.Run(t, new(AccountTestSuite))
}
