package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"goRest/controllers"
	"goRest/handler"
	middlewre "goRest/middleware"
	"goRest/models"

	"github.com/gofiber/fiber/v2"
)

func TestMain(m *testing.M) {
	os.Setenv("JWT_SECRET", "segredo-de-teste")
	os.Exit(m.Run())
}

// fakes que implementam as interfaces de serviço sem banco de dados
type fakeCustommers struct{}

func (fakeCustommers) CustommerInsert(models.Custommer) error { return nil }
func (fakeCustommers) CustommerGet() []models.Custommer {
	return []models.Custommer{{Body: "segredo"}}
}

type fakeUsers struct {
	users map[string]models.User
}

func (f *fakeUsers) UserInsert(u models.User) error {
	hash, err := handler.GenerateHashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = hash
	u.ID = "id-" + u.Username
	f.users[u.Email] = u
	return nil
}

func (f *fakeUsers) UserFind(email string) (models.User, error) {
	u, ok := f.users[email]
	if !ok {
		return models.User{}, errors.New("not found")
	}
	return u, nil
}

func newTestApp(t *testing.T) (*fiber.App, *fakeUsers) {
	t.Helper()
	users := &fakeUsers{users: map[string]models.User{}}
	app := fiber.New()
	SetupCustommersRoutes(app,
		controllers.NewCustommersController(fakeCustommers{}),
		controllers.NewUserController(users),
	)
	return app, users
}

func do(t *testing.T, app *fiber.App, method, path, token string, body any) (int, string) {
	t.Helper()
	var r io.Reader
	if body != nil {
		b, _ := json.Marshal(body)
		r = bytes.NewReader(b)
	}
	req := httptest.NewRequest(method, path, r)
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("%s %s: %v", method, path, err)
	}
	out, _ := io.ReadAll(resp.Body)
	return resp.StatusCode, string(out)
}

func TestAdminRoutesRequireToken(t *testing.T) {
	app, _ := newTestApp(t)

	for _, path := range []string{"/api/v1/admin/custommers", "/api/v1/admin"} {
		if code, body := do(t, app, http.MethodGet, path, "", nil); code != http.StatusUnauthorized {
			t.Errorf("%s sem token: esperado 401, veio %d (%s)", path, code, body)
		}
	}
}

func TestUserTokenCannotAccessAdmin(t *testing.T) {
	app, _ := newTestApp(t)
	tok, _ := middlewre.GenerateToken("u1", middlewre.RoleUser)

	if code, _ := do(t, app, http.MethodGet, "/api/v1/admin/custommers", tok, nil); code != http.StatusForbidden {
		t.Fatalf("usuário comum: esperado 403, veio %d", code)
	}
}

func TestAdminTokenAccessesAdmin(t *testing.T) {
	app, _ := newTestApp(t)
	tok, _ := middlewre.GenerateToken("a1", middlewre.RoleAdmin)

	for _, path := range []string{"/api/v1/admin/custommers", "/api/v1/admin"} {
		if code, body := do(t, app, http.MethodGet, path, tok, nil); code != http.StatusOK {
			t.Errorf("%s com admin: esperado 200, veio %d (%s)", path, code, body)
		}
	}
}

func TestRegisterAndLoginFlow(t *testing.T) {
	app, users := newTestApp(t)

	code, body := do(t, app, http.MethodPost, "/api/v1/create-user", "", map[string]string{
		"username": "ana", "email": "ana@exemplo.com", "password": "senha-forte", "role": "admin",
	})
	if code != http.StatusCreated {
		t.Fatalf("cadastro: esperado 201, veio %d (%s)", code, body)
	}
	if strings.Contains(body, "senha-forte") {
		t.Fatalf("resposta do cadastro não pode conter a senha: %s", body)
	}
	if got := users.users["ana@exemplo.com"].Role; got != "user" {
		t.Fatalf("papel enviado no JSON não pode ser aceito; role=%q", got)
	}

	// senha errada -> 401
	if code, _ := do(t, app, http.MethodPost, "/api/v1/login", "", map[string]string{
		"email": "ana@exemplo.com", "password": "errada",
	}); code != http.StatusUnauthorized {
		t.Fatalf("senha errada: esperado 401, veio %d", code)
	}

	// login correto -> token de usuário comum, sem acesso admin
	code, body = do(t, app, http.MethodPost, "/api/v1/login", "", map[string]string{
		"email": "ana@exemplo.com", "password": "senha-forte",
	})
	if code != http.StatusOK {
		t.Fatalf("login: esperado 200, veio %d (%s)", code, body)
	}
	var res struct{ Token string }
	_ = json.Unmarshal([]byte(body), &res)

	role, userID, err := middlewre.ValidateToken(res.Token)
	if err != nil || role != middlewre.RoleUser || userID != "id-ana" {
		t.Fatalf("token do login: role=%q userID=%q err=%v", role, userID, err)
	}
	if code, _ := do(t, app, http.MethodGet, "/api/v1/admin/custommers", res.Token, nil); code != http.StatusForbidden {
		t.Fatalf("usuário recém-cadastrado não pode acessar admin; veio %d", code)
	}
}
