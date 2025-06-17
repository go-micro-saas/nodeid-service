package global

const (
	KeyPrefix = "nid_s_"
)

func Key(k string) string {
	return KeyPrefix + k
}
