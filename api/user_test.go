package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	mockdb "simpleBank/db/mock"
	db "simpleBank/db/sqlc"
	"simpleBank/util"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg_     db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}
	err := util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}
	e.arg_.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg_, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matched arg %v and password %s", e.arg_, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestCreateUserAPI(t *testing.T) {
	user, password := randomUser(t)
	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"username":  user.Username,
				"password":  password,
				"full_name": user.FullName,
				"email":     user.Email,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateUserParams{
					Username: user.Username,
					FullName: user.FullName,
					Email:    user.Email,
				}
				store.EXPECT().
					CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		ctrl := gomock.NewController(t)
		store := mockdb.NewMockStore(ctrl)
		tc.buildStubs(store)

		server := newTestServer(t, store)
		recorder := httptest.NewRecorder()
		url := "/users"
		jsondata, err := json.Marshal(tc.body)
		require.NoError(t, err)
		request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsondata))
		require.NoError(t, err)

		server.router.ServeHTTP(recorder, request)
		tc.checkResponse(t, recorder)
	}
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var rst UserResponse
	err = json.Unmarshal(data, &rst)
	require.NoError(t, err)
	require.Equal(t, rst.Username, user.Username)
	require.Equal(t, rst.FullName, user.FullName)
	require.Equal(t, rst.Email, user.Email)
	require.WithinDuration(t, rst.CreatedAt, user.CreatedAt, time.Second)
	require.WithinDuration(t, rst.PasswordChangedAt, user.PasswordChangedAt, time.Second)
}

func randomUser(t *testing.T) (db.User, string) {
	password := util.RandomString(8)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)
	return db.User{
		Username:          util.RandomOwner(),
		FullName:          util.RandomOwner(),
		HashedPassword:    hashedPassword,
		Email:             util.RandomEmail(),
		CreatedAt:         time.Now(),
		PasswordChangedAt: time.Time{},
	}, password
}
