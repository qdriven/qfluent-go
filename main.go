package main

func main() {

	// serves static files from the provided public dir (if exists)
	//app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	//	e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
	//	return nil
	//})
	//
	//if err := app.Start(); err != nil {
	//	log.Fatal(err)
	//}
}
