package decryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RunDecryptor decrypts files in the directory with the provided hex key.
func RunDecryptor(targetDir, keyString string) {
	globalKey, err := hex.DecodeString(keyString)
	if err != nil {
		fmt.Println("Error decoding key:", err)
		return
	}
	const encryptedExt = ".thunderkitty.encrypted"

	if err := decryptDir(targetDir, globalKey, encryptedExt); err != nil {
		fmt.Println("Decryption error:", err)
	}
}

// decryptDir recursively decrypts files with the specified extension in the directory using the provided key.
func decryptDir(directory string, globalKey []byte, encryptedExt string) error {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return fmt.Errorf("reading directory %s: %w", directory, err)
	}

	for _, entry := range entries {
		path := filepath.Join(directory, entry.Name())
		if entry.IsDir() {
			if err := decryptDir(path, globalKey, encryptedExt); err != nil {
				return err
			}
		} else if strings.HasSuffix(entry.Name(), encryptedExt) {
			if err := decryptFile(path, globalKey, encryptedExt); err != nil {
				return err
			}
		}
	}
	return nil
}

// decryptFile decrypts a single file encrypted with AES-GCM using the provided key.
func decryptFile(filePath string, globalKey []byte, encryptedExt string) error {
	encryptedData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("reading encrypted file %s: %w", filePath, err)
	}

	block, err := aes.NewCipher(globalKey)
	if err != nil {
		return fmt.Errorf("creating cipher: %w", err)
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("creating AES GCM: %w", err)
	}

	nonceSize := aesgcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return errors.New("invalid encrypted data: too short")
	}

	nonce := encryptedData[:nonceSize]
	ciphertext := encryptedData[nonceSize:]

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("decrypting file %s: %w", filePath, err)
	}

	originalFilePath := strings.TrimSuffix(filePath, encryptedExt)
	if err := os.WriteFile(originalFilePath, plaintext, 0644); err != nil {
		return fmt.Errorf("writing decrypted file %s: %w", originalFilePath, err)
	}

	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("removing encrypted file %s: %w", filePath, err)
	}

	fmt.Println("Decrypted file:", originalFilePath)
	return nil
}