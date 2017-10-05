package cynomys

import (
	"context"
	"errors"

	"cloud.google.com/go/pubsub"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

const (
	// OAuthScope is ...
	OAuthScope = pubsub.ScopePubSub
	// GCPCredentialTypeDefault is ...
	GCPCredentialTypeDefault = "default"
)

// PerRPCCredential is ...
type PerRPCCredential credentials.PerRPCCredentials

// TLSCredential is ...
type TLSCredential credentials.TransportCredentials

// NewPerRPCCredential is ...
func NewPerRPCCredential(ctx context.Context, credentialType string) (PerRPCCredential, error) {
	var c PerRPCCredential
	switch credentialType {
	case GCPCredentialTypeDefault:
		c, err := oauth.NewApplicationDefault(ctx, OAuthScope)
		if err != nil {
			return c, err
		}
	default:
		return c, errors.New("auth method is not implemented")
	}
	return c, nil
}

// NewTLSCredential is ...
func NewTLSCredential() TLSCredential {
	return credentials.NewTLS(nil)
}
