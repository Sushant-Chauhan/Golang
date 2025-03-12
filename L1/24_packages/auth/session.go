package auth

func extractSession() string { //private to package
	return "loggedin"
}

func GetSession() string { //public
	return extractSession()
}
