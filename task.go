package crosschain

import "fmt"

// Task represents a tx, e.g. smart contract function call, on a blockchain.
type Task string

// TaskConfig is the model used to represent a task read from config file or db
type TaskConfig struct {
	Name  string   `yaml:"name"`
	Code  string   `yaml:"code"`
	Chain string   `yaml:"chain"`
	Allow []string `yaml:"allow"`
	// Contract   string                `yaml:"contract"`
	Operations []TaskConfigOperation `yaml:"operations"`

	// internal
	AllowList []*AllowEntry `yaml:"-"`
	SrcAsset  ITask         `yaml:"-"`
	DstAsset  ITask         `yaml:"-"`
}

// PipelineConfig is the model used to represent a pipeline (list of tasks) read from config file or db
type PipelineConfig struct {
	ID    string   `yaml:"name"`
	Allow []string `yaml:"allow"`
	Tasks []string `yaml:"tasks"`

	// internal
	AllowList []*AllowEntry `yaml:"-"`
}

func (p PipelineConfig) String() string {
	return fmt.Sprintf(
		"PipelineConfig(id=%s)",
		p.ID,
	)
}

type AllowEntry struct {
	Src AssetID
	Dst AssetID
}

type TaskConfigOperation struct {
	Function  string                     `yaml:"function"`
	Signature string                     `yaml:"signature"`
	Contract  string                     `yaml:"contract"`
	Payable   bool                       `yaml:"payable"`
	Params    []TaskConfigOperationParam `yaml:"params"`
}

type TaskConfigOperationParam struct {
	Name     string                             `yaml:"name"`
	Type     string                             `yaml:"type"`
	Bind     string                             `yaml:"bind"`
	Defaults []TaskConfigOperationParamDefaults `yaml:"defaults"`
	// Fields   []TaskConfigOperationParamField    `yaml:"fields"`
}

type TaskConfigOperationParamDefaults struct {
	Match string `yaml:"match"`
	Value string `yaml:"value"`
}

type ITask interface {
	ID() AssetID
	GetDriver() string
	GetAssetConfig() *AssetConfig
	GetNativeAsset() *NativeAssetConfig
	GetTask() *TaskConfig
}

func (task TaskConfig) String() string {
	src := "not-set"
	if task.SrcAsset != nil {
		src = string(task.SrcAsset.ID())
	}
	dst := "not-set"
	if task.DstAsset != nil {
		dst = string(task.DstAsset.ID())
	}
	return fmt.Sprintf(
		"TaskConfig(id=%s src=%s dst=%s)",
		task.ID(), src, dst,
	)
}

func (task *TaskConfig) ID() AssetID {
	return AssetID(task.Name)
}

func (task TaskConfig) GetAssetConfig() *AssetConfig {
	return task.SrcAsset.GetAssetConfig()
}

func (task TaskConfig) GetDriver() string {
	return task.SrcAsset.GetAssetConfig().Driver
}

func (task TaskConfig) GetAsset() string {
	return task.SrcAsset.GetAssetConfig().Asset
}

func (task TaskConfig) GetNativeAsset() *NativeAssetConfig {
	return task.SrcAsset.GetNativeAsset()
}

func (task TaskConfig) GetTask() *TaskConfig {
	return &task
}
