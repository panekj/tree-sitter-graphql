package tree_sitter_graphql_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-graphql"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_graphql.Language())
	if language == nil {
		t.Errorf("Error loading Graphql grammar")
	}
}
