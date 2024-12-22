package db

type dbImplementation struct {
	// если говорить о базах данных
	// то тут лежит соединение с базой
	// redis.Client или gorm.DB
	connection map[string]string
}

func (db *dbImplementation) Get(
	id string,
) (string, error) {
	name, ok := db.connection[id]
	if !ok {
		return "", ErrNotFound
	}

	return name, nil
}

func (db *dbImplementation) Set(
	id, name string,
) error {
	if _, ok := db.connection[id]; ok {
		return ErrAlreadyExist
	}

	db.connection[id] = name

	return nil
}

func New() *dbImplementation {
	return &dbImplementation{
		connection: make(map[string]string),
	}
}
