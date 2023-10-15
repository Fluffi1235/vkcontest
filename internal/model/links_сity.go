package model

const (
	Spb            string = "https://yandex.ru/pogoda/ru-RU/details?lat=59.938951&lon=30.315635&via=ms"
	Mosckow        string = "https://yandex.ru/pogoda/ru-RU/moscow/details?lat=55.755863&lon=37.6177"
	Novosibirsk    string = "https://yandex.ru/pogoda/ru-RU/details?lat=55.03020096&lon=82.92043304&via=ms"
	Ekb            string = "https://yandex.ru/pogoda/ru-RU/details?lat=56.8380127&lon=60.59747314&via=ms"
	Kazan          string = "https://yandex.ru/pogoda/ru-RU/details?lat=55.79612732&lon=49.10641479&via=ms"
	Samara         string = "https://yandex.ru/pogoda/ru-RU/details?lat=53.19587708&lon=50.10020065&via=ms"
	NizhnyNovgorod string = "https://yandex.ru/pogoda/ru-RU/details?lat=56.32679749&lon=44.00651932&via=ms"
	Rostov         string = "https://yandex.ru/pogoda/ru-RU/details?lat=47.22208023&lon=39.72035599&via=ms"
	Ufa            string = "https://yandex.ru/pogoda/ru-RU/details?lat=54.73514938&lon=55.9587326&via=ms"
	Zheleznogorsk  string = "https://yandex.ru/pogoda/ru-RU/details?lat=52.33920288&lon=35.35087204&via=ms"
)

func Cities() map[string]string {
	cities := make(map[string]string)
	cities["санкт-петербург"] = Spb
	cities["москва"] = Mosckow
	cities["новосибирск"] = Novosibirsk
	cities["екатеренбург"] = Ekb
	cities["казань"] = Kazan
	cities["самара"] = Samara
	cities["нижний новгород"] = NizhnyNovgorod
	cities["ростов"] = Rostov
	cities["уфа"] = Ufa
	cities["железногорск"] = Zheleznogorsk
	return cities
}
