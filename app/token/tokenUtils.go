package token

func TokensEqual(a, b Token) bool {
	return a.TokenType == b.TokenType && a.Value == b.Value && a.Line == b.Line && a.Column == b.Column
}

func TokenSlicesEqual(a, b []Token) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if !TokensEqual(v, b[i]) {
			return false
		}
	}
	return true
}
