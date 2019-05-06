package bootkube

import (
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	etcdMetricSignerSecretFileName = "etcd-metric-signer-secret.yaml.template"
)

var _ asset.WritableAsset = (*EtcdMetricSignerSecret)(nil)

// EtcdMetricSignerSecret is an asset for the etcd serving signer
type EtcdMetricSignerSecret struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset
func (t *EtcdMetricSignerSecret) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *EtcdMetricSignerSecret) Name() string {
	return "EtcdMetricSignerSecret"
}

// Generate generates the actual files by this asset
func (t *EtcdMetricSignerSecret) Generate(parents asset.Parents) error {
	fileName := etcdMetricSignerSecretFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, fileName),
			Data:     []byte(data),
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *EtcdMetricSignerSecret) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *EtcdMetricSignerSecret) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, etcdMetricSignerSecretFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
