// exemplos_ponteiros.go
// Arquivo auxiliar para demonstração - NÃO é um programa completo

package main

import "fmt"

// ============================================
// EXEMPLO 1: Pass by Value (o problema)
// ============================================

type CarteiraErrada struct {
	saldo int
}

// Receiver por VALOR - faz cópia!
func (c CarteiraErrada) Depositar(quantia int) {
	c.saldo += quantia // modifica a CÓPIA
}

func exemploPassByValue() {
	carteira := CarteiraErrada{saldo: 0}
	carteira.Depositar(10)
	fmt.Println(carteira.saldo) // 0 - não mudou!
}

// ============================================
// EXEMPLO 2: Operadores & e *
// ============================================

func exemploPonteirosBasicos() {
	x := 10
	p := &x // p = ponteiro para x (endereço)

	fmt.Println("Valor de x:", x)    // 10
	fmt.Println("Endereço de x:", p) // 0xc00001a0a8
	fmt.Println("Valor via *p:", *p) // 10

	*p = 20                            // modifica valor via ponteiro
	fmt.Println("Novo valor de x:", x) // 20 - mudou!
}

// ============================================
// EXEMPLO 3: Receiver por ponteiro (solução)
// ============================================

type CarteiraCorrigida struct {
	saldo int
}

// Receiver por PONTEIRO - modifica o original!
func (c *CarteiraCorrigida) Depositar(quantia int) {
	c.saldo += quantia
}

func exemploPassByPointer() {
	carteira := CarteiraCorrigida{saldo: 0}
	carteira.Depositar(10)
	fmt.Println(carteira.saldo) // 10 - funcionou!
}

// ============================================
// EXEMPLO 4: nil pointers
// ============================================

func exemploPonteiroNil() {
	var c *CarteiraCorrigida // nil!
	// c.Depositar(10)       // 💥 PANIC! nil pointer dereference

	// Sempre inicializar:
	c = &CarteiraCorrigida{saldo: 0} // ✅
	c.Depositar(10)
	fmt.Println(c.saldo) // 10
}

// ============================================
// EXEMPLO 5: Aplicações práticas
// ============================================

// Contador que precisa modificar estado
type Contador struct {
	contagem int
}

func (c *Contador) Incrementar() {
	c.contagem++ // modifica estado - precisa ser ponteiro
}

// Struct grande - evitar cópias
type ConfiguracaoGrande struct {
	URLBancoDados string
	ChavesAPI     map[string]string
	Configuracoes []string
	// ... dezenas de campos
}

func (cfg *ConfiguracaoGrande) Validar() error {
	// ponteiro evita copiar struct grande
	if cfg.URLBancoDados == "" {
		return fmt.Errorf("URL do banco de dados obrigatória")
	}
	return nil
}

// Consistência - se um method usa *, todos usam
type BancoDados struct {
	conexao string
}

func (bd *BancoDados) Consultar(sql string) error {
	// usa ponteiro
	return nil
}

func (bd *BancoDados) Fechar() error {
	// também usa ponteiro (consistência)
	return nil
}

func (bd *BancoDados) Ping() error {
	// todos com * (consistência)
	return nil
}
