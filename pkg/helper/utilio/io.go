package utilio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFileBuffered(filePath string, bufferSize int) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReaderSize(file, bufferSize)
	var sb strings.Builder
	buf := make([]byte, bufferSize)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			sb.Write(buf[:n])
		}
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return "", fmt.Errorf("error reading file: %v", err)
		}
	}
	return sb.String(), nil
}
