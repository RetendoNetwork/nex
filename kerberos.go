package nex

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rc4"
	"errors"
)

// EncryptionHandler provides methods for encryption and integrity verification
type EncryptionHandler struct {
	secretKey []byte
}

// VerifyChecksum ensures the integrity of data using HMAC validation
func (eh *EncryptionHandler) VerifyChecksum(data []byte) bool {
	content := data[:len(data)-16]
	checksum := data[len(data)-16:]
	mac := hmac.New(md5.New, eh.secretKey)
	mac.Write(content)
	return hmac.Equal(checksum, mac.Sum(nil))
}

// DecryptData decrypts data if it passes the HMAC integrity check
func (eh *EncryptionHandler) DecryptData(data []byte) ([]byte, error) {
	if !eh.VerifyChecksum(data) {
		return nil, errors.New("data integrity validation failed: invalid checksum")
	}
	cipher, err := rc4.NewCipher(eh.secretKey)
	if err != nil {
		return nil, err
	}
	decryptedData := make([]byte, len(data)-16)
	cipher.XORKeyStream(decryptedData, data[:len(data)-16])
	return decryptedData, nil
}

// EncryptData encrypts data and appends an HMAC checksum for validation
func (eh *EncryptionHandler) EncryptData(data []byte) []byte {
	cipher, _ := rc4.NewCipher(eh.secretKey)
	encryptedData := make([]byte, len(data))
	cipher.XORKeyStream(encryptedData, data)
	mac := hmac.New(md5.New, eh.secretKey)
	mac.Write(encryptedData)
	checksum := mac.Sum(nil)
	return append(encryptedData, checksum...)
}

// NewEncryptionHandler initializes a new encryption handler instance
func NewEncryptionHandler(secret []byte) *EncryptionHandler {
	return &EncryptionHandler{secretKey: secret}
}

type Ticket struct {
	SessionKey []byte
	Target     int
	Internal   int
}

func (t *Ticket) EncryptData(secretKey []byte) []byte {
	encryption := NewEncryptionHandler(secretKey)
	data := append(t.SessionKey, byte(t.Target), byte(t.Internal))
	return encryption.EncryptData(data)
}

func NewTicket() *Ticket {
	return &Ticket{}
}

// GenerateEncryptionKey creates an encryption key based on user ID and password
func GenerateEncryptionKey(user int, password []byte) []byte {
	iterations := int(65000 + user%1024)
	currentKey := password
	hash := make([]byte, md5.Size)

	for i := 0; i < iterations; i++ {
		sum := md5.Sum(currentKey)
		copy(hash, sum[:])
		currentKey = hash
	}

	return currentKey
}

// NOT FINISHED YET
