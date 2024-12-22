package business_logic
// конечно пакет называею как то осмыйсленно

type Logic interface {
	Get(id string) (string, error)
	Set(id, name string) error
}
