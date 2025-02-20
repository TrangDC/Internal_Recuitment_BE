package dto

import (
	"trec/models"

	"github.com/golang-jwt/jwt"
)

type AzureDto interface {
	ConvertTokenClam(claims jwt.MapClaims) *models.Token
}

type azureDtoImpl struct {
}

func NewAzureDto() AzureDto {
	return &azureDtoImpl{}
}

func (d azureDtoImpl) ConvertTokenClam(claims jwt.MapClaims) *models.Token {
	return &models.Token{
		Audience:           claims["aud"].(string),
		Issuer:             claims["iss"].(string),
		IssuedAt:           int64(claims["iat"].(float64)),
		NotBefore:          int64(claims["nbf"].(float64)),
		Expiration:         int64(claims["exp"].(float64)),
		AuthenticationInfo: claims["aio"].(string),
		AuthorizedParty:    claims["azp"].(string),
		AuthorizedPartyCR:  claims["azpacr"].(string),
		Name:               claims["name"].(string),
		ObjectID:           claims["oid"].(string),
		PreferredUsername:  claims["preferred_username"].(string),
		RefreshTokenHash:   claims["rh"].(string),
		Scopes:             claims["scp"].(string),
		Subject:            claims["sub"].(string),
		TenantID:           claims["tid"].(string),
		UniqueTokenID:      claims["uti"].(string),
		Version:            claims["ver"].(string),
	}
}
