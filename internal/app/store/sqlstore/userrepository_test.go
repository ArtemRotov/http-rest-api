package sqlstore_test

import (
	"testing"

	"github.com/ArtemRotov/http-rest-api/internal/app/model"
	"github.com/ArtemRotov/http-rest-api/internal/app/store"
	"github.com/ArtemRotov/http-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	u := model.TestUser(t)
	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepositoryFindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecrodNotFound.Error())

	_ = s.User().Create(model.TestUser(t))

	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepositoryFind(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)

	_, err := s.User().Find(1)
	assert.EqualError(t, err, store.ErrRecrodNotFound.Error())

	u := model.TestUser(t)
	_ = s.User().Create(u)

	u, err = s.User().Find(u.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
