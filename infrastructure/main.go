package infrastructure

// initializeServer function uses http.ListenAndServe
// which is a blocking operation, so it must run last.
func Run() {
	initializeDatabase()
	initializeServer()
}
