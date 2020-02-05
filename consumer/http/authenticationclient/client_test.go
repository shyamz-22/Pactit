package authenticationclient

import (
	"errors"
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"net/http"
	"testing"
)

func TestClient_AuthenticateUser(t *testing.T) {

	var username = "alice"
	var password = "s3cr3t"
	var apiKey = "2bc90bbb0c7be4e5"

	pact := &dsl.Pact{
		Consumer: "Quoki",
		Provider: "UserManager",
		PactDir:  "../pacts",
		LogDir:   "../pacts/logs",
	}
	defer pact.Teardown()

	t.Run("user exists", func(t *testing.T) {
		pact.
			AddInteraction().
			Given("user exists").
			UponReceiving("a request to authenticate").
			WithRequest(dsl.Request{
				Method:  "POST",
				Path:    dsl.String(fmt.Sprintf("/users/%s/authentication", username)),
				Headers: dsl.MapMatcher{
					"Content-Type": dsl.Like("application/x-www-form-urlencoded"),
					"X-Api-Key": dsl.Like(apiKey),
				},
				Body: dsl.MapMatcher{
					"password": dsl.Like(password),
				},
			}).
			WillRespondWith(dsl.Response{
				Status: http.StatusNoContent,
			})

		test := func() error {
			subject := New(fmt.Sprintf("http://localhost:%d", pact.Server.Port), apiKey)
			ok := subject.AuthenticateUser(username, password)

			if !ok {
				return errors.New("failed to authenticate user")
			}

			return nil
		}

		if err := pact.Verify(test); err != nil {
			t.Log(err)
			t.Fail()
		}

	})

	t.Run("invalid username and password combination", func(t *testing.T) {
		pact.
			AddInteraction().
			Given("user exists").
			UponReceiving("a request to authenticate with invalid username and password").
			WithRequest(dsl.Request{
				Method:  "POST",
				Path:    dsl.String(fmt.Sprintf("/users/%s/authentication", username)),
				Headers: dsl.MapMatcher{
					"Content-Type": dsl.Like("application/x-www-form-urlencoded"),
					"X-Api-Key": dsl.Like(apiKey),
				},
				Body: dsl.MapMatcher{
					"password": dsl.Like("invalidPassword"),
				},
			}).
			WillRespondWith(dsl.Response{
				Status: http.StatusBadRequest,
			})

		test := func() error {
			subject := New(fmt.Sprintf("http://localhost:%d", pact.Server.Port), apiKey)
			ok := subject.AuthenticateUser(username, "invalidPassword")

			if ok {
				return errors.New("expected to not authenticate user, when provided with bad credentials")
			}

			return nil
		}

		if err := pact.Verify(test); err != nil {
			t.Log(err)
			t.Fail()
		}

	})

	t.Run("returns unAuthorized", func(t *testing.T) {
		pact.AddInteraction().
			Given("user exists").
			UponReceiving("a request to authenticate without api key").
			WithRequest(dsl.Request{
				Method:  "POST",
				Path:    dsl.String(fmt.Sprintf("/users/%s/authentication", username)),
				Headers: dsl.MapMatcher{
					"Content-Type": dsl.Like("application/x-www-form-urlencoded"),
				},
				Body: dsl.MapMatcher{
					"password": dsl.Like(password),
				},
			}).WillRespondWith(dsl.Response{
				Status: http.StatusUnauthorized,
			})

		test := func() error {
			subject := New(fmt.Sprintf("http://localhost:%d", pact.Server.Port), apiKey)
			ok := subject.AuthenticateUser(username, password)

			if ok {
				return errors.New("excepted the user to be not authenticated successfully")
			}

			return nil
		}

		if err := pact.Verify(test); err != nil {
			t.Log(err)
			t.Fail()
		}
	})
}
