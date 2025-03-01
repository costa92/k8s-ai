

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// PersistentVolumeSourceApplyConfiguration represents a declarative configuration of the PersistentVolumeSource type for use
// with apply.
type PersistentVolumeSourceApplyConfiguration struct {
	GCEPersistentDisk    *GCEPersistentDiskVolumeSourceApplyConfiguration    `json:"gcePersistentDisk,omitempty"`
	AWSElasticBlockStore *AWSElasticBlockStoreVolumeSourceApplyConfiguration `json:"awsElasticBlockStore,omitempty"`
	HostPath             *HostPathVolumeSourceApplyConfiguration             `json:"hostPath,omitempty"`
	Glusterfs            *GlusterfsPersistentVolumeSourceApplyConfiguration  `json:"glusterfs,omitempty"`
	NFS                  *NFSVolumeSourceApplyConfiguration                  `json:"nfs,omitempty"`
	RBD                  *RBDPersistentVolumeSourceApplyConfiguration        `json:"rbd,omitempty"`
	ISCSI                *ISCSIPersistentVolumeSourceApplyConfiguration      `json:"iscsi,omitempty"`
	Cinder               *CinderPersistentVolumeSourceApplyConfiguration     `json:"cinder,omitempty"`
	CephFS               *CephFSPersistentVolumeSourceApplyConfiguration     `json:"cephfs,omitempty"`
	FC                   *FCVolumeSourceApplyConfiguration                   `json:"fc,omitempty"`
	Flocker              *FlockerVolumeSourceApplyConfiguration              `json:"flocker,omitempty"`
	FlexVolume           *FlexPersistentVolumeSourceApplyConfiguration       `json:"flexVolume,omitempty"`
	AzureFile            *AzureFilePersistentVolumeSourceApplyConfiguration  `json:"azureFile,omitempty"`
	VsphereVolume        *VsphereVirtualDiskVolumeSourceApplyConfiguration   `json:"vsphereVolume,omitempty"`
	Quobyte              *QuobyteVolumeSourceApplyConfiguration              `json:"quobyte,omitempty"`
	AzureDisk            *AzureDiskVolumeSourceApplyConfiguration            `json:"azureDisk,omitempty"`
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSourceApplyConfiguration `json:"photonPersistentDisk,omitempty"`
	PortworxVolume       *PortworxVolumeSourceApplyConfiguration             `json:"portworxVolume,omitempty"`
	ScaleIO              *ScaleIOPersistentVolumeSourceApplyConfiguration    `json:"scaleIO,omitempty"`
	Local                *LocalVolumeSourceApplyConfiguration                `json:"local,omitempty"`
	StorageOS            *StorageOSPersistentVolumeSourceApplyConfiguration  `json:"storageos,omitempty"`
	CSI                  *CSIPersistentVolumeSourceApplyConfiguration        `json:"csi,omitempty"`
}

// PersistentVolumeSourceApplyConfiguration constructs a declarative configuration of the PersistentVolumeSource type for use with
// apply.
func PersistentVolumeSource() *PersistentVolumeSourceApplyConfiguration {
	return &PersistentVolumeSourceApplyConfiguration{}
}

// WithGCEPersistentDisk sets the GCEPersistentDisk field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GCEPersistentDisk field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithGCEPersistentDisk(value *GCEPersistentDiskVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.GCEPersistentDisk = value
	return b
}

// WithAWSElasticBlockStore sets the AWSElasticBlockStore field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AWSElasticBlockStore field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithAWSElasticBlockStore(value *AWSElasticBlockStoreVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.AWSElasticBlockStore = value
	return b
}

// WithHostPath sets the HostPath field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HostPath field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithHostPath(value *HostPathVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.HostPath = value
	return b
}

// WithGlusterfs sets the Glusterfs field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Glusterfs field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithGlusterfs(value *GlusterfsPersistentVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.Glusterfs = value
	return b
}

// WithNFS sets the NFS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NFS field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithNFS(value *NFSVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.NFS = value
	return b
}

// WithRBD sets the RBD field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RBD field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithRBD(value *RBDPersistentVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.RBD = value
	return b
}

// WithISCSI sets the ISCSI field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ISCSI field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithISCSI(value *ISCSIPersistentVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.ISCSI = value
	return b
}

// WithCinder sets the Cinder field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Cinder field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithCinder(value *CinderPersistentVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.Cinder = value
	return b
}

// WithCephFS sets the CephFS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CephFS field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithCephFS(value *CephFSPersistentVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.CephFS = value
	return b
}

// WithFC sets the FC field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FC field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithFC(value *FCVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.FC = value
	return b
}

// WithFlocker sets the Flocker field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Flocker field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithFlocker(value *FlockerVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.Flocker = value
	return b
}

// WithFlexVolume sets the FlexVolume field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FlexVolume field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithFlexVolume(value *FlexPersistentVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.FlexVolume = value
	return b
}

// WithAzureFile sets the AzureFile field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AzureFile field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithAzureFile(value *AzureFilePersistentVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.AzureFile = value
	return b
}

// WithVsphereVolume sets the VsphereVolume field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VsphereVolume field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithVsphereVolume(value *VsphereVirtualDiskVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.VsphereVolume = value
	return b
}

// WithQuobyte sets the Quobyte field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Quobyte field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithQuobyte(value *QuobyteVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.Quobyte = value
	return b
}

// WithAzureDisk sets the AzureDisk field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AzureDisk field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithAzureDisk(value *AzureDiskVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.AzureDisk = value
	return b
}

// WithPhotonPersistentDisk sets the PhotonPersistentDisk field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PhotonPersistentDisk field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithPhotonPersistentDisk(value *PhotonPersistentDiskVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.PhotonPersistentDisk = value
	return b
}

// WithPortworxVolume sets the PortworxVolume field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PortworxVolume field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithPortworxVolume(value *PortworxVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.PortworxVolume = value
	return b
}

// WithScaleIO sets the ScaleIO field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ScaleIO field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithScaleIO(value *ScaleIOPersistentVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.ScaleIO = value
	return b
}

// WithLocal sets the Local field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Local field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithLocal(value *LocalVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.Local = value
	return b
}

// WithStorageOS sets the StorageOS field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StorageOS field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithStorageOS(value *StorageOSPersistentVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.StorageOS = value
	return b
}

// WithCSI sets the CSI field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CSI field is set to the value of the last call.
func (b *PersistentVolumeSourceApplyConfiguration) WithCSI(value *CSIPersistentVolumeSourceApplyConfiguration) *PersistentVolumeSourceApplyConfiguration {
	b.CSI = value
	return b
}
