package database

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/taise-hub/webchat/src/domain/model"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestSave(t *testing.T) {
	type args struct {
		message *model.Message
	}
	tests := map[string]struct {
		args args
	}{
		"データベースに値を保存できる": {
			args: args {
				message: &model.Message{
					Text: "test message",
					UserID: 1,
				},
			},
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			gdb, err := gorm.Open(mysql.New(mysql.Config{Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
			if err != nil {
				t.Fatalf("%s", err)
			}
			mock.ExpectBegin()
			mock.ExpectExec("INSERT INTO `messages`").WillReturnResult(sqlmock.NewResult(1, 1))
			mock.ExpectCommit()

			sut := NewMessageRepository(gdb)
			err = sut.Save(test.args.message)
			assert.Nil(t, err)
		})

	}
}