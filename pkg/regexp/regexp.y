%{
package regexp
%}

%union {
	char rune
	expr string
	Node Node
	And AndNode
	Or OrNode
	Repeated RepeatedNode
}

%token <char> CHAR
%token '*' '.' '\\' '|' '(' ')'

%type <Node> NODE NODE2 NODE3
%type <And> AndExpr
%type <Repeated> RepeatedExpr

%%

top:
	NODE
	{
		SetResult($1)
	}

NODE:
	NODE3
	{
		$$ = $1
	}
|	NODE '|' NODE3
	{
		$$ = OrNode{Left: $1, Right: $3}
	}

NODE3:
	NODE2
	{
		$$ = $1
	}
| 	RepeatedExpr
	{
		$$ = $1
	}
|	AndExpr
	{
		$$ = $1
	}


AndExpr:
	NODE2 CHAR
	{
		$$ = AndNode{Left: $1, Right: $2}
	}
| 	RepeatedExpr CHAR
	{
		$$ = AndNode{Left: $1, Right: $2}
	}
|	AndExpr CHAR
	{
		$$ = AndNode{Left: $1, Right: $2}
	}


RepeatedExpr:
	NODE2 '*'
	{
		$$ = RepeatedNode{Val: $1}
	}

NODE2:
	CHAR
	{
		$$ = $1
	}
| 	'(' NODE ')'
	{
		$$ = $2
	}

%%
