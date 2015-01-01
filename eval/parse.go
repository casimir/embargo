package eval

import (
	"regexp"
	"strings"
)

// DefaultModule is used as module name when the one provided is empty.
var DefaultModule string

type (
	// Token represents a part of an evaluated block.
	Token struct {
		Module string
		Text   string
		Tokens []string
	}

	// TokenList represents a list of token. No kidding.
	TokenList map[string]Token
)

var (
	reToken = regexp.MustCompile(`'([^']*)'|(\S+)`)
	reVar   = regexp.MustCompile(`\$\{([^}]*)\}`)
)

// Parse tokenizes the input.
func Parse(input string) (vars TokenList) {
	vars = TokenList{}
	parts := reVar.FindAllStringSubmatch(input, -1)
	for _, it := range parts {
		tokens := lex(it[1])
		module, cmd := parseCmd(tokens[0])

		v := Token{Module: module, Text: it[1]}
		v.Tokens = append([]string{cmd}, tokens[1:]...)
		vars[it[0]] = v
	}
	return
}

func lex(input string) (tokens []string) {
	parts := reToken.FindAllStringSubmatch(input, -1)
	for _, it := range parts {
		tokens = append(tokens, strings.Join(it[1:], ""))
	}
	return
}

func parseCmd(input string) (module, cmd string) {
	cmd = input
	module = DefaultModule
	parts := strings.SplitN(cmd, ".", 2)
	if len(parts) > 1 {
		module = parts[0]
		cmd = parts[1]
	}
	return
}
