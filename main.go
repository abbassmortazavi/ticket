package main

import (
	"ticket/cmd"
)

func main() {
	cmd.Execute()

	//TODO::
	/*storage := store.NewStorage(database.DB)
	jwt := viper.GetString("JwtSecret")
	authenticator := auth.NewJwtAuthenticator(jwt)

	app := &api.Application{
		Store:         storage,
		Authenticator: authenticator,
	}

	mux := app.Start()
	if err := app.Run(mux); err != nil {
		log.Fatal().Err(err).Msg("server failed to start")
	}*/
}
