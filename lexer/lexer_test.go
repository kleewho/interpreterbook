package lexer

import (
	"monkey/token"
	"testing"
)

func singleChar(tokenType token.TokenType) struct {
	expectedType    token.TokenType
	expectedLiteral string
} {
	return struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{expectedType: tokenType, expectedLiteral: string(tokenType)}
}

func TestNextTokenExtended(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
  x + y;
};
let result = add(five, ten);
!-/*5;
5 < 10 > 5;
if (5 < 10) {
	return true; 
} else {
	return false;
}

10 == 10;
10 != 9;
"foobar"
"foo bar"
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		singleChar(token.ASSIGN),
		{token.INT, "5"},
		singleChar(token.SEMICOLON),
		{token.LET, "let"},
		{token.IDENT, "ten"},
		singleChar(token.ASSIGN),
		{token.INT, "10"},
		singleChar(token.SEMICOLON),
		{token.LET, "let"},
		{token.IDENT, "add"},
		singleChar(token.ASSIGN),
		{token.FUNCTION, "fn"},
		singleChar(token.LPAREN),
		{token.IDENT, "x"},
		singleChar(token.COMMA),
		{token.IDENT, "y"},
		singleChar(token.RPAREN),
		singleChar(token.LBRACE),
		{token.IDENT, "x"},
		singleChar(token.PLUS),
		{token.IDENT, "y"},
		singleChar(token.SEMICOLON),
		{token.RBRACE, "}"},
		singleChar(token.SEMICOLON),
		{token.LET, "let"},
		{token.IDENT, "result"},
		singleChar(token.ASSIGN),
		{token.IDENT, "add"},
		singleChar(token.LPAREN),
		{token.IDENT, "five"},
		singleChar(token.COMMA),
		{token.IDENT, "ten"},
		singleChar(token.RPAREN),
		singleChar(token.SEMICOLON),
		singleChar(token.BANG),
		singleChar(token.MINUS),
		singleChar(token.SLASH),
		singleChar(token.ASTERISK),
		{token.INT, "5"},
		singleChar(token.SEMICOLON),
		{token.INT, "5"},
		singleChar(token.LT),
		{token.INT, "10"},
		singleChar(token.GT),
		{token.INT, "5"},
		singleChar(token.SEMICOLON),
		{token.IF, "if"},
		singleChar(token.LPAREN),
		{token.INT, "5"},
		singleChar(token.LT),
		{token.INT, "10"},
		singleChar(token.RPAREN),
		singleChar(token.LBRACE),
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		singleChar(token.SEMICOLON),
		singleChar(token.RBRACE),
		{token.ELSE, "else"},
		singleChar(token.LBRACE),
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		singleChar(token.SEMICOLON),
		singleChar(token.RBRACE),
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		singleChar(token.SEMICOLON),
		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
		singleChar(token.SEMICOLON),
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tt := range tests {
		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tt := range tests {
		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
	}
}
