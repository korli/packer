/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-10-18T06:16:13Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type DetachBlockStorageInstancesRequest struct {

	// 블록스토리지인스턴스번호리스트
	BlockStorageInstanceNoList []*string `json:"blockStorageInstanceNoList"`
}
