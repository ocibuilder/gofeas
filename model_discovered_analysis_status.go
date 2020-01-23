/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gofeas

// DiscoveredAnalysisStatus : Analysis status for a resource. Currently for initial analysis only (not updated in continuous analysis).   - ANALYSIS_STATUS_UNSPECIFIED: Unknown.  - PENDING: Resource is known but no action has been taken yet.  - SCANNING: Resource is being analyzed.  - FINISHED_SUCCESS: Analysis has finished successfully.  - FINISHED_FAILED: Analysis has finished unsuccessfully, the analysis itself is in a bad state.  - FINISHED_UNSUPPORTED: The resource is known not to be supported
type DiscoveredAnalysisStatus string

// List of DiscoveredAnalysisStatus
const (
	ANALYSIS_STATUS_UNSPECIFIED_DiscoveredAnalysisStatus DiscoveredAnalysisStatus = "ANALYSIS_STATUS_UNSPECIFIED"
	PENDING_DiscoveredAnalysisStatus                     DiscoveredAnalysisStatus = "PENDING"
	SCANNING_DiscoveredAnalysisStatus                    DiscoveredAnalysisStatus = "SCANNING"
	FINISHED_SUCCESS_DiscoveredAnalysisStatus            DiscoveredAnalysisStatus = "FINISHED_SUCCESS"
	FINISHED_FAILED_DiscoveredAnalysisStatus             DiscoveredAnalysisStatus = "FINISHED_FAILED"
	FINISHED_UNSUPPORTED_DiscoveredAnalysisStatus        DiscoveredAnalysisStatus = "FINISHED_UNSUPPORTED"
)
