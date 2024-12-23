package po

const (
	KeyPrefix = "nid_"
)

func Key(k string) string {
	return KeyPrefix + k
}
