package main

// Use server's environment variables in production
// import "os"

func main() {
	a := App{}
	a.Initialize("root", "", "127.0.0.1:3306", "test")

	//   a.Initialize(
	// // os.Getenv("APP_DB_USERNAME"),
	//       // os.Getenv("APP_DB_PASSWORD"),
	//       // os.Getenv("APP_DB_HOST"),
	//       // os.Getenv("APP_DB_NAME")
	//   )

	a.Run(":8000")
}
