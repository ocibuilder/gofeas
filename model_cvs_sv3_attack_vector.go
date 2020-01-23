/*
 * proto/v1beta1/grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gofeas

type CvsSv3AttackVector string

// List of CVSSv3AttackVector
const (
	UNSPECIFIED_CvsSv3AttackVector CvsSv3AttackVector = "ATTACK_VECTOR_UNSPECIFIED"
	NETWORK_CvsSv3AttackVector     CvsSv3AttackVector = "ATTACK_VECTOR_NETWORK"
	ADJACENT_CvsSv3AttackVector    CvsSv3AttackVector = "ATTACK_VECTOR_ADJACENT"
	LOCAL_CvsSv3AttackVector       CvsSv3AttackVector = "ATTACK_VECTOR_LOCAL"
	PHYSICAL_CvsSv3AttackVector    CvsSv3AttackVector = "ATTACK_VECTOR_PHYSICAL"
)
