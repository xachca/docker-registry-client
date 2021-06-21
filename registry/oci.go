package registry

const (
	OCIContentDescriptor           string = "application/vnd.oci.descriptor.v1+json"
	OCIImageIndex                  string = "application/vnd.oci.image.index.v1+json"
	OCIManifestSchema1             string = "application/vnd.oci.image.manifest.v1+json"
	OCIConfigJSON                  string = "application/vnd.oci.image.config.v1+json"
	OCILayer                       string = "application/vnd.oci.image.layer.v1.tar+gzip"
	OCIRestrictedLayer             string = "application/vnd.oci.image.layer.nondistributable.v1.tar+gzip"
	OCIUncompressedLayer           string = "application/vnd.oci.image.layer.v1.tar"
	OCIUncompressedRestrictedLayer string = "application/vnd.oci.image.layer.nondistributable.v1.tar"
)
