package teststore_test

import (
	"testing"

	"github.com/ArtemRotov/http-rest-api/internal/app/model"
	"github.com/ArtemRotov/http-rest-api/internal/app/store"
	"github.com/ArtemRotov/http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepositoryCreate(t *testing.T) {
	s := teststore.New()

	u := model.TestUser(t)
	err := s.User().Create(u)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepositoryFindByEmail(t *testing.T) {
	s := teststore.New()

	email := "user@example.org"

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecrodNotFound.Error())

	_ = s.User().Create(model.TestUser(t))

	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
