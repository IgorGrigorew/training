package lessoncodebasic

import (
	"encoding/json"
	"errors"
)

//DecodeAndValidateRequest([]byte("{\"email\":\"\",\"password\":\"test\",\"password_confirmation\":\"test\"}"))
//DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"\",\"password_confirmation\":\"test\"}")) // CreateUserRequest{}, "password is required"
//DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"\"}")) // CreateUserRequest{}, "password confirmation is required"
//DecodeAndValidateRequest([]byte("{\"email\":\"test\",\"password\":\"test\",\"password_confirmation\":\"test2\"}")) // CreateUserRequest{}, "password does not match with the confirmation"
//DecodeAndValidateRequest([]byte("{\"email\":\"email@test.com\",\"password\":\"passwordtest\",\"password_confirmation\":\"passwordtest\"}")) // CreateUserRequest{Email:"email@test.com", Password:"passwordtest", PasswordConfirmation:"passwordtest"}, nil

// CreateUserRequest is a request to create a new user.
type CreateUserRequest struct {
Email                string `json:"email"`
Password             string `json:"password"`
PasswordConfirmation string `json:"password_confirmation"`
}

// validation errors
var (
errEmailRequired                = errors.New("email is required")
errPasswordRequired             = errors.New("password is required")
errPasswordConfirmationRequired = errors.New("password confirmation is required")
errPasswordDoesNotMatch         = errors.New("password does not match with the confirmation")
)

//которая декодирует тело запроса из JSON в структуру CreateUserRequest и валидирует ее.
//Eсли приходит невалидный JSON или структура заполнена неверно, функция должна вернуть ошибку

func DecodeAndValidateRequest(requestBody []byte) (CreateUserRequest, error) {	
 cur:= CreateUserRequest{}
err := json.Unmarshal(requestBody, &cur)
if err != nil{
return cur, err   // чтоб пройти тесты в задании
}
if cur.Email == ""{
return CreateUserRequest{}, errEmailRequired
}else if cur.Password == ""{
return CreateUserRequest{}, errPasswordRequired
}else if cur.PasswordConfirmation == ""{
return CreateUserRequest{}, errPasswordConfirmationRequired
}else if cur.Password != cur.PasswordConfirmation{
return CreateUserRequest{}, errPasswordDoesNotMatch
}
return cur, nil
}