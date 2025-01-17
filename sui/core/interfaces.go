package core

// SUIs the loaded SUI instances
var SUIs = map[string]SUI{}

// SUI is the interface for the SUI
type SUI interface {
	GetTemplates() ([]ITemplate, error)
	GetTemplate(name string) (ITemplate, error)
	UploadTemplate(src string, dst string) (ITemplate, error)
}

// ITemplate is the interface for the ITemplate
type ITemplate interface {
	Pages() ([]IPage, error)
	PageTree(route string) ([]*PageTreeNode, error)
	Page(route string) (IPage, error)
	GetPageFromAsset(asset string) (IPage, error)

	Blocks() ([]IBlock, error)
	Block(name string) (IBlock, error)

	Components() ([]IComponent, error)
	Component(name string) (IComponent, error)

	Assets() []string
	Locales() []SelectOption
	Themes() []SelectOption

	Asset(file string) (*Asset, error)
}

// IPage is the interface for the page
type IPage interface {
	Load() error

	Get() *Page

	EditorRender(request *Request) (*ResponseEditor, error)
	EditorPageSource() ResponseSource
	EditorScriptSource() ResponseSource
	EditorStyleSource() ResponseSource
	EditorDataSource() ResponseSource

	AssetScript() (*Asset, error)
	AssetStyle() (*Asset, error)

	// Render()

	// Html()
	// Script()
	// Style()
	// Data()

	// Compile()
	// Locale()
}

// IBlock is the interface for the block
type IBlock interface {
	Compile() (string, error)
	Load() error
	Source() string
}

// IComponent is the interface for the component
type IComponent interface {
	Compile() (string, error)
	Load() error
	Source() string
}
