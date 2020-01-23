/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gofeas

// Request to create occurrences in batch.
type V1beta1BatchCreateOccurrencesRequest struct {
	// The name of the project in the form of `projects/[PROJECT_ID]`, under which the occurrences are to be created.
	Parent string `json:"parent,omitempty"`
	// The occurrences to create. Max allowed length is 1000.
	Occurrences []V1beta1Occurrence `json:"occurrences,omitempty"`
}
