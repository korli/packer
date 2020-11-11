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

// FiltersNet One or more filters.
type FiltersNet struct {
	// The IDs of the DHCP options sets.
	DhcpOptionsSetIds []string `json:"DhcpOptionsSetIds,omitempty"`
	// The IP ranges for the Nets, in CIDR notation (for example, 10.0.0.0/16).
	IpRanges []string `json:"IpRanges,omitempty"`
	// If `true`, the Net used is the default one.
	IsDefault bool `json:"IsDefault,omitempty"`
	// The IDs of the Nets.
	NetIds []string `json:"NetIds,omitempty"`
	// The states of the Nets (`pending` \\| `available`).
	States []string `json:"States,omitempty"`
	// The keys of the tags associated with the Nets.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the Nets.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the Nets, in the following format: \"Filters\":{\"Tags\":[\"TAGKEY=TAGVALUE\"]}.
	Tags []string `json:"Tags,omitempty"`
}
