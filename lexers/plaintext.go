package lexers

import (
	. "github.com/alecthomas/chroma" // nolint
)

var Plaintext = Register(MustNewLexer(
	&Config{
		Name:      "plaintext",
		Aliases:   []string{"text", "plain", "no-highlight"},
		Filenames: []string{"*.txt"},
		MimeTypes: []string{"text/plain"},
		Priority:  -1,
	},
	PlaintextRules,
))
