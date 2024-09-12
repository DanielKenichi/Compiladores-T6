package generator

import (
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

	programResult = append(programResult, "import networkx as nx\n")
	programResult = append(programResult, "import matplotlib.pyplot as plt\n")

	// len(ctx.AllDraw_command()))

	result := g.VisitDrawCommands(ctx.AllDraw_command())
	programResult = append(programResult, result...)

	return programResult
}

func (g *GraphGenGenerator) VisitDrawCommands(ctxs []parser.IDraw_commandContext) []string {
	var drawCommandResult = make([]string, 0)

	for _, ctx := range ctxs {
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

			drawCommandResult = append(drawCommandResult, "nx.draw(G, with_labels=True, font_weight='bold')\n")
		}

		if varType == symboltable.PERSON {
			person := g.Visitor.Relations.Persons[ctx.IDENT(0).GetText()]

			drawCommandResult = append(drawCommandResult, "edge_labels={}\n")

			// TODO: add different color to person (root)

			// has filter
			if ctx.FILTER_BY() != nil {
				filter := ctx.GetFilter().GetText()
				for _, relatedPerson := range person.Relationships[filter] {
					drawCommandResult = append(drawCommandResult, "G.add_edge(\""+person.Name+"\", \""+relatedPerson.Name+"\")\n")
					drawCommandResult = append(drawCommandResult, "edge_labels[(\""+person.Name+"\", \""+relatedPerson.Name+"\")] = \""+filter+"\"\n")
				}
			} else {
				// no filter
				for key, value := range person.Relationships {
					for _, relatedPerson := range value {
						drawCommandResult = append(drawCommandResult, "G.add_edge(\""+person.Name+"\", \""+relatedPerson.Name+"\")\n")
						drawCommandResult = append(drawCommandResult, "edge_labels[(\""+person.Name+"\", \""+relatedPerson.Name+"\")] = \""+key+"\"\n")
					}
				}
			}

			drawCommandResult = append(drawCommandResult, "pos = nx.spring_layout(G)\n")
			drawCommandResult = append(drawCommandResult, "nx.draw(G, pos, node_color='pink', alpha=0.9, node_size=500, labels={node: node for node in G.nodes()})\n")
			drawCommandResult = append(drawCommandResult, "nx.draw_networkx_edge_labels(G, pos, edge_labels=edge_labels, font_color='red')\n")
		}

		drawCommandResult = append(drawCommandResult, "plt.show()\n")

	}

	return drawCommandResult
}
