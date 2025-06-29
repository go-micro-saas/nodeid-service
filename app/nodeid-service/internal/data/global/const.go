package global

const (
	KeyPrefix = "nodeid_service_"
)

func Key(k string) string {
	return KeyPrefix + k
}
