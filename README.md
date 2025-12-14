# MGSmartFlow

MGSmartFlow é um layout flexível para Fyne que distribui widgets em linhas, calculando tamanhos e espaçamentos automaticamente.

## Instalação

`go get github.com/mugomes/mgsmartflow`

## Documentação

O MGSmartFlow foi desenvolvido para respeitar o tamanho da largura e altura da janela, portanto, caso o texto for muito longo, o mesmo será cortado e não exibirá a outra parte até que você aumente a janela ou use wrap para quebrar o texto.

**Exemplo de Uso**

```
flow := mgsmartflow.New()

/* Define o gap padrão para todas as linhas e colunas */
flow.SetGlobalGap(10, 8)

/* Label */
lbl1 := widget.NewLabel("Exemplo de Linha")
flow.AddRow(lbl1)
	
/* Botões em Colunas */
a := widget.NewButton("A", nil)
b := widget.NewButton("B", nil)
flow.AddColumn(a, b)

/* Redimensiona as colunas */
flow.SetResize(a, fyne.NewSize(120, 40))
flow.SetResize(b, fyne.NewSize(120, 40))

/* Ajusta o gap do button B */
flow.SetGap(b, fyne.NewPos(30, 0))

window.SetContent(flow.Container)
```

O SetResize e SetGap podem ser utilizados para ajustar linhas e colunas, o Gap afeta sempre o que vem na frente e não antes.

- Caso o SetResize não seja definido, será aplicado o espaço total da largura da janela ou da coluna.
- Caso o SetGap não seja definido, será aplicado automaticamente um espaço por padrão.

## Information

 - [Page MGSmartFlow](https://github.com/mugomes/mgsmartflow)

## Requirement

 - Go 1.24.6
 - Fyne 2.7.1

## Support

- GitHub: https://github.com/sponsors/mugomes
- More: https://www.mugomes.com.br/p/apoie.html

## License

Copyright (c) 2025 Murilo Gomes Julio

Licensed under the [MIT](https://github.com/mugomes/mgsmartflow/blob/main/LICENSE) license.

All contributions to the MGSmartFlow are subject to this license.
