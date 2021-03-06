// Code generated by the feature package; DO NOT EDIT.

package feature

var appMetrics = MakeBoolFlag(
	"App Metrics",
	"appMetrics",
	"Bucky, Monitoring Team",
	false,
	Permanent,
	true,
)

// AppMetrics - Send UI Telementry to Tools cluster - should always be false in OSS
func AppMetrics() BoolFlag {
	return appMetrics
}

var backendExample = MakeBoolFlag(
	"Backend Example",
	"backendExample",
	"Gavin Cabbage",
	false,
	Permanent,
	false,
)

// BackendExample - A permanent backend example boolean flag
func BackendExample() BoolFlag {
	return backendExample
}

var communityTemplates = MakeBoolFlag(
	"Community Templates",
	"communityTemplates",
	"Bucky, Johnny Steenbergen (Berg)",
	false,
	Permanent,
	true,
)

// CommunityTemplates - Replace current template uploading functionality with community driven templates
func CommunityTemplates() BoolFlag {
	return communityTemplates
}

var frontendExample = MakeIntFlag(
	"Frontend Example",
	"frontendExample",
	"Gavin Cabbage",
	42,
	Temporary,
	true,
)

// FrontendExample - A temporary frontend example integer flag
func FrontendExample() IntFlag {
	return frontendExample
}

var pushDownWindowAggregateMean = MakeBoolFlag(
	"Push Down Window Aggregate Mean",
	"pushDownWindowAggregateMean",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownWindowAggregateMean - Enable Mean variant of PushDownWindowAggregateRule and PushDownBareAggregateRule
func PushDownWindowAggregateMean() BoolFlag {
	return pushDownWindowAggregateMean
}

var groupWindowAggregateTranspose = MakeBoolFlag(
	"Group Window Aggregate Transpose",
	"groupWindowAggregateTranspose",
	"Query Team",
	false,
	Temporary,
	false,
)

// GroupWindowAggregateTranspose - Enables the GroupWindowAggregateTransposeRule for all enabled window aggregates
func GroupWindowAggregateTranspose() BoolFlag {
	return groupWindowAggregateTranspose
}

var newAuth = MakeBoolFlag(
	"New Auth Package",
	"newAuth",
	"Alirie Gray",
	false,
	Temporary,
	false,
)

// NewAuthPackage - Enables the refactored authorization api
func NewAuthPackage() BoolFlag {
	return newAuth
}

var newLabels = MakeBoolFlag(
	"New Label Package",
	"newLabels",
	"Alirie Gray",
	false,
	Temporary,
	false,
)

// NewLabelPackage - Enables the refactored labels api
func NewLabelPackage() BoolFlag {
	return newLabels
}

var hydratevars = MakeBoolFlag(
	"New Hydrate Vars Functionality",
	"hydratevars",
	"Ariel Salem / Monitoring Team",
	false,
	Temporary,
	true,
)

// NewHydrateVarsFunctionality - Enables a minimalistic variable hydration
func NewHydrateVarsFunctionality() BoolFlag {
	return hydratevars
}

var queryCacheForDashboards = MakeBoolFlag(
	"Query Cache for Dashboards UI",
	"queryCacheForDashboards",
	"Ariel Salem / Monitoring Team",
	false,
	Temporary,
	true,
)

// QueryCacheForDashboardsUi - Enables a Dashboard Cache on the uI
func QueryCacheForDashboardsUi() BoolFlag {
	return queryCacheForDashboards
}

var memoryOptimizedFill = MakeBoolFlag(
	"Memory Optimized Fill",
	"memoryOptimizedFill",
	"Query Team",
	false,
	Temporary,
	false,
)

// MemoryOptimizedFill - Enable the memory optimized fill()
func MemoryOptimizedFill() BoolFlag {
	return memoryOptimizedFill
}

var memoryOptimizedSchemaMutation = MakeBoolFlag(
	"Memory Optimized Schema Mutation",
	"memoryOptimizedSchemaMutation",
	"Query Team",
	false,
	Temporary,
	false,
)

// MemoryOptimizedSchemaMutation - Enable the memory optimized schema mutation functions
func MemoryOptimizedSchemaMutation() BoolFlag {
	return memoryOptimizedSchemaMutation
}

var simpleTaskOptionsExtraction = MakeBoolFlag(
	"Simple Task Options Extraction",
	"simpleTaskOptionsExtraction",
	"Brett Buddin",
	false,
	Temporary,
	false,
)

// SimpleTaskOptionsExtraction - Simplified task options extraction to avoid undefined functions when saving tasks
func SimpleTaskOptionsExtraction() BoolFlag {
	return simpleTaskOptionsExtraction
}

var useUserPermission = MakeBoolFlag(
	"Use User Permission",
	"useUserPermission",
	"Lyon Hill",
	false,
	Temporary,
	false,
)

// UseUserPermission - When enabled we will use the new user service permission function
func UseUserPermission() BoolFlag {
	return useUserPermission
}

var mergeFiltersRule = MakeBoolFlag(
	"Merged Filters Rule",
	"mergeFiltersRule",
	"Query Team",
	false,
	Temporary,
	false,
)

// MergedFiltersRule - Create one filter combining multiple single return statements
func MergedFiltersRule() BoolFlag {
	return mergeFiltersRule
}

var bandPlotType = MakeBoolFlag(
	"Band Plot Type",
	"bandPlotType",
	"Monitoring Team",
	false,
	Temporary,
	true,
)

// BandPlotType - Enables the creation of a band plot in Dashboards
func BandPlotType() BoolFlag {
	return bandPlotType
}

var mosaicGraphType = MakeBoolFlag(
	"Mosaic Graph Type",
	"mosaicGraphType",
	"Monitoring Team",
	false,
	Temporary,
	true,
)

// MosaicGraphType - Enables the creation of a mosaic graph in Dashboards
func MosaicGraphType() BoolFlag {
	return mosaicGraphType
}

var notebooks = MakeBoolFlag(
	"Notebooks",
	"notebooks",
	"Monitoring Team",
	false,
	Temporary,
	true,
)

// Notebooks - Determine if the notebook feature's route and navbar icon are visible to the user
func Notebooks() BoolFlag {
	return notebooks
}

var pushDownGroupAggregateMinMax = MakeBoolFlag(
	"Push Down Group Aggregate Min Max",
	"pushDownGroupAggregateMinMax",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownGroupAggregateMinMax - Enable the min and max variants of the PushDownGroupAggregate planner rule
func PushDownGroupAggregateMinMax() BoolFlag {
	return pushDownGroupAggregateMinMax
}

var orgOnlyMemberList = MakeBoolFlag(
	"Org Only Member list",
	"orgOnlyMemberList",
	"Compute Team",
	false,
	Temporary,
	false,
)

// OrgOnlyMemberList - Enforce only org members have access to view members of org related resorces
func OrgOnlyMemberList() BoolFlag {
	return orgOnlyMemberList
}

var enforceOrgDashboardLimits = MakeBoolFlag(
	"Enforce Organization Dashboard Limits",
	"enforceOrgDashboardLimits",
	"Compute Team",
	false,
	Temporary,
	false,
)

// EnforceOrganizationDashboardLimits - Enforces the default limit params for the dashboards api when orgs are set
func EnforceOrganizationDashboardLimits() BoolFlag {
	return enforceOrgDashboardLimits
}

var all = []Flag{
	appMetrics,
	backendExample,
	communityTemplates,
	frontendExample,
	pushDownWindowAggregateMean,
	groupWindowAggregateTranspose,
	newAuth,
	newLabels,
	hydratevars,
	queryCacheForDashboards,
	memoryOptimizedFill,
	memoryOptimizedSchemaMutation,
	simpleTaskOptionsExtraction,
	useUserPermission,
	mergeFiltersRule,
	bandPlotType,
	mosaicGraphType,
	notebooks,
	pushDownGroupAggregateMinMax,
	orgOnlyMemberList,
	enforceOrgDashboardLimits,
}

var byKey = map[string]Flag{
	"appMetrics":                    appMetrics,
	"backendExample":                backendExample,
	"communityTemplates":            communityTemplates,
	"frontendExample":               frontendExample,
	"pushDownWindowAggregateMean":   pushDownWindowAggregateMean,
	"groupWindowAggregateTranspose": groupWindowAggregateTranspose,
	"newAuth":                       newAuth,
	"newLabels":                     newLabels,
	"hydratevars":                   hydratevars,
	"queryCacheForDashboards":       queryCacheForDashboards,
	"memoryOptimizedFill":           memoryOptimizedFill,
	"memoryOptimizedSchemaMutation": memoryOptimizedSchemaMutation,
	"simpleTaskOptionsExtraction":   simpleTaskOptionsExtraction,
	"useUserPermission":             useUserPermission,
	"mergeFiltersRule":              mergeFiltersRule,
	"bandPlotType":                  bandPlotType,
	"mosaicGraphType":               mosaicGraphType,
	"notebooks":                     notebooks,
	"pushDownGroupAggregateMinMax":  pushDownGroupAggregateMinMax,
	"orgOnlyMemberList":             orgOnlyMemberList,
	"enforceOrgDashboardLimits":     enforceOrgDashboardLimits,
}
