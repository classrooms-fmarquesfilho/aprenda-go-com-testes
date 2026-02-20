// exemplos_erros.go
// Arquivo auxiliar para demonstração - NÃO é um programa completo

package main

import (
	"errors"
	"fmt"
)

// ============================================
// EXEMPLO 1: error é uma interface
// ============================================

// A interface error da standard library:
/*
type error interface {
    Error() string
}
*/

// ============================================
// EXEMPLO 2: Padrão idiomático if err != nil
// ============================================

func exemploPadraoErro() {
	resultado, err := FazerAlgo()
	if err != nil {
		// tratar erro
		fmt.Println("Erro:", err)
		return
	}
	// usar resultado com segurança
	fmt.Println("Sucesso:", resultado)
}

func FazerAlgo() (string, error) {
	// simula operação que pode falhar
	return "", errors.New("algo deu errado")
}

// ============================================
// EXEMPLO 3: Criar erros simples
// ============================================

func exemploCriarErros() error {
	// Forma 1: errors.New
	err1 := errors.New("saldo insuficiente")

	// Forma 2: fmt.Errorf (com formatação)
	err2 := fmt.Errorf("não é possível sacar %d, saldo é %d", 100, 50)

	fmt.Println(err1)
	fmt.Println(err2)
	return nil
}

// ============================================
// EXEMPLO 4: Tipo customizado de erro
// ============================================

type ErroSaldoInsuficiente struct {
	Solicitado int
	Disponivel int
}

// Satisfaz interface error automaticamente!
func (e ErroSaldoInsuficiente) Error() string {
	return fmt.Sprintf("não é possível sacar %d, disponível %d",
		e.Solicitado, e.Disponivel)
}

func Sacar(saldo int, quantia int) error {
	if quantia > saldo {
		return ErroSaldoInsuficiente{
			Solicitado: quantia,
			Disponivel: saldo,
		}
	}
	return nil
}

// ============================================
// EXEMPLO 5: Error wrapping (Go 1.13+)
// ============================================

type CarteiraSimples struct {
	saldo int
}

func ObterCarteira(usuarioID int) (*CarteiraSimples, error) {
	// simula busca
	if usuarioID == 0 {
		return nil, errors.New("ID de usuário inválido")
	}
	return &CarteiraSimples{saldo: 100}, nil
}

func ProcessarPagamento(usuarioID int, quantia int) error {
	carteira, err := ObterCarteira(usuarioID)
	if err != nil {
		// %w envolve o erro original
		return fmt.Errorf("falha ao obter carteira do usuário %d: %w", usuarioID, err)
	}

	if quantia > carteira.saldo {
		return fmt.Errorf("saldo insuficiente")
	}

	return nil
}

// ============================================
// EXEMPLO 6: errors.As e errors.Is
// ============================================

func exemploInspecaoErro() {
	err := ProcessarPagamento(0, 100)

	// errors.As - extrai tipo específico
	var erroInsuf ErroSaldoInsuficiente
	if errors.As(err, &erroInsuf) {
		fmt.Printf("Saldo insuficiente! Disponível: %d\n", erroInsuf.Disponivel)
	}

	// errors.Is - verifica erro específico
	var ErrNaoEncontrado = errors.New("não encontrado")
	if errors.Is(err, ErrNaoEncontrado) {
		fmt.Println("Carteira não encontrada")
	}
}

// ============================================
// EXEMPLO 7: Boas práticas
// ============================================

// ✅ FAÇA: Sempre verificar
func boaPratica() error {
	dados, err := FazerAlgo()
	if err != nil {
		return err
	}
	fmt.Println(dados)
	return nil
}

// ❌ NÃO FAÇA: Ignorar erro
func maPratica() {
	dados, _ := FazerAlgo() // ❌ ignorando erro!
	fmt.Println(dados)
}

// ✅ FAÇA: Retornar erro
func Dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero")
	}
	return a / b, nil
}

// ❌ NÃO FAÇA: Panic sem necessidade
func DividirRuim(a, b float64) float64 {
	if b == 0 {
		panic("divisão por zero") // ❌ evitar!
	}
	return a / b
}

// ✅ FAÇA: Mensagens em minúsculas
func mensagemErroBoa() error {
	return errors.New("saldo insuficiente") // ✅
}

// ❌ NÃO FAÇA: Mensagens maiúsculas
func mensagemErroRuim() error {
	return errors.New("Saldo Insuficiente") // ❌
}
