/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gofeas

// Version contains structured information about the version of a package.
type PackageVersion struct {
	// Used to correct mistakes in the version numbering scheme.
	Epoch int32 `json:"epoch,omitempty"`
	// Required only when version kind is NORMAL. The main part of the version name.
	Name string `json:"name,omitempty"`
	// The iteration of the package build from the above version.
	Revision string `json:"revision,omitempty"`
	// Required. Distinguishes between sentinel MIN/MAX versions and normal versions.
	Kind *VersionVersionKind `json:"kind,omitempty"`
}
