package manifest

type chunckMap = map[string]interface{}
type manifestMap = map[string]chunckMap

type Chunck struct {
	Src            string   `json:"src"`
	File           string   `json:"file"`
	Css            []string `json:"css"`
	Assets         []string `json:"assets"`
	IsEntry        bool     `json:"isEntry"`
	IsDynamicEntry bool     `json:"isDynamicEntry"`
	Imports        []string `json:"imports"`
	DynamicImports []string `json:"dynamicImports"`
}

type ManifestData = map[string]*Chunck
