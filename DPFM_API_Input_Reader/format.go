package dpfm_api_input_reader

import (
	"data-platform-api-address-exconf-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToAddress() *requests.Address {
	data := sdc.Address
	return &requests.Address{
		AddressID:       data.AddressID,
		ValidityEndDate: data.ValidityEndDate,
	}
}
