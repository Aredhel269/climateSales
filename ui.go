package main

func (app *Config) makeUI() {
	// obtenir les dades de l'API (Probabilitat de precipitacions, Temperatura Max. i Min. i Humitat)
	precipitacio, tempMax, tempMin, humitat := app.getClimaText()
	//insertar la informació dins del contenidor

	//afegir el contenidor a la finestra
}
