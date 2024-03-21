package backend

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

/*

type User struct {
	Id int `json:"id"`
}

type Credential struct {
	Id               int    `json:"id"`
	CredentialId     string `json:"credentialID"`
	CredentialSecret string `json:"credentialSecret"`
}
*/

// URLs for the backend calls
const (
	BACKEND_AUTH_GET_JWT_URL    = ""
	BACKEND_GET_SUMMARY_URL     = ""
	BACKEND_ALL_PLATFORM_URL    = ""
	BACKEND_ADD_PLATFORM_URL    = ""
	BACKEND_REMOVE_PLATFORM_URL = ""
)

// Get the summaries from the backend
// Returns the list of tasks, events, an summaries
// if successful. Otherwise returns all nil
func getSummary(writer http.ResponseWriter,
	request *http.Request, session *sessions.Session) ([]Task, []Event, []Summary) {

	backendRequest, err := http.NewRequest(
		http.MethodPost,
		BACKEND_GET_SUMMARY_URL,
		nil,
	)
	if err != nil {
		log.Println("auth.go - getSummary(), request")
		log.Println(err)
		return nil, nil, nil
	}

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("auth.go - getSummary(), response")
		log.Println(err)
		return nil, nil, nil
	}

	type ResponseBody struct {
		Tasks     []Task    `json:"tasks"`
		Events    []Event   `json:"events"`
		Summaries []Summary `json:"summaries"`
	}
	var responseBody ResponseBody
	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("auth.go - getSummary(), decode")
		log.Println(err)
		return nil, nil, nil
	}

	return responseBody.Tasks, responseBody.Events, responseBody.Summaries
}

// Get all bot entries from the backend
func getAllBots(session *sessions.Session) []PlatformEntry {

	backendRequest, err := http.NewRequest(
		http.MethodPost,
		BACKEND_ADD_PLATFORM_URL,
		nil,
	)
	if err != nil {
		log.Println("auth.go - getAllBots(), request")
		log.Println(err)
		return nil
	}
	backendRequest.Header.Set(
		"Authorization", session.Values[COOKIE_JWT].(string),
	)

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("auth.go - getAllBots(), response")
		log.Println(err)
		return nil
	}

	var PlatformEntries []PlatformEntry
	err = json.NewDecoder(backendResponse.Body).Decode(&PlatformEntries)
	if err != nil {
		log.Println("auth.go - getAllBots(), decode")
		log.Println(err)
		return nil
	}

	return PlatformEntries
}

// Remove a linked platform account from the backend
func removePlatform(session *sessions.Session,
	verificationCode string) error {

	type PostBody struct {
		VerificationCode string `json:"verificationCode"`
	}
	postBody := PostBody{
		VerificationCode: verificationCode,
	}

	jsonBody, err := json.Marshal(postBody)
	if err != nil {
		log.Println("auth.go - addBot(), marshal")
		log.Println(err)
		return err
	}

	backendRequest, err := http.NewRequest(
		http.MethodPost,
		BACKEND_ADD_PLATFORM_URL,
		bytes.NewReader(jsonBody),
	)
	if err != nil {
		log.Println("auth.go - addBot(), request")
		log.Println(err)
		return err
	}
	backendRequest.Header.Set(
		"Authorization", session.Values[COOKIE_JWT].(string),
	)

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("auth.go - addBot(), response")
		log.Println(err)
		return err
	}

	type ResponseBody struct {
		Message string `json:"message"`
	}
	var responseBody ResponseBody
	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("auth.go - addBot(), decode")
		log.Println(err)
		return err
	}

	if responseBody.Message == "Platform link added successfully." {
		return nil
	}

	return nil
}

// Add a platform account to the backend
func addPlatform(session *sessions.Session,
	verificationCode string) error {

	type PostBody struct {
		VerificationCode string `json:"verificationCode"`
	}
	postBody := PostBody{
		VerificationCode: verificationCode,
	}

	jsonBody, err := json.Marshal(postBody)
	if err != nil {
		log.Println("auth.go - addBot(), marshal")
		log.Println(err)
		return err
	}

	backendRequest, err := http.NewRequest(
		http.MethodPost,
		BACKEND_ADD_PLATFORM_URL,
		bytes.NewReader(jsonBody),
	)
	if err != nil {
		log.Println("auth.go - addBot(), request")
		log.Println(err)
		return err
	}
	backendRequest.Header.Set(
		"Authorization", session.Values[COOKIE_JWT].(string),
	)

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("auth.go - addBot(), response")
		log.Println(err)
		return err
	}

	type ResponseBody struct {
		Message string `json:"message"`
	}
	var responseBody ResponseBody
	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("auth.go - addBot(), decode")
		log.Println(err)
		return err
	}

	if responseBody.Message == "Platform link added successfully." {
		return nil
	}

	return nil
}

// Get jwt from backend auth and set to
// the cookie jwt field
// Returns any errors found, or nil if succeeds
func getJWT(writer http.ResponseWriter,
	request *http.Request, session *sessions.Session) error {

	type PostBody struct {
		Email string `json:"email"`
	}
	postBody := PostBody{
		Email: session.Values["email"].(string),
	}

	jsonBody, err := json.Marshal(postBody)
	if err != nil {
		log.Println("auth.go - getJWT(), marshal")
		log.Println(err)
		return err
	}

	backendRequest, err := http.NewRequest(
		http.MethodPost,
		BACKEND_AUTH_GET_JWT_URL,
		bytes.NewReader(jsonBody),
	)
	if err != nil {
		log.Println("auth.go - getJWT(), request")
		log.Println(err)
		return err
	}

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("auth.go - getJWT(), response")
		log.Println(err)
		return err
	}

	type ResponseBody struct {
		Token string `json:"token"`
		Error string `json:"error"`
	}
	var responseBody ResponseBody
	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("auth.go - getJWT(), decode")
		log.Println(err)
		return err
	} else if responseBody.Error == "" {
		log.Println("auth.go - getJWT(), backend return error")
		log.Println(responseBody.Error)
		return errors.New("Backend returned error")
	}

	// cookie will contain the jwt token inclusive of the
	// "Bearer " prefix so future calls will just have
	// to type cast it into string without further addition
	session.Values[COOKIE_JWT] = "Bearer " + responseBody.Token
	session.Save(request, writer)
	return nil
}
