package firebase_create_test_token

import (
	"bytes"
	"context"
	"encoding/json"
	"firebase.google.com/go/auth"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CreateTestToken creates a custom token for the specified UID and verifies it with the Firebase API
func CreateTestToken(client *auth.Client, firebaseKey, uid string) (string, error) {
	token, err := client.CustomToken(context.Background(), uid)

	if err != nil {
		return "", err
	}

	requestBody, _ := json.Marshal(map[string]interface{}{
		"token":             token,
		"returnSecureToken": true,
	})

	res, err := http.Post("https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key="+firebaseKey, "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return "", err
	}

	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	bodyMap := map[string]interface{}{}

	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		return "", err
	}

	generatedToken, _ := bodyMap["idToken"]

	return fmt.Sprintf("%v", generatedToken), nil
}
