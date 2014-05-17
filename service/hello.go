package service

// public method
func GetGreeting() string {
	return "Hello, " + getUsername()
}

// private method
func getUsername() string {
	return "User"
}
