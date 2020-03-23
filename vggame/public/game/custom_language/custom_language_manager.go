package icuslan

type ICustomLanguageManager interface {
	GetLanguageValue(key string, lan int32) string
	AddLanguageValue(key string, mapLanVal map[int32]string)
}
