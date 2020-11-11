/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-10-18T06:16:13Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type GetPublicIpInstanceListRequest struct {

	// 할당여부
	IsAssociated *bool `json:"isAssociated,omitempty"`

	// 공인IP인스턴스번호리스트
	PublicIpInstanceNoList []*string `json:"publicIpInstanceNoList,omitempty"`

	// 공인IP리스트
	PublicIpList []*string `json:"publicIpList,omitempty"`

	// 검색할필터명
	SearchFilterName *string `json:"searchFilterName,omitempty"`

	// 검색할필터값
	SearchFilterValue *string `json:"searchFilterValue,omitempty"`

	// 인터넷라인구분코드
	InternetLineTypeCode *string `json:"internetLineTypeCode,omitempty"`

	// 리전번호
	RegionNo *string `json:"regionNo,omitempty"`

	// ZONE번호
	ZoneNo *string `json:"zoneNo,omitempty"`

	// 페이지번호
	PageNo *int32 `json:"pageNo,omitempty"`

	// 페이지사이즈
	PageSize *int32 `json:"pageSize,omitempty"`

	// 소팅대상
	SortedBy *string `json:"sortedBy,omitempty"`

	// 소팅순서
	SortingOrder *string `json:"sortingOrder,omitempty"`
}
