// Code generated by swagger-doc. DO NOT EDIT.

package v1alpha1

func (VirtualMachineExport) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachineExport defines the operation of exporting a VM source\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"status": "+optional",
	}
}

func (VirtualMachineExportList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "VirtualMachineExportList is a list of VirtualMachineExport resources\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "+listType=atomic",
	}
}

func (VirtualMachineExportSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":               "VirtualMachineExportSpec is the spec for a VirtualMachineExport resource",
		"tokenSecretRef": "TokenSecretRef is the name of the secret that contains the token used by the export server pod",
	}
}

func (VirtualMachineExportStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":            "VirtualMachineExportStatus is the status for a VirtualMachineExport resource",
		"phase":       "+optional",
		"links":       "+optional",
		"serviceName": "+optional\nServiceName is the name of the service created associated with the Virtual Machine export. It will be used to\ncreate the internal URLs for downloading the images",
		"conditions":  "+optional\n+listType=atomic",
	}
}

func (VirtualMachineExportLinks) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "VirtualMachineExportLinks contains the links that point the exported VM resources",
		"internal": "+optional",
		"external": "+optional",
	}
}

func (VirtualMachineExportLink) SwaggerDoc() map[string]string {
	return map[string]string{
		"":        "VirtualMachineExportLink contains a list of volumes available for export, as well as the URLs to obtain these volumes",
		"cert":    "Cert is the public CA certificate base64 encoded",
		"volumes": "Volumes is a list of available volumes to export\n+listType=atomic\n+optional",
	}
}

func (VirtualMachineExportVolume) SwaggerDoc() map[string]string {
	return map[string]string{
		"":        "VirtualMachineExportVolume contains the name and available formats for the exported volume",
		"name":    "Name is the name of the exported volume",
		"formats": "+listType=atomic\n+optional",
	}
}

func (VirtualMachineExportVolumeFormat) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "VirtualMachineExportVolumeFormat contains the format type and URL to get the volume in that format",
		"format": "Format is the format of the image at the specified URL",
		"url":    "Url is the url that contains the volume in the format specified",
	}
}

func (Condition) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                   "Condition defines conditions",
		"lastProbeTime":      "+optional\n+nullable",
		"lastTransitionTime": "+optional\n+nullable",
		"reason":             "+optional",
		"message":            "+optional",
	}
}
