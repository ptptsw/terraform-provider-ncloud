/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-28T05:05:08Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type RecreateServerInstanceRequest struct {

	// 서버인스턴스번호
ServerInstanceNo *string `json:"serverInstanceNo,omitempty"`

	// 서버인스턴스이름
ServerInstanceName *string `json:"serverInstanceName,omitempty"`

	// 서버이미지상품코드
ServerImageProductCode *string `json:"serverImageProductCode,omitempty"`

	// 사용자데이터
UserData *string `json:"userData,omitempty"`

	// 인스턴스태그리스트
InstanceTagList []*InstanceTagParameter `json:"instanceTagList,omitempty"`
}