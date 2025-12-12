package direct

// Event names from direct-js handleNotification.
// These are server-to-client notifications prefixed with "notify_".
const (
	// Connection events
	EventSessionCreated     = "session_created"
	EventSessionError       = "session_error"
	EventDataRecovered      = "data_recovered"
	EventNotificationError  = "notification_error"
	EventError              = "error"
	EventDecodeError        = "decode_error"
	EventAccessTokenChanged = "access_token_changed"

	// Message notifications
	EventNotifyCreateMessage = "notify_create_message"
	EventNotifyDeleteMessage = "notify_delete_message"

	// Talk/Room notifications
	EventNotifyCreateGroupTalk = "notify_create_group_talk"
	EventNotifyCreatePairTalk  = "notify_create_pair_talk"
	EventNotifyAddTalkers      = "notify_add_talkers"
	EventNotifyDeleteTalker    = "notify_delete_talker"
	EventNotifyUpdateTalk      = "notify_update_talk"

	// User/Friend notifications
	EventNotifyAddFriend        = "notify_add_friend"
	EventNotifyDeleteFriend     = "notify_delete_friend"
	EventNotifyAddAcquaintance  = "notify_add_acquaintance"
	EventNotifyAddAcquaintances = "notify_add_acquaintances"
	EventNotifyUpdateUser       = "notify_update_user"

	// Domain notifications
	EventNotifyJoinDomain         = "notify_join_domain"
	EventNotifyLeaveDomain        = "notify_leave_domain"
	EventNotifyAddDomainInvite    = "notify_add_domain_invite"
	EventNotifyDeleteDomainInvite = "notify_delete_domain_invite"

	// Attachment notifications
	EventNotifyCreateAttachment = "notify_create_attachment"
	EventNotifyDeleteAttachment = "notify_delete_attachment"

	// Note notifications
	EventNotifyCreateNote = "notify_create_note"
	EventNotifyUpdateNote = "notify_update_note"
	EventNotifyDeleteNote = "notify_delete_note"

	// Favorite notifications
	EventNotifyAddFavoriteTalk    = "notify_add_favorite_talk"
	EventNotifyDeleteFavoriteTalk = "notify_delete_favorite_talk"

	// Announcement notifications
	EventNotifyCreateAnnouncement = "notify_create_announcement"
	EventNotifyDeleteAnnouncement = "notify_delete_announcement"

	// Read status notifications
	EventNotifyUpdateReadStatus = "notify_update_read_status"
	EventNotifyUpdateTalkStatus = "notify_update_talk_status"

	// Conference notifications
	EventNotifyCreateConference = "notify_create_conference"
	EventNotifyCloseConference  = "notify_close_conference"
	EventNotifyConferenceJoin   = "notify_conference_participant_join"
	EventNotifyConferenceReject = "notify_conference_participant_reject"
)

// API method names for RPC calls.
const (
	// Session
	MethodCreateSession     = "create_session"
	MethodStartNotification = "start_notification"
	MethodResetNotification = "reset_notification"
	MethodUpdateLastUsedAt  = "update_last_used_at"

	// Authentication
	MethodCreateAccessToken     = "create_access_token"
	MethodCreateAccessTokenByID = "create_access_token_by_id"
	MethodAuthorizeDevice       = "authorize_device"

	// Users
	MethodGetMe              = "get_me"
	MethodGetUsers           = "get_users"
	MethodGetProfile         = "get_profile"
	MethodUpdateUser         = "update_user"
	MethodUpdateProfile      = "update_profile"
	MethodGetPresences       = "get_presences"
	MethodGetUserIdentifiers = "get_user_identifiers"

	// Friends
	MethodAddFriend        = "add_friend"
	MethodDeleteFriend     = "delete_friend"
	MethodGetFriends       = "get_friends"
	MethodGetAcquaintances = "get_acquaintances"

	// Domains
	MethodGetDomains         = "get_domains"
	MethodLeaveDomain        = "leave_domain"
	MethodGetDomainInvites   = "get_domain_invites"
	MethodAcceptDomainInvite = "accept_domain_invite"
	MethodDeleteDomainInvite = "delete_domain_invite"
	MethodGetDomainUsers     = "get_domain_users"
	MethodSearchDomainUsers  = "search_domain_users"

	// Departments
	MethodGetDepartmentTree      = "get_department_tree"
	MethodGetDepartmentUsers     = "get_department_users"
	MethodGetDepartmentUserCount = "get_department_user_count"

	// Talks
	MethodGetTalks        = "get_talks"
	MethodGetTalkStatuses = "get_talk_statuses"
	MethodCreateGroupTalk = "create_group_talk"
	MethodCreatePairTalk  = "create_pair_talk"
	MethodUpdateGroupTalk = "update_group_talk"
	MethodAddTalkers      = "add_talkers"
	MethodDeleteTalker    = "delete_talker"

	// Favorites
	MethodAddFavoriteTalk    = "add_favorite_talk"
	MethodDeleteFavoriteTalk = "delete_favorite_talk"

	// Messages
	MethodGetMessages                  = "get_messages"
	MethodCreateMessage                = "create_message"
	MethodDeleteMessage                = "delete_message"
	MethodScheduleMessage              = "schedule_message"
	MethodSearchMessages               = "search_messages"
	MethodSearchMessagesAroundDateTime = "search_messages_around_datetime"
	MethodGetFavoriteMessages          = "get_favorite_messages"
	MethodAddFavoriteMessage           = "add_favorite_message"
	MethodDeleteFavoriteMessage        = "delete_favorite_message"
	MethodGetScheduledMessages         = "get_scheduled_messages"
	MethodDeleteScheduledMessage       = "delete_scheduled_message"
	MethodRescheduleMessage            = "reschedule_message"
	MethodGetAvailableMessageReactions = "get_available_message_reactions"
	MethodSetMessageReaction           = "set_message_reaction"
	MethodResetMessageReaction         = "reset_message_reaction"
	MethodGetMessageReactionUsers      = "get_message_reaction_users"

	// File & Attachment
	MethodCreateUploadAuth  = "create_upload_auth"
	MethodGetAttachments    = "get_attachments"
	MethodDeleteAttachment  = "delete_attachment"
	MethodSearchAttachments = "search_attachments"
	MethodCreateFilePreview = "create_file_preview"
	MethodGetFilePreview    = "get_file_preview"

	// Read status
	MethodGetReadStatus = "get_read_status"

	// Push notifications
	MethodDisablePushNotification = "disable_push_notification"
	MethodEnablePushNotification  = "enable_push_notification"

	// Announcements
	MethodCreateAnnouncement       = "create_announcement"
	MethodGetAnnouncements         = "get_announcements"
	MethodGetAnnouncementStatuses  = "get_announcement_statuses"
	MethodGetAnnouncementStatus    = "get_announcement_status"
	MethodUpdateAnnouncementStatus = "update_announcement_status"

	// Conference/Call
	MethodGetConferences            = "get_conferences"
	MethodGetConferenceParticipants = "get_conference_participants"
	MethodJoinConference            = "join_conference"
	MethodLeaveConference           = "leave_conference"
	MethodRejectConference          = "reject_conference"
)

// Message types from direct API.
// NOTE: For action stamps (types 13-21), these are INTERNAL enum values.
// When SENDING to the API, use the WireType constants below instead.
const (
	MsgTypeSystem           = 0  // System message
	MsgTypeText             = 1  // Text message
	MsgTypeStamp            = 2  // Stamp
	MsgTypeLocation         = 3  // Location (geo)
	MsgTypeFile             = 4  // Single file
	MsgTypeTextMultipleFile = 5  // Text with multiple files
	MsgTypeUnused           = 6  // Reserved / unused
	MsgTypeDeleted          = 7  // Deleted message
	MsgTypeNoteShared       = 8  // Note shared
	MsgTypeNoteDeleted      = 9  // Note deleted
	MsgTypeNoteCreated      = 10 // Note created
	MsgTypeNoteUpdated      = 11 // Note updated
	MsgTypeOriginalStamp    = 12 // Original stamp
	MsgTypeYesNo            = 13 // Yes/No action stamp (internal)
	MsgTypeYesNoReply       = 14 // Yes/No reply (internal)
	MsgTypeSelect           = 15 // Select action stamp (internal)
	MsgTypeSelectReply      = 16 // Select reply (internal)
	MsgTypeTask             = 17 // Task action stamp (internal)
	MsgTypeTaskDone         = 18 // Task done reply (internal)
	MsgTypeYesNoClosed      = 19 // Yes/No closed (internal)
	MsgTypeSelectClosed     = 20 // Select closed (internal)
	MsgTypeTaskClosed       = 21 // Task closed (internal)
)

// Wire message types for action stamps.
// The API expects wire types for action stamps (internal types 13-21).
// Formula: wireType = 500 + internalType - 13
// These are the values that must be used in create_message API calls.
const (
	WireTypeYesNo        = 500 // 500 + 13 - 13 = 500
	WireTypeYesNoReply   = 501 // 500 + 14 - 13 = 501
	WireTypeSelect       = 502 // 500 + 15 - 13 = 502
	WireTypeSelectReply  = 503 // 500 + 16 - 13 = 503
	WireTypeTask         = 504 // 500 + 17 - 13 = 504
	WireTypeTaskDone     = 505 // 500 + 18 - 13 = 505
	WireTypeYesNoClosed  = 506 // 500 + 19 - 13 = 506
	WireTypeSelectClosed = 507 // 500 + 20 - 13 = 507
	WireTypeTaskClosed   = 508 // 500 + 21 - 13 = 508
)
