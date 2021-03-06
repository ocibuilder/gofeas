/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gofeas

// DeploymentPlatform : Types of platforms.   - PLATFORM_UNSPECIFIED: Unknown.  - GKE: Google Container Engine.  - FLEX: Google App Engine: Flexible Environment.  - CUSTOM: Custom user-defined platform.
type DeploymentPlatform string

// List of DeploymentPlatform
const (
	PLATFORM_UNSPECIFIED_DeploymentPlatform DeploymentPlatform = "PLATFORM_UNSPECIFIED"
	GKE_DeploymentPlatform                  DeploymentPlatform = "GKE"
	FLEX_DeploymentPlatform                 DeploymentPlatform = "FLEX"
	CUSTOM_DeploymentPlatform               DeploymentPlatform = "CUSTOM"
)
