package token

import "strings"

type TokenType string
type ValueType string

type Token struct {
	Type      TokenType
	Literal   string
	ValueType ValueType // value type is only associated with identifiers and literals
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// literals
	INT    = "INT"
	STRING = "STRING"
	DATE   = "DATE"

	// Operators
	ASSIGN = "="

	LT = "<"
	GT = ">"
	LE = "<="
	GE = ">="

	EQ     = "=="
	NOT_EQ = "!="

	// Keywords
	AND = "AND"
	OR  = "OR"

	LPAREN = "("
	RPAREN = ")"

	// Size units
	UNIT = "UNIT"

	// Search Identifiers
	SIZE      = "SIZE"
	TYPE      = "TYPE"
	EXTENSION = "EXTENSION"
	NAME      = "NAME"
	TRID      = "TRID"
	PACKER    = "PACKER"
	MAGIC     = "MAGIC"
	TAG       = "TAG"
	FS        = "FS"
	LS        = "LS"
	POSITIVES = "POSITIVES"
	CRC32     = "CRC32"
	ENGINES   = "ENGINES"
	AV        = "AV"

	// Value Types
	TYPE_INT    ValueType = "INT"
	TYPE_STRING ValueType = "STRING"
	TYPE_DATE   ValueType = "DATE"
	TYPE_ENUM   ValueType = "ENUM"
	TYPE_ARRAY  ValueType = "ARRAY"
)

var TypeEnum = []string{"pe", "elf", "macho", "txt"}
var ExtensionEnum = []string{"dll", "exe", "ps1"}

var Identifiers = map[string]struct {
	Token     TokenType
	ValueType ValueType
	EnumVals  []string // Optional: only populated for enum types
}{
	"size":        {SIZE, TYPE_INT, nil},
	"type":        {TYPE, TYPE_ENUM, TypeEnum},
	"extension":   {EXTENSION, TYPE_ENUM, ExtensionEnum},
	"name":        {NAME, TYPE_STRING, nil},
	"trid":        {TRID, TYPE_ARRAY, nil},
	"packer":      {PACKER, TYPE_ARRAY, nil},
	"magic":       {MAGIC, TYPE_STRING, nil},
	"tag":         {TAG, TYPE_STRING, nil},
	"fs":          {FS, TYPE_DATE, nil},
	"ls":          {LS, TYPE_DATE, nil},
	"positives":   {POSITIVES, TYPE_INT, nil},
	"crc32":       {CRC32, TYPE_STRING, nil},
	"engines":     {ENGINES, TYPE_STRING, nil},
	"avast":       {AV, TYPE_STRING, nil},
	"avira":       {AV, TYPE_STRING, nil},
	"bitdefender": {AV, TYPE_STRING, nil},
	"clamav":      {AV, TYPE_STRING, nil},
	"comodo":      {AV, TYPE_STRING, nil},
	"drweb":       {AV, TYPE_STRING, nil},
	"eset":        {AV, TYPE_STRING, nil},
	"fsecure":     {AV, TYPE_STRING, nil},
	"kaspersky":   {AV, TYPE_STRING, nil},
	"mcafee":      {AV, TYPE_STRING, nil},
	"sophos":      {AV, TYPE_STRING, nil},
	"symantec":    {AV, TYPE_STRING, nil},
	"trendmicro":  {AV, TYPE_STRING, nil},
	"windefender": {AV, TYPE_STRING, nil},
}

var keywords = map[string]TokenType{
	"or":  OR,
	"and": AND,
}

var sizeUnits = map[string]TokenType{
	"kb": UNIT,
	"mb": UNIT,
	"gb": UNIT,
	"tb": UNIT,
}

// Update LookupIdent to return both token type and value type
func LookupIdent(ident string) (TokenType, ValueType) {
	if tok, ok := keywords[ident]; ok {
		return tok, ""
	}
	if info, ok := Identifiers[ident]; ok {
		return info.Token, info.ValueType
	}
	if tok, ok := sizeUnits[strings.ToLower(ident)]; ok {
		return tok, ""
	}
	return STRING, TYPE_STRING
}
