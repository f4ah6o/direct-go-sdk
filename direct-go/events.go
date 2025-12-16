// events.go defines event names, RPC method names, and message type constants
// used by the direct API.
package direct

// Server-to-client event names that can be received via Client.On().
// These events are emitted when notifications arrive from the direct service.
const (
	// Connection and session events
	// EventSessionCreated is emitted when the session is successfully created after connect.
	EventSessionCreated = "session_created"

	// EventSessionError is emitted when session creation fails (authentication error).
	EventSessionError = "session_error"

	// EventDataRecovered is emitted when initial data sync is complete and the bot is ready.
	EventDataRecovered = "data_recovered"

	// EventNotificationError is emitted when a notification system error occurs.
	EventNotificationError = "notification_error"

	// EventError is emitted for general connection or protocol errors.
	EventError = "error"

	// EventDecodeError is emitted when a message cannot be decoded.
	EventDecodeError = "decode_error"

	// EventAccessTokenChanged is emitted when the access token changes.
	EventAccessTokenChanged = "access_token_changed"

	// Message notifications - emitted when messages are sent/deleted
	// EventNotifyCreateMessage is emitted when a new message is received.
	EventNotifyCreateMessage = "notify_create_message"

	// EventNotifyDeleteMessage is emitted when a message is deleted.
	EventNotifyDeleteMessage = "notify_delete_message"

	// Talk/Room notifications - emitted for room/conversation changes
	// EventNotifyCreateGroupTalk is emitted when a new group talk is created.
	EventNotifyCreateGroupTalk = "notify_create_group_talk"

	// EventNotifyCreatePairTalk is emitted when a new pair (1:1) talk is created.
	EventNotifyCreatePairTalk = "notify_create_pair_talk"

	// EventNotifyAddTalkers is emitted when users are added to a group talk.
	EventNotifyAddTalkers = "notify_add_talkers"

	// EventNotifyDeleteTalker is emitted when a user is removed from a talk.
	EventNotifyDeleteTalker = "notify_delete_talker"

	// EventNotifyUpdateTalk is emitted when a talk is updated (name, settings, etc.).
	EventNotifyUpdateTalk = "notify_update_talk"

	// User/Friend notifications
	// EventNotifyAddFriend is emitted when a new friend is added.
	EventNotifyAddFriend = "notify_add_friend"

	// EventNotifyDeleteFriend is emitted when a friend is removed.
	EventNotifyDeleteFriend = "notify_delete_friend"

	// EventNotifyAddAcquaintance is emitted when an acquaintance is added.
	EventNotifyAddAcquaintance = "notify_add_acquaintance"

	// EventNotifyAddAcquaintances is emitted when multiple acquaintances are added.
	EventNotifyAddAcquaintances = "notify_add_acquaintances"

	// EventNotifyUpdateUser is emitted when user information is updated.
	EventNotifyUpdateUser = "notify_update_user"

	// Domain/Organization notifications
	// EventNotifyJoinDomain is emitted when the user joins a domain.
	EventNotifyJoinDomain = "notify_join_domain"

	// EventNotifyLeaveDomain is emitted when the user leaves a domain.
	EventNotifyLeaveDomain = "notify_leave_domain"

	// EventNotifyAddDomainInvite is emitted when a domain invitation is received.
	EventNotifyAddDomainInvite = "notify_add_domain_invite"

	// EventNotifyDeleteDomainInvite is emitted when a domain invitation is removed.
	EventNotifyDeleteDomainInvite = "notify_delete_domain_invite"

	// Attachment notifications
	// EventNotifyCreateAttachment is emitted when a file attachment is created.
	EventNotifyCreateAttachment = "notify_create_attachment"

	// EventNotifyDeleteAttachment is emitted when a file attachment is deleted.
	EventNotifyDeleteAttachment = "notify_delete_attachment"

	// Note notifications
	// EventNotifyCreateNote is emitted when a note is created.
	EventNotifyCreateNote = "notify_create_note"

	// EventNotifyUpdateNote is emitted when a note is updated.
	EventNotifyUpdateNote = "notify_update_note"

	// EventNotifyDeleteNote is emitted when a note is deleted.
	EventNotifyDeleteNote = "notify_delete_note"

	// Favorite notifications
	// EventNotifyAddFavoriteTalk is emitted when a talk is added to favorites.
	EventNotifyAddFavoriteTalk = "notify_add_favorite_talk"

	// EventNotifyDeleteFavoriteTalk is emitted when a talk is removed from favorites.
	EventNotifyDeleteFavoriteTalk = "notify_delete_favorite_talk"

	// Announcement notifications
	// EventNotifyCreateAnnouncement is emitted when a new announcement is created.
	EventNotifyCreateAnnouncement = "notify_create_announcement"

	// EventNotifyDeleteAnnouncement is emitted when an announcement is deleted.
	EventNotifyDeleteAnnouncement = "notify_delete_announcement"

	// Read status notifications
	// EventNotifyUpdateReadStatus is emitted when message read status changes.
	EventNotifyUpdateReadStatus = "notify_update_read_status"

	// EventNotifyUpdateTalkStatus is emitted when talk status changes (unread count, etc.).
	EventNotifyUpdateTalkStatus = "notify_update_talk_status"

	// Conference/Call notifications
	// EventNotifyCreateConference is emitted when a new conference/call is started.
	EventNotifyCreateConference = "notify_create_conference"

	// EventNotifyCloseConference is emitted when a conference/call ends.
	EventNotifyCloseConference = "notify_close_conference"

	// EventNotifyConferenceJoin is emitted when a participant joins a conference.
	EventNotifyConferenceJoin = "notify_conference_participant_join"

	// EventNotifyConferenceReject is emitted when a participant rejects a conference invitation.
	EventNotifyConferenceReject = "notify_conference_participant_reject"
)

// RPC method names used with Client.Call() to invoke direct API operations.
// Use these constants when calling Client.Call() to ensure correct method names.
const (
	// Session management methods
	// MethodCreateSession authenticates with the direct service using an access token.
	MethodCreateSession = "create_session"

	// MethodStartNotification enables receiving server notifications.
	MethodStartNotification = "start_notification"

	// MethodResetNotification resets the notification state.
	MethodResetNotification = "reset_notification"

	// MethodUpdateLastUsedAt updates the session's last-used timestamp.
	MethodUpdateLastUsedAt = "update_last_used_at"

	// Authentication methods
	// MethodCreateAccessToken creates a new access token.
	MethodCreateAccessToken = "create_access_token"

	// MethodCreateAccessTokenByID creates a new access token using a user ID.
	MethodCreateAccessTokenByID = "create_access_token_by_id"

	// MethodAuthorizeDevice authorizes a device for the current session.
	MethodAuthorizeDevice = "authorize_device"

	// User management methods
	// MethodGetMe retrieves the current authenticated user's information.
	MethodGetMe = "get_me"

	// MethodGetUsers retrieves information about specific users.
	MethodGetUsers = "get_users"

	// MethodGetProfile retrieves detailed profile information for a user.
	MethodGetProfile = "get_profile"

	// MethodUpdateUser updates user information.
	MethodUpdateUser = "update_user"

	// MethodUpdateProfile updates the current user's profile.
	MethodUpdateProfile = "update_profile"

	// MethodGetPresences retrieves online/offline status for users.
	MethodGetPresences = "get_presences"

	// MethodGetUserIdentifiers retrieves user identity information (email, alias).
	MethodGetUserIdentifiers = "get_user_identifiers"

	// Friend management methods
	// MethodAddFriend adds a user to the current user's friends list.
	MethodAddFriend = "add_friend"

	// MethodDeleteFriend removes a user from the friends list.
	MethodDeleteFriend = "delete_friend"

	// MethodGetFriends retrieves the current user's friends list.
	MethodGetFriends = "get_friends"

	// MethodGetAcquaintances retrieves the current user's acquaintances.
	MethodGetAcquaintances = "get_acquaintances"

	// Domain/Organization methods
	// MethodGetDomains retrieves the list of organizations the user belongs to.
	MethodGetDomains = "get_domains"

	// MethodLeaveDomain removes the current user from an organization.
	MethodLeaveDomain = "leave_domain"

	// MethodGetDomainInvites retrieves pending organization invitations.
	MethodGetDomainInvites = "get_domain_invites"

	// MethodAcceptDomainInvite accepts an organization invitation.
	MethodAcceptDomainInvite = "accept_domain_invite"

	// MethodDeleteDomainInvite deletes an organization invitation.
	MethodDeleteDomainInvite = "delete_domain_invite"

	// MethodGetDomainUsers retrieves users in an organization.
	MethodGetDomainUsers = "get_domain_users"

	// MethodSearchDomainUsers searches for users in an organization.
	MethodSearchDomainUsers = "search_domain_users"

	// Department methods
	// MethodGetDepartmentTree retrieves the organization's department hierarchy.
	MethodGetDepartmentTree = "get_department_tree"

	// MethodGetDepartmentUsers retrieves users in a specific department.
	MethodGetDepartmentUsers = "get_department_users"

	// MethodGetDepartmentUserCount retrieves the user count for a department.
	MethodGetDepartmentUserCount = "get_department_user_count"

	// Talk/Conversation methods
	// MethodGetTalks retrieves the list of conversation rooms.
	MethodGetTalks = "get_talks"

	// MethodGetTalkStatuses retrieves status information for all talks (unread counts, etc.).
	MethodGetTalkStatuses = "get_talk_statuses"

	// MethodCreateGroupTalk creates a new group conversation.
	MethodCreateGroupTalk = "create_group_talk"

	// MethodCreatePairTalk creates a new 1:1 conversation.
	MethodCreatePairTalk = "create_pair_talk"

	// MethodUpdateGroupTalk updates a group conversation's properties.
	MethodUpdateGroupTalk = "update_group_talk"

	// MethodAddTalkers adds users to a group conversation.
	MethodAddTalkers = "add_talkers"

	// MethodDeleteTalker removes a user from a conversation.
	MethodDeleteTalker = "delete_talker"

	// Favorite methods
	// MethodAddFavoriteTalk adds a conversation to favorites.
	MethodAddFavoriteTalk = "add_favorite_talk"

	// MethodDeleteFavoriteTalk removes a conversation from favorites.
	MethodDeleteFavoriteTalk = "delete_favorite_talk"

	// Message methods
	// MethodGetMessages retrieves messages from a conversation.
	MethodGetMessages = "get_messages"

	// MethodCreateMessage sends a message to a conversation.
	MethodCreateMessage = "create_message"

	// MethodDeleteMessage deletes a message.
	MethodDeleteMessage = "delete_message"

	// MethodScheduleMessage schedules a message to be sent at a future time.
	MethodScheduleMessage = "schedule_message"

	// MethodSearchMessages searches for messages across conversations.
	MethodSearchMessages = "search_messages"

	// MethodSearchMessagesAroundDateTime searches for messages around a specific time.
	MethodSearchMessagesAroundDateTime = "search_messages_around_datetime"

	// MethodGetFavoriteMessages retrieves the user's favorite messages.
	MethodGetFavoriteMessages = "get_favorite_messages"

	// MethodAddFavoriteMessage adds a message to favorites.
	MethodAddFavoriteMessage = "add_favorite_message"

	// MethodDeleteFavoriteMessage removes a message from favorites.
	MethodDeleteFavoriteMessage = "delete_favorite_message"

	// MethodGetScheduledMessages retrieves scheduled messages.
	MethodGetScheduledMessages = "get_scheduled_messages"

	// MethodDeleteScheduledMessage cancels a scheduled message.
	MethodDeleteScheduledMessage = "delete_scheduled_message"

	// MethodRescheduleMessage reschedules a previously scheduled message.
	MethodRescheduleMessage = "reschedule_message"

	// MethodGetAvailableMessageReactions retrieves the list of available reactions for messages.
	MethodGetAvailableMessageReactions = "get_available_message_reactions"

	// MethodSetMessageReaction adds a reaction to a message.
	MethodSetMessageReaction = "set_message_reaction"

	// MethodResetMessageReaction removes a reaction from a message.
	MethodResetMessageReaction = "reset_message_reaction"

	// MethodGetMessageReactionUsers retrieves users who reacted to a message.
	MethodGetMessageReactionUsers = "get_message_reaction_users"

	// File and attachment methods
	// MethodCreateUploadAuth creates credentials for uploading a file.
	MethodCreateUploadAuth = "create_upload_auth"

	// MethodGetAttachments retrieves attachments from a conversation.
	MethodGetAttachments = "get_attachments"

	// MethodDeleteAttachment deletes a file attachment.
	MethodDeleteAttachment = "delete_attachment"

	// MethodSearchAttachments searches for file attachments.
	MethodSearchAttachments = "search_attachments"

	// MethodCreateFilePreview creates a preview (thumbnail) for a file.
	MethodCreateFilePreview = "create_file_preview"

	// MethodGetFilePreview retrieves a file preview/thumbnail.
	MethodGetFilePreview = "get_file_preview"

	// Read status methods
	// MethodGetReadStatus retrieves read status for messages.
	MethodGetReadStatus = "get_read_status"

	// Push notification methods
	// MethodDisablePushNotification disables push notifications.
	MethodDisablePushNotification = "disable_push_notification"

	// MethodEnablePushNotification enables push notifications.
	MethodEnablePushNotification = "enable_push_notification"

	// Announcement methods
	// MethodCreateAnnouncement creates a new announcement.
	MethodCreateAnnouncement = "create_announcement"

	// MethodGetAnnouncements retrieves announcements.
	MethodGetAnnouncements = "get_announcements"

	// MethodGetAnnouncementStatuses retrieves read status for announcements.
	MethodGetAnnouncementStatuses = "get_announcement_statuses"

	// MethodGetAnnouncementStatus retrieves the read status for a specific announcement.
	MethodGetAnnouncementStatus = "get_announcement_status"

	// MethodUpdateAnnouncementStatus updates the read status of an announcement.
	MethodUpdateAnnouncementStatus = "update_announcement_status"

	// Conference/Call methods
	// MethodGetConferences retrieves active conferences/calls.
	MethodGetConferences = "get_conferences"

	// MethodGetConferenceParticipants retrieves participants in a conference.
	MethodGetConferenceParticipants = "get_conference_participants"

	// MethodJoinConference joins a conference/call.
	MethodJoinConference = "join_conference"

	// MethodLeaveConference leaves a conference/call.
	MethodLeaveConference = "leave_conference"

	// MethodRejectConference rejects a conference invitation.
	MethodRejectConference = "reject_conference"
)

// Message type constants for received messages.
// These are used in ReceivedMessage.Type to identify the message content type.
// IMPORTANT: For action stamps (types 13-21), these are INTERNAL enum values used
// when receiving messages from the server. When SENDING action stamps via Client.Send,
// use the WireType constants below instead.
const (
	// MsgTypeSystem is a system-generated message.
	MsgTypeSystem = 0

	// MsgTypeText is a plain text message.
	MsgTypeText = 1

	// MsgTypeStamp is an emoji/stamp reaction.
	MsgTypeStamp = 2

	// MsgTypeLocation is a location/map share.
	MsgTypeLocation = 3

	// MsgTypeFile is a single file attachment.
	MsgTypeFile = 4

	// MsgTypeTextMultipleFile is text with multiple file attachments.
	MsgTypeTextMultipleFile = 5

	// MsgTypeUnused is reserved for future use.
	MsgTypeUnused = 6

	// MsgTypeDeleted indicates the original message was deleted.
	MsgTypeDeleted = 7

	// MsgTypeNoteShared indicates a note was shared in the conversation.
	MsgTypeNoteShared = 8

	// MsgTypeNoteDeleted indicates a previously shared note was deleted.
	MsgTypeNoteDeleted = 9

	// MsgTypeNoteCreated indicates a new note was created and shared.
	MsgTypeNoteCreated = 10

	// MsgTypeNoteUpdated indicates a shared note was updated.
	MsgTypeNoteUpdated = 11

	// MsgTypeOriginalStamp is a legacy stamp type.
	MsgTypeOriginalStamp = 12

	// MsgTypeYesNo is a yes/no poll question (INTERNAL, for receiving only).
	MsgTypeYesNo = 13

	// MsgTypeYesNoReply is a response to a yes/no poll (INTERNAL, for receiving only).
	MsgTypeYesNoReply = 14

	// MsgTypeSelect is a multiple-choice poll (INTERNAL, for receiving only).
	MsgTypeSelect = 15

	// MsgTypeSelectReply is a response to a multiple-choice poll (INTERNAL, for receiving only).
	MsgTypeSelectReply = 16

	// MsgTypeTask is a task assignment poll (INTERNAL, for receiving only).
	MsgTypeTask = 17

	// MsgTypeTaskDone is a task completion response (INTERNAL, for receiving only).
	MsgTypeTaskDone = 18

	// MsgTypeYesNoClosed indicates a yes/no poll was closed (INTERNAL, for receiving only).
	MsgTypeYesNoClosed = 19

	// MsgTypeSelectClosed indicates a multiple-choice poll was closed (INTERNAL, for receiving only).
	MsgTypeSelectClosed = 20

	// MsgTypeTaskClosed indicates a task poll was closed (INTERNAL, for receiving only).
	MsgTypeTaskClosed = 21
)

// Wire message types for action stamps (polls/interactive messages).
// These constants are used when SENDING action stamps via Client.Send().
// Do NOT use the MsgType constants for sending - use these WireType constants instead.
// The API internally converts wire types to internal enum values.
// Formula: wireType = 500 + (internalType - 13)
//
// Example:
//
//	content := map[string]interface{}{
//		"question": "What do you think?",
//		"options": []string{"Good", "Bad", "Neutral"},
//	}
//	err := client.Send(roomID, direct.WireTypeSelect, content)
const (
	// WireTypeYesNo sends a yes/no poll (500 + 13 - 13 = 500).
	WireTypeYesNo = 500

	// WireTypeYesNoReply sends a yes/no poll response.
	WireTypeYesNoReply = 501

	// WireTypeSelect sends a multiple-choice poll (500 + 15 - 13 = 502).
	WireTypeSelect = 502

	// WireTypeSelectReply sends a multiple-choice poll response.
	WireTypeSelectReply = 503

	// WireTypeTask sends a task assignment (500 + 17 - 13 = 504).
	WireTypeTask = 504

	// WireTypeTaskDone sends a task completion response.
	WireTypeTaskDone = 505

	// WireTypeYesNoClosed closes a yes/no poll.
	WireTypeYesNoClosed = 506

	// WireTypeSelectClosed closes a multiple-choice poll.
	WireTypeSelectClosed = 507

	// WireTypeTaskClosed closes a task poll.
	WireTypeTaskClosed = 508
)
