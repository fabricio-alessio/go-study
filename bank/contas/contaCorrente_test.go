package contas

import (
	"testing"

	"alessio.com/study/clientes"
	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	contaDoGuilherme := ContaCorrente{
		Titular:       clientes.Titular{Nome: "Guilherme"},
		NumeroAgencia: 2344,
		NumeroConta:   83929283,
	}

	if contaDoGuilherme.saldo != 0 {
		t.Error("Saldo de conta criada precisa ser zero")
	}
}

func TestXx2(t *testing.T) {
	contaDoGuilherme := ContaCorrente{
		Titular:       clientes.Titular{Nome: "Guilherme"},
		NumeroAgencia: 2344,
		NumeroConta:   83929283,
	}

	assert.Equal(t, 1, contaDoGuilherme.saldo, "Saldo de conta criada precisa ser zero")
}
