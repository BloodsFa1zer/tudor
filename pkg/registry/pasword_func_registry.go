package registry

import "golang.org/x/crypto/bcrypt"

func hashPasFunc() func(password string) string {
	return func(password string) string {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		return string(hashedPassword)
	}
}

func comparePasFunc() func(hashedPassword string, candidatePassword string) error {
	return func(hashedPassword string, candidatePassword string) error {
		return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
	}
}
