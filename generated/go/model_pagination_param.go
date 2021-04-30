/*
 * Scaffold Servise
 *
 * Base for other projects
 *
 * API version: 0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PaginationParam struct {

	Page int64 `json:"page,omitempty"`

	PageSize int64 `json:"pageSize,omitempty"`

	Order string `json:"order,omitempty"`

	Total int64 `json:"total,omitempty"`

	Records []ReturnedJson `json:"records,omitempty"`
}
