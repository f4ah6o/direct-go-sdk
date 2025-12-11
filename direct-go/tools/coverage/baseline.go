package main

// jsMethodsByCategory contains all 82 RPC methods from direct-js organized by functional category
var jsMethodsByCategory = map[string][]string{
	"Session & Auth": {
		"start_notification",
		"reset_notification",
		"update_last_used_at",
		"get_joined_account_control_group",
		"accept_account_control_request",
		"reject_account_control_request",
		"get_account_control_requests",
	},
	"User Management": {
		"get_me",
		"get_users",
		"get_profile",
		"update_user",
		"update_profile",
		"get_presences",
		"get_user_identifiers",
		"add_friend",
		"delete_friend",
		"get_friends",
		"get_acquaintances",
	},
	"Domain Management": {
		"get_domains",
		"leave_domain",
		"get_domain_invites",
		"accept_domain_invite",
		"delete_domain_invite",
		"get_domain_users",
		"search_domain_users",
	},
	"Department Management": {
		"get_department_tree",
		"get_department_users",
		"get_department_user_count",
	},
	"Talk/Room Management": {
		"get_talks",
		"get_talk_statuses",
		"create_group_talk",
		"create_pair_talk",
		"update_group_talk",
		"add_talkers",
		"delete_talker",
		"add_favorite_talk",
		"delete_favorite_talk",
	},
	"Message Operations": {
		"get_messages",
		"create_message",
		"delete_message",
		"schedule_message",
		"reschedule_message",
		"get_scheduled_messages",
		"delete_scheduled_message",
		"search_messages",
		"search_messages_around_datetime",
		"add_favorite_message",
		"delete_favorite_message",
		"get_favorite_messages",
		"set_message_reaction",
		"reset_message_reaction",
		"get_message_reaction_users",
		"get_available_message_reactions",
		"get_read_status",
	},
	"File & Attachment Management": {
		"create_upload_auth",
		"get_file_preview",
		"create_file_preview",
		"delete_attachment",
		"get_attachments",
		"search_attachments",
	},
	"Note Management": {
		"get_note",
		"get_note_statuses",
		"delete_note",
		"lock_note",
		"unlock_note",
		"update_note_setting",
	},
	"Announcement Management": {
		"create_announcement",
		"get_announcements",
		"get_announcement_statuses",
		"update_announcement_status",
	},
	"Push Notification Management": {
		"disable_push_notification",
		"enable_push_notification",
	},
	"Conference/Call Management": {
		"get_conferences",
		"join_conference",
		"leave_conference",
		"reject_conference",
		"get_conference_participants",
	},
	"Miscellaneous": {
		"get_actions",
		"get_solutions",
		"get_stampsets",
		"get_direct_apps",
		"get_flow_notification_badges",
	},
}

// categoryOrder defines the display order of categories
var categoryOrder = []string{
	"Session & Auth",
	"User Management",
	"Domain Management",
	"Department Management",
	"Talk/Room Management",
	"Message Operations",
	"File & Attachment Management",
	"Note Management",
	"Announcement Management",
	"Push Notification Management",
	"Conference/Call Management",
	"Miscellaneous",
}

// getAllJSMethods returns a flat list of all JS methods
func getAllJSMethods() []string {
	var methods []string
	for _, category := range categoryOrder {
		if categoryMethods, ok := jsMethodsByCategory[category]; ok {
			methods = append(methods, categoryMethods...)
		}
	}
	return methods
}

// categorizeMethod returns the category name for a given method
func categorizeMethod(method string) string {
	for category, methods := range jsMethodsByCategory {
		for _, m := range methods {
			if m == method {
				return category
			}
		}
	}
	return "Unknown"
}

// getTotalMethodCount returns the total number of JS methods
func getTotalMethodCount() int {
	count := 0
	for _, methods := range jsMethodsByCategory {
		count += len(methods)
	}
	return count
}
