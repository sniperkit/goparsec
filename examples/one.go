/*
Sniperkit-Bot
- Status: analyzed
*/

package main

import (
	"fmt"

	"github.com/sniperkit/snk.fork.goparsec"
)

var input = `[a,[aa,[aaa],ab,ac],b,c,[ca,cb,cc,[cca]]]`

func main() {
	var array parsec.Parser
	ast := parsec.NewAST("one", 100)

	id := parsec.And(
		func(ns []parsec.ParsecNode) parsec.ParsecNode {
			t := ns[0].(*parsec.Terminal)
			t.Value = `"` + t.Value + `"`
			return t
		},
		parsec.Token(`[a-z]+`, "ID"),
	)
	opensqr := parsec.Atom(`[`, "OPENSQR")
	closesqr := parsec.Atom(`]`, "CLOSESQR")
	comma := ast.Maybe("comma", nil, parsec.Atom(`,`, "COMMA"))

	item := ast.OrdChoice("item", nil, id, &array)
	itemsep := ast.And("itemsep", nil, item, comma)
	items := ast.Kleene("items", nil, itemsep, nil)
	array = ast.And("array", nil, opensqr, items, closesqr)

	s := parsec.NewScanner([]byte(input))
	node, _ := ast.Parsewith(array, s)
	fmt.Println(node.GetValue())
}
