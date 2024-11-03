package services

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func passwordHashing(password string) (string, error) {
	log.Println("Hashing the password using Argon2 algorithm")


	// Recommended parameters for Argon2id
	p := &params{
		memory:   64 * 1024, // 64 MB
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}
	
	salt := make([]byte, p.saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		log.Println("Error generating the salt: ", err)
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	b64Salt := base64.URLEncoding.EncodeToString(salt)
	b64Hash := base64.URLEncoding.EncodeToString(hash)

	passwordHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return passwordHash, nil
}


func verifyPasswordHash(password string, passwordHash string) (bool) {
	log.Println("Verifying the password hash using Argon2 algorithm")
	
	// Extract the parameters, salt and hash from the passwordHash
	p, salt, hash, err := extractPasswordHash(passwordHash)
	if err != nil {
		log.Println("Error extracting the parameters, salt and hash from the password hash: ", err)
		return false
	}

	// Verify the password
	loginHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, uint32(len(hash)))
	return subtle.ConstantTimeCompare(hash, loginHash) == 1
}

func extractPasswordHash(passwordHash string) (p *params, salt, hash []byte, err error) {
	parts := strings.Split(passwordHash, "$")
	if len(parts) != 6 {
		return nil, nil, nil, errors.New("invalid password hash")
	}

	var version int
    _, err = fmt.Sscanf(parts[2], "v=%d", &version)
    if err != nil {
        return nil, nil, nil, err
    }
    if version != argon2.Version {
        return nil, nil, nil, errors.New("incompatible version")
    }

    p = &params{}
    _, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
    if err != nil {
        return nil, nil, nil, err
    }

	// Decode the salt and hash from base64
    salt, err = base64.URLEncoding.DecodeString(parts[4])
    if err != nil {
        return nil, nil, nil, err
    }
    p.saltLength = uint32(len(salt))

    hash, err = base64.URLEncoding.DecodeString(parts[5])
    if err != nil {
        return nil, nil, nil, err
    }
    p.keyLength = uint32(len(hash))

    return p, salt, hash, nil
}