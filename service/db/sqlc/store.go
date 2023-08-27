package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func newStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type RecordTestTxParams struct {
	UserID         int32 `json:"from_user"`
	ReadingScore   int32 `json:"reading_score"`
	ListeningScore int32 `json:"listening_score"`
}
type RecordTestTxResult struct {
	Test           Test  `json:"test"`
	UserID         int32 `json:"from_user"`
	ReadingScore   Score `json:"reading_score"`
	ListeningScore Score `json:"listening_score"`
}

func (store *Store) SubmitTx(ctx context.Context, arg RecordTestTxParams) (RecordTestTxResult, error) {
	var result RecordTestTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Test, err = q.CreateTest(ctx, CreateTestParams{
			UserID: arg.UserID,
		})
		if err != nil {
			return err
		}
		result.UserID = arg.UserID
		result.ReadingScore, err = q.CreateScore(ctx, CreateScoreParams{
			Score: arg.ReadingScore,
			Type:  "reading",
		})
		if err != nil {
			return err
		}
		result.ListeningScore, err = q.CreateScore(ctx, CreateScoreParams{
			TestID: result.Test.TestID,
			Score:  arg.ListeningScore,
			Type:   "listening",
		})
		if err != nil {
			return err
		}
		return nil
	})
	// 1. Insert a new test record for the user
	// 2. Insert the scores for reading and listening
	return result, err
}
