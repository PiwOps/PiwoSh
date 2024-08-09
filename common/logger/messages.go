// logger/messages.go
package logger

type Resource string

const (
	ResourceCommand        Resource = "COMMAND"
	ResourceConfig         Resource = "CONFIG"
	ResourceCmdConfigEntry Resource = "CMD_CONFIG_ENTRY"
	ResourceEnv            Resource = "ENV"
	ResourceChannel        Resource = "CHANNEL"
	ResourceGuild          Resource = "GUILD"
	ResourceUser           Resource = "USER"
)

var (
	MsgLoadingResource    = NewMessageWrapper("MsgLoadingResource", "loading resource %s (%s)", InfoLevel)
	MsgInitializing       = NewMessageWrapper("MsgInitializing", "initializing", InfoLevel)
	MsgCleaningUp         = NewMessageWrapper("MsgCleaningUp", "cleaning up", InfoLevel)
	MsgRegisteredResource = NewMessageWrapper("MsgRegisteredResource", "registered resource %s (%s)", SuccessLevel)
	MsgRemovedResource    = NewMessageWrapper("MsgRemovedResource", "removed resource %s (%s)", SuccessLevel)
)

var (
	ErrNoHandlerFound          = NewMessageWrapper("ErrNoHandlerFound", "no handler found for command: %s", WarnLevel)
	ErrRespondingToInteraction = NewMessageWrapper("ErrRespondingToInteraction", "error responding to interaction: %s", WarnLevel)
	ErrReadingResource         = NewMessageWrapper("ErrReadingResource", "error reading %s (%s)", FatalLevel)
	ErrFormattingResource      = NewMessageWrapper("ErrFormattingResource", "error formatting %s", FatalLevel)
	ErrResourcesDirectoryEmpty = NewMessageWrapper("ErrResourcesDirectoryEmpty", "resources (%s) directory is empty", FatalLevel)
	ErrResourceAlreadyExists   = NewMessageWrapper("ErrResourceAlreadyExists", "resource %s (%s) already exists", ErrorLevel)
	ErrResourceNotFound        = NewMessageWrapper("ErrResourceNotFound", "resource %s (%s) not found", ErrorLevel)
	ErrRegisteringResource     = NewMessageWrapper("ErrRegisteringResource", "error registering resource %s (%s)", ErrorLevel)
	ErrFetchingResource		   = NewMessageWrapper("ErrFetchingResource", "error fetching resource %s (%s): %s", ErrorLevel)
	ErrRemovingResource        = NewMessageWrapper("ErrRemovingResource", "error removing resource %s (%s): %s", ErrorLevel)
	ErrLoadingResource         = NewMessageWrapper("ErrLoadingResource", "error loading resource %s (%s): %s", FatalLevel)
	ErrInitializing            = NewMessageWrapper("ErrInitializing", "error initializing: %s", FatalLevel)
	ErrCleaningUp              = NewMessageWrapper("ErrCleaningUp", "error cleaning up: %s", ErrorLevel)
)
