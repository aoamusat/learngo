// Code generated by goa v3.16.1, DO NOT EDIT.
//
// signin HTTP server encoders and decoders
//
// Command:
// $ goa gen learngo.io/firstgoa/design

package server

import (
	"context"
	"errors"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	signin "learngo.io/firstgoa/gen/signin"
)

// EncodeAuthenticateResponse returns an encoder for responses returned by the
// signin authenticate endpoint.
func EncodeAuthenticateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*signin.Creds)
		enc := encoder(ctx, w)
		body := NewAuthenticateResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeAuthenticateRequest returns a decoder for requests sent to the signin
// authenticate endpoint.
func DecodeAuthenticateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (any, error) {
	return func(r *http.Request) (any, error) {
		payload := NewAuthenticatePayload()
		user, pass, ok := r.BasicAuth()
		if !ok {
			return nil, goa.MissingFieldError("Authorization", "header")
		}
		payload.Username = user
		payload.Password = pass

		return payload, nil
	}
}

// EncodeAuthenticateError returns an encoder for errors returned by the
// authenticate signin endpoint.
func EncodeAuthenticateError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(ctx context.Context, err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en goa.GoaErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.GoaErrorName() {
		case "unauthorized":
			var res signin.Unauthorized
			errors.As(v, &res)
			enc := encoder(ctx, w)
			body := res
			w.Header().Set("goa-error", res.GoaErrorName())
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}