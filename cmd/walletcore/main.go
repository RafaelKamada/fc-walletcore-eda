package main

import (
	"database/sql"
	"fmt"

	"github.com/RafaelKamada/fc-ms-wallet/internal/database"
	"github.com/RafaelKamada/fc-ms-wallet/internal/event"
	"github.com/RafaelKamada/fc-ms-wallet/internal/usecase/create_account"
	"github.com/RafaelKamada/fc-ms-wallet/internal/usecase/create_client"
	"github.com/RafaelKamada/fc-ms-wallet/internal/usecase/create_transaction"
	"github.com/RafaelKamada/fc-ms-wallet/pkg/events"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreateEvent := event.NewTransactionCreated()
	//eventDispatcher.Register("TransactionCreated", handler)

	clientDb := database.NewClientDB(db)
	accountDB := database.NewAccountDB(db)
	transactionDB := database.NewTransactionDB(db)

	createClientUseCase := create_client.NewCreateClientUseCase(clientDb)
	createAccountUseCase := create_account.NewCreateAccountUseCase(accountDB, clientDb)
	createTransactionUseCase := create_transaction.NewCreateTransactionUseCase(transactionDB, accountDB, eventDispatcher, transactionCreateEvent)

}
