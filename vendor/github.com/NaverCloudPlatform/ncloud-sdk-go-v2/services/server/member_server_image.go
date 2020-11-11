/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-10-18T06:16:13Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type MemberServerImage struct {

	// 회원서버이미지번호
	MemberServerImageNo *string `json:"memberServerImageNo,omitempty"`

	// 회원서버이미지명
	MemberServerImageName *string `json:"memberServerImageName,omitempty"`

	// 회원서버이미지설명
	MemberServerImageDescription *string `json:"memberServerImageDescription,omitempty"`

	// 원본서버인스턴스번호
	OriginalServerInstanceNo *string `json:"originalServerInstanceNo,omitempty"`

	// 원본서버상품코드
	OriginalServerProductCode *string `json:"originalServerProductCode,omitempty"`

	// 원본서버명
	OriginalServerName *string `json:"originalServerName,omitempty"`

	// 원본서버기본블록스토리지디스크유형
	OriginalBaseBlockStorageDiskType *CommonCode `json:"originalBaseBlockStorageDiskType,omitempty"`

	// 원본서버이미지상품코드
	OriginalServerImageProductCode *string `json:"originalServerImageProductCode,omitempty"`

	// 원본OS정보
	OriginalOsInformation *string `json:"originalOsInformation,omitempty"`

	// 원본서버이미지명
	OriginalServerImageName *string `json:"originalServerImageName,omitempty"`

	// 원본서버이미지상태명
	MemberServerImageStatusName *string `json:"memberServerImageStatusName,omitempty"`

	// 원본서버이미지상태
	MemberServerImageStatus *CommonCode `json:"memberServerImageStatus,omitempty"`

	// 원본서버이미지OP
	MemberServerImageOperation *CommonCode `json:"memberServerImageOperation,omitempty"`

	// 회원서버이미지플랫폼구분
	MemberServerImagePlatformType *CommonCode `json:"memberServerImagePlatformType,omitempty"`

	// 리전
	Region *Region `json:"region,omitempty"`

	// ZONE
	Zone *Zone `json:"zone,omitempty"`

	// 생성일시
	CreateDate *string `json:"createDate,omitempty"`

	// 회원서버이미지블록스토리지인스턴스총 개수
	MemberServerImageBlockStorageTotalRows *int32 `json:"memberServerImageBlockStorageTotalRows,omitempty"`

	// 회원서버이미지총사이즈
	MemberServerImageBlockStorageTotalSize *int64 `json:"memberServerImageBlockStorageTotalSize,omitempty"`
}
