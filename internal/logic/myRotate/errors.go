package myRotate

const (
	errNoFiles     = "не были указаны файлы для ротации"
	errArchiveDir  = "ошибка при проверке директории архивации %s: %w"
	errCwd         = "не удалось получить текущую рабочую директорию: %w"
	errParseFlags  = "ошибка при разборе флагов: %w"
	errNotALogFile = "файл %s не является лог-файлом (.log)"
	errFileInfo    = "ошибка получения информации о файле %s: %w"
	errFileAccess  = "не удалось получить доступ к файлу %s: %w"
	errArchiveFile = "не удалось архивировать файл %s: %w"
	errCreateTar   = "не удалось создать tar-файл для файла %s: %w"
	errTarHeader   = "не удалось записать заголовок tar-файла для %s: %w"
)
