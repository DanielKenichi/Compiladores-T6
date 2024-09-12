package generator

import (
	"log"

	parser "github.com/DanielKenichi/Compiladores-T6/Antlr"
	"github.com/DanielKenichi/Compiladores-T6/GraphGen/symboltable"
	"github.com/DanielKenichi/Compiladores-T6/GraphGen/visitor"
)

type GraphGenGenerator struct {
	parser.BaseGraphGenVisitor
	Visitor *visitor.GraphGenVisitor
}

func New(visitor *visitor.GraphGenVisitor) *GraphGenGenerator {
	return &GraphGenGenerator{
		Visitor: visitor,
	}
}

func (g *GraphGenGenerator) VisitProgram(ctx parser.IProgramContext) []string {
	var programResult = make([]string, 0)

	log.Print("Generating code...")

	programResult = append(programResult, "import networkx as nx\n")
	programResult = append(programResult, "import matplotlib.pyplot as plt\n")

	// len(ctx.AllDraw_command()))

	result := g.VisitDrawCommands(ctx.AllDraw_command())
	programResult = append(programResult, result...)

	return programResult
}

func (g *GraphGenGenerator) VisitDrawCommands(ctxs []parser.IDraw_commandContext) []string {
	var drawCommandResult = make([]string, 0)

	log.Print("Visiting draw commands...")

	for _, ctx := range ctxs {
		// log.Print("command: ", ctx.GetText())

		drawCommandResult = append(drawCommandResult, "G = nx.Graph()\n")

		varType := g.Visitor.Scopes.CurrentScope().GetType(ctx.IDENT(0).GetText())

		if varType == symboltable.GROUP {
			entities := g.Visitor.Relations.Groups[ctx.IDENT(0).GetText()]

			for i := 0; i < len(entities); i++ {
				for j := i + 1; j < len(entities); j++ {
					if entities[i] != entities[j] {
						drawCommandResult = append(drawCommandResult, "G.add_edge(\""+entities[i].Name+"\", \""+entities[j].Name+"\")\n")
					}
				}
			}
		}

		drawCommandResult = append(drawCommandResult, "nx.draw(G, with_labels=True, font_weight='bold')\n")
		drawCommandResult = append(drawCommandResult, "plt.show()\n")

	}

	return drawCommandResult
}
