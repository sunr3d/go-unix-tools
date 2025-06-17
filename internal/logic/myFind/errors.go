package myFind

const (
	errInvalidPath    = "указанный путь %s не является директорией"
	errFileInfo       = "ошибка получения информации о файле %s: %w"
	errWalkDir        = "ошибка при обходе директории %s: %w"
	errExtWithoutFile = "флаг -ext требует флаг -f"
	errInvalidExt     = "расширение должно состоять только из букв и цифр - %s"
	errSinglePath     = "необходимо указать только одну директорию"
	errFlagParse      = "ошибка парсинга набора флагов: %w"
)
