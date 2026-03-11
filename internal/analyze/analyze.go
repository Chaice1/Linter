package analyze

import (
	"go/ast"
	"go/token"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var ForbiddenWords = []string{"token", "apiKey", "password"}
var Analyzer = &analysis.Analyzer{
	Name: "LogLinter",
	Doc:  "Checks logging messages",
	Run:  run,
}

func Islog(n ast.Node) (*ast.CallExpr, bool) {
	call, ok := n.(*ast.CallExpr)

	if !ok {
		return nil, false
	}

	SelExpr, ok := call.Fun.(*ast.SelectorExpr)

	if !ok {
		return nil, false
	}

	x, ok := SelExpr.X.(*ast.Ident)

	if !ok {
		return nil, false
	}

	IsLogPkg := x.Name == "log" || x.Name == "slog" || x.Name == "zap"

	IsMethodName := SelExpr.Sel.Name == "Error" || SelExpr.Sel.Name == "Info" || SelExpr.Sel.Name == "Debug" || SelExpr.Sel.Name == "Warn" ||
		SelExpr.Sel.Name == "Errorf" || SelExpr.Sel.Name == "Infof"

	if IsMethodName && IsLogPkg {
		return call, true
	}

	return nil, false

}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, item := range pass.Files {
		ast.Inspect(item, func(n ast.Node) bool {
			call, ok := Islog(n)
			if !ok {
				return true
			}

			if len(call.Args) > 0 {
				arg := call.Args[0]
				CheckMessage(pass, arg)

				for _, item := range call.Args {
					CheckSensitiveData(pass, item)
				}

			}

			return true

		})
	}
	return nil, nil
}

func CheckMessage(pass *analysis.Pass, arg ast.Expr) {
	str, ok := arg.(*ast.BasicLit)

	if !ok || str.Kind != token.STRING {
		return
	}

	message, err := strconv.Unquote(str.Value)

	if err != nil {
		return
	}

	if unicode.IsUpper(rune(message[0])) {
		pass.Reportf(arg.Pos(), "First symbol should be low")
	}

	ValidString := regexp.MustCompile(`^[a-z0-9\s.,:;_\-\(\)\[\]]+$`)

	if !ValidString.MatchString(message) {
		pass.Reportf(arg.Pos(), "log message can contain only a-z 0-9 and basic punctuation symbols")
	}

}

func CheckSensitiveData(pass *analysis.Pass, arg ast.Expr) {
	switch expr := arg.(type) {

	case *ast.BasicLit:
		if expr.Kind == token.STRING {

			val := strings.ToLower(expr.Value)

			for _, key := range ForbiddenWords {
				if strings.Contains(val, key) {
					pass.Reportf(expr.Pos(), "log message can't containt this keyword %s", key)
					return
				}
			}
		}

	case *ast.Ident:
		name := strings.ToLower(expr.Name)
		for _, key := range ForbiddenWords {
			if strings.Contains(name, key) {
				pass.Reportf(expr.Pos(), "logging variable with sensitive name '%s'", expr.Name)
				return
			}
		}

	case *ast.BinaryExpr:
		if expr.Op == token.ADD {
			CheckSensitiveData(pass, expr.X)
			CheckSensitiveData(pass, expr.Y)
		}
	}

}
