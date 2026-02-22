---
marp: true
theme: default
paginate: true
backgroundColor: #1a1a2e
color: #e0e0e0
style: |
  section {
    font-family: 'JetBrains Mono', 'Fira Code', monospace;
  }
  h1 {
    color: #00d4aa;
    border-bottom: 2px solid #00d4aa;
    padding-bottom: 0.3em;
  }
  h2 {
    color: #7ec8e3;
  }
  h3 {
    color: #c4b5fd;
  }
  strong {
    color: #f59e0b;
  }
  code {
    background-color: #2d2d44;
    color: #00d4aa;
    padding: 0.1em 0.3em;
    border-radius: 4px;
  }
  pre {
    background-color: #0d1117 !important;
    border: 1px solid #333;
    border-radius: 8px;
  }
  pre code {
    background-color: transparent;
    color: #e0e0e0;
  }
  table {
    font-size: 0.85em;
  }
  th {
    background-color: #2d2d44;
    color: #00d4aa;
  }
  td {
    background-color: #1e1e32;
  }
  a {
    color: #7ec8e3;
  }
  .columns { display: flex; gap: 2em; }
  .col { flex: 1; }
  blockquote {
    border-left: 4px solid #00d4aa;
    background-color: #2d2d44;
    padding: 0.5em 1em;
    font-style: italic;
  }
---

<!-- Slide 1 — Título (00:00) -->

# Inversão de Dependência

### Curso básico de Go — Aula 05

Pré-requisitos: Aulas 01–04 (todas as anteriores)

---

<!-- Slide 2 — O problema (00:00) -->

# O problema

### `Hello` retorna string — e daí?

```go
func Hello(options ...HelloOption) string {
    // ... monta saudação ...
    return greeting + ", " + cfg.Name
}

func main() {
    fmt.Println(Hello(WithName("Maria")))  // sempre stdout
}
```

❌ E se eu quiser escrever num **arquivo**?
❌ E se eu quiser mandar pra uma **resposta HTTP**?
❌ E se eu quiser capturar num **buffer** pra testar?

> `Printf` tem destino fixo no **stdout** — dependência **hard-coded**

---

<!-- Slide 3 — Investigação: Printf → Fprintf (02:00) -->

# Investigando o código-fonte

### `Printf` é só um atalho

```go
func Printf(format string, a ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, format, a...)
}
```

`Printf` = `Fprintf` com destino **fixo** em `os.Stdout`

### `Fprintf` aceita qualquer destino

```go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
```

O primeiro parâmetro é `io.Writer` — uma **interface**

---

<!-- Slide 4 — io.Writer (06:00) -->

# `io.Writer` — uma interface muito importante no Go

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

**Um único método.** Qualquer tipo que o tenha é um `io.Writer`.

---

### Quem implementa `io.Writer`?

```
                    ┌───────────────┐
                    │  io.Writer    │
                    │               │
                    │  Write([]byte)│
                    │  (int, error) │
                    └──────┬────────┘
                           │
              ┌────────────┼────────────┐
              │            │            │
     ┌────────▼───┐  ┌─────▼─────┐  ┌──▼──────────────┐
     │ *os.File   │  │  *bytes.  │  │ http.Response   │
     │ (os.Stdout)│  │  Buffer   │  │ Writer          │
     │            │  │           │  │                 │
     │ → terminal │  │ → memória │  │ → navegador     │
     └────────────┘  └───────────┘  └─────────────────┘
```

---

<!-- Slide 9 — Três contextos (18:30) -->

# Uma função, três contextos

| Contexto | Escritor | O que acontece |
|----------|----------|----------------|
| **Teste** | `*bytes.Buffer` | Dados vão pro buffer → inspecionar |
| **Terminal** | `*os.File` (stdout) | Dados aparecem no terminal |
| **HTTP** | `http.ResponseWriter` | Dados vão pro navegador |

```go
// Teste
buffer := bytes.Buffer{}
Saudar(&buffer, WithName("Maria"))

// Terminal
Saudar(os.Stdout, WithName("João"))

// HTTP
func MeuHandler(w http.ResponseWriter, r *http.Request) {
    Saudar(w, WithName("Renata"))
}
```

**A mesma função `Saudar`. Zero modificações.**

---

### Composição, não substituição

```
Hello(options...)        → gera a saudação (retorna string)
Saudar(escritor, opts...) → escreve no destino (usa Hello por dentro)
```

> Nenhum código anterior foi alterado. Só adicionamos uma camada.
