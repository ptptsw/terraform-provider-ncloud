/*
 * cdn
 *
 * <br/>https://ncloud.apigw.ntruss.com/cdn/v2
 *
 * API version: 2018-08-07T06:43:44Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package cdn

type GlobalCdnPurgeHistory struct {

	// CDN인스턴스번호
CdnInstanceNo *string `json:"cdnInstanceNo,omitempty"`

	// 퍼지ID
PurgeId *string `json:"purgeId,omitempty"`

	// 전체퍼지여부
IsWholePurge *bool `json:"isWholePurge,omitempty"`

	// 전체도메인퍼지여부
IsWholeDomain *bool `json:"isWholeDomain,omitempty"`

	// Global CDN서비스도메인리스트
GlobalCdnServiceDomainList []*GlobalCdnServiceDomain `json:"globalCdnServiceDomainList,omitempty"`

	// 타겟파일리스트
TargetFileList []*string `json:"targetFileList,omitempty"`

	// 예상완료날짜
EstimatedCompletionDate *string `json:"estimatedCompletionDate,omitempty"`

	// 성공여부
IsSuccess *bool `json:"isSuccess,omitempty"`

	// 요청날짜
RequestDate *string `json:"requestDate,omitempty"`
}