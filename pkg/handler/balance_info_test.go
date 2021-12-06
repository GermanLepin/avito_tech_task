package handler_test

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/mock"
// )

// const (
// 	ServeAddress = "localhost:9000"
// )

// type mockRepository struct {
// 	mock.Mock
// }

// func (r *mockRepository) BalanceInfoDB(ctx context.Context, w http.ResponseWriter, id int64) (userId int64, balance float64) {
// 	args := r.Called(ctx, w, id)
// 	argId := args.Get(0)
// 	argBalance := args.Get(1)
// 	if argId == nil {
// 		return 0, 0
// 	}

// 	fmt.Println(argId, argBalance)

// 	return 0, 0
// }

// func TestTestBalanceInfo(t *testing.T) {
// 	JSONParams := bytes.NewBuffer([]byte(`{
// 			"id":"1"
// 			}`))

// 	w := httptest.NewRecorder()
// 	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/balance-info", ServeAddress), JSONParams)
// 	ctx := context.Background()

// 	r := new(mockRepository)
// 	r.On("BalanceInfoDB", ctx, w, 1).Return(1, 100.55)

// }

// func TestBalanceInfo(t *testing.T) {
// 	t.Run("good user_id", func(t *testing.T) {
// 		JSONParams := bytes.NewBuffer([]byte(
// 			`{
// 		"id":"1"
// 		}`))

// 		w := httptest.NewRecorder()
// 		req, err := http.NewRequest("GET", "localhost:9000/balance-info", JSONParams)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		req = req.WithContext(ctx)
// 		repo := new(mockRepository)
// 		repo.On("BalanceInfo", 1, 1059.55).Return(mUser.Balance, nil)
// 		res := w.Result()
// 		defer res.Body.Close()
// 		x, err := io.ReadAll(res.Body)
// 		fmt.Println(repo)
// 		fmt.Println(x)
// 		fmt.Println("rrr")

// 	})

// 	r := httptest.NewRecorder()
// 	handler := http.HandlerFunc(BalanceInfo)
// 	handler.ServeHTTP(r, req)

// 	if status := r.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	result := string(body)
// 	expected := "{\"user id\":1,\"top up an amount\":1059.55}\n"

// 	repo.On("BalanceInfo", 1, 1059.55).Return(mUser.mockUser, nil)

//service := HttpService{
//	UserService: repo,
//}
//service.GetReceivers(w, r)
// 	res := w.Result()
// 	defer res.Body.Close()
// 	str, err := io.ReadAll(res.Body)
// 	require.Nil(t, err)
// 	var resBody mockReceivers
// 	err = json.Unmarshal(str, &resBody)
// 	require.Nil(t, err)
// 	require.Equal(t, http.StatusOK, w.Result().StatusCode)
// 	require.Equal(t, mUser.Receivers, resBody.Receivers)

// 	if result != expected {
// 		t.Errorf("handler body: got %v want %s",
// 			result, expected)
// 	}
// }

// func TestpBalanceErrorUserId(t *testing.T) {
// 	JSONparams := bytes.NewBuffer([]byte(
// 		`{
// 				"id":"-1",
// 				"amount":"10590.55"
// 			}`))

// 	req, err := http.NewRequest("POST", "localhost:9000/balance-info", JSONparams)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ctx := context.Background()
// 	w := httptest.NewRecorder()
// 	r := new(mockRepository)
// 	r.On("BalanceInfoDB", ctx, w, 1).Return(1, 100.55)
// 	service := handler.HttpService{
// 		UserService: r,
// 	}
// 	handler := http.HandlerFunc(service.BalanceInfo)
// 	handler.ServeHTTP(w, req)

// 	if status := w.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	body, err := ioutil.ReadAll(w.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	result := string(body)
// 	expected := "{\"error\":\"Incorrect value id user\"}\n"

// 	if result != expected {
// 		t.Errorf("handler body: got %v want %v",
// 			result, expected)
// 	}
// }

// func TestUpBalanceErrorAmoun(t *testing.T) {
// 	JSONparams := bytes.NewBuffer([]byte(
// 		`{
// 				"id":"1",
// 				"amount":"-10590.55"
// 			}`))

// 	req, err := http.NewRequest("POST", "localhost:9000/balance-info", JSONparams)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ctx := context.Background()
// 	w := httptest.NewRecorder()
// 	r := new(mockRepository)
// 	r.On("BalanceInfoDB", ctx, w, 1).Return(1, 100.55)
// 	service := handler.HttpService{
// 		UserService: r,
// 	}
// 	handler := http.HandlerFunc(service.BalanceInfo)
// 	handler.ServeHTTP(w, req)

// 	if status := w.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	body, err := ioutil.ReadAll(w.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	result := string(body)
// 	expected := "{\"error\":\"The amount is negative\"}\n"

// 	if result != expected {
// 		t.Errorf("handler body: got %v want %v",
// 			result, expected)
// 	}
// }

// func TestBalanceInfoErrorFindUserIdDB(t *testing.T) {
// 	JSONparams := bytes.NewBuffer([]byte(
// 		`{
// 				"id":"9999999",
// 				"amount":"10590.55"
// 			}`))

// 	req, err := http.NewRequest("POST", "localhost:9000/balance-info", JSONparams)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	ctx := context.Background()
// 	w := httptest.NewRecorder()
// 	r := new(mockRepository)
// 	r.On("BalanceInfoDB", ctx, w, 1).Return(1, 100.55)
// 	service := handler.HttpService{
// 		UserService: r,
// 	}
// 	handler := http.HandlerFunc(service.BalanceInfo)
// 	handler.ServeHTTP(w, req)

// 	if status := w.Code; status != http.StatusBadRequest {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	body, err := ioutil.ReadAll(w.Body)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	result := string(body)
// 	expected := "{\"error\":\"User not found in database\"}\n"

// 	if result != expected {
// 		t.Errorf("handler body: got %v want %v",
// 			result, expected)
// 	}
// }
