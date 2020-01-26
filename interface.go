package gofeas

import (
	"context"
	"net/http"
)

type CommonAPIClient interface {
	BatchCreateNotes(ctx context.Context, parent string, body V1beta1BatchCreateNotesRequest) (V1beta1BatchCreateNotesResponse, *http.Response, error)
	BatchCreateOccurrences(ctx context.Context, parent string, body V1beta1BatchCreateOccurrencesRequest) (V1beta1BatchCreateOccurrencesResponse, *http.Response, error)
	CreateNote(ctx context.Context, parent string, body V1beta1Note) (V1beta1Note, *http.Response, error)
	CreateOccurrence(ctx context.Context, parent string, body V1beta1Occurrence) (V1beta1Occurrence, *http.Response, error)
	DeleteNote(ctx context.Context, name string) (interface{}, *http.Response, error)
	DeleteOccurrence(ctx context.Context, name string) (interface{}, *http.Response, error)
	GetNote(ctx context.Context, name string) (V1beta1Note, *http.Response, error)
	GetOccurrence(ctx context.Context, name string) (V1beta1Occurrence, *http.Response, error)
	GetOccurrenceNote(ctx context.Context, name string) (V1beta1Note, *http.Response, error)
	GetVulnerabilityOccurrencesSummary(ctx context.Context, parent string, localVarOptionals *GetVulnerabilityOccurrencesSummaryOpts) (V1beta1VulnerabilityOccurrencesSummary, *http.Response, error)
	ListNoteOccurrences(ctx context.Context, name string, localVarOptionals *ListNoteOccurrencesOpts) (V1beta1ListNoteOccurrencesResponse, *http.Response, error)
	ListNotes(ctx context.Context, parent string, localVarOptionals *ListNotesOpts) (V1beta1ListNotesResponse, *http.Response, error)
	ListOccurrences(ctx context.Context, parent string, localVarOptionals *ListOccurrencesOpts) (V1beta1ListOccurrencesResponse, *http.Response, error)
	UpdateNote(ctx context.Context, name string, body V1beta1Note) (V1beta1Note, *http.Response, error)
	UpdateOccurrence(ctx context.Context, name string, body V1beta1Occurrence) (V1beta1Occurrence, *http.Response, error)
}
