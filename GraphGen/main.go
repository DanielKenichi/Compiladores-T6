package main

import (
	"log"
	"os"

	graphgen "github.com/DanielKenichi/Compiladores-T6/Antlr"
	"github.com/DanielKenichi/Compiladores-T6/GraphGen/errortokens"
	"github.com/DanielKenichi/Compiladores-T6/GraphGen/visitor"
	"github.com/DanielKenichi/Compiladores-T6/GraphGen/vocabulary"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	input, err := antlr.NewFileStream(inputFile)

	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	output, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalf("Error opening output file: %v", err)
	}

	lexer := graphgen.NewGraphGenLexer(input)

	//Instanciando vocabulario para retornar nome de display dos tokens
	vocabulary := vocabulary.New(lexer.LiteralNames, lexer.SymbolicNames)

	for {
		t := lexer.NextToken()

		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		tokenName := *vocabulary.GetDisplayName(t.GetTokenType())

		if errortokens.IsTokenAError(tokenName) {
			_, err = output.WriteString(errortokens.WriteError(tokenName, t))

			if err != nil {
				log.Fatalf("Failed writing to output file: %v", err)
			}

			break
		}

		if err != nil {
			log.Fatalf("Failed writing to output file: %v", err)
		}
	}

	lexer.Reset()

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := graphgen.NewGraphGenParser(tokens)

	parser.RemoveErrorListeners()
	customErrorListener := NewCustomErrorListener(output)
	parser.AddErrorListener(customErrorListener)

	tree := parser.Program()

	treeVisitor := visitor.New()

	semanticErrors := treeVisitor.VisitProgram(tree)

	log.Print(semanticErrors)

	if len(semanticErrors) > 0 {
		for _, semanticError := range semanticErrors {
			output.WriteString(semanticError)
		}
		return
	}
}
