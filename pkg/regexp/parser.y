%{
package regexp
%}

%union {
	char RUNE
	node Node
}
// 表达式
%type <node> expr repatedExpr andExpr orExpr

// token
%token <char> CHAR
%token '*' '.' '|' '(' ')' '\\'

%%

top:
	orExpr
	{
		SetResult(YoYolex, $1)
	}

orExpr:
	andExpr
	{
		$$ = $1
	}
|
	orExpr '|' andExpr
	{
		$$ = OrNode{Left:$1, Right:$3}
	}

andExpr:
	repatedExpr
	{
		$$ = $1
	}
|
	andExpr repatedExpr
	{
		$$ = AndNode{Left:$1, Right:$2}
	}

repatedExpr:
	expr
	{
		$$ = $1
	}
|
	repatedExpr '*'
	{
		$$ = RepeatedNode{Val:$1}
	}

expr:
	CHAR
	{
		$$ = $1
	}
|
	'.'
	{
		$$ = AnyChar('.')
	}
|
	'\\' '.'
	{
		$$ = RUNE('.')
	}
|
	'\\' '|'
	{
		$$ = RUNE('|')
	}
|
	'\\' '*'
	{
		$$ = RUNE('*')
	}
|
	'\\' '('
	{
		$$ = RUNE('(')
	}
|
	'\\' ')'
	{
		$$ = RUNE(')')
	}
|
	'\\' '\\'
	{
		$$ = RUNE('\\')
	}
| 	'(' orExpr ')'
	{
		$$ = $2
	}

%%
