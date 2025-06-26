package analyzer

import "fmt"

type FileError struct {
	Path string
	Err  error
}

func (e *FileError) Error() string {
	return fmt.Sprintf("Erreur fichier: %s - %v", e.Path, e.Err)
}

func (e *FileError) Unwrap() error {
	return e.Err
}

type ParseError struct {
	ID string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Erreur de parsing du log: %s", e.ID)
}
