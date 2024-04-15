package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"time"
	"infind/internal/user"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type Provider struct {
	secret *rsa.PrivateKey
}

func NewProvider(secert string) (*Provider, error) {
	byteSecret, err := base64.StdEncoding.DecodeString(secert)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	block, _ := pem.Decode(byteSecret)
	if block == nil {
		log.Error().Msg("failed to decode JWT secret as PEM block")
		return nil, err
	}

	prePrivateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	privateKey, ok := prePrivateKey.(*rsa.PrivateKey)
	if !ok {
		log.Error().Msg("failed to assert JWT secret as RSA private key")
		return nil, err
	}

	return &Provider{
		secret: privateKey,
	}, nil
}

func (s *Provider) GenerateToken(user *user.User) (string, error) {
	claims := &Claims{
		UserID: user.UserID.String(),
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "simple-api",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(s.secret)
	if err != nil {
		log.Error().Err(err).Send()
		return "", err
	}

	return token, nil
}
