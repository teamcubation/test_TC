package dto

import (
	jwt "github.com/teamcubation/teamcandidates/pkg/authe/jwt/v5"

	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/authe/usecases/domain"
)

func ToTokenDomain(token *jwt.Token) *domain.Token {
	return &domain.Token{
		AccessToken:      token.AccessToken,
		RefreshToken:     token.RefreshToken,
		AccessExpiresAt:  token.AccessExpiresAt,
		RefreshExpiresAt: token.RefreshExpiresAt,
		IssuedAt:         token.IssuedAt,
		Subject:          token.Subject,
		TokenType:        token.TokenType,
	}
}

func ToTokenClaimsDomain(token *jwt.TokenClaims) *domain.TokenClaims {
	return &domain.TokenClaims{
		Subject:   token.Subject,
		ExpiresAt: token.ExpiresAt,
		IssuedAt:  token.IssuedAt,
	}
}
