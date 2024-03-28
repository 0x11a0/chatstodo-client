package backend

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/lucasodra/chatstodo-client/internal/constants"
	"github.com/lucasodra/chatstodo-client/internal/utils"
)

func RefreshData(session *sessions.Session) int {

	backendRequest, err := http.NewRequest(
		http.MethodGet,
		os.Getenv("BACKEND_REFRESH_DATA_URL"),
		nil,
	)
	if err != nil {
		log.Println("user-manager.go - RefreshData(), request")
		log.Println(err)
		return http.StatusInternalServerError
	}

	backendRequest.Header.Set(
		"Authorization", session.Values[constants.COOKIE_JWT].(string),
	)

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("user-manager.go - RefreshData(), response")
		log.Println(err)
		return http.StatusInternalServerError
	}

	type ResponseBody struct {
		Success string `json:"success"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}
	var responseBody ResponseBody

	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("user-manager.go - RefreshData(), decode")
		log.Println(err)
		return http.StatusInternalServerError
	}

	return backendResponse.StatusCode
}

// Get the summaries from the backend.
// Also returns the appropriate http status code.
// Expected: http.StatusOK
func GetSummaries(session *sessions.Session) ([]*Summary, int) {

	backendRequest, err := http.NewRequest(
		http.MethodGet,
		os.Getenv("BACKEND_GET_SUMMARY_URL"),
		nil,
	)
	if err != nil {
		log.Println("user-manager.go - GetSummaries(), request")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	backendRequest.Header.Set(
		"Authorization", session.Values[constants.COOKIE_JWT].(string),
	)

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("user-manager.go - GetSummaries(), response")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}
	var responseBody []*Summary
	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("user-manager.go - GetSummaries(), decode")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	for _, summary := range responseBody {
		summary.DisplayTags = "[" + strings.Join(summary.Tags, ", ") + "]"
	}

	return responseBody, http.StatusOK
}

// Get the tasks from the backend.
// Also returns the appropriate http status code.
// Expected: http.StatusOK
func GetTasks(session *sessions.Session) ([]*Task, int) {

	backendRequest, err := http.NewRequest(
		http.MethodGet,
		os.Getenv("BACKEND_GET_TASKS_URL"),
		nil,
	)
	if err != nil {
		log.Println("user-manager.go - getTasks(), request")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	backendRequest.Header.Set(
		"Authorization", session.Values[constants.COOKIE_JWT].(string),
	)

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("user-manager.go - getTasks(), response")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}
	var responseBody []*Task
	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("user-manager.go - getTasks(), decode")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	for _, task := range responseBody {
		deadlineTime := utils.ParseISOString(task.Deadline)
		task.HTMLDeadline = utils.GetLocalDateTimeDatePicker(
			deadlineTime,
			"Singapore",
		)
		task.DisplayDeadline = utils.GetLocalDateTimePretty(
			deadlineTime,
			"Singapore",
		)
		if task.HTMLDeadline == "" {
			log.Println("DateTime format error for task: ", task)
		}
		task.DisplayTags = "[" + strings.Join(task.Tags, ", ") + "]"
	}

	return responseBody, http.StatusOK
}

// Get the tasks from the backend.
// Also returns the appropriate http status code.
// Expected: http.StatusOK
func GetEvents(session *sessions.Session) ([]*Event, int) {

	backendRequest, err := http.NewRequest(
		http.MethodGet,
		os.Getenv("BACKEND_GET_EVENTS_URL"),
		nil,
	)
	if err != nil {
		log.Println("user-manager.go - getEvents(), request")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	backendRequest.Header.Set(
		"Authorization", session.Values[constants.COOKIE_JWT].(string),
	)

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("user-manager.go - getEvents(), response")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}
	var responseBody []*Event
	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("user-manager.go - getEvents(), decode")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	for _, event := range responseBody {
		startTime := utils.ParseISOString(event.DateStart)
		event.HTMLDateStart = utils.GetLocalDateTimeDatePicker(
			startTime,
			"Singapore",
		)
		event.DisplayDateStart = utils.GetLocalDateTimePretty(
			startTime,
			"Singapore",
		)
		if event.HTMLDateStart == "" {
			log.Println("DateTime format error for event DateStart: ", event)
		}

		endTime := utils.ParseISOString(event.DateEnd)
		event.HTMLDateEnd = utils.GetLocalDateTimeDatePicker(
			endTime,
			"Singapore",
		)
		event.DisplayDateEnd = utils.GetLocalDateTimePretty(
			endTime,
			"Singapore",
		)
		if event.HTMLDateEnd == "" {
			log.Println("DateTime format error for event DateEnd: ", event)
		}

		event.DisplayTags = "[" + strings.Join(event.Tags, ", ") + "]"
	}

	return responseBody, http.StatusOK
}

// Get all platforms linked from the backend. Returns the
// platform entries obtained or the appropriate http status code
func GetAllPlatforms(session *sessions.Session) ([]PlatformEntry, int) {

	backendRequest, err := http.NewRequest(
		http.MethodGet,
		os.Getenv("BACKEND_GET_ALL_PLATFORM_URL"),
		nil,
	)
	if err != nil {
		log.Println("user-manager.go - GetAllplatforms(), request")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}
	backendRequest.Header.Set(
		"Authorization", session.Values[constants.COOKIE_JWT].(string),
	)

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("user-manager.go - GetAllPlatforms(), response")
		log.Println(err)
		return nil, http.StatusInternalServerError
	} else if backendResponse.StatusCode != http.StatusOK {
		type responseBody struct {
			Error string `json:"error"`
		}
		var backendResponseBody responseBody
		err = json.NewDecoder(backendResponse.Body).Decode(&backendResponseBody)
		if err != nil {
			log.Println("user-manager.go - GetAllPlatforms(), decode")
			log.Println(err)
			return nil, http.StatusInternalServerError
		}
		log.Println("user-manager.go - GetAllPlatforms(), backend server error")
		log.Println(backendResponseBody.Error)
		return nil, backendResponse.StatusCode
	}

	type ResponseBody struct {
		Platforms []PlatformEntry `json:"platforms"`
	}
	var responseBody ResponseBody

	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("user-manager.go - GetAllPlatforms(), decode")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}
	log.Println(responseBody.Platforms)
	return responseBody.Platforms, http.StatusOK
}

// Remove a linked platform account from the backend.
// Returns the appropriate http status code
func RemovePlatform(session *sessions.Session,
	platformId int) int {

	type DeleteBody struct {
		PlatformId int `json:"platformId"`
	}
	deleteBody := DeleteBody{
		PlatformId: platformId,
	}

	jsonBody, err := json.Marshal(deleteBody)
	if err != nil {
		log.Println("user-manager.go - RemovePlatform(), marshal")
		log.Println(err)
		return http.StatusInternalServerError
	}

	backendRequest, err := http.NewRequest(
		http.MethodDelete,
		os.Getenv("BACKEND_REMOVE_PLATFORM_URL"),
		bytes.NewReader(jsonBody),
	)
	if err != nil {
		log.Println("user-manager.go - RemovePlatform(), request")
		log.Println(err)
		return http.StatusInternalServerError
	}
	backendRequest.Header.Set(
		"Authorization", session.Values[constants.COOKIE_JWT].(string),
	)

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("user-manager.go - RemovePlatform(), response")
		log.Println(err)
		return http.StatusInternalServerError
	}
	type ResponseBody struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}
	var responseBody ResponseBody
	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("user-manager.go - RemovePlatform(), decode")
		log.Println(err)
		return http.StatusInternalServerError
	} else if backendResponse.StatusCode != http.StatusCreated {
		log.Println("user-manager.go - RemovePlatform(), backend response")
		log.Println(backendResponse.Status, responseBody.Error)
		return backendResponse.StatusCode
	}

	return http.StatusOK
}

// Add a platform account to the backend. Returns
// the appropriate http status code.
// Expected: http.StatusCreated
func AddPlatform(session *sessions.Session,
	verificationCode string) int {
	type PostBody struct {
		VerificationCode string `json:"verificationCode"`
	}
	postBody := PostBody{
		VerificationCode: verificationCode,
	}

	jsonBody, err := json.Marshal(postBody)
	if err != nil {
		log.Println("user-manager.go - AddPlatform(), marshal")
		log.Println(err)
		return http.StatusInternalServerError
	}
	log.Println("jsonBody", string(jsonBody))
	backendRequest, err := http.NewRequest(
		http.MethodPost,
		os.Getenv("BACKEND_ADD_PLATFORM_URL"),
		bytes.NewReader(jsonBody),
	)
	if err != nil {
		log.Println("user-manager.go - AddPlatform(), request")
		log.Println(err)
		return http.StatusInternalServerError
	}

	backendRequest.Header.Set(
		"Authorization", session.Values[constants.COOKIE_JWT].(string),
	)
	backendRequest.Header.Set(
		"Content-Type", "application/json",
	)

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("user-manager.go - AddPlatform(), response")
		log.Println(err)
		return http.StatusInternalServerError
	}

	type ResponseBody struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}

	var responseBody ResponseBody
	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("user-manager.go - AddPlatform(), decode")
		log.Println(err)
		return http.StatusInternalServerError
	} else if backendResponse.StatusCode != http.StatusOK {
		log.Println("user-manager.go - AddPlatform(), backend response")
		log.Println(backendResponse.Status, responseBody.Error)
		return backendResponse.StatusCode
	}

	return http.StatusOK
}

// Gets all the groups which the user is in.
// Returns the groups found and the appropriate
// http status code. Expected: http.StatusOk
func GetAllGroups(session sessions.Session) ([]PlatformGroups, int) {

	backendRequest, err := http.NewRequest(
		http.MethodGet,
		os.Getenv("BACKEND_GET_ALL_GROUPS_URL"),
		nil,
	)
	if err != nil {
		log.Println("user-manager.go - GetAllGroups(), request")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	backendRequest.Header.Set(
		"Authorization", session.Values[constants.COOKIE_JWT].(string),
	)

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("user-manager.go - GetAllGroups(), response")
		log.Println(err)
		return nil, http.StatusInternalServerError
	} else if backendResponse.StatusCode != http.StatusOK {
		type responseBody struct {
			Error string `json:"error"`
		}
		var backendResponseBody responseBody
		err = json.NewDecoder(backendResponse.Body).Decode(&backendResponseBody)
		if err != nil {
			log.Println("user-manager.go - GetAllGroups(), decode")
			log.Println(err)
			return nil, http.StatusInternalServerError
		}

		log.Println("user-manager.go - GetAllGroups() backend error")
		log.Println(backendResponseBody.Error)
		return nil, backendResponse.StatusCode
	}

	type ResponseBody struct {
		Platforms []PlatformGroups `json:"platforms"`
	}
	var responseBody ResponseBody

	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("user-manager.go - GetAllGroups(), decode")
		log.Println(err)
		return nil, http.StatusInternalServerError
	}

	return responseBody.Platforms, http.StatusOK
}

// Removes a group in which the
// Returns the groups found and the appropriate
// http status code. Expected: http.StatusNoContent
func DeleteGroup(session *sessions.Session,
	groupId int, platform string) int {

	type DeleteBody struct {
		GroupId  int    `json:"groupId"`
		Platform string `json:"platform"`
	}
	deleteBody := DeleteBody{
		GroupId:  groupId,
		Platform: platform,
	}

	jsonBody, err := json.Marshal(deleteBody)
	if err != nil {
		log.Println("user-manager.go - DeleteGroup(), marshal")
		log.Println(err)
		return http.StatusInternalServerError
	}

	backendRequest, err := http.NewRequest(
		http.MethodDelete,
		os.Getenv("BACKEND_DELETE_GROUP_URL"),
		bytes.NewReader(jsonBody),
	)
	if err != nil {
		log.Println("user-manager.go - DeleteGroup(), request")
		log.Println(err)
		return http.StatusInternalServerError
	}
	backendRequest.Header.Set(
		"Authorization", session.Values[constants.COOKIE_JWT].(string),
	)

	backendResponse, err := http.DefaultClient.Do(backendRequest)
	if err != nil {
		log.Println("user-manager.go - DeleteGroup(), response")
		log.Println(err)
		return http.StatusInternalServerError
	} else if backendResponse.StatusCode == http.StatusNoContent {
		// NOTE: Successful return here
		return http.StatusNoContent
	}

	type ResponseBody struct {
		Error string `json:"error"`
	}
	var responseBody ResponseBody
	err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)
	if err != nil {
		log.Println("user-manager.go - DeleteGroup(), decode")
		log.Println(err)
		return http.StatusInternalServerError
	}

	log.Println("user-manager.go - DeleteGroup(), backend response")
	log.Println(backendResponse.Status, responseBody.Error)
	return backendResponse.StatusCode
}
