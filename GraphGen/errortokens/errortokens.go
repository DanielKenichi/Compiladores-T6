package errortokens

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func IsTokenAError(tokenName string) bool {
	return strings.Contains(tokenName, "ERRO")
}

func WriteError(tokenName string, token antlr.Token) string {
	switch tokenName {
	case "ERROR":
		return fmt.Sprintf("Linha %s: %s - simbolo nao identificado\n", strconv.Itoa(token.GetLine()), token.GetText())
	case "ERROR_OPEN_COMMENT":
		return fmt.Sprintf("Linha %s: comentario nao fechado\n", strconv.Itoa(token.GetLine()))
	default:
		return fmt.Sprintf("Linha %s: %s - erro nao identificado\n", strconv.Itoa(token.GetLine()), token.GetText())
	}
}
