package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Tiago", "")
		esperado := "Olá, Tiago"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, mundo' quando uma stringvazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Olá, mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("caso o idioma seja em espanhol, retornar traduzido", func(t *testing.T) {
		resultado := Ola("Ana", "espanhol")
		esperado := "Hola, Ana"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("caso o idioma seja em frances, retornar traduzido", func(t *testing.T) {
		resultado := Ola("Ana", "frances")
		esperado := "Bonjour, Ana"

		verificaMensagemCorreta(t, resultado, esperado)
	})
}
