package gorae

import (
	"strings"

	"golang.org/x/net/html"
)

// <p class="k5" id="EmYUHVi"><u>actividad</u> específica</p> comienza definición compuesta
func parseClassK5(z *html.Tokenizer) (definicion string) {
	level := 1
	for level > 0 {
		tt := z.Next()
		tag := z.Token()

		if tt == html.StartTagToken {
			level++
		} else if tt == html.TextToken {
			definicion += tag.Data
		} else if tt == html.EndTagToken {
			level--
		}
	}
	return "*" + definicion + "*"
}

// <p class="j" id="NWrtL6E"> comienza una definición
func parseClassJ(z *html.Tokenizer) (definicion string) {
	level := 1
	for level > 0 {
		tt := z.Next()
		tag := z.Token()

		if tt == html.StartTagToken {
			level++
		} else if tt == html.TextToken {
			definicion += tag.Data
		} else if tt == html.EndTagToken {
			level--
		}
	}

	//fmt.Println("======> " + definicion)
	return
}

// <header class="f">actividad.</header> la palabra a definir
func parseHeader(z *html.Tokenizer) (definicion string) {
	z.Next() // StartTag Text EndTag
	tag := z.Token()
	definicion = "*" + tag.Data + "*"
	return
}

func htmlToText(ht string) (text string) {
	tokenizer := html.NewTokenizer(strings.NewReader(ht))
	for {
		tt := tokenizer.Next()

		if tt == html.ErrorToken {
			break
		}

		tag := tokenizer.Token()

		if tt == html.StartTagToken {
			if tag.Data == "header" {
				text += parseHeader(tokenizer)
			} else if tag.Data == "p" { // comienza un bloque
				for _, att := range tag.Attr {
					if att.Key == "class" {
						switch att.Val {
						case "j", "m": // definicion o uso
							text += "\n" + parseClassJ(tokenizer)
						case "k5", "k", "k6": // palabra compuesta
							text += "\n\n" + parseClassK5(tokenizer)
						default:
							// nada, simplemente tiene muchos atributos que no nos interesan
						}
					}
				} // for attr
			}
		}
	}

	if len(text) > 2000 {
		text = text[0:2000] + "... cortado"
	}

	return
}
