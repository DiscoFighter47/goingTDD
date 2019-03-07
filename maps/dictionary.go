package maps

// Dictionary representes a key-word pair storage
type Dictionary map[string]string

// Search finds a value from the dictionary for a given key
func (d Dictionary) Search(key string) (string, error) {
	word, found := d[key]
	if !found {
		return "", ErrKeyNotFound
	}
	return word, nil
}

// Add adds a new key-word pair in the dictionary
func (d Dictionary) Add(key, word string) error {
	_, err := d.Search(key)
	if err == nil {
		return ErrDuplicateKey
	}
	d[key] = word
	return nil
}

// Update updates word for a key in the dictionary
func (d Dictionary) Update(key, word string) error {
	_, err := d.Search(key)
	if err != nil {
		return err
	}
	d[key] = word
	return nil
}

// Delete deletes a key form the dictionary
func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	if err != nil {
		return err
	}
	delete(d, key)
	return nil
}
