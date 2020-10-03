// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDcimPowerPanelsListParams creates a new DcimPowerPanelsListParams object
// with the default values initialized.
func NewDcimPowerPanelsListParams() *DcimPowerPanelsListParams {
	var ()
	return &DcimPowerPanelsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPanelsListParamsWithTimeout creates a new DcimPowerPanelsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerPanelsListParamsWithTimeout(timeout time.Duration) *DcimPowerPanelsListParams {
	var ()
	return &DcimPowerPanelsListParams{

		timeout: timeout,
	}
}

// NewDcimPowerPanelsListParamsWithContext creates a new DcimPowerPanelsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerPanelsListParamsWithContext(ctx context.Context) *DcimPowerPanelsListParams {
	var ()
	return &DcimPowerPanelsListParams{

		Context: ctx,
	}
}

// NewDcimPowerPanelsListParamsWithHTTPClient creates a new DcimPowerPanelsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerPanelsListParamsWithHTTPClient(client *http.Client) *DcimPowerPanelsListParams {
	var ()
	return &DcimPowerPanelsListParams{
		HTTPClient: client,
	}
}

/*DcimPowerPanelsListParams contains all the parameters to send to the API endpoint
for the dcim power panels list operation typically these are written to a http.Request
*/
type DcimPowerPanelsListParams struct {

	/*ID*/
	ID *string
	/*IDGt*/
	IDGt *string
	/*IDGte*/
	IDGte *string
	/*IDLt*/
	IDLt *string
	/*IDLte*/
	IDLte *string
	/*IDn*/
	IDn *string
	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*NameIc*/
	NameIc *string
	/*NameIe*/
	NameIe *string
	/*NameIew*/
	NameIew *string
	/*NameIsw*/
	NameIsw *string
	/*Namen*/
	Namen *string
	/*NameNic*/
	NameNic *string
	/*NameNie*/
	NameNie *string
	/*NameNiew*/
	NameNiew *string
	/*NameNisw*/
	NameNisw *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*RackGroupID*/
	RackGroupID *string
	/*RackGroupIDn*/
	RackGroupIDn *string
	/*Region*/
	Region *string
	/*Regionn*/
	Regionn *string
	/*RegionID*/
	RegionID *string
	/*RegionIDn*/
	RegionIDn *string
	/*Site*/
	Site *string
	/*Siten*/
	Siten *string
	/*SiteID*/
	SiteID *string
	/*SiteIDn*/
	SiteIDn *string
	/*Tag*/
	Tag *string
	/*Tagn*/
	Tagn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithTimeout(timeout time.Duration) *DcimPowerPanelsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithContext(ctx context.Context) *DcimPowerPanelsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithHTTPClient(client *http.Client) *DcimPowerPanelsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithID(id *string) *DcimPowerPanelsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithIDGt(iDGt *string) *DcimPowerPanelsListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithIDGte(iDGte *string) *DcimPowerPanelsListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithIDLt(iDLt *string) *DcimPowerPanelsListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithIDLte(iDLte *string) *DcimPowerPanelsListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithIDn(iDn *string) *DcimPowerPanelsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLimit adds the limit to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithLimit(limit *int64) *DcimPowerPanelsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithName(name *string) *DcimPowerPanelsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithNameIc(nameIc *string) *DcimPowerPanelsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithNameIe(nameIe *string) *DcimPowerPanelsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithNameIew(nameIew *string) *DcimPowerPanelsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithNameIsw(nameIsw *string) *DcimPowerPanelsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithNamen(namen *string) *DcimPowerPanelsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithNameNic(nameNic *string) *DcimPowerPanelsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithNameNie(nameNie *string) *DcimPowerPanelsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithNameNiew(nameNiew *string) *DcimPowerPanelsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithNameNisw(nameNisw *string) *DcimPowerPanelsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithOffset(offset *int64) *DcimPowerPanelsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithQ(q *string) *DcimPowerPanelsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRackGroupID adds the rackGroupID to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithRackGroupID(rackGroupID *string) *DcimPowerPanelsListParams {
	o.SetRackGroupID(rackGroupID)
	return o
}

// SetRackGroupID adds the rackGroupId to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetRackGroupID(rackGroupID *string) {
	o.RackGroupID = rackGroupID
}

// WithRackGroupIDn adds the rackGroupIDn to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithRackGroupIDn(rackGroupIDn *string) *DcimPowerPanelsListParams {
	o.SetRackGroupIDn(rackGroupIDn)
	return o
}

// SetRackGroupIDn adds the rackGroupIdN to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetRackGroupIDn(rackGroupIDn *string) {
	o.RackGroupIDn = rackGroupIDn
}

// WithRegion adds the region to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithRegion(region *string) *DcimPowerPanelsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithRegionn(regionn *string) *DcimPowerPanelsListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithRegionID(regionID *string) *DcimPowerPanelsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithRegionIDn(regionIDn *string) *DcimPowerPanelsListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithSite(site *string) *DcimPowerPanelsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithSiten(siten *string) *DcimPowerPanelsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithSiteID(siteID *string) *DcimPowerPanelsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithSiteIDn(siteIDn *string) *DcimPowerPanelsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTag adds the tag to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithTag(tag *string) *DcimPowerPanelsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim power panels list params
func (o *DcimPowerPanelsListParams) WithTagn(tagn *string) *DcimPowerPanelsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim power panels list params
func (o *DcimPowerPanelsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPanelsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string
		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {
			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}

	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string
		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {
			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}

	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string
		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {
			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}

	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string
		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {
			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}

	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string
		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {
			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string
		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {
			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}

	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string
		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {
			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}

	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string
		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {
			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}

	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string
		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {
			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}

	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string
		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {
			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}

	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string
		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {
			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}

	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string
		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {
			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}

	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string
		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {
			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}

	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string
		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {
			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.RackGroupID != nil {

		// query param rack_group_id
		var qrRackGroupID string
		if o.RackGroupID != nil {
			qrRackGroupID = *o.RackGroupID
		}
		qRackGroupID := qrRackGroupID
		if qRackGroupID != "" {
			if err := r.SetQueryParam("rack_group_id", qRackGroupID); err != nil {
				return err
			}
		}

	}

	if o.RackGroupIDn != nil {

		// query param rack_group_id__n
		var qrRackGroupIDn string
		if o.RackGroupIDn != nil {
			qrRackGroupIDn = *o.RackGroupIDn
		}
		qRackGroupIDn := qrRackGroupIDn
		if qRackGroupIDn != "" {
			if err := r.SetQueryParam("rack_group_id__n", qRackGroupIDn); err != nil {
				return err
			}
		}

	}

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if o.Regionn != nil {

		// query param region__n
		var qrRegionn string
		if o.Regionn != nil {
			qrRegionn = *o.Regionn
		}
		qRegionn := qrRegionn
		if qRegionn != "" {
			if err := r.SetQueryParam("region__n", qRegionn); err != nil {
				return err
			}
		}

	}

	if o.RegionID != nil {

		// query param region_id
		var qrRegionID string
		if o.RegionID != nil {
			qrRegionID = *o.RegionID
		}
		qRegionID := qrRegionID
		if qRegionID != "" {
			if err := r.SetQueryParam("region_id", qRegionID); err != nil {
				return err
			}
		}

	}

	if o.RegionIDn != nil {

		// query param region_id__n
		var qrRegionIDn string
		if o.RegionIDn != nil {
			qrRegionIDn = *o.RegionIDn
		}
		qRegionIDn := qrRegionIDn
		if qRegionIDn != "" {
			if err := r.SetQueryParam("region_id__n", qRegionIDn); err != nil {
				return err
			}
		}

	}

	if o.Site != nil {

		// query param site
		var qrSite string
		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {
			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}

	}

	if o.Siten != nil {

		// query param site__n
		var qrSiten string
		if o.Siten != nil {
			qrSiten = *o.Siten
		}
		qSiten := qrSiten
		if qSiten != "" {
			if err := r.SetQueryParam("site__n", qSiten); err != nil {
				return err
			}
		}

	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string
		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {
			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}

	}

	if o.SiteIDn != nil {

		// query param site_id__n
		var qrSiteIDn string
		if o.SiteIDn != nil {
			qrSiteIDn = *o.SiteIDn
		}
		qSiteIDn := qrSiteIDn
		if qSiteIDn != "" {
			if err := r.SetQueryParam("site_id__n", qSiteIDn); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string
		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {
			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
