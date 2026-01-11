# MGSmartFlow

MGSmartFlow é um **layout customizado para Fyne** que facilita a criação de interfaces dinâmicas baseadas em **linhas e colunas**, com controle inteligente de espaçamento, redimensionamento e posicionamento de widgets.

Ele foi projetado para simplificar layouts complexos sem depender apenas dos layouts padrões do Fyne.

---

## ✨ Recursos

* 📐 Layout por **linhas** e **colunas**
* 📏 Redimensionamento automático ou fixo por widget
* 🧭 Posicionamento manual opcional
* ↔️ Espaçamento global ou individual
* 🔄 Atualização dinâmica do layout
* 🧩 Wrapper simples para uso direto

---

## 📦 Instalação

```bash
go get github.com/mugomes/mgsmartflow
```

---

## 🚀 Uso Básico

### Criando o SmartFlow

```go
flow := mgsmartflow.New()
```

### Adicionando uma linha

```go
flow.AddRow(widget.NewLabel("Linha única"))
```

### Adicionando colunas na mesma linha

```go
flow.AddColumn(
	widget.NewButton("Botão 1", nil),
	widget.NewButton("Botão 2", nil),
)
```

---

## 📐 Controle de Layout

### Redimensionar um widget

```go
flow.Resize(btn, 120, 40)
```

### Mover manualmente um widget

```go
flow.Move(btn, 10, 20)
```

### Espaçamento individual entre widgets

```go
flow.Gap(btn, 15, 10)
```

### Espaçamento global do layout

```go
flow.GlobalGap(10, 10)
```

---

## ⚠️ Funções Depreciadas

As funções abaixo ainda funcionam, mas foram mantidas apenas por compatibilidade:

* `SetResize` → use `Resize`
* `SetMove` → use `Move`
* `SetGap` → use `Gap`
* `SetGlobalGap` → use `GlobalGap`

Essas funções serão removidas na versão 1.2.0

---

## 🧩 Compatibilidade

* Go 1.25.5+
* Fyne 2.7.1

---

## 👤 Autor

**Murilo Gomes Julio**

🔗 [https://mugomes.github.io](https://mugomes.github.io)

📺 [https://youtube.com/@mugomesoficial](https://youtube.com/@mugomesoficial)

---

## License

Copyright (c) 2025-2026 Murilo Gomes Julio

Licensed under the [MIT](https://github.com/mugomes/mgsmartflow/blob/main/LICENSE) license.

All contributions to the MGSmartFlow are subject to this license.