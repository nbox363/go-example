package business_logic

import "example/internal/db"

// как правило структуры не назвыют Implementation
// так в качестве примера
type logicImplementation struct {
	db db.DB

	// тут могут быть другие зависимости я не стал усложнять
	// например зависимость от пакета прометеуса,
	// где будут скрыты http запросы или вызовы либы
}

func (l *logicImplementation) Get(
	id string,
) (string, error) {
	// вызовы других зависимостей

	// тут конечно нужно проверить на ошибку и обернуть ее
	return l.db.Get(id)
}
func (l *logicImplementation) Set(
	id, name string,
) error {
	return l.db.Set(id, name)
}

func New(
	db db.DB,
) *logicImplementation {
	return &logicImplementation{
		db: db,
	}
}
