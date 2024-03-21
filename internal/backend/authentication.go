package backend

import (
	_ "bytes"
	_ "encoding/json"
	"github.com/gorilla/sessions"
	"github.com/lucasodra/chatstodo-client/internal/constants"
	_ "log"
	"net/http"
	_ "strings"
)

// Get jwt from backend auth and set to
// the cookie jwt field
// Returns the appropriate http status code
func GetJWT(writer http.ResponseWriter,
	request *http.Request, session *sessions.Session) int {

	// WARN: DUMMY IMPLEMENTATION
	// UNCOMMENT THE CODE BELOW FOR PROPER IMPLEMENTATION

	session.Values[constants.COOKIE_JWT] = "jwt123"
	session.Save(request, writer)
	return http.StatusOK

	/*
		type PostBody struct {
			Email string `json:"email"`
		}

		postBody := PostBody{
			Email: strings.ToLower(session.Values["email"].(string)),
		}

		jsonBody, err := json.Marshal(postBody)

		if err != nil {
			log.Println("user-manager.go - getJWT(), marshal")
			log.Println(err)
			return http.StatusInternalServerError
		}

		backendRequest, err := http.NewRequest(
			http.MethodPost,
			BACKEND_AUTH_GET_JWT_URL,
			bytes.NewReader(jsonBody),
		)

		if err != nil {
			log.Println("user-manager.go - getJWT(), request")
			log.Println(err)
			return http.StatusInternalServerError
		}

		backendResponse, err := http.DefaultClient.Do(backendRequest)
		if err != nil {
			log.Println("user-manager.go - getJWT(), response")
			log.Println(err)
			return http.StatusInternalServerError
		}

		type ResponseBody struct {
			Token string `json:"token"`
			Error string `json:"error"`
		}

		var responseBody ResponseBody
			err = json.NewDecoder(backendResponse.Body).Decode(&responseBody)

			if err != nil {
				log.Println("user-manager.go - getJWT(), decode")
				log.Println(err)
				return http.StatusInternalServerError
			} else if backendResponse.StatusCode != http.StatusOK {
				log.Println("user-manager.go - getJWT(), backend response")
				log.Println(backendResponse.Status, responseBody.Error)
				return backendResponse.StatusCode
			}

		// cookie will contain the jwt token inclusive of the
		// "Bearer " prefix so future calls will just have
		// to type cast it into string without further addition
		session.Values[constants.COOKIE_JWT] = "Bearer " + responseBody.Token
		session.Save(request, writer)
		return http.StatusOK

	*/
}
