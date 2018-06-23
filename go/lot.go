/*
 * SlugSpace
 *
 * SlugSpace API - Realtime Parking Data and Metrics.
 *
 * API version: 0.0.1
 * Contact: colby.leiske@gmail.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package slugspace

type Lot struct {

	Id int64 `json:"id"`

	Name string `json:"name"`

	FreeSpaces int64 `json:"freeSpaces"`

	TotalSpaces int64 `json:"totalSpaces"`

	LastUpdated string `json:"lastUpdated"`
}
