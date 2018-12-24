package login

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	Repository Repository
}

type updatePwd struct {
	New_password string `json:"New_password"`
}

//Login API which takes Email_id and Password then give jwt as a result
func (c *Controller) Login(w http.ResponseWriter, req *http.Request) {
	var response Response
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.Message = "Not Readable body"
		response.Status_code = 500
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
	var person Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		response.Message = "Could not unmarshal body"
		response.Status_code = 500
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
	flag := c.Repository.checkEmailId(person.Email_id)
	byteHash := []byte(flag.Password)
	plainPwd := []byte(person.Password)
	err = bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err == nil {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email_id": person.Email_id,
			"password": person.Password,
		})
		tokenString, err := token.SignedString([]byte("secret-key"))
		if err != nil {
			response.Message = "Invalid Secret Key"
			response.Status_code = 401
			jsonResponse, err := json.Marshal(response)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			// w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
		} else {
			response.Message = "JWT Created"
			response.Status_code = 200
			response.Data = tokenString
			jsonResponse, err := json.Marshal(response)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			// w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
		}
	} else {
		response.Message = "Invalid Email_id and Password"
		response.Status_code = 401
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}

// Signup API which takes Email_id and Password as input and insert it into DB.
// It also check Email_id exist or not.
func (c *Controller) SignUp(w http.ResponseWriter, req *http.Request) {
	var response Response
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var person Person
	err = json.Unmarshal(body, &person)
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	check := re.MatchString(person.Email_id)
	if check {
		flag := c.Repository.checkEmailId(person.Email_id)
		if flag.Email_id == "" && flag.Password == "" {
			password := []byte(person.Password)
			hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
			if err != nil {
				panic(err)
			}
			person.Password = string(hash)
			c.Repository.InsertCredential(person)
			response.Message = "Successfully Registered"
			response.Status_code = 201
			jsonResponse, err := json.Marshal(response)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			// w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
		} else {
			response.Message = "Email ID already exist"
			response.Status_code = 409
			jsonResponse, err := json.Marshal(response)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			// w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
		}
	} else {
		response.Message = "Not a Email ID"
		response.Status_code = 400
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}

}

func (c *Controller) UpdatePassword(w http.ResponseWriter, req *http.Request) {
	var response Response
	reqToken := req.Header.Get("Authorization")
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	if err != nil {
		response.Message = "Invalid Token"
		response.Status_code = 401
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
	claims := token.Claims.(jwt.MapClaims)
	var updatePwd updatePwd
	var person Person
	person.Email_id = claims["email_id"].(string)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response.Message = "Not Readable new password"
		response.Status_code = 500
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
	err = json.Unmarshal(body, &updatePwd)
	if err != nil {
		response.Message = "Could not unmarshal body data"
		response.Status_code = 500
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
	person.Password = updatePwd.New_password
	bytePwd := []byte(person.Password)
	hashPwd, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	person.Password = string(hashPwd)
	c.Repository.Update(person)
	response.Message = "Password Successfully Updated"
	response.Status_code = 200
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
