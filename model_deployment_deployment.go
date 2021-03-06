/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gofeas

import (
	"time"
)

// The period during which some deployable was active in a runtime.
type DeploymentDeployment struct {
	// Identity of the user that triggered this deployment.
	UserEmail string `json:"user_email,omitempty"`
	// Required. Beginning of the lifetime of this deployment.
	DeployTime time.Time `json:"deploy_time,omitempty"`
	// End of the lifetime of this deployment.
	UndeployTime time.Time `json:"undeploy_time,omitempty"`
	// Configuration used to create this deployment.
	Config string `json:"config,omitempty"`
	// Address of the runtime element hosting this deployment.
	Address string `json:"address,omitempty"`
	// Output only. Resource URI for the artifact being deployed taken from the deployable field with the same name.
	ResourceUri []string `json:"resource_uri,omitempty"`
	// Platform hosting this deployment.
	Platform *DeploymentPlatform `json:"platform,omitempty"`
}
