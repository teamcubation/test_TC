package authe

import (
	"context"
	"errors"
	"time"

	redis0 "github.com/go-redis/redis/v8"

	redis "github.com/teamcubation/teamcandidates/pkg/databases/cache/redis/v8"
	types "github.com/teamcubation/teamcandidates/pkg/types"

	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/authe/redis/dto"
	"github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/authe/usecases/domain"
)

type cache struct {
	cache redis.Cache
}

func NewRedisCache(c redis.Cache) Cache {
	return &cache{
		cache: c,
	}
}

func (c *cache) StoreToken(ctx context.Context, userID string, token *domain.Token) error {
	data, err := dto.FromDomainToJSON(token)
	if err != nil {
		return err
	}

	expiration := time.Until(token.AccessExpiresAt)
	return c.cache.Set(ctx, userID, data, expiration)
}

func (c *cache) RetrieveToken(ctx context.Context, userID string) (*domain.Token, error) {
	data, err := c.cache.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, redis0.Nil) {
			return nil, types.NewError(types.ErrTokenNotFound, "token not found in cache", nil)
		}
		return nil, types.NewError(types.ErrConnection, "failed to retrieve token from cache", err)
	}

	token, parseErr := dto.FromJSONToDomain(data)
	if parseErr != nil {
		return nil, types.NewError(types.ErrInvalidInput, "failed to parse token data", parseErr)
	}

	return token, nil
}

func (c *cache) Close() {
	c.cache.Close()
}
