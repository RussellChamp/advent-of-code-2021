package files

import (
	"errors"
	"fmt"
	"os"
)

func CreateOrReplace(fileName string) (*os.File, error) {
	var output *os.File
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		output, err = os.Create(fileName)
		if err != nil {
			return nil, fmt.Errorf("Error creating file %s", fileName)
		}
	} else {
		err = os.Truncate(fileName, 0)
		if err != nil {
			return nil, fmt.Errorf("Error truncating file %s", fileName)
		}
		output, err = os.Open(fileName)
		if err != nil {
			return nil, fmt.Errorf("Error opening file %s", fileName)
		}
	}
	return output, nil
}
