package authenticationclient

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"net/http"
	"testing"
)

func TestClient_AuthenticateUser(t *testing.T) {

	var username = "alice"
	var password = "s3cr3t"

	t.Run("user exists", func(t *testing.T) {
		pact := &dsl.Pact{
			Consumer: "Quoki",
			Provider: "UserManager",
			PactDir:  "../pacts",
			LogDir:   "../pacts/logs",
		}
		defer pact.Teardown()

		pact.
			AddInteraction().
			Given("user exists").
			UponReceiving("a request to authenticate").
			WithRequest(dsl.Request{
				Method:  "POST",
				Path:    dsl.String(fmt.Sprintf("/users/%s/authentication", username)),
				Headers: dsl.MapMatcher{"Content-Type": dsl.Like("application/x-www-form-urlencoded")},
				Body: dsl.MapMatcher{
					"password": dsl.Like(password),
				},
			}).
			WillRespondWith(dsl.Response{
				Status: http.StatusNoContent,
			})

		pact.Verify(func() error {
			subject := New(fmt.Sprintf("http://localhost:%d", pact.Server.Port))
			ok := subject.AuthenticateUser(username, password)

			if !ok {
				t.Fail()
			}

			return nil
		})

	})
}
