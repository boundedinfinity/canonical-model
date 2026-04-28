package message

// OBS WebSocket v5.x Kind values
type Kind string

var Kinds = kinds{
	// ---------------------------
	// General
	// ---------------------------
	GetVersion:                 "GetVersion",
	GetStats:                   "GetStats",
	BroadcastCustomEvent:       "BroadcastCustomEvent",
	CallVendorRequest:          "CallVendorRequest",
	GetHotkeyList:              "GetHotkeyList",
	TriggerHotkeyByName:        "TriggerHotkeyByName",
	TriggerHotkeyByKeySequence: "TriggerHotkeyByKeySequence",
	Sleep:                      "Sleep",

	// ---------------------------
	// Scenes
	// ---------------------------
	GetSceneList:           "GetSceneList",
	GetCurrentProgramScene: "GetCurrentProgramScene",
	SetCurrentProgramScene: "SetCurrentProgramScene",
	GetCurrentPreviewScene: "GetCurrentPreviewScene",
	SetCurrentPreviewScene: "SetCurrentPreviewScene",
	CreateScene:            "CreateScene",
	RemoveScene:            "RemoveScene",
	SetSceneName:           "SetSceneName",

	// ---------------------------
	// Scene Items
	// ---------------------------
	GetSceneItemList:      "GetSceneItemList",
	GetSceneItemId:        "GetSceneItemId",
	CreateSceneItem:       "CreateSceneItem",
	RemoveSceneItem:       "RemoveSceneItem",
	DuplicateSceneItem:    "DuplicateSceneItem",
	GetSceneItemEnabled:   "GetSceneItemEnabled",
	SetSceneItemEnabled:   "SetSceneItemEnabled",
	GetSceneItemIndex:     "GetSceneItemIndex",
	SetSceneItemIndex:     "SetSceneItemIndex",
	GetSceneItemTransform: "GetSceneItemTransform",
	SetSceneItemTransform: "SetSceneItemTransform",
	GetSceneItemBlendMode: "GetSceneItemBlendMode",
	SetSceneItemBlendMode: "SetSceneItemBlendMode",

	// ---------------------------
	// Inputs (Sources)
	// ---------------------------
	GetInputList:             "GetInputList",
	GetInputKindList:         "GetInputKindList",
	GetSpecialInputs:         "GetSpecialInputs",
	CreateInput:              "CreateInput",
	RemoveInput:              "RemoveInput",
	SetInputName:             "SetInputName",
	GetInputSettings:         "GetInputSettings",
	SetInputSettings:         "SetInputSettings",
	GetInputDefaultSettings:  "GetInputDefaultSettings",
	GetInputMute:             "GetInputMute",
	SetInputMute:             "SetInputMute",
	ToggleInputMute:          "ToggleInputMute",
	GetInputVolume:           "GetInputVolume",
	SetInputVolume:           "SetInputVolume",
	GetInputAudioBalance:     "GetInputAudioBalance",
	SetInputAudioBalance:     "SetInputAudioBalance",
	GetInputAudioSyncOffset:  "GetInputAudioSyncOffset",
	SetInputAudioSyncOffset:  "SetInputAudioSyncOffset",
	GetInputAudioMonitorType: "GetInputAudioMonitorType",
	SetInputAudioMonitorType: "SetInputAudioMonitorType",
	GetInputAudioTracks:      "GetInputAudioTracks",
	SetInputAudioTracks:      "SetInputAudioTracks",

	// ---------------------------
	// Source Filters
	// ---------------------------
	GetSourceFilterList:            "GetSourceFilterList",
	GetSourceFilter:                "GetSourceFilter",
	CreateSourceFilter:             "CreateSourceFilter",
	RemoveSourceFilter:             "RemoveSourceFilter",
	SetSourceFilterName:            "SetSourceFilterName",
	SetSourceFilterIndex:           "SetSourceFilterIndex",
	SetSourceFilterEnabled:         "SetSourceFilterEnabled",
	GetSourceFilterDefaultSettings: "GetSourceFilterDefaultSettings",
	SetSourceFilterSettings:        "SetSourceFilterSettings",

	// ---------------------------
	// Transitions
	// ---------------------------
	GetTransitionKindList:             "GetTransitionKindList",
	GetSceneTransitionList:            "GetSceneTransitionList",
	GetCurrentSceneTransition:         "GetCurrentSceneTransition",
	SetCurrentSceneTransition:         "SetCurrentSceneTransition",
	SetCurrentSceneTransitionDuration: "SetCurrentSceneTransitionDuration",
	GetCurrentSceneTransitionCursor:   "GetCurrentSceneTransitionCursor",
	TriggerStudioModeTransition:       "TriggerStudioModeTransition",

	// ---------------------------
	// Studio Mode
	// ---------------------------
	GetStudioModeEnabled: "GetStudioModeEnabled",
	SetStudioModeEnabled: "SetStudioModeEnabled",

	// ---------------------------
	// Streaming
	// ---------------------------
	GetStreamStatus:   "GetStreamStatus",
	StartStream:       "StartStream",
	StopStream:        "StopStream",
	ToggleStream:      "ToggleStream",
	SendStreamCaption: "SendStreamCaption",

	// ---------------------------
	// Recording
	// ---------------------------
	GetRecordStatus: "GetRecordStatus",
	StartRecord:     "StartRecord",
	StopRecord:      "StopRecord",
	ToggleRecord:    "ToggleRecord",
	PauseRecord:     "PauseRecord",
	ResumeRecord:    "ResumeRecord",

	// ---------------------------
	// Replay Buffer
	// ---------------------------
	GetReplayBufferStatus: "GetReplayBufferStatus",
	StartReplayBuffer:     "StartReplayBuffer",
	StopReplayBuffer:      "StopReplayBuffer",
	SaveReplayBuffer:      "SaveReplayBuffer",

	// ---------------------------
	// Virtual Camera
	// ---------------------------
	GetVirtualCamStatus: "GetVirtualCamStatus",
	StartVirtualCam:     "StartVirtualCam",
	StopVirtualCam:      "StopVirtualCam",
	ToggleVirtualCam:    "ToggleVirtualCam",

	// ---------------------------
	// Media Inputs
	// ---------------------------
	GetMediaInputStatus:     "GetMediaInputStatus",
	SetMediaInputCursor:     "SetMediaInputCursor",
	OffsetMediaInputCursor:  "OffsetMediaInputCursor",
	TriggerMediaInputAction: "TriggerMediaInputAction",

	// ---------------------------
	// Profiles
	// ---------------------------
	GetProfileList:    "GetProfileList",
	GetCurrentProfile: "GetCurrentProfile",
	SetCurrentProfile: "SetCurrentProfile",

	// ---------------------------
	// Scene Collections
	// ---------------------------
	GetSceneCollectionList:    "GetSceneCollectionList",
	GetCurrentSceneCollection: "GetCurrentSceneCollection",
	SetCurrentSceneCollection: "SetCurrentSceneCollection",
	CreateSceneCollection:     "CreateSceneCollection",

	// ---------------------------
	// Outputs (Advanced)
	// ---------------------------
	GetOutputList:   "GetOutputList",
	GetOutputStatus: "GetOutputStatus",
	StartOutput:     "StartOutput",
	StopOutput:      "StopOutput",

	// ---------------------------
	// Configuration / Video
	// ---------------------------
	GetVideoSettings: "GetVideoSettings",
	SetVideoSettings: "SetVideoSettings",

	// ---------------------------
	// Screenshot
	// ---------------------------
	GetSourceScreenshot: "GetSourceScreenshot",

	// ---------------------------
	// Vendor / Extensions
	// ---------------------------
	GetVendorList: "GetVendorList",
}

func init() {
	Kinds.All = []Kind{
		Kinds.GetVersion,
		Kinds.GetStats,
		Kinds.BroadcastCustomEvent,
		Kinds.CallVendorRequest,
		Kinds.GetHotkeyList,
		Kinds.TriggerHotkeyByName,
		Kinds.TriggerHotkeyByKeySequence,
		Kinds.Sleep,

		Kinds.GetSceneList,
		Kinds.GetCurrentProgramScene,
		Kinds.SetCurrentProgramScene,
		Kinds.GetCurrentPreviewScene,
		Kinds.SetCurrentPreviewScene,
		Kinds.CreateScene,
		Kinds.RemoveScene,
		Kinds.SetSceneName,

		Kinds.GetSceneItemList,
		Kinds.GetSceneItemId,
		Kinds.CreateSceneItem,
		Kinds.RemoveSceneItem,
		Kinds.DuplicateSceneItem,
		Kinds.GetSceneItemEnabled,
		Kinds.SetSceneItemEnabled,
		Kinds.GetSceneItemIndex,
		Kinds.SetSceneItemIndex,
		Kinds.GetSceneItemTransform,
		Kinds.SetSceneItemTransform,
		Kinds.GetSceneItemBlendMode,
		Kinds.SetSceneItemBlendMode,

		Kinds.GetInputList,
		Kinds.GetInputKindList,
		Kinds.GetSpecialInputs,
		Kinds.CreateInput,
		Kinds.RemoveInput,
		Kinds.SetInputName,
		Kinds.GetInputSettings,
		Kinds.SetInputSettings,
		Kinds.GetInputDefaultSettings,
		Kinds.GetInputMute,
		Kinds.SetInputMute,
		Kinds.ToggleInputMute,
		Kinds.GetInputVolume,
		Kinds.SetInputVolume,
		Kinds.GetInputAudioBalance,
		Kinds.SetInputAudioBalance,
		Kinds.GetInputAudioSyncOffset,
		Kinds.SetInputAudioSyncOffset,
		Kinds.GetInputAudioMonitorType,
		Kinds.SetInputAudioMonitorType,
		Kinds.GetInputAudioTracks,
		Kinds.SetInputAudioTracks,

		Kinds.GetSourceFilterList,
		Kinds.GetSourceFilter,
		Kinds.CreateSourceFilter,
		Kinds.RemoveSourceFilter,
		Kinds.SetSourceFilterName,
		Kinds.SetSourceFilterIndex,
		Kinds.SetSourceFilterEnabled,
		Kinds.GetSourceFilterDefaultSettings,
		Kinds.SetSourceFilterSettings,

		Kinds.GetTransitionKindList,
		Kinds.GetSceneTransitionList,
		Kinds.GetCurrentSceneTransition,
		Kinds.SetCurrentSceneTransition,
		Kinds.SetCurrentSceneTransitionDuration,
		Kinds.GetCurrentSceneTransitionCursor,
		Kinds.TriggerStudioModeTransition,

		Kinds.GetStudioModeEnabled,
		Kinds.SetStudioModeEnabled,

		Kinds.GetStreamStatus,
		Kinds.StartStream,
		Kinds.StopStream,
		Kinds.ToggleStream,
		Kinds.SendStreamCaption,

		Kinds.GetRecordStatus,
		Kinds.StartRecord,
		Kinds.StopRecord,
		Kinds.ToggleRecord,
		Kinds.PauseRecord,
		Kinds.ResumeRecord,

		Kinds.GetReplayBufferStatus,
		Kinds.StartReplayBuffer,
		Kinds.StopReplayBuffer,
		Kinds.SaveReplayBuffer,

		Kinds.GetVirtualCamStatus,
		Kinds.StartVirtualCam,
		Kinds.StopVirtualCam,
		Kinds.ToggleVirtualCam,

		Kinds.GetMediaInputStatus,
		Kinds.SetMediaInputCursor,
		Kinds.OffsetMediaInputCursor,
		Kinds.TriggerMediaInputAction,

		Kinds.GetProfileList,
		Kinds.GetCurrentProfile,
		Kinds.SetCurrentProfile,

		Kinds.GetSceneCollectionList,
		Kinds.GetCurrentSceneCollection,
		Kinds.SetCurrentSceneCollection,
		Kinds.CreateSceneCollection,

		Kinds.GetOutputList,
		Kinds.GetOutputStatus,
		Kinds.StartOutput,
		Kinds.StopOutput,

		Kinds.GetVideoSettings,
		Kinds.SetVideoSettings,

		Kinds.GetSourceScreenshot,

		Kinds.GetVendorList,
	}
}

type kinds struct {
	All []Kind
	// ---------------------------
	// General
	// ---------------------------
	GetVersion                 Kind
	GetStats                   Kind
	BroadcastCustomEvent       Kind
	CallVendorRequest          Kind
	GetHotkeyList              Kind
	TriggerHotkeyByName        Kind
	TriggerHotkeyByKeySequence Kind
	Sleep                      Kind

	// ---------------------------
	// Scenes
	// ---------------------------
	GetSceneList           Kind
	GetCurrentProgramScene Kind
	SetCurrentProgramScene Kind
	GetCurrentPreviewScene Kind
	SetCurrentPreviewScene Kind
	CreateScene            Kind
	RemoveScene            Kind
	SetSceneName           Kind

	// ---------------------------
	// Scene Items
	// ---------------------------
	GetSceneItemList      Kind
	GetSceneItemId        Kind
	CreateSceneItem       Kind
	RemoveSceneItem       Kind
	DuplicateSceneItem    Kind
	GetSceneItemEnabled   Kind
	SetSceneItemEnabled   Kind
	GetSceneItemIndex     Kind
	SetSceneItemIndex     Kind
	GetSceneItemTransform Kind
	SetSceneItemTransform Kind
	GetSceneItemBlendMode Kind
	SetSceneItemBlendMode Kind

	// ---------------------------
	// Inputs (Sources)
	// ---------------------------
	GetInputList             Kind
	GetInputKindList         Kind
	GetSpecialInputs         Kind
	CreateInput              Kind
	RemoveInput              Kind
	SetInputName             Kind
	GetInputSettings         Kind
	SetInputSettings         Kind
	GetInputDefaultSettings  Kind
	GetInputMute             Kind
	SetInputMute             Kind
	ToggleInputMute          Kind
	GetInputVolume           Kind
	SetInputVolume           Kind
	GetInputAudioBalance     Kind
	SetInputAudioBalance     Kind
	GetInputAudioSyncOffset  Kind
	SetInputAudioSyncOffset  Kind
	GetInputAudioMonitorType Kind
	SetInputAudioMonitorType Kind
	GetInputAudioTracks      Kind
	SetInputAudioTracks      Kind

	// ---------------------------
	// Source Filters
	// ---------------------------
	GetSourceFilterList            Kind
	GetSourceFilter                Kind
	CreateSourceFilter             Kind
	RemoveSourceFilter             Kind
	SetSourceFilterName            Kind
	SetSourceFilterIndex           Kind
	SetSourceFilterEnabled         Kind
	GetSourceFilterDefaultSettings Kind
	SetSourceFilterSettings        Kind

	// ---------------------------
	// Transitions
	// ---------------------------
	GetTransitionKindList             Kind
	GetSceneTransitionList            Kind
	GetCurrentSceneTransition         Kind
	SetCurrentSceneTransition         Kind
	SetCurrentSceneTransitionDuration Kind
	GetCurrentSceneTransitionCursor   Kind
	TriggerStudioModeTransition       Kind

	// ---------------------------
	// Studio Mode
	// ---------------------------
	GetStudioModeEnabled Kind
	SetStudioModeEnabled Kind

	// ---------------------------
	// Streaming
	// ---------------------------
	GetStreamStatus   Kind
	StartStream       Kind
	StopStream        Kind
	ToggleStream      Kind
	SendStreamCaption Kind

	// ---------------------------
	// Recording
	// ---------------------------
	GetRecordStatus Kind
	StartRecord     Kind
	StopRecord      Kind
	ToggleRecord    Kind
	PauseRecord     Kind
	ResumeRecord    Kind

	// ---------------------------
	// Replay Buffer
	// ---------------------------
	GetReplayBufferStatus Kind
	StartReplayBuffer     Kind
	StopReplayBuffer      Kind
	SaveReplayBuffer      Kind

	// ---------------------------
	// Virtual Camera
	// ---------------------------
	GetVirtualCamStatus Kind
	StartVirtualCam     Kind
	StopVirtualCam      Kind
	ToggleVirtualCam    Kind

	// ---------------------------
	// Media Inputs
	// ---------------------------
	GetMediaInputStatus     Kind
	SetMediaInputCursor     Kind
	OffsetMediaInputCursor  Kind
	TriggerMediaInputAction Kind

	// ---------------------------
	// Profiles
	// ---------------------------
	GetProfileList    Kind
	GetCurrentProfile Kind
	SetCurrentProfile Kind

	// ---------------------------
	// Scene Collections
	// ---------------------------
	GetSceneCollectionList    Kind
	GetCurrentSceneCollection Kind
	SetCurrentSceneCollection Kind
	CreateSceneCollection     Kind

	// ---------------------------
	// Outputs (Advanced)
	// ---------------------------
	GetOutputList   Kind
	GetOutputStatus Kind
	StartOutput     Kind
	StopOutput      Kind

	// ---------------------------
	// Configuration / Video
	// ---------------------------
	GetVideoSettings Kind
	SetVideoSettings Kind

	// ---------------------------
	// Screenshot
	// ---------------------------
	GetSourceScreenshot Kind

	// ---------------------------
	// Vendor / Extensions
	// ---------------------------
	GetVendorList Kind
}
