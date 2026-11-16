package middlewre

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("JWT_SECRET", "segredo-de-teste")
	os.Exit(m.Run())
}

func TestGenerateAndValidateToken(t *testing.T) {
	tok, err := GenerateToken("abc-123", RoleAdmin)
	if err != nil {
		t.Fatalf("GenerateToken: %v", err)
	}

	role, userID, err := ValidateToken(tok)
	if err != nil {
		t.Fatalf("token gerado pela API deveria ser válido: %v", err)
	}
	if role != RoleAdmin || userID != "abc-123" {
		t.Fatalf("got role=%q userID=%q", role, userID)
	}
}

func TestValidateTokenRejectsGarbage(t *testing.T) {
	if _, _, err := ValidateToken("nao-e-um-jwt"); err == nil {
		t.Fatal("esperava erro para token inválido")
	}
}
