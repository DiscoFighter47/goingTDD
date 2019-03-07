package maps

const (
	// ErrKeyNotFound means key was not found in the dictionary
	ErrKeyNotFound = DictionaryErr("key not found")
	// ErrDuplicateKey means key already exists in the dictionary
	ErrDuplicateKey = DictionaryErr("duplicate key")
)

// DictionaryErr represents errors regarding dictionary
type DictionaryErr string

func (d DictionaryErr) Error() string {
	return string(d)
}
