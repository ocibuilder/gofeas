/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gofeas

// PackageArchitecture : Instruction set architectures supported by various package managers.   - ARCHITECTURE_UNSPECIFIED: Unknown architecture.  - X86: X86 architecture.  - X64: X64 architecture.
type PackageArchitecture string

// List of packageArchitecture
const (
	ARCHITECTURE_UNSPECIFIED_PackageArchitecture PackageArchitecture = "ARCHITECTURE_UNSPECIFIED"
	X86_PackageArchitecture                      PackageArchitecture = "X86"
	X64_PackageArchitecture                      PackageArchitecture = "X64"
)
