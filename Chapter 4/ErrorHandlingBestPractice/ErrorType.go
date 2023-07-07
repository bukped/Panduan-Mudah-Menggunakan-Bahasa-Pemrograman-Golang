func OpenFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Lakukan operasi pada file

	return nil
}
