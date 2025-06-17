package myWc

const (
	errMultipleFlags = "за раз можно использовать только один из флагов"
	errNoFiles       = "не указан файл для обработки или файл не найден"
	errFileOpen      = "ошибка открытия файла %s: %w"
	errFileProcess   = "ошибка обработки файла %s: %w"
	errFlagParse     = "ошибка парсинга набора флагов: %w"
)
