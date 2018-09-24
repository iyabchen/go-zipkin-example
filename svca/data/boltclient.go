package data

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/iyabchen/go-zipkin-example/model"
	"github.com/iyabchen/go-zipkin-example/svca/common"
	zipkin "github.com/openzipkin/zipkin-go"
	zipkinmodel "github.com/openzipkin/zipkin-go/model"
)

// BoltClient implementation.
type BoltClient struct {
	boltDB *bolt.DB
}

// OpenDb opens a bolt connection.
func (bc *BoltClient) OpenDb() error {
	var err error
	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		return err
	}
	return nil
}

// Seed starts seeding accounts
func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts(100)
}

// Creates an "AccountBucket" in our BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

// Seed (n) make-believe account objects into the AccountBucket bucket.
func (bc *BoltClient) seedAccounts(count int) {

	total := count
	for i := 0; i < total; i++ {

		// Generate a key 10000 or larger
		key := strconv.Itoa(10000 + i)

		// Create an instance of our Account struct
		acc := model.Account{
			ID:   key,
			Name: "Person_" + strconv.Itoa(i),
		}

		// Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(acc)

		// Write the data to the AccountBucket
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake accounts...\n", total)
}

// QueryAccount query account by account ID from bolt db.
func (bc *BoltClient) QueryAccount(ctx context.Context, accountID string) (model.Account, error) {

	tracer := common.GetTracer()
	var parentContext zipkinmodel.SpanContext
	if span := zipkin.SpanFromContext(ctx); span != nil {
		parentContext = span.Context()
	}
	appSpan := tracer.StartSpan("bolt_db", zipkin.Parent(parentContext))
	defer appSpan.Finish()

	// Allocate an empty Account instance we'll let json.Unmarhal populate for us in a bit.
	account := model.Account{}

	// Read an object from the bucket using boltDB.View
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the DB
		b := tx.Bucket([]byte("AccountBucket"))

		// Read the value identified by our accountID supplied as []byte
		accountBytes := b.Get([]byte(accountID))
		if accountBytes == nil {
			return fmt.Errorf("No account found for " + accountID)
		}
		// Unmarshal the returned bytes into the account struct we created at
		// the top of the function
		json.Unmarshal(accountBytes, &account)

		// Return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	// If there were an error, return the error
	if err != nil {
		return model.Account{}, err
	}
	// Return the Account struct and nil as error.
	return account, nil
}

// Check is naive healthcheck, just makes sure the DB connection has been initialized.
func (bc *BoltClient) Check() bool {
	return bc.boltDB != nil
}
