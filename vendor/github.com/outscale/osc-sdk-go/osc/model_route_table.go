/*
 * 3DS OUTSCALE API
 *
 * Welcome to the 3DS OUTSCALE's API documentation.<br /><br />  The 3DS OUTSCALE API enables you to manage your resources in the 3DS OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the 3DS OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the 3DS OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.2
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// RouteTable Information about the route table.
type RouteTable struct {
	// One or more associations between the route table and Subnets.
	LinkRouteTables []LinkRouteTable `json:"LinkRouteTables,omitempty"`
	// The ID of the Net for the route table.
	NetId string `json:"NetId,omitempty"`
	// Information about virtual gateways propagating routes.
	RoutePropagatingVirtualGateways []RoutePropagatingVirtualGateway `json:"RoutePropagatingVirtualGateways,omitempty"`
	// The ID of the route table.
	RouteTableId string `json:"RouteTableId,omitempty"`
	// One or more routes in the route table.
	Routes []Route `json:"Routes,omitempty"`
	// One or more tags associated with the route table.
	Tags []ResourceTag `json:"Tags,omitempty"`
}
