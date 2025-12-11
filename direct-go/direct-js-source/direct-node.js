!(function (y, S) {
  function a() {
    return G.__string_rec(this, "");
  }
  var t = {};
  function e(e, t) {
    function n() {}
    n.prototype = e;
    var i,
      r = new n();
    for (i in t) r[i] = t[i];
    return (
      t.toString !== Object.prototype.toString && (r.toString = t.toString),
      r
    );
  }
  function v(e, t) {
    this.r = new RegExp(e, t.split("u").join(""));
  }
  function T() {}
  function I() {}
  function N() {}
  function x() {}
  function A() {}
  function b() {}
  function M() {}
  function C(e) {
    Fn.call(this, e);
  }
  function W() {}
  function K() {}
  function V() {}
  function q() {}
  function Q() {}
  function X() {}
  function J() {}
  function Z() {}
  function $() {}
  function ee() {}
  function te() {}
  function ne() {}
  function ie() {}
  function re() {}
  function oe(e, t, n) {
    ((this.email = e), (this.pass = t), (this.accessToken = n));
  }
  function ae(e) {
    null != e &&
      ((this.id = e.id),
      (this.group = new Mi(e.group)),
      (this.profilePolicy = new Ci(e.profile_policy)));
  }
  function se(e) {
    null != e &&
      ((this.id = e.id),
      null != e.group && (this.group = new Mi(e.group)),
      null != e.profile_policy) &&
      (this.profilePolicy = new Ci(e.profile_policy));
  }
  function le(e) {
    null != e &&
      ((this.id = D.fromIntOrInt64(e.id)),
      (this.groupName = e.group_name),
      (this.groupOwnerName = e.group_owner_name),
      (this.groupOwnerEmail = e.group_owner_email),
      (this.hasDomainInvite = e.has_domain_invite),
      (this.updatedAt = e.updated_at),
      (this.version = e.version));
  }
  function ue(e) {
    ((this.unreadCount = 0),
      null != e &&
        ((this.domainId = e.domain_id),
        (this.unreadCount = null != e.unread_count ? e.unread_count : 0),
        (this.maxAnnouncementId = e.max_announcement_id),
        null != e.max_announcement &&
          (this.maxAnnouncement = new Ri(e.max_announcement)),
        (this.maxReadAnnouncementId = e.max_read_announcement_id),
        null == this.maxReadAnnouncementId) &&
        0 == this.unreadCount &&
        (this.maxReadAnnouncementId = this.maxAnnouncementId));
  }
  function ce(e) {
    null != e &&
      ((this.domainId = e.domain_id),
      (this.readAnnouncementIds = e.read_announcement_ids));
  }
  function _e(e) {
    if (
      ((this.notificationInvisibleTalkSound = !0),
      (this.notificationInvisibleTalkPopup = !1),
      (this.notificationVisibleTalkSound = !1),
      (this.notificationVisibleTalkPopup = !1),
      (this.notificationIncludingMessageContent = !0),
      (this.notificationPopupTime = 5),
      (this.talkAutoScroll = 1),
      (this.legacyNoteEnabled = !1),
      (this.leaveTalksEnabled = !1),
      (this.talksSortMode = "new_messages_first"),
      (this.talkRightPanelOpened = !0),
      (this.talkRightPanelWidth = 0),
      (this.keywordWatchingText = ""),
      (this.keywordWatchingEmphasis = !1),
      (this.keywordWatchingSelfMessage = !1),
      (this.keywordWatchingActionReply = !1),
      (this.userDataContainer = []),
      (this.language = "ja"),
      (this.conferenceIncomingSound = !0),
      (this.domainsSortMode = "name_order_asc"),
      (this.appLeftPanelOpened = !1),
      (this.appLeftPanelWidth = 0),
      (this.scheduledMessageDefaultTimeStr = "09:00"),
      null != e)
    ) {
      var t;
      if (
        (null != e.notificationInvisibleTalkSound &&
          (this.notificationInvisibleTalkSound =
            e.notificationInvisibleTalkSound),
        null != e.notificationInvisibleTalkPopup &&
          (this.notificationInvisibleTalkPopup =
            e.notificationInvisibleTalkPopup),
        null != e.notificationVisibleTalkSound &&
          (this.notificationVisibleTalkSound = e.notificationVisibleTalkSound),
        null != e.notificationVisibleTalkPopup &&
          (this.notificationVisibleTalkPopup = e.notificationVisibleTalkPopup),
        null != e.notificationIncludingMessageContent &&
          (this.notificationIncludingMessageContent =
            e.notificationIncludingMessageContent),
        null != e.notificationPopupTime &&
          (this.notificationPopupTime = e.notificationPopupTime),
        null != e.notificationVisibleTalkAutoScroll &&
          (this.talkAutoScroll = e.notificationVisibleTalkAutoScroll ? 1 : 3),
        null != e.talkAutoScroll && (this.talkAutoScroll = e.talkAutoScroll),
        null != e.legacyNoteEnabled &&
          (this.legacyNoteEnabled = e.legacyNoteEnabled),
        null != e.leaveTalksEnabled &&
          (this.leaveTalksEnabled = e.leaveTalksEnabled),
        null != e.talksSortMode && (this.talksSortMode = e.talksSortMode),
        (this.talkRightPanelOpened =
          null != e.talkRightPanelOpened
            ? e.talkRightPanelOpened
            : null == o.getItem("right_panel_opened")),
        null != e.talkRightPanelWidth
          ? (this.talkRightPanelWidth = e.talkRightPanelWidth)
          : ((t = P.parseInt(o.getItem("right_panel_width"))),
            (this.talkRightPanelWidth = null == t ? 0 : t)),
        null != e.keywordWatchingText &&
          (this.keywordWatchingText = e.keywordWatchingText),
        null != e.keywordWatchingEmphasis &&
          (this.keywordWatchingEmphasis = e.keywordWatchingEmphasis),
        null != e.keywordWatchingSelfMessage &&
          (this.keywordWatchingSelfMessage = e.keywordWatchingSelfMessage),
        null != e.keywordWatchingActionReply &&
          (this.keywordWatchingActionReply = e.keywordWatchingActionReply),
        null != e.userDataContainer)
      )
        for (
          var n = 0, i = G.__cast(e.userDataContainer, Array);
          n < i.length;
        ) {
          var r = i[n];
          (++n, this.userDataContainer.push(xi.fromJson(r)));
        }
      (null != e.language && (this.language = e.language),
        null != e.conferenceIncomingSound &&
          (this.conferenceIncomingSound = e.conferenceIncomingSound),
        null != e.domainsSortMode && (this.domainsSortMode = e.domainsSortMode),
        (this.appLeftPanelOpened =
          null != e.appLeftPanelOpened
            ? e.appLeftPanelOpened
            : null != o.getItem("app_left_panel_opened")),
        null != e.appLeftPanelWidth &&
          (this.appLeftPanelWidth = e.appLeftPanelWidth),
        null != e.scheduledMessageDefaultTimeStr &&
          (this.scheduledMessageDefaultTimeStr =
            e.scheduledMessageDefaultTimeStr));
    }
  }
  function he(e) {
    null != e &&
      ((this.userId = e.user_id),
      (this.id = e.conference_id),
      (this.domainId = e.domain_id),
      (this.talkId = e.talk_id),
      (this.messageId = e.message_id),
      (this.createdAt = e.created_at),
      (this.expiredAt = e.expired_at),
      (this.participants = e.participants),
      (this.skywayVersion = null == e.skyway_version ? 1 : e.skyway_version));
  }
  function de() {}
  function fe(e) {
    null != e &&
      ((this.domainId = e.domain_id),
      (this.departments = this.getDepartments(e.departments)));
  }
  function me(e) {
    null != e &&
      ((this.departmentId = e.department_id),
      (this.all = e.all),
      (this.partial = e.partial));
  }
  function pe(e) {
    null != e &&
      (this.departments = this.getDepartmentUserCounts(e.departments));
  }
  function ge(e) {
    if (null != e) {
      if (
        ((this.id = D.fromIntOrInt64(e.direct_app_id)),
        (this.uniqueKey = e.unique_key),
        (this.appUrl = e.app_url),
        (this.locales = new B()),
        null != e.locales)
      )
        for (var t = 0, n = L.fields(e.locales); t < n.length; ) {
          var i = n[t],
            r = (++t, L.field(e.locales, i)),
            o = this.locales,
            r = new Wi(r);
          null != Y[i] ? o.setReserved(i, r) : (o.h[i] = r);
        }
      if (((this.features = []), null != e.features))
        for (var a = 0, s = G.__cast(e.features, Array); a < s.length; ) {
          var l = s[a];
          (++a, this.features.push(new Ki(l)));
        }
    }
  }
  function ye(e) {
    null != e &&
      ((this.id = e.domain_id),
      (this.updatedAt = e.updated_at),
      (this.domainInfo = new Xi(e.domain)),
      (this.contract = new Hi(e.contract)),
      (this.profileDefinition = new Ir(e.profile_definition)),
      (this.setting = new Ji(e.setting)),
      (this.role = new $i(e.role)),
      (this.stampsetSetting = new Kr(e.stampset_setting)),
      (this.builtinStampSetting = new Li(e.builtin_stamp_setting)),
      (this.directAppsSetting = new Qi(e.direct_apps_setting)),
      (this.closed = !1));
  }
  function Se(e) {
    null != e &&
      ((this.id = e.domain_id),
      (this.name = e.domain_name),
      (this.accountControlRequestId = D.fromIntOrInt64(
        e.account_control_request_id,
      )),
      (this.updatedAt = e.updated_at));
  }
  function ve(e) {
    (ir.call(this, e),
      null != e &&
        ((this.domainId = e.domain_id),
        null != e.profile_contact &&
          (this.profileContact = G.__cast(e.profile_contact, Array).map(
            function (e) {
              return new wr(e);
            },
          )),
        (this.departments = e.departments),
        (this.allowSelectedForTwoPersonsGroupTalk =
          0 != e.allow_selected_for_two_persons_group_talk)));
  }
  function Te(e) {
    this.values = e;
  }
  function Ie(e) {
    null != e &&
      ((this.fileId = e.file_id),
      (this.url = e.get_url),
      (this.headers = new or(e.get_headers)));
  }
  function Ee() {}
  function we() {}
  function Ne(e) {
    null != e &&
      ((this.domainId = e.domain_id),
      (this.talkId = e.talk_id),
      (this.orderInFavorites = e.order_in_favorites),
      (this.favoriteVersion = e.favorite_version));
  }
  function Ae(e) {
    null != e &&
      ((this.messageId = e[0]), (this.talkId = e[1]), (this.fileId = e[2]));
  }
  function be(e) {
    null != e &&
      ((this.domainId = e.domain_id),
      (this.show = e.show),
      (this.updatedAt = e.updated_at),
      (this.version = e.version));
  }
  function Me() {}
  function Ce() {}
  function ke(e) {
    null != e &&
      ((this.conferenceId = e.conference_id),
      (this.roomName = e.room_name),
      (this.credential = e.credential),
      (this.mode = this.parseMode(e.mode)),
      (this.timestamp = e.timestamp),
      (this.skywayVersion = null == e.skyway_version ? 1 : e.skyway_version));
  }
  function De() {}
  function Re(e) {
    null != e &&
      ((this.talkId = e[0]),
      (this.messageId = e[1]),
      (this.isMention = !!e[2]));
  }
  function Oe(e) {
    ((this.id = e.message_id),
      (this.talkId = e.talk_id),
      (this.readUserIds = null == e.read_user_ids ? [] : e.read_user_ids),
      (this.unreadUserIds =
        null == e.unread_user_ids ? [] : e.unread_user_ids));
  }
  function Fe(e) {
    ((this.talkId = e.talk_id),
      (this.messageIds = e.message_ids),
      (this.mentionMessageIds =
        null == e.mention_message_ids ? [] : e.mention_message_ids),
      (this.readUserIds = null == e.read_user_ids ? [] : e.read_user_ids),
      (this.messageIdsExcludingUnreadCountTargets =
        e.message_ids_excluding_unread_count_targets),
      null == this.messageIdsExcludingUnreadCountTargets &&
        (this.messageIdsExcludingUnreadCountTargets = []));
  }
  function xe() {}
  function Ue() {}
  function Le(e, t, n) {
    null == e
      ? (this.readingLastMessageId = n)
      : ((this.messages = e.map(function (e) {
          return new Ur(e, t);
        })),
        (this.readingLastMessageId = 0 < e.length ? e[e.length - 1].id : null));
  }
  function Pe(e) {
    if (
      null != e &&
      ((this.total = e.total),
      (this.marker = e.marker),
      (this.nextMarker = e.next_marker),
      (this.contents = []),
      null != e.contents)
    )
      for (var t = 0, n = G.__cast(e.contents, Array); t < n.length; ) {
        var i = n[t];
        (++t, this.contents.push(new ar(i)));
      }
  }
  function Be(e) {
    null != e &&
      ((this.total = e.total),
      (this.marker = e.marker),
      (this.nextMarker = e.next_marker),
      (this.contents = []),
      null != e.contents) &&
      (this.contents = G.__cast(e.contents, Array).map(function (e) {
        return new cr(e);
      }));
  }
  function He(e) {
    null != e &&
      (null != e.domain_id_str &&
        (this.domainId = D.makeFromIdStr(e.domain_id_str)),
      null != e.talk_id_str && (this.talkId = D.makeFromIdStr(e.talk_id_str)),
      (this.keyword = e.keyword),
      null == e.search_type || e.search_type == Pr.toLocalizeText(Lr.MESSAGES)
        ? (this.searchType = Lr.MESSAGES)
        : (this.searchType = Lr.ATTACHMENTS),
      (this.sinceText = e.since),
      (this.untilText = e.until),
      null != e.from_id_str) &&
      (this.fromUserId = D.makeFromIdStr(e.from_id_str));
  }
  function Ge() {}
  function je(e) {
    null != e &&
      ((this.user = new fr(e.user)),
      (this.device = e.device),
      (this.notification = e.notification),
      (this.passwordExpiration = new Sr(e.password_expiration)),
      (this.configuration = new Bi(e.configuration)));
  }
  function ze(e) {
    null != e &&
      ((this.solutionId = e.solution_id),
      (this.name = e.name),
      null != e.link && (this.link = new Hr(e.link)),
      null != e.plugin) &&
      (this.plugin = new Gr(e.plugin));
  }
  function Ye() {}
  function We(e) {
    this.commonStampId = e;
  }
  function Ke(e, t, n) {
    ((this.stampsetId = e), (this.id = t), (this.illust = n));
  }
  function Ve(e) {
    null != e &&
      ((this.stampsetId = D.fromIntOrInt64(e.stampset_id)),
      (this.name = e.name),
      (this.icon = e.icon),
      (this.version = e.version));
  }
  function qe() {}
  function Qe(e, t) {
    ((this.talk = e), (this.status = t));
  }
  function Xe() {}
  function Je(e, t) {
    ((this.domainId = e.domain_id),
      (this.talkId = e.talk_id),
      (this.type = t ? br.fromRawValue(e.type) : Mr.None),
      (this.version = e.version));
  }
  function Ze(e) {
    ((this.unreadMentionCount = 0),
      (this.unreadCount = 0),
      null == e
        ? (this.pushNotificationSetting = new Ar())
        : ((this.id = e.talk_id),
          (this.unreadCount = null != e.unread_count ? e.unread_count : 0),
          (this.unreadMentionCount =
            null != e.unread_mention_count ? e.unread_mention_count : 0),
          (this.maxMessageId = e.max_message_id),
          (this.maxMessage =
            null != e.max_message ? new O(e.max_message) : null),
          (this.maxReadMessageId = e.max_read_message_id),
          (this.maxEveryoneReadMessageId = e.max_everyone_read_message_id),
          (this.talkOrderingTimestamp = e.talk_ordering_timestamp),
          (this.orderInFavorites = e.order_in_favorites),
          (this.favoriteVersion = e.favorite_version),
          (this.pushNotificationSetting = new Ar(e.push_notification))));
  }
  function $e() {}
  function et(e) {
    null != e &&
      ((this.talkId = e.talk_id),
      (this.maxEveryoneReadMessageId = e.max_everyone_read_message_id),
      (this.maxReadMessageId = h.option(e.max_read_message_id)));
  }
  function tt(e, t, n) {
    ((this.file = e),
      (this.dimension = to.createWithSize(t, n)),
      (this.auth = null));
  }
  function nt(e, t, n) {
    var i = new E(0, 0);
    ((this.domainId = i),
      (this.domainId = e),
      (this.oldSetting = t),
      (this.newSetting = n));
  }
  function it() {}
  function rt(e) {
    null != e &&
      (null != e.user_id && (this.userId = e.user_id),
      null != e.email && (this.email = e.email),
      null != e.sub_email && (this.subEmail = e.sub_email),
      null != e.group_alias && (this.groupAlias = e.group_alias),
      null != e.signin_id) &&
      (this.signinId = e.signin_id);
  }
  function ot(e, t) {
    ((this.lastUpdatedAt = t),
      null != e &&
        ((this.userId = e.user_id),
        (this.lastUsedAtOpt = h.option(e.last_used_at))));
  }
  function at(e) {
    this.success = e;
  }
  function st(e) {
    this.success = e;
  }
  function lt(e, t, n, i, r) {
    ((this.domainId = e),
      (this.talkId = t),
      (this.marker = n),
      (this.nextMarker = i),
      (this.messages = r));
  }
  function ut(e, t, n) {
    ((this.domainId = e),
      (this.talkId = t),
      (this.marker = n.marker),
      (this.nextMarker = n.nextMarker),
      (this.contents =
        null == n.contents
          ? []
          : G.__cast(n.contents, Array).map(function (e) {
              return new fo(e);
            })));
  }
  function ct(i, e) {
    var r = this;
    k.iter(e, function (e) {
      var e = new io(e),
        t = e.createdAt,
        n = (i.high - t.high) | 0;
      (0 == n && (n = w.ucompare(i.low, t.low)),
        (i.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1) < 0
          ? (r.messageCreatedAtNew = e)
          : (r.messageCreatedAtOld = e));
    });
  }
  function _t(e, t, n) {
    ((this.callerId = e), (this.error = t), (this.emitterKey = n));
  }
  function ht(e) {
    ((this.note = new po(e)), (this.emitterKey = null));
  }
  function dt(e) {
    this.compressedNote = new ro(e);
  }
  function ft(e) {
    this.result = e;
  }
  function mt(e) {
    this.note = new po(e);
  }
  function pt(e) {
    this.compressedNote = new ro(e);
  }
  function gt(e) {
    null != e &&
      ((this.marker = e.marker),
      (this.nextMarker = e.next_marker),
      (this.noteStatuses = Un.asArray(e.contents).map(function (e) {
        return new Ao(e);
      })));
  }
  function yt(e) {
    this.note = new po(e);
  }
  function St(e) {
    this.compressedNote = new ro(e);
  }
  function vt(e) {
    this.note = new po(e);
  }
  function Tt(e) {
    this.compressedNote = new ro(e);
  }
  function It() {}
  function Et(e) {
    null != e &&
      ((this.fileId = e.file_id),
      (this.status = e.status),
      (this.filePreviewFileId = e.file_preview_file_id),
      (this.url = e.url),
      (this.key = e.key));
  }
  function wt() {}
  function Nt(e, t) {
    ((this.talkId = e), (this.paneType = t));
  }
  function At() {}
  function bt() {}
  function Mt() {}
  function Ct() {}
  function kt(e) {
    ((this.message = new O(e.message)),
      (this.domainId = e.domain_id),
      (this.talkId = e.talk_id));
  }
  function Dt(e) {
    ((this.messageId = e.message_id),
      (this.domainId = e.domain_id),
      (this.talkId = e.talk_id));
  }
  function Rt() {}
  function Ot(e, t, n, i, r) {
    (null == r && (r = !1),
      (this.type = e),
      (this.content = t),
      (this.files = n),
      (this.createMessageAt = i),
      (this.hasError = r));
  }
  function Ft(e) {
    null != e &&
      ((this.noteId = new So(e.note_id)),
      (this.revision = e.revision),
      (this.userId = e.user_id),
      (this.deviceId = e.device_id),
      (this.expiredAt = e.expired_at));
  }
  function xt(e) {
    this.value = e;
  }
  function Ut(e) {
    this.value = e;
  }
  function Lt(e) {
    null != e &&
      ((this.noteId = new So(e.note_id)),
      (this.talkId = e.talk_id),
      (this.setting = new No(e.setting)),
      (this.noteRevisionSummary = new wo(e.note_revision_summary)));
  }
  function Pt(e) {
    null != e &&
      ((this.noteId = new So(e.note_id)),
      (this.talkId = e.talk_id),
      (this.setting = new No(e.setting)));
  }
  function Bt(e) {
    null != e &&
      ((this.noteId = new So(e.note_id)),
      (this.talkId = e.talk_id),
      (this.noteRevisionSummary = new wo(e.note_revision_summary)));
  }
  function Ht() {}
  function Gt(e) {
    null != e &&
      ((this.noteId = new So(e.note_id)),
      (this.revision = e.revision),
      (this.userId = e.user_id));
  }
  function jt(i) {
    this.reactionCounts = new B();
    var r = this;
    null != i &&
      ((this.messageId = i.message_id),
      (this.myReactionId = D.fromIntOrInt64(i.my_reaction_id)),
      k.iter(L.fields(i.reaction_counts), function (e) {
        var t = r.reactionCounts,
          n = L.field(i.reaction_counts, e);
        null != Y[e] ? t.setReserved(e, n) : (t.h[e] = n);
      }));
  }
  function zt(i) {
    this.reactionUserIds = new B();
    var r = this;
    null != i &&
      ((this.messageId = i.message_id),
      k.iter(L.fields(i.reaction_user_ids), function (e) {
        var t = L.field(i.reaction_user_ids, e),
          n = r.reactionUserIds;
        null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
      }));
  }
  function Yt() {}
  function Wt(e, t, n) {
    ((this.ignoreCase = t),
      (this.normalizeFunc = n),
      t && (e = e.toUpperCase()),
      (this.normalizedPattern = n(e)));
  }
  function Kt() {}
  function Vt() {}
  function qt() {}
  function Qt() {}
  function Xt() {}
  function Jt() {}
  function Zt() {}
  function $t(e) {
    ((this.rawValue = e),
      (this.post_url = e.post_url),
      (this.post_form = e.post_form),
      (this.put_url = e.put_url));
  }
  function en() {}
  function p() {}
  function tn() {}
  function nn(e, t) {
    ((this.updatedAt = e.updatedAt),
      (this.domainInfo = e.domainInfo),
      (this.contract = e.contract),
      (this.profileDefinition = e.profileDefinition),
      (this.setting = e.setting),
      (this.role = e.role),
      (this.closed = e.closed),
      (this.id = e.id),
      null != t && (this.talks = t));
  }
  function rn(e) {
    (Aa.call(this, e),
      (this.domainId = e.domainId),
      (this.profileContact = this.convertProfileContact(e.profileContact)),
      (this.departments = e.departments),
      (this.email = e.email),
      (this.subEmail = e.subEmail),
      (this.groupAlias = e.groupAlias),
      (this.signinId = e.signinId));
  }
  function on(e) {
    (Aa.call(this, e),
      (this.email = e.email),
      (this.signinId = e.signinId),
      (this.profiles = e.profiles),
      (this.departments = e.departments));
  }
  function an(e) {
    var t = e.id;
    ((this.id = "_" + t.high + "_" + t.low),
      (this.id_i64 = e.id),
      (this.name = e.name),
      (this.content_type = e.contentType),
      (this.content_size = e.contentSize),
      (this.url = e.url));
  }
  function sn(e, t, n) {
    ((this.name = e.name),
      (this.iconUrl = e.iconUrl),
      (this.updatedAt = e.updatedAt),
      (this.leftUsers = e.leftUsers),
      (this.groupTalkSettings = e.groupTalkSettings),
      (this.id = e.id),
      (this.type = e.type),
      (this.userIds = e.userIds),
      (this.users = t),
      (this.domain = n),
      (this.domainId = e.domainId));
  }
  function ln(e) {
    ((this.id = e.noteId.toString()), (this.id_i64 = e.noteId.value));
    var t = e.talkId,
      t =
        ((this.talkId = "_" + t.high + "_" + t.low),
        (this.talkId_i64 = e.talkId),
        ka.of(e.noteRevisionSummary));
    (Object.defineProperty(this, "noteRevisionSummary", { value: t }),
      Object.defineProperty(this, "title", {
        get: this.get_title,
        enumerable: !0,
      }),
      Object.defineProperty(this, "contentType", {
        get: this.get_contentType,
        enumerable: !0,
      }),
      Object.defineProperty(this, "contentSummary", {
        get: this.get_contentSummary,
        enumerable: !0,
      }),
      Object.defineProperty(this, "attachments", {
        get: this.get_attachments,
        enumerable: !0,
      }),
      Object.defineProperty(this, "lastUpdatedBy", {
        get: this.get_lastUpdatedBy,
      }),
      Object.defineProperty(this, "lastUpdatedAt", {
        get: this.get_lastUpdatedAt,
      }));
  }
  function un(e) {
    ((this.id = e.noteId.toString()), (this.id_i64 = e.noteId.value));
    var t = e.talkId,
      t =
        ((this.talkId = "_" + t.high + "_" + t.low),
        (this.talkId_i64 = e.talkId),
        (this.createdBy = e.createdBy),
        (this.createdAt = e.createdAt),
        ka.of(e.noteRevisionSummary));
    (Object.defineProperty(this, "noteRevisionSummary", { value: t }),
      Object.defineProperty(this, "title", {
        get: this.get_title,
        enumerable: !0,
      }),
      Object.defineProperty(this, "contentType", {
        get: this.get_contentType,
        enumerable: !0,
      }),
      Object.defineProperty(this, "contentSummary", {
        get: this.get_contentSummary,
        enumerable: !0,
      }),
      Object.defineProperty(this, "attachments", {
        get: this.get_attachments,
        enumerable: !0,
      }),
      Object.defineProperty(this, "lastUpdatedBy", {
        get: this.get_lastUpdatedBy,
      }),
      Object.defineProperty(this, "lastUpdatedAt", {
        get: this.get_lastUpdatedAt,
      }));
  }
  function cn() {}
  function _n() {}
  function hn() {}
  function dn() {}
  function fn() {}
  function mn() {}
  function pn() {}
  function gn() {}
  function yn() {}
  function Sn() {}
  var vn = (y.DirectAPI = function () {
      this.eventEmitter = new In();
    }),
    Tn =
      (((t.DirectAPI = vn).__name__ = ["DirectAPI"]),
      (vn.getInstance = function () {
        return (null == vn.instance && (vn.instance = new vn()), vn.instance);
      }),
      (vn.main = function () {}),
      (vn.prototype = {
        setOptions: function (e) {
          (null != e &&
            ((qs.host = e.host),
            (qs.endpoint = e.endpoint),
            (qs.accessToken = e.access_token),
            (qs.proxyURL = e.proxyURL),
            (qs.account = e.account),
            (qs.talkWithBot = e.talkWithBot),
            (qs.acceptableEventTimeDiff = e.acceptableEventTimeDiff),
            (qs.name = e.name),
            (qs.storagePath = e.storage_path),
            (qs.storageQuota = e.storage_quota),
            (qs.wsConfig = e.ws_config),
            null != e.internalLogger) &&
            F.setInternalLogger(e.internalLogger),
            F.d("current settings", qs),
            null == qs.accessToken &&
              F.e("Not enough parameters provided. I need a access token"));
        },
        announce: function (e, t) {
          var n = e.room;
          if (null != n) {
            ((n = n.split("_")),
              (n =
                2 < n.length
                  ? _.make(P.parseInt(n[1]), P.parseInt(n[2]))
                  : null));
            if (null == n || null == t) return;
            n = this.data.getTalk(n);
            if (null == n) return;
            i = n.domainId;
          } else
            var n = e.id.split("_"),
              i =
                2 < n.length
                  ? _.make(P.parseInt(n[1]), P.parseInt(n[2]))
                  : null;
          null != i && this.sendQueue.sendAnnouncement(i, t);
        },
        send: function (e, t) {
          ((e = e.room.split("_")),
            (e =
              2 < e.length
                ? _.make(P.parseInt(e[1]), P.parseInt(e[2]))
                : null));
          null != e && null != t && this.sendQueue.sendMessage(e, t);
        },
        topic: function (e, t) {
          ((e = e.room.split("_")),
            (e =
              2 < e.length
                ? _.make(P.parseInt(e[1]), P.parseInt(e[2]))
                : null));
          this.facade.sendNotification("Talk", Ii.UPDATE_FOR_HUBOT(e, t));
        },
        download: function (e, t, n) {
          var i,
            r = null,
            o = null;
          ("string" == typeof t
            ? (i = t)
            : ((i = t.url), (r = t.path), (o = t.name)),
            null == i
              ? n(null, new Error("target URL is required"))
              : (null == o && (o = ms.basename(i)),
                null == r && (r = ms.join(fs.tmpdir(), o)),
                this.facade.sendNotification(
                  "FilePath",
                  Xn.DOWNLOAD_PATH(i, r, n),
                )));
        },
        leave: function (e, t) {
          var n,
            i = this,
            e = e.room.split("_"),
            r =
              2 < e.length ? _.make(P.parseInt(e[1]), P.parseInt(e[2])) : null;
          ((n =
            null != t && 2 < (e = t.id.split("_")).length
              ? _.make(P.parseInt(e[1]), P.parseInt(e[2]))
              : null),
            Xa.delay(function () {
              i.facade.sendNotification("Talk", Ii.DELETE_FOR_HUBOT(r, n));
            }, 500));
        },
        userForId: function (e, t) {
          t = D.fromNullableIdStr(t);
          return this.hubotObject.userObjectByIdStr(t, e);
        },
        userObjects: function (e) {
          e = D.fromNullableIdStr(e);
          return this.hubotObject.userObjects(e);
        },
        talkObjects: function () {
          return this.hubotObject.talkObjects();
        },
        domainObjects: function () {
          return this.hubotObject.domainObjects();
        },
        parseInt64: function (e) {
          return ja.parseInt64(e);
        },
        stringifyInt64: function (e, t) {
          return ja.stringifyInt64(e, (t = null == t ? !1 : t));
        },
        listen: function () {
          ((this.facade = i.getInstance()),
            (this.api = G.__cast(this.facade.retrieveProxy("api"), zo)),
            (this.data = G.__cast(this.facade.retrieveProxy("dataStore"), qo)),
            (this.ID = new ja()),
            (this.hubotObject = G.__cast(
              this.facade.retrieveProxy("hubotObject"),
              xa,
            )),
            (this.sendQueue = G.__cast(
              this.facade.retrieveProxy("sendQueue"),
              Pa,
            )));
          var e = G.__cast(this.facade.retrieveProxy("fileService"), Xo);
          ((this.notes = new wa(new Na(this, this.api, e))),
            (this.mention = new An()),
            this.facade.startup());
        },
        emit: function (e, t, n, i) {
          this.eventEmitter.emit(e, t, n, i);
        },
        on: function (e, t) {
          return this.eventEmitter.on(e, t);
        },
        off: function (e, t) {
          return this.eventEmitter.removeListener(e, t);
        },
        getDomainInvites: function (e) {
          this.api._getDomainInvites(e);
        },
        acceptDomainInvite: function (e) {
          this.api.acceptDomainInvite(e);
        },
        __class__: vn,
      }),
      ((t.EReg = v).__name__ = ["EReg"]),
      (v.prototype = {
        match: function (e) {
          return (
            this.r.global && (this.r.lastIndex = 0),
            (this.r.m = this.r.exec(e)),
            (this.r.s = e),
            null != this.r.m
          );
        },
        matched: function (e) {
          if (null != this.r.m && 0 <= e && e < this.r.m.length)
            return this.r.m[e];
          throw new H("EReg::matched");
        },
        matchedPos: function () {
          if (null == this.r.m) throw new H("No string matched");
          return { pos: this.r.m.index, len: this.r.m[0].length };
        },
        matchSub: function (e, t, n) {
          var i, r;
          return (
            null == n && (n = -1),
            this.r.global
              ? ((this.r.lastIndex = t),
                (i = this.r),
                (r = n < 0 ? e : U.substr(e, 0, t + n)),
                (this.r.m = i.exec(r)),
                (i = null != this.r.m) && (this.r.s = e),
                i)
              : ((r = this.match(
                  n < 0 ? U.substr(e, t, null) : U.substr(e, t, n),
                )) && ((this.r.s = e), (this.r.m.index += t)),
                r)
          );
        },
        split: function (e) {
          var t = "#__delim__#";
          return e.replace(this.r, t).split(t);
        },
        map: function (e, t) {
          for (var n = 0, i = ""; !(n >= e.length); ) {
            if (!this.matchSub(e, n)) {
              i += P.string(U.substr(e, n, null));
              break;
            }
            var r = this.matchedPos(),
              i =
                (i += P.string(U.substr(e, n, r.pos - n))) + P.string(t(this)),
              n =
                0 == r.len
                  ? ((i += P.string(U.substr(e, r.pos, 1))), r.pos + 1)
                  : r.pos + r.len;
            if (!this.r.global) break;
          }
          return (
            !this.r.global &&
              0 < n &&
              n < e.length &&
              (i += P.string(U.substr(e, n, null))),
            i
          );
        },
        __class__: v,
      }),
      require("events").EventEmitter),
    In = function () {
      Tn.call(this);
    },
    U =
      (((t.EventEmitter = In).__name__ = ["EventEmitter"]),
      (In.__super__ = Tn),
      (In.prototype = e(Tn.prototype, { __class__: In })),
      function () {}),
    k =
      (((t.HxOverrides = U).__name__ = ["HxOverrides"]),
      (U.dateStr = function (e) {
        var t = e.getMonth() + 1,
          n = e.getDate(),
          i = e.getHours(),
          r = e.getMinutes(),
          o = e.getSeconds();
        return (
          e.getFullYear() +
          "-" +
          (t < 10 ? "0" + t : "" + t) +
          "-" +
          (n < 10 ? "0" + n : "" + n) +
          " " +
          (i < 10 ? "0" + i : "" + i) +
          ":" +
          (r < 10 ? "0" + r : "" + r) +
          ":" +
          (o < 10 ? "0" + o : "" + o)
        );
      }),
      (U.strDate = function (e) {
        switch (e.length) {
          case 8:
            var t = e.split(":"),
              n = new Date();
            return (
              n.setTime(0),
              n.setUTCHours(t[0]),
              n.setUTCMinutes(t[1]),
              n.setUTCSeconds(t[2]),
              n
            );
          case 10:
            t = e.split("-");
            return new Date(t[0], t[1] - 1, t[2], 0, 0, 0);
          case 19:
            ((n = e.split(" ")), (t = n[0].split("-")), (n = n[1].split(":")));
            return new Date(t[0], t[1] - 1, t[2], n[0], n[1], n[2]);
          default:
            throw new H("Invalid date format : " + e);
        }
      }),
      (U.cca = function (e, t) {
        e = e.charCodeAt(t);
        if (e == e) return e;
      }),
      (U.substr = function (e, t, n) {
        if (null == n) n = e.length;
        else if (n < 0) {
          if (0 != t) return "";
          n = e.length + n;
        }
        return e.substr(t, n);
      }),
      (U.remove = function (e, t) {
        t = e.indexOf(t);
        return -1 != t && (e.splice(t, 1), !0);
      }),
      (U.iter = function (e) {
        return {
          cur: 0,
          arr: e,
          hasNext: function () {
            return this.cur < this.arr.length;
          },
          next: function () {
            return this.arr[this.cur++];
          },
        };
      }),
      function () {}),
    En =
      (((t.Lambda = k).__name__ = ["Lambda"]),
      (k.array = function (e) {
        for (var t = [], n = Ns(e)(); n.hasNext(); ) {
          var i = n.next();
          t.push(i);
        }
        return t;
      }),
      (k.map = function (e, t) {
        for (var n = new En(), i = Ns(e)(); i.hasNext(); ) {
          var r = i.next();
          n.add(t(r));
        }
        return n;
      }),
      (k.mapi = function (e, t) {
        for (var n = new En(), i = 0, r = Ns(e)(); r.hasNext(); ) {
          var o = r.next();
          n.add(t(i++, o));
        }
        return n;
      }),
      (k.has = function (e, t) {
        for (var n = Ns(e)(); n.hasNext(); ) if (n.next() == t) return !0;
        return !1;
      }),
      (k.exists = function (e, t) {
        for (var n = Ns(e)(); n.hasNext(); ) if (t(n.next())) return !0;
        return !1;
      }),
      (k.foreach = function (e, t) {
        for (var n = Ns(e)(); n.hasNext(); ) if (!t(n.next())) return !1;
        return !0;
      }),
      (k.iter = function (e, t) {
        for (var n = Ns(e)(); n.hasNext(); ) t(n.next());
      }),
      (k.filter = function (e, t) {
        for (var n = new En(), i = Ns(e)(); i.hasNext(); ) {
          var r = i.next();
          t(r) && n.add(r);
        }
        return n;
      }),
      (k.fold = function (e, t, n) {
        for (var i = Ns(e)(); i.hasNext(); ) n = t(i.next(), n);
        return n;
      }),
      (k.count = function (e, t) {
        var n = 0;
        if (null == t)
          for (var i = Ns(e)(); i.hasNext(); ) {
            i.next();
            ++n;
          }
        else for (var r = Ns(e)(); r.hasNext(); ) t(r.next()) && ++n;
        return n;
      }),
      (k.find = function (e, t) {
        for (var n = Ns(e)(); n.hasNext(); ) {
          var i = n.next();
          if (t(i)) return i;
        }
        return null;
      }),
      function () {
        this.length = 0;
      }),
    wn =
      (((t.List = En).__name__ = ["List"]),
      (En.prototype = {
        add: function (e) {
          e = new wn(e, null);
          (null == this.h ? (this.h = e) : (this.q.next = e),
            (this.q = e),
            this.length++);
        },
        isEmpty: function () {
          return null == this.h;
        },
        remove: function (e) {
          for (var t = null, n = this.h; null != n; ) {
            if (n.item == e)
              return (
                null == t ? (this.h = n.next) : (t.next = n.next),
                this.q == n && (this.q = t),
                this.length--,
                !0
              );
            n = (t = n).next;
          }
          return !1;
        },
        iterator: function () {
          return new Nn(this.h);
        },
        map: function (e) {
          for (var t = new En(), n = this.h; null != n; ) {
            var i = n.item,
              n = n.next;
            t.add(e(i));
          }
          return t;
        },
        __class__: En,
      }),
      function (e, t) {
        ((this.item = e), (this.next = t));
      }),
    Nn =
      (((t["_List.ListNode"] = wn).__name__ = ["_List", "ListNode"]),
      (wn.prototype = { __class__: wn }),
      function (e) {
        this.head = e;
      }),
    An =
      (((t["_List.ListIterator"] = Nn).__name__ = ["_List", "ListIterator"]),
      (Nn.prototype = {
        hasNext: function () {
          return null != this.head;
        },
        next: function () {
          var e = this.head.item;
          return ((this.head = this.head.next), e);
        },
        __class__: Nn,
      }),
      (Math.__name__ = ["Math"]),
      function () {
        this.forAll = "ALL";
      }),
    bn =
      (((t.MentionUtil = An).__name__ = ["MentionUtil"]),
      (An.extractMentionParts = function (e) {
        return Gs.extractMentionParts(e).map(function (e) {
          return new bn(e);
        });
      }),
      (An.removeMentionMarkup = function (e) {
        return Gs.removeMentionMarkup(e);
      }),
      (An.prototype = {
        markup: function (e) {
          var t, n;
          return e && e.id && (e.name || e.displayName)
            ? ((t = D.makeFromIdStr(e.id)),
              (n = e.displayName || e.name),
              "{@:" +
                Va.toString(t) +
                "," +
                (Gs.codePointCount(n) + 1) +
                "}@" +
                n)
            : "string" == typeof e && "ALL" == e
              ? "{@:ALL,4}@ALL"
              : null;
        },
        __class__: An,
      }),
      function (e) {
        ((this.start = e.start),
          (this.end = e.end),
          (this.userIdNumStr = e.userIdNumStr),
          (this.mentionText = e.mentionText));
      }),
    Mn =
      (((t.MentionPart = bn).__name__ = ["MentionPart"]),
      (bn.prototype = {
        isALL: function () {
          return "ALL" == this.userIdNumStr;
        },
        __class__: bn,
      }),
      ((t.ObjectHelper = T).__name__ = ["ObjectHelper"]),
      (T.deepCopy = function (e) {
        return null == e ? null : JSON.parse(JSON.stringify(e));
      }),
      function (e, t) {
        ((this.basetime = 0),
          (this.count = 0),
          (this.max = e),
          (this.interval = t));
      }),
    L =
      (((t.RateLimit = Mn).__name__ = ["RateLimit"]),
      (Mn.prototype = {
        _now: function () {
          return new Date().getTime();
        },
        calculateRetryAfter: function (e) {
          e =
            0 < this.basetime
              ? this.interval - (e - this.basetime)
              : this.interval;
          return Math.ceil(e / 1e3);
        },
        createError: function (e) {
          var t =
            this.max +
            " calls every " +
            Math.floor(this.interval / 1e3) +
            " sec";
          return Ha.createTooManyRequestsError(t, this.calculateRetryAfter(e));
        },
        apply: function (i) {
          var r = this;
          return new Promise(function (e, t) {
            var n = r._now();
            if (0 < r.max && n - r.basetime > r.interval)
              ((r.basetime = n), (r.count = 1));
            else {
              if (!(0 < r.max && r.count < r.max))
                return void t(r.createError(n));
              r.count++;
            }
            i().then(e, t);
          });
        },
        __class__: Mn,
      }),
      function () {}),
    P =
      (((t.Reflect = L).__name__ = ["Reflect"]),
      (L.field = function (e, t) {
        try {
          return e[t];
        } catch (e) {
          return null;
        }
      }),
      (L.fields = function (e) {
        var t = [];
        if (null != e) {
          var n,
            i = Object.prototype.hasOwnProperty;
          for (n in e)
            "__id__" != n && "hx__closures__" != n && i.call(e, n) && t.push(n);
        }
        return t;
      }),
      (L.isFunction = function (e) {
        return "function" == typeof e && !(e.__name__ || e.__ename__);
      }),
      (L.compare = function (e, t) {
        return e == t ? 0 : t < e ? 1 : -1;
      }),
      (L.isObject = function (e) {
        var t;
        return (
          null != e &&
          ("string" == (t = typeof e) ||
            ("object" == t && null == e.__enum__) ||
            ("function" == t && null != (e.__name__ || e.__ename__)))
        );
      }),
      (L.isEnumValue = function (e) {
        return null != e && null != e.__enum__;
      }),
      (L.deleteField = function (e, t) {
        return (
          !!Object.prototype.hasOwnProperty.call(e, t) && (delete e[t], !0)
        );
      }),
      (L.copy = function (e) {
        for (var t = {}, n = 0, i = L.fields(e); n < i.length; ) {
          var r = i[n];
          (++n, (t[r] = L.field(e, r)));
        }
        return t;
      }),
      function () {}),
    Cn =
      (((t.Std = P).__name__ = ["Std"]),
      (P.string = function (e) {
        return G.__string_rec(e, "");
      }),
      (P.parseInt = function (e) {
        var t = parseInt(e, 10);
        return (
          0 != t ||
            (120 != U.cca(e, 1) && 88 != U.cca(e, 1)) ||
            (t = parseInt(e)),
          isNaN(t) ? null : t
        );
      }),
      function () {
        this.b = "";
      }),
    n =
      (((t.StringBuf = Cn).__name__ = ["StringBuf"]),
      (Cn.prototype = {
        toString: function () {
          return this.b;
        },
        __class__: Cn,
      }),
      ((t.StringTools = I).__name__ = ["StringTools"]),
      (I.htmlEscape = function (e, t) {
        return (
          (e = e
            .split("&")
            .join("&amp;")
            .split("<")
            .join("&lt;")
            .split(">")
            .join("&gt;")),
          t ? e.split('"').join("&quot;").split("'").join("&#039;") : e
        );
      }),
      (I.startsWith = function (e, t) {
        return e.length >= t.length && U.substr(e, 0, t.length) == t;
      }),
      (I.isSpace = function (e, t) {
        e = U.cca(e, t);
        return (8 < e && e < 14) || 32 == e;
      }),
      (I.ltrim = function (e) {
        for (var t = e.length, n = 0; n < t && I.isSpace(e, n); ) ++n;
        return 0 < n ? U.substr(e, n, t - n) : e;
      }),
      (I.rtrim = function (e) {
        for (var t = e.length, n = 0; n < t && I.isSpace(e, t - n - 1); ) ++n;
        return 0 < n ? U.substr(e, 0, t - n) : e;
      }),
      (I.trim = function (e) {
        return I.ltrim(I.rtrim(e));
      }),
      (I.replace = function (e, t, n) {
        return e.split(t).join(n);
      }),
      (I.hex = function (e, t) {
        for (
          var n = "";
          (n = "0123456789ABCDEF".charAt(15 & e) + n), 0 < (e >>>= 4);
        );
        if (null != t) for (; n.length < t; ) n = "0" + n;
        return n;
      }),
      ((t.TextHelper = N).__name__ = ["TextHelper"]),
      (N.slice = function (e, n) {
        for (
          var i = [],
            t = function (e) {
              for (; e.length > n; ) {
                var t = U.substr(e, 0, n);
                (i.push(t), (e = U.substr(e, n, null)));
              }
              0 < e.length && i.push(e);
            },
            r = "",
            o = e.split("\n");
          0 < o.length;
        ) {
          var a = o.shift();
          (r.length + a.length > n && (t(r), (r = "")),
            0 < r.length && (r += "\n"),
            (r += a));
        }
        return (t(r), i);
      }),
      (t.ValueType = {
        __ename__: ["ValueType"],
        __constructs__: [
          "TNull",
          "TInt",
          "TFloat",
          "TBool",
          "TObject",
          "TFunction",
          "TClass",
          "TEnum",
          "TUnknown",
        ],
      })),
    kn =
      ((n.TNull = ["TNull", 0]),
      (n.TNull.toString = a),
      ((n.TNull.__enum__ = n).TInt = ["TInt", 1]),
      (n.TInt.toString = a),
      ((n.TInt.__enum__ = n).TFloat = ["TFloat", 2]),
      (n.TFloat.toString = a),
      ((n.TFloat.__enum__ = n).TBool = ["TBool", 3]),
      (n.TBool.toString = a),
      ((n.TBool.__enum__ = n).TObject = ["TObject", 4]),
      (n.TObject.toString = a),
      ((n.TObject.__enum__ = n).TFunction = ["TFunction", 5]),
      (n.TFunction.toString = a),
      ((n.TFunction.__enum__ = n).TClass = function (e) {
        e = ["TClass", 6, e];
        return ((e.__enum__ = n), (e.toString = a), e);
      }),
      (n.TEnum = function (e) {
        e = ["TEnum", 7, e];
        return ((e.__enum__ = n), (e.toString = a), e);
      }),
      (n.TUnknown = ["TUnknown", 8]),
      (n.TUnknown.toString = a),
      ((n.TUnknown.__enum__ = n).__empty_constructs__ = [
        n.TNull,
        n.TInt,
        n.TFloat,
        n.TBool,
        n.TObject,
        n.TFunction,
        n.TUnknown,
      ]),
      ((t.Type = x).__name__ = ["Type"]),
      (x.getClassName = function (e) {
        e = e.__name__;
        return null == e ? null : e.join(".");
      }),
      (x.getEnumName = function (e) {
        return e.__ename__.join(".");
      }),
      (x.resolveClass = function (e) {
        e = t[e];
        return null != e && e.__name__ ? e : null;
      }),
      (x.resolveEnum = function (e) {
        e = t[e];
        return null != e && e.__ename__ ? e : null;
      }),
      (x.createInstance = function (e, t) {
        switch (t.length) {
          case 0:
            return new e();
          case 1:
            return new e(t[0]);
          case 2:
            return new e(t[0], t[1]);
          case 3:
            return new e(t[0], t[1], t[2]);
          case 4:
            return new e(t[0], t[1], t[2], t[3]);
          case 5:
            return new e(t[0], t[1], t[2], t[3], t[4]);
          case 6:
            return new e(t[0], t[1], t[2], t[3], t[4], t[5]);
          case 7:
            return new e(t[0], t[1], t[2], t[3], t[4], t[5], t[6]);
          case 8:
            return new e(t[0], t[1], t[2], t[3], t[4], t[5], t[6], t[7]);
          case 9:
            return new e(t[0], t[1], t[2], t[3], t[4], t[5], t[6], t[7], t[8]);
          case 10:
            return new e(
              t[0],
              t[1],
              t[2],
              t[3],
              t[4],
              t[5],
              t[6],
              t[7],
              t[8],
              t[9],
            );
          case 11:
            return new e(
              t[0],
              t[1],
              t[2],
              t[3],
              t[4],
              t[5],
              t[6],
              t[7],
              t[8],
              t[9],
              t[10],
            );
          case 12:
            return new e(
              t[0],
              t[1],
              t[2],
              t[3],
              t[4],
              t[5],
              t[6],
              t[7],
              t[8],
              t[9],
              t[10],
              t[11],
            );
          case 13:
            return new e(
              t[0],
              t[1],
              t[2],
              t[3],
              t[4],
              t[5],
              t[6],
              t[7],
              t[8],
              t[9],
              t[10],
              t[11],
              t[12],
            );
          case 14:
            return new e(
              t[0],
              t[1],
              t[2],
              t[3],
              t[4],
              t[5],
              t[6],
              t[7],
              t[8],
              t[9],
              t[10],
              t[11],
              t[12],
              t[13],
            );
          default:
            throw new H("Too many arguments");
        }
      }),
      (x.createEmptyInstance = function (e) {
        function t() {}
        return ((t.prototype = e.prototype), new t());
      }),
      (x.createEnum = function (e, t, n) {
        var i = L.field(e, t);
        if (null == i) throw new H("No such constructor " + t);
        if (L.isFunction(i)) {
          if (null == n) throw new H("Constructor " + t + " need parameters");
          return i.apply(e, n);
        }
        if (null != n && 0 != n.length)
          throw new H("Constructor " + t + " does not need parameters");
        return i;
      }),
      (x.typeof = function (e) {
        var t;
        switch (typeof e) {
          case "boolean":
            return n.TBool;
          case "function":
            return e.__name__ || e.__ename__ ? n.TObject : n.TFunction;
          case "number":
            return Math.ceil(e) == e % 2147483648 ? n.TInt : n.TFloat;
          case "object":
            return null == e
              ? n.TNull
              : null != (t = e.__enum__)
                ? n.TEnum(t)
                : null != (t = G.getClass(e))
                  ? n.TClass(t)
                  : n.TObject;
          case "string":
            return n.TClass(String);
          case "undefined":
            return n.TNull;
          default:
            return n.TUnknown;
        }
      }),
      (x.enumEq = function (e, t) {
        if (e != t)
          try {
            if (e[0] != t[0]) return !1;
            for (var n = 2, i = e.length; n < i; ) {
              var r = n++;
              if (!x.enumEq(e[r], t[r])) return !1;
            }
            var o = e.__enum__;
            if (o != t.__enum__ || null == o) return !1;
          } catch (e) {
            return !1;
          }
        return !0;
      }),
      ((t["puremvc.interfaces.IFacade"] = A).__name__ = [
        "puremvc",
        "interfaces",
        "IFacade",
      ]),
      (A.prototype = { __class__: A }),
      function () {
        (kn.instance = this).initializeFacade();
      }),
    i =
      (((t["puremvc.patterns.facade.Facade"] = kn).__name__ = [
        "puremvc",
        "patterns",
        "facade",
        "Facade",
      ]),
      (kn.__interfaces__ = [A]),
      (kn.getInstance = function () {
        return (null == kn.instance && (kn.instance = new kn()), kn.instance);
      }),
      (kn.prototype = {
        initializeFacade: function () {
          (this.initializeModel(),
            this.initializeController(),
            this.initializeView());
        },
        initializeController: function () {
          null == this.controller && (this.controller = vs.getInstance());
        },
        initializeModel: function () {
          null == this.model && (this.model = Ts.getInstance());
        },
        initializeView: function () {
          null == this.view && (this.view = Is.getInstance());
        },
        registerCommand: function (e, t) {
          this.controller.registerCommand(e, t);
        },
        removeCommand: function (e) {
          this.controller.removeCommand(e);
        },
        hasCommand: function (e) {
          return this.controller.hasCommand(e);
        },
        registerProxy: function (e) {
          this.model.registerProxy(e);
        },
        retrieveProxy: function (e) {
          return this.model.retrieveProxy(e);
        },
        removeProxy: function (e) {
          var t = null;
          return (t = null != this.model ? this.model.removeProxy(e) : t);
        },
        hasProxy: function (e) {
          return this.model.hasProxy(e);
        },
        registerMediator: function (e) {
          null != this.view && this.view.registerMediator(e);
        },
        retrieveMediator: function (e) {
          return this.view.retrieveMediator(e);
        },
        removeMediator: function (e) {
          var t = null;
          return (t = null != this.view ? this.view.removeMediator(e) : t);
        },
        hasMediator: function (e) {
          return this.view.hasMediator(e);
        },
        sendNotification: function (e, t, n) {
          this.notifyObservers(new Es(e, t, n));
        },
        notifyObservers: function (e) {
          null != this.view && this.view.notifyObservers(e);
        },
        __class__: kn,
      }),
      function () {
        kn.call(this);
      }),
    Dn =
      (((t["albero.AppFacade"] = i).__name__ = ["albero", "AppFacade"]),
      (i.getInstance = function () {
        return (
          null == kn.instance && (kn.instance = new i()),
          G.__cast(kn.instance, i)
        );
      }),
      (i.__super__ = kn),
      (i.prototype = e(kn.prototype, {
        initializeFacade: function () {
          kn.prototype.initializeFacade.call(this);
        },
        initializeModel: function () {
          kn.prototype.initializeModel.call(this);
          var e = [];
          (e = (e = e.concat([new xa(), new Pa(), new Ua()])).concat([
            new sa(),
            new Ko(),
            new qo(),
            new Vo(),
            new Sa(),
            new jo(),
            new zo(),
            new Ia(),
            new Yo(),
            new Wo(),
            Jo.newInstance(),
            new fa(),
            new ga(),
            new ea(),
            new Qo(),
            new ya(),
            new da(),
            new aa(),
            new Ta(),
            new va(),
            new oa(),
            Po.newInstance(),
          ])).push(Ho.newInstance());
          for (var t = 0; t < e.length; ) {
            var n = e[t];
            (++t, this.registerProxy(n));
          }
          for (var i = 0; i < e.length; ) {
            var r = e[i];
            (++i, this.autoBind(r));
          }
        },
        initializeView: function () {
          kn.prototype.initializeView.call(this);
          for (var e = 0, t = this.getMediators(); e < t.length; ) {
            var n = t[e];
            (++e, this.registerMediator(n));
          }
        },
        getMediators: function () {
          return [new Ra()];
        },
        initializeController: function () {
          kn.prototype.initializeController.call(this);
          for (
            var e = [
                Vn,
                Si,
                vi,
                ci,
                yi,
                Ti,
                ti,
                li,
                Zn,
                pi,
                wi,
                Ei,
                Ni,
                Ai,
                fi,
                Gn,
                ni,
                Wn,
                zn,
                ai,
                Jn,
                hi,
                ri,
              ],
              t = (e.push(Qn), 0);
            t < e.length;
          ) {
            var n = e[t],
              i =
                (++t,
                I.replace(x.getClassName(n).split(".").pop(), "Command", ""));
            this.registerCommand(i, n);
          }
        },
        registerMediator: function (e) {
          (this.autoBind(e), kn.prototype.registerMediator.call(this, e));
        },
        autoBind: function (e) {
          for (
            var t = _s.getFields(null == e ? null : G.getClass(e)),
              n = 0,
              i = L.fields(t);
            n < i.length;
          ) {
            var r = i[n],
              o = (++n, L.field(t, r));
            Object.prototype.hasOwnProperty.call(o, "inject") &&
              (null == (o = this.retrieveProxy(r))
                ? F.v(
                    "[autoBind:Error]" +
                      x.getClassName(null == e ? null : G.getClass(e)) +
                      "." +
                      r,
                  )
                : (F.v(
                    "[autoBind] " +
                      x.getClassName(null == e ? null : G.getClass(e)) +
                      "." +
                      r +
                      " <= " +
                      x.getClassName(null == o ? null : G.getClass(o)),
                  ),
                  (e[r] = o)));
          }
        },
        startup: function () {
          G.__cast(this.retrieveProxy("appState"), Ko).start();
        },
        __class__: i,
      })),
      (t["albero.AppState"] = {
        __ename__: ["albero", "AppState"],
        __constructs__: ["Active", "Inactive"],
      })),
    Rn =
      ((Dn.Active = ["Active", 0]),
      (Dn.Active.toString = a),
      ((Dn.Active.__enum__ = Dn).Inactive = ["Inactive", 1]),
      (Dn.Inactive.toString = a),
      ((Dn.Inactive.__enum__ = Dn).__empty_constructs__ = [
        Dn.Active,
        Dn.Inactive,
      ]),
      (t["albero.ConnectionStatus"] = {
        __ename__: ["albero", "ConnectionStatus"],
        __constructs__: [
          "Ok",
          "Error",
          "ConcurrentAccessError",
          "ForcibliyClosedError",
        ],
      })),
    On =
      ((Rn.Ok = ["Ok", 0]),
      (Rn.Ok.toString = a),
      ((Rn.Ok.__enum__ = Rn).Error = ["Error", 1]),
      (Rn.Error.toString = a),
      ((Rn.Error.__enum__ = Rn).ConcurrentAccessError = [
        "ConcurrentAccessError",
        2,
      ]),
      (Rn.ConcurrentAccessError.toString = a),
      ((Rn.ConcurrentAccessError.__enum__ = Rn).ForcibliyClosedError = [
        "ForcibliyClosedError",
        3,
      ]),
      (Rn.ForcibliyClosedError.toString = a),
      ((Rn.ForcibliyClosedError.__enum__ = Rn).__empty_constructs__ = [
        Rn.Ok,
        Rn.Error,
        Rn.ConcurrentAccessError,
        Rn.ForcibliyClosedError,
      ]),
      ((t["albero.ContentTypeResolver"] = b).__name__ = [
        "albero",
        "ContentTypeResolver",
      ]),
      (b.prototype = { __class__: b }),
      ((t["albero.DefaultContentTypeResolver"] = M).__name__ = [
        "albero",
        "DefaultContentTypeResolver",
      ]),
      (M.__interfaces__ = [b]),
      (M.getContentTypeByFileName = function (e, t) {
        if (Ln.isNotEmpty(t))
          return "csv" == M.getExtension(e) ? "text/csv" : t;
        switch (M.getExtension(e)) {
          case "csv":
            return "text/csv";
          case "doc":
            return "application/msword";
          case "docm":
            return "application/vnd.ms-word.document.macroenabled.12";
          case "docx":
            return "application/vnd.openxmlformats-officedocument.wordprocessingml.document";
          case "gif":
            return "image/gif";
          case "htm":
          case "html":
            return "text/html";
          case "jpeg":
          case "jpg":
            return "image/jpeg";
          case "lzh":
            return "application/x-lzh";
          case "md":
            return "text/markdown";
          case "pdf":
            return "application/pdf";
          case "png":
            return "image/png";
          case "ppt":
            return "application/vnd.ms-powerpoint";
          case "pptm":
            return "application/vnd.ms-powerpoint.presentation.macroenabled.12";
          case "pptx":
            return "application/vnd.openxmlformats-officedocument.presentationml.presentation";
          case "rar":
            return "application/vnd.rar";
          case "svg":
            return "image/svg+xml";
          case "txt":
            return "text/plain";
          case "webp":
            return "image/webp";
          case "xls":
            return "application/vnd.ms-excel";
          case "xlsb":
            return "application/vnd.ms-excel.sheet.binary.macroenabled.12";
          case "xlsm":
            return "application/vnd.ms-excel.sheet.macroenabled.12";
          case "xlsx":
            return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet";
          case "xml":
            return "text/xml";
          default:
            return "application/octet-stream";
        }
      }),
      (M.getExtension = function (e) {
        var t;
        return null == e || 0 == (t = (e = e.split(".")).length)
          ? ""
          : e[t - 1];
      }),
      (M.prototype = {
        resolve: function (e) {
          return M.getContentTypeByFileName(e.name, e.type);
        },
        __class__: M,
      }),
      function () {}),
    Fn =
      (((t["haxe._Unserializer.DefaultResolver"] = On).__name__ = [
        "haxe",
        "_Unserializer",
        "DefaultResolver",
      ]),
      (On.prototype = {
        resolveClass: function (e) {
          return x.resolveClass(e);
        },
        resolveEnum: function (e) {
          return x.resolveEnum(e);
        },
        __class__: On,
      }),
      function (e) {
        ((this.buf = e),
          (this.length = e.length),
          (this.pos = 0),
          (this.scache = []),
          (this.cache = []));
        e = Fn.DEFAULT_RESOLVER;
        (null == e && ((e = new On()), (Fn.DEFAULT_RESOLVER = e)),
          (this.resolver = e));
      }),
    D =
      (((t["haxe.Unserializer"] = Fn).__name__ = ["haxe", "Unserializer"]),
      (Fn.initCodes = function () {
        for (var e = [], t = 0, n = Fn.BASE64.length; t < n; ) {
          var i = t++;
          e[Fn.BASE64.charCodeAt(i)] = i;
        }
        return e;
      }),
      (Fn.prototype = {
        readDigits: function () {
          for (var e = 0, t = !1, n = this.pos; ; ) {
            var i = this.buf.charCodeAt(this.pos);
            if (i != i) break;
            if (45 == i) {
              if (this.pos != n) break;
              t = !0;
            } else {
              if (i < 48 || 57 < i) break;
              e = 10 * e + (i - 48);
            }
            this.pos++;
          }
          return (t && (e *= -1), e);
        },
        readFloat: function () {
          for (var e = this.pos; ; ) {
            var t = this.buf.charCodeAt(this.pos);
            if (t != t) break;
            if (!((43 <= t && t < 58) || 101 == t || 69 == t)) break;
            this.pos++;
          }
          return parseFloat(U.substr(this.buf, e, this.pos - e));
        },
        unserializeObject: function (e) {
          for (;;) {
            if (this.pos >= this.length) throw new H("Invalid object");
            if (103 == this.buf.charCodeAt(this.pos)) break;
            var t = this.unserialize();
            if ("string" != typeof t) throw new H("Invalid object key");
            var n = this.unserialize();
            e[t] = n;
          }
          this.pos++;
        },
        unserializeEnum: function (e, t) {
          if (58 != this.buf.charCodeAt(this.pos++))
            throw new H("Invalid enum format");
          var n = this.readDigits();
          if (0 == n) return x.createEnum(e, t);
          for (var i = []; 0 < n--; ) i.push(this.unserialize());
          return x.createEnum(e, t, i);
        },
        unserialize: function () {
          switch (this.buf.charCodeAt(this.pos++)) {
            case 65:
              var e = this.unserialize(),
                t = this.resolver.resolveClass(e);
              if (null == t) throw new H("Class not found " + e);
              return t;
            case 66:
              var e = this.unserialize(),
                t = this.resolver.resolveEnum(e);
              if (null == t) throw new H("Enum not found " + e);
              return t;
            case 67:
              ((e = this.unserialize()), (t = this.resolver.resolveClass(e)));
              if (null == t) throw new H("Class not found " + e);
              e = x.createEmptyInstance(t);
              if (
                (this.cache.push(e),
                e.hxUnserialize(this),
                103 != this.buf.charCodeAt(this.pos++))
              )
                throw new H("Invalid custom data");
              return e;
            case 77:
              var n = new ns();
              for (
                this.cache.push(n), this.buf;
                104 != this.buf.charCodeAt(this.pos);
              ) {
                var i = this.unserialize();
                n.set(i, this.unserialize());
              }
              return (this.pos++, n);
            case 82:
              t = this.readDigits();
              if (t < 0 || t >= this.scache.length)
                throw new H("Invalid string reference");
              return this.scache[t];
            case 97:
              this.buf;
              var r = [];
              for (this.cache.push(r); ; ) {
                var o = this.buf.charCodeAt(this.pos);
                if (104 == o) {
                  this.pos++;
                  break;
                }
                117 == o
                  ? (this.pos++,
                    (o = this.readDigits()),
                    (r[r.length + o - 1] = null))
                  : r.push(this.unserialize());
              }
              return r;
            case 98:
              var a = new B();
              for (
                this.cache.push(a), this.buf;
                104 != this.buf.charCodeAt(this.pos);
              ) {
                var s = this.unserialize(),
                  l = this.unserialize();
                null != Y[s] ? a.setReserved(s, l) : (a.h[s] = l);
              }
              return (this.pos++, a);
            case 99:
              ((e = this.unserialize()), (t = this.resolver.resolveClass(e)));
              if (null == t) throw new H("Class not found " + e);
              e = x.createEmptyInstance(t);
              return (this.cache.push(e), this.unserializeObject(e), e);
            case 100:
              return this.readFloat();
            case 102:
              return !1;
            case 105:
              return this.readDigits();
            case 106:
              ((t = this.unserialize()), (e = this.resolver.resolveEnum(t)));
              if (null == e) throw new H("Enum not found " + t);
              this.pos++;
              var u = this.readDigits(),
                c = e.__constructs__.slice()[u];
              if (null == c) throw new H("Unknown enum index " + t + "@" + u);
              t = this.unserializeEnum(e, c);
              return (this.cache.push(t), t);
            case 107:
              return NaN;
            case 108:
              var _ = new En();
              for (
                this.cache.push(_), this.buf;
                104 != this.buf.charCodeAt(this.pos);
              )
                _.add(this.unserialize());
              return (this.pos++, _);
            case 109:
              return -1 / 0;
            case 110:
              return null;
            case 111:
              u = {};
              return (this.cache.push(u), this.unserializeObject(u), u);
            case 112:
              return 1 / 0;
            case 113:
              for (
                var h = new ts(),
                  d =
                    (this.cache.push(h),
                    this.buf,
                    this.buf.charCodeAt(this.pos++));
                58 == d;
              ) {
                var f = this.readDigits(),
                  m = this.unserialize();
                ((h.h[f] = m), (d = this.buf.charCodeAt(this.pos++)));
              }
              if (104 != d) throw new H("Invalid IntMap format");
              return h;
            case 114:
              e = this.readDigits();
              if (e < 0 || e >= this.cache.length)
                throw new H("Invalid reference");
              return this.cache[e];
            case 115:
              var c = this.readDigits(),
                p = this.buf;
              if (
                58 != this.buf.charCodeAt(this.pos++) ||
                this.length - this.pos < c
              )
                throw new H("Invalid bytes length");
              for (
                var g = Fn.CODES,
                  y =
                    (null == g && ((g = Fn.initCodes()), (Fn.CODES = g)),
                    this.pos),
                  t = 3 & c,
                  S = y + (c - t),
                  v = new rs(
                    new ArrayBuffer(3 * (c >> 2) + (2 <= t ? t - 1 : 0)),
                  ),
                  T = 0;
                y < S;
              ) {
                var I = g[p.charCodeAt(y++)],
                  E = g[p.charCodeAt(y++)],
                  I =
                    ((v.b[T++] = 255 & ((I << 2) | (E >> 4))),
                    g[p.charCodeAt(y++)]),
                  E =
                    ((v.b[T++] = 255 & ((E << 4) | (I >> 2))),
                    g[p.charCodeAt(y++)]);
                v.b[T++] = 255 & ((I << 6) | E);
              }
              return (
                2 <= t &&
                  ((u = g[p.charCodeAt(y++)]),
                  (e = g[p.charCodeAt(y++)]),
                  (v.b[T++] = 255 & ((u << 2) | (e >> 4))),
                  3 == t) &&
                  ((u = g[p.charCodeAt(y++)]),
                  (v.b[T++] = 255 & ((e << 4) | (u >> 2)))),
                (this.pos += c),
                this.cache.push(v),
                v
              );
            case 116:
              return !0;
            case 118:
              return (
                48 <= this.buf.charCodeAt(this.pos) &&
                this.buf.charCodeAt(this.pos) <= 57 &&
                48 <= this.buf.charCodeAt(this.pos + 1) &&
                this.buf.charCodeAt(this.pos + 1) <= 57 &&
                48 <= this.buf.charCodeAt(this.pos + 2) &&
                this.buf.charCodeAt(this.pos + 2) <= 57 &&
                48 <= this.buf.charCodeAt(this.pos + 3) &&
                this.buf.charCodeAt(this.pos + 3) <= 57 &&
                45 == this.buf.charCodeAt(this.pos + 4)
                  ? ((w = U.strDate(U.substr(this.buf, this.pos, 19))),
                    (this.pos += 19))
                  : ((t = this.readFloat()), (w = new Date(t))),
                this.cache.push(w),
                w
              );
            case 119:
              ((e = this.unserialize()), (u = this.resolver.resolveEnum(e)));
              if (null == u) throw new H("Enum not found " + e);
              c = this.unserializeEnum(u, this.unserialize());
              return (this.cache.push(c), c);
            case 120:
              throw H.wrap(this.unserialize());
            case 121:
              t = this.readDigits();
              if (
                58 != this.buf.charCodeAt(this.pos++) ||
                this.length - this.pos < t
              )
                throw new H("Invalid string length");
              var w = U.substr(this.buf, this.pos, t);
              return (
                (this.pos += t),
                (w = decodeURIComponent(w.split("+").join(" "))),
                this.scache.push(w),
                w
              );
            case 122:
              return 0;
          }
          throw (
            this.pos--,
            new H(
              "Invalid char " +
                this.buf.charAt(this.pos) +
                " at position " +
                this.pos,
            )
          );
        },
        __class__: Fn,
      }),
      ((t["albero.CustomUnserializer"] = C).__name__ = [
        "albero",
        "CustomUnserializer",
      ]),
      (C.run = function (e) {
        return new Fn(e).unserialize();
      }),
      (C.__super__ = Fn),
      (C.prototype = e(Fn.prototype, {
        unserialize: function () {
          if (99 != this.buf.charCodeAt(this.pos))
            return Fn.prototype.unserialize.call(this);
          this.pos++;
          var e = this.unserialize(),
            t = this.resolver.resolveClass(e);
          if (null == t) throw new H("Class not found " + e);
          e = x.createEmptyInstance(t);
          return (
            this.cache.push(e),
            this.unserializeObject(e),
            G.__instanceof(e, E) ? _.getOrStore(G.__cast(e, E)) : e
          );
        },
        __class__: C,
      })),
      ((t["albero.DateTimeHelper"] = W).__name__ = [
        "albero",
        "DateTimeHelper",
      ]),
      (W.isSameDate = function (e, t) {
        return (
          e.getFullYear() == t.getFullYear() &&
          e.getMonth() == t.getMonth() &&
          e.getDate() == t.getDate()
        );
      }),
      (W.dateString = function (e, t) {
        var n = new Date(),
          i = "";
        return (
          e.getFullYear() != n.getFullYear() && (i += e.getFullYear() + "/"),
          (i += e.getMonth() + 1 + "/" + e.getDate()),
          t && (i += " (" + W.dayOfWeekString(e) + ")"),
          i
        );
      }),
      (W.dayOfWeekString = function (e) {
        switch (e.getDay()) {
          case 0:
            return c.localize("DateTimeHelper.sun");
          case 1:
            return c.localize("DateTimeHelper.mon");
          case 2:
            return c.localize("DateTimeHelper.tue");
          case 3:
            return c.localize("DateTimeHelper.wed");
          case 4:
            return c.localize("DateTimeHelper.thu");
          case 5:
            return c.localize("DateTimeHelper.fri");
          default:
            return c.localize("DateTimeHelper.sat");
        }
      }),
      (W.dateStringForPicker = function (e) {
        return (
          e.getFullYear() +
          "-" +
          W.toDoubleDigits(e.getMonth() + 1) +
          "-" +
          W.toDoubleDigits(e.getDate())
        );
      }),
      (W.getDateStringForFile = function () {
        var e = new Date();
        return (
          e.getFullYear() +
          W.toDoubleDigits(e.getMonth() + 1) +
          W.toDoubleDigits(e.getDate()) +
          "_" +
          W.toDoubleDigits(e.getHours()) +
          W.toDoubleDigits(e.getMinutes()) +
          W.toDoubleDigits(e.getSeconds())
        );
      }),
      (W.getTimeStringForVoice = function (e) {
        return (
          W.toDoubleDigits((e / 60) | 0) + ":" + W.toDoubleDigits((0 | e) % 60)
        );
      }),
      (W.toDoubleDigits = function (e) {
        e = null == e ? "null" : "" + e;
        return (e = 1 == e.length ? "0" + e : e);
      }),
      (W.nowAsInt64 = function () {
        return W.dateAsInt64(new Date().getTime());
      }),
      (W.oneDayAfterAsInt64 = function () {
        return W.dateAsInt64(new Date().getTime() + 864e5);
      }),
      (W.afterAsInt64 = function (e) {
        return W.dateAsInt64(new Date().getTime() + e);
      }),
      (W.dateAsInt64 = function (e) {
        var t = 4294967296;
        return new E((e / t) | 0, (e % t) | 0);
      }),
      (W.abbreviateDatetimeString = function (e) {
        var t, n, i;
        return null == e
          ? ""
          : ((t = new Date()),
            (n = 4294967296),
            (i = e.high),
            (e = e.low),
            (i = new Date(i * n + (0 <= e ? e : e + n))),
            t.getFullYear() == i.getFullYear() &&
            t.getMonth() == i.getMonth() &&
            t.getDate() == i.getDate()
              ? i.getHours() +
                ":" +
                (i.getMinutes() < 10 ? "0" : "") +
                i.getMinutes()
              : W.dateString(i, !1));
      }),
      (W.getDateStringForNotification = function (e) {
        var t, n, i;
        return null == e
          ? ""
          : ((t = new Date()),
            (n = 4294967296),
            (i = e.high),
            (e = e.low),
            (i = new Date(i * n + (0 <= e ? e : e + n))),
            t.getFullYear() == i.getFullYear() &&
            t.getMonth() == i.getMonth() &&
            t.getDate() == i.getDate()
              ? i.getHours() +
                ":" +
                (i.getMinutes() < 10 ? "0" : "") +
                i.getMinutes()
              : i.getMonth() + 1 + "/" + i.getDate());
      }),
      (W.datetimeStringIn11Chars = function (e) {
        var t, n, i;
        return null == e
          ? ""
          : ((t = new Date()),
            (n = 4294967296),
            (i = e.high),
            (e = e.low),
            (i = new Date(i * n + (0 <= e ? e : e + n))),
            (e = W.dateString(i, !1)),
            t.getFullYear() == i.getFullYear() &&
              (e +=
                " " +
                i.getHours() +
                ":" +
                (i.getMinutes() < 10 ? "0" : "") +
                i.getMinutes()),
            e);
      }),
      (W.timeString = function (e) {
        var t, n;
        return null == e
          ? ""
          : ((t = 4294967296),
            (n = e.high),
            (e = e.low),
            (n = new Date(n * t + (0 <= e ? e : e + t))).getHours() +
              ":" +
              (n.getMinutes() < 10 ? "0" : "") +
              n.getMinutes());
      }),
      (W.datetimeString = function (e) {
        return null == e
          ? ""
          : ((e = parseFloat(Va.toString(e))),
            (e = new Date(e)).getFullYear() +
              "/" +
              (e.getMonth() + 1) +
              "/" +
              e.getDate() +
              " " +
              e.getHours() +
              ":" +
              (e.getMinutes() < 10 ? "0" : "") +
              e.getMinutes());
      }),
      ((t["albero.FileDynamicHelper"] = K).__name__ = [
        "albero",
        "FileDynamicHelper",
      ]),
      (K.filterDeleted = function (e) {
        return null == e ? [] : e.filter(K.isNotDeleted);
      }),
      (K.isNotDeleted = function (e) {
        return null == e.deleted || !e.deleted;
      }),
      (K.createFileInfoDynamic = function (e, t, n, i) {
        var r = {
          url: i.get_url,
          content_type: e.type,
          content_size: e.size,
          name: t,
          file_id: i.file_id,
        };
        return (
          h.foreach(n, function (e) {
            (null != e.auth && (r.thumbnail_url = e.auth.get_url),
              null != e.dimension &&
                (r.thumbnail_dimension = e.dimension.toObject()));
          }),
          r
        );
      }),
      ((t["albero.FileHelper"] = V).__name__ = ["albero", "FileHelper"]),
      (V.createByteSizeStringWithUnit = function (e) {
        if (null == e) return "0 B";
        for (var t = ["", "K", "M", "G", "T"]; ; ) {
          var n,
            i = (n = Va.divMod(e, _.make(0, 1024)).quotient),
            r = new E(0, 0),
            o = (i.high - r.high) | 0;
          if (
            (0 == o && (o = w.ucompare(i.low, r.low)),
            !(0 < (i.high < 0 ? (r.high < 0 ? o : -1) : 0 <= r.high ? o : 1)))
          )
            break;
          ((e = n), t.shift());
        }
        return Va.toString(e) + " " + t.shift() + "B";
      }),
      (V.extractFilenameExceptExtension = function (e) {
        return new v("\\.(?=[^.]+$)", "").split(e)[0];
      }),
      (V.getUniqueFilename = function (e, t) {
        for (
          var n = e, i = 0;
          null != Y[n] ? t.existsReserved(n) : t.h.hasOwnProperty(n);
        ) {
          ++i;
          var r = V.extractFilenameExceptExtension(e),
            n = r + "(" + i + ")" + U.substr(e, r.length, null);
        }
        return n;
      }),
      ((t["albero.History"] = q).__name__ = ["albero", "History"]),
      (q.replaceState = function (e, t, n) {
        return (
          null != window.history && (window.history.replaceState(e, t, n), !0)
        );
      }),
      function () {}),
    xn =
      (((t["albero.Int64Helper"] = D).__name__ = ["albero", "Int64Helper"]),
      (D.parse = function (e) {
        if (!new v("^\\d+$", "").match(e)) return null;
        for (
          var t = _.make(0, 0), n = _.make(0, 10), i = 0, r = e.length;
          i < r;
        ) {
          var o = i++,
            a = 65535 & t.low,
            s = t.low >>> 16,
            l = 65535 & n.low,
            u = n.low >>> 16,
            c = w._mul(a, l),
            l = w._mul(s, l),
            a = w._mul(a, u),
            s = (((w._mul(s, u) + (a >>> 16)) | 0) + (l >>> 16)) | 0,
            s =
              ((a <<= 16),
              w.ucompare((u = (c + a) | 0), a) < 0 && (s++, (s |= 0)),
              (l <<= 16),
              w.ucompare((u = (u + l) | 0), l) < 0 && (s++, (s |= 0)),
              (s + ((w._mul(t.low, n.high) + w._mul(t.high, n.low)) | 0)) | 0),
            c = new E(s, u),
            a = _.make(0, P.parseInt(e.charAt(o))),
            l = (c.high + a.high) | 0,
            s = (c.low + a.low) | 0;
          w.ucompare(s, c.low) < 0 && (l++, (l |= 0));
          t = new E(l, s);
        }
        return t;
      }),
      (D.getHigh = function (e) {
        return e.high;
      }),
      (D.getLow = function (e) {
        return e.low;
      }),
      (D.idStr = function (e) {
        return "_" + e.high + "_" + e.low;
      }),
      (D.makeFromIdStr = function (e) {
        var t = new v("^_(-?\\d*)_(-?\\d*)$", "");
        return t.match(e)
          ? _.make(P.parseInt(t.matched(1)), P.parseInt(t.matched(2)))
          : null;
      }),
      (D.fromNullableIdStr = function (e) {
        return null != e ? D.makeFromIdStr(e) : null;
      }),
      (D.toStr = function (e) {
        return Va.toString(e);
      }),
      (D.isNeg = function (e) {
        return e.high < 0;
      }),
      (D.add = function (e, t) {
        var n = (e.high + t.high) | 0,
          t = (e.low + t.low) | 0;
        return (w.ucompare(t, e.low) < 0 && (n++, (n |= 0)), new E(n, t));
      }),
      (D.sub = function (e, t) {
        var n = (e.high - t.high) | 0,
          i = (e.low - t.low) | 0;
        return (w.ucompare(e.low, t.low) < 0 && (n--, (n |= 0)), new E(n, i));
      }),
      (D.compare = function (e, t) {
        var n = (e.high - t.high) | 0;
        return (
          0 == n && (n = w.ucompare(e.low, t.low)),
          e.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1
        );
      }),
      (D.eq = function (e, t) {
        return null != e && null != t && e.high == t.high && e.low == t.low;
      }),
      (D.eqOrBothNull = function (e, t) {
        return (
          (null == e && null == t) ||
          (null != e && null != t && e.high == t.high && e.low == t.low)
        );
      }),
      (D.toFloat = function (e) {
        var t = 4294967296,
          n = e.high,
          e = e.low;
        return n * t + (0 <= e ? e : e + t);
      }),
      (D.idStrToInt64 = function (e) {
        e = e.split("_");
        return 2 < e.length ? _.make(P.parseInt(e[1]), P.parseInt(e[2])) : null;
      }),
      (D.decrement = function (e) {
        var t, n, i;
        return null == e
          ? null
          : ((t = new E(0, 1)),
            (n = (e.high - t.high) | 0),
            (i = (e.low - t.low) | 0),
            w.ucompare(e.low, t.low) < 0 && (n--, (n |= 0)),
            new E(n, i));
      }),
      (D.increment = function (e) {
        var t, n;
        return null == e
          ? null
          : ((n = new E(0, 1)),
            (t = (e.high + n.high) | 0),
            (n = (e.low + n.low) | 0),
            w.ucompare(n, e.low) < 0 && (t++, (t |= 0)),
            new E(t, n));
      }),
      (D.unique = function (e) {
        if (null == e) return [];
        for (var t = [], n = 0, i = e.length; n < i; ) {
          var r = n++,
            o = e[r];
          D.indexOf(e, o) == r && t.push(o);
        }
        return t;
      }),
      (D.indexOf = function (e, t) {
        if (null != e)
          for (var n = 0, i = e.length; n < i; ) {
            var r = n++,
              o = e[r];
            if (null != o && null != t && o.high == t.high && o.low == t.low)
              return r;
          }
        return -1;
      }),
      (D.contains = function (e, t) {
        if (null != e)
          for (var n = 0; n < e.length; ) {
            var i = e[n];
            if (
              (++n,
              null != i && null != t && i.high == t.high && i.low == t.low)
            )
              return !0;
          }
        return !1;
      }),
      (D.notContains = function (e, t) {
        return !D.contains(e, t);
      }),
      (D.remove = function (e, t) {
        if (null != e)
          for (var n = 0; n < e.length; ) {
            var i = e[n];
            if (
              (++n,
              null != i && null != t && i.high == t.high && i.low == t.low)
            )
              return (U.remove(e, i), !0);
          }
        return !1;
      }),
      (D.fromIntOrInt64 = function (e) {
        return null == e || G.__instanceof(e, E)
          ? e
          : ((e = G.__cast(e, js)), new E(e >> 31, e));
      }),
      (D.eqArray = function (e, t) {
        if (e != t) {
          if (null == e || null == t) return !1;
          if (e.length != t.length) return !1;
          for (var n = 0, i = e.length; n < i; ) {
            var r = n++,
              o = e[r],
              r = t[r];
            if (null == o || null == r || o.high != r.high || o.low != r.low)
              return !1;
          }
        }
        return !0;
      }),
      (D.splitUsingMaxCount = function (e, t) {
        for (var n = [], i = 0; i < e.length; ) {
          var r = i + t,
            o = e.slice(i, r);
          (n.push(o), (i = r));
        }
        return n;
      }),
      (t["albero.AccessMode"] = {
        __ename__: ["albero", "AccessMode"],
        __constructs__: ["Disabled", "ReadOnly", "ReadWrite"],
      })),
    _ =
      ((xn.Disabled = ["Disabled", 0]),
      (xn.Disabled.toString = a),
      ((xn.Disabled.__enum__ = xn).ReadOnly = ["ReadOnly", 1]),
      (xn.ReadOnly.toString = a),
      ((xn.ReadOnly.__enum__ = xn).ReadWrite = ["ReadWrite", 2]),
      (xn.ReadWrite.toString = a),
      ((xn.ReadWrite.__enum__ = xn).__empty_constructs__ = [
        xn.Disabled,
        xn.ReadOnly,
        xn.ReadWrite,
      ]),
      function () {}),
    Un =
      (((t["albero.Int64Pool"] = _).__name__ = ["albero", "Int64Pool"]),
      (_.make = function (e, t) {
        if (_.accessMode == xn.Disabled) return new E(e, t);
        null == _.instancePool && (_.instancePool = new B());
        var n = "_" + e + "_" + t,
          i = _.instancePool,
          i = null != Y[n] ? i.getReserved(n) : i.h[n];
        return (
          null != i ||
            ((i = new E(e, t)),
            _.accessMode == xn.ReadWrite &&
              ((e = _.instancePool),
              null != Y[n] ? e.setReserved(n, i) : (e.h[n] = i))),
          i
        );
      }),
      (_.getOrStore = function (e) {
        if (_.accessMode != xn.Disabled) {
          null == _.instancePool && (_.instancePool = new B());
          var t = "_" + e.high + "_" + e.low,
            n = _.instancePool,
            n = null != Y[t] ? n.getReserved(t) : n.h[t];
          if (null != n) return n;
          _.accessMode == xn.ReadWrite &&
            ((n = _.instancePool),
            null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e));
        }
        return e;
      }),
      (_.enable = function () {
        _.updateAccessMode(xn.ReadWrite);
      }),
      (_.disable = function () {
        _.updateAccessMode(xn.Disabled);
      }),
      (_.updateAccessMode = function (e) {
        _.accessMode = e;
      }),
      (_.clear = function () {
        _.instancePool = null;
      }),
      ((t["albero.LambdaHelper"] = Q).__name__ = ["albero", "LambdaHelper"]),
      (Q.flatten = function (e) {
        return k
          .filter(e, function (e) {
            return h.isDefined(e);
          })
          .map(function (e) {
            return h.get(e);
          });
      }),
      (Q.headOption = function (e) {
        for (var t = Ns(e)(); t.hasNext(); ) {
          var n = t.next();
          return m.Some(n);
        }
        return m.None;
      }),
      (Q.lastOption = function (e) {
        for (var t = m.None, n = Ns(e)(); n.hasNext(); )
          var i = n.next(), t = m.Some(i);
        return t;
      }),
      (Q.findOption = function (e, t) {
        return h.option(k.find(e, t));
      }),
      (Q.indexWhere = function (e, t) {
        for (var n = 0, i = Ns(e)(); i.hasNext(); ) {
          if (t(i.next())) return n;
          ++n;
        }
        return -1;
      }),
      (Q.forall = function (e, t) {
        return !k.exists(e, function (e) {
          return !t(e);
        });
      }),
      function () {}),
    h =
      (((t["albero.ArrayHelper"] = Un).__name__ = ["albero", "ArrayHelper"]),
      (Un.lastOption = function (e) {
        return 0 == e.length ? m.None : m.Some(e[e.length - 1]);
      }),
      (Un.sortAndReturn = function (e, t) {
        return (e.sort(t), e);
      }),
      (Un.reverseAndReturn = function (e) {
        return (e.reverse(), e);
      }),
      (Un.asArray = function (e) {
        return G.__cast(e, Array);
      }),
      ((t["albero.MapHelper"] = X).__name__ = ["albero", "MapHelper"]),
      (X.copyStringMap = function (e) {
        for (var t = new B(), n = e.keys(); n.hasNext(); ) {
          var i = n.next(),
            r = null != Y[i] ? e.getReserved(i) : e.h[i];
          null != Y[i] ? t.setReserved(i, r) : (t.h[i] = r);
        }
        return t;
      }),
      (X.toKeyArray = function (e) {
        for (var t = [], n = e.keys(); n.hasNext(); ) t.push(n.next());
        return t;
      }),
      (X.toValueArray = function (e) {
        for (var t = [], n = e.iterator(); n.hasNext(); ) t.push(n.next());
        return t;
      }),
      ((t["albero.NoteContentCompressor"] = J).__name__ = [
        "albero",
        "NoteContentCompressor",
      ]),
      (J.prototype = {
        run: function (n) {
          var e;
          return this.shouldCompress(n)
            ? ((e = n.rich_text),
              this.compress(e).then(
                function (e) {
                  var t = L.copy(n);
                  return (
                    L.deleteField(t, "rich_text"),
                    (t.rich_text_compressed = e),
                    t
                  );
                },
                function (e) {
                  return Promise.reject({ error: e });
                },
              ))
            : Promise.resolve(n);
        },
        shouldCompress: function (e) {
          return null != e && null != e.rich_text;
        },
        compress: function (e) {
          e = new TextEncoder().encode(e);
          return new Response(
            new Blob([e]).stream().pipeThrough(new CompressionStream("gzip")),
          )
            .arrayBuffer()
            .then(function (e) {
              return rs.ofData(e);
            });
        },
        __class__: J,
      }),
      ((t["albero.NoteContentUncompressor"] = Z).__name__ = [
        "albero",
        "NoteContentUncompressor",
      ]),
      (Z.prototype = {
        run: function (n) {
          var e;
          return null == n
            ? Promise.reject({ error: null })
            : this.shouldUncompress(n)
              ? ((e = n.rich_text_compressed),
                this.uncompress(e).then(
                  function (e) {
                    var t = L.copy(n);
                    return (
                      L.deleteField(t, "rich_text_compressed"),
                      (t.rich_text = e),
                      t
                    );
                  },
                  function (e) {
                    return Promise.reject({ error: e });
                  },
                ))
              : Promise.resolve(n);
        },
        shouldUncompress: function (e) {
          return null != e && null != e.rich_text_compressed;
        },
        uncompress: function (e) {
          e = e.b.bufferValue;
          return new Response(
            new Blob([e]).stream().pipeThrough(new DecompressionStream("gzip")),
          )
            .arrayBuffer()
            .then(function (e) {
              return new TextDecoder().decode(e);
            });
        },
        __class__: Z,
      }),
      function () {}),
    o =
      (((t["albero.OptionHelper"] = h).__name__ = ["albero", "OptionHelper"]),
      (h.map = function (e, t) {
        switch (e[1]) {
          case 0:
            return m.Some(t(e[2]));
          case 1:
            return m.None;
        }
      }),
      (h.flatMap = function (e, t) {
        switch (e[1]) {
          case 0:
            return t(e[2]);
          case 1:
            return m.None;
        }
      }),
      (h.foreach = function (e, t) {
        0 === e[1] && t(e[2]);
      }),
      (h.filter = function (e, t) {
        switch (e[1]) {
          case 0:
            return t(e[2]) ? e : m.None;
          case 1:
            return m.None;
        }
      }),
      (h.exists = function (e, t) {
        return h.isDefined(h.filter(e, t));
      }),
      (h.isDefined = function (e) {
        switch (e[1]) {
          case 0:
            return !0;
          case 1:
            return !1;
        }
      }),
      (h.isEmpty = function (e) {
        switch (e[1]) {
          case 0:
            return !1;
          case 1:
            return !0;
        }
      }),
      (h.get = function (e) {
        switch (e[1]) {
          case 0:
            return e[2];
          case 1:
            throw new H("None.get called");
        }
      }),
      (h.orNull = function (e) {
        switch (e[1]) {
          case 0:
            return e[2];
          case 1:
            return null;
        }
      }),
      (h.flatten = function (e) {
        return h.flatMap(e, function (e) {
          return e;
        });
      }),
      (h.eq = function (e, t) {
        return x.enumEq(e, t);
      }),
      (h.option = function (e) {
        return null != e ? m.Some(e) : m.None;
      }),
      function () {}),
    Ln =
      (((t["albero.StorageHelper"] = o).__name__ = ["albero", "StorageHelper"]),
      (o.getItemAsJson = function (e) {
        e = o.getItem(e);
        if (null == e) return null;
        try {
          return JSON.parse(e);
        } catch (e) {
          return null;
        }
      }),
      (o.getItemAsInt = function (e) {
        e = o.getItem(e);
        return null == e ? null : P.parseInt(e);
      }),
      (o.getItem = function (e) {
        var t = Ro.getLocalStorage();
        if (null == t) F.e("storage not found");
        else
          try {
            return t.getItem(e);
          } catch (e) {
            (e instanceof H && (e = e.val), F.e("storage error occered.", e));
          }
        return null;
      }),
      (o.removeItem = function (e) {
        var t = Ro.getLocalStorage();
        if (null == t) F.e("storage not found");
        else
          try {
            return (t.removeItem(e), !0);
          } catch (e) {
            (e instanceof H && (e = e.val), F.e("storage error occered.", e));
          }
        return !1;
      }),
      (o.setItem = function (e, t) {
        var n = Ro.getLocalStorage();
        if (null == n) F.e("storage not found");
        else
          try {
            return (n.setItem(e, t), !0);
          } catch (e) {
            (e instanceof H && (e = e.val), F.e("storage error occered.", e));
          }
        return !1;
      }),
      (o.getItemWithUnserialize = function (e) {
        e = o.getItem(e);
        if (null != e)
          try {
            return C.run(e);
          } catch (e) {
            (e instanceof H && (e = e.val), F.w(e));
          }
        return null;
      }),
      (o.setItemWithSerialize = function (e, t) {
        o.setItem(e, Qa.run(t));
      }),
      function () {}),
    g =
      (((t["albero.StringHelper"] = Ln).__name__ = ["albero", "StringHelper"]),
      (Ln.trim = function (e) {
        return e.trim();
      }),
      (Ln.compare = function (e, t) {
        return t < e ? 1 : e < t ? -1 : 0;
      }),
      (Ln.isEmpty = function (e) {
        return null == e || 0 == e.length;
      }),
      (Ln.isNotEmpty = function (e) {
        return !Ln.isEmpty(e);
      }),
      (Ln.randomString = function () {
        var e = "",
          t = 8;
        for (e += P.string(I.hex(0 | new Date().getTime(), 8)); t++ < 36; )
          e += P.string(
            0 != ((51 * t) & 52)
              ? I.hex(
                  0 != (15 ^ t)
                    ? 8 ^ ((Math.random() * (0 != (20 ^ t) ? 16 : 4)) | 0)
                    : 4,
                )
              : "-",
          );
        return e.toLowerCase();
      }),
      (t["albero.Urls"] = {
        __ename__: ["albero", "Urls"],
        __constructs__: [
          "auto",
          "domains",
          "domain",
          "members",
          "talks",
          "actions",
          "bookmarks",
          "scheduled_messages",
          "settings",
          "announcements",
          "error",
          "loading",
          "environment",
          "apps",
          "unknown",
        ],
      })),
    Pn =
      ((g.auto = ["auto", 0]),
      (g.auto.toString = a),
      ((g.auto.__enum__ = g).domains = ["domains", 1]),
      (g.domains.toString = a),
      ((g.domains.__enum__ = g).domain = function (e) {
        e = ["domain", 2, e];
        return ((e.__enum__ = g), (e.toString = a), e);
      }),
      (g.members = function (e) {
        e = ["members", 3, e];
        return ((e.__enum__ = g), (e.toString = a), e);
      }),
      (g.talks = function (e, t) {
        e = ["talks", 4, e, t];
        return ((e.__enum__ = g), (e.toString = a), e);
      }),
      (g.actions = function (e) {
        e = ["actions", 5, e];
        return ((e.__enum__ = g), (e.toString = a), e);
      }),
      (g.bookmarks = function (e) {
        e = ["bookmarks", 6, e];
        return ((e.__enum__ = g), (e.toString = a), e);
      }),
      (g.scheduled_messages = function (e) {
        e = ["scheduled_messages", 7, e];
        return ((e.__enum__ = g), (e.toString = a), e);
      }),
      (g.settings = function (e) {
        e = ["settings", 8, e];
        return ((e.__enum__ = g), (e.toString = a), e);
      }),
      (g.announcements = ["announcements", 9]),
      (g.announcements.toString = a),
      ((g.announcements.__enum__ = g).error = ["error", 10]),
      (g.error.toString = a),
      ((g.error.__enum__ = g).loading = ["loading", 11]),
      (g.loading.toString = a),
      ((g.loading.__enum__ = g).environment = function (e) {
        e = ["environment", 12, e];
        return ((e.__enum__ = g), (e.toString = a), e);
      }),
      (g.apps = function (e) {
        e = ["apps", 13, e];
        return ((e.__enum__ = g), (e.toString = a), e);
      }),
      (g.unknown = ["unknown", 14]),
      (g.unknown.toString = a),
      ((g.unknown.__enum__ = g).__empty_constructs__ = [
        g.auto,
        g.domains,
        g.announcements,
        g.error,
        g.loading,
        g.unknown,
      ]),
      ((t["albero.UrlsHelper"] = $).__name__ = ["albero", "UrlsHelper"]),
      ($.isDisabledForDirectAppsOnly = function (e) {
        if (null == e) return !1;
        switch (e[1]) {
          case 2:
          case 3:
          case 4:
          case 5:
          case 6:
          case 7:
            return !0;
          default:
            return !1;
        }
      }),
      ($.isDirectAppsOnly = function (e) {
        return null != e && 13 == e[1];
      }),
      (t["albero.TalkUrlParam"] = {
        __ename__: ["albero", "TalkUrlParam"],
        __constructs__: ["multi", "single", "auto"],
      })),
    Bn =
      ((Pn.multi = function (e) {
        e = ["multi", 0, e];
        return ((e.__enum__ = Pn), (e.toString = a), e);
      }),
      (Pn.single = function (e) {
        e = ["single", 1, e];
        return ((e.__enum__ = Pn), (e.toString = a), e);
      }),
      (Pn.auto = ["auto", 2]),
      (Pn.auto.toString = a),
      ((Pn.auto.__enum__ = Pn).__empty_constructs__ = [Pn.auto]),
      ((t["albero.TalkUrlParamHelper"] = ee).__name__ = [
        "albero",
        "TalkUrlParamHelper",
      ]),
      (ee.getTalkSelectionForSingle = function (e) {
        switch (e[1]) {
          case 0:
            return m.None;
          case 1:
            return m.Some(e[2]);
          case 2:
            return m.None;
        }
      }),
      (ee.getTalkSelectionsForMulti = function (e) {
        switch (e[1]) {
          case 0:
            return m.Some(e[2]);
          case 1:
          case 2:
            return m.None;
        }
      }),
      ((t["albero.UrlConverter"] = te).__name__ = ["albero", "UrlConverter"]),
      (te.createUrls = function (e, t) {
        return null == e || "" != (e = e.split("/"))[0] || e.length < 2
          ? g.unknown
          : 2 == e.length
            ? te.createUrlsWithoutDomain(e, t)
            : te.createUrlsWithDomain(e, t);
      }),
      (te.createUrlsWithoutDomain = function (e, t) {
        switch (e[1]) {
          case "":
            return g.domains;
          case "apps":
            var n = t.getFallbackDomainId();
            return (
              h.foreach(n, function (e) {
                q.replaceState(
                  "apps",
                  null,
                  "#" + te.createFragment(g.apps(e), t),
                );
              }),
              g.apps(h.orNull(n))
            );
          case "environment":
            n = t.getFallbackDomainId();
            return (
              h.foreach(n, function (e) {
                q.replaceState(
                  "environment",
                  null,
                  "#" + te.createFragment(g.environment(e), t),
                );
              }),
              g.environment(h.orNull(n))
            );
          case "error":
            return g.error;
          case "loading":
            return g.loading;
          case "settings":
            n = t.getFallbackDomainId();
            return (
              h.foreach(n, function (e) {
                q.replaceState(
                  "settings",
                  null,
                  "#" + te.createFragment(g.settings(e), t),
                );
              }),
              g.settings(h.orNull(n))
            );
          default:
            return g.unknown;
        }
      }),
      (te.createUrlsWithDomain = function (e, t) {
        var n = D.parse(e[1]);
        if (null != n)
          switch (e[2]) {
            case "":
              return g.domain(n);
            case "actions":
              return g.actions(n);
            case "apps":
              return g.apps(n);
            case "bookmarks":
              return g.bookmarks(n);
            case "environment":
              return g.environment(n);
            case "members":
              return g.members(n);
            case "scheduled_messages":
              return g.scheduled_messages(n);
            case "settings":
              return g.settings(n);
            case "talks":
              if (e.length < 4) return g.talks(n, Pn.single($r.NotSelected));
              if ("auto" == e[3]) return g.talks(n, Pn.auto);
              if ("multi" == e[3] || ",," == e[3])
                return g.talks(
                  n,
                  Pn.multi([$r.NotSelected, $r.NotSelected, $r.NotSelected]),
                );
              if ("single" == e[3] || "" == e[3])
                return g.talks(n, Pn.single($r.NotSelected));
              var i,
                r = e[3].split(",");
              if (1 == r.length)
                return (
                  (i = e.length < 5 ? m.None : h.option(D.parse(e[4]))),
                  (i = te.createTalkSelection(r[0], i)),
                  g.talks(n, Pn.single(i))
                );
              if (3 == r.length)
                return (
                  (i = r.map(function (e) {
                    return te.createTalkSelection(e, m.None);
                  })),
                  g.talks(n, Pn.multi(i))
                );
          }
        return g.unknown;
      }),
      (te.createTalkSelection = function (e, t) {
        return e == te.ANNOUNCEMENTS_KEY
          ? $r.AnnouncementsSelected
          : null == (e = D.parse(e))
            ? $r.NotSelected
            : $r.TalkSelected(e, t);
      }),
      (te.createFragment = function (e, t) {
        switch (e[1]) {
          case 1:
            return "/";
          case 2:
            return "/" + Va.toString(e[2]) + "/";
          case 3:
            return "/" + Va.toString(e[2]) + "/members";
          case 4:
            var n = e[3],
              i = e[2];
            switch (n[1]) {
              case 0:
                var r = "/" + n[2].map(te.createTalkPartFragment).join(",");
                return "/" + Va.toString(i) + "/talks" + r;
              case 1:
                ((r = te.createTalkPartFragment(n[2])),
                  (r = Ln.isEmpty(r) ? "" : "/" + r));
                return "/" + Va.toString(i) + "/talks" + r;
              case 2:
                return "/" + Va.toString(i) + "/talks/auto";
            }
            break;
          case 5:
            return "/" + Va.toString(e[2]) + "/actions";
          case 6:
            return "/" + Va.toString(e[2]) + "/bookmarks";
          case 7:
            return "/" + Va.toString(e[2]) + "/scheduled_messages";
          case 8:
            var o,
              a = e[2];
            switch (
              (F.v(
                "convert urls.settings to fragment. domainId:",
                null != a ? a : "",
              ),
              h.option(a)[1])
            ) {
              case 0:
                o = h.option(a);
                break;
              case 1:
                o = t.getLastSelectedDomainId();
            }
            var s = h.map(o, function (e) {
              return "/" + Va.toString(e) + "/settings";
            });
            switch (s[1]) {
              case 0:
                return s[2];
              case 1:
                return "/settings";
            }
            break;
          case 10:
          case 11:
            return "/" + P.string(e);
          case 12:
            var l,
              u = e[2];
            switch (
              (F.v(
                "convert urls.environment to fragment. domainId:",
                null != u ? u : "",
              ),
              h.option(u)[1])
            ) {
              case 0:
                l = h.option(u);
                break;
              case 1:
                l = t.getLastSelectedDomainId();
            }
            var c = h.map(l, function (e) {
              return "/" + Va.toString(e) + "/environment";
            });
            switch (c[1]) {
              case 0:
                return c[2];
              case 1:
                return "/environment";
            }
            break;
          case 13:
            return "/" + Va.toString(e[2]) + "/apps";
          default:
            return null;
        }
      }),
      (te.createTalkPartFragment = function (e) {
        switch (e[1]) {
          case 0:
            return "";
          case 1:
            var t = e[2],
              n = h.map(e[3], function (e) {
                return "/" + Va.toString(e);
              });
            switch (n[1]) {
              case 0:
                var i = n[2];
                break;
              case 1:
                i = "";
            }
            return Va.toString(t) + i;
          case 2:
            return te.ANNOUNCEMENTS_KEY;
        }
      }),
      ((t["albero.IUrlConverterDelegate"] = ne).__name__ = [
        "albero",
        "IUrlConverterDelegate",
      ]),
      (ne.prototype = { __class__: ne }),
      ((t["puremvc.interfaces.INotifier"] = ie).__name__ = [
        "puremvc",
        "interfaces",
        "INotifier",
      ]),
      (ie.prototype = { __class__: ie }),
      function () {
        this.facade = kn.getInstance();
      }),
    Hn =
      (((t["puremvc.patterns.observer.Notifier"] = Bn).__name__ = [
        "puremvc",
        "patterns",
        "observer",
        "Notifier",
      ]),
      (Bn.__interfaces__ = [ie]),
      (Bn.prototype = {
        sendNotification: function (e, t, n) {
          this.facade.sendNotification(e, t, n);
        },
        __class__: Bn,
      }),
      ((t["puremvc.interfaces.ICommand"] = re).__name__ = [
        "puremvc",
        "interfaces",
        "ICommand",
      ]),
      (re.prototype = { __class__: re }),
      function () {
        Bn.call(this);
      }),
    r =
      (((t["puremvc.patterns.command.SimpleCommand"] = Hn).__name__ = [
        "puremvc",
        "patterns",
        "command",
        "SimpleCommand",
      ]),
      (Hn.__interfaces__ = [re]),
      (Hn.__super__ = Bn),
      (Hn.prototype = e(Bn.prototype, {
        execute: function (e) {},
        __class__: Hn,
      })),
      function () {
        (Hn.call(this), i.getInstance().autoBind(this));
      }),
    Gn =
      (((t["albero.command.AutoBindCommand"] = r).__name__ = [
        "albero",
        "command",
        "AutoBindCommand",
      ]),
      (r.__super__ = Hn),
      (r.prototype = e(Hn.prototype, { __class__: r })),
      function () {
        r.call(this);
      }),
    jn =
      (((t["albero.command.AccountControlRequestCommand"] = Gn).__name__ = [
        "albero",
        "command",
        "AccountControlRequestCommand",
      ]),
      (Gn.__super__ = r),
      (Gn.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t[1]) {
            case 0:
              this.api.acceptAccountControlRequest(t[2]);
              break;
            case 1:
              this.api.rejectAccountControlRequest(t[2]);
          }
        },
        __class__: Gn,
      })),
      (t["albero.command.AccountControlRequestAction"] = {
        __ename__: ["albero", "command", "AccountControlRequestAction"],
        __constructs__: ["ACCEPT_REQUEST", "REJECT_REQUEST"],
      })),
    zn =
      ((jn.ACCEPT_REQUEST = function (e) {
        e = ["ACCEPT_REQUEST", 0, e];
        return ((e.__enum__ = jn), (e.toString = a), e);
      }),
      (jn.REJECT_REQUEST = function (e) {
        e = ["REJECT_REQUEST", 1, e];
        return ((e.__enum__ = jn), (e.toString = a), e);
      }),
      (jn.__empty_constructs__ = []),
      function () {
        r.call(this);
      }),
    Yn =
      (((t["albero.command.ConferenceCommand"] = zn).__name__ = [
        "albero",
        "command",
        "ConferenceCommand",
      ]),
      (zn.__super__ = r),
      (zn.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t[1]) {
            case 0:
              this.api.joinConference(t[2], t[3]);
              break;
            case 1:
              var n = t[2],
                n =
                  (this.api.leaveConference(n),
                  this.dataStore.getConference(n));
              null != n &&
                n.isExpired() &&
                (this.dataStore.removeConference(n),
                this.sendNotification("notify_close_conference", n));
              break;
            case 2:
              this.api.rejectConference(t[2]);
              break;
            case 3:
              this.api.getConferenceParticipants(t[2]);
          }
        },
        __class__: zn,
      })),
      (t["albero.command.ConferenceAction"] = {
        __ename__: ["albero", "command", "ConferenceAction"],
        __constructs__: [
          "JOIN_CONFERENCE",
          "LEAVE_CONFERENCE",
          "REJECT_CONFERENCE",
          "GET_CONFERENCE_PARTICIPANTS",
        ],
      })),
    Wn =
      ((Yn.JOIN_CONFERENCE = function (e, t) {
        e = ["JOIN_CONFERENCE", 0, e, t];
        return ((e.__enum__ = Yn), (e.toString = a), e);
      }),
      (Yn.LEAVE_CONFERENCE = function (e) {
        e = ["LEAVE_CONFERENCE", 1, e];
        return ((e.__enum__ = Yn), (e.toString = a), e);
      }),
      (Yn.REJECT_CONFERENCE = function (e) {
        e = ["REJECT_CONFERENCE", 2, e];
        return ((e.__enum__ = Yn), (e.toString = a), e);
      }),
      (Yn.GET_CONFERENCE_PARTICIPANTS = function (e) {
        e = ["GET_CONFERENCE_PARTICIPANTS", 3, e];
        return ((e.__enum__ = Yn), (e.toString = a), e);
      }),
      (Yn.__empty_constructs__ = []),
      function () {
        r.call(this);
      }),
    Kn =
      (((t["albero.command.DeviceCommand"] = Wn).__name__ = [
        "albero",
        "command",
        "DeviceCommand",
      ]),
      (Wn.__super__ = r),
      (Wn.prototype = e(r.prototype, {
        execute: function (e) {
          var n = this;
          e.getBody();
          require("read")({ prompt: "Code: " }, function (e, t) {
            n.api.authorizeDevice(t, function () {
              n.sendNotification("SignIn");
            });
          });
        },
        __class__: Wn,
      })),
      (t["albero.command.DeviceAction"] = {
        __ename__: ["albero", "command", "DeviceAction"],
        __constructs__: ["AUTHORIZE"],
      })),
    Vn =
      ((Kn.AUTHORIZE = ["AUTHORIZE", 0]),
      (Kn.AUTHORIZE.toString = a),
      ((Kn.AUTHORIZE.__enum__ = Kn).__empty_constructs__ = [Kn.AUTHORIZE]),
      function () {
        r.call(this);
      }),
    qn =
      (((t["albero.command.DomainCommand"] = Vn).__name__ = [
        "albero",
        "command",
        "DomainCommand",
      ]),
      (Vn.__super__ = r),
      (Vn.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t[1]) {
            case 0:
              this.api.acceptDomainInvite(t[2]);
              break;
            case 1:
              this.api.deleteDomainInvite(t[2]);
              break;
            case 2:
              this.api.leaveDomain(t[2]);
          }
        },
        __class__: Vn,
      })),
      (t["albero.command.DomainAction"] = {
        __ename__: ["albero", "command", "DomainAction"],
        __constructs__: ["ACCEPT_INVITE", "DELETE_INVITE", "LEAVE"],
      })),
    Qn =
      ((qn.ACCEPT_INVITE = function (e) {
        e = ["ACCEPT_INVITE", 0, e];
        return ((e.__enum__ = qn), (e.toString = a), e);
      }),
      (qn.DELETE_INVITE = function (e) {
        e = ["DELETE_INVITE", 1, e];
        return ((e.__enum__ = qn), (e.toString = a), e);
      }),
      (qn.LEAVE = function (e) {
        e = ["LEAVE", 2, e];
        return ((e.__enum__ = qn), (e.toString = a), e);
      }),
      (qn.__empty_constructs__ = []),
      function () {
        r.call(this);
      }),
    Xn =
      (((t["albero.command.FilePathCommand"] = Qn).__name__ = [
        "albero",
        "command",
        "FilePathCommand",
      ]),
      (Qn.__super__ = r),
      (Qn.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t[1]) {
            case 0:
              var n = t[6],
                i = t[2],
                r = this.fileService.createDummyFile(t[3], t[4], t[5]);
              this.api.uploadForHubot(i, r, n);
              break;
            case 1:
              for (
                var i = t[7],
                  o = t[6],
                  a = t[5],
                  s = t[4],
                  r = t[3],
                  n = t[2],
                  l = [],
                  u = 0,
                  c = s.length;
                u < c;
              ) {
                var _ = u++;
                l.push(this.fileService.createDummyFile(s[_], a[_], o[_]));
              }
              this.api.uploadMultiForHubot(n, r, l, i);
              break;
            case 2:
              this.fileService.download(t[2], t[3], t[4]);
          }
        },
        __class__: Qn,
      })),
      (t["albero.command.FilePathAction"] = {
        __ename__: ["albero", "command", "FilePathAction"],
        __constructs__: ["UPLOAD_PATH", "UPLOAD_MULTI_PATH", "DOWNLOAD_PATH"],
      })),
    Jn =
      ((Xn.UPLOAD_PATH = function (e, t, n, i, r) {
        e = ["UPLOAD_PATH", 0, e, t, n, i, r];
        return ((e.__enum__ = Xn), (e.toString = a), e);
      }),
      (Xn.UPLOAD_MULTI_PATH = function (e, t, n, i, r, o) {
        e = ["UPLOAD_MULTI_PATH", 1, e, t, n, i, r, o];
        return ((e.__enum__ = Xn), (e.toString = a), e);
      }),
      (Xn.DOWNLOAD_PATH = function (e, t, n) {
        e = ["DOWNLOAD_PATH", 2, e, t, n];
        return ((e.__enum__ = Xn), (e.toString = a), e);
      }),
      (Xn.__empty_constructs__ = []),
      function () {
        r.call(this);
      }),
    Zn =
      (((t["albero.command.FilePreviewCommand"] = Jn).__name__ = [
        "albero",
        "command",
        "FilePreviewCommand",
      ]),
      (Jn.__super__ = r),
      (Jn.prototype = e(r.prototype, {
        get: function (e) {
          return this.api.getFilePreview(e);
        },
        create: function (e) {
          return this.api.createFilePreview(e);
        },
        __class__: Jn,
      })),
      function () {
        r.call(this);
      }),
    $n =
      (((t["albero.command.LoadStampsetCommand"] = Zn).__name__ = [
        "albero",
        "command",
        "LoadStampsetCommand",
      ]),
      (Zn.__super__ = r),
      (Zn.prototype = e(r.prototype, {
        execute: function (e) {
          e.getBody();
          this.loadCommonStampset();
        },
        getOriginalStampset: function (t) {
          var n = ei.getInstance(),
            e = n.get(t);
          return (
            null != e ||
              ((e = this.api.getStampsets([t])),
              n.add(t, e),
              e.then(
                function (e) {
                  n.remove(t);
                },
                function (e) {
                  n.remove(t);
                },
              )),
            e
          );
        },
        getOriginalStamp: function (n) {
          var i = this,
            e = this.dataStore.getOriginalStamp(n.stampId);
          return null != e
            ? Promise.resolve(e)
            : this.getOriginalStampset(n.stampsetId).then(function (e) {
                var t = i.dataStore.getOriginalStamp(n.stampId);
                return null == t
                  ? Promise.reject("stamp_not_found")
                  : Promise.resolve(t);
              });
        },
        loadCommonStampset: function () {},
        __class__: Zn,
      })),
      (t["albero.command.LoadStampsetAction"] = {
        __ename__: ["albero", "command", "LoadStampsetAction"],
        __constructs__: ["COMMON"],
      })),
    ei =
      (($n.COMMON = ["COMMON", 0]),
      ($n.COMMON.toString = a),
      (($n.COMMON.__enum__ = $n).__empty_constructs__ = [$n.COMMON]),
      function () {
        this.loadings = new B();
      }),
    ti =
      (((t["albero.command.StampsetLoadingPromises"] = ei).__name__ = [
        "albero",
        "command",
        "StampsetLoadingPromises",
      ]),
      (ei.getInstance = function () {
        return (null == ei.instance && (ei.instance = new ei()), ei.instance);
      }),
      (ei.prototype = {
        get: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.loadings;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        add: function (e, t) {
          var e = "_" + e.high + "_" + e.low,
            n = this.loadings;
          null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
        },
        remove: function (e) {
          this.loadings.remove("_" + e.high + "_" + e.low);
        },
        __class__: ei,
      }),
      function () {
        r.call(this);
      }),
    $n =
      (((t["albero.command.ManageFriendsCommand"] = ti).__name__ = [
        "albero",
        "command",
        "ManageFriendsCommand",
      ]),
      (ti.__super__ = r),
      (ti.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t.action[1]) {
            case 0:
              for (var n = 0, i = t.users; n < i.length; ) {
                var r = i[n];
                (++n, this.api.addFriend(r));
              }
              break;
            case 1:
              for (var o = 0, a = t.users; o < a.length; ) {
                var s = a[o];
                (++o, this.api.deleteFriend(s));
              }
          }
        },
        __class__: ti,
      })),
      (t["albero.command.ManageFriendsAction"] = {
        __ename__: ["albero", "command", "ManageFriendsAction"],
        __constructs__: ["ADD", "DELETE"],
      })),
    ni =
      (($n.ADD = ["ADD", 0]),
      ($n.ADD.toString = a),
      (($n.ADD.__enum__ = $n).DELETE = ["DELETE", 1]),
      ($n.DELETE.toString = a),
      (($n.DELETE.__enum__ = $n).__empty_constructs__ = [$n.ADD, $n.DELETE]),
      function () {
        r.call(this);
      }),
    ii =
      (((t["albero.command.MessageCommand"] = ni).__name__ = [
        "albero",
        "command",
        "MessageCommand",
      ]),
      (ni.__super__ = r),
      (ni.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t[1]) {
            case 0:
              this.api.deleteMessage(t[2], t[3]);
              break;
            case 1:
              this.api.forwardMessages(t[2], t[3], 0);
              break;
            case 2:
              this.api.addFavoriteMessage(t[2], t[3], t[4]);
              break;
            case 3:
              this.api.deleteFavoriteMessage(t[2], t[3]);
              break;
            case 4:
              this.api.getFavoriteMessages(t[2], t[3], t[4]);
          }
        },
        __class__: ni,
      })),
      (t["albero.command.MessageAction"] = {
        __ename__: ["albero", "command", "MessageAction"],
        __constructs__: [
          "DELETE",
          "FORWARD",
          "ADD_FAVORITE",
          "DELETE_FAVORITE",
          "GET_FAVORITES",
        ],
      })),
    ri =
      ((ii.DELETE = function (e, t) {
        e = ["DELETE", 0, e, t];
        return ((e.__enum__ = ii), (e.toString = a), e);
      }),
      (ii.FORWARD = function (e, t) {
        e = ["FORWARD", 1, e, t];
        return ((e.__enum__ = ii), (e.toString = a), e);
      }),
      (ii.ADD_FAVORITE = function (e, t, n) {
        e = ["ADD_FAVORITE", 2, e, t, n];
        return ((e.__enum__ = ii), (e.toString = a), e);
      }),
      (ii.DELETE_FAVORITE = function (e, t) {
        e = ["DELETE_FAVORITE", 3, e, t];
        return ((e.__enum__ = ii), (e.toString = a), e);
      }),
      (ii.GET_FAVORITES = function (e, t, n) {
        e = ["GET_FAVORITES", 4, e, t, n];
        return ((e.__enum__ = ii), (e.toString = a), e);
      }),
      (ii.__empty_constructs__ = []),
      function () {
        r.call(this);
      }),
    oi =
      (((t["albero.command.MessageReactionCommand"] = ri).__name__ = [
        "albero",
        "command",
        "MessageReactionCommand",
      ]),
      (ri.__super__ = r),
      (ri.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t[1]) {
            case 0:
              this.api.setMessageReaction(t[2], t[3], t[4]);
              break;
            case 1:
              this.api.resetMessageReaction(t[2], t[3]);
              break;
            case 2:
              this.api.getMessageReactionSummaries(t[2], t[3]);
              break;
            case 3:
              this.api.getAvailableMessageReactions(t[2]);
          }
        },
        getMessageReactionUsersUsers: function (e, t) {
          var n = this.dataStore.getMessageReactionUsers(t);
          return null != n
            ? Promise.resolve(n)
            : this.api.getMessageReactionUsers(e, t);
        },
        __class__: ri,
      })),
      (t["albero.command.MessageReactionAction"] = {
        __ename__: ["albero", "command", "MessageReactionAction"],
        __constructs__: [
          "SET_MY_REACTION",
          "RESET_MY_REACTION",
          "GET_SUMMARIES",
          "GET_AVAILABLE_REACTIONS",
        ],
      })),
    ai =
      ((oi.SET_MY_REACTION = function (e, t, n) {
        e = ["SET_MY_REACTION", 0, e, t, n];
        return ((e.__enum__ = oi), (e.toString = a), e);
      }),
      (oi.RESET_MY_REACTION = function (e, t) {
        e = ["RESET_MY_REACTION", 1, e, t];
        return ((e.__enum__ = oi), (e.toString = a), e);
      }),
      (oi.GET_SUMMARIES = function (e, t) {
        e = ["GET_SUMMARIES", 2, e, t];
        return ((e.__enum__ = oi), (e.toString = a), e);
      }),
      (oi.GET_AVAILABLE_REACTIONS = function (e) {
        e = ["GET_AVAILABLE_REACTIONS", 3, e];
        return ((e.__enum__ = oi), (e.toString = a), e);
      }),
      (oi.__empty_constructs__ = []),
      function () {
        r.call(this);
      }),
    si =
      (((t["albero.command.NoteCommand"] = ai).__name__ = [
        "albero",
        "command",
        "NoteCommand",
      ]),
      (ai.__super__ = r),
      (ai.prototype = e(r.prototype, {
        execute: function (e) {
          var t = this,
            n = e.getBody();
          switch (n[1]) {
            case 0:
              this.api.getNoteStatuses(n[2], n[3]);
              break;
            case 1:
              this.api.getNote(n[2]);
              break;
            case 2:
              var i = n[2];
              Yo.delayForReplicationLag().then(function (e) {
                t.api.getNote(i);
              });
              break;
            case 3:
              this.api.createNoteV2(n[2], n[3], n[4]);
              break;
            case 4:
              var r = n[4],
                o = n[2],
                a = this.createNoteLocalEdit(o, n[3], n[5]);
              this.api.createNoteV2(o, a, r);
              break;
            case 5:
              this.api.updateNoteSetting(n[2], n[3], n[4]);
              break;
            case 6:
              this.api.updateNoteV2(n[2], n[3], n[4], n[5]);
              break;
            case 7:
              this.api.deleteNote(n[2]);
              break;
            case 8:
              this.dataStore.setNoteLocalEdit(n[2], null);
          }
        },
        createNoteLocalEdit: function (e, t, n) {
          var i = So.createDummyId(),
            r = this.createNoteAttachmentFileInfoList(e, t.attachments),
            o = t.contentType == go.RICH_TEXT;
          return new vo(e, i, t.title, t.content, r, n, null, o);
        },
        createNoteAttachmentFileInfoList: function (e, t) {
          for (var n = [], i = 0; i < t.length; ) {
            var r = t[i],
              r =
                (++i, this.fileService.createDummyFile(r.path, r.name, r.type));
            n.push(ar.createStagedFileInfoFromDummyFile(e, r, null));
          }
          return n;
        },
        updateNote: function (e, t, n, i) {
          return this.api.updateNoteV2(e, t, n, i);
        },
        lockNote: function (e, t) {
          return this.api.lockNote(e, t);
        },
        unlockNote: function (e, t) {
          return this.api.unlockNote(e, t);
        },
        __class__: ai,
      })),
      (t["albero.command.NoteAction"] = {
        __ename__: ["albero", "command", "NoteAction"],
        __constructs__: [
          "GET_NOTE_STATUSES",
          "GET_NOTE",
          "GET_NOTE_ON_NOTIFY_UPDATE",
          "CREATE_NOTE",
          "CREATE_NOTE_FOR_HUBOT",
          "UPDATE_NOTE_SETTING",
          "UPDATE_NOTE",
          "DELETE_NOTE",
          "DELETE_NOTE_LOCAL_EDIT",
        ],
      })),
    li =
      ((si.GET_NOTE_STATUSES = function (e, t) {
        e = ["GET_NOTE_STATUSES", 0, e, t];
        return ((e.__enum__ = si), (e.toString = a), e);
      }),
      (si.GET_NOTE = function (e) {
        e = ["GET_NOTE", 1, e];
        return ((e.__enum__ = si), (e.toString = a), e);
      }),
      (si.GET_NOTE_ON_NOTIFY_UPDATE = function (e, t) {
        e = ["GET_NOTE_ON_NOTIFY_UPDATE", 2, e, t];
        return ((e.__enum__ = si), (e.toString = a), e);
      }),
      (si.CREATE_NOTE = function (e, t, n) {
        e = ["CREATE_NOTE", 3, e, t, n];
        return ((e.__enum__ = si), (e.toString = a), e);
      }),
      (si.CREATE_NOTE_FOR_HUBOT = function (e, t, n, i) {
        e = ["CREATE_NOTE_FOR_HUBOT", 4, e, t, n, i];
        return ((e.__enum__ = si), (e.toString = a), e);
      }),
      (si.UPDATE_NOTE_SETTING = function (e, t, n) {
        e = ["UPDATE_NOTE_SETTING", 5, e, t, n];
        return ((e.__enum__ = si), (e.toString = a), e);
      }),
      (si.UPDATE_NOTE = function (e, t, n, i) {
        e = ["UPDATE_NOTE", 6, e, t, n, i];
        return ((e.__enum__ = si), (e.toString = a), e);
      }),
      (si.DELETE_NOTE = function (e) {
        e = ["DELETE_NOTE", 7, e];
        return ((e.__enum__ = si), (e.toString = a), e);
      }),
      (si.DELETE_NOTE_LOCAL_EDIT = function (e) {
        e = ["DELETE_NOTE_LOCAL_EDIT", 8, e];
        return ((e.__enum__ = si), (e.toString = a), e);
      }),
      (si.__empty_constructs__ = []),
      function () {
        r.call(this);
      }),
    ui =
      (((t["albero.command.ReadCommand"] = li).__name__ = [
        "albero",
        "command",
        "ReadCommand",
      ]),
      (li.__super__ = r),
      (li.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t[1]) {
            case 0:
              this.readStatusUpdater.updateReadStatuses(t[2], t[3]);
              break;
            case 1:
              this.readStatusUpdater.updateAnnouncementReadStatus(t[2], t[3]);
              break;
            case 2:
              this.api.getReadStatus(t[2], t[3]);
              break;
            case 3:
              this.keywordWatcher.readKeywordDetection(t[2], null, t[3]);
              break;
            case 4:
              this.keywordWatcher.readKeywordDetection(null, t[2], t[3]);
          }
        },
        __class__: li,
      })),
      (t["albero.command.ReadType"] = {
        __ename__: ["albero", "command", "ReadType"],
        __constructs__: [
          "TALK",
          "ANNOUNCEMENT",
          "READ_STATUS",
          "KEYWORD_IN_ANNOUNCEMENTS",
          "KEYWORD_IN_MESSAGES",
        ],
      })),
    ci =
      ((ui.TALK = function (e, t) {
        e = ["TALK", 0, e, t];
        return ((e.__enum__ = ui), (e.toString = a), e);
      }),
      (ui.ANNOUNCEMENT = function (e, t) {
        e = ["ANNOUNCEMENT", 1, e, t];
        return ((e.__enum__ = ui), (e.toString = a), e);
      }),
      (ui.READ_STATUS = function (e, t) {
        e = ["READ_STATUS", 2, e, t];
        return ((e.__enum__ = ui), (e.toString = a), e);
      }),
      (ui.KEYWORD_IN_ANNOUNCEMENTS = function (e, t) {
        e = ["KEYWORD_IN_ANNOUNCEMENTS", 3, e, t];
        return ((e.__enum__ = ui), (e.toString = a), e);
      }),
      (ui.KEYWORD_IN_MESSAGES = function (e, t) {
        e = ["KEYWORD_IN_MESSAGES", 4, e, t];
        return ((e.__enum__ = ui), (e.toString = a), e);
      }),
      (ui.__empty_constructs__ = []),
      function () {
        r.call(this);
      }),
    _i =
      (((t["albero.command.ReloadDataCommand"] = ci).__name__ = [
        "albero",
        "command",
        "ReloadDataCommand",
      ]),
      (ci.__super__ = r),
      (ci.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t[1]) {
            case 0:
              this.api.getMessages(t[2], t[3], t[4], t[5]);
              break;
            case 1:
              this.api.getAllUsers(t[2]);
              break;
            case 2:
              this.api.getUsers(t[2], t[3]);
              break;
            case 3:
              this.api.getProfile(t[2], t[3]);
              break;
            case 4:
              this.api.searchDomainUsers(t[2], t[3]);
              break;
            case 5:
              this.api.getAnnouncements(t[2], t[3], t[4]);
              break;
            case 6:
              this.api.getQuestions(t[2], t[3], t[4], t[5]);
              break;
            case 7:
              this.api.getQuestion(t[2], t[3]);
              break;
            case 8:
              this.api.getAttachments(t[2], t[3]);
              break;
            case 9:
              this.api.getSolutions(t[2], t[3]);
              break;
            case 10:
              this.api.getDepartmentTree(t[2]);
              break;
            case 11:
              this.api.getDepartmentUsers(t[2], t[3]);
              break;
            case 12:
              this.api.getDepartmentAllUsers(t[2]);
              break;
            case 13:
              this.api.getDepartmentUserCount(t[2], t[3]);
          }
        },
        __class__: ci,
      })),
      (t["albero.command.ReloadDataType"] = {
        __ename__: ["albero", "command", "ReloadDataType"],
        __constructs__: [
          "Messages",
          "AllUsers",
          "GetUsers",
          "GetProfile",
          "SearchDomainUsers",
          "Announcements",
          "Questions",
          "Question",
          "Files",
          "Solutions",
          "DepartmentTree",
          "DepartmentUsers",
          "DepartmentAllUsers",
          "DepartmentUserCount",
        ],
      })),
    hi =
      ((_i.Messages = function (e, t, n, i) {
        e = ["Messages", 0, e, t, n, i];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.AllUsers = function (e) {
        e = ["AllUsers", 1, e];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.GetUsers = function (e, t) {
        e = ["GetUsers", 2, e, t];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.GetProfile = function (e, t) {
        e = ["GetProfile", 3, e, t];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.SearchDomainUsers = function (e, t) {
        e = ["SearchDomainUsers", 4, e, t];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.Announcements = function (e, t, n) {
        e = ["Announcements", 5, e, t, n];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.Questions = function (e, t, n, i) {
        e = ["Questions", 6, e, t, n, i];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.Question = function (e, t) {
        e = ["Question", 7, e, t];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.Files = function (e, t) {
        e = ["Files", 8, e, t];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.Solutions = function (e, t) {
        e = ["Solutions", 9, e, t];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.DepartmentTree = function (e) {
        e = ["DepartmentTree", 10, e];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.DepartmentUsers = function (e, t) {
        e = ["DepartmentUsers", 11, e, t];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.DepartmentAllUsers = function (e) {
        e = ["DepartmentAllUsers", 12, e];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.DepartmentUserCount = function (e, t) {
        e = ["DepartmentUserCount", 13, e, t];
        return ((e.__enum__ = _i), (e.toString = a), e);
      }),
      (_i.__empty_constructs__ = []),
      function () {
        r.call(this);
      }),
    di =
      (((t["albero.command.ScheduleMessageCommand"] = hi).__name__ = [
        "albero",
        "command",
        "ScheduleMessageCommand",
      ]),
      (hi.__super__ = r),
      (hi.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t[1]) {
            case 0:
              this.api.getScheduledMessages(t[2], t[3], !0);
              break;
            case 1:
              this.api.getScheduledMessages(t[2], t[3], !1);
              break;
            case 2:
              this.api.scheduleMessage(t[2], t[3], t[4], t[5]);
              break;
            case 3:
              var n = t[3];
              this.api.scheduleFileMessage(t[2], n.domainId, n.id, t[4]);
              break;
            case 4:
              n = t[3];
              this.api.scheduleFilesMessage(t[2], n.domainId, n.id, t[4], t[5]);
              break;
            case 5:
              this.api.deleteScheduledMessage(t[2]);
              break;
            case 6:
              this.api.rescheduleMessage(t[2], t[3], t[4], t[5], t[6]);
              break;
            case 7:
              n = t[4];
              this.api.rescheduleFileMessage(
                t[2],
                t[3],
                n.domainId,
                n.id,
                t[5],
              );
              break;
            case 8:
              n = t[4];
              this.api.rescheduleFilesMessage(
                t[2],
                t[3],
                n.domainId,
                n.id,
                t[5],
                t[6],
              );
              break;
            case 9:
              this.api.sendScheduledMessageNow(t[2]);
          }
        },
        __class__: hi,
      })),
      (t["albero.command.ScheduleMessageAction"] = {
        __ename__: ["albero", "command", "ScheduleMessageAction"],
        __constructs__: [
          "RELOAD",
          "GET",
          "SCHEDULE",
          "SCHEDULE_UPLOAD",
          "SCHEDULE_UPLOAD_MULTI",
          "DELETE",
          "RESCHEDULE",
          "RESCHEDULE_UPLOAD",
          "RESCHEDULE_UPLOAD_MULTI",
          "SEND_NOW",
        ],
      })),
    fi =
      ((di.RELOAD = function (e, t) {
        e = ["RELOAD", 0, e, t];
        return ((e.__enum__ = di), (e.toString = a), e);
      }),
      (di.GET = function (e, t) {
        e = ["GET", 1, e, t];
        return ((e.__enum__ = di), (e.toString = a), e);
      }),
      (di.SCHEDULE = function (e, t, n, i) {
        e = ["SCHEDULE", 2, e, t, n, i];
        return ((e.__enum__ = di), (e.toString = a), e);
      }),
      (di.SCHEDULE_UPLOAD = function (e, t, n) {
        e = ["SCHEDULE_UPLOAD", 3, e, t, n];
        return ((e.__enum__ = di), (e.toString = a), e);
      }),
      (di.SCHEDULE_UPLOAD_MULTI = function (e, t, n, i) {
        e = ["SCHEDULE_UPLOAD_MULTI", 4, e, t, n, i];
        return ((e.__enum__ = di), (e.toString = a), e);
      }),
      (di.DELETE = function (e) {
        e = ["DELETE", 5, e];
        return ((e.__enum__ = di), (e.toString = a), e);
      }),
      (di.RESCHEDULE = function (e, t, n, i, r) {
        e = ["RESCHEDULE", 6, e, t, n, i, r];
        return ((e.__enum__ = di), (e.toString = a), e);
      }),
      (di.RESCHEDULE_UPLOAD = function (e, t, n, i) {
        e = ["RESCHEDULE_UPLOAD", 7, e, t, n, i];
        return ((e.__enum__ = di), (e.toString = a), e);
      }),
      (di.RESCHEDULE_UPLOAD_MULTI = function (e, t, n, i, r) {
        e = ["RESCHEDULE_UPLOAD_MULTI", 8, e, t, n, i, r];
        return ((e.__enum__ = di), (e.toString = a), e);
      }),
      (di.SEND_NOW = function (e) {
        e = ["SEND_NOW", 9, e];
        return ((e.__enum__ = di), (e.toString = a), e);
      }),
      (di.__empty_constructs__ = []),
      function () {
        r.call(this);
      }),
    mi =
      (((t["albero.command.SearchCommand"] = fi).__name__ = [
        "albero",
        "command",
        "SearchCommand",
      ]),
      (fi.__super__ = r),
      (fi.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t[1]) {
            case 0:
              var n = t[4],
                i = t[3],
                r = t[2];
              if (!this.searchService.isSearching())
                switch (
                  (null == n && this.searchService.saveParams(r),
                  r.searchType[1])
                ) {
                  case 0:
                    this.api.searchMessages(r.copy(), i, n);
                    break;
                  case 1:
                    this.api.searchAttachments(r.copy(), i, n);
                }
              break;
            case 1:
              (this.searchService.clearEditingParams(),
                this.searchService.clearRecentParams(),
                this.searchService.clearSearching(),
                this.sendNotification("notify_search_clear", {}));
          }
        },
        searchMessagesAroundDateTime: function (e, t) {
          return this.api.searchMessagesAroundDateTime(e, t);
        },
        __class__: fi,
      })),
      (t["albero.command.SearchAction"] = {
        __ename__: ["albero", "command", "SearchAction"],
        __constructs__: ["SEARCH", "CLEAR"],
      })),
    pi =
      ((mi.SEARCH = function (e, t, n) {
        e = ["SEARCH", 0, e, t, n];
        return ((e.__enum__ = mi), (e.toString = a), e);
      }),
      (mi.CLEAR = ["CLEAR", 1]),
      (mi.CLEAR.toString = a),
      ((mi.CLEAR.__enum__ = mi).__empty_constructs__ = [mi.CLEAR]),
      function () {
        r.call(this);
      }),
    gi =
      (((t["albero.command.SelectTalkCommand"] = pi).__name__ = [
        "albero",
        "command",
        "SelectTalkCommand",
      ]),
      (pi.__super__ = r),
      (pi.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t[1]) {
            case 0:
              var n = t[3],
                i = new Zr(t[2], $r.NotSelected);
              this.sendNotification(
                "Url",
                bi.REDIRECT(i.createUrls(n, this.settings)),
              );
              break;
            case 1:
              ((i = t[3]),
                (n = new Zr(t[2], $r.TalkSelected(i.id, m.Some(t[4])))));
              this.sendNotification(
                "Url",
                bi.REDIRECT(n.createUrls(i.domainId, this.settings)),
              );
              break;
            case 2:
              ((n = t[3]), (i = new Zr(t[2], $r.TalkSelected(n.id, m.None))));
              this.sendNotification(
                "Url",
                bi.REDIRECT(i.createUrls(n.domainId, this.settings)),
              );
              break;
            case 3:
              ((i = t[3]), (n = new Zr(t[2], $r.AnnouncementsSelected)));
              this.sendNotification(
                "Url",
                bi.REDIRECT(n.createUrls(i, this.settings)),
              );
              break;
            case 4:
              var r,
                o,
                a,
                n = t[4],
                i = t[3],
                s = t[2],
                l = this.getFallbackUrls(s, k.array(i));
              h.isEmpty(l)
                ? ((r = z(this, this.updateSelectionIfNeed)),
                  (o = s),
                  (a = n),
                  k.iter(i, function (e) {
                    r(o, a, e);
                  }))
                : ((s = h.get(l)),
                  this.sendNotification("routing_talk_fallback", s));
          }
        },
        updateSelectionIfNeed: function (e, t, n) {
          if (t || !this.isSelectionParametersNotChanged(e, n)) {
            var i = n.talkSelection;
            switch (i[1]) {
              case 0:
                this.settings.clearSelectedTalk(n.paneType);
                break;
              case 1:
                this.settings.selectTalk(n.paneType, e, i[2], i[3]);
                break;
              case 2:
                this.settings.selectAnnouncement(n.paneType);
            }
          }
        },
        isSelectionParametersNotChanged: function (e, t) {
          var n = this.settings.getSelectedDomainIdOnTalkSelectionChanged();
          return (
            null != e &&
            null != n &&
            e.high == n.high &&
            e.low == n.low &&
            eo.eqTalkSelection(
              t.talkSelection,
              this.settings.getTalkSelection(e, t.paneType),
            )
          );
        },
        getFallbackUrls: function (e, t) {
          var n,
            i = z(this, this.getValidTalkSelection),
            r = e,
            t = t.map(function (e) {
              return i(r, e);
            });
          return k.exists(t, h.isEmpty)
            ? ((n = t.map(z(this, this.getOrNotSelected))),
              (t = 1 == t.length ? Pn.single(n[0]) : Pn.multi(n)),
              m.Some(g.talks(e, t)))
            : m.None;
        },
        getOrNotSelected: function (e) {
          switch (e[1]) {
            case 0:
              return e[2];
            case 1:
              return $r.NotSelected;
          }
        },
        getValidTalkSelection: function (e, t) {
          var n = t.talkSelection;
          switch (n[1]) {
            case 0:
              break;
            case 1:
              var i = this.dataStore.getTalk(n[2]);
              if (
                null == i ||
                !(
                  null != (i = i.domainId) &&
                  null != e &&
                  i.high == e.high &&
                  i.low == e.low
                )
              )
                return m.None;
              break;
            case 2:
              var i = this.dataStore.getDomain(e),
                r = this.dataStore.getAnnouncementStatus(e);
              if (Oi.hasAnnouncement(i, r)) break;
              return m.None;
          }
          return m.Some(t.talkSelection);
        },
        __class__: pi,
      })),
      (t["albero.command.SelectTalkAction"] = {
        __ename__: ["albero", "command", "SelectTalkAction"],
        __constructs__: [
          "SelectNone",
          "SelectTalkAndMessage",
          "SelectTalk",
          "SelectAnnouncement",
          "UpdateSelection",
        ],
      })),
    yi =
      ((gi.SelectNone = function (e, t) {
        e = ["SelectNone", 0, e, t];
        return ((e.__enum__ = gi), (e.toString = a), e);
      }),
      (gi.SelectTalkAndMessage = function (e, t, n) {
        e = ["SelectTalkAndMessage", 1, e, t, n];
        return ((e.__enum__ = gi), (e.toString = a), e);
      }),
      (gi.SelectTalk = function (e, t) {
        e = ["SelectTalk", 2, e, t];
        return ((e.__enum__ = gi), (e.toString = a), e);
      }),
      (gi.SelectAnnouncement = function (e, t) {
        e = ["SelectAnnouncement", 3, e, t];
        return ((e.__enum__ = gi), (e.toString = a), e);
      }),
      (gi.UpdateSelection = function (e, t, n) {
        e = ["UpdateSelection", 4, e, t, n];
        return ((e.__enum__ = gi), (e.toString = a), e);
      }),
      (gi.__empty_constructs__ = []),
      function () {
        r.call(this);
      }),
    Si =
      (((t["albero.command.SendCommand"] = yi).__name__ = [
        "albero",
        "command",
        "SendCommand",
      ]),
      (yi.__super__ = r),
      (yi.prototype = e(r.prototype, {
        execute: function (e) {
          var t,
            e = e.getBody();
          ((t = null != e.id ? "_" + (t = e.id).high + "_" + t.low : null),
            null != e.talkId
              ? this.api.createMessage(e.talkId, e.type, e.content, t)
              : this.api.createAnnouncement(e.domainId, e.type, e.content, t));
        },
        __class__: yi,
      })),
      function () {
        r.call(this);
      }),
    vi =
      (((t["albero.command.SignInCommand"] = Si).__name__ = [
        "albero",
        "command",
        "SignInCommand",
      ]),
      (Si.__super__ = r),
      (Si.prototype = e(r.prototype, {
        execute: function (e) {
          this.accessTokenResolver.asyncGetAccessToken(
            e.getBody(),
            z((j = this.session), j.start),
          );
        },
        __class__: Si,
      })),
      function () {
        r.call(this);
      }),
    Ti =
      (((t["albero.command.SignOutCommand"] = vi).__name__ = [
        "albero",
        "command",
        "SignOutCommand",
      ]),
      (vi.__super__ = r),
      (vi.prototype = e(r.prototype, {
        execute: function (e) {
          (this.api.deleteSession(),
            this.settings.clearInputTextForAll(),
            this.dataStore.clear(!0),
            this.sendNotification("signed_out"));
        },
        __class__: vi,
      })),
      function () {
        r.call(this);
      }),
    Ii =
      (((t["albero.command.TalkCommand"] = Ti).__name__ = [
        "albero",
        "command",
        "TalkCommand",
      ]),
      (Ti.__super__ = r),
      (Ti.prototype = e(r.prototype, {
        execute: function (e) {
          var t = this,
            n = e.getBody();
          switch (n[1]) {
            case 0:
              var i = n[4],
                r = n[3],
                o = n[2],
                a = this.dataStore.getValidPairTalk(r, o);
              null != a
                ? (this.sendNotification("SelectTalk", gi.SelectTalk(i, a)),
                  this.sendNotification("open_existing_pair_talk_completed"))
                : this.api.createPairTalk(o, r, i);
              break;
            case 1:
              this.api.createGroupTalk(n[2], n[3], n[4], n[5]);
              break;
            case 2:
              this.api.addTalkers(n[2], n[3], n[4]);
              break;
            case 3:
              a = this.dataStore.me.id;
              this.api.deleteTalker(n[2], a);
              break;
            case 4:
              ((o = n[3]), (r = this.dataStore.getTalk(n[2])));
              (null == o && (o = this.dataStore.me.id),
                this.api.deleteTalker(r, o));
              break;
            case 5:
              var s = n[6],
                i = n[5],
                a = n[4],
                l = n[3],
                u = n[2];
              u.type == Xr.GroupTalk &&
                (null == s && (s = G.__cast(u, qr).isSystemMessageVisibled()),
                this.prepareIconUrl(u.domainId, a, i, function (e) {
                  t.api.updateGroupTalk(u, l, e, s);
                }));
              break;
            case 6:
              var c = n[6],
                r = n[5],
                o = n[4],
                _ = n[3],
                h = this.dataStore.getTalk(n[2]);
              h.type == Xr.GroupTalk &&
                (null == c && (c = G.__cast(h, qr).isSystemMessageVisibled()),
                this.prepareIconUrl(h.domainId, o, r, function (e) {
                  t.api.updateGroupTalk(h, _, e, c);
                }));
              break;
            case 7:
              this.api.addFavoriteTalk(n[2], n[3]);
              break;
            case 8:
              this.api.deleteFavoriteTalk(n[2], n[3]);
              break;
            case 9:
              this.api.disablePushNotification(n[2], n[3]);
              break;
            case 10:
              this.api.enablePushNotification(n[2], n[3], n[4]);
          }
        },
        prepareIconUrl: function (e, t, n, i) {
          var r = this;
          null == n && null != t
            ? this.api.uploadFile(
                t,
                e,
                u.TALK_ICON,
                function (e) {
                  i(e.get_url);
                },
                function (e) {
                  r.sendNotification("update_group_talk_fail");
                },
              )
            : i(n);
        },
        __class__: Ti,
      })),
      (t["albero.command.TalkAction"] = {
        __ename__: ["albero", "command", "TalkAction"],
        __constructs__: [
          "OPEN_PAIR_TALK",
          "OPEN_GROUP_TALK",
          "ADD",
          "DELETE",
          "DELETE_FOR_HUBOT",
          "UPDATE",
          "UPDATE_FOR_HUBOT",
          "ADD_FAVORITE",
          "DELETE_FAVORITE",
          "DISABLE_PUSH_NOTIFICATION",
          "ENABLE_PUSH_NOTIFICATION",
        ],
      })),
    Ei =
      ((Ii.OPEN_PAIR_TALK = function (e, t, n) {
        e = ["OPEN_PAIR_TALK", 0, e, t, n];
        return ((e.__enum__ = Ii), (e.toString = a), e);
      }),
      (Ii.OPEN_GROUP_TALK = function (e, t, n, i) {
        e = ["OPEN_GROUP_TALK", 1, e, t, n, i];
        return ((e.__enum__ = Ii), (e.toString = a), e);
      }),
      (Ii.ADD = function (e, t, n) {
        e = ["ADD", 2, e, t, n];
        return ((e.__enum__ = Ii), (e.toString = a), e);
      }),
      (Ii.DELETE = function (e) {
        e = ["DELETE", 3, e];
        return ((e.__enum__ = Ii), (e.toString = a), e);
      }),
      (Ii.DELETE_FOR_HUBOT = function (e, t) {
        e = ["DELETE_FOR_HUBOT", 4, e, t];
        return ((e.__enum__ = Ii), (e.toString = a), e);
      }),
      (Ii.UPDATE = function (e, t, n, i, r) {
        e = ["UPDATE", 5, e, t, n, i, r];
        return ((e.__enum__ = Ii), (e.toString = a), e);
      }),
      (Ii.UPDATE_FOR_HUBOT = function (e, t, n, i, r) {
        e = ["UPDATE_FOR_HUBOT", 6, e, t, n, i, r];
        return ((e.__enum__ = Ii), (e.toString = a), e);
      }),
      (Ii.ADD_FAVORITE = function (e, t) {
        e = ["ADD_FAVORITE", 7, e, t];
        return ((e.__enum__ = Ii), (e.toString = a), e);
      }),
      (Ii.DELETE_FAVORITE = function (e, t) {
        e = ["DELETE_FAVORITE", 8, e, t];
        return ((e.__enum__ = Ii), (e.toString = a), e);
      }),
      (Ii.DISABLE_PUSH_NOTIFICATION = function (e, t) {
        e = ["DISABLE_PUSH_NOTIFICATION", 9, e, t];
        return ((e.__enum__ = Ii), (e.toString = a), e);
      }),
      (Ii.ENABLE_PUSH_NOTIFICATION = function (e, t, n) {
        e = ["ENABLE_PUSH_NOTIFICATION", 10, e, t, n];
        return ((e.__enum__ = Ii), (e.toString = a), e);
      }),
      (Ii.__empty_constructs__ = []),
      function () {
        r.call(this);
      }),
    wi =
      (((t["albero.command.UpdateProfileCommand"] = Ei).__name__ = [
        "albero",
        "command",
        "UpdateProfileCommand",
      ]),
      (Ei.__super__ = r),
      (Ei.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody().domainId,
            n = e.getBody().profileItemValues,
            e = e.getBody().copyToAllDomains,
            i = this.dataStore.getDomain(t);
          if (null != i && null != i.profileDefinition.itemDefinitions) {
            var r = [];
            if (e) {
              for (var o = [], a = 0; a < n.length; ) {
                var s = n[a],
                  l =
                    (++a,
                    Er.findById(
                      i.profileDefinition.itemDefinitions,
                      s.profileItemId,
                    ));
                null != l && o.push({ name: l.name, value: s.value });
              }
              for (var u = 0, c = this.dataStore.getDomains(); u < c.length; ) {
                var _ = c[u];
                if ((++u, !_.domainInfo.frozen)) {
                  for (
                    var h = [], d = 0, f = _.profileDefinition.itemDefinitions;
                    d < f.length;
                  ) {
                    var m = f[d];
                    if ((++d, m.visible && !m.locked))
                      for (var p = 0; p < o.length; ) {
                        var g = o[p];
                        if ((++p, g.name == m.name)) {
                          var y = new wr();
                          ((y.profileItemId = m.profileItemId),
                            (y.value = g.value),
                            h.push(y));
                          break;
                        }
                      }
                  }
                  0 < h.length && r.push(new Nr(_.id, h));
                }
              }
            } else {
              for (var S = [], v = (U.iter(n), 0); v < n.length; ) {
                var T = n[v],
                  I =
                    (++v,
                    Er.findById(
                      i.profileDefinition.itemDefinitions,
                      T.profileItemId,
                    ));
                null != I && I.visible && !I.locked && S.push(T);
              }
              0 < S.length && r.push(new Nr(i.id, S));
            }
            0 == r.length
              ? this.sendNotification(
                  "update_profile_responsed",
                  this.dataStore.me,
                )
              : this.api.updateProfile(r);
          }
        },
        __class__: Ei,
      })),
      function () {
        r.call(this);
      }),
    Ni =
      (((t["albero.command.UpdateUserCommand"] = wi).__name__ = [
        "albero",
        "command",
        "UpdateUserCommand",
      ]),
      (wi.__super__ = r),
      (wi.prototype = e(r.prototype, {
        execute: function (e) {
          e = e.getBody();
          if (
            null != e.profileImage &&
            !I.startsWith(e.profileImage.type, "image/")
          )
            return;
          this.api.updateUser(
            e.displayName,
            e.profileImage,
            e.profileImageUrl,
            e.phoneticDisplayName,
            e.status,
          );
        },
        __class__: wi,
      })),
      function () {
        r.call(this);
      }),
    Ai =
      (((t["albero.command.UpdateUserPresencesCommand"] = Ni).__name__ = [
        "albero",
        "command",
        "UpdateUserPresencesCommand",
      ]),
      (Ni.__super__ = r),
      (Ni.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody().domainId,
            e = e.getBody().userIds;
          null != t &&
            null != e &&
            0 != (e = this.filterOnlyExpired(D.unique(e))).length &&
            this.updateIfNeedPerSub(t, e);
        },
        filterOnlyExpired: function (e) {
          var n = this,
            i = W.nowAsInt64(),
            t = this.settings.getConfiguration().presenceExpiration,
            r = new E(t >> 31, t);
          return e.filter(function (e) {
            var t = h.map(n.userPresences.getUserPresence(e), function (e) {
              return e.isExpired(i, r);
            });
            switch (t[1]) {
              case 0:
                return t[2];
              case 1:
                return !0;
            }
          });
        },
        updateIfNeedPerSub: function (n, e) {
          for (
            var i = this,
              t = Promise.resolve([]),
              r = D.splitUsingMaxCount(e, 100),
              o = 0;
            o < r.length;
          ) {
            var a = [r[o]];
            (++o,
              (t = t.then(
                (function (t) {
                  return function (e) {
                    return i.api.getPresences(n, t[0]);
                  };
                })(a),
              )));
          }
          t.catch(function (e) {});
        },
        __class__: Ni,
      })),
      function () {
        r.call(this);
      }),
    bi =
      (((t["albero.command.UrlCommand"] = Ai).__name__ = [
        "albero",
        "command",
        "UrlCommand",
      ]),
      (Ai.__super__ = r),
      (Ai.prototype = e(r.prototype, {
        execute: function (e) {
          var t = e.getBody();
          switch (t[1]) {
            case 0:
              this.routing.forward(t[2]);
              break;
            case 1:
              this.routing.redirect(t[2]);
              break;
            case 2:
              this.routing.redirect(t[2], t[3]);
              break;
            case 3:
              this.routing.back();
          }
        },
        __class__: Ai,
      })),
      (t["albero.command.UrlAction"] = {
        __ename__: ["albero", "command", "UrlAction"],
        __constructs__: [
          "FORWARD",
          "REDIRECT",
          "REDIRECT_WITH_HISTORY_STATE",
          "BACK",
        ],
      })),
    Mi =
      ((bi.FORWARD = function (e) {
        e = ["FORWARD", 0, e];
        return ((e.__enum__ = bi), (e.toString = a), e);
      }),
      (bi.REDIRECT = function (e) {
        e = ["REDIRECT", 1, e];
        return ((e.__enum__ = bi), (e.toString = a), e);
      }),
      (bi.REDIRECT_WITH_HISTORY_STATE = function (e, t) {
        e = ["REDIRECT_WITH_HISTORY_STATE", 2, e, t];
        return ((e.__enum__ = bi), (e.toString = a), e);
      }),
      (bi.BACK = ["BACK", 3]),
      (bi.BACK.toString = a),
      ((bi.BACK.__enum__ = bi).__empty_constructs__ = [bi.BACK]),
      ((t["albero.entity.Account"] = oe).__name__ = [
        "albero",
        "entity",
        "Account",
      ]),
      (oe.prototype = { __class__: oe }),
      ((t["albero.entity.AccountControlGroup"] = ae).__name__ = [
        "albero",
        "entity",
        "AccountControlGroup",
      ]),
      (ae.prototype = {
        isAllowedUpdateProfileImage: function () {
          return (
            null == this.profilePolicy ||
            this.profilePolicy.allowUpdateProfileImage
          );
        },
        isAllowedUpdateDisplayName: function () {
          return (
            null == this.profilePolicy ||
            this.profilePolicy.allowUpdateDisplayName
          );
        },
        __class__: ae,
      }),
      function (e) {
        null != e &&
          ((this.name = e.name),
          (this.alias = e.alias),
          (this.ownerName = e.owner_name),
          (this.version = e.version));
      }),
    Ci =
      (((t["albero.entity.AccountControlGroupGroup"] = Mi).__name__ = [
        "albero",
        "entity",
        "AccountControlGroupGroup",
      ]),
      (Mi.prototype = { __class__: Mi }),
      function (e) {
        null != e &&
          ((this.allowUpdateDisplayName = e.allow_update_display_name),
          (this.allowUpdateProfileImage = e.allow_update_profile_image),
          (this.version = e.version));
      }),
    ki =
      (((t["albero.entity.AccountControlGroupProfilePolicy"] = Ci).__name__ = [
        "albero",
        "entity",
        "AccountControlGroupProfilePolicy",
      ]),
      (Ci.prototype = { __class__: Ci }),
      ((t["albero.entity.AccountControlGroupPartialUpdate"] = se).__name__ = [
        "albero",
        "entity",
        "AccountControlGroupPartialUpdate",
      ]),
      (se.prototype = { __class__: se }),
      ((t["albero.entity.AccountControlRequest"] = le).__name__ = [
        "albero",
        "entity",
        "AccountControlRequest",
      ]),
      (le.prototype = { __class__: le }),
      (t["albero.entity.AllowAttachmentTypeValue"] = {
        __ename__: ["albero", "entity", "AllowAttachmentTypeValue"],
        __constructs__: [
          "ng",
          "ok",
          "image",
          "imageOrVideo",
          "imageOrVideoOrAudio",
          "other",
        ],
      })),
    Di =
      ((ki.ng = ["ng", 0]),
      (ki.ng.toString = a),
      ((ki.ng.__enum__ = ki).ok = ["ok", 1]),
      (ki.ok.toString = a),
      ((ki.ok.__enum__ = ki).image = ["image", 2]),
      (ki.image.toString = a),
      ((ki.image.__enum__ = ki).imageOrVideo = ["imageOrVideo", 3]),
      (ki.imageOrVideo.toString = a),
      ((ki.imageOrVideo.__enum__ = ki).imageOrVideoOrAudio = [
        "imageOrVideoOrAudio",
        4,
      ]),
      (ki.imageOrVideoOrAudio.toString = a),
      ((ki.imageOrVideoOrAudio.__enum__ = ki).other = ["other", 5]),
      (ki.other.toString = a),
      ((ki.other.__enum__ = ki).__empty_constructs__ = [
        ki.ng,
        ki.ok,
        ki.image,
        ki.imageOrVideo,
        ki.imageOrVideoOrAudio,
        ki.other,
      ]),
      function (e) {
        this.value = e;
      }),
    Ri =
      (((t["albero.entity.AllowAttachmentType"] = Di).__name__ = [
        "albero",
        "entity",
        "AllowAttachmentType",
      ]),
      (Di.fromInt = function (e) {
        if (null == e) return new Di(ki.other);
        var t;
        if (null == e) t = ki.other;
        else
          switch (e) {
            case 0:
              t = ki.ng;
              break;
            case 1:
              t = ki.ok;
              break;
            case 2:
              t = ki.image;
              break;
            case 3:
              t = ki.imageOrVideo;
              break;
            case 4:
              t = ki.imageOrVideoOrAudio;
              break;
            default:
              t = ki.other;
          }
        return new Di(t);
      }),
      (Di.fromValue = function (e) {
        return new Di(e);
      }),
      (Di.createAttachmentErrorMessage = function (e, t, n) {
        return (
          null == n && (n = ""),
          0 < e.compareStrongness(t)
            ? e.createAttachmentErrorMessageByNetwork(n)
            : t.createAttachmentErrorMessageByDomain(n)
        );
      }),
      (Di.prototype = {
        compareStrongness: function (e) {
          return this.getStrongness() - e.getStrongness();
        },
        getStrongness: function () {
          switch (this.value[1]) {
            case 0:
              return 40;
            case 2:
              return 30;
            case 3:
              return 20;
            case 4:
              return 10;
            case 1:
            case 5:
              return 0;
          }
        },
        createAttachmentErrorMessageByDomain: function (e) {
          switch ((null == e && (e = ""), this.value[1])) {
            case 0:
              return c.localize("AllowAttachmentType.ban_attach_file") + e;
            case 2:
              return (
                c.localize("AllowAttachmentType.ban_attach_file_without_img") +
                e
              );
            case 3:
              return (
                c.localize(
                  "AllowAttachmentType.ban_attach_file_without_img_movie",
                ) + e
              );
            case 4:
              return (
                c.localize(
                  "AllowAttachmentType.ban_attach_file_without_img_movie_sound",
                ) + e
              );
            default:
              return "";
          }
        },
        createAttachmentErrorMessageByNetwork: function (e) {
          switch ((null == e && (e = ""), this.value[1])) {
            case 0:
              return (
                c.localize("AllowAttachmentType.network_ban_attach_file") + e
              );
            case 2:
              return (
                c.localize(
                  "AllowAttachmentType.network_ban_attach_file_without_img",
                ) + e
              );
            case 3:
              return (
                c.localize(
                  "AllowAttachmentType.network_ban_attach_file_without_img_movie",
                ) + e
              );
            case 4:
              return (
                c.localize(
                  "AllowAttachmentType.network_ban_attach_file_without_img_movie_sound",
                ) + e
              );
            default:
              return "";
          }
        },
        createAttachmentFilter: function () {
          switch (this.value[1]) {
            case 0:
              return function (e) {
                return !1;
              };
            case 2:
              return function (e) {
                return new v("image.*", "").match(e.type);
              };
            case 3:
              return function (e) {
                return new v("image.*|video.*", "").match(e.type);
              };
            case 4:
              return function (e) {
                return new v("image.*|video.*|audio.*", "").match(e.type);
              };
            default:
              return function (e) {
                return !0;
              };
          }
        },
        createDataTransferItemFilter: function () {
          switch (this.value[1]) {
            case 0:
              return function (e) {
                return !1;
              };
            case 2:
              var t = new v("image.*", "");
              return function (e) {
                return t.match(e.type);
              };
            case 3:
              var n = new v("image.*|video.*", "");
              return function (e) {
                return n.match(e.type);
              };
            case 4:
              var i = new v("image.*|video.*|audio.*", "");
              return function (e) {
                return i.match(e.type);
              };
            default:
              return function (e) {
                return !0;
              };
          }
        },
        __class__: Di,
      }),
      function (e) {
        null != e &&
          ((this.id = e.announcement_id),
          (this.domainId = e.domain_id),
          (this.groupId = e.group_id),
          (this.groupName = e.group_name),
          (this.type = O.typeOf(e.type)),
          (this.content = e.content),
          (this.userId = e.user_id),
          (this.userName = e.user_name),
          (this.createdAt = e.created_at));
      }),
    Oi =
      (((t["albero.entity.Announcement"] = Ri).__name__ = [
        "albero",
        "entity",
        "Announcement",
      ]),
      (Ri.prototype = {
        getDisplayTextWithoutEscape: function () {
          return O.getDisplayTextWithoutEscape(this.type, this.content);
        },
        getNotificatoinTextWithoutEscape: function (e) {
          switch (
            (null == (e = null == e ? this.userName : e) && (e = ""),
            this.type[1])
          ) {
            case 1:
            case 7:
              return e + ":" + this.getDisplayTextWithoutEscape();
            default:
              return (
                F.e(
                  "Unsupported announcement.type is detected. announcment:%o",
                  this,
                ),
                O.getUnsupportText()
              );
          }
        },
        markAsDeleted: function () {
          ((this.type = d.deleted), (this.content = ""));
        },
        __class__: Ri,
      }),
      ((t["albero.entity.AnnouncementStatus"] = ue).__name__ = [
        "albero",
        "entity",
        "AnnouncementStatus",
      ]),
      (ue.fromDomainId = function (e) {
        var t = new ue();
        return ((t.domainId = e), t);
      }),
      (ue.prototype = {
        calcUnreadCount: function () {
          var e = null == this.unreadCount ? 0 : this.unreadCount,
            t = this.maxReadAnnouncementId;
          return null != this.readingAnnouncements &&
            (this.readingAnnouncements.isReadingAll() ||
              (e =
                e -
                this.readingAnnouncements.getUnreadAnnouncementIds(t).length) <
                0)
            ? 0
            : e;
        },
        calcMaxReadAnnouncementId: function () {
          var e, t, n;
          return null == this.readingAnnouncements
            ? this.maxReadAnnouncementId
            : this.readingAnnouncements.isReadingAll()
              ? this.maxAnnouncementId
              : ((e = this.readingAnnouncements.getLastAnnouncementId()),
                null != this.maxReadAnnouncementId &&
                ((t = this.maxReadAnnouncementId),
                0 == (n = (e.high - t.high) | 0) &&
                  (n = w.ucompare(e.low, t.low)),
                (e.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1) < 0)
                  ? this.maxReadAnnouncementId
                  : e);
        },
        updateReadWithStatusUpdate: function (e) {
          var t = 0,
            n = null;
          if (null != e.readAnnouncementIds)
            for (var i = 0, r = e.readAnnouncementIds; i < r.length; ) {
              var o,
                a,
                s = r[i];
              (++i,
                (null == this.maxReadAnnouncementId ||
                  ((a = this.maxReadAnnouncementId),
                  0 == (o = (s.high - a.high) | 0) &&
                    (o = w.ucompare(s.low, a.low)),
                  0 <
                    (s.high < 0
                      ? a.high < 0
                        ? o
                        : -1
                      : 0 <= a.high
                        ? o
                        : 1))) &&
                  ((null == n ||
                    (0 == (a = (s.high - n.high) | 0) &&
                      (a = w.ucompare(s.low, n.low)),
                    0 <
                      (s.high < 0
                        ? n.high < 0
                          ? a
                          : -1
                        : 0 <= n.high
                          ? a
                          : 1))) &&
                    (n = s),
                  ++t));
            }
          null != n && this.read(n, t);
        },
        updateByAnnouncementDeletion: function (e) {
          var t,
            n = !1;
          return (
            this.isUnread(e) && (this.decrementUnreadCount(1), (n = !0)),
            null != this.maxAnnouncementId &&
              null != (t = this.maxAnnouncementId) &&
              null != e &&
              t.high == e.high &&
              t.low == e.low &&
              null != this.maxAnnouncement &&
              ((this.maxAnnouncement.type = d.deleted),
              (n = !(this.maxAnnouncement.content = ""))),
            n
          );
        },
        updateByAnnouncement: function (e) {
          (this.isNewMax(e.id) &&
            ((this.maxAnnouncementId = e.id), (this.maxAnnouncement = e)),
            this.unreadCount++);
        },
        updateByReadingAnnouncementsStarted: function (e) {
          if (null == this.maxAnnouncementId) return !1;
          var t = new xr(e, this.maxAnnouncementId);
          if (t.canRegardAnnouncementReadCountZero(this.maxAnnouncementId)) {
            if (!Oi.isUnreadAnnouncementExisted(this)) return !1;
          } else if (0 == e.filter(z(this, this.isUnread)).length) return !1;
          return null != this.readingAnnouncements
            ? this.readingAnnouncements.update(t)
            : ((this.readingAnnouncements = t), !0);
        },
        updateByReadingAnnouncementsFinished: function (e) {
          var t, n;
          null == this.readingAnnouncements ||
            null == (n = this.readingAnnouncements.getLastAnnouncementId()) ||
            (0 == (t = (n.high - e.high) | 0) && (t = w.ucompare(n.low, e.low)),
            0 < (n.high < 0 ? (e.high < 0 ? t : -1) : 0 <= e.high ? t : 1)) ||
            (null != e &&
              null != this.maxReadAnnouncementId &&
              (0 ==
                (t = ((n = this.maxReadAnnouncementId).high - e.high) | 0) &&
                (t = w.ucompare(n.low, e.low)),
              0 <=
                (n.high < 0 ? (e.high < 0 ? t : -1) : 0 <= e.high ? t : 1))) ||
            (this.readingAnnouncements.canRegardAnnouncementReadCountZero(
              this.maxAnnouncementId,
            )
              ? this.readAll()
              : ((n = this.readingAnnouncements.getLastAnnouncementId()),
                (e = this.readingAnnouncements.getUnreadAnnouncementIds(
                  this.maxReadAnnouncementId,
                )),
                this.read(n, e.length)),
            (this.readingAnnouncements = null));
        },
        updateByReadingAnnouncementsCanceled: function (e) {
          var t, n;
          null == this.readingAnnouncements ||
            null == (t = this.readingAnnouncements.getLastAnnouncementId()) ||
            (0 == (n = (t.high - e.high) | 0) && (n = w.ucompare(t.low, e.low)),
            0 < (t.high < 0 ? (e.high < 0 ? n : -1) : 0 <= e.high ? n : 1)) ||
            (this.readingAnnouncements = null);
        },
        read: function (e, t) {
          ((this.maxReadAnnouncementId = e),
            this.decrementUnreadCount(t),
            F.v(
              "Announcement status is updated. [domainId, maxReadAnnoundementId, unreadCount]:",
              this.domainId,
              this.maxReadAnnouncementId,
              this.unreadCount,
            ));
        },
        decrementUnreadCount: function (e) {
          ((this.unreadCount -= e),
            this.unreadCount < 0 && (this.unreadCount = 0));
        },
        readAll: function () {
          ((this.maxReadAnnouncementId = this.maxAnnouncementId),
            (this.unreadCount = 0));
        },
        isUnread: function (e) {
          var t, n;
          return (
            null == this.maxReadAnnouncementId ||
            (0 == (n = ((t = this.maxReadAnnouncementId).high - e.high) | 0) &&
              (n = w.ucompare(t.low, e.low)),
            (t.high < 0 ? (e.high < 0 ? n : -1) : 0 <= e.high ? n : 1) < 0)
          );
        },
        isMax: function (e) {
          var t;
          return (
            null != this.maxAnnouncementId &&
            null != (t = this.maxAnnouncementId) &&
            null != e &&
            t.high == e.high &&
            t.low == e.low
          );
        },
        isNewMax: function (e) {
          var t, n;
          return (
            null == this.maxAnnouncementId ||
            (0 == (n = ((t = this.maxAnnouncementId).high - e.high) | 0) &&
              (n = w.ucompare(t.low, e.low)),
            (t.high < 0 ? (e.high < 0 ? n : -1) : 0 <= e.high ? n : 1) < 0)
          );
        },
        clearReadingAnnouncements: function () {
          this.readingAnnouncements = null;
        },
        __class__: ue,
      }),
      function () {}),
    Fi =
      (((t["albero.entity.AnnouncementStatusHelper"] = Oi).__name__ = [
        "albero",
        "entity",
        "AnnouncementStatusHelper",
      ]),
      (Oi.isUnreadAnnouncementExisted = function (e) {
        return (
          null != e &&
          null != e.maxAnnouncement &&
          e.isUnread(e.maxAnnouncement.id)
        );
      }),
      (Oi.hasAnnouncement = function (e, t) {
        return (
          !(null == e || !e.role.allowReadAnnouncements || null == t) &&
          null != t.maxAnnouncementId
        );
      }),
      ((t["albero.entity.AnnouncementStatusUpdate"] = ce).__name__ = [
        "albero",
        "entity",
        "AnnouncementStatusUpdate",
      ]),
      (ce.prototype = { __class__: ce }),
      function () {
        this.announcements = new B();
      }),
    xi =
      (((t["albero.entity.AnnouncementStore"] = Fi).__name__ = [
        "albero",
        "entity",
        "AnnouncementStore",
      ]),
      (Fi.prototype = {
        getAnnouncement: function (e) {
          var t;
          return null == e
            ? null
            : ((e = "_" + e.high + "_" + e.low),
              (t = this.announcements),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        setAnnouncement: function (e) {},
        removeAnnouncement: function (e) {
          this.announcements.remove("_" + e.high + "_" + e.low);
        },
        __class__: Fi,
      }),
      ((t["albero.entity.BrowserSettings"] = _e).__name__ = [
        "albero",
        "entity",
        "BrowserSettings",
      ]),
      (_e.fromJson = function (e) {
        return new _e(e);
      }),
      (_e.prototype = {
        createCopy: function () {
          var e = new _e(null);
          return (this.copyTo(e), e);
        },
        equalNotificationSettings: function (e) {
          return (
            e.notificationInvisibleTalkSound ==
              this.notificationInvisibleTalkSound &&
            e.notificationInvisibleTalkPopup ==
              this.notificationInvisibleTalkPopup &&
            e.notificationVisibleTalkSound ==
              this.notificationVisibleTalkSound &&
            e.notificationVisibleTalkPopup ==
              this.notificationVisibleTalkPopup &&
            e.notificationIncludingMessageContent ==
              this.notificationIncludingMessageContent &&
            e.notificationPopupTime == this.notificationPopupTime
          );
        },
        equalTalkSettings: function (e) {
          return (
            e.talkAutoScroll == this.talkAutoScroll &&
            e.talksSortMode == this.talksSortMode &&
            e.talkRightPanelOpened == this.talkRightPanelOpened &&
            e.talkRightPanelWidth == this.talkRightPanelWidth &&
            e.legacyNoteEnabled == this.legacyNoteEnabled &&
            e.leaveTalksEnabled == this.leaveTalksEnabled
          );
        },
        equalKeywordWatching: function (e) {
          return (
            e.keywordWatchingText == this.keywordWatchingText &&
            e.keywordWatchingEmphasis == this.keywordWatchingEmphasis &&
            e.keywordWatchingSelfMessage == this.keywordWatchingSelfMessage &&
            e.keywordWatchingActionReply == this.keywordWatchingActionReply
          );
        },
        equalLanguageSettings: function (e) {
          return e.language == this.language;
        },
        equalConferenceSettings: function (e) {
          return e.conferenceIncomingSound == this.conferenceIncomingSound;
        },
        equalUserData: function (e, t) {
          var n = e.getUserData(t),
            i = this.getUserData(t);
          if (null == n) return !0;
          e = i.stampHistories.length;
          if (n.stampHistories.length != e) return !0;
          for (var r = 0, o = e; r < o; ) {
            var a = r++;
            if (!Ui.eq(n.stampHistories[a], i.stampHistories[a])) return !0;
          }
          return !1;
        },
        copyTo: function (e) {
          (this.copyNotificationSettingsTo(e),
            this.copyKeywordSettingsTo(e),
            this.copyTalkSettingsTo(e),
            this.copyUserDataContainerTo(e),
            this.copyLanguageSettingsTo(e),
            this.copyConferenceSettingsTo(e),
            this.copyDomainsSortModeSettingsTo(e),
            this.copyAppLeftPanelOpenedSettingsTo(e),
            this.copyAppLeftPanelWidthSettingsTo(e),
            this.copyScheduledMessageDefaultTimeStrSettingsTo(e));
        },
        copyNotificationSettingsTo: function (e) {
          ((e.notificationInvisibleTalkSound =
            this.notificationInvisibleTalkSound),
            (e.notificationInvisibleTalkPopup =
              this.notificationInvisibleTalkPopup),
            (e.notificationVisibleTalkSound =
              this.notificationVisibleTalkSound),
            (e.notificationVisibleTalkPopup =
              this.notificationVisibleTalkPopup),
            (e.notificationIncludingMessageContent =
              this.notificationIncludingMessageContent),
            (e.notificationPopupTime = this.notificationPopupTime));
        },
        copyKeywordSettingsTo: function (e) {
          ((e.keywordWatchingText = this.keywordWatchingText),
            (e.keywordWatchingEmphasis = this.keywordWatchingEmphasis),
            (e.keywordWatchingSelfMessage = this.keywordWatchingSelfMessage),
            (e.keywordWatchingActionReply = this.keywordWatchingActionReply));
        },
        copyTalkSettingsTo: function (e) {
          ((e.talkAutoScroll = this.talkAutoScroll),
            (e.talksSortMode = this.talksSortMode),
            (e.talkRightPanelOpened = this.talkRightPanelOpened),
            (e.talkRightPanelWidth = this.talkRightPanelWidth),
            (e.legacyNoteEnabled = this.legacyNoteEnabled),
            (e.leaveTalksEnabled = this.leaveTalksEnabled));
        },
        copyUserDataContainerTo: function (e) {
          e.userDataContainer = this.userDataContainer;
        },
        getUserData: function (t) {
          return k.find(this.userDataContainer, function (e) {
            return e.userIdString == t;
          });
        },
        copyLanguageSettingsTo: function (e) {
          e.language = this.language;
        },
        copyConferenceSettingsTo: function (e) {
          e.conferenceIncomingSound = this.conferenceIncomingSound;
        },
        copyDomainsSortModeSettingsTo: function (e) {
          e.domainsSortMode = this.domainsSortMode;
        },
        copyAppLeftPanelOpenedSettingsTo: function (e) {
          e.appLeftPanelOpened = this.appLeftPanelOpened;
        },
        copyAppLeftPanelWidthSettingsTo: function (e) {
          e.appLeftPanelWidth = this.appLeftPanelWidth;
        },
        copyScheduledMessageDefaultTimeStrSettingsTo: function (e) {
          e.scheduledMessageDefaultTimeStr =
            this.scheduledMessageDefaultTimeStr;
        },
        assureUserData: function (e) {
          var t = this.getUserData(e);
          return (
            null == t &&
              ((t = xi.fromUserIdString(e)), this.userDataContainer.push(t)),
            t
          );
        },
        isThumbnailExpansionEnabled: function (e) {
          e = this.getUserData(e);
          return null != e && e.useThumbnailExpansion;
        },
        setThumbnailExpansionEnabled: function (e, t) {
          this.assureUserData(e).useThumbnailExpansion = t;
        },
        setPrevSelectedMic: function (e, t) {
          this.assureUserData(e).setPrevSelectedMic(t);
        },
        setPrevSelectedSpeaker: function (e, t) {
          this.assureUserData(e).setPrevSelectedSpeaker(t);
        },
        setPrevSelectedCamera: function (e, t) {
          this.assureUserData(e).setPrevSelectedCamera(t);
        },
        setPrevUseBlur: function (e, t) {
          this.assureUserData(e).setPrevUseBlur(t);
        },
        setSearchHistories: function (e, t, n) {
          e = this.assureUserData(e).assureDomainIdStrToSearchHistories();
          null != Y[t] ? e.setReserved(t, n) : (e.h[t] = n);
        },
        getSearchHistories: function (e, t) {
          e = this.assureUserData(e).assureDomainIdStrToSearchHistories();
          return null != Y[t] ? e.getReserved(t) : e.h[t];
        },
        getPrevSelectedDevice: function (e) {
          e = this.getUserData(e);
          return null == e || null == e.prevSelectedDevice
            ? xi.createDefaultConfrenceDeviceSetting()
            : e.prevSelectedDevice;
        },
        getStampHistories: function (e) {
          e = this.getUserData(e);
          return null == e ? [] : e.stampHistories;
        },
        addUsedStampJson: function (e, t) {
          var e = this.assureUserData(e),
            n = e.stampHistories.filter(function (e) {
              return !Ui.eq(t, e);
            });
          (n.unshift(t), (e.stampHistories = n.slice(0, Ui.MAX_SIZE)));
        },
        isJaMode: function () {
          return null == this.language || "ja" == this.language;
        },
        __class__: _e,
      }),
      function () {}),
    Ui =
      (((t["albero.entity.UserData"] = xi).__name__ = [
        "albero",
        "entity",
        "UserData",
      ]),
      (xi.createDefaultConfrenceDeviceSetting = function () {
        return { mic: "", speaker: "", camera: "", useBlur: !1 };
      }),
      (xi.fromUserIdString = function (e) {
        var t = new xi();
        return (
          (t.userIdString = e),
          (t.stampHistories = []),
          (t.useThumbnailExpansion = !1),
          (t.prevSelectedDevice = xi.createDefaultConfrenceDeviceSetting()),
          (t.domainIdStrToSearchHistories = new B()),
          t
        );
      }),
      (xi.fromJson = function (e) {
        var t = new xi();
        return (
          null != e &&
            (null != e.userIdString && (t.userIdString = e.userIdString),
            null != e.stampHistories &&
              (t.stampHistories = e.stampHistories.map(Ui.migrate)),
            null != e.useThumbnailExpansion &&
              (t.useThumbnailExpansion = e.useThumbnailExpansion),
            null != e.prevSelectedDevice &&
              ((t.prevSelectedDevice = e.prevSelectedDevice),
              null == t.prevSelectedDevice.useBlur) &&
              (t.prevSelectedDevice.useBlur = !1),
            null != e.domainIdStrToSearchHistories) &&
            (t.domainIdStrToSearchHistories = e.domainIdStrToSearchHistories),
          t
        );
      }),
      (xi.prototype = {
        assureDomainIdStrToSearchHistories: function () {
          return (
            null == this.domainIdStrToSearchHistories &&
              (this.domainIdStrToSearchHistories = new B()),
            this.domainIdStrToSearchHistories
          );
        },
        setPrevSelectedMic: function (e) {
          this.assurePrevSelectedDevice().mic = e;
        },
        setPrevSelectedSpeaker: function (e) {
          this.assurePrevSelectedDevice().speaker = e;
        },
        setPrevSelectedCamera: function (e) {
          this.assurePrevSelectedDevice().camera = e;
        },
        setPrevUseBlur: function (e) {
          this.assurePrevSelectedDevice().useBlur = e;
        },
        assurePrevSelectedDevice: function () {
          return (
            null == this.prevSelectedDevice &&
              (this.prevSelectedDevice =
                xi.createDefaultConfrenceDeviceSetting()),
            this.prevSelectedDevice
          );
        },
        __class__: xi,
      }),
      function () {}),
    Li =
      (((t["albero.entity.StampHisory"] = Ui).__name__ = [
        "albero",
        "entity",
        "StampHisory",
      ]),
      (Ui.migrate = function (e) {
        return "string" == typeof e ? { commonStampId: e } : e;
      }),
      (Ui.eq = function (e, t) {
        return null == e || null == t
          ? e != t
          : null != e.commonStampId && null != t.commonStampId
            ? e.commonStampId == t.commonStampId
            : null != e.originalStamp &&
              null != t.originalStamp &&
              e.originalStamp.stampIdStr == t.originalStamp.stampIdStr;
      }),
      function (e) {
        ((this.allowCreateMessageStampCategories = []),
          (this.version = 0),
          null != e &&
            ((this.version = e.version),
            (this.allowCreateMessageStampCategories =
              e.allow_create_message_stamp_categories)));
      }),
    Pi =
      (((t["albero.entity.BuiltinStampSetting"] = Li).__name__ = [
        "albero",
        "entity",
        "BuiltinStampSetting",
      ]),
      (Li.prototype = {
        isNewerThan: function (e) {
          return this.version > e.version;
        },
        isStampsetEnabled: function (e) {
          return k.has(this.allowCreateMessageStampCategories, e);
        },
        isStampCategoriesEmpty: function () {
          return 0 == this.allowCreateMessageStampCategories.length;
        },
        getCategoryIndex: function (e) {
          return this.allowCreateMessageStampCategories.indexOf(e);
        },
        __class__: Li,
      }),
      ((t["albero.entity.Conference"] = he).__name__ = [
        "albero",
        "entity",
        "Conference",
      ]),
      (he.prototype = {
        isIncomingExpired: function () {
          var e = this.createdAt,
            t = 4294967296,
            n = e.high,
            e = e.low;
          return (
            n * t + (0 <= e ? e : e + t) + he.INCOMING_TIME <
            new Date().getTime()
          );
        },
        isExpired: function () {
          var e = this.expiredAt,
            t = 4294967296,
            n = e.high,
            e = e.low;
          return n * t + (0 <= e ? e : e + t) < new Date().getTime();
        },
        __class__: he,
      }),
      function () {
        ((this.conferencesMap = new B()),
          (this.talkIdConferenceIdMap = new B()));
      }),
    Bi =
      (((t["albero.entity.ConferenceStore"] = Pi).__name__ = [
        "albero",
        "entity",
        "ConferenceStore",
      ]),
      (Pi.prototype = {
        addConference: function (e) {
          var t = e.id,
            t = "_" + t.high + "_" + t.low,
            n = this.conferencesMap,
            n = (null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e), e.talkId),
            e = "_" + n.high + "_" + n.low,
            n = this.talkIdConferenceIdMap;
          null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
        },
        removeConference: function (e) {
          var t = e.id,
            t = "_" + t.high + "_" + t.low,
            e = (this.conferencesMap.remove(t), e.talkId),
            e = "_" + e.high + "_" + e.low,
            n = this.talkIdConferenceIdMap;
          (null != Y[e] ? n.getReserved(e) : n.h[e]) == t &&
            this.talkIdConferenceIdMap.remove(e);
        },
        getConferences: function () {
          return X.toValueArray(this.conferencesMap);
        },
        getConference: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.conferencesMap;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        hasConferenceInTalk: function (e) {
          var t;
          return (
            null != e &&
            ((e = "_" + e.high + "_" + e.low),
            (t = this.talkIdConferenceIdMap),
            null != Y[e] ? t.existsReserved(e) : t.h.hasOwnProperty(e))
          );
        },
        getConferenceIdInTalk: function (e) {
          var t;
          return null == e ||
            ((e = "_" + e.high + "_" + e.low),
            (t = this.talkIdConferenceIdMap),
            null == (t = null != Y[e] ? t.getReserved(e) : t.h[e]))
            ? null
            : D.makeFromIdStr(t);
        },
        getConferenceTalkIds: function () {
          return X.toKeyArray(this.talkIdConferenceIdMap).map(D.makeFromIdStr);
        },
        hasConference: function () {
          return this.talkIdConferenceIdMap.keys().hasNext();
        },
        __class__: Pi,
      }),
      function (e) {
        null != e &&
          ((this.maxFriends = e.max_friends),
          (this.maxTalks = e.max_talks),
          (this.maxTalkers = e.max_talkers),
          (this.maxMessageContentLength = e.max_message_content_length),
          (this.botExpiredVersion = e.bot_expired_version),
          (this.presenceExpiration =
            null != e.presence_expiration
              ? e.presence_expiration
              : Bi.DEFAULT_PRESENCE_EXPIRATION),
          (this.allowAttachmentType = Di.fromInt(e.allow_attachment_type)),
          (this.imageSanitizationEnabled = !!e.image_sanitization_enabled),
          (this.updateLastUsedAtInterval =
            null != e.update_last_used_at_interval
              ? e.update_last_used_at_interval
              : Bi.DEFAULT_PRESENCE_EXPIRATION));
      }),
    Hi =
      (((t["albero.entity.Configuration"] = Bi).__name__ = [
        "albero",
        "entity",
        "Configuration",
      ]),
      (Bi.prototype = { __class__: Bi }),
      function (e) {
        ((this.contractConference = m.None),
          null != e &&
            ((this.id = e.contract_id),
            (this.plan = new vr(e.plan)),
            (this.quota = new Fr(e.quota)),
            (this.solutionIds = e.solution_ids),
            null != e.conference) &&
            (this.contractConference = m.Some(new Gi(e.conference))));
      }),
    Gi =
      (((t["albero.entity.Contract"] = Hi).__name__ = [
        "albero",
        "entity",
        "Contract",
      ]),
      (Hi.prototype = {
        getMaxMessageContentLength: function () {
          return h.map(h.option(this.quota), function (e) {
            return e.maxMessageContentLength;
          });
        },
        getMaxTalkers: function () {
          return h.flatMap(h.option(this.quota), function (e) {
            return h.option(e.maxTalkers);
          });
        },
        getMaxTalks: function () {
          return h.flatMap(h.option(this.quota), function (e) {
            return h.option(e.maxTalks);
          });
        },
        isConferenceEnalbed: function () {
          return h.isDefined(this.contractConference);
        },
        isVideoEnalbed: function () {
          var e = h.map(this.contractConference, function (e) {
            return e.videoEnabled;
          });
          switch (e[1]) {
            case 0:
              return e[2];
            case 1:
              return !1;
          }
        },
        getConferenceTtl: function () {
          return h.map(this.contractConference, function (e) {
            return e.ttl;
          });
        },
        getConferenceMaxParticipants: function () {
          return h.map(this.contractConference, function (e) {
            return e.maxConferenceParticipants;
          });
        },
        isGroupConfenreceEnalbed: function () {
          var e = h.map(this.contractConference, function (e) {
            return e.groupConferenceEnabled;
          });
          switch (e[1]) {
            case 0:
              return e[2];
            case 1:
              return !1;
          }
        },
        isTargetOfPlanAd: function () {
          return null != this.quota && this.quota.withAd;
        },
        isFreePlan: function () {
          return null != this.plan && this.plan.free;
        },
        __class__: Hi,
      }),
      function (e) {
        ((this.groupConferenceEnabled = !1),
          (this.maxConferenceParticipants = Gi.DEFAULT_MAX_PARTICIPANTS),
          (this.ttl = Gi.DEFAULT_TTL),
          (this.videoEnabled = !1),
          null != e &&
            ((this.videoEnabled = e.video_enabled),
            (this.ttl = e.ttl),
            (this.maxConferenceParticipants = e.max_conference_participants),
            (this.groupConferenceEnabled = e.group_conference_enabled)));
      }),
    ji =
      (((t["albero.entity.ContractConference"] = Gi).__name__ = [
        "albero",
        "entity",
        "ContractConference",
      ]),
      (Gi.prototype = { __class__: Gi }),
      function (e) {
        if (null != e) {
          switch (
            ((this.id = e.department_id),
            (this.order = e.order),
            (this.parentId = e.parent),
            (this.nodeType =
              null == e.node
                ? zi.Normal
                : this.parseDepartmentNodeType(e.node)),
            this.nodeType[1])
          ) {
            case 0:
              this.name = e.name;
              break;
            case 1:
              this.name = ji.getLabelRoot();
              break;
            case 2:
              this.name = ji.getLabelFree();
          }
          ((this.depth = 0),
            (this.childrenIds = null),
            (this.userCount = null),
            (this.userIds = null),
            (this.userCountLoading = !1));
        }
      }),
    zi =
      (((t["albero.entity.Department"] = ji).__name__ = [
        "albero",
        "entity",
        "Department",
      ]),
      (ji.getLabelFree = function () {
        return c.localize("Department.label_free");
      }),
      (ji.getLabelRoot = function () {
        return c.localize("Department.label_root");
      }),
      (ji.prototype = {
        addChild: function (e) {
          (null == this.childrenIds && (this.childrenIds = []),
            this.childrenIds.push(e));
        },
        isRoot: function () {
          return this.nodeType == zi.Root;
        },
        isFree: function () {
          return this.nodeType == zi.Free;
        },
        parseDepartmentNodeType: function (e) {
          switch (e) {
            case 1:
              return zi.Root;
            case 2:
              return zi.Free;
            default:
              return zi.Normal;
          }
        },
        getChildIdStrList: function () {
          return null == this.childrenIds ? [] : this.childrenIds.map(D.idStr);
        },
        __class__: ji,
      }),
      ((t["albero.entity.DepartmentHelper"] = de).__name__ = [
        "albero",
        "entity",
        "DepartmentHelper",
      ]),
      (de.getName = function (e) {
        return null == e ? null : e.name;
      }),
      (t["albero.entity.DepartmentNodeType"] = {
        __ename__: ["albero", "entity", "DepartmentNodeType"],
        __constructs__: ["Normal", "Root", "Free"],
      })),
    Yi =
      ((zi.Normal = ["Normal", 0]),
      (zi.Normal.toString = a),
      ((zi.Normal.__enum__ = zi).Root = ["Root", 1]),
      (zi.Root.toString = a),
      ((zi.Root.__enum__ = zi).Free = ["Free", 2]),
      (zi.Free.toString = a),
      ((zi.Free.__enum__ = zi).__empty_constructs__ = [
        zi.Normal,
        zi.Root,
        zi.Free,
      ]),
      ((t["albero.entity.DepartmentList"] = fe).__name__ = [
        "albero",
        "entity",
        "DepartmentList",
      ]),
      (fe.prototype = {
        getDepartments: function (e) {
          return null == e
            ? []
            : e.map(function (e) {
                return new ji(e);
              });
        },
        __class__: fe,
      }),
      function () {
        ((this.departmentMap = new B()), (this.rootDepartmentIdMap = new B()));
      }),
    Wi =
      (((t["albero.entity.DepartmentStore"] = Yi).__name__ = [
        "albero",
        "entity",
        "DepartmentStore",
      ]),
      (Yi.prototype = {
        getDepartment: function (e) {
          var t;
          return null == e
            ? null
            : ((e = "_" + e.high + "_" + e.low),
              (t = this.departmentMap),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        setDepartment: function (e) {
          null != e.parentId &&
            null != (t = this.getDepartment(e.parentId)) &&
            (t.addChild(e.id), (e.depth = t.depth + 1));
          var t = e.id,
            t = "_" + t.high + "_" + t.low,
            n = this.departmentMap;
          null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e);
        },
        clearDomainDepartment: function (e) {
          var t = this,
            e = "_" + e.high + "_" + e.low,
            n = this.rootDepartmentIdMap,
            n = null != Y[e] ? n.getReserved(e) : n.h[e];
          null != n &&
            (k.iter(this.getDepartmentsDescendantOrSelf(n), function (e) {
              e = e.id;
              t.departmentMap.remove("_" + e.high + "_" + e.low);
            }),
            this.rootDepartmentIdMap.remove(e));
        },
        getRootDepartmentId: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.rootDepartmentIdMap;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        setRootDepartmentId: function (e, t) {
          var e = "_" + e.high + "_" + e.low,
            n = this.rootDepartmentIdMap;
          null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
        },
        setDepartmentUserCount: function (e) {
          var t = this.getDepartment(e.departmentId);
          null != t && ((t.userCount = e), (t.userCountLoading = !1));
        },
        setUserCountLoading: function (e, t) {
          e = this.getDepartment(e);
          null != e && (e.userCountLoading = t);
        },
        isUserCountLoadingRequired: function (e) {
          e = this.getDepartment(e);
          return null == e || (null == e.userCount && !e.userCountLoading);
        },
        clearDomainDepartmentUsers: function (e) {
          e = this.getRootDepartmentId(e);
          return (
            null != e &&
            (k.iter(this.getDepartmentsDescendantOrSelf(e), function (e) {
              ((e.userCount = null),
                (e.userIds = null),
                (e.userCountLoading = !1));
            }),
            !0)
          );
        },
        setDepartmentUsers: function (e, t) {
          e = this.getDepartment(e);
          null != e && (e.userIds = t);
        },
        isInDepartment: function (e) {
          e = this.getDepartment(e);
          return null != e && !e.isFree();
        },
        getDepartmentPath: function (e) {
          return this.getDepartmentsAncestorOrSelf(e);
        },
        getDepartmentPathWithoutRoot: function (e) {
          return this.getDepartmentPath(e).filter(function (e) {
            return !e.isRoot();
          });
        },
        getDepartmentPathSlashSeparatedIfNeeded: function (e) {
          return this.getDepartmentPathWithoutRoot(e)
            .map(de.getName)
            .filter(Ln.isNotEmpty)
            .join("/");
        },
        getDepartmentsAncestorOrSelf: function (e) {
          var t = [];
          return (this._getDepartmentsAncestorOrSelf(e, t), t);
        },
        _getDepartmentsAncestorOrSelf: function (e, t) {
          e = this.getDepartment(e);
          null != e &&
            (null != e.parentId &&
              this._getDepartmentsAncestorOrSelf(e.parentId, t),
            t.push(e));
        },
        getDepartmentsDescendantOrSelf: function (e) {
          var t = [];
          return (this._getDepartmentsDescendantOrSelf(e, t), t);
        },
        _getDepartmentsDescendantOrSelf: function (e, t) {
          var n = this,
            e = this.getDepartment(e);
          null != e &&
            (t.push(e), null != e.childrenIds) &&
            k.iter(e.childrenIds, function (e) {
              n._getDepartmentsDescendantOrSelf(e, t);
            });
        },
        __class__: Yi,
      }),
      ((t["albero.entity.DepartmentUserCount"] = me).__name__ = [
        "albero",
        "entity",
        "DepartmentUserCount",
      ]),
      (me.prototype = { __class__: me }),
      ((t["albero.entity.DepartmentUserCountList"] = pe).__name__ = [
        "albero",
        "entity",
        "DepartmentUserCountList",
      ]),
      (pe.prototype = {
        getDepartmentUserCounts: function (e) {
          return null == e
            ? []
            : e.map(function (e) {
                return new me(e);
              });
        },
        __class__: pe,
      }),
      ((t["albero.entity.DirectApp"] = ge).__name__ = [
        "albero",
        "entity",
        "DirectApp",
      ]),
      (ge.prototype = {
        getLinkFeature: function () {
          return k.find(this.features, function (e) {
            return e.isLink();
          });
        },
        getMessageExtensionFeature: function () {
          return k.find(this.features, function (e) {
            return e.isMessageMenuExtension();
          });
        },
        getLocalizedLabel: function (e, t) {
          var n = this.locales,
            n = null != Y[t] ? n.getReserved(t) : n.h[t];
          return null == n || null == (t = n.getResourceText(e)) ? "" : t;
        },
        __class__: ge,
      }),
      function (e) {
        this.map = new B();
        for (var t = 0, n = L.fields(e); t < n.length; ) {
          var i = n[t],
            r = (++t, L.field(e, i)),
            o = this.map;
          null != Y[i] ? o.setReserved(i, r) : (o.h[i] = r);
        }
      }),
    Ki =
      (((t["albero.entity.DirectAppLangResources"] = Wi).__name__ = [
        "albero",
        "entity",
        "DirectAppLangResources",
      ]),
      (Wi.prototype = {
        getResourceText: function (e) {
          var t = this.map;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        __class__: Wi,
      }),
      function (e) {
        null != e &&
          ((this.type = e.type),
          (this.label = e.label),
          (this.iconUrl = e.icon_url),
          (this.priority = e.priority),
          (this.messageCategory = this.parseMessageCategory(
            e.message_category,
          )),
          (this.multiple = !!e.multiple),
          (this.actionUrl = e.action_url));
      }),
    Vi =
      (((t["albero.entity.DirectAppFeature"] = Ki).__name__ = [
        "albero",
        "entity",
        "DirectAppFeature",
      ]),
      (Ki.prototype = {
        parseMessageCategory: function (e) {
          return 1 == e ? mr.textOrFiles : mr.unknown;
        },
        isLink: function () {
          return "link" == this.type;
        },
        isMessageMenuExtension: function () {
          return "message_menu_extension" == this.type;
        },
        __class__: Ki,
      }),
      function () {
        this.store = new B();
      }),
    qi =
      (((t["albero.entity.DirectAppStore"] = Vi).__name__ = [
        "albero",
        "entity",
        "DirectAppStore",
      ]),
      (Vi.prototype = {
        clear: function () {
          this.store = new B();
        },
        setDirectApp: function (e, t) {
          null == this.store && (this.store = new B());
          var e = "_" + e.high + "_" + e.low,
            n = this.store;
          null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
        },
        getDirectApp: function (e) {
          var t;
          return null == this.store
            ? null
            : ((e = "_" + e.high + "_" + e.low),
              (t = this.store),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        getDirectApps: function (e) {
          e = e.map(z(this, this.getDirectApp)).filter(function (e) {
            return null != e;
          });
          return new qi(e);
        },
        getAllDirectApps: function () {
          return new qi(X.toValueArray(this.store));
        },
        removeDirectApp: function (e) {
          null != this.store &&
            ((e = "_" + e.high + "_" + e.low), this.store.remove(e));
        },
        __class__: Vi,
      }),
      function (e) {
        this.values = e;
      }),
    Qi =
      (((t["albero.entity.DirectApps"] = qi).__name__ = [
        "albero",
        "entity",
        "DirectApps",
      ]),
      (qi.prototype = {
        getSortedLinkFeatures: function () {
          return Un.sortAndReturn(
            this.values
              .map(function (e) {
                return e.getLinkFeature();
              })
              .filter(function (e) {
                return null != e;
              }),
            function (e, t) {
              return e.priority - t.priority;
            },
          );
        },
        getSortedLinkFeatureTuples: function () {
          return Un.sortAndReturn(
            this.values
              .map(function (e) {
                return { app: e, feature: e.getLinkFeature() };
              })
              .filter(function (e) {
                return null != e.feature;
              }),
            function (e, t) {
              return e.feature.priority - t.feature.priority;
            },
          );
        },
        getSortedMessageExtensionFeatureTuples: function () {
          return Un.sortAndReturn(
            this.values
              .map(function (e) {
                return { app: e, feature: e.getMessageExtensionFeature() };
              })
              .filter(function (e) {
                return null != e.feature;
              }),
            function (e, t) {
              return e.feature.priority - t.feature.priority;
            },
          );
        },
        getAppUrls: function () {
          return this.values.map(function (e) {
            return e.appUrl;
          });
        },
        __class__: qi,
      }),
      function (e) {
        null == e || null == e.enabled
          ? (this.enabled = [])
          : (this.enabled = G.__cast(e.enabled, Array).map(D.fromIntOrInt64));
      }),
    Xi =
      (((t["albero.entity.DirectAppsSetting"] = Qi).__name__ = [
        "albero",
        "entity",
        "DirectAppsSetting",
      ]),
      (Qi.prototype = { __class__: Qi }),
      ((t["albero.entity.Domain"] = ye).__name__ = [
        "albero",
        "entity",
        "Domain",
      ]),
      (ye.isChangedContractTreeEnabled = function (e, t) {
        return ye.isDepartmentEnabled(e) != ye.isDepartmentEnabled(t);
      }),
      (ye.isDepartmentEnabled = function (e) {
        return (
          null != e && !!ye.isAllowListUser(e) && !!e.setting.contactTreeEnabled
        );
      }),
      (ye.isAllowListUser = function (e) {
        return null == e || e.role.allowListUsers;
      }),
      (ye.isChangedRoleType = function (e, t) {
        return (
          (null == e ? er.user : e.role.type) !=
          (null == t ? er.user : t.role.type)
        );
      }),
      (ye.isChangedMaxTalkers = function (e, t) {
        var n = h.flatMap(h.option(e), function (e) {
          return e.getMaxTalkers();
        });
        switch (n[1]) {
          case 0:
            var i = n[2];
            break;
          case 1:
            i = 0;
        }
        var r = h.flatMap(h.option(t), function (e) {
          return e.getMaxTalkers();
        });
        switch (r[1]) {
          case 0:
            var o = r[2];
            break;
          case 1:
            o = 0;
        }
        return i != o;
      }),
      (ye.isAllowReadAnnouncement = function (e) {
        return null != e && !!e.role.allowReadAnnouncements;
      }),
      (ye.isSolutionChanged = function (e, t) {
        var n = null == e ? [] : e.getSolutionIds(),
          i = null == t ? [] : t.getSolutionIds();
        if (n.length != i.length) return !0;
        for (var r = 0, o = n.length; r < o; ) {
          var a = r++;
          if (n[a] != i[a]) return !0;
        }
        return !1;
      }),
      (ye.isForwardTypeChanged = function (e, t) {
        return (
          null == e ||
          e.setting.allowForwardMessageType != t.setting.allowForwardMessageType
        );
      }),
      (ye.isSaveAllowedChanged = function (e, t, n) {
        return (
          (null == e || e.isSaveAllowed(n)) != (null == t || t.isSaveAllowed(n))
        );
      }),
      (ye.prototype = {
        getMaxMessageContentLengthOrElse: function (e) {
          var t = h.flatMap(h.option(this.contract), function (e) {
            return e.getMaxMessageContentLength();
          });
          switch (t[1]) {
            case 0:
              return t[2];
            case 1:
              return e;
          }
        },
        getMaxTalkers: function () {
          return h.flatMap(h.option(this.contract), function (e) {
            return e.getMaxTalkers();
          });
        },
        getMaxTalks: function () {
          return h.flatMap(h.option(this.contract), function (e) {
            return e.getMaxTalks();
          });
        },
        isDirectAppsOnly: function (e, t) {
          var n = this.getMaxTalks();
          switch (n[1]) {
            case 0:
              var i = n[2];
              break;
            case 1:
              i = 0;
          }
          t = t.getDomainTalksCount(e.id);
          return 0 == i && 0 == t && !ye.isAllowReadAnnouncement(e);
        },
        isSaveAllowed: function (e) {
          return e
            ? this.setting.allowSaveAttachments.desktop
            : this.setting.allowSaveAttachments.web;
        },
        isFreePlan: function () {
          return null != this.contract && this.contract.isFreePlan();
        },
        isTargetOfPlanAd: function () {
          return null != this.contract && this.contract.isTargetOfPlanAd();
        },
        isMyOwnDomain: function () {
          return this.role.isOwner();
        },
        updateStampsetSetting: function (e) {
          e.isNewerThan(this.stampsetSetting) && (this.stampsetSetting = e);
        },
        updateStampsetInfo: function (e) {
          this.stampsetSetting.updateStampsetInfo(e);
        },
        deleteStampsetInfo: function (e) {
          this.stampsetSetting.deleteStampsetInfo(e);
        },
        updateBuiltinStampSetting: function (e) {
          e.isNewerThan(this.builtinStampSetting) &&
            (this.builtinStampSetting = e);
        },
        hasAnyEnabledStampset: function () {
          return (
            !this.stampsetSetting.isStampsetIdsEmpty() ||
            !this.builtinStampSetting.isStampCategoriesEmpty()
          );
        },
        getSolutionIds: function () {
          return null == this.contract || null == this.contract.solutionIds
            ? []
            : this.contract.solutionIds;
        },
        __class__: ye,
      }),
      function (e) {
        null != e &&
          ((this.name = e.domain_name),
          (this.logoUrl = e.logo_url),
          (this.frozen = null != e.frozen && e.frozen));
      }),
    Ji =
      (((t["albero.entity.DomainInfo"] = Xi).__name__ = [
        "albero",
        "entity",
        "DomainInfo",
      ]),
      (Xi.prototype = { __class__: Xi }),
      function (e) {
        null != e &&
          ((this.allowAttachmentType = Di.fromInt(e.allow_attachment_type)),
          (this.allowSaveAttachmentsToDevice =
            e.allow_save_attachments_to_device),
          null != e.allow_save_attachments &&
            (this.allowSaveAttachments = new Zi(e.allow_save_attachments)),
          (this.allowForwardMessageType = this.allowForwardMessageTypeOf(
            e.allow_forward_message_type,
          )),
          (this.contactTreeEnabled = !!e.contact_tree_enabled),
          (this.dropboxIntegrationEnabled = !!e.dropbox_integration_enabled),
          (this.boxIntegrationEnabled = !!e.box_integration_enabled),
          (this.allowForwardNote = !!e.allow_forward_note));
      }),
    Zi =
      (((t["albero.entity.DomainSetting"] = Ji).__name__ = [
        "albero",
        "entity",
        "DomainSetting",
      ]),
      (Ji.prototype = {
        allowForwardMessageTypeOf: function (e) {
          switch (e) {
            case 0:
              return tr.ng;
            case 1:
              return tr.ok;
            case 2:
              return tr.onlyTextOrStamp;
            default:
              return tr.other;
          }
        },
        __class__: Ji,
      }),
      function (e) {
        null != e &&
          ((this.web = null == e.web || e.web),
          (this.ios = null == e.ios || e.ios),
          (this.android = null == e.android || e.android),
          (this.desktop = null == e.desktop || e.desktop));
      }),
    $i =
      (((t["albero.entity.DomainSettingAllowSaveAttachments"] = Zi).__name__ = [
        "albero",
        "entity",
        "DomainSettingAllowSaveAttachments",
      ]),
      (Zi.prototype = { __class__: Zi }),
      function (e) {
        var t;
        null != e &&
          ((this.type = this.typeOf(e.type)),
          (t = e.allow_guests),
          (this.allowGuests = t || !0),
          (this.allowCreateAttachments = e.allow_create_attachments || !0),
          (this.allowReadAttachments = e.allow_read_attachments || !0),
          (this.allowCreateAnnouncement = e.allow_create_announcement),
          (this.allowReadAnnouncements = e.allow_read_announcements),
          null == this.allowReadAnnouncements &&
            (this.allowReadAnnouncements = !0),
          (this.allowListUsers = e.allow_list_users),
          null == this.allowListUsers) &&
          (this.allowListUsers = !0);
      }),
    er =
      (((t["albero.entity.DomainRole"] = $i).__name__ = [
        "albero",
        "entity",
        "DomainRole",
      ]),
      ($i.prototype = {
        def: function (e, t) {
          return e || t;
        },
        typeOf: function (e) {
          switch (e) {
            case 10:
              return er.owner;
            case 20:
              return er.manager;
            case 30:
              return er.user;
            default:
              return er.guest;
          }
        },
        isOwner: function () {
          return x.enumEq(this.type, er.owner);
        },
        __class__: $i,
      }),
      (t["albero.entity.DomainRoleType"] = {
        __ename__: ["albero", "entity", "DomainRoleType"],
        __constructs__: ["owner", "manager", "user", "guest"],
      })),
    tr =
      ((er.owner = ["owner", 0]),
      (er.owner.toString = a),
      ((er.owner.__enum__ = er).manager = ["manager", 1]),
      (er.manager.toString = a),
      ((er.manager.__enum__ = er).user = ["user", 2]),
      (er.user.toString = a),
      ((er.user.__enum__ = er).guest = ["guest", 3]),
      (er.guest.toString = a),
      ((er.guest.__enum__ = er).__empty_constructs__ = [
        er.owner,
        er.manager,
        er.user,
        er.guest,
      ]),
      (t["albero.entity.DomainAllowForwardMessageType"] = {
        __ename__: ["albero", "entity", "DomainAllowForwardMessageType"],
        __constructs__: ["ng", "ok", "onlyTextOrStamp", "other"],
      })),
    nr =
      ((tr.ng = ["ng", 0]),
      (tr.ng.toString = a),
      ((tr.ng.__enum__ = tr).ok = ["ok", 1]),
      (tr.ok.toString = a),
      ((tr.ok.__enum__ = tr).onlyTextOrStamp = ["onlyTextOrStamp", 2]),
      (tr.onlyTextOrStamp.toString = a),
      ((tr.onlyTextOrStamp.__enum__ = tr).other = ["other", 3]),
      (tr.other.toString = a),
      ((tr.other.__enum__ = tr).__empty_constructs__ = [
        tr.ng,
        tr.ok,
        tr.onlyTextOrStamp,
        tr.other,
      ]),
      function (e) {
        null != e &&
          ((this.domainId = e.domain_id), (this.departments = e.departments));
      }),
    ir =
      (((t["albero.entity.DomainDepartmentPath"] = nr).__name__ = [
        "albero",
        "entity",
        "DomainDepartmentPath",
      ]),
      (nr.prototype = { __class__: nr }),
      ((t["albero.entity.DomainInvite"] = Se).__name__ = [
        "albero",
        "entity",
        "DomainInvite",
      ]),
      (Se.prototype = { __class__: Se }),
      function (e) {
        null != e &&
          ((this.id = e.user_id),
          (this.status = e.status),
          (this.displayName = e.display_name),
          (this.canonicalDisplayName = e.canonical_display_name),
          (this.phoneticDisplayName = e.phonetic_display_name),
          (this.canonicalPhoneticDisplayName =
            e.canonical_phonetic_display_name),
          (this.profileImageUrl = e.profile_image_url),
          (this.updatedAt = e.updated_at));
      }),
    rr =
      (((t["albero.entity.User"] = ir).__name__ = ["albero", "entity", "User"]),
      (ir.prototype = { __class__: ir }),
      ((t["albero.entity.DomainUser"] = ve).__name__ = [
        "albero",
        "entity",
        "DomainUser",
      ]),
      (ve.__super__ = ir),
      (ve.prototype = e(ir.prototype, {
        setUserIdentifier: function (e) {
          ((this.email = e.email),
            (this.subEmail = e.subEmail),
            (this.groupAlias = e.groupAlias),
            (this.signinId = e.signinId));
        },
        match: function (e, t) {
          if (this.matchNameOrPhonetic(e)) return !0;
          if (
            null != this.profileContact &&
            null != t.profileDefinition.itemDefinitions
          )
            for (var n = 0, i = this.profileContact; n < i.length; )
              for (
                var r = i[n],
                  o = (++n, 0),
                  a = t.profileDefinition.itemDefinitions;
                o < a.length;
              ) {
                var s = a[o];
                if ((++o, s.profileItemId == r.profileItemId)) {
                  if (
                    s.visible &&
                    10 == s.group &&
                    -1 < r.canonicalValue.indexOf(e)
                  )
                    return !0;
                  break;
                }
              }
          return !1;
        },
        matchNameOrPhonetic: function (e) {
          return (
            !!Ln.isEmpty(e) ||
            (Ln.isEmpty(this.canonicalDisplayName) &&
              ((this.canonicalDisplayName = xo.canonicalize(this.displayName)),
              F.v(
                "create canonical display name from %s to %s.",
                this.canonicalDisplayName,
                this.displayName,
              )),
            -1 < this.canonicalDisplayName.indexOf(e)) ||
            (null != this.canonicalPhoneticDisplayName &&
              -1 < this.canonicalPhoneticDisplayName.indexOf(e))
          );
        },
        matchNameOrPhoneticRoman: function (e) {
          return (
            !!this.matchNameOrPhonetic(e) ||
            (Ln.isEmpty(this.canonicalDisplayNameRoman) &&
              ((this.canonicalDisplayNameRoman = xo.canonicalizeForRoman(
                this.displayName,
              )),
              F.v(
                "create canonical display name(roman) from %s to %s.",
                this.canonicalDisplayNameRoman,
                this.displayName,
              )),
            -1 < this.canonicalDisplayNameRoman.indexOf(e)) ||
            (null != this.canonicalPhoneticDisplayName &&
              (Ln.isEmpty(this.canonicalPhoneticDisplayNameRoman) &&
                (this.canonicalPhoneticDisplayNameRoman =
                  xo.canonicalizeForRoman(this.canonicalPhoneticDisplayName)),
              -1 < this.canonicalPhoneticDisplayNameRoman.indexOf(e)))
          );
        },
        __class__: ve,
      })),
      ((t["albero.entity.Domains"] = Te).__name__ = [
        "albero",
        "entity",
        "Domains",
      ]),
      (Te.prototype = {
        orderByName: function () {
          var e = Un.sortAndReturn(
            this.values.map(rr.fromDomain),
            rr.compare,
          ).map(rr.toDomain);
          return new Te(e);
        },
        getOldestDomain: function () {
          return null == this.values || 0 == this.values.length
            ? m.None
            : k.fold(
                this.values,
                function (i, e) {
                  switch (
                    h.filter(e, function (e) {
                      var e = e.id,
                        t = i.id,
                        n = (e.high - t.high) | 0;
                      return (
                        0 == n && (n = w.ucompare(e.low, t.low)),
                        (e.high < 0
                          ? t.high < 0
                            ? n
                            : -1
                          : 0 <= t.high
                            ? n
                            : 1) <= 0
                      );
                    })[1]
                  ) {
                    case 0:
                      return h.filter(e, function (e) {
                        var e = e.id,
                          t = i.id,
                          n = (e.high - t.high) | 0;
                        return (
                          0 == n && (n = w.ucompare(e.low, t.low)),
                          (e.high < 0
                            ? t.high < 0
                              ? n
                              : -1
                            : 0 <= t.high
                              ? n
                              : 1) <= 0
                        );
                      });
                    case 1:
                      return m.Some(i);
                  }
                },
                m.None,
              );
        },
        getDirectAppsIds: function () {
          var n;
          return null == this.values || 0 == this.values.length
            ? []
            : ((n = new B()),
              k.iter(this.values, function (e) {
                k.iter(e.directAppsSetting.enabled, function (e) {
                  var t = "_" + e.high + "_" + e.low;
                  null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e);
                });
              }),
              X.toValueArray(n));
        },
        __class__: Te,
      }),
      function (e) {
        ((this.domain = e),
          (this.name = e.domainInfo.name),
          (this.canonicalizedName = xo.canonicalize(this.name)));
      }),
    or =
      (((t["albero.entity.DomainForSort"] = rr).__name__ = [
        "albero",
        "entity",
        "DomainForSort",
      ]),
      (rr.fromDomain = function (e) {
        return new rr(e);
      }),
      (rr.toDomain = function (e) {
        return e.domain;
      }),
      (rr.compare = function (e, t) {
        var n = Ln.compare(e.canonicalizedName, t.canonicalizedName);
        return 0 != n ? n : Ln.compare(e.name, t.name);
      }),
      (rr.prototype = { __class__: rr }),
      ((t["albero.entity.DownloadAuth"] = Ie).__name__ = [
        "albero",
        "entity",
        "DownloadAuth",
      ]),
      (Ie.prototype = { __class__: Ie }),
      function (e) {
        this.rawData = e;
      }),
    ar =
      (((t["albero.entity.DownloadAuthHeaders"] = or).__name__ = [
        "albero",
        "entity",
        "DownloadAuthHeaders",
      ]),
      (or.prototype = { __class__: or }),
      ((t["albero.entity.DownloadAuthErrorFactory"] = Ee).__name__ = [
        "albero",
        "entity",
        "DownloadAuthErrorFactory",
      ]),
      (Ee.fromApiError = function (e) {
        return Ee.isPendingFileError(e)
          ? { type: "pendingFile", reason: e.detail.reason }
          : { type: "failed" };
      }),
      (Ee.isPendingFileError = function (e) {
        return (
          null != e &&
          403 == e.code &&
          "pending file" == e.message &&
          null != e.detail &&
          null != e.detail.reason
        );
      }),
      ((t["albero.entity.DownloadAuthPendingReasonFactory"] = we).__name__ = [
        "albero",
        "entity",
        "DownloadAuthPendingReasonFactory",
      ]),
      (we.fromDownloadAuthError = function (e) {
        return null == e || "pendingFile" != e.type || null == e.reason
          ? -1
          : e.reason;
      }),
      ((t["albero.entity.FavoriteTalkEvent"] = Ne).__name__ = [
        "albero",
        "entity",
        "FavoriteTalkEvent",
      ]),
      (Ne.prototype = { __class__: Ne }),
      function (e) {
        ((this.uploadResult = sr.NONE),
          (this.preview = m.None),
          (this.thumbnailDimension = m.None),
          null != e &&
            ((this.attachmentId = e.attachment_id),
            (this.messageId = e.message_id),
            (this.talkId = e.talk_id),
            (this.id = e.file_id),
            (this.userId = e.user_id),
            (this.name = e.name),
            (this.contentType = e.content_type),
            (this.contentSize = e.content_size),
            (this.url = e.url),
            (this.thumbUrl = e.thumbnail_url),
            (this.updatedAt = e.updated_at),
            (this.file = e.file),
            (this.localThumbInfo = e.localThumbInfo),
            (this.thumbnailDimension = m.None),
            (this.deleted = null != e.deleted && e.deleted),
            (this.previewSupported =
              null != e.preview_supported && e.preview_supported)));
      }),
    sr =
      (((t["albero.entity.FileInfo"] = ar).__name__ = [
        "albero",
        "entity",
        "FileInfo",
      ]),
      (ar.fromMessageAndFile = function (e, t) {
        var n = new ar();
        return (
          (n.messageId = e.id),
          (n.talkId = e.talkId),
          (n.userId = e.userId),
          (n.updatedAt = e.createdAt),
          (n.id = t.file_id),
          (n.name = t.name),
          (n.contentType = t.content_type),
          (n.contentSize = t.content_size),
          (n.url = t.url),
          (n.thumbUrl = t.thumbnail_url),
          (n.thumbnailDimension =
            null == t.thumbnail_dimension
              ? m.None
              : m.Some(new to(t.thumbnail_dimension))),
          (n.deleted = null != t.deleted && t.deleted),
          (n.file = t.file),
          (n.localThumbInfo = t.localThumbInfo),
          (n.previewSupported = t.preview_supported),
          n
        );
      }),
      (ar.fromScheduledMessageFileContent = function (e) {
        var t = new ar();
        return (
          (t.id = e.file_id),
          (t.name = e.name),
          (t.contentType = e.content_type),
          (t.contentSize = e.content_size),
          (t.url = e.url),
          (t.thumbUrl = e.thumbnail_url),
          (t.thumbnailDimension =
            null == e.thumbnail_dimension
              ? m.None
              : m.Some(new to(e.thumbnail_dimension))),
          (t.deleted = null != e.deleted && e.deleted),
          (t.file = e.file),
          (t.localThumbInfo = e.localThumbInfo),
          (t.previewSupported = e.preview_supported),
          t
        );
      }),
      (ar.fromScheduledMessageFileContentWithFileId = function (e, t) {
        e = ar.fromScheduledMessageFileContent(e);
        return ((e.id = t), e);
      }),
      (ar.fromMessageAndFileWithFileId = function (e, t, n) {
        e = ar.fromMessageAndFile(e, t);
        return ((e.id = n), e);
      }),
      (ar.createUploadedFileInfo = function (e) {
        var t = new ar();
        return (
          (t.id = e.id),
          (t.contentType = e.content_type),
          (t.contentSize = e.content_size),
          (t.url = e.url),
          (t.name = e.name),
          t
        );
      }),
      (ar.createStagedFileInfo = function (e, t, n) {
        var i = t,
          r = new ar();
        return (
          (r.id = _.make(0, --ar.stagedFileInfosMaxId)),
          (r.talkId = e),
          null != t &&
            ((r.contentType = i.type),
            (r.contentSize = i.size),
            r.addLocalFile(t),
            (r.name = i.name)),
          null != n && r.addLocalThumbInfo(n),
          r
        );
      }),
      (ar.createStagedFileInfoFromDummyFile = function (e, t, n) {
        var i = new ar();
        return (
          (i.id = _.make(0, --ar.stagedFileInfosMaxId)),
          (i.talkId = e),
          null != t &&
            ((i.contentType = t.type),
            (i.contentSize = t.size),
            (i.file = t),
            (i.name = t.name)),
          null != n && i.addLocalThumbInfo(n),
          i
        );
      }),
      (ar.prototype = {
        isRemote: function () {
          return null == this.file;
        },
        isUploaded: function () {
          if (this.isRemote()) return !0;
          switch (this.uploadResult[1]) {
            case 0:
              return !0;
            case 1:
            case 2:
              return !1;
          }
        },
        isUploadFailed: function () {
          if (this.isRemote()) return !1;
          switch (this.uploadResult[1]) {
            case 1:
              return !0;
            case 0:
            case 2:
              return !1;
          }
        },
        setUploadResult: function (e) {
          this.uploadResult = e;
        },
        isNotUploaded: function () {
          return !this.isUploaded();
        },
        hasThumb: function () {
          return null != this.thumbUrl || this.hasLocalThumbInfo();
        },
        hasLocalThumbInfo: function () {
          return null != this.localThumbInfo;
        },
        hasLocalFile: function () {
          return null != this.file;
        },
        hasUrl: function () {
          return null != this.url;
        },
        isImage: function () {
          return I.startsWith(this.contentType, "image/");
        },
        isImageGif: function () {
          return I.startsWith(this.contentType, "image/gif");
        },
        isImageJpeg: function () {
          return I.startsWith(this.contentType, "image/jpeg");
        },
        isPdf: function () {
          return I.startsWith(this.contentType, "application/pdf");
        },
        isVideo: function () {
          return I.startsWith(this.contentType, "video/");
        },
        canDisplayAsThumb: function () {
          return !!this.hasThumb() || this.isImage();
        },
        getNoThumbText: function () {
          return null != this.name ? this.name : this.url;
        },
        getUploadedFileData: function () {
          var e;
          if (null == this.file)
            return (
              (e = {
                url: this.url,
                content_type: this.contentType,
                content_size: this.contentSize,
                name: this.name,
                file_id: this.id,
              }),
              null != this.thumbUrl && (e.thumbnail_url = this.thumbUrl),
              h.isDefined(this.thumbnailDimension) &&
                (e.thumbnail_dimension = h
                  .get(this.thumbnailDimension)
                  .toObject()),
              e
            );
          var t = this.uploadResult;
          switch (t[1]) {
            case 0:
              return t[2];
            case 1:
            case 2:
              return (F.e("file not uploaded."), null);
          }
        },
        addLocalFile: function (e) {
          this.file = e;
        },
        addLocalThumbInfo: function (e) {
          this.localThumbInfo = e;
        },
        getUploadableFilePack: function () {
          return {
            file: this.file,
            thumb: this.localThumbInfo,
            uploadedFile: this.isUploaded() ? this.getUploadedFileData() : null,
          };
        },
        __class__: ar,
      }),
      (t["albero.entity.UploadResult"] = {
        __ename__: ["albero", "entity", "UploadResult"],
        __constructs__: ["SUCCESS", "FAILUER", "NONE"],
      })),
    lr =
      ((sr.SUCCESS = function (e) {
        e = ["SUCCESS", 0, e];
        return ((e.__enum__ = sr), (e.toString = a), e);
      }),
      (sr.FAILUER = ["FAILUER", 1]),
      (sr.FAILUER.toString = a),
      ((sr.FAILUER.__enum__ = sr).NONE = ["NONE", 2]),
      (sr.NONE.toString = a),
      ((sr.NONE.__enum__ = sr).__empty_constructs__ = [sr.FAILUER, sr.NONE]),
      ((t["albero.entity.FileInfoDeletion"] = Ae).__name__ = [
        "albero",
        "entity",
        "FileInfoDeletion",
      ]),
      (Ae.prototype = { __class__: Ae }),
      function () {
        ((this.fileInfoMap = new B()),
          (this.fileIdStrsMap = new B()),
          (this.messageFileIdStrsMap = new B()),
          (this.noteFileIdStrsMap = new B()),
          (this.filePreviewMap = new B()),
          (this.scheduledMessageFileIdStrsMap = new B()));
      }),
    ur =
      (((t["albero.entity.FileInfoStore"] = lr).__name__ = [
        "albero",
        "entity",
        "FileInfoStore",
      ]),
      (lr.prototype = {
        ensureStagedFileInfos: function (e) {
          return this.ensureStagedFileIdStrs(e)
            .map(D.makeFromIdStr)
            .filter(function (e) {
              return null != e;
            })
            .map(z(this, this.getFileInfo));
        },
        moveStagedFile: function (e, t, n) {
          e = this.ensureStagedFileIdStrs(e);
          null != e && ((t = e[t]), U.remove(e, t), e.splice(n, 0, t));
        },
        ensureStagedFileIdStrs: function (e) {
          var t,
            e = e.toString(),
            n = this.fileIdStrsMap,
            n = null != Y[e] ? n.getReserved(e) : n.h[e];
          return (
            null == n &&
              ((n = []),
              (t = this.fileIdStrsMap),
              null != Y[e] ? t.setReserved(e, n) : (t.h[e] = n)),
            n
          );
        },
        addStagedFileInfo: function (e, t, n, i) {
          ((e = ar.createStagedFileInfo(e, n, i)),
            (n = e.id),
            (i = "_" + n.high + "_" + n.low),
            (n = this.fileInfoMap),
            null != Y[i] ? n.setReserved(i, e) : (n.h[i] = e),
            (n = e.id),
            (i = "_" + n.high + "_" + n.low));
          return (this.ensureStagedFileIdStrs(t).push(i), e);
        },
        replaceStagedFileInfo: function (e, t, n, i) {
          n = ar.createStagedFileInfo(e.talkId, n, i);
          return (this.replaceStagedFileInfoFromFileId(e.id, t, n), n);
        },
        replaceStagedFileInfoFromFileId: function (e, t, n) {
          var i = n.id,
            i = "_" + i.high + "_" + i.low,
            r = this.fileInfoMap,
            r =
              (null != Y[i] ? r.setReserved(i, n) : (r.h[i] = n),
              this.ensureStagedFileIdStrs(t)),
            i = r.indexOf("_" + e.high + "_" + e.low);
          -1 != i &&
            (r.splice(i, 1),
            (n = n.id),
            r.splice(i, 0, "_" + n.high + "_" + n.low),
            this.removeStagedFileInfo(t, e));
        },
        restoreStagedFile: function (e, t) {
          var n = t.id,
            n = "_" + n.high + "_" + n.low,
            i = this.fileInfoMap,
            i = (null != Y[n] ? i.setReserved(n, t) : (i.h[n] = t), t.id),
            n = "_" + i.high + "_" + i.low;
          this.ensureStagedFileIdStrs(e).push(n);
        },
        removeStagedFileInfo: function (e, t) {
          ((t = "_" + t.high + "_" + t.low),
            U.remove(this.ensureStagedFileIdStrs(e), t),
            (e = this.fileInfoMap),
            (e = null != Y[t] ? e.getReserved(t) : e.h[t]));
          null == e || e.isRemote() || this.fileInfoMap.remove(t);
        },
        removeStagedFileInfos: function (e) {
          for (var t = 0, n = this.ensureStagedFileIdStrs(e); t < n.length; ) {
            var i = n[t];
            (++t, this.fileInfoMap.remove(i));
          }
          this.fileIdStrsMap.remove(e.toString());
        },
        ensureTalkFileInfos: function (e) {
          return this.ensureTalkFileIds(e).map(z(this, this.fileInfoGetter));
        },
        ensureTalkFileIds: function (e) {
          var t,
            e = "_" + e.high + "_" + e.low,
            n = this.fileIdStrsMap,
            n = null != Y[e] ? n.getReserved(e) : n.h[e];
          return (
            null == n &&
              ((n = []),
              (t = this.fileIdStrsMap),
              null != Y[e] ? t.setReserved(e, n) : (t.h[e] = n)),
            n
          );
        },
        getTalkFileInfo: function (e, t, n) {
          return this.getFileInfoForMessages(n, t);
        },
        setTalkFileInfo: function (e, t) {
          var n = e.id,
            n = "_" + n.high + "_" + n.low,
            i = e.messageId,
            i = "_" + i.high + "_" + i.low,
            e =
              (this.messageFileInfoSetter(e), this.ensureTalkFileIds(e.talkId)),
            n = n + i;
          -1 == e.indexOf(n) && (t ? e.push(n) : e.unshift(n));
        },
        deleteFile: function (e) {
          e = this.getTalkFileInfo(e.talkId, e.messageId, e.fileId);
          null != e && (e.deleted = !0);
        },
        getFileInfo: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.fileInfoMap;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        getMessageFileInfos: function (t) {
          var n = this,
            e = h.map(
              h.option(
                this.messageFileIdStrsMap.get("_" + t.high + "_" + t.low),
              ),
              function (e) {
                return e.map(function (e) {
                  return n.messageFileInfoGetter(e, "_" + t.high + "_" + t.low);
                });
              },
            );
          switch (e[1]) {
            case 0:
              return e[2];
            case 1:
              return [];
          }
        },
        setMessageFileInfos: function (e, t) {
          for (var n = 0; n < t.length; ) {
            var i = t[n];
            (++n, this.messageFileInfoSetter(i));
          }
          var r = this.messageFileIdStrsMap,
            e = "_" + e.high + "_" + e.low,
            o = t.map(z(this, this.fileIdStrGetter));
          null != Y[e] ? r.setReserved(e, o) : (r.h[e] = o);
        },
        removeMessageFileInfos: function (e) {
          var t = "_" + e.high + "_" + e.low,
            n = this.messageFileIdStrsMap,
            i = null != Y[t] ? n.getReserved(t) : n.h[t];
          if (null != i)
            for (var r = 0; r < i.length; ) {
              var o = i[r];
              (++r, this.messageFileInfoRemover(o, t));
            }
          this.messageFileIdStrsMap.remove("_" + e.high + "_" + e.low);
        },
        isAllFileDeletedOnMessage: function (e) {
          return !k.exists(this.getMessageFileInfos(e), function (e) {
            return !e.deleted;
          });
        },
        getFileInfoForMessages: function (e, t) {
          return this.messageFileInfoGetter(
            "_" + e.high + "_" + e.low,
            "_" + t.high + "_" + t.low,
          );
        },
        getScheduledMessageFileInfos: function (e) {
          var t = this,
            n = "_" + e.high + "_" + e.low,
            i = h.map(
              h.option(this.scheduledMessageFileIdStrsMap.get(n)),
              function (e) {
                return e.map(function (e) {
                  return t.fileInfoMap.get(t.getScheduledMessgaeFileKey(e, n));
                });
              },
            );
          switch (i[1]) {
            case 0:
              return i[2];
            case 1:
              return [];
          }
        },
        setScheduledMessageFileInfos: function (e, t) {
          for (var n = "_" + e.high + "_" + e.low, i = 0; i < t.length; ) {
            var r = t[i],
              o = (++i, this.fileInfoMap),
              a = r.id,
              a = this.getScheduledMessgaeFileKey(
                "_" + a.high + "_" + a.low,
                n,
              );
            null != Y[a] ? o.setReserved(a, r) : (o.h[a] = r);
          }
          var e = this.scheduledMessageFileIdStrsMap,
            s = t.map(z(this, this.fileIdStrGetter));
          null != Y[n] ? e.setReserved(n, s) : (e.h[n] = s);
        },
        removeScheduledMessageFileInfos: function (e) {
          var t = "_" + e.high + "_" + e.low,
            e = this.scheduledMessageFileIdStrsMap,
            n = null != Y[t] ? e.getReserved(t) : e.h[t];
          if (null != n)
            for (var i = 0; i < n.length; ) {
              var r = n[i];
              (++i,
                this.fileInfoMap.remove(this.getScheduledMessgaeFileKey(r, t)));
            }
          this.scheduledMessageFileIdStrsMap.remove(t);
        },
        getFileInfoForSheduledMessages: function (e, t) {
          var n = this.fileInfoMap,
            e = this.getScheduledMessgaeFileKey(
              "_" + e.high + "_" + e.low,
              "_" + t.high + "_" + t.low,
            ),
            t = n;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        getNoteFileInfos: function (e) {
          var t = this,
            e = e.toString(),
            n = h.map(h.option(this.noteFileIdStrsMap.get(e)), function (e) {
              return e.map(z(t, t.fileInfoGetter));
            });
          switch (n[1]) {
            case 0:
              return n[2];
            case 1:
              return [];
          }
        },
        setNoteFileInfos: function (e, t) {
          var e = e.toString(),
            n = this.noteFileIdStrsMap,
            i = null != Y[e] ? n.getReserved(e) : n.h[e];
          if (null != i)
            for (var r = 0; r < i.length; ) {
              var o = i[r];
              (++r, this.fileInfoMap.remove(o));
            }
          if ((this.noteFileIdStrsMap.remove(e), null != t && 0 != t.length)) {
            for (var a = 0; a < t.length; ) {
              var s = t[a],
                l = (++a, s.id),
                l = "_" + l.high + "_" + l.low,
                u = this.fileInfoMap;
              null != Y[l] ? u.setReserved(l, s) : (u.h[l] = s);
            }
            var n = this.noteFileIdStrsMap,
              c = t.map(z(this, this.fileIdStrGetter));
            null != Y[e] ? n.setReserved(e, c) : (n.h[e] = c);
          }
        },
        getNoteFileInfosForEdit: function (e) {
          var t = this,
            e = e.toString() + "edit",
            n = h.map(h.option(this.noteFileIdStrsMap.get(e)), function (e) {
              return e.map(z(t, t.fileInfoGetter));
            });
          switch (n[1]) {
            case 0:
              return n[2];
            case 1:
              return [];
          }
        },
        setNoteFileInfosForEdit: function (e, t) {
          var e = e.toString() + "edit",
            n = this.noteFileIdStrsMap,
            i = null != Y[e] ? n.getReserved(e) : n.h[e];
          if (null != i)
            for (var r = 0; r < i.length; ) {
              var o = i[r];
              (++r, this.fileInfoMap.remove(o));
            }
          if ((this.noteFileIdStrsMap.remove(e), null != t && 0 != t.length)) {
            for (var a = 0; a < t.length; ) {
              var s = t[a],
                l = (++a, s.id),
                l = "_" + l.high + "_" + l.low,
                u = this.fileInfoMap;
              null != Y[l] ? u.setReserved(l, s) : (u.h[l] = s);
            }
            var n = this.noteFileIdStrsMap,
              c = t.map(z(this, this.fileIdStrGetter));
            null != Y[e] ? n.setReserved(e, c) : (n.h[e] = c);
          }
        },
        getFilePreview: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.filePreviewMap;
          return h.option(null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        setFilePreview: function (e, t) {
          var e = "_" + e.high + "_" + e.low,
            n = this.filePreviewMap;
          null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
        },
        fileIdStrGetter: function (e) {
          e = e.id;
          return "_" + e.high + "_" + e.low;
        },
        fileInfoGetter: function (e) {
          var t = this.fileInfoMap;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        messageFileInfoSetter: function (e) {
          var t = e.id,
            n = e.messageId,
            t = "_" + t.high + "_" + t.low + "_" + n.high + "_" + n.low,
            n = this.fileInfoMap;
          null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e);
        },
        messageFileInfoGetter: function (e, t) {
          var n = this.fileInfoMap,
            e = e + t;
          return null != Y[e] ? n.getReserved(e) : n.h[e];
        },
        messageFileInfoRemover: function (e, t) {
          this.fileInfoMap.remove(e + t);
        },
        getScheduledMessgaeFileKey: function (e, t) {
          return e + t + "scheduled";
        },
        __class__: lr,
      }),
      ((t["albero.entity.FlowNotificationBadge"] = be).__name__ = [
        "albero",
        "entity",
        "FlowNotificationBadge",
      ]),
      (be.prototype = { __class__: be }),
      function () {}),
    cr =
      (((t["albero.entity.FlowNotificationBadgeStore"] = ur).__name__ = [
        "albero",
        "entity",
        "FlowNotificationBadgeStore",
      ]),
      (ur.prototype = {
        clear: function () {
          this.badgeMap = new B();
        },
        setBadge: function (e, t) {
          null == this.badgeMap && (this.badgeMap = new B());
          var e = "_" + e.high + "_" + e.low,
            n = this.badgeMap;
          null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
        },
        getBadge: function (e) {
          var t;
          return null == this.badgeMap
            ? null
            : ((e = "_" + e.high + "_" + e.low),
              (t = this.badgeMap),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        removeBadge: function (e) {
          null != this.badgeMap &&
            ((e = "_" + e.high + "_" + e.low), this.badgeMap.remove(e));
        },
        isBadgeShownForOneDomain: function (e) {
          return null != e && this.isBadgeShown("_" + e.high + "_" + e.low);
        },
        isBadgeShown: function (e) {
          var t;
          return (
            null != this.badgeMap &&
            ((t = this.badgeMap),
            null != (t = null != Y[e] ? t.getReserved(e) : t.h[e])) &&
            t.show
          );
        },
        __class__: ur,
      }),
      function (e) {
        null != e &&
          ((this.talkId = e.talk_id), null != e.messages) &&
          (this.messages = G.__cast(e.messages, Array).map(function (e) {
            return new _r(e);
          }));
      }),
    _r =
      (((t["albero.entity.HitContext"] = cr).__name__ = [
        "albero",
        "entity",
        "HitContext",
      ]),
      (cr.prototype = {
        findHitMessage: function (t) {
          return k.find(this.messages, function (e) {
            e = e.id;
            return null != e && null != t && e.high == t.high && e.low == t.low;
          });
        },
        __class__: cr,
      }),
      function (e) {
        null != e &&
          ((this.id = e.message_id),
          (this.talkId = e.talk_id),
          (this.userId = e.user_id),
          (this.type = O.typeOf(e.type)),
          (this.content = e.content),
          (this.createdAt = e.created_at),
          (this.hit = e.hit));
      }),
    hr =
      (((t["albero.entity.HitMessage"] = _r).__name__ = [
        "albero",
        "entity",
        "HitMessage",
      ]),
      (_r.prototype = { __class__: _r }),
      ((t["albero.entity.IdSpec"] = Me).__name__ = [
        "albero",
        "entity",
        "IdSpec",
      ]),
      (Me.timestamp = function (e) {
        var t = Me.TIMESTAMP_SHIFT;
        return 0 == (t &= 63)
          ? new E(e.high, e.low)
          : t < 32
            ? new E(e.high >> t, (e.high << (32 - t)) | (e.low >>> t))
            : new E(e.high >> 31, e.high >> (t - 32));
      }),
      (Me.subTimestampPart = function (e, t) {
        var e = Me.timestamp(e),
          t = Me.timestamp(t),
          n = (e.high - t.high) | 0,
          i = (e.low - t.low) | 0;
        return (w.ucompare(e.low, t.low) < 0 && (n--, (n |= 0)), new E(n, i));
      }),
      (Me.createForTest = function (e) {
        var t = Me.TIMESTAMP_SHIFT;
        return 0 == (t &= 63)
          ? new E(e.high, e.low)
          : t < 32
            ? new E((e.high << t) | (e.low >>> (32 - t)), e.low << t)
            : new E(e.low << (t - 32), 0);
      }),
      ((t["albero.entity.IdValue"] = Ce).__name__ = [
        "albero",
        "entity",
        "IdValue",
      ]),
      (Ce.prototype = { __class__: Ce }),
      ((t["albero.entity.JoinConference"] = ke).__name__ = [
        "albero",
        "entity",
        "JoinConference",
      ]),
      (ke.prototype = {
        parseMode: function (e) {
          return "sfu" == e
            ? hr.sfu
            : "fullMesh" == e
              ? hr.fullMesh
              : hr.unknown;
        },
        isSfu: function (e) {
          switch (this.mode[1]) {
            case 0:
              return !0;
            case 1:
              return !1;
            case 2:
              return e.type == Xr.GroupTalk;
          }
        },
        getEllaspedSecOnJoin: function () {
          var e = D.fromIntOrInt64(this.credential.timestamp),
            t = D.fromIntOrInt64(this.timestamp),
            n = (t.high - e.high) | 0,
            i = (t.low - e.low) | 0,
            t = (w.ucompare(t.low, e.low) < 0 && (n--, (n |= 0)), new E(n, i));
          return t.high < 0
            ? 0
            : 0 < t.high || this.credential.ttl < t.low
              ? this.credential.ttl
              : t.low;
        },
        __class__: ke,
      }),
      (t["albero.entity.JoinConferenceMode"] = {
        __ename__: ["albero", "entity", "JoinConferenceMode"],
        __constructs__: ["sfu", "fullMesh", "unknown"],
      })),
    R =
      ((hr.sfu = ["sfu", 0]),
      (hr.sfu.toString = a),
      ((hr.sfu.__enum__ = hr).fullMesh = ["fullMesh", 1]),
      (hr.fullMesh.toString = a),
      ((hr.fullMesh.__enum__ = hr).unknown = ["unknown", 2]),
      (hr.unknown.toString = a),
      ((hr.unknown.__enum__ = hr).__empty_constructs__ = [
        hr.sfu,
        hr.fullMesh,
        hr.unknown,
      ]),
      (t["albero.entity.LocalErrorType"] = {
        __ename__: ["albero", "entity", "LocalErrorType"],
        __constructs__: [
          "FailedUser",
          "RequestTimeout",
          "TalksCountOvered",
          "TalkersCountOvered",
          "SendingMessageConflicted",
          "DeletingMessageConflicted",
          "UploadSizeOvered",
          "StorageSizeOvered",
          "GeneralFileError",
          "UploadForbiddenForNote",
          "UploadSizeOveredForNote",
          "StorageSizeOveredForNote",
          "GeneralFileErrorForNote",
          "AddingSomeUsersFailure",
          "NoteForwardingForbidden",
          "ScheduleMessageFailed",
          "SendingScheduledMessageFailed",
          "CannotEditFile",
          "CannotPlayFile",
          "InvalidFileName",
          "MaxBatchDownloadSize",
          "BanForwardNote",
          "DeletedNote",
          "MaxAttachmentFileNum",
          "MaxActionTextLength",
          "CannotReadFileAsImage",
          "FailedToLoadImage",
          "ForbiddenFileExists",
          "FailedToAddReaction",
          "FailedToChangeReaction",
          "FailedToCancelReaction",
          "SearchTooManyTalks",
          "NoteContentUncompressError",
          "NoteTextTooLong",
          "NoteTextTooComplex",
        ],
      })),
    dr =
      ((R.FailedUser = ["FailedUser", 0]),
      (R.FailedUser.toString = a),
      ((R.FailedUser.__enum__ = R).RequestTimeout = ["RequestTimeout", 1]),
      (R.RequestTimeout.toString = a),
      ((R.RequestTimeout.__enum__ = R).TalksCountOvered = function (e) {
        e = ["TalksCountOvered", 2, e];
        return ((e.__enum__ = R), (e.toString = a), e);
      }),
      (R.TalkersCountOvered = function (e) {
        e = ["TalkersCountOvered", 3, e];
        return ((e.__enum__ = R), (e.toString = a), e);
      }),
      (R.SendingMessageConflicted = ["SendingMessageConflicted", 4]),
      (R.SendingMessageConflicted.toString = a),
      ((R.SendingMessageConflicted.__enum__ = R).DeletingMessageConflicted = [
        "DeletingMessageConflicted",
        5,
      ]),
      (R.DeletingMessageConflicted.toString = a),
      ((R.DeletingMessageConflicted.__enum__ = R).UploadSizeOvered = function (
        e,
      ) {
        e = ["UploadSizeOvered", 6, e];
        return ((e.__enum__ = R), (e.toString = a), e);
      }),
      (R.StorageSizeOvered = function (e) {
        e = ["StorageSizeOvered", 7, e];
        return ((e.__enum__ = R), (e.toString = a), e);
      }),
      (R.GeneralFileError = ["GeneralFileError", 8]),
      (R.GeneralFileError.toString = a),
      ((R.GeneralFileError.__enum__ = R).UploadForbiddenForNote = function (e) {
        e = ["UploadForbiddenForNote", 9, e];
        return ((e.__enum__ = R), (e.toString = a), e);
      }),
      (R.UploadSizeOveredForNote = function (e, t) {
        e = ["UploadSizeOveredForNote", 10, e, t];
        return ((e.__enum__ = R), (e.toString = a), e);
      }),
      (R.StorageSizeOveredForNote = function (e, t) {
        e = ["StorageSizeOveredForNote", 11, e, t];
        return ((e.__enum__ = R), (e.toString = a), e);
      }),
      (R.GeneralFileErrorForNote = function (e) {
        e = ["GeneralFileErrorForNote", 12, e];
        return ((e.__enum__ = R), (e.toString = a), e);
      }),
      (R.AddingSomeUsersFailure = ["AddingSomeUsersFailure", 13]),
      (R.AddingSomeUsersFailure.toString = a),
      ((R.AddingSomeUsersFailure.__enum__ = R).NoteForwardingForbidden = [
        "NoteForwardingForbidden",
        14,
      ]),
      (R.NoteForwardingForbidden.toString = a),
      ((R.NoteForwardingForbidden.__enum__ = R).ScheduleMessageFailed =
        function (e) {
          e = ["ScheduleMessageFailed", 15, e];
          return ((e.__enum__ = R), (e.toString = a), e);
        }),
      (R.SendingScheduledMessageFailed = ["SendingScheduledMessageFailed", 16]),
      (R.SendingScheduledMessageFailed.toString = a),
      ((R.SendingScheduledMessageFailed.__enum__ = R).CannotEditFile = [
        "CannotEditFile",
        17,
      ]),
      (R.CannotEditFile.toString = a),
      ((R.CannotEditFile.__enum__ = R).CannotPlayFile = ["CannotPlayFile", 18]),
      (R.CannotPlayFile.toString = a),
      ((R.CannotPlayFile.__enum__ = R).InvalidFileName = [
        "InvalidFileName",
        19,
      ]),
      (R.InvalidFileName.toString = a),
      ((R.InvalidFileName.__enum__ = R).MaxBatchDownloadSize = [
        "MaxBatchDownloadSize",
        20,
      ]),
      (R.MaxBatchDownloadSize.toString = a),
      ((R.MaxBatchDownloadSize.__enum__ = R).BanForwardNote = [
        "BanForwardNote",
        21,
      ]),
      (R.BanForwardNote.toString = a),
      ((R.BanForwardNote.__enum__ = R).DeletedNote = ["DeletedNote", 22]),
      (R.DeletedNote.toString = a),
      ((R.DeletedNote.__enum__ = R).MaxAttachmentFileNum = function (e) {
        e = ["MaxAttachmentFileNum", 23, e];
        return ((e.__enum__ = R), (e.toString = a), e);
      }),
      (R.MaxActionTextLength = function (e) {
        e = ["MaxActionTextLength", 24, e];
        return ((e.__enum__ = R), (e.toString = a), e);
      }),
      (R.CannotReadFileAsImage = ["CannotReadFileAsImage", 25]),
      (R.CannotReadFileAsImage.toString = a),
      ((R.CannotReadFileAsImage.__enum__ = R).FailedToLoadImage = [
        "FailedToLoadImage",
        26,
      ]),
      (R.FailedToLoadImage.toString = a),
      ((R.FailedToLoadImage.__enum__ = R).ForbiddenFileExists = function (
        e,
        t,
      ) {
        e = ["ForbiddenFileExists", 27, e, t];
        return ((e.__enum__ = R), (e.toString = a), e);
      }),
      (R.FailedToAddReaction = ["FailedToAddReaction", 28]),
      (R.FailedToAddReaction.toString = a),
      ((R.FailedToAddReaction.__enum__ = R).FailedToChangeReaction = [
        "FailedToChangeReaction",
        29,
      ]),
      (R.FailedToChangeReaction.toString = a),
      ((R.FailedToChangeReaction.__enum__ = R).FailedToCancelReaction = [
        "FailedToCancelReaction",
        30,
      ]),
      (R.FailedToCancelReaction.toString = a),
      ((R.FailedToCancelReaction.__enum__ = R).SearchTooManyTalks = [
        "SearchTooManyTalks",
        31,
      ]),
      (R.SearchTooManyTalks.toString = a),
      ((R.SearchTooManyTalks.__enum__ = R).NoteContentUncompressError =
        function (e) {
          e = ["NoteContentUncompressError", 32, e];
          return ((e.__enum__ = R), (e.toString = a), e);
        }),
      (R.NoteTextTooLong = ["NoteTextTooLong", 33]),
      (R.NoteTextTooLong.toString = a),
      ((R.NoteTextTooLong.__enum__ = R).NoteTextTooComplex = [
        "NoteTextTooComplex",
        34,
      ]),
      (R.NoteTextTooComplex.toString = a),
      ((R.NoteTextTooComplex.__enum__ = R).__empty_constructs__ = [
        R.FailedUser,
        R.RequestTimeout,
        R.SendingMessageConflicted,
        R.DeletingMessageConflicted,
        R.GeneralFileError,
        R.AddingSomeUsersFailure,
        R.NoteForwardingForbidden,
        R.SendingScheduledMessageFailed,
        R.CannotEditFile,
        R.CannotPlayFile,
        R.InvalidFileName,
        R.MaxBatchDownloadSize,
        R.BanForwardNote,
        R.DeletedNote,
        R.CannotReadFileAsImage,
        R.FailedToLoadImage,
        R.FailedToAddReaction,
        R.FailedToChangeReaction,
        R.FailedToCancelReaction,
        R.SearchTooManyTalks,
        R.NoteTextTooLong,
        R.NoteTextTooComplex,
      ]),
      (t["albero.entity.ScheduleMessageFailedType"] = {
        __ename__: ["albero", "entity", "ScheduleMessageFailedType"],
        __constructs__: [
          "InvalidCreateAt",
          "Conflict",
          "LimitReached",
          "GeneralFileError",
          "UploadSizeOvered",
          "StorageSizeOvered",
          "Unknown",
        ],
      })),
    fr =
      ((dr.InvalidCreateAt = ["InvalidCreateAt", 0]),
      (dr.InvalidCreateAt.toString = a),
      ((dr.InvalidCreateAt.__enum__ = dr).Conflict = ["Conflict", 1]),
      (dr.Conflict.toString = a),
      ((dr.Conflict.__enum__ = dr).LimitReached = ["LimitReached", 2]),
      (dr.LimitReached.toString = a),
      ((dr.LimitReached.__enum__ = dr).GeneralFileError = [
        "GeneralFileError",
        3,
      ]),
      (dr.GeneralFileError.toString = a),
      ((dr.GeneralFileError.__enum__ = dr).UploadSizeOvered = function (e) {
        e = ["UploadSizeOvered", 4, e];
        return ((e.__enum__ = dr), (e.toString = a), e);
      }),
      (dr.StorageSizeOvered = function (e) {
        e = ["StorageSizeOvered", 5, e];
        return ((e.__enum__ = dr), (e.toString = a), e);
      }),
      (dr.Unknown = ["Unknown", 6]),
      (dr.Unknown.toString = a),
      ((dr.Unknown.__enum__ = dr).__empty_constructs__ = [
        dr.InvalidCreateAt,
        dr.Conflict,
        dr.LimitReached,
        dr.GeneralFileError,
        dr.Unknown,
      ]),
      ((t["albero.entity.LocalErrorTypeMixin"] = De).__name__ = [
        "albero",
        "entity",
        "LocalErrorTypeMixin",
      ]),
      (De.toTitleText = function (e) {
        switch (e[1]) {
          case 0:
            return c.localize("LocalErrorType.failed_user");
          case 1:
            return c.localize("LocalErrorType.request_timeout");
          case 2:
            var t = e[2];
            return 0 == t
              ? c.localize("LocalErrorType.talks_forbidden")
              : c.localize("LocalErrorType.number_of_talks", { limitMax: t });
          case 3:
            return c.localize("LocalErrorType.number_of_talkers", {
              limitMax: e[2],
            });
          case 4:
            return c.localize("LocalErrorType.retry_sending");
          case 5:
            return c.localize("LocalErrorType.retry_deletion");
          case 6:
            t = V.createByteSizeStringWithUnit(e[2]);
            return c.localize("LocalErrorType.file_size", { sizeString: t });
          case 7:
            t = V.createByteSizeStringWithUnit(e[2]);
            return c.localize("LocalErrorType.storage_size", { sizeString: t });
          case 9:
            return c.localize("LocalErrorType.file_forbidden_for_note");
          case 10:
            t = V.createByteSizeStringWithUnit(e[2]);
            return c.localize("LocalErrorType.file_size", { sizeString: t });
          case 11:
            t = V.createByteSizeStringWithUnit(e[2]);
            return c.localize("LocalErrorType.storage_size", { sizeString: t });
          case 8:
          case 12:
            return c.localize("LocalErrorType.file");
          case 13:
            return c.localize("LocalErrorType.adding_users");
          case 14:
            return c.localize("LocalErrorType.note_forwarding_forbidden");
          case 15:
            return c.localize("LocalErrorType.schedule_message_failed");
          case 16:
            return c.localize(
              "LocalErrorType.sending_scheduled_message_failed",
            );
          case 17:
            return c.localize("LocalErrorType.cannot_edit_file");
          case 18:
            return c.localize("LocalErrorType.cannot_play_file");
          case 19:
            return c.localize("LocalErrorType.invalid_file_name");
          case 20:
            return c.localize("LocalErrorType.max_batch_download_size");
          case 21:
            return c.localize("LocalErrorType.ban_forward_note");
          case 22:
            return c.localize("LocalErrorType.deleted_note");
          case 23:
            return c.localize("LocalErrorType.max_attachment_file_num", {
              maxAttachmentFileNum: e[2],
            });
          case 24:
            return c.localize("LocalErrorType.max_action_text_length", {
              maxActionTextLength: e[2],
            });
          case 25:
            return c.localize("LocalErrorType.cannot_read_file_as_image");
          case 26:
            return c.localize("LocalErrorType.failed_to_load_image");
          case 27:
            var t = e[3],
              n = e[2],
              i = c.localize("LocalErrorType.period");
            return Di.createAttachmentErrorMessage(n, t, i);
          case 28:
            return c.localize("LocalErrorType.failed_to_add_reaction");
          case 29:
            return c.localize("LocalErrorType.failed_to_change_reaction");
          case 30:
            return c.localize("LocalErrorType.failed_to_cancel_reaction");
          case 31:
            return c.localize("LocalErrorType.search_too_many_talks");
          case 32:
            return c.localize("LocalErrorType.note_content_uncompress_error");
          case 33:
            return c.localize("LocalErrorType.note_text_too_long");
          case 34:
            return c.localize("LocalErrorType.note_text_too_complex");
        }
      }),
      (De.toBodyText = function (e) {
        switch (e[1]) {
          case 15:
            return De.getScheduledMessageFailedBodyText(e[2]);
          case 16:
            return c.localize(
              "LocalErrorType.sending_scheduled_message_failed_by_unknown",
            );
          default:
            return null;
        }
      }),
      (De.getScheduledMessageFailedBodyText = function (e) {
        switch (e[1]) {
          case 0:
            return c.localize(
              "LocalErrorType.schedule_message_failed_by_invalid_create_at",
            );
          case 1:
            return c.localize(
              "LocalErrorType.schedule_message_failed_by_conflict",
            );
          case 2:
            return c.localize(
              "LocalErrorType.schedule_message_failed_by_limit_reached",
            );
          case 3:
            return c.localize("LocalErrorType.file");
          case 4:
            var t = V.createByteSizeStringWithUnit(e[2]);
            return c.localize("LocalErrorType.file_size", { sizeString: t });
          case 5:
            t = V.createByteSizeStringWithUnit(e[2]);
            return c.localize("LocalErrorType.storage_size", { sizeString: t });
          case 6:
            return c.localize(
              "LocalErrorType.schedule_message_failed_by_unknown",
            );
        }
      }),
      (De.getNoteId = function (e) {
        switch (e[1]) {
          case 9:
            return m.Some(e[2]);
          case 10:
            return m.Some(e[3]);
          case 11:
            return m.Some(e[3]);
          case 12:
            return m.Some(e[2]);
          default:
            return m.None;
        }
      }),
      (De.convertFileErrorForNote = function (e, t) {
        if (null == t) return null;
        if (403 == t.code) return R.UploadForbiddenForNote(e);
        if (502 == t.code && null != t.detail) {
          var n = D.fromIntOrInt64(t.detail.limit_max);
          if ("file_size" == t.detail.limit_target)
            return R.UploadSizeOveredForNote(n, e);
          if ("storage_size" == t.detail.limit_target)
            return R.UploadSizeOveredForNote(n, e);
        }
        return R.GeneralFileErrorForNote(e);
      }),
      (De.getNoteForwardingForbiddenError = function () {
        return R.NoteForwardingForbidden;
      }),
      (De.getNoteTextTooLongError = function () {
        return R.NoteTextTooLong;
      }),
      (De.getNoteTextTooComplexError = function () {
        return R.NoteTextTooComplex;
      }),
      function (e) {
        (ir.call(this, e),
          null != e &&
            ((this.email = e.email),
            (this.signinId = e.signin_id),
            null != e.profiles &&
              (this.profiles = G.__cast(e.profiles, Array).map(function (e) {
                return new Tr(e);
              })),
            null != e.departments) &&
            (this.departments = G.__cast(e.departments, Array).map(
              function (e) {
                return new nr(e);
              },
            )));
      }),
    O =
      (((t["albero.entity.Me"] = fr).__name__ = ["albero", "entity", "Me"]),
      (fr.__super__ = ir),
      (fr.prototype = e(ir.prototype, {
        toDomainUser: function (e) {
          var t = new ve();
          if (
            ((t.id = this.id),
            (t.status = this.status),
            (t.domainId = e),
            (t.displayName = this.displayName),
            (t.canonicalDisplayName = this.canonicalDisplayName),
            (t.phoneticDisplayName = this.phoneticDisplayName),
            (t.canonicalPhoneticDisplayName =
              this.canonicalPhoneticDisplayName),
            (t.profileImageUrl = this.profileImageUrl),
            (t.updatedAt = this.updatedAt),
            (t.email = this.email),
            null != this.profiles)
          )
            for (var n = 0, i = this.profiles; n < i.length; ) {
              var r = i[n],
                o = (++n, r.domainId);
              if (
                null != o &&
                null != e &&
                o.high == e.high &&
                o.low == e.low
              ) {
                t.profileContact = r.itemValues;
                break;
              }
            }
          if (null != this.departments)
            for (var a = 0, s = this.departments; a < s.length; ) {
              var l = s[a],
                u = (++a, l.domainId);
              if (
                null != u &&
                null != e &&
                u.high == e.high &&
                u.low == e.low
              ) {
                t.departments = l.departments;
                break;
              }
            }
          return t;
        },
        __class__: fr,
      })),
      function (e) {
        ((this.scheduleAt = null),
          (this.mention = !1),
          null == e
            ? ((this.readUserIds = []), (this.favorite = new ho(null)))
            : ((this.id = e.message_id),
              (this.talkId = e.talk_id),
              (this.userId = e.user_id),
              (this.type = O.typeOf(e.type)),
              (this.content = e.content),
              (this.createdAt = e.created_at),
              (this.readCount = e.read_count),
              (this.readUserIds =
                null == e.read_user_ids ? [] : e.read_user_ids),
              (this.mention = null != e.mention && e.mention),
              (this.isUnreadCountTarget =
                null == e.is_unread_count_target || e.is_unread_count_target),
              (this.isBotMessage = null != e.bot_message && e.bot_message),
              (this.favorite = new ho(e.favorite))));
      }),
    d =
      (((t["albero.entity.Message"] = O).__name__ = [
        "albero",
        "entity",
        "Message",
      ]),
      (O.typeOf = function (e) {
        switch (e) {
          case 0:
            return d.system;
          case 1:
            return d.text;
          case 2:
            return d.stamp;
          case 3:
            return d.geo;
          case 4:
            return d.file;
          case 5:
            return d.textMultipleFile;
          case 7:
            return d.deleted;
          case 8:
            return d.noteShared;
          case 9:
            return d.noteDeleted;
          case 10:
            return d.noteCreated;
          case 11:
            return d.noteUpdated;
          case 12:
            return d.originalStamp;
          case 500:
            return d.yesOrNo;
          case 501:
            return d.yesOrNoReply;
          case 502:
            return d.selectOne;
          case 503:
            return d.selectOneReply;
          case 504:
            return d.todo;
          case 505:
            return d.todoDone;
          case 506:
            return d.yesOrNoClosed;
          case 507:
            return d.selectOneClosed;
          case 508:
            return d.todoClosed;
          case 509:
            return d.openConference;
          case 600:
            return d.phoneCall;
          case 601:
            return d.phoneReceive;
          default:
            return d.unknown;
        }
      }),
      (O.enumIndex = function (e) {
        switch (e[1]) {
          case 0:
          case 1:
          case 2:
          case 3:
          case 4:
          case 5:
          case 7:
          case 8:
          case 9:
          case 10:
          case 11:
          case 12:
            return e[1];
          case 13:
          case 14:
          case 15:
          case 16:
          case 17:
          case 18:
          case 19:
          case 20:
          case 21:
            return 500 + e[1] - 13;
          case 23:
          case 24:
            return 600 + e[1] - 23;
          default:
            return -1;
        }
      }),
      (O.typeString = function (e) {
        return P.string(e);
      }),
      (O.getDisplayTextWithoutEscape = function (e, t) {
        switch (e[1]) {
          case 1:
            return t;
          case 7:
            return O.getDeletedText();
          default:
            return (
              F.e("Unsupported message.type is detected.", e),
              O.getUnsupportText()
            );
        }
      }),
      (O.getDeletedText = function () {
        return c.localize("MessageHelper.deleted");
      }),
      (O.getUnsupportText = function () {
        return c.localize("MessageHelper.not_support");
      }),
      (O.getCorrectedFilesContentForForwarding = function (e) {
        return K.filterDeleted(e).map(function (e) {
          return O.getCorrectedFileContentForForwarding(e);
        });
      }),
      (O.getCorrectedFileContentForForwarding = function (e) {
        var t = {
          url: O.getCorrectedUrlForForwarding(e.url),
          content_type: e.content_type,
          content_size: e.content_size,
          name: e.name,
          file_id: e.file_id,
        };
        return (
          null != e.thumbnail_url &&
            (t.thumbnail_url = O.getCorrectedUrlForForwarding(e.thumbnail_url)),
          t
        );
      }),
      (O.getCorrectedUrlForForwarding = function (e) {
        e = e.split("?message_id=");
        return 0 != e.length ? e[0] : "";
      }),
      (O.prototype = {
        isDummy: function () {
          return null != this.id && this.id.high < 0;
        },
        isScheduledMessage: function () {
          return null != this.scheduleAt;
        },
        messageStringForKeywordDetection: function (e) {
          switch (this.type[1]) {
            case 0:
              return "";
            case 1:
              return this.content;
            case 3:
            case 4:
              return "";
            case 5:
              return null != this.content.text && 0 < this.content.text.length
                ? this.content.text
                : "";
            case 7:
              return "";
            case 2:
            case 12:
              return null != this.content.text && 0 < this.content.text.length
                ? this.content.text
                : "";
            case 13:
              return [
                this.content.question,
                c.localize("MessageStringMaker.yes"),
                c.localize("MessageStringMaker.no"),
              ].join(" ");
            case 14:
              return e && null != this.content.response
                ? this.content.response
                  ? c.localize("MessageStringMaker.yes")
                  : c.localize("MessageStringMaker.no")
                : "";
            case 15:
              return null != this.content.options
                ? P.string(this.content.question) +
                    " " +
                    P.string(this.content.options.join(" "))
                : this.content.question;
            case 16:
              return e &&
                null != this.content.response &&
                null != this.content.options &&
                this.content.options.length > this.content.response
                ? this.content.options[this.content.response]
                : "";
            case 17:
              return [
                this.content.title,
                c.localize("MessageStringMaker.achieve"),
                c.localize("MessageStringMaker.remand"),
              ].join(" ");
            case 18:
              return e && null != this.content.done
                ? this.content.done
                  ? c.localize("MessageStringMaker.achieve")
                  : c.localize("MessageStringMaker.remand")
                : "";
            default:
              return "";
          }
        },
        messageStringForCompactMessage: function () {
          switch (this.type[1]) {
            case 1:
              return this.content;
            case 5:
              if (null != this.content.text && 0 < this.content.text.length)
                return this.content.text;
              break;
            case 7:
              return c.localize("Message.bookmark_label_deleted");
            case 8:
              return [
                c.localize("Message.bookmark_label_note_shared"),
                new oo(this.content).title,
              ].join(" ");
            case 10:
              return [
                c.localize("Message.bookmark_label_note_created"),
                new oo(this.content).title,
              ].join(" ");
            case 11:
              return [
                c.localize("Message.bookmark_label_note_updated"),
                new oo(this.content).title,
              ].join(" ");
            case 2:
            case 12:
              if (null != this.content.text && 0 < this.content.text.length)
                return this.content.text;
              break;
            case 14:
            case 16:
              return [
                c.localize("Message.bookmark_label_answer"),
                this.content.question,
              ].join(" ");
            case 13:
            case 15:
              return [
                c.localize("Message.bookmark_label_question"),
                this.content.question,
              ].join(" ");
            case 17:
              return [
                c.localize("Message.bookmark_label_task"),
                this.content.title,
              ].join(" ");
            case 18:
              return [
                this.content.done
                  ? c.localize("Message.bookmark_label_completed")
                  : c.localize("Message.bookmark_label_reverted"),
                this.content.title,
              ].join(" ");
            case 19:
            case 20:
              return [
                c.localize("Message.bookmark_label_closed"),
                this.content.question,
              ].join(" ");
            case 21:
              return [
                c.localize("Message.bookmark_label_closed"),
                this.content.title,
              ].join(" ");
          }
          return "";
        },
        getContentsForNote: function () {
          switch (this.type[1]) {
            case 8:
            case 10:
            case 11:
              return new oo(this.content);
            default:
              return null;
          }
        },
        messageStringRaw: function () {
          switch (this.type[1]) {
            case 0:
              return "";
            case 1:
              return this.content;
            case 3:
            case 4:
              return "";
            case 5:
              return null != this.content.text && 0 < this.content.text.length
                ? this.content.text
                : "";
            case 8:
            case 9:
            case 10:
            case 11:
              return "";
            case 2:
            case 12:
              return null != this.content.text && 0 < this.content.text.length
                ? this.content.text
                : "";
            case 14:
            case 16:
              return "";
            case 13:
            case 15:
              return this.content.question;
            case 17:
              return this.content.title;
            case 18:
              return "";
            default:
              return "";
          }
        },
        getContentsForNoteDelete: function () {
          return 9 == this.type[1] ? new ao(this.content) : null;
        },
        getContentsForOriginalStamp: function () {
          return 12 == this.type[1] ? new so(this.content) : null;
        },
        getContentsForStamp: function () {
          return 2 == this.type[1] ? new uo(this.content) : null;
        },
        getChoices: function () {
          return 15 != this.type[1] || null == this.content.options
            ? []
            : this.content.options;
        },
        isAllowedToDelete: function (e) {
          return (
            !!(
              this.isSentBy(e) &&
              this.isNotSystemMessage() &&
              this.isNotActionStampMessage() &&
              this.isNotNoteMessage()
            ) && this.isNotDeletedMessage()
          );
        },
        isAllowedToReply: function (e) {
          return (
            !(!this.isNotSentBy(e) || !this.isNotSystemMessage()) &&
            this.isNotDeletedMessage()
          );
        },
        isAllowedToShowSelectionMessageMenu: function () {
          return (
            !(
              this.type != d.text &&
              !this.isStampMessageWithText() &&
              !this.isTextMultipleFileWithText()
            ) || this.isActionStampMessage()
          );
        },
        isAllowedToQuote: function (e) {
          return (
            !!(
              this.isNotSystemMessage() &&
              this.isNotDeletedMessage() &&
              this.isNotFileMessage() &&
              this.isNotStampMessageWithEmptyText()
            ) && this.isNotTextMultipleFileWithEmptyText()
          );
        },
        isAllowedToCopy: function () {
          return (
            !!(
              this.isNotSystemMessage() &&
              this.isNotDeletedMessage() &&
              this.isNotFileMessage() &&
              this.isNotStampMessageWithEmptyText()
            ) && this.isNotTextMultipleFileWithEmptyText()
          );
        },
        isAllowedToCopyLink: function () {
          return (
            !(!this.isNotSystemMessage() || !this.isNotDeletedMessage()) &&
            this.isNotStampMessageWithEmptyText()
          );
        },
        isAllowedToForward: function (e) {
          return (
            !!(
              this.isNotSystemMessage() &&
              this.isNotDeletedMessage() &&
              this.isNotActionStampMessage() &&
              this.isNotDeletedFileMessage(e) &&
              this.isNotTextMultipleFileAllDeletedWithEmptyText(e) &&
              this.isNotNoteMessage()
            ) && this.isNotOpenConference()
          );
        },
        isAllowedToForwardOnOnlyTextOrStamp: function () {
          return (
            !(!this.isTextMessage() && !this.isStampMessage()) ||
            this.isTextMultipleFileWithText()
          );
        },
        isAllowedToAddToFavorites: function () {
          return (
            !!(
              this.isNotSystemMessage() &&
              this.isNotDeletedMessage() &&
              this.isNotOpenConference() &&
              this.isNotStampMessageWithEmptyText() &&
              this.isNotNoteDeletedMessage()
            ) && this.isNotFavorite()
          );
        },
        isAllowedToRemoveFromFavorites: function () {
          return (
            !!(
              this.isNotSystemMessage() &&
              this.isNotDeletedMessage() &&
              this.isNotOpenConference() &&
              this.isNotStampMessageWithEmptyText() &&
              this.isNotNoteDeletedMessage()
            ) && this.isFavorite()
          );
        },
        isMatchCategory: function (e) {
          switch (e[1]) {
            case 0:
              return (
                !!(
                  this.isTextMessage() ||
                  this.isFileMessage() ||
                  this.isTextMultipleFile()
                ) || this.isStampMessageWithText()
              );
            case 1:
              return !1;
          }
        },
        isSentBy: function (e) {
          var t = this.userId;
          return null != e && null != t && e.high == t.high && e.low == t.low;
        },
        isNotSentBy: function (e) {
          return !this.isSentBy(e);
        },
        isSystemMessage: function () {
          return this.type == d.system;
        },
        isNotSystemMessage: function () {
          return !this.isSystemMessage();
        },
        isActionStampMessage: function () {
          return (
            this.type == d.yesOrNo ||
            this.type == d.yesOrNoReply ||
            this.type == d.selectOne ||
            this.type == d.selectOneReply ||
            this.type == d.todo ||
            this.type == d.todoDone ||
            this.type == d.yesOrNoClosed ||
            this.type == d.selectOneClosed ||
            this.type == d.todoClosed
          );
        },
        isNotActionStampMessage: function () {
          return !this.isActionStampMessage();
        },
        isReplyActionStampMessage: function () {
          return (
            this.type == d.selectOneReply ||
            this.type == d.yesOrNoReply ||
            this.type == d.todoDone
          );
        },
        isClosedActionStampMessage: function () {
          return (
            this.type == d.yesOrNoClosed ||
            this.type == d.selectOneClosed ||
            this.type == d.todoClosed
          );
        },
        isTextMultipleFile: function () {
          return this.type == d.textMultipleFile;
        },
        isSendingActionStampMessage: function () {
          return (
            this.type == d.yesOrNo ||
            this.type == d.selectOne ||
            this.type == d.todo
          );
        },
        isNoteMessage: function () {
          return (
            this.type == d.noteShared ||
            this.type == d.noteDeleted ||
            this.type == d.noteCreated ||
            this.type == d.noteUpdated
          );
        },
        isNoteDeletedMessage: function () {
          return this.type == d.noteDeleted;
        },
        isNotNoteMessage: function () {
          return !this.isNoteMessage();
        },
        isNotNoteDeletedMessage: function () {
          return !this.isNoteDeletedMessage();
        },
        isDeletedMessage: function () {
          return this.type == d.deleted;
        },
        isNotDeletedMessage: function () {
          return !this.isDeletedMessage();
        },
        isFileMessage: function () {
          return this.type == d.file;
        },
        isNotFileMessage: function () {
          return !this.isFileMessage();
        },
        isStampMessage: function () {
          return this.type == d.stamp || this.type == d.originalStamp;
        },
        isNotStampMessage: function () {
          return !this.isStampMessage();
        },
        isStampMessageWithText: function () {
          return !!this.isStampMessage() && this.isNotObjectWithEmptyText();
        },
        isStampMessageWithEmptyText: function () {
          return !!this.isStampMessage() && this.isObjectWithEmptyText();
        },
        isNotStampMessageWithEmptyText: function () {
          return !!this.isNotStampMessage() || this.isNotObjectWithEmptyText();
        },
        isObjectWithEmptyText: function () {
          return Ln.isEmpty(this.content.text);
        },
        isNotObjectWithEmptyText: function () {
          return !this.isObjectWithEmptyText();
        },
        isNotTextMultipleFile: function () {
          return !this.isTextMultipleFile();
        },
        isTextMultipleFileWithText: function () {
          return !!this.isTextMultipleFile() && this.isNotObjectWithEmptyText();
        },
        isTextMultipleFileWithEmptyText: function () {
          return !!this.isTextMultipleFile() && this.isObjectWithEmptyText();
        },
        isNotTextMultipleFileWithEmptyText: function () {
          return (
            !!this.isNotTextMultipleFile() || this.isNotObjectWithEmptyText()
          );
        },
        isNotDeletedFileMessage: function (e) {
          return !!this.isNotFileMessage() || !e;
        },
        isNotTextMultipleFileAllDeletedWithEmptyText: function (e) {
          return !!this.isNotTextMultipleFileWithEmptyText() || !e;
        },
        isTextMessage: function () {
          return this.type == d.text;
        },
        isOpenConference: function () {
          return this.type == d.openConference;
        },
        isNotOpenConference: function () {
          return !this.isOpenConference();
        },
        getInReplyTo: function () {
          return this.isReplyActionStampMessage() ||
            this.isClosedActionStampMessage()
            ? m.Some(this.content.in_reply_to)
            : m.None;
        },
        isNotFavorite: function () {
          return !this.isFavorite();
        },
        isFavorite: function () {
          return this.favorite.isFavorite();
        },
        isNotCalloutHolder: function () {
          return (
            !(
              !this.isFileMessage() && !this.isTextMultipleFileWithEmptyText()
            ) || this.isStampMessageWithEmptyText()
          );
        },
        isNotReactionTarget: function () {
          return (
            !!(
              this.isSystemMessage() ||
              this.isDeletedMessage() ||
              this.isOpenConference() ||
              this.isNoteDeletedMessage()
            ) || this.isDummy()
          );
        },
        isReactionTarget: function () {
          return !this.isNotReactionTarget();
        },
        close: function () {
          this.content.closed = !0;
        },
        markAsDeleted: function () {
          ((this.type = d.deleted), (this.content = ""));
        },
        updateFavorite: function (e) {
          this.favorite.update(e);
        },
        handOver: function (e) {
          this.addReadUsersUntillMaxCount(e.readUserIds);
        },
        addReadUsersUntillMaxCount: function (e) {
          var t = this;
          this.type == d.deleted ||
            16 <= this.readUserIds.length ||
            (16 <= e.length
              ? (this.readUserIds = e.slice(0, 16))
              : ((e = e.filter(function (e) {
                  return D.notContains(t.readUserIds, e);
                })),
                (this.readUserIds = this.readUserIds.concat(e).slice(0, 16))));
        },
        createReadCountText: function () {
          return 16 <= this.readUserIds.length
            ? "15+"
            : P.string(this.readUserIds.length);
        },
        hasNoReadUsers: function () {
          return 0 == this.readUserIds.length;
        },
        deleteFile: function (n) {
          var e, t;
          switch (this.type[1]) {
            case 4:
              e = [this.content];
              break;
            case 5:
              e = this.content.files;
              break;
            default:
              e = [];
          }
          0 != e.length &&
            null !=
              (t = k.find(e, function (e) {
                var e = e.file_id,
                  t = n.fileId;
                return (
                  null != e && null != t && e.high == t.high && e.low == t.low
                );
              })) &&
            (t.deleted = !0);
        },
        getQuestionTitle: function () {
          switch (this.type[1]) {
            case 17:
            case 18:
            case 21:
              return null == this.content.title ? "" : this.content.title;
            case 13:
            case 14:
            case 15:
            case 16:
            case 19:
            case 20:
              return null == this.content.question ? "" : this.content.question;
            default:
              return "";
          }
        },
        getChoiceCloseResponse: function (n) {
          var e,
            t,
            i = c.localize("QuestionHelper.close_answer") + "\n";
          return (
            null == this.content.responses
              ? (i += "...")
              : ((e = this.content.responses),
                (t = this.content.last_response),
                0 == Dr.closingTypeOf(this.content.closing_type)[1]
                  ? (i +=
                      null != t
                        ? c.localize("QuestionHelper.last_response", {
                            content: n(e[t].content),
                            interpolation: { escapeValue: !1 },
                          })
                        : c.localize("QuestionHelper.no_answer"))
                  : (i +=
                      null != t
                        ? null == e
                          ? ""
                          : e
                              .map(function (e) {
                                var t = null == e.count ? 0 : e.count;
                                return (
                                  n(e.content) +
                                  ": " +
                                  t +
                                  c.localize("QuestionHelper.name")
                                );
                              })
                              .join("\n")
                        : c.localize("QuestionHelper.no_answer"))),
            i
          );
        },
        getTodoCloseResponse: function () {
          var e = c.localize("QuestionHelper.close_task") + "\n";
          if (null == this.content.responses) e += "...";
          else {
            var t = this.content.responses,
              n = this.content.last_response;
            if (0 == Dr.closingTypeOf(this.content.closing_type)[1])
              e +=
                null != n && "DONE" == t[n].content
                  ? c.localize("QuestionHelper.achieve_task")
                  : c.localize("QuestionHelper.not_achieve_task");
            else {
              var i = 0;
              if (null != t)
                for (var r = 0; r < t.length; ) {
                  var o = t[r];
                  if ((++r, "DONE" == o.content)) {
                    i = null == o.count ? 0 : o.count;
                    break;
                  }
                }
              e +=
                0 != i
                  ? c.localize("QuestionHelper.completed") +
                    ": " +
                    i +
                    c.localize("QuestionHelper.name")
                  : c.localize("QuestionHelper.no_achiever");
            }
          }
          return e;
        },
        createMessageForForwarding: function (e, t) {
          switch (e[1]) {
            case 0:
              return null;
            case 1:
              switch (this.type[1]) {
                case 4:
                  var n = O.getCorrectedFilesContentForForwarding([
                    this.content,
                  ]);
                  return 0 == n.length
                    ? null
                    : (((i = new O()).type = d.file), (i.content = n[0]), i);
                case 5:
                  var i,
                    n = O.getCorrectedFilesContentForForwarding(
                      this.content.files,
                    );
                  return 0 == n.length
                    ? ((i = t(this.content.text)), this.createTextMessage(i))
                    : (((i = new O()).type = d.textMultipleFile),
                      (i.content = { files: n }),
                      null != this.content.text &&
                        ((n = t(this.content.text)), (i.content.text = n)),
                      i);
              }
              break;
            case 2:
              var r;
              switch (this.type[1]) {
                case 4:
                  return null;
                case 5:
                  return null == this.content.text
                    ? null
                    : ((r = t(this.content.text)), this.createTextMessage(r));
              }
              break;
            case 3:
              return null;
          }
          switch (this.type[1]) {
            case 1:
              var o = t(this.content);
              return this.createTextMessage(o);
            case 2:
              o = new O();
              return (
                (o.type = d.stamp),
                (o.content = {
                  stamp_set: this.content.stamp_set,
                  stamp_index: this.content.stamp_index,
                }),
                null != this.content.text &&
                  ((a = t(this.content.text)), (o.content.text = a)),
                o
              );
            case 5:
              var a = new O();
              return (
                (a.type = d.textMultipleFile),
                (a.content = { files: this.content.files }),
                null != this.content.text &&
                  ((o = t(this.content.text)), (a.content.text = o)),
                a
              );
            case 12:
              o = new O();
              return (
                (o.type = d.originalStamp),
                (o.content = {
                  stampset_id: this.content.stampset_id,
                  stamp_id: this.content.stamp_id,
                }),
                null != this.content.text &&
                  ((a = t(this.content.text)), (o.content.text = a)),
                o
              );
            default:
              return null;
          }
        },
        createTextMessage: function (e) {
          var t;
          return null == e || 0 == e.length
            ? null
            : (((t = new O()).type = d.text), (t.content = e), t);
        },
        __class__: O,
      }),
      (t["albero.entity.MessageType"] = {
        __ename__: ["albero", "entity", "MessageType"],
        __constructs__: [
          "system",
          "text",
          "stamp",
          "geo",
          "file",
          "textMultipleFile",
          "unused",
          "deleted",
          "noteShared",
          "noteDeleted",
          "noteCreated",
          "noteUpdated",
          "originalStamp",
          "yesOrNo",
          "yesOrNoReply",
          "selectOne",
          "selectOneReply",
          "todo",
          "todoDone",
          "yesOrNoClosed",
          "selectOneClosed",
          "todoClosed",
          "openConference",
          "phoneCall",
          "phoneReceive",
          "unknown",
        ],
      })),
    mr =
      ((d.system = ["system", 0]),
      (d.system.toString = a),
      ((d.system.__enum__ = d).text = ["text", 1]),
      (d.text.toString = a),
      ((d.text.__enum__ = d).stamp = ["stamp", 2]),
      (d.stamp.toString = a),
      ((d.stamp.__enum__ = d).geo = ["geo", 3]),
      (d.geo.toString = a),
      ((d.geo.__enum__ = d).file = ["file", 4]),
      (d.file.toString = a),
      ((d.file.__enum__ = d).textMultipleFile = ["textMultipleFile", 5]),
      (d.textMultipleFile.toString = a),
      ((d.textMultipleFile.__enum__ = d).unused = ["unused", 6]),
      (d.unused.toString = a),
      ((d.unused.__enum__ = d).deleted = ["deleted", 7]),
      (d.deleted.toString = a),
      ((d.deleted.__enum__ = d).noteShared = ["noteShared", 8]),
      (d.noteShared.toString = a),
      ((d.noteShared.__enum__ = d).noteDeleted = ["noteDeleted", 9]),
      (d.noteDeleted.toString = a),
      ((d.noteDeleted.__enum__ = d).noteCreated = ["noteCreated", 10]),
      (d.noteCreated.toString = a),
      ((d.noteCreated.__enum__ = d).noteUpdated = ["noteUpdated", 11]),
      (d.noteUpdated.toString = a),
      ((d.noteUpdated.__enum__ = d).originalStamp = ["originalStamp", 12]),
      (d.originalStamp.toString = a),
      ((d.originalStamp.__enum__ = d).yesOrNo = ["yesOrNo", 13]),
      (d.yesOrNo.toString = a),
      ((d.yesOrNo.__enum__ = d).yesOrNoReply = ["yesOrNoReply", 14]),
      (d.yesOrNoReply.toString = a),
      ((d.yesOrNoReply.__enum__ = d).selectOne = ["selectOne", 15]),
      (d.selectOne.toString = a),
      ((d.selectOne.__enum__ = d).selectOneReply = ["selectOneReply", 16]),
      (d.selectOneReply.toString = a),
      ((d.selectOneReply.__enum__ = d).todo = ["todo", 17]),
      (d.todo.toString = a),
      ((d.todo.__enum__ = d).todoDone = ["todoDone", 18]),
      (d.todoDone.toString = a),
      ((d.todoDone.__enum__ = d).yesOrNoClosed = ["yesOrNoClosed", 19]),
      (d.yesOrNoClosed.toString = a),
      ((d.yesOrNoClosed.__enum__ = d).selectOneClosed = [
        "selectOneClosed",
        20,
      ]),
      (d.selectOneClosed.toString = a),
      ((d.selectOneClosed.__enum__ = d).todoClosed = ["todoClosed", 21]),
      (d.todoClosed.toString = a),
      ((d.todoClosed.__enum__ = d).openConference = ["openConference", 22]),
      (d.openConference.toString = a),
      ((d.openConference.__enum__ = d).phoneCall = ["phoneCall", 23]),
      (d.phoneCall.toString = a),
      ((d.phoneCall.__enum__ = d).phoneReceive = ["phoneReceive", 24]),
      (d.phoneReceive.toString = a),
      ((d.phoneReceive.__enum__ = d).unknown = ["unknown", 25]),
      (d.unknown.toString = a),
      ((d.unknown.__enum__ = d).__empty_constructs__ = [
        d.system,
        d.text,
        d.stamp,
        d.geo,
        d.file,
        d.textMultipleFile,
        d.unused,
        d.deleted,
        d.noteShared,
        d.noteDeleted,
        d.noteCreated,
        d.noteUpdated,
        d.originalStamp,
        d.yesOrNo,
        d.yesOrNoReply,
        d.selectOne,
        d.selectOneReply,
        d.todo,
        d.todoDone,
        d.yesOrNoClosed,
        d.selectOneClosed,
        d.todoClosed,
        d.openConference,
        d.phoneCall,
        d.phoneReceive,
        d.unknown,
      ]),
      (t["albero.entity.MessageCategory"] = {
        __ename__: ["albero", "entity", "MessageCategory"],
        __constructs__: ["textOrFiles", "unknown"],
      })),
    pr =
      ((mr.textOrFiles = ["textOrFiles", 0]),
      (mr.textOrFiles.toString = a),
      ((mr.textOrFiles.__enum__ = mr).unknown = ["unknown", 1]),
      (mr.unknown.toString = a),
      ((mr.unknown.__enum__ = mr).__empty_constructs__ = [
        mr.textOrFiles,
        mr.unknown,
      ]),
      ((t["albero.entity.MessageDeletion"] = Re).__name__ = [
        "albero",
        "entity",
        "MessageDeletion",
      ]),
      (Re.prototype = { __class__: Re }),
      ((t["albero.entity.MessageReadStatus"] = Oe).__name__ = [
        "albero",
        "entity",
        "MessageReadStatus",
      ]),
      (Oe.prototype = { __class__: Oe }),
      ((t["albero.entity.MessageReadStatusesUpdate"] = Fe).__name__ = [
        "albero",
        "entity",
        "MessageReadStatusesUpdate",
      ]),
      (Fe.prototype = { __class__: Fe }),
      function () {
        ((this.messageIdsFromThisDevice = new B()), (this.messages = new B()));
      }),
    gr =
      (((t["albero.entity.MessageStore"] = pr).__name__ = [
        "albero",
        "entity",
        "MessageStore",
      ]),
      (pr.prototype = {
        getMessage: function (e) {
          var t;
          return null == e
            ? null
            : ((e = "_" + e.high + "_" + e.low),
              (t = this.messages),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        setMessage: function (e) {},
        removeMessage: function (e) {
          this.messages.remove("_" + e.high + "_" + e.low);
        },
        setFromThisDevice: function (e) {},
        isFromThisDevice: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.messageIdsFromThisDevice;
          return null != Y[e] ? t.existsReserved(e) : t.h.hasOwnProperty(e);
        },
        deleteFile: function (e) {
          var t = this.getMessage(e.messageId);
          null != t && t.deleteFile(e);
        },
        closeConference: function (e) {
          e = this.getMessage(e.messageId);
          null != e && e.close();
        },
        closeAction: function (e) {
          e = this.getMessage(e);
          null != e && e.close();
        },
        deleteFavoriteMessage: function (e) {
          var t = this.getMessage(e.getMessageId());
          null != t && t.updateFavorite(e);
        },
        addFavoriteMessage: function (e) {
          var t = this.getMessage(e.getMessageId());
          null != t ? t.updateFavorite(e) : this.setMessage(e.getMessage());
        },
        __class__: pr,
      }),
      (t["albero.entity.MessagesOrder"] = {
        __ename__: ["albero", "entity", "MessagesOrder"],
        __constructs__: ["asc", "desc"],
      })),
    yr =
      ((gr.asc = ["asc", 0]),
      (gr.asc.toString = a),
      ((gr.asc.__enum__ = gr).desc = ["desc", 1]),
      (gr.desc.toString = a),
      ((gr.desc.__enum__ = gr).__empty_constructs__ = [gr.asc, gr.desc]),
      ((t["albero.entity.MessagesOrderProc"] = xe).__name__ = [
        "albero",
        "entity",
        "MessagesOrderProc",
      ]),
      (xe.orderToInt = function (e) {
        switch (e[1]) {
          case 0:
            return 1;
          case 1:
            return 2;
        }
      }),
      (t["albero.entity.PaneType"] = {
        __ename__: ["albero", "entity", "PaneType"],
        __constructs__: [
          "SingleTalkPane",
          "MultiTalkPane1",
          "MultiTalkPane2",
          "MultiTalkPane3",
        ],
      })),
    Sr =
      ((yr.SingleTalkPane = ["SingleTalkPane", 0]),
      (yr.SingleTalkPane.toString = a),
      ((yr.SingleTalkPane.__enum__ = yr).MultiTalkPane1 = [
        "MultiTalkPane1",
        1,
      ]),
      (yr.MultiTalkPane1.toString = a),
      ((yr.MultiTalkPane1.__enum__ = yr).MultiTalkPane2 = [
        "MultiTalkPane2",
        2,
      ]),
      (yr.MultiTalkPane2.toString = a),
      ((yr.MultiTalkPane2.__enum__ = yr).MultiTalkPane3 = [
        "MultiTalkPane3",
        3,
      ]),
      (yr.MultiTalkPane3.toString = a),
      ((yr.MultiTalkPane3.__enum__ = yr).__empty_constructs__ = [
        yr.SingleTalkPane,
        yr.MultiTalkPane1,
        yr.MultiTalkPane2,
        yr.MultiTalkPane3,
      ]),
      ((t["albero.entity.PaneTypeHelper"] = Ue).__name__ = [
        "albero",
        "entity",
        "PaneTypeHelper",
      ]),
      (Ue.toPanePrefix = function (e) {
        switch (e[1]) {
          case 0:
            return "";
          case 1:
            return "multi1-";
          case 2:
            return "multi2-";
          case 3:
            return "multi3-";
        }
      }),
      (Ue.getPaneType = function (e) {
        switch (e) {
          case 0:
            return yr.SingleTalkPane;
          case 1:
            return yr.MultiTalkPane1;
          case 2:
            return yr.MultiTalkPane2;
          case 3:
            return yr.MultiTalkPane3;
          default:
            return null;
        }
      }),
      (Ue.getMultiPaneType = function (e) {
        switch (e) {
          case 0:
            return yr.MultiTalkPane1;
          case 1:
            return yr.MultiTalkPane2;
          case 2:
            return yr.MultiTalkPane3;
          default:
            return yr.MultiTalkPane1;
        }
      }),
      (Ue.getMultiPaneTypes = function () {
        return [yr.MultiTalkPane1, yr.MultiTalkPane2, yr.MultiTalkPane3];
      }),
      (Ue.fromString = function (e) {
        switch (e) {
          case "multi1":
            return yr.MultiTalkPane1;
          case "multi2":
            return yr.MultiTalkPane2;
          case "multi3":
            return yr.MultiTalkPane3;
          default:
            return yr.SingleTalkPane;
        }
      }),
      (Ue.toString = function (e) {
        switch (e[1]) {
          case 0:
            return "single";
          case 1:
            return "multi1";
          case 2:
            return "multi2";
          case 3:
            return "multi3";
        }
      }),
      function (e) {
        null != e &&
          ((this.expiration = e.expiration), (this.warning = e.warning));
      }),
    vr =
      (((t["albero.entity.PasswordExpiration"] = Sr).__name__ = [
        "albero",
        "entity",
        "PasswordExpiration",
      ]),
      (Sr.prototype = {
        isExpired: function (e) {
          return this.timeOvered(this.expiration, e);
        },
        needWarning: function (e, t) {
          return (
            !!this.timeOveredOrNotExist(t, e) &&
            this.timeOvered(this.warning, e)
          );
        },
        timeOvered: function (e, t) {
          var n;
          return (
            null != e &&
            (0 == (n = (t.high - e.high) | 0) && (n = w.ucompare(t.low, e.low)),
            0 < (t.high < 0 ? (e.high < 0 ? n : -1) : 0 <= e.high ? n : 1))
          );
        },
        timeOveredOrNotExist: function (e, t) {
          return null == e || this.timeOvered(e, t);
        },
        __class__: Sr,
      }),
      function (e) {
        null != e &&
          ((this.name = e.plan_name),
          (this.trial = !!e.trial),
          (this.free = !!e.free));
      }),
    Tr =
      (((t["albero.entity.Plan"] = vr).__name__ = ["albero", "entity", "Plan"]),
      (vr.prototype = { __class__: vr }),
      function (e) {
        null != e &&
          ((this.domainId = e.domain_id),
          (this.userId = e.user_id),
          (this.itemValues = null) != e.item_values) &&
          (this.itemValues = G.__cast(e.item_values, Array).map(function (e) {
            return new wr(e);
          }));
      }),
    Ir =
      (((t["albero.entity.Profile"] = Tr).__name__ = [
        "albero",
        "entity",
        "Profile",
      ]),
      (Tr.prototype = { __class__: Tr }),
      function (e) {
        if (
          null != e &&
          ((this.domainId = e.domain_id), null != e.item_definitions)
        ) {
          this.itemDefinitions = [];
          for (
            var t = 0, n = G.__cast(e.item_definitions, Array);
            t < n.length;
          ) {
            var i = n[t];
            (++t, this.itemDefinitions.push(new Er(i)));
          }
        }
      }),
    Er =
      (((t["albero.entity.ProfileDefinition"] = Ir).__name__ = [
        "albero",
        "entity",
        "ProfileDefinition",
      ]),
      (Ir.isSameProfileItemDefinitions = function (e, t) {
        if (t.length != e.length) return !1;
        for (var n = 0, i = t.length; n < i; ) {
          var r = n++;
          if (t[r].profileItemId != e[r].profileItemId) return !1;
        }
        return !0;
      }),
      (Ir.prototype = {
        getVisibleItemDefinitions: function (t, n) {
          return (
            null == n && (n = !1),
            null == t && (t = !1),
            null == this.itemDefinitions
              ? []
              : this.itemDefinitions.filter(function (e) {
                  return (
                    !!e.visible &&
                    !(
                      (!t && 10 != e.group) ||
                      (n && (1 == e.profileItemId || 2 == e.profileItemId))
                    )
                  );
                })
          );
        },
        __class__: Ir,
      }),
      function (e) {
        ((this.locked = !1),
          (this.visible = !0),
          null != e &&
            ((this.profileItemId = e.profile_item_id),
            (this.group = e.group),
            (this.name = e.name),
            (this.valueType = e.value_type),
            null != e.visible && (this.visible = e.visible),
            null != e.locked) &&
            (this.locked = e.locked));
      }),
    wr =
      (((t["albero.entity.ProfileItemDefinition"] = Er).__name__ = [
        "albero",
        "entity",
        "ProfileItemDefinition",
      ]),
      (Er.findById = function (e, t) {
        for (var n = 0; n < e.length; ) {
          var i = e[n];
          if ((++n, i.profileItemId == t)) return i;
        }
        return null;
      }),
      (Er.prototype = { __class__: Er }),
      function (e) {
        null != e &&
          ((this.profileItemId = e.profile_item_id),
          (this.value = e.value),
          (this.canonicalValue = e.canonical_value));
      }),
    Nr =
      (((t["albero.entity.ProfileItemValue"] = wr).__name__ = [
        "albero",
        "entity",
        "ProfileItemValue",
      ]),
      (wr.findById = function (e, t) {
        for (var n = 0; n < e.length; ) {
          var i = e[n];
          if ((++n, i.profileItemId == t)) return i;
        }
        return null;
      }),
      (wr.prototype = { __class__: wr }),
      function (e, t) {
        ((this.domainId = e), (this.profileItemValues = t));
      }),
    Ar =
      (((t["albero.entity.ProfileItemValuesPerDomain"] = Nr).__name__ = [
        "albero",
        "entity",
        "ProfileItemValuesPerDomain",
      ]),
      (Nr.convertToApiParams = function (e) {
        return e.map(function (e) {
          return e.convertToApiParamsPerDomain();
        });
      }),
      (Nr.prototype = {
        convertToApiParamsPerDomain: function () {
          return [
            this.domainId,
            this.profileItemValues.map(function (e) {
              return { profile_item_id: e.profileItemId, value: e.value };
            }),
          ];
        },
        __class__: Nr,
      }),
      function (e) {
        ((this.type =
          null == e
            ? Mr.All
            : this.createPushNotificationType(e.type, e.enabled)),
          (this.version = null == e ? 0 : e.version));
      }),
    br =
      (((t["albero.entity.PushNotificationSetting"] = Ar).__name__ = [
        "albero",
        "entity",
        "PushNotificationSetting",
      ]),
      (Ar.createWithParams = function (e, t) {
        var n = new Ar();
        return ((n.type = e), (n.version = t), n);
      }),
      (Ar.prototype = {
        isDisabled: function () {
          return this.type == Mr.None;
        },
        createPushNotificationType: function (e, t) {
          return null != t ? (t ? Mr.All : Mr.None) : br.fromRawValue(e);
        },
        __class__: Ar,
      }),
      function () {}),
    Mr =
      (((t["albero.entity.PushNotificationTypeHelper"] = br).__name__ = [
        "albero",
        "entity",
        "PushNotificationTypeHelper",
      ]),
      (br.fromRawValue = function (e) {
        if (null == e) return Mr.All;
        if (null == e) return Mr.All;
        switch (e) {
          case 0:
            return Mr.None;
          case 1:
            return Mr.All;
          case 2:
            return Mr.Mention;
          default:
            return Mr.All;
        }
      }),
      (br.toRawValue = function (e) {
        switch (e[1]) {
          case 0:
            return 1;
          case 1:
            return 2;
          case 2:
            return 0;
        }
      }),
      (br.fromString = function (e) {
        return x.createEnum(Mr, e);
      }),
      (br.toString = function (e) {
        return e[0];
      }),
      (t["albero.entity.PushNotificationType"] = {
        __ename__: ["albero", "entity", "PushNotificationType"],
        __constructs__: ["All", "Mention", "None"],
      })),
    Cr =
      ((Mr.All = ["All", 0]),
      (Mr.All.toString = a),
      ((Mr.All.__enum__ = Mr).Mention = ["Mention", 1]),
      (Mr.Mention.toString = a),
      ((Mr.Mention.__enum__ = Mr).None = ["None", 2]),
      (Mr.None.toString = a),
      ((Mr.None.__enum__ = Mr).__empty_constructs__ = [
        Mr.All,
        Mr.Mention,
        Mr.None,
      ]),
      function (e) {
        null != e &&
          ((this.id = e.message_id),
          (this.talkId = e.talk_id),
          (this.type = O.typeOf(e.type)),
          (this.content = e.content),
          (this.userId = e.user_id),
          (this.recipientIds = e.assigned_user_ids),
          (this.responses = Cr.createResponses(e.responses)),
          (this.listing = e.listing),
          (this.closingType = Dr.closingTypeOf(e.closing_type)),
          (this.maxResponseId = e.max_response_id),
          (this.lastResponse = e.last_response),
          (this.lastResponseUserId = e.last_response_user_id),
          (this.createdAt = e.created_at),
          (this.updatedAt = e.updated_at),
          (this.responded = e.responded),
          (this.closed = e.closed));
      }),
    $n =
      (((t["albero.entity.Question"] = Cr).__name__ = [
        "albero",
        "entity",
        "Question",
      ]),
      (Cr.createResponses = function (e) {
        var t = h.option(e);
        switch (t[1]) {
          case 0:
            var n = t[2];
            break;
          case 1:
            n = [];
        }
        return n.map(function (e) {
          return new Rr(e);
        });
      }),
      (Cr.createResponsesFromStrings = function (e) {
        var t = h.option(e);
        switch (t[1]) {
          case 0:
            var n = t[2];
            break;
          case 1:
            n = [];
        }
        return n.map(function (e) {
          return new Rr({ content: e, count: 0, user_ids: [] });
        });
      }),
      (Cr.createResponseTexts = function (e, t) {
        switch (e[1]) {
          case 13:
            return ["YES", "NO"];
          case 15:
            return t;
          case 17:
            return ["DONE", "UNDONE"];
          default:
            return [];
        }
      }),
      (Cr.fromMessageAndTalk = function (e, t) {
        var n = new Cr(),
          i =
            ((n.id = e.id),
            (n.talkId = e.talkId),
            (n.type = e.type),
            h.option(e.content.title));
        switch (i[1]) {
          case 0:
            var r = i[2];
            break;
          case 1:
            r = e.content.question;
        }
        ((n.content = r),
          (n.userId = e.userId),
          (n.recipientIds = t.userIds),
          (n.responses = Cr.createResponsesFromStrings(
            Cr.createResponseTexts(e.type, e.content.options),
          )));
        var o = h.option(e.content.listing);
        switch (o[1]) {
          case 0:
            var a = o[2];
            break;
          case 1:
            a = !0;
        }
        ((n.listing = a),
          (n.closingType = Dr.closingTypeOf(e.content.closing_type)),
          (n.maxResponseId = null),
          (n.lastResponse = null),
          (n.lastResponseUserId = null),
          (n.updatedAt = n.createdAt = e.createdAt),
          (n.responded = !1));
        var s = h.option(e.content.closed);
        switch (s[1]) {
          case 0:
            var l = s[2];
            break;
          case 1:
            l = !1;
        }
        return ((n.closed = l), n);
      }),
      (Cr.createTodoDoneMessageWithParams = function (e, t, n) {
        var i = new O();
        return (
          (i.talkId = e),
          (i.type = d.todoDone),
          (i.content = { in_reply_to: t, done: n }),
          i
        );
      }),
      (Cr.createYesOrNoReplyMessageWithParams = function (e, t, n) {
        var i = new O();
        return (
          (i.talkId = e),
          (i.content = { in_reply_to: t, response: n }),
          (i.type = d.yesOrNoReply),
          i
        );
      }),
      (Cr.createSelectOneReplyMessageWithParams = function (e, t, n) {
        var i = new O();
        return (
          (i.talkId = e),
          (i.content = { in_reply_to: t, response: n }),
          (i.type = d.selectOneReply),
          i
        );
      }),
      (Cr.fromTypeToInt = function (e) {
        switch (e[1]) {
          case 0:
            return 0;
          case 1:
            return 1;
        }
      }),
      (Cr.filterToOpBool = function (e) {
        switch (e[1]) {
          case 0:
            return m.Some(!0);
          case 1:
            return m.Some(!1);
          case 2:
            return m.None;
        }
      }),
      (Cr.isDoneResponse = function (e) {
        return "DONE" == e.content;
      }),
      (Cr.prototype = {
        isCompleted: function (e) {
          if (x.enumEq(this.type, d.todo)) {
            if (x.enumEq(this.closingType, kr.any)) {
              if (
                null != this.lastResponse &&
                "DONE" == this.responses[this.lastResponse].content
              )
                return !0;
            } else {
              if (!this.responded) return !1;
              for (var t = 0, n = this.responses; t < n.length; ) {
                var i = n[t];
                if ((++t, i.contains(e))) return "DONE" == i.content;
              }
            }
            return !1;
          }
          return x.enumEq(this.closingType, kr.any)
            ? null != this.lastResponse
            : this.responded;
        },
        createTodoDoneMessage: function (e) {
          return Cr.createTodoDoneMessageWithParams(this.talkId, this.id, e);
        },
        createYesOrNoReplyMessage: function (e) {
          return Cr.createYesOrNoReplyMessageWithParams(
            this.talkId,
            this.id,
            e,
          );
        },
        createSelectOneReplyMessage: function (e) {
          return Cr.createSelectOneReplyMessageWithParams(
            this.talkId,
            this.id,
            e,
          );
        },
        createActionCloseMessage: function () {
          var e = new O();
          switch (
            ((e.talkId = this.talkId),
            (e.content = { in_reply_to: this.id }),
            this.type[1])
          ) {
            case 13:
              e.type = d.yesOrNoClosed;
              break;
            case 15:
              e.type = d.selectOneClosed;
              break;
            case 17:
              e.type = d.todoClosed;
              break;
            default:
              F.e("message type is not action.");
          }
          return e;
        },
        getLastResponse: function () {
          var t = this;
          return h.map(h.option(this.lastResponse), function (e) {
            return t.responses[e];
          });
        },
        getLastDoneResponse: function () {
          return h.filter(this.getLastResponse(), Cr.isDoneResponse);
        },
        getDoneResponse: function () {
          return h.option(k.find(this.responses, Cr.isDoneResponse));
        },
        isSendingQuestion: function () {
          return (
            this.type == d.yesOrNo ||
            this.type == d.selectOne ||
            this.type == d.todo
          );
        },
        isNotSendingQuestion: function () {
          return !this.isSendingQuestion();
        },
        getChoices: function () {
          return this.responses.map(function (e) {
            return e.content;
          });
        },
        updateByMessage: function (e, t) {
          switch (e.type[1]) {
            case 14:
              var n = new _o(e.content).getResponseIndex();
              this.updateByReply(n, e, t);
              break;
            case 16:
              n = new lo(e.content).getResponseIndex();
              this.updateByReply(n, e, t);
              break;
            case 18:
              n = new co(e.content).getResponseIndex();
              this.updateByReply(n, e, t);
              break;
            case 19:
            case 20:
            case 21:
              this.updateByClose();
          }
        },
        updateByClose: function () {
          this.closed = !0;
        },
        updateByReply: function (e, t, n) {
          this.updateByReplySub(e, t.id, t.userId, t.createdAt, n);
        },
        updateByReplySub: function (e, t, n, i, r) {
          (k.iter(this.responses, function (e) {
            e.sub(n);
          }),
            this.responses[e].add(n),
            r && (this.responded = !0));
          var r = this.updatedAt,
            o = (i.high - r.high) | 0;
          (0 == o && (o = w.ucompare(i.low, r.low)),
            0 < (i.high < 0 ? (r.high < 0 ? o : -1) : 0 <= r.high ? o : 1) &&
              ((this.updatedAt = i),
              (this.maxResponseId = t),
              (this.lastResponse = e),
              (this.lastResponseUserId = n)));
        },
        __class__: Cr,
      }),
      (t["albero.entity.QuestionFromType"] = {
        __ename__: ["albero", "entity", "QuestionFromType"],
        __constructs__: ["fromSelf", "fromOther"],
      })),
    $n =
      (($n.fromSelf = ["fromSelf", 0]),
      ($n.fromSelf.toString = a),
      (($n.fromSelf.__enum__ = $n).fromOther = ["fromOther", 1]),
      ($n.fromOther.toString = a),
      (($n.fromOther.__enum__ = $n).__empty_constructs__ = [
        $n.fromSelf,
        $n.fromOther,
      ]),
      (t["albero.entity.QuestionFilter"] = {
        __ename__: ["albero", "entity", "QuestionFilter"],
        __constructs__: ["onlyClosed", "onlyUnclosed", "noFilter"],
      })),
    kr =
      (($n.onlyClosed = ["onlyClosed", 0]),
      ($n.onlyClosed.toString = a),
      (($n.onlyClosed.__enum__ = $n).onlyUnclosed = ["onlyUnclosed", 1]),
      ($n.onlyUnclosed.toString = a),
      (($n.onlyUnclosed.__enum__ = $n).noFilter = ["noFilter", 2]),
      ($n.noFilter.toString = a),
      (($n.noFilter.__enum__ = $n).__empty_constructs__ = [
        $n.onlyClosed,
        $n.onlyUnclosed,
        $n.noFilter,
      ]),
      (t["albero.entity.QuestionClosingType"] = {
        __ename__: ["albero", "entity", "QuestionClosingType"],
        __constructs__: ["any", "all", "unknown"],
      })),
    Dr =
      ((kr.any = ["any", 0]),
      (kr.any.toString = a),
      ((kr.any.__enum__ = kr).all = ["all", 1]),
      (kr.all.toString = a),
      ((kr.all.__enum__ = kr).unknown = ["unknown", 2]),
      (kr.unknown.toString = a),
      ((kr.unknown.__enum__ = kr).__empty_constructs__ = [
        kr.any,
        kr.all,
        kr.unknown,
      ]),
      function () {}),
    Rr =
      (((t["albero.entity.QuestionClosingTypeHelper"] = Dr).__name__ = [
        "albero",
        "entity",
        "QuestionClosingTypeHelper",
      ]),
      (Dr.closingTypeOf = function (e) {
        switch (e) {
          case 0:
            return kr.any;
          case 1:
            return kr.all;
          default:
            return kr.unknown;
        }
      }),
      function (e) {
        ((this.content = e.content),
          (this.count = null == e.count ? 0 : e.count),
          (this.userIds = null == e.user_ids ? [] : e.user_ids));
      }),
    Or =
      (((t["albero.entity.QuestionResponse"] = Rr).__name__ = [
        "albero",
        "entity",
        "QuestionResponse",
      ]),
      (Rr.prototype = {
        add: function (e) {
          this.contains(e) ||
            ((this.count += 1), (this.userIds = this.userIds.concat([e])));
        },
        sub: function (t) {
          this.contains(t) &&
            (--this.count,
            (this.userIds = this.userIds.filter(function (e) {
              return !(t.high == e.high && t.low == e.low);
            })));
        },
        contains: function (e) {
          return D.contains(this.userIds, e);
        },
        __class__: Rr,
      }),
      function () {
        this.questions = new B();
      }),
    Fr =
      (((t["albero.entity.QuestionStore"] = Or).__name__ = [
        "albero",
        "entity",
        "QuestionStore",
      ]),
      (Or.prototype = {
        removeTalkQuestions: function (e) {
          for (var t = this.getQuestions(e), n = 0; n < t.length; ) {
            var i = t[n],
              i = (++n, i.id);
            this.questions.remove("_" + i.high + "_" + i.low);
          }
        },
        getQuestions: function (n, i) {
          return k.array(this.questions).filter(function (e) {
            var t = e.talkId;
            return (
              null != t &&
              null != n &&
              t.high == n.high &&
              t.low == n.low &&
              (null == i ||
                ((t = e.userId),
                null != i && null != t && i.high == t.high && i.low == t.low))
            );
          });
        },
        getQuestion: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.questions;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        setQuestion: function (e) {},
        __class__: Or,
      }),
      function (e) {
        null != e &&
          ((this.msgRetentionPeriod = e.msg_retention_period),
          (this.msgFileSize = e.msg_file_size),
          (this.maxMessageContentLength = e.max_message_content_length),
          (this.maxTalkers = e.max_talkers),
          (this.maxTalks = e.max_talks),
          (this.withAd = !!e.with_ad));
      }),
    xr =
      (((t["albero.entity.Quota"] = Fr).__name__ = [
        "albero",
        "entity",
        "Quota",
      ]),
      (Fr.prototype = { __class__: Fr }),
      function (e, t) {
        null == e
          ? (this.readingLastAnnouncementId = t)
          : ((this.announcementIds = e),
            (this.readingLastAnnouncementId =
              0 < e.length ? e[e.length - 1] : null));
      }),
    Ur =
      (((t["albero.entity.ReadingAnnouncements"] = xr).__name__ = [
        "albero",
        "entity",
        "ReadingAnnouncements",
      ]),
      (xr.prototype = {
        update: function (e) {
          var t,
            n,
            i,
            r = this;
          return (
            null != this.readingLastAnnouncementId &&
            null != e.readingLastAnnouncementId &&
            ((i = this.readingLastAnnouncementId),
            (t = e.readingLastAnnouncementId),
            0 == (n = (i.high - t.high) | 0) && (n = w.ucompare(i.low, t.low)),
            !(
              0 <= (i.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1) ||
              (null != this.announcementIds &&
                null != e.announcementIds &&
                ((i = e.announcementIds.filter(function (e) {
                  var t = r.readingLastAnnouncementId,
                    n = (e.high - t.high) | 0;
                  return (
                    0 == n && (n = w.ucompare(e.low, t.low)),
                    0 <
                      (e.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1)
                  );
                })),
                (this.announcementIds = this.announcementIds.concat(i))),
              (this.readingLastAnnouncementId = e.readingLastAnnouncementId),
              0)
            ))
          );
        },
        getLastAnnouncementId: function () {
          return this.readingLastAnnouncementId;
        },
        canRegardAnnouncementReadCountZero: function (e) {
          var t;
          return (
            null == this.announcementIds ||
            ((t = this.getLastAnnouncementId()),
            null != e && null != t && e.high == t.high && e.low == t.low)
          );
        },
        getUnreadAnnouncementIds: function (n) {
          return null == this.announcementIds
            ? []
            : this.announcementIds.filter(function (e) {
                var t;
                return (
                  null == n ||
                  (0 == (t = (n.high - e.high) | 0) &&
                    (t = w.ucompare(n.low, e.low)),
                  (n.high < 0 ? (e.high < 0 ? t : -1) : 0 <= e.high ? t : 1) <
                    0)
                );
              });
        },
        isReadingAll: function () {
          return null == this.announcementIds;
        },
        __class__: xr,
      }),
      ((t["albero.entity.ReadingMessages"] = Le).__name__ = [
        "albero",
        "entity",
        "ReadingMessages",
      ]),
      (Le.prototype = {
        update: function (e) {
          var t,
            n,
            i,
            r = this;
          return (
            null != this.readingLastMessageId &&
            null != e.readingLastMessageId &&
            ((i = this.readingLastMessageId),
            (t = e.readingLastMessageId),
            0 == (n = (i.high - t.high) | 0) && (n = w.ucompare(i.low, t.low)),
            !(
              0 <= (i.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1) ||
              (null != this.messages &&
                null != e.messages &&
                ((i = e.messages.filter(function (e) {
                  var e = e.messageId,
                    t = r.readingLastMessageId,
                    n = (e.high - t.high) | 0;
                  return (
                    0 == n && (n = w.ucompare(e.low, t.low)),
                    0 <
                      (e.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1)
                  );
                })),
                (this.messages = this.messages.concat(i))),
              (this.readingLastMessageId = e.readingLastMessageId),
              0)
            ))
          );
        },
        getLastMesssgeId: function () {
          return this.readingLastMessageId;
        },
        canRegardTalkReadCountZero: function (e) {
          var t;
          return (
            null == this.messages ||
            ((t = this.getLastMesssgeId()),
            null != e && null != t && e.high == t.high && e.low == t.low)
          );
        },
        getUnreadMessages: function (i) {
          return null == this.messages
            ? []
            : this.messages.filter(function (e) {
                var t =
                    null == i ||
                    ((t = e.messageId),
                    0 == (n = (i.high - t.high) | 0) &&
                      (n = w.ucompare(i.low, t.low)),
                    (i.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1) <
                      0),
                  n = !!t && e.isUnreadCountTarget;
                return !!n && !e.isSendByMe;
              });
        },
        getUnreadMentionMessages: function (e) {
          return this.getUnreadMessages(e).filter(function (e) {
            return e.hasMention;
          });
        },
        isReadingAll: function () {
          return null == this.messages;
        },
        __class__: Le,
      }),
      function (e, t) {
        this.messageId = e.id;
        var n = e.userId;
        ((this.isSendByMe =
          null != n && null != t && n.high == t.high && n.low == t.low),
          (this.isUnreadCountTarget = e.isUnreadCountTarget),
          (this.hasMention = e.mention));
      }),
    Lr =
      (((t["albero.entity.ReadingMessage"] = Ur).__name__ = [
        "albero",
        "entity",
        "ReadingMessage",
      ]),
      (Ur.prototype = { __class__: Ur }),
      ((t["albero.entity.SearchAttachmentsResult"] = Pe).__name__ = [
        "albero",
        "entity",
        "SearchAttachmentsResult",
      ]),
      (Pe.prototype = { __class__: Pe }),
      ((t["albero.entity.SearchMessagesResult"] = Be).__name__ = [
        "albero",
        "entity",
        "SearchMessagesResult",
      ]),
      (Be.prototype = { __class__: Be }),
      ((t["albero.entity.SearchParams"] = He).__name__ = [
        "albero",
        "entity",
        "SearchParams",
      ]),
      (He.dateQueryToText = function (e, t) {
        if (null == e) return t;
        switch (e) {
          case "-1d":
            return c.localize("SearchParams.text_date_yesterday");
          case "-3d":
            return c.localize("SearchParams.text_date_three_days_before");
          case "t":
            return c.localize("SearchParams.text_date_today");
          default:
            return I.replace(e, "-", "/");
        }
      }),
      (He.dateQueryToBaseValue = function (e) {
        if (null == e) return null;
        switch (e) {
          case "-1d":
            return "-1d";
          case "-3d":
            return "-3d";
          case "t":
            return "0d";
          default:
            return I.replace(e, "-", "/");
        }
      }),
      (He.searchTypeToKey = function (e) {
        switch (e[1]) {
          case 0:
            return "messages";
          case 1:
            return "attachments";
        }
      }),
      (He.prototype = {
        toProps: function () {
          var e,
            t = {};
          return (
            null != this.domainId &&
              ((e = this.domainId),
              (t.domain_id_str = "_" + e.high + "_" + e.low)),
            null != this.talkId &&
              ((e = this.talkId), (t.talk_id_str = "_" + e.high + "_" + e.low)),
            (t.keyword = this.keyword),
            (t.search_type = Pr.toLocalizeText(this.searchType)),
            (t.since = this.sinceText),
            (t.until = this.untilText),
            null != this.fromUserId &&
              ((e = this.fromUserId),
              (t.from_id_str = "_" + e.high + "_" + e.low)),
            t
          );
        },
        copy: function () {
          return new He(this.toProps());
        },
        eq: function (e) {
          return !!(
            this.isSameInt64(this.domainId, e.domainId) &&
            this.isSameInt64(this.talkId, e.talkId) &&
            this.keyword == e.keyword &&
            x.enumEq(this.searchType, e.searchType) &&
            this.sinceText == e.sinceText &&
            this.untilText == e.untilText &&
            this.isSameInt64(this.fromUserId, e.fromUserId)
          );
        },
        isSameInt64: function (e, t) {
          return (
            (null == e && null == t) ||
            (null != e && null != t && e.high == t.high && e.low == t.low)
          );
        },
        getQueryText: function () {
          var e = [];
          return (
            null != this.keyword &&
              0 != this.keyword.length &&
              e.push(this.keyword),
            null != this.sinceText && e.push("since:" + this.sinceText),
            null != this.untilText && e.push("until:" + this.untilText),
            null != this.fromUserId &&
              e.push("from:" + Va.toString(this.fromUserId)),
            e.join(" ")
          );
        },
        getPeriodText: function () {
          return (
            He.dateQueryToText(this.sinceText, "") +
            "~" +
            He.dateQueryToText(this.untilText, "")
          );
        },
        __class__: He,
      }),
      (t["albero.entity.SearchType"] = {
        __ename__: ["albero", "entity", "SearchType"],
        __constructs__: ["MESSAGES", "ATTACHMENTS"],
      })),
    Pr =
      ((Lr.MESSAGES = ["MESSAGES", 0]),
      (Lr.MESSAGES.toString = a),
      ((Lr.MESSAGES.__enum__ = Lr).ATTACHMENTS = ["ATTACHMENTS", 1]),
      (Lr.ATTACHMENTS.toString = a),
      ((Lr.ATTACHMENTS.__enum__ = Lr).__empty_constructs__ = [
        Lr.MESSAGES,
        Lr.ATTACHMENTS,
      ]),
      function () {}),
    Br =
      (((t["albero.entity.SearchTypeUtil"] = Pr).__name__ = [
        "albero",
        "entity",
        "SearchTypeUtil",
      ]),
      (Pr.toLocalizeText = function (e) {
        switch (e[1]) {
          case 0:
            return c.localize("SearchParams.text_type_messages");
          case 1:
            return c.localize("SearchParams.text_type_attachments");
        }
      }),
      (t["albero.entity.SearchDate"] = {
        __ename__: ["albero", "entity", "SearchDate"],
        __constructs__: [
          "DATE_NONE",
          "DATE_TODAY",
          "DATE_YESTERDAY",
          "DATE_THREE_DAYS_BEFORE",
          "DATE_CALENDAR",
        ],
      })),
    Hr =
      ((Br.DATE_NONE = ["DATE_NONE", 0]),
      (Br.DATE_NONE.toString = a),
      ((Br.DATE_NONE.__enum__ = Br).DATE_TODAY = ["DATE_TODAY", 1]),
      (Br.DATE_TODAY.toString = a),
      ((Br.DATE_TODAY.__enum__ = Br).DATE_YESTERDAY = ["DATE_YESTERDAY", 2]),
      (Br.DATE_YESTERDAY.toString = a),
      ((Br.DATE_YESTERDAY.__enum__ = Br).DATE_THREE_DAYS_BEFORE = [
        "DATE_THREE_DAYS_BEFORE",
        3,
      ]),
      (Br.DATE_THREE_DAYS_BEFORE.toString = a),
      ((Br.DATE_THREE_DAYS_BEFORE.__enum__ = Br).DATE_CALENDAR = [
        "DATE_CALENDAR",
        4,
      ]),
      (Br.DATE_CALENDAR.toString = a),
      ((Br.DATE_CALENDAR.__enum__ = Br).__empty_constructs__ = [
        Br.DATE_NONE,
        Br.DATE_TODAY,
        Br.DATE_YESTERDAY,
        Br.DATE_THREE_DAYS_BEFORE,
        Br.DATE_CALENDAR,
      ]),
      ((t["albero.entity.SearchDateUtil"] = Ge).__name__ = [
        "albero",
        "entity",
        "SearchDateUtil",
      ]),
      (Ge.createSearchDateArray = function (e) {
        return e
          ? [
              Br.DATE_TODAY,
              Br.DATE_YESTERDAY,
              Br.DATE_THREE_DAYS_BEFORE,
              Br.DATE_CALENDAR,
            ]
          : [Br.DATE_YESTERDAY, Br.DATE_THREE_DAYS_BEFORE, Br.DATE_CALENDAR];
      }),
      (Ge.toLocalizeText = function (e) {
        switch (e[1]) {
          case 0:
            return c.localize("SearchParams.text_date_none");
          case 1:
            return c.localize("SearchParams.text_date_today");
          case 2:
            return c.localize("SearchParams.text_date_yesterday");
          case 3:
            return c.localize("SearchParams.text_date_three_days_before");
          case 4:
            return c.localize("SearchParams.text_date_calendar");
        }
      }),
      ((t["albero.entity.Session"] = je).__name__ = [
        "albero",
        "entity",
        "Session",
      ]),
      (je.prototype = { __class__: je }),
      ((t["albero.entity.Solution"] = ze).__name__ = [
        "albero",
        "entity",
        "Solution",
      ]),
      (ze.getSolutionType = function (e) {
        switch (e) {
          case 1:
            return jr.PhotoViewTrial;
          case 2:
            return jr.PhotoView;
          case 3:
            return jr.MultiView;
          case 4:
            return jr.KeywordWatching;
          case 26:
            return jr.Flow;
          default:
            return jr.Unknown;
        }
      }),
      (ze.getSolutionIdFromType = function (e) {
        switch (e[1]) {
          case 0:
            return 1;
          case 1:
            return 2;
          case 2:
            return 3;
          case 3:
            return 4;
          case 4:
            return 26;
          default:
            return null;
        }
      }),
      (ze.prototype = { __class__: ze }),
      function (e) {
        null != e && ((this.iconUrl = e.icon_url), (this.url = e.url));
      }),
    Gr =
      (((t["albero.entity.SolutionLink"] = Hr).__name__ = [
        "albero",
        "entity",
        "SolutionLink",
      ]),
      (Hr.prototype = {
        getUrlWithPlaceholderParams: function (e, t) {
          return null == this.url
            ? null
            : I.replace(
                I.replace(this.url, "${domain_id}", Va.toString(e)),
                "${user_id}",
                Va.toString(t),
              );
        },
        __class__: Hr,
      }),
      function (e) {
        null != e && ((this.url = e.url), (this.loaded = !1));
      }),
    jr =
      (((t["albero.entity.SolutionPlugin"] = Gr).__name__ = [
        "albero",
        "entity",
        "SolutionPlugin",
      ]),
      (Gr.prototype = { __class__: Gr }),
      (t["albero.entity.SolutionType"] = {
        __ename__: ["albero", "entity", "SolutionType"],
        __constructs__: [
          "PhotoViewTrial",
          "PhotoView",
          "MultiView",
          "KeywordWatching",
          "Flow",
          "Unknown",
        ],
      })),
    zr =
      ((jr.PhotoViewTrial = ["PhotoViewTrial", 0]),
      (jr.PhotoViewTrial.toString = a),
      ((jr.PhotoViewTrial.__enum__ = jr).PhotoView = ["PhotoView", 1]),
      (jr.PhotoView.toString = a),
      ((jr.PhotoView.__enum__ = jr).MultiView = ["MultiView", 2]),
      (jr.MultiView.toString = a),
      ((jr.MultiView.__enum__ = jr).KeywordWatching = ["KeywordWatching", 3]),
      (jr.KeywordWatching.toString = a),
      ((jr.KeywordWatching.__enum__ = jr).Flow = ["Flow", 4]),
      (jr.Flow.toString = a),
      ((jr.Flow.__enum__ = jr).Unknown = ["Unknown", 5]),
      (jr.Unknown.toString = a),
      ((jr.Unknown.__enum__ = jr).__empty_constructs__ = [
        jr.PhotoViewTrial,
        jr.PhotoView,
        jr.MultiView,
        jr.KeywordWatching,
        jr.Flow,
        jr.Unknown,
      ]),
      function () {
        ((this.loadingSolutionIdsMap = new ts()),
          (this.solutionsMap = new ts()));
      }),
    Yr =
      (((t["albero.entity.SolutionStore"] = zr).__name__ = [
        "albero",
        "entity",
        "SolutionStore",
      ]),
      (zr.prototype = {
        setSolution: function (e) {
          this.solutionsMap.h[e.solutionId] = e;
        },
        getSolution: function (e) {
          return this.solutionsMap.h[e];
        },
        isLoadedSolutionId: function (e) {
          return this.solutionsMap.h.hasOwnProperty(e);
        },
        addLoadingSolutionId: function (e) {
          this.loadingSolutionIdsMap.h[e] = !0;
        },
        removeLoadingSolutionId: function (e) {
          this.loadingSolutionIdsMap.remove(e);
        },
        isLoadingSolutionId: function (e) {
          return this.loadingSolutionIdsMap.h.hasOwnProperty(e);
        },
        __class__: zr,
      }),
      ((t["albero.entity.Stamp"] = Ye).__name__ = [
        "albero",
        "entity",
        "Stamp",
      ]),
      (Ye.prototype = { __class__: Ye }),
      ((t["albero.entity.CommonStamp"] = We).__name__ = [
        "albero",
        "entity",
        "CommonStamp",
      ]),
      (We.__interfaces__ = [Ye]),
      (We.fromCommonStampId = function (e) {
        return new We(e);
      }),
      (We.fromJson = function (e) {
        return null == e.commonStampId
          ? null
          : We.fromCommonStampId(e.commonStampId);
      }),
      (We.prototype = {
        getImageUrl: function () {
          return "./assets/images/stamp/3/" + this.commonStampId + ".png";
        },
        getKey: function () {
          return "common_" + this.commonStampId;
        },
        toJson: function () {
          return { commonStampId: this.commonStampId };
        },
        createMessage: function (e, t) {
          var n = new O();
          return (
            (n.talkId = e),
            (n.type = d.stamp),
            (n.content = {
              stamp_set: 3,
              stamp_index: this.createStampIndex(),
            }),
            null != t && 0 != t.length && (n.content.text = t),
            n
          );
        },
        createStampIndex: function () {
          for (
            var e = _.make(0, 0),
              t = _.make(0, 10),
              n = 0,
              i = this.commonStampId.length;
            n < i;
          ) {
            var r = n++,
              o = 65535 & e.low,
              a = e.low >>> 16,
              s = 65535 & t.low,
              l = t.low >>> 16,
              u = w._mul(o, s),
              s = w._mul(a, s),
              o = w._mul(o, l),
              a = (((w._mul(a, l) + (o >>> 16)) | 0) + (s >>> 16)) | 0,
              a =
                ((o <<= 16),
                w.ucompare((l = (u + o) | 0), o) < 0 && (a++, (a |= 0)),
                (s <<= 16),
                w.ucompare((l = (l + s) | 0), s) < 0 && (a++, (a |= 0)),
                (a + ((w._mul(e.low, t.high) + w._mul(e.high, t.low)) | 0)) |
                  0),
              u = new E(a, l),
              o = _.make(0, P.parseInt(this.commonStampId.charAt(r))),
              s = (u.high + o.high) | 0,
              a = (u.low + o.low) | 0;
            w.ucompare(a, u.low) < 0 && (s++, (s |= 0));
            e = new E(s, a);
          }
          return e;
        },
        __class__: We,
      }),
      ((t["albero.entity.OriginalStamp"] = Ke).__name__ = [
        "albero",
        "entity",
        "OriginalStamp",
      ]),
      (Ke.__interfaces__ = [Ye]),
      (Ke.fromJson = function (e) {
        var t, n;
        return null == e.originalStamp ||
          null == (e = e.originalStamp).stampsetIdStr ||
          null == e.stampIdStr
          ? null
          : ((t = D.makeFromIdStr(e.stampsetIdStr)),
            (n = D.makeFromIdStr(e.stampIdStr)),
            new Ke(t, n, e.illust));
      }),
      (Ke.fromPropsAndStampsetId = function (e, t) {
        return new Ke(t, D.fromIntOrInt64(e.id), e.illust);
      }),
      (Ke.createKey = function (e) {
        return "original__" + e.high + "_" + e.low;
      }),
      (Ke.prototype = {
        getImageUrl: function () {
          return this.illust;
        },
        getKey: function () {
          return Ke.createKey(this.id);
        },
        toJson: function () {
          var e = this.stampsetId,
            t = this.id;
          return {
            originalStamp: {
              stampsetIdStr: "_" + e.high + "_" + e.low,
              stampIdStr: "_" + t.high + "_" + t.low,
              illust: this.illust,
            },
          };
        },
        createMessage: function (e, t) {
          var n = new O();
          return (
            (n.talkId = e),
            (n.type = d.originalStamp),
            (n.content = { stampset_id: this.stampsetId, stamp_id: this.id }),
            null != t && 0 != t.length && (n.content.text = t),
            n
          );
        },
        __class__: Ke,
      }),
      function () {
        ((this.stampsetStore = new B()),
          (this.stampStore = new B()),
          (this.stampKeyCategoryIdMap = new B()));
      }),
    Wr =
      (((t["albero.entity.StampStore"] = Yr).__name__ = [
        "albero",
        "entity",
        "StampStore",
      ]),
      (Yr.prototype = {
        addStampsets: function (e) {
          for (var t = 0; t < e.length; ) {
            var n = e[t];
            (++t, this.addStampset(n));
          }
        },
        addStampset: function (e, t) {
          null == t && (t = !0);
          for (
            var n = this.stampsetStore,
              i = e.toTabId(),
              r = (null != Y[i] ? n.setReserved(i, e) : (n.h[i] = e), 0),
              o = e.getStamps();
            r < o.length;
          ) {
            var a,
              s,
              l = o[r],
              u = (++r, l.getKey());
            (!(
              !t &&
              ((a = this.stampStore),
              null != Y[u] ? a.existsReserved(u) : a.h.hasOwnProperty(u))
            ) &&
              ((a = this.stampStore),
              null != Y[u] ? a.setReserved(u, l) : (a.h[u] = l)),
              t &&
                ((l = this.stampKeyCategoryIdMap),
                (s = e.getCategoryId()),
                (l = l),
                null != Y[u] ? l.setReserved(u, s) : (l.h[u] = s)));
          }
        },
        updateStampsetInfo: function (e) {
          var t = e.toTabId(),
            n = this.stampsetStore,
            n = null != Y[t] ? n.getReserved(t) : n.h[t];
          if (null != n && n.isOlderThan(e)) {
            this.stampsetStore.remove(t);
            for (var i = 0, r = n.getStamps(); i < r.length; ) {
              var o = r[i];
              (++i, this.stampStore.remove(o.getKey()));
            }
          }
        },
        removeStampsetInfo: function (e) {
          var e = Va.toString(e),
            t = this.stampsetStore,
            t = null != Y[e] ? t.getReserved(e) : t.h[e];
          if (null != t) {
            this.stampsetStore.remove(e);
            for (var n = 0, i = t.getStamps(); n < i.length; ) {
              var r = i[n];
              (++n, this.stampStore.remove(r.getKey()));
            }
          }
        },
        addStampsetInfos: function (e) {
          var i = this;
          k.iter(e, function (e) {
            var t = i.stampsetStore,
              n = e.toTabId();
            (null != Y[n] ? t.existsReserved(n) : t.h.hasOwnProperty(n)) ||
              ((t = Wr.fromStampsetInfo(e)), i.addStampset(t));
          });
        },
        getStampsetByStampsetId: function (e) {
          return this.getStampsetByTabId("_" + e.high + "_" + e.low);
        },
        getStampsetsByTabIds: function (e) {
          return e.map(z(this, this.getStampsetByTabId)).filter(function (e) {
            return null != e;
          });
        },
        getStampsetByTabId: function (e) {
          var t = this.stampsetStore;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        getStampByStampKey: function (e) {
          var t = this.stampStore;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        getOriginalStamp: function (e) {
          e = this.getStampByStampKey(Ke.createKey(e));
          return null == e ? null : G.__cast(e, Ke);
        },
        applyHistories: function (r, e) {
          var o = this,
            e = e
              .map(function (e) {
                var t = We.fromJson(e);
                if (null != t)
                  return (
                    (i = o.stampKeyCategoryIdMap),
                    (n = t.getKey()),
                    (i = i),
                    (i = null != Y[n] ? i.getReserved(n) : i.h[n]),
                    r.builtinStampSetting.isStampsetEnabled(i) ? t : null
                  );
                var n = Ke.fromJson(e);
                if (null != n) {
                  var i = G.__cast(n, Ke);
                  if (
                    null != r &&
                    r.stampsetSetting.containSendableStampset(i.stampsetId)
                  )
                    return n;
                }
                return (F.d("A history stamp is disabled.", e, r), null);
              })
              .filter(function (e) {
                return null != e;
              }),
            t = Wr.historyStampset();
          t.setStamps(e);
          this.addStampset(t, !1);
        },
        __class__: Yr,
      }),
      function (e, t, n, i) {
        ((this.stampsetType = e),
          (this.name = t),
          (this.icon = n),
          (this.version = i),
          (this.stamps = []));
      }),
    Kr =
      (((t["albero.entity.Stampset"] = Wr).__name__ = [
        "albero",
        "entity",
        "Stampset",
      ]),
      (Wr.historyStampset = function () {
        return new Wr(Vr.getHistoryStampsetType(), "stamp-history", null, 0);
      }),
      (Wr.commonStampset = function (e, t, n) {
        return new Wr(
          Vr.fromCommonStampCategoryIndex(e),
          t,
          "./assets/images/stamp/3/" + n,
          0,
        );
      }),
      (Wr.originalStampset = function (e) {
        var t = D.fromIntOrInt64(e.stampset_id),
          n = Vr.fromStampsetId(t),
          n = new Wr(n, e.name, e.icon, e.version);
        return (
          null != e.stamps &&
            ((e = G.__cast(e.stamps, Array).map(function (e) {
              return Ke.fromPropsAndStampsetId(e, t);
            })),
            n.setStamps(e)),
          n
        );
      }),
      (Wr.fromStampsetInfo = function (e) {
        var t = Vr.fromStampsetId(e.stampsetId);
        return new Wr(t, e.name, e.icon, e.version);
      }),
      (Wr.prototype = {
        getCategoryId: function () {
          return this.stampsetType.toCategoryId();
        },
        toTabId: function () {
          return this.stampsetType.toTabId();
        },
        getIconUrl: function () {
          return this.icon;
        },
        isStampHistory: function () {
          return "stamp-history" == this.name;
        },
        isOriginalStampset: function () {
          return this.stampsetType.isOriginalStampset();
        },
        isEmpty: function () {
          return 0 == this.stamps.length;
        },
        setStamps: function (e) {
          this.stamps = e;
        },
        getStamps: function () {
          return this.stamps;
        },
        isOlderThan: function (e) {
          return this.version < e.version;
        },
        __class__: Wr,
      }),
      ((t["albero.entity.StampsetInfo"] = Ve).__name__ = [
        "albero",
        "entity",
        "StampsetInfo",
      ]),
      (Ve.prototype = {
        toTabId: function () {
          var e = this.stampsetId;
          return "_" + e.high + "_" + e.low;
        },
        eqStampsetId: function (e) {
          var t = this.stampsetId,
            e = e.stampsetId;
          return null != t && null != e && t.high == e.high && t.low == e.low;
        },
        isNewerThan: function (e) {
          return this.version > e.version;
        },
        __class__: Ve,
      }),
      function (e) {
        ((this.stampsetInfos = []),
          (this.allowCreateMessageStampsetIds = []),
          (this.version = 0),
          null != e &&
            ((this.version = e.version),
            (this.allowCreateMessageStampsetIds =
              e.allow_create_message_stampset_ids.map(D.fromIntOrInt64)),
            (this.stampsetInfos = this.getStampsetInfos(e.stampset_infos))));
      }),
    s =
      (((t["albero.entity.StampsetSetting"] = Kr).__name__ = [
        "albero",
        "entity",
        "StampsetSetting",
      ]),
      (Kr.prototype = {
        getStampsetInfos: function (e) {
          return null == e
            ? []
            : e.map(function (e) {
                return new Ve(e);
              });
        },
        isForwardForbidden: function (e) {
          ((e = e.getContentsForOriginalStamp()),
            (e = null == e ? null : e.stampsetId));
          return null != e && !this.containSendableStampset(e);
        },
        containSendableStampset: function (t) {
          return (
            null != this.allowCreateMessageStampsetIds &&
            k.exists(this.allowCreateMessageStampsetIds, function (e) {
              return (
                null != t && null != e && t.high == e.high && t.low == e.low
              );
            })
          );
        },
        containVisibleStampset: function (t) {
          return k.exists(this.stampsetInfos, function (e) {
            e = e.stampsetId;
            return null != t && null != e && t.high == e.high && t.low == e.low;
          });
        },
        getSendableStampsetInfos: function () {
          var t = this;
          return null == this.stampsetInfos ||
            null == this.allowCreateMessageStampsetIds
            ? []
            : this.stampsetInfos.filter(function (e) {
                return t.containSendableStampset(e.stampsetId);
              });
        },
        isNewerThan: function (e) {
          return this.version > e.version;
        },
        updateStampsetInfo: function (t) {
          this.stampsetInfos = this.stampsetInfos.map(function (e) {
            return t.eqStampsetId(e) && t.isNewerThan(e) ? t : e;
          });
        },
        deleteStampsetInfo: function (t) {
          ((this.stampsetInfos = this.stampsetInfos.filter(function (e) {
            e = e.stampsetId;
            return !(
              null != e &&
              null != t &&
              e.high == t.high &&
              e.low == t.low
            );
          })),
            (this.allowCreateMessageStampsetIds =
              this.allowCreateMessageStampsetIds.filter(function (e) {
                return !(
                  null != e &&
                  null != t &&
                  e.high == t.high &&
                  e.low == t.low
                );
              })));
        },
        getStampsetIdsDiff: function (e) {
          for (
            var t = new B(),
              n = e.stampsetInfos.concat(this.stampsetInfos),
              i = 0;
            i < n.length;
          ) {
            var r = n[i],
              r = (++i, r.stampsetId),
              r = "_" + r.high + "_" + r.low,
              o =
                ((null != Y[r] ? t.existsReserved(r) : t.h.hasOwnProperty(r))
                  ? null != Y[r]
                    ? t.getReserved(r)
                    : t.h[r]
                  : 0) + 1;
            null != Y[r] ? t.setReserved(r, o) : (t.h[r] = o);
          }
          for (var a = [], s = t.keys(); s.hasNext(); ) {
            var l = s.next();
            1 == (null != Y[l] ? t.getReserved(l) : t.h[l]) &&
              a.push(D.makeFromIdStr(l));
          }
          return a;
        },
        isStampsetIdsEmpty: function () {
          return 0 == this.allowCreateMessageStampsetIds.length;
        },
        __class__: Kr,
      }),
      (t["albero.entity.StampsetTypeValue"] = {
        __ename__: ["albero", "entity", "StampsetTypeValue"],
        __constructs__: [
          "HISTORY",
          "COMMON_CIMMON",
          "COMMON_BUSINESS",
          "COMMON_BUSINESS_TEXT",
          "COMMON_LARGE_TEXT",
          "COMMON_EFFECT",
          "COMMON_EFFECT2",
          "COMMON_HANKO",
          "COMMON_YUZU",
          "COMMON_YUZU2",
          "COMMON_YUZU3",
          "ORIGINAL",
        ],
      })),
    Vr =
      ((s.HISTORY = ["HISTORY", 0]),
      (s.HISTORY.toString = a),
      ((s.HISTORY.__enum__ = s).COMMON_CIMMON = ["COMMON_CIMMON", 1]),
      (s.COMMON_CIMMON.toString = a),
      ((s.COMMON_CIMMON.__enum__ = s).COMMON_BUSINESS = ["COMMON_BUSINESS", 2]),
      (s.COMMON_BUSINESS.toString = a),
      ((s.COMMON_BUSINESS.__enum__ = s).COMMON_BUSINESS_TEXT = [
        "COMMON_BUSINESS_TEXT",
        3,
      ]),
      (s.COMMON_BUSINESS_TEXT.toString = a),
      ((s.COMMON_BUSINESS_TEXT.__enum__ = s).COMMON_LARGE_TEXT = [
        "COMMON_LARGE_TEXT",
        4,
      ]),
      (s.COMMON_LARGE_TEXT.toString = a),
      ((s.COMMON_LARGE_TEXT.__enum__ = s).COMMON_EFFECT = ["COMMON_EFFECT", 5]),
      (s.COMMON_EFFECT.toString = a),
      ((s.COMMON_EFFECT.__enum__ = s).COMMON_EFFECT2 = ["COMMON_EFFECT2", 6]),
      (s.COMMON_EFFECT2.toString = a),
      ((s.COMMON_EFFECT2.__enum__ = s).COMMON_HANKO = ["COMMON_HANKO", 7]),
      (s.COMMON_HANKO.toString = a),
      ((s.COMMON_HANKO.__enum__ = s).COMMON_YUZU = ["COMMON_YUZU", 8]),
      (s.COMMON_YUZU.toString = a),
      ((s.COMMON_YUZU.__enum__ = s).COMMON_YUZU2 = ["COMMON_YUZU2", 9]),
      (s.COMMON_YUZU2.toString = a),
      ((s.COMMON_YUZU2.__enum__ = s).COMMON_YUZU3 = ["COMMON_YUZU3", 10]),
      (s.COMMON_YUZU3.toString = a),
      ((s.COMMON_YUZU3.__enum__ = s).ORIGINAL = function (e) {
        e = ["ORIGINAL", 11, e];
        return ((e.__enum__ = s), (e.toString = a), e);
      }),
      (s.__empty_constructs__ = [
        s.HISTORY,
        s.COMMON_CIMMON,
        s.COMMON_BUSINESS,
        s.COMMON_BUSINESS_TEXT,
        s.COMMON_LARGE_TEXT,
        s.COMMON_EFFECT,
        s.COMMON_EFFECT2,
        s.COMMON_HANKO,
        s.COMMON_YUZU,
        s.COMMON_YUZU2,
        s.COMMON_YUZU3,
      ]),
      function (e) {
        this.value = e;
      }),
    qr =
      (((t["albero.entity.StampsetType"] = Vr).__name__ = [
        "albero",
        "entity",
        "StampsetType",
      ]),
      (Vr.fromStampsetId = function (e) {
        return new Vr(s.ORIGINAL(e));
      }),
      (Vr.fromCommonStampCategoryIndex = function (e) {
        switch (e) {
          case 0:
            return new Vr(s.COMMON_CIMMON);
          case 1:
            return new Vr(s.COMMON_BUSINESS);
          case 2:
            return new Vr(s.COMMON_BUSINESS_TEXT);
          case 3:
            return new Vr(s.COMMON_LARGE_TEXT);
          case 4:
            return new Vr(s.COMMON_EFFECT);
          case 5:
            return new Vr(s.COMMON_EFFECT2);
          case 6:
            return new Vr(s.COMMON_HANKO);
          case 7:
            return new Vr(s.COMMON_YUZU);
          case 8:
            return new Vr(s.COMMON_YUZU2);
          case 9:
            return new Vr(s.COMMON_YUZU3);
          default:
            return null;
        }
      }),
      (Vr.fromTabId = function (e) {
        var t = P.parseInt(e);
        if (null != t)
          if (null == t) {
            var n = Vr.fromCommonStampCategoryIndex(t - 1);
            if (null != n) return n;
          } else {
            if (0 == t) return Vr.getHistoryStampsetType();
            n = Vr.fromCommonStampCategoryIndex(t - 1);
            if (null != n) return n;
          }
        t = D.makeFromIdStr(e);
        if (null != t) return new Vr(s.ORIGINAL(t));
        try {
          return new Vr(x.createEnum(s, e));
        } catch (e) {
          return Vr.getDefaultStampsetType();
        }
      }),
      (Vr.getHistoryStampsetType = function () {
        return new Vr(s.HISTORY);
      }),
      (Vr.getDefaultStampsetType = function () {
        return new Vr(s.COMMON_CIMMON);
      }),
      (Vr.prototype = {
        toTabId: function () {
          var e = this.value;
          return 11 == e[1]
            ? "_" + (e = e[2]).high + "_" + e.low
            : P.string(this.value);
        },
        toCategoryId: function () {
          var e = this.value;
          switch (e[1]) {
            case 0:
              return null;
            case 1:
              return 1;
            case 2:
              return 2;
            case 3:
              return 3;
            case 4:
              return 4;
            case 5:
              return 5;
            case 6:
              return 6;
            case 7:
              return 7;
            case 8:
              return 8;
            case 9:
              return 9;
            case 10:
              return 10;
            case 11:
              return null;
          }
        },
        getStampsetId: function () {
          var e = this.value;
          return 11 == e[1] ? e[2] : null;
        },
        isOriginalStampset: function () {
          var e = this.value;
          return 11 == e[1] || null;
        },
        fallbackIfInvalid: function (e) {
          return this.isValid(e) ? this : Vr.getDefaultStampsetType();
        },
        isValid: function (e) {
          var t = this.value;
          return (
            11 != t[1] ||
            (null != e &&
              null != e.stampsetSetting &&
              e.stampsetSetting.containSendableStampset(t[2]))
          );
        },
        __class__: Vr,
      }),
      function (e) {
        null != e &&
          ((this.id = e.talk_id),
          (this.domainId = e.domain_id),
          (this.type = this.typeOf(e.type)),
          (this.name = e.talk_name),
          (this.iconUrl = e.icon_url),
          (this.userIds = e.user_ids),
          (this.updatedAt = e.updated_at),
          (this.leftUsers = this.getUsers(e.left_users)),
          (this.groupTalkSettings = this.getGroupTalkSetting(e.settings)));
      }),
    Qr =
      (((t["albero.entity.Talk"] = qr).__name__ = ["albero", "entity", "Talk"]),
      (qr.prototype = {
        typeOf: function (e) {
          switch (e) {
            case 1:
              return Xr.PairTalk;
            case 2:
              return Xr.GroupTalk;
            default:
              return Xr.Unknown;
          }
        },
        getUsers: function (e) {
          return null == e
            ? null
            : e.map(function (e) {
                return new ve(e);
              });
        },
        getGroupTalkSetting: function (e) {
          return null == e ? m.None : m.Some(new Qr(e));
        },
        isSystemMessageVisibled: function () {
          var e = h.map(this.groupTalkSettings, function (e) {
            return e.systemMessageVisible;
          });
          switch (e[1]) {
            case 0:
              return e[2];
            case 1:
              return !0;
          }
        },
        isAllowDisplayPastMessages: function () {
          var e = h.map(this.groupTalkSettings, function (e) {
            return e.allowDisplayPastMessages;
          });
          switch (e[1]) {
            case 0:
              return e[2];
            case 1:
              return !1;
          }
        },
        isMember: function (e) {
          for (var t = 0, n = this.userIds; t < n.length; ) {
            var i = n[t];
            if (
              (++t,
              null != e && null != i && e.high == i.high && e.low == i.low)
            )
              return !0;
          }
          return !1;
        },
        __class__: qr,
      }),
      function (e) {
        null != e &&
          ((this.allowDisplayPastMessages = !!e.allow_display_past_messages),
          (this.systemMessageVisible =
            null == e.system_message_visible || e.system_message_visible));
      }),
    Xr =
      (((t["albero.entity.GroupTalkSettings"] = Qr).__name__ = [
        "albero",
        "entity",
        "GroupTalkSettings",
      ]),
      (Qr.createDefault = function () {
        var e = new Qr();
        return (
          (e.allowDisplayPastMessages = !0),
          (e.systemMessageVisible = !0),
          e
        );
      }),
      (Qr.prototype = { __class__: Qr }),
      (t["albero.entity.TalkType"] = {
        __ename__: ["albero", "entity", "TalkType"],
        __constructs__: ["Unknown", "PairTalk", "GroupTalk"],
      })),
    Jr =
      ((Xr.Unknown = ["Unknown", 0]),
      (Xr.Unknown.toString = a),
      ((Xr.Unknown.__enum__ = Xr).PairTalk = ["PairTalk", 1]),
      (Xr.PairTalk.toString = a),
      ((Xr.PairTalk.__enum__ = Xr).GroupTalk = ["GroupTalk", 2]),
      (Xr.GroupTalk.toString = a),
      ((Xr.GroupTalk.__enum__ = Xr).__empty_constructs__ = [
        Xr.Unknown,
        Xr.PairTalk,
        Xr.GroupTalk,
      ]),
      (t["albero.entity.TalkOrAnnouncements"] = {
        __ename__: ["albero", "entity", "TalkOrAnnouncements"],
        __constructs__: ["talk", "announcements"],
      })),
    Zr =
      ((Jr.talk = function (e) {
        e = ["talk", 0, e];
        return ((e.__enum__ = Jr), (e.toString = a), e);
      }),
      (Jr.announcements = ["announcements", 1]),
      (Jr.announcements.toString = a),
      ((Jr.announcements.__enum__ = Jr).__empty_constructs__ = [
        Jr.announcements,
      ]),
      ((t["albero.entity.TalkOrAnnouncementsHelper"] = qe).__name__ = [
        "albero",
        "entity",
        "TalkOrAnnouncementsHelper",
      ]),
      (qe.getTalkId = function (e) {
        switch (e[1]) {
          case 0:
            return m.Some(e[2]);
          case 1:
            return m.None;
        }
      }),
      ((t["albero.entity.TalkPack"] = Qe).__name__ = [
        "albero",
        "entity",
        "TalkPack",
      ]),
      (Qe.compareTalkOrderingTimestamp = function (e, t) {
        var e = e.getTalkOrderingTimestamp(),
          t = t.getTalkOrderingTimestamp(),
          n = (e.high - t.high) | 0;
        return (
          0 == n && (n = w.ucompare(e.low, t.low)),
          e.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1
        );
      }),
      (Qe.prototype = {
        getTalkOrderingTimestamp: function () {
          return null != this.status &&
            null != this.status.talkOrderingTimestamp
            ? this.status.talkOrderingTimestamp
            : this.talk.updatedAt;
        },
        getTalkOrderInFavorites: function () {
          return null != this.status ? this.status.orderInFavorites : null;
        },
        __class__: Qe,
      }),
      function (e, t) {
        ((this.paneType = e), (this.talkSelection = t));
      }),
    $r =
      (((t["albero.entity.TalkPaneSelection"] = Zr).__name__ = [
        "albero",
        "entity",
        "TalkPaneSelection",
      ]),
      (Zr.createArrayForMultiPanes = function (e) {
        return k.mapi(e, function (e, t) {
          return new Zr(Ue.getMultiPaneType(e), t);
        });
      }),
      (Zr.prototype = {
        createUrls: function (e, t) {
          switch (this.paneType[1]) {
            case 0:
              return g.talks(e, Pn.single(this.talkSelection));
            case 1:
              var n = z(t, t.getTalkSelection),
                i = e,
                r = Ue.getMultiPaneTypes().map(function (e) {
                  return n(i, e);
                });
              return ((r[0] = this.talkSelection), g.talks(e, Pn.multi(r)));
            case 2:
              var o = z(t, t.getTalkSelection),
                a = e,
                r = Ue.getMultiPaneTypes().map(function (e) {
                  return o(a, e);
                });
              return ((r[1] = this.talkSelection), g.talks(e, Pn.multi(r)));
            case 3:
              var s = z(t, t.getTalkSelection),
                l = e,
                r = Ue.getMultiPaneTypes().map(function (e) {
                  return s(l, e);
                });
              return ((r[2] = this.talkSelection), g.talks(e, Pn.multi(r)));
          }
        },
        getTalkIdOrNull: function () {
          return eo.getTalkIdOrNull(this.talkSelection);
        },
        __class__: Zr,
      }),
      ((t["albero.entity.ITalkSelectionGetter"] = Xe).__name__ = [
        "albero",
        "entity",
        "ITalkSelectionGetter",
      ]),
      (Xe.prototype = { __class__: Xe }),
      ((t["albero.entity.TalkPushNotificationEvent"] = Je).__name__ = [
        "albero",
        "entity",
        "TalkPushNotificationEvent",
      ]),
      (Je.prototype = { __class__: Je }),
      (t["albero.entity.TalkSelection"] = {
        __ename__: ["albero", "entity", "TalkSelection"],
        __constructs__: [
          "NotSelected",
          "TalkSelected",
          "AnnouncementsSelected",
        ],
      })),
    eo =
      (($r.NotSelected = ["NotSelected", 0]),
      ($r.NotSelected.toString = a),
      (($r.NotSelected.__enum__ = $r).TalkSelected = function (e, t) {
        e = ["TalkSelected", 1, e, t];
        return ((e.__enum__ = $r), (e.toString = a), e);
      }),
      ($r.AnnouncementsSelected = ["AnnouncementsSelected", 2]),
      ($r.AnnouncementsSelected.toString = a),
      (($r.AnnouncementsSelected.__enum__ = $r).__empty_constructs__ = [
        $r.NotSelected,
        $r.AnnouncementsSelected,
      ]),
      function () {}),
    to =
      (((t["albero.entity.TalkSelectionHelper"] = eo).__name__ = [
        "albero",
        "entity",
        "TalkSelectionHelper",
      ]),
      (eo.eqTalkSelection = function (e, t) {
        var n, i;
        return (
          !!(
            e[0] == t[0] &&
            ((n = eo.getTalkIdOrNull(e)),
            (i = eo.getTalkIdOrNull(t)),
            (null == n && null == i) ||
              (null != n && null != i && n.high == i.high && n.low == i.low))
          ) &&
          ((n = eo.getMessageIdOrNull(e)),
          (i = eo.getMessageIdOrNull(t)),
          (null == n && null == i) ||
            (null != n && null != i && n.high == i.high && n.low == i.low))
        );
      }),
      (eo.isSelected = function (e, t) {
        switch (e[1]) {
          case 0:
            return !1;
          case 1:
            var n = e[2],
              i = h.orNull(qe.getTalkId(t));
            return null != n && null != i && n.high == i.high && n.low == i.low;
          case 2:
            return t == Jr.announcements;
        }
      }),
      (eo.getTalkIdOrNull = function (e) {
        switch (e[1]) {
          case 0:
            return null;
          case 1:
            return e[2];
          case 2:
            return null;
        }
      }),
      (eo.getMessageIdOrNull = function (e) {
        switch (e[1]) {
          case 0:
            return null;
          case 1:
            return h.orNull(e[3]);
          case 2:
            return null;
        }
      }),
      (eo.toTalkSelectionStr = function (e) {
        switch (e[1]) {
          case 0:
            return "";
          case 1:
            var t = e[2];
            return "_" + t.high + "_" + t.low;
          case 2:
            return "announcement";
        }
      }),
      ((t["albero.entity.TalkStatus"] = Ze).__name__ = [
        "albero",
        "entity",
        "TalkStatus",
      ]),
      (Ze.prototype = {
        calcUnreadMentionCount: function () {
          var e = this.unreadMentionCount,
            t = this.maxReadMessageId;
          if (null != this.readingMessages) {
            if (this.readingMessages.isReadingAll()) return 0;
            t = this.readingMessages.getUnreadMentionMessages(t);
            ((e -= t.length), 0 < t.length && t[t.length - 1].messageId);
          }
          return e;
        },
        calcUnreadCount: function () {
          var e = this.unreadCount,
            t = this.maxReadMessageId;
          return null != this.readingMessages &&
            (this.readingMessages.isReadingAll() ||
              (e = e - this.readingMessages.getUnreadMessages(t).length) < 0)
            ? 0
            : e;
        },
        calcMaxReadMessageId: function () {
          var e, t, n;
          return null == this.readingMessages
            ? this.maxReadMessageId
            : this.readingMessages.isReadingAll()
              ? this.maxMessageId
              : ((e = this.readingMessages.getLastMesssgeId()),
                null != this.maxReadMessageId &&
                ((t = this.maxReadMessageId),
                0 == (n = (e.high - t.high) | 0) &&
                  (n = w.ucompare(e.low, t.low)),
                (e.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1) < 0)
                  ? this.maxReadMessageId
                  : e);
        },
        update: function (e) {
          var t = this,
            n = !1;
          return (
            this.needIdUpdate(
              this.maxEveryoneReadMessageId,
              e.maxEveryoneReadMessageId,
            ) &&
              ((this.maxEveryoneReadMessageId = e.maxEveryoneReadMessageId),
              (n = !0)),
            h.foreach(e.maxReadMessageId, function (e) {
              t.needIdUpdate(t.maxReadMessageId, e) &&
                ((t.maxReadMessageId = e), (n = !0));
            }),
            n
          );
        },
        updateByMessageDeletion: function (e) {
          var t = !1;
          return (
            this.isUnread(e.messageId) &&
              (this.decrementUnreadCount(1),
              e.isMention && this.decrementUnreadMentionCount(1),
              (t = !0)),
            this.isMax(e.messageId) &&
              (null != this.maxMessage && (this.maxMessage.type = d.deleted),
              (t = !0)),
            t
          );
        },
        updateByMessageReadStatusesUpdate: function (e) {
          var t = e.messageIds,
            n = e.mentionMessageIds,
            i = t.filter(z(this, this.isUnread));
          if (0 == i.length) return !1;
          for (var r = i[0], o = 0; o < i.length; ) {
            var a = i[o],
              s = (++o, (r.high - a.high) | 0);
            (0 == s && (s = w.ucompare(r.low, a.low)),
              (r.high < 0 ? (a.high < 0 ? s : -1) : 0 <= a.high ? s : 1) < 0 &&
                (r = a),
              D.notContains(e.messageIdsExcludingUnreadCountTargets, a) &&
                (this.decrementUnreadCount(1), D.contains(n, a)) &&
                this.decrementUnreadMentionCount(1));
          }
          return ((this.maxReadMessageId = r), !0);
        },
        needIdUpdate: function (e, t) {
          var n;
          return (
            null != t &&
            (null == e ||
              (0 == (n = (e.high - t.high) | 0) &&
                (n = w.ucompare(e.low, t.low)),
              (e.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1) < 0))
          );
        },
        read: function (e, t, n) {
          ((this.maxReadMessageId = e),
            this.decrementUnreadCount(t),
            this.decrementUnreadMentionCount(n));
        },
        decrementUnreadCount: function (e) {
          ((this.unreadCount -= e),
            this.unreadCount < 0 && (this.unreadCount = 0));
        },
        decrementUnreadMentionCount: function (e) {
          ((this.unreadMentionCount -= e),
            this.unreadMentionCount < 0 && (this.unreadMentionCount = 0));
        },
        readAll: function () {
          ((this.maxReadMessageId = this.maxMessageId),
            (this.unreadCount = 0),
            (this.unreadMentionCount = 0));
        },
        updateByMessage: function (e, t) {
          if (
            (this.isNewMax(e.id) &&
              ((this.maxMessageId = e.id), (this.maxMessage = e)),
            e.isUnreadCountTarget)
          ) {
            if (((this.talkOrderingTimestamp = e.createdAt), t))
              return this.isUnread(e.id) && 0 == this.calcUnreadCount()
                ? null != this.readingMessages
                  ? ((t = new Le(null, e.userId, e.id)),
                    void this.readingMessages.update(t))
                  : void this.readAll()
                : void 0;
            (this.unreadCount++, e.mention && this.unreadMentionCount++);
          }
        },
        updateByReadingMessagesStarted: function (e, t) {
          if (null == this.maxMessageId) return !1;
          e = new Le(e, t, this.maxMessageId);
          if (e.canRegardTalkReadCountZero(this.maxMessageId)) {
            if (!this.hasUnreadMessage()) return !1;
          } else if (!this.isUnread(e.getLastMesssgeId())) return !1;
          return null != this.readingMessages
            ? this.readingMessages.update(e)
            : ((this.readingMessages = e), !0);
        },
        updateByReadingMessagesFinished: function (e) {
          var t, n;
          null == this.readingMessages ||
            null == (t = this.readingMessages.getLastMesssgeId()) ||
            (0 == (n = (t.high - e.high) | 0) && (n = w.ucompare(t.low, e.low)),
            0 < (t.high < 0 ? (e.high < 0 ? n : -1) : 0 <= e.high ? n : 1)) ||
            (null != e &&
              null != this.maxReadMessageId &&
              (0 == (n = ((t = this.maxReadMessageId).high - e.high) | 0) &&
                (n = w.ucompare(t.low, e.low)),
              0 <=
                (t.high < 0 ? (e.high < 0 ? n : -1) : 0 <= e.high ? n : 1))) ||
            (this.readingMessages.canRegardTalkReadCountZero(this.maxMessageId)
              ? this.readAll()
              : ((t = this.readingMessages.getLastMesssgeId()),
                (n = (e = this.readingMessages.getUnreadMessages(
                  this.maxReadMessageId,
                )).filter(function (e) {
                  return e.hasMention;
                })),
                this.read(t, e.length, n.length)),
            (this.readingMessages = null));
        },
        updateByReadingMessagesCanceled: function (e) {
          var t, n;
          null == this.readingMessages ||
            null == (t = this.readingMessages.getLastMesssgeId()) ||
            (0 == (n = (t.high - e.high) | 0) && (n = w.ucompare(t.low, e.low)),
            0 < (t.high < 0 ? (e.high < 0 ? n : -1) : 0 <= e.high ? n : 1)) ||
            (this.readingMessages = null);
        },
        updateByTalkPushNotificationEvent: function (e) {
          return (
            e.version > this.pushNotificationSetting.version &&
            ((this.pushNotificationSetting = Ar.createWithParams(
              e.type,
              e.version,
            )),
            !0)
          );
        },
        isMessageAcceptable: function (e, t) {
          var n;
          return (
            !!this.isUnread(e.id) ||
            (0 ==
              (n =
                ((e = Me.subTimestampPart(this.maxReadMessageId, e.id)).high -
                  t.high) |
                0) && (n = w.ucompare(e.low, t.low)),
            (e.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1) <= 0)
          );
        },
        isUnread: function (e) {
          var t, n;
          return (
            null == this.maxReadMessageId ||
            (0 == (n = ((t = this.maxReadMessageId).high - e.high) | 0) &&
              (n = w.ucompare(t.low, e.low)),
            (t.high < 0 ? (e.high < 0 ? n : -1) : 0 <= e.high ? n : 1) < 0)
          );
        },
        isMax: function (e) {
          var t;
          return (
            null != this.maxMessageId &&
            null != (t = this.maxMessageId) &&
            null != e &&
            t.high == e.high &&
            t.low == e.low
          );
        },
        isNewMax: function (e) {
          var t, n;
          return (
            null == this.maxMessageId ||
            (0 == (n = ((t = this.maxMessageId).high - e.high) | 0) &&
              (n = w.ucompare(t.low, e.low)),
            (t.high < 0 ? (e.high < 0 ? n : -1) : 0 <= e.high ? n : 1) < 0)
          );
        },
        isNotificationDisabled: function () {
          return this.pushNotificationSetting.isDisabled();
        },
        getNotificationType: function () {
          return this.pushNotificationSetting.type;
        },
        hasUnreadMessage: function () {
          return null != this.maxMessage && this.isUnread(this.maxMessage.id);
        },
        clearReadingMessages: function () {
          this.readingMessages = null;
        },
        __class__: Ze,
      }),
      ((t["albero.entity.TalkStatusHelper"] = $e).__name__ = [
        "albero",
        "entity",
        "TalkStatusHelper",
      ]),
      ($e.isNotificationDisabled = function (e) {
        return null != e && e.isNotificationDisabled();
      }),
      ($e.isFavoriteTalk = function (e) {
        return null != e && null != e.orderInFavorites;
      }),
      ($e.isUnreadMesasgeExisted = function (e) {
        return null != e && e.hasUnreadMessage();
      }),
      ((t["albero.entity.TalkStatusUpdate"] = et).__name__ = [
        "albero",
        "entity",
        "TalkStatusUpdate",
      ]),
      (et.prototype = { __class__: et }),
      ((t["albero.entity.ThumbInfo"] = tt).__name__ = [
        "albero",
        "entity",
        "ThumbInfo",
      ]),
      (tt.prototype = {
        setAuth: function (e) {
          this.auth = e;
        },
        __class__: tt,
      }),
      function (e) {
        null != e && ((this.width = e.width), (this.height = e.height));
      }),
    u =
      (((t["albero.entity.ThumbDimension"] = to).__name__ = [
        "albero",
        "entity",
        "ThumbDimension",
      ]),
      (to.createWithSize = function (e, t) {
        var n = new to();
        return ((n.width = e), (n.height = t), n);
      }),
      (to.prototype = {
        toObject: function () {
          return { width: this.width, height: this.height };
        },
        __class__: to,
      }),
      ((t["albero.entity.UpdateDomainStampSettingEvent"] = nt).__name__ = [
        "albero",
        "entity",
        "UpdateDomainStampSettingEvent",
      ]),
      (nt.prototype = {
        isValid: function () {
          return this.newSetting.isNewerThan(this.oldSetting);
        },
        getStampsetIdsDiff: function () {
          return this.oldSetting.getStampsetIdsDiff(this.newSetting);
        },
        __class__: nt,
      }),
      (t["albero.entity.UploadUseType"] = {
        __ename__: ["albero", "entity", "UploadUseType"],
        __constructs__: [
          "PROFILE_IMAGE",
          "MESSAGE",
          "TALK_ICON",
          "THUMBNAIL",
          "NOTE_ATTACHMENT",
          "NOTE_THUMBNAIL",
          "SCHEDULE_MESSAGE",
        ],
      })),
    E =
      ((u.PROFILE_IMAGE = ["PROFILE_IMAGE", 0]),
      (u.PROFILE_IMAGE.toString = a),
      ((u.PROFILE_IMAGE.__enum__ = u).MESSAGE = ["MESSAGE", 1]),
      (u.MESSAGE.toString = a),
      ((u.MESSAGE.__enum__ = u).TALK_ICON = ["TALK_ICON", 2]),
      (u.TALK_ICON.toString = a),
      ((u.TALK_ICON.__enum__ = u).THUMBNAIL = ["THUMBNAIL", 3]),
      (u.THUMBNAIL.toString = a),
      ((u.THUMBNAIL.__enum__ = u).NOTE_ATTACHMENT = ["NOTE_ATTACHMENT", 4]),
      (u.NOTE_ATTACHMENT.toString = a),
      ((u.NOTE_ATTACHMENT.__enum__ = u).NOTE_THUMBNAIL = ["NOTE_THUMBNAIL", 5]),
      (u.NOTE_THUMBNAIL.toString = a),
      ((u.NOTE_THUMBNAIL.__enum__ = u).SCHEDULE_MESSAGE = [
        "SCHEDULE_MESSAGE",
        6,
      ]),
      (u.SCHEDULE_MESSAGE.toString = a),
      ((u.SCHEDULE_MESSAGE.__enum__ = u).__empty_constructs__ = [
        u.PROFILE_IMAGE,
        u.MESSAGE,
        u.TALK_ICON,
        u.THUMBNAIL,
        u.NOTE_ATTACHMENT,
        u.NOTE_THUMBNAIL,
        u.SCHEDULE_MESSAGE,
      ]),
      ((t["albero.entity.UploadUseTypeHelper"] = it).__name__ = [
        "albero",
        "entity",
        "UploadUseTypeHelper",
      ]),
      (it.getUseTypeInt = function (e) {
        switch (e[1]) {
          case 0:
            return 0;
          case 1:
            return 1;
          case 2:
            return 2;
          case 3:
            return 4;
          case 4:
            return 5;
          case 5:
            return 6;
          case 6:
            return 1;
        }
      }),
      ((t["albero.entity.UserIdentifier"] = rt).__name__ = [
        "albero",
        "entity",
        "UserIdentifier",
      ]),
      (rt.prototype = { __class__: rt }),
      function (e, t) {
        ((this.high = e), (this.low = t));
      }),
    no =
      (((t["haxe._Int64.___Int64"] = E).__name__ = [
        "haxe",
        "_Int64",
        "___Int64",
      ]),
      (E.prototype = { __class__: E }),
      ((t["albero.entity.UserPresence"] = ot).__name__ = [
        "albero",
        "entity",
        "UserPresence",
      ]),
      (ot.createUserPresenceState = function (e, t) {
        var n = (t.high - e.high) | 0,
          i = (t.low - e.low) | 0;
        w.ucompare(t.low, e.low) < 0 && (n--, (n |= 0));
        var t = new E(n, i),
          e = ot.SIXTY_MINUTES,
          n = (t.high - e.high) | 0;
        return (
          0 == n && (n = w.ucompare(t.low, e.low)),
          0 < (t.high < 0 ? (e.high < 0 ? n : -1) : 0 <= e.high ? n : 1)
            ? no.GT_60
            : (0 == (e = (t.high - (i = ot.FIFTEEN_MINUTES).high) | 0) &&
                (e = w.ucompare(t.low, i.low)),
              0 < (t.high < 0 ? (i.high < 0 ? e : -1) : 0 <= i.high ? e : 1)
                ? no.LT_OR_EQ_60
                : no.LT_OR_EQ_15)
        );
      }),
      (ot.prototype = {
        toState: function (e) {
          var t = e,
            n = h.map(this.lastUsedAtOpt, function (e) {
              return ot.createUserPresenceState(e, t);
            });
          switch (n[1]) {
            case 0:
              return n[2];
            case 1:
              return no.GT_60;
          }
        },
        isExpired: function (e, t) {
          var n = this.lastUpdatedAt,
            i = (e.high - n.high) | 0,
            r = (e.low - n.low) | 0;
          w.ucompare(e.low, n.low) < 0 && (i--, (i |= 0));
          ((e = new E(i, r)), (n = (e.high - t.high) | 0));
          return (
            0 == n && (n = w.ucompare(e.low, t.low)),
            0 < (e.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1)
          );
        },
        __class__: ot,
      }),
      (t["albero.entity.UserPresenceState"] = {
        __ename__: ["albero", "entity", "UserPresenceState"],
        __constructs__: ["LT_OR_EQ_15", "LT_OR_EQ_60", "GT_60"],
      })),
    io =
      ((no.LT_OR_EQ_15 = ["LT_OR_EQ_15", 0]),
      (no.LT_OR_EQ_15.toString = a),
      ((no.LT_OR_EQ_15.__enum__ = no).LT_OR_EQ_60 = ["LT_OR_EQ_60", 1]),
      (no.LT_OR_EQ_60.toString = a),
      ((no.LT_OR_EQ_60.__enum__ = no).GT_60 = ["GT_60", 2]),
      (no.GT_60.toString = a),
      ((no.GT_60.__enum__ = no).__empty_constructs__ = [
        no.LT_OR_EQ_15,
        no.LT_OR_EQ_60,
        no.GT_60,
      ]),
      ((t["albero.entity.api.message.AddFavoriteMessageResult"] = at).__name__ =
        ["albero", "entity", "api", "message", "AddFavoriteMessageResult"]),
      (at.prototype = { __class__: at }),
      ((t["albero.entity.api.message.DeleteFavoriteMessageResult"] =
        st).__name__ = [
        "albero",
        "entity",
        "api",
        "message",
        "DeleteFavoriteMessageResult",
      ]),
      (st.prototype = { __class__: st }),
      ((t["albero.entity.api.message.GetFavoriteMessagesResultWithParams"] =
        lt).__name__ = [
        "albero",
        "entity",
        "api",
        "message",
        "GetFavoriteMessagesResultWithParams",
      ]),
      (lt.prototype = {
        getNextMarker: function () {
          return this.nextMarker;
        },
        getMessages: function () {
          return this.messages;
        },
        __class__: lt,
      }),
      ((t["albero.entity.api.message.GetScheduledMessagesResultWithParams"] =
        ut).__name__ = [
        "albero",
        "entity",
        "api",
        "message",
        "GetScheduledMessagesResultWithParams",
      ]),
      (ut.prototype = { __class__: ut }),
      ((t["albero.entity.api.message.SearchMessagesAroundDateTimeResult"] =
        ct).__name__ = [
        "albero",
        "entity",
        "api",
        "message",
        "SearchMessagesAroundDateTimeResult",
      ]),
      (ct.prototype = { __class__: ct }),
      function (e) {
        ((this.messageId = e.message_id), (this.createdAt = e.created_at));
      }),
    ro =
      (((t["albero.entity.api.message.MessageCreatedAt"] = io).__name__ = [
        "albero",
        "entity",
        "api",
        "message",
        "MessageCreatedAt",
      ]),
      (io.prototype = { __class__: io }),
      function (e) {
        this.noteRaw = e;
      }),
    oo =
      (((t["albero.entity.api.note.CompressedNote"] = ro).__name__ = [
        "albero",
        "entity",
        "api",
        "note",
        "CompressedNote",
      ]),
      (ro.prototype = {
        uncompress: function () {
          var n = this,
            e = this.getCompressedContentRaw();
          return new Z().run(e).then(function (e) {
            var t = L.copy(n.noteRaw.note_revision),
              e = ((t.content = e), L.copy(n.noteRaw));
            return ((e.note_revision = t), e);
          });
        },
        getCompressedContentRaw: function () {
          return null == this.noteRaw || null == this.noteRaw.note_revision
            ? null
            : this.noteRaw.note_revision.content;
        },
        __class__: ro,
      }),
      ((t["albero.entity.api.note.CreateNoteError"] = _t).__name__ = [
        "albero",
        "entity",
        "api",
        "note",
        "CreateNoteError",
      ]),
      (_t.prototype = { __class__: _t }),
      ((t["albero.entity.api.note.CreateNoteResult"] = ht).__name__ = [
        "albero",
        "entity",
        "api",
        "note",
        "CreateNoteResult",
      ]),
      (ht.prototype = {
        updateWithEmitterKey: function (e) {
          var t = new ht();
          return ((t.note = this.note), (t.emitterKey = e), t);
        },
        __class__: ht,
      }),
      ((t["albero.entity.api.note.CreateNoteResultCompressed"] = dt).__name__ =
        ["albero", "entity", "api", "note", "CreateNoteResultCompressed"]),
      (dt.prototype = {
        uncompress: function () {
          return this.compressedNote.uncompress().then(function (e) {
            return new ht(e);
          });
        },
        __class__: dt,
      }),
      ((t["albero.entity.api.note.DeleteNoteResult"] = ft).__name__ = [
        "albero",
        "entity",
        "api",
        "note",
        "DeleteNoteResult",
      ]),
      (ft.prototype = { __class__: ft }),
      ((t["albero.entity.api.note.GetNoteResult"] = mt).__name__ = [
        "albero",
        "entity",
        "api",
        "note",
        "GetNoteResult",
      ]),
      (mt.prototype = {
        getNoteId: function () {
          return this.note.noteId;
        },
        getNoteCreateUserId: function () {
          return this.note.createdBy;
        },
        getRevisionCreateUserId: function () {
          return this.note.getRevisionCreateUserId();
        },
        getNoteCreateDate: function () {
          return this.note.createdAt;
        },
        getRevisionCreateDate: function () {
          return this.note.getRevisionCreateDate();
        },
        getLockedUserIdWithoutSelf: function (e) {
          return this.note.getLockedUserIdWithoutSelf(e);
        },
        __class__: mt,
      }),
      ((t["albero.entity.api.note.GetNoteResultCompressed"] = pt).__name__ = [
        "albero",
        "entity",
        "api",
        "note",
        "GetNoteResultCompressed",
      ]),
      (pt.prototype = {
        uncompress: function () {
          return this.compressedNote.uncompress().then(function (e) {
            return new mt(e);
          });
        },
        __class__: pt,
      }),
      ((t["albero.entity.api.note.GetNoteStatusesResult"] = gt).__name__ = [
        "albero",
        "entity",
        "api",
        "note",
        "GetNoteStatusesResult",
      ]),
      (gt.prototype = {
        isNotEmptyNoteStatuses: function () {
          return 0 < this.noteStatuses.length;
        },
        __class__: gt,
      }),
      ((t["albero.entity.api.note.UpdateNoteResult"] = yt).__name__ = [
        "albero",
        "entity",
        "api",
        "note",
        "UpdateNoteResult",
      ]),
      (yt.prototype = { __class__: yt }),
      ((t["albero.entity.api.note.UpdateNoteResultCompressed"] = St).__name__ =
        ["albero", "entity", "api", "note", "UpdateNoteResultCompressed"]),
      (St.prototype = {
        uncompress: function () {
          return this.compressedNote.uncompress().then(function (e) {
            return new yt(e);
          });
        },
        __class__: St,
      }),
      ((t["albero.entity.api.note.UpdateNoteSettingResult"] = vt).__name__ = [
        "albero",
        "entity",
        "api",
        "note",
        "UpdateNoteSettingResult",
      ]),
      (vt.prototype = { __class__: vt }),
      ((t["albero.entity.api.note.UpdateNoteSettingResultCompressed"] =
        Tt).__name__ = [
        "albero",
        "entity",
        "api",
        "note",
        "UpdateNoteSettingResultCompressed",
      ]),
      (Tt.prototype = {
        uncompress: function () {
          return this.compressedNote.uncompress().then(function (e) {
            return new vt(e);
          });
        },
        __class__: Tt,
      }),
      ((t["albero.entity.file.DummyFileUtil"] = It).__name__ = [
        "albero",
        "entity",
        "file",
        "DummyFileUtil",
      ]),
      (It.getBlobFromUpdatableFile = function (e) {
        return null == e.blob ? e : e.blob;
      }),
      ((t["albero.entity.file.FilePreview"] = Et).__name__ = [
        "albero",
        "entity",
        "file",
        "FilePreview",
      ]),
      (Et.prototype = { __class__: Et }),
      ((t["albero.entity.file.StageType"] = wt).__name__ = [
        "albero",
        "entity",
        "file",
        "StageType",
      ]),
      (wt.prototype = { __class__: wt }),
      ((t["albero.entity.file.StageTypeTalk"] = Nt).__name__ = [
        "albero",
        "entity",
        "file",
        "StageTypeTalk",
      ]),
      (Nt.__interfaces__ = [wt]),
      (Nt.prototype = {
        toString: function () {
          var e = this.talkId;
          return (
            "staged__" + e.high + "_" + e.low + Ue.toPanePrefix(this.paneType)
          );
        },
        isForPane: function (e) {
          return this.paneType == e;
        },
        isForNote: function () {
          return !1;
        },
        isForPreview: function () {
          return !1;
        },
        __class__: Nt,
      }),
      ((t["albero.entity.file.StageTypeLegacyNote"] = At).__name__ = [
        "albero",
        "entity",
        "file",
        "StageTypeLegacyNote",
      ]),
      (At.__interfaces__ = [wt]),
      (At.prototype = {
        toString: function () {
          return "staged_note_legacy_";
        },
        isForPane: function (e) {
          return !1;
        },
        isForNote: function () {
          return !0;
        },
        isForPreview: function () {
          return !1;
        },
        __class__: At,
      }),
      ((t["albero.entity.file.StageTypeNote"] = bt).__name__ = [
        "albero",
        "entity",
        "file",
        "StageTypeNote",
      ]),
      (bt.__interfaces__ = [wt]),
      (bt.prototype = {
        toString: function () {
          return "staged_note_";
        },
        isForPane: function (e) {
          return !1;
        },
        isForNote: function () {
          return !0;
        },
        isForPreview: function () {
          return !1;
        },
        __class__: bt,
      }),
      ((t["albero.entity.file.UploadableFilePackConverter"] = Mt).__name__ = [
        "albero",
        "entity",
        "file",
        "UploadableFilePackConverter",
      ]),
      (Mt.toFileInfoDynamicArray = function (e) {
        return k.array(
          e
            .filter(function (e) {
              return null != e.file || null != e.uploadedFile;
            })
            .map(Mt.toFileInfoDynamic),
        );
      }),
      (Mt.toFileInfoDynamic = function (e) {
        var t;
        return null != e.uploadedFile
          ? e.uploadedFile
          : {
              content_type: (t = e.file).type,
              content_size: t.size,
              name: xo.normalizeForFile(t.name),
              file: t,
              localThumbInfo: e.thumb,
            };
      }),
      ((t["albero.entity.message.IFavoriteMessageEvent"] = Ct).__name__ = [
        "albero",
        "entity",
        "message",
        "IFavoriteMessageEvent",
      ]),
      (Ct.prototype = { __class__: Ct }),
      ((t["albero.entity.message.AddFavoriteMessageEvent"] = kt).__name__ = [
        "albero",
        "entity",
        "message",
        "AddFavoriteMessageEvent",
      ]),
      (kt.__interfaces__ = [Ct]),
      (kt.prototype = {
        getMessage: function () {
          return this.message;
        },
        getMessageId: function () {
          return this.message.id;
        },
        isFavorite: function () {
          return !0;
        },
        getDomainId: function () {
          return this.domainId;
        },
        getTalkId: function () {
          return this.talkId;
        },
        __class__: kt,
      }),
      ((t["albero.entity.message.DeleteFavoriteMessageEvent"] = Dt).__name__ = [
        "albero",
        "entity",
        "message",
        "DeleteFavoriteMessageEvent",
      ]),
      (Dt.__interfaces__ = [Ct]),
      (Dt.prototype = {
        getMessageId: function () {
          return this.messageId;
        },
        getTalkId: function () {
          return this.talkId;
        },
        getDomainId: function () {
          return this.domainId;
        },
        isFavorite: function () {
          return !1;
        },
        __class__: Dt,
      }),
      ((t["albero.entity.message.IMessageContentForActionReply"] =
        Rt).__name__ = [
        "albero",
        "entity",
        "message",
        "IMessageContentForActionReply",
      ]),
      (Rt.prototype = { __class__: Rt }),
      function (e) {
        null != e &&
          ((this.noteId = new So(e.note_id)),
          (this.title = e.title),
          (this.revision = e.revision),
          (this.deleted = e.deleted),
          (this.hasAttachments = !!e.has_attachments));
      }),
    ao =
      (((t["albero.entity.message.MessageContentForNote"] = oo).__name__ = [
        "albero",
        "entity",
        "message",
        "MessageContentForNote",
      ]),
      (oo.prototype = { __class__: oo }),
      function (e) {
        null != e &&
          ((this.noteId = new So(e.note_id)), (this.title = e.title));
      }),
    so =
      (((t["albero.entity.message.MessageContentForNoteDeleted"] =
        ao).__name__ = [
        "albero",
        "entity",
        "message",
        "MessageContentForNoteDeleted",
      ]),
      (ao.prototype = { __class__: ao }),
      function (e) {
        ((this.text = m.None),
          (this.stampsetId = D.fromIntOrInt64(e.stampset_id)),
          (this.stampId = D.fromIntOrInt64(e.stamp_id)),
          null != e.text && (this.text = m.Some(e.text)));
      }),
    lo =
      (((t["albero.entity.message.MessageContentForOriginalStamp"] =
        so).__name__ = [
        "albero",
        "entity",
        "message",
        "MessageContentForOriginalStamp",
      ]),
      (so.fromIds = function (e, t) {
        return new so({ stampset_id: e, stamp_id: t });
      }),
      (so.prototype = {
        hasValidText: function () {
          return h.isDefined(this.getValidText());
        },
        getValidText: function () {
          return h.filter(this.text, Ln.isNotEmpty);
        },
        __class__: so,
      }),
      function (e) {
        ((this.inReplyTo = e.in_reply_to),
          (this.response = e.response),
          (this.options = e.options),
          (this.listing = null == e.listing || e.listing),
          (this.closingType = null == e.closing_type ? 1 : e.closing_type),
          (this.closed = null != e.closed && e.closed));
      }),
    uo =
      (((t["albero.entity.message.MessageContentForSelectOneReply"] =
        lo).__name__ = [
        "albero",
        "entity",
        "message",
        "MessageContentForSelectOneReply",
      ]),
      (lo.__interfaces__ = [Rt]),
      (lo.prototype = {
        getResponseIndex: function () {
          return this.response;
        },
        __class__: lo,
      }),
      function (e) {
        ((this.text = m.None),
          (this.stampSet = e.stamp_set),
          (this.stampIndex = e.stamp_index),
          null != e.text && (this.text = m.Some(e.text)));
      }),
    co =
      (((t["albero.entity.message.MessageContentForStamp"] = uo).__name__ = [
        "albero",
        "entity",
        "message",
        "MessageContentForStamp",
      ]),
      (uo.prototype = {
        hasValidText: function () {
          return h.isDefined(this.getValidText());
        },
        getValidText: function () {
          return h.filter(this.text, Ln.isNotEmpty);
        },
        getUrl: function () {
          return (
            "./assets/images/stamp/" +
            this.stampSet +
            "/" +
            Va.toString(this.stampIndex) +
            ".png"
          );
        },
        isValidStamp: function () {
          return 0 <= this.stampSet;
        },
        __class__: uo,
      }),
      function (e) {
        ((this.inReplyTo = e.in_reply_to),
          (this.title = e.title),
          (this.listing = null == e.listing || e.listing),
          (this.closingType = null == e.closing_type ? 1 : e.closing_type),
          (this.done = e.done),
          (this.closed = null != e.closed && e.closed));
      }),
    _o =
      (((t["albero.entity.message.MessageContentForTodoReply"] = co).__name__ =
        ["albero", "entity", "message", "MessageContentForTodoReply"]),
      (co.__interfaces__ = [Rt]),
      (co.prototype = {
        getResponseIndex: function () {
          return this.done ? 0 : 1;
        },
        __class__: co,
      }),
      function (e) {
        ((this.inReplyTo = e.in_reply_to),
          (this.response = e.response),
          (this.question = e.question),
          (this.listing = null == e.listing || e.listing),
          (this.closingType = null == e.closing_type ? 1 : e.closing_type),
          (this.closed = null != e.closed && e.closed));
      }),
    ho =
      (((t["albero.entity.message.MessageContentForYesNoReply"] = _o).__name__ =
        ["albero", "entity", "message", "MessageContentForYesNoReply"]),
      (_o.__interfaces__ = [Rt]),
      (_o.prototype = {
        getResponseIndex: function () {
          return this.response ? 0 : 1;
        },
        __class__: _o,
      }),
      function (e) {
        null == e
          ? ((this.favorite = !1), (this.createdAt = null), (this.tags = []))
          : ((this.favorite = !0),
            (this.createdAt = e.created_at),
            (this.tags = e.tags));
      }),
    fo =
      (((t["albero.entity.message.MessageFavorite"] = ho).__name__ = [
        "albero",
        "entity",
        "message",
        "MessageFavorite",
      ]),
      (ho.prototype = {
        isFavorite: function () {
          return this.favorite;
        },
        update: function (e) {
          this.favorite = e.isFavorite();
        },
        __class__: ho,
      }),
      function (e) {
        ((this.isDummy = !1),
          null != e &&
            ((this.id = e.scheduled_message_id),
            (this.createMessageAt = e.create_message_at),
            (this.talkId = e.talk_id),
            (this.type = O.typeOf(e.type)),
            (this.content = e.content)));
      }),
    mo =
      (((t["albero.entity.message.ScheduledMessage"] = fo).__name__ = [
        "albero",
        "entity",
        "message",
        "ScheduledMessage",
      ]),
      (fo.fromParams = function (e) {
        var t = new fo(null);
        return (
          (t.id = e.id),
          (t.createMessageAt = e.createMessageAt),
          (t.talkId = e.talkId),
          (t.type = e.type),
          (t.content = e.content),
          t
        );
      }),
      (fo.fromNotifyScheduleMessage = function (e) {
        var t = new fo(null);
        return (
          (t.id = e.id),
          (t.talkId = e.talkId),
          (t.createMessageAt = W.oneDayAfterAsInt64()),
          (t.type = d.text),
          (t.content = ""),
          (t.isDummy = !0),
          t
        );
      }),
      (fo.compare = function (e, t) {
        var n = e.createMessageAt,
          i = t.createMessageAt,
          r = (n.high - i.high) | 0,
          n =
            (0 == r && (r = w.ucompare(n.low, i.low)),
            n.high < 0 ? (i.high < 0 ? r : -1) : 0 <= i.high ? r : 1);
        return 0 != n
          ? n
          : (0 == (n = ((i = e.id).high - (r = t.id).high) | 0) &&
              (n = w.ucompare(i.low, r.low)),
            i.high < 0 ? (r.high < 0 ? n : -1) : 0 <= r.high ? n : 1);
      }),
      (fo.prototype = { __class__: fo }),
      ((t["albero.entity.message.ScheduledMessageEdit"] = Ot).__name__ = [
        "albero",
        "entity",
        "message",
        "ScheduledMessageEdit",
      ]),
      (Ot.prototype = {
        toEditWithError: function () {
          return new Ot(
            this.type,
            this.content,
            this.files,
            this.createMessageAt,
            !0,
          );
        },
        __class__: Ot,
      }),
      function () {
        ((this.scheduledMessageEdits = new B()),
          (this.scheduledMessages = new B()));
      }),
    po =
      (((t["albero.entity.message.ScheduledMessageStore"] = mo).__name__ = [
        "albero",
        "entity",
        "message",
        "ScheduledMessageStore",
      ]),
      (mo.prototype = {
        getScheduledMessage: function (e) {
          var t;
          return null == e
            ? null
            : ((e = "_" + e.high + "_" + e.low),
              (t = this.scheduledMessages),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        getScheduledMessages: function () {
          return k.array(this.scheduledMessages);
        },
        setScheduledMessage: function (e) {
          var t = e.id,
            t = "_" + t.high + "_" + t.low,
            n = this.scheduledMessages;
          null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e);
        },
        removeScheduledMessage: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.scheduledMessages,
            t = null != Y[e] ? t.getReserved(e) : t.h[e];
          return (this.scheduledMessages.remove(e), t);
        },
        getScheduledMessageEdit: function (e) {
          var t;
          return null == e
            ? null
            : ((e = "_" + e.high + "_" + e.low),
              (t = this.scheduledMessageEdits),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        setScheduledMessageEdit: function (e, t) {
          var e = "_" + e.high + "_" + e.low,
            n = this.scheduledMessageEdits;
          null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
        },
        removeScheduledMessageEdit: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.scheduledMessageEdits,
            t = null != Y[e] ? t.getReserved(e) : t.h[e];
          return (this.scheduledMessageEdits.remove(e), t);
        },
        __class__: mo,
      }),
      function (e) {
        null != e &&
          ((this.noteId = new So(e.note_id)),
          (this.talkId = e.talk_id),
          (this.createdBy = e.created_by),
          (this.createdAt = e.created_at),
          (this.setting = new No(e.setting)),
          (this.noteRevision = new Eo(e.note_revision)),
          (this.noteLocked = new Io(e.locked)));
      }),
    go =
      (((t["albero.entity.note.Note"] = po).__name__ = [
        "albero",
        "entity",
        "note",
        "Note",
      ]),
      (po.prototype = {
        isSameNoteId: function (e) {
          return this.noteId.equals(e);
        },
        getRevisionCreateUserId: function () {
          return this.noteRevision.createdBy;
        },
        getRevisionCreateDate: function () {
          return this.noteRevision.createdAt;
        },
        compareRevision: function (e) {
          return this.noteRevision.compareRevision(e);
        },
        lock: function (e, t, n, i) {
          this.noteRevision.revision == e && this.noteLocked.lock(t, n, i);
        },
        unlock: function (e, t) {
          this.noteRevision.revision == e && this.noteLocked.unlock(t);
        },
        getLockedUserIdWithoutSelf: function (e) {
          return this.noteLocked.getLockedUserIdWithoutSelf(e);
        },
        isSameTalkId: function (e) {
          var t = this.talkId;
          return null != t && null != e && t.high == e.high && t.low == e.low;
        },
        __class__: po,
      }),
      (t["albero.entity.note.NoteContentType"] = {
        __ename__: ["albero", "entity", "note", "NoteContentType"],
        __constructs__: ["TEXT", "FILES", "RICH_TEXT", "UNKNOWN"],
      })),
    yo =
      ((go.TEXT = ["TEXT", 0]),
      (go.TEXT.toString = a),
      ((go.TEXT.__enum__ = go).FILES = ["FILES", 1]),
      (go.FILES.toString = a),
      ((go.FILES.__enum__ = go).RICH_TEXT = ["RICH_TEXT", 2]),
      (go.RICH_TEXT.toString = a),
      ((go.RICH_TEXT.__enum__ = go).UNKNOWN = ["UNKNOWN", 3]),
      (go.UNKNOWN.toString = a),
      ((go.UNKNOWN.__enum__ = go).__empty_constructs__ = [
        go.TEXT,
        go.FILES,
        go.RICH_TEXT,
        go.UNKNOWN,
      ]),
      function () {}),
    So =
      (((t["albero.entity.note.NoteContentTypeHelper"] = yo).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteContentTypeHelper",
      ]),
      (yo.createNoteContentType = function (e) {
        if (null == e) return (F.e("no value."), null);
        switch (e) {
          case 1:
            return go.TEXT;
          case 5:
            return go.FILES;
          case 13:
            return go.RICH_TEXT;
          default:
            return (F.e("value is not expected."), go.UNKNOWN);
        }
      }),
      (yo.getValue = function (e) {
        if (null == e) return (F.e("no content type."), null);
        switch (e[1]) {
          case 0:
            return 1;
          case 1:
            return 5;
          case 2:
            return 13;
          default:
            return (F.e("content type is not expected."), null);
        }
      }),
      function (e) {
        this.value = e;
      }),
    vo =
      (((t["albero.entity.note.NoteId"] = So).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteId",
      ]),
      (So.__interfaces__ = [Ce]),
      (So.createDummyId = function () {
        var e,
          t,
          n,
          i = So.nextDummyId,
          r = So.nextDummyId;
        return (
          (r =
            null == r
              ? null
              : ((e = new E(0, 1)),
                (t = (r.high - e.high) | 0),
                (n = (r.low - e.low) | 0),
                w.ucompare(r.low, e.low) < 0 && (t--, (t |= 0)),
                new E(t, n))),
          (So.nextDummyId = r),
          new So(i)
        );
      }),
      (So.prototype = {
        toString: function () {
          var e = this.value;
          return "_" + e.high + "_" + e.low;
        },
        equals: function (e) {
          var t;
          return (
            null != e &&
            ((t = this.value), (e = e.value), null != t) &&
            null != e &&
            t.high == e.high &&
            t.low == e.low
          );
        },
        getHtmlAttributeId: function () {
          return "note-" + this.toString();
        },
        isDummy: function () {
          return this.value.high < 0;
        },
        __class__: So,
      }),
      function (e, t, n, i, r, o, a, s, l) {
        ((this.state = f.SENDING),
          (this.talkId = e),
          (this.noteId = t),
          (this.title = n),
          (this.text = i),
          (this.fileInfos = r),
          (this.emitterKey = o),
          (this.createMessage = null == a || a),
          (this.isRichText = null != s && s),
          (this.plainText = null != l ? l : ""));
      }),
    f =
      (((t["albero.entity.note.NoteLocalEdit"] = vo).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteLocalEdit",
      ]),
      (vo.prototype = {
        hasEmitterKey: function () {
          return null != this.emitterKey;
        },
        getEmitterKey: function () {
          return this.emitterKey;
        },
        setState: function (e) {
          this.state = e;
        },
        getTalkId: function () {
          return this.talkId;
        },
        getNoteId: function () {
          return this.noteId;
        },
        getTitle: function () {
          return this.title;
        },
        getText: function () {
          return this.text;
        },
        getContent: function () {
          switch (this.getContentType()[1]) {
            case 0:
              return this.text;
            case 1:
              return {
                files: this.getUploadedFilesData(),
                text: null == this.text ? "" : this.text,
              };
            case 2:
              return {
                files: this.getUploadedFilesData(),
                rich_text: null == this.text ? "" : this.text,
              };
            case 3:
              return (F.e("unknown content type"), null);
          }
        },
        getFileInfos: function () {
          return this.fileInfos;
        },
        getState: function () {
          return this.state;
        },
        isSending: function () {
          return this.state == f.SENDING;
        },
        isFailed: function () {
          switch (this.state[1]) {
            case 0:
            case 8:
              return !1;
            case 1:
            case 2:
            case 3:
            case 4:
            case 5:
            case 6:
            case 7:
              return !0;
          }
        },
        isForCreate: function () {
          return this.noteId.isDummy();
        },
        canNotSend: function () {
          return 0 != this.getUnuploadedFileInfos().length;
        },
        getContentType: function () {
          return this.isRichText
            ? go.RICH_TEXT
            : 0 < this.fileInfos.length
              ? go.FILES
              : go.TEXT;
        },
        getUnuploadedFileInfos: function () {
          return this.fileInfos.filter(function (e) {
            return e.isNotUploaded();
          });
        },
        getUploadedFilesData: function () {
          return this.fileInfos
            .map(function (e) {
              return e.getUploadedFileData();
            })
            .filter(function (e) {
              return null != e;
            });
        },
        createNoteRevisionSummary: function (e) {
          var t = new wo();
          return (
            (t.revision = 0),
            (t.contentType = this.getContentType()),
            (t.title = this.title),
            (t.contentSummary =
              "" != this.plainText ? this.plainText : this.text),
            (t.createdBy = e),
            (t.createdAt = W.nowAsInt64()),
            (t.contentFiles = this.fileInfos),
            t
          );
        },
        conflict: function () {
          this.state = f.FAILED_BY_CONFLICT;
        },
        deleted: function () {
          this.state = f.FAILED_BY_NOT_FOUND;
        },
        getCreateMessage: function () {
          return this.createMessage;
        },
        __class__: vo,
      }),
      (t["albero.entity.note.NoteLocalEditState"] = {
        __ename__: ["albero", "entity", "note", "NoteLocalEditState"],
        __constructs__: [
          "SENDING",
          "FAILED_BY_CONFLICT",
          "FAILED_BY_NOT_FOUND",
          "FAILED_BY_NOT_FOUND_FORWARDING",
          "FAILED_BY_FILE",
          "FAILED_BY_UNKNOWN",
          "FAILED_BY_TEXT_TOO_LONG",
          "FAILED_BY_TEXT_TOO_COMPLEX",
          "COMPLETED",
        ],
      })),
    To =
      ((f.SENDING = ["SENDING", 0]),
      (f.SENDING.toString = a),
      ((f.SENDING.__enum__ = f).FAILED_BY_CONFLICT = ["FAILED_BY_CONFLICT", 1]),
      (f.FAILED_BY_CONFLICT.toString = a),
      ((f.FAILED_BY_CONFLICT.__enum__ = f).FAILED_BY_NOT_FOUND = [
        "FAILED_BY_NOT_FOUND",
        2,
      ]),
      (f.FAILED_BY_NOT_FOUND.toString = a),
      ((f.FAILED_BY_NOT_FOUND.__enum__ = f).FAILED_BY_NOT_FOUND_FORWARDING = [
        "FAILED_BY_NOT_FOUND_FORWARDING",
        3,
      ]),
      (f.FAILED_BY_NOT_FOUND_FORWARDING.toString = a),
      ((f.FAILED_BY_NOT_FOUND_FORWARDING.__enum__ = f).FAILED_BY_FILE =
        function (e, t) {
          e = ["FAILED_BY_FILE", 4, e, t];
          return ((e.__enum__ = f), (e.toString = a), e);
        }),
      (f.FAILED_BY_UNKNOWN = ["FAILED_BY_UNKNOWN", 5]),
      (f.FAILED_BY_UNKNOWN.toString = a),
      ((f.FAILED_BY_UNKNOWN.__enum__ = f).FAILED_BY_TEXT_TOO_LONG = [
        "FAILED_BY_TEXT_TOO_LONG",
        6,
      ]),
      (f.FAILED_BY_TEXT_TOO_LONG.toString = a),
      ((f.FAILED_BY_TEXT_TOO_LONG.__enum__ = f).FAILED_BY_TEXT_TOO_COMPLEX = [
        "FAILED_BY_TEXT_TOO_COMPLEX",
        7,
      ]),
      (f.FAILED_BY_TEXT_TOO_COMPLEX.toString = a),
      ((f.FAILED_BY_TEXT_TOO_COMPLEX.__enum__ = f).COMPLETED = [
        "COMPLETED",
        8,
      ]),
      (f.COMPLETED.toString = a),
      ((f.COMPLETED.__enum__ = f).__empty_constructs__ = [
        f.SENDING,
        f.FAILED_BY_CONFLICT,
        f.FAILED_BY_NOT_FOUND,
        f.FAILED_BY_NOT_FOUND_FORWARDING,
        f.FAILED_BY_UNKNOWN,
        f.FAILED_BY_TEXT_TOO_LONG,
        f.FAILED_BY_TEXT_TOO_COMPLEX,
        f.COMPLETED,
      ]),
      ((t["albero.entity.note.NoteLockEvent"] = Ft).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteLockEvent",
      ]),
      (Ft.prototype = { __class__: Ft }),
      ((t["albero.entity.note.NoteLockExpiredAt"] = xt).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteLockExpiredAt",
      ]),
      (xt.prototype = {
        fromNow: function () {
          var e = this.value,
            t = e.high,
            e = e.low;
          return (
            (4294967296 * t +
              (0 <= e ? e : e + 4294967296) -
              new Date().getTime()) |
            0
          );
        },
        calcRelockDelay: function () {
          return (0.9 * this.fromNow()) | 0;
        },
        __class__: xt,
      }),
      (t["albero.entity.note.NoteLockStateValue"] = {
        __ename__: ["albero", "entity", "note", "NoteLockStateValue"],
        __constructs__: [
          "FAILED_BY_CONFLICT",
          "FAILED_BY_LOCKED",
          "FAILED_BY_NOT_FOUND",
          "FAILED_BY_UNKNOWN",
          "COMPLETED",
          "RELEASED",
        ],
      })),
    Io =
      ((To.FAILED_BY_CONFLICT = ["FAILED_BY_CONFLICT", 0]),
      (To.FAILED_BY_CONFLICT.toString = a),
      ((To.FAILED_BY_CONFLICT.__enum__ = To).FAILED_BY_LOCKED = function (e) {
        e = ["FAILED_BY_LOCKED", 1, e];
        return ((e.__enum__ = To), (e.toString = a), e);
      }),
      (To.FAILED_BY_NOT_FOUND = ["FAILED_BY_NOT_FOUND", 2]),
      (To.FAILED_BY_NOT_FOUND.toString = a),
      ((To.FAILED_BY_NOT_FOUND.__enum__ = To).FAILED_BY_UNKNOWN = [
        "FAILED_BY_UNKNOWN",
        3,
      ]),
      (To.FAILED_BY_UNKNOWN.toString = a),
      ((To.FAILED_BY_UNKNOWN.__enum__ = To).COMPLETED = function (e) {
        e = ["COMPLETED", 4, e];
        return ((e.__enum__ = To), (e.toString = a), e);
      }),
      (To.RELEASED = ["RELEASED", 5]),
      (To.RELEASED.toString = a),
      ((To.RELEASED.__enum__ = To).__empty_constructs__ = [
        To.FAILED_BY_CONFLICT,
        To.FAILED_BY_NOT_FOUND,
        To.FAILED_BY_UNKNOWN,
        To.RELEASED,
      ]),
      ((t["albero.entity.note.NoteLockState"] = Ut).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteLockState",
      ]),
      (Ut.prototype = {
        isCompleted: function () {
          return 4 == this.value[1];
        },
        getExpiredAtOrNull: function () {
          var e = this.value;
          return 4 == e[1] ? e[2] : null;
        },
        isLockedByOther: function () {
          return 1 == this.value[1];
        },
        getLockedUser: function () {
          var e = this.value;
          return 1 == e[1] ? e[2] : null;
        },
        __class__: Ut,
      }),
      function (e) {
        null != e &&
          ((this.userId = e.user_id),
          (this.deviceId = e.device_id),
          (this.expiredAt = e.expired_at));
      }),
    Eo =
      (((t["albero.entity.note.NoteLocked"] = Io).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteLocked",
      ]),
      (Io.prototype = {
        lock: function (e, t, n) {
          ((this.expiredAt = n), (this.deviceId = t), (this.userId = e));
        },
        unlock: function (e) {
          var t = this.userId;
          null != t &&
            null != e &&
            t.high == e.high &&
            t.low == e.low &&
            ((this.expiredAt = null),
            (this.userId = null),
            (this.deviceId = null));
        },
        getLockedUserIdWithoutSelf: function (e) {
          var t = this.deviceId;
          return (null != t &&
            null != e &&
            t.high == e.high &&
            t.low == e.low) ||
            null == this.expiredAt ||
            4294967296 * (t = this.expiredAt).high +
              (0 <= (e = t.low) ? e : e + 4294967296) <=
              new Date().getTime()
            ? null
            : this.userId;
        },
        __class__: Io,
      }),
      ((t["albero.entity.note.NotePartialUpdate"] = Lt).__name__ = [
        "albero",
        "entity",
        "note",
        "NotePartialUpdate",
      ]),
      (Lt.prototype = {
        getRevision: function () {
          return this.noteRevisionSummary.revision;
        },
        __class__: Lt,
      }),
      ((t["albero.entity.note.NotePartialUpdateForSetting"] = Pt).__name__ = [
        "albero",
        "entity",
        "note",
        "NotePartialUpdateForSetting",
      ]),
      (Pt.prototype = { __class__: Pt }),
      ((t["albero.entity.note.NotePartialUpdateForSummary"] = Bt).__name__ = [
        "albero",
        "entity",
        "note",
        "NotePartialUpdateForSummary",
      ]),
      (Bt.prototype = {
        getRevision: function () {
          return this.noteRevisionSummary.revision;
        },
        getUpdatedBy: function () {
          return this.noteRevisionSummary.createdBy;
        },
        getUpdatedAt: function () {
          return this.noteRevisionSummary.createdAt;
        },
        __class__: Bt,
      }),
      function (e) {
        if (
          ((this.contentFiles = []),
          null != e &&
            ((this.revision = e.revision),
            (this.title = e.title),
            (this.contentType = yo.createNoteContentType(e.content_type)),
            (this.createdBy = e.created_by),
            (this.createdAt = e.created_at),
            null != this.contentType))
        ) {
          var t = e.content;
          switch (this.contentType[1]) {
            case 0:
              this.contentText = null == t ? "" : t;
              break;
            case 1:
            case 3:
              if (
                ((this.contentText = null == t || null == t.text ? "" : t.text),
                null != t && null != t.files)
              )
                for (var n = 0, i = G.__cast(t.files, Array); n < i.length; ) {
                  var r = i[n];
                  (++n, this.contentFiles.push(new ar(r)));
                }
              break;
            case 2:
              if (
                ((this.contentText = null == t ? "" : t.rich_text),
                null != t && null != t.files)
              )
                for (var o = 0, a = G.__cast(t.files, Array); o < a.length; ) {
                  var s = a[o];
                  (++o, this.contentFiles.push(new ar(s)));
                }
          }
        }
      }),
    wo =
      (((t["albero.entity.note.NoteRevision"] = Eo).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteRevision",
      ]),
      (Eo.prototype = {
        isEmptyTitle: function () {
          return "" == this.title || null == this.title;
        },
        compareRevision: function (e) {
          return this.revision == e ? 0 : this.revision < e ? -1 : 1;
        },
        __class__: Eo,
      }),
      function (e) {
        if (
          null != e &&
          ((this.revision = e.revision),
          (this.title = e.title),
          (this.contentType = yo.createNoteContentType(e.content_type)),
          (this.contentSummary = e.content_summary),
          (this.createdBy = e.created_by),
          (this.createdAt = e.created_at),
          (this.contentFiles = []),
          null != e.content_files)
        )
          for (
            var t = 0, n = G.__cast(e.content_files, Array);
            t < n.length;
          ) {
            var i = n[t];
            (++t, this.contentFiles.push(new ar(i)));
          }
      }),
    No =
      (((t["albero.entity.note.NoteRevisionSummary"] = wo).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteRevisionSummary",
      ]),
      (wo.prototype = {
        isEmptyTitle: function () {
          return "" == this.title || null == this.title;
        },
        compareRevision: function (e) {
          return this.revision == e ? 0 : this.revision < e ? -1 : 1;
        },
        getContentSummaryWithEllipsis: function () {
          var e = I.replace(
            I.replace(this.contentSummary, "- [x] ", ""),
            "- [ ] ",
            "",
          );
          return e.length < 100 ? e : e.substring(0, 99) + "…";
        },
        getContentSummaryWithoutCheckboxes: function () {
          return I.replace(
            I.replace(this.contentSummary, "- [x] ", ""),
            "- [ ] ",
            "",
          );
        },
        __class__: wo,
      }),
      function (e) {
        null != e &&
          ((this.collaborativeEditing = e.collaborative_editing),
          (this.version = e.version));
      }),
    Ao =
      (((t["albero.entity.note.NoteSetting"] = No).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteSetting",
      ]),
      (No.prototype = { __class__: No }),
      function (e) {
        null != e &&
          ((this.noteId = new So(e.note_id)),
          (this.noteSummary = new bo(e.note_summary)),
          (this.commentIds = e.comment_ids),
          (this.readUserIds = e.read_user_ids));
      }),
    bo =
      (((t["albero.entity.note.NoteStatus"] = Ao).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteStatus",
      ]),
      (Ao.prototype = {
        getNoteRevisionSummary: function () {
          return this.noteSummary.noteRevisionSummary;
        },
        __class__: Ao,
      }),
      function (e) {
        null != e &&
          ((this.noteId = new So(e.note_id)),
          (this.talkId = e.talk_id),
          (this.createdBy = e.created_by),
          (this.createdAt = e.created_at),
          (this.setting = new No(e.setting)),
          (this.noteRevisionSummary = new wo(e.note_revision_summary)),
          (this.noteLocked = new Io(e.locked)));
      }),
    Mo =
      (((t["albero.entity.note.NoteSummary"] = bo).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteSummary",
      ]),
      (bo.prototype = {
        getNoteIdByString: function () {
          return this.noteId.toString();
        },
        getRevisionCreatedAt: function () {
          return this.noteRevisionSummary.createdAt;
        },
        getRevisionContentSummary: function () {
          return this.noteRevisionSummary.contentSummary;
        },
        getRevisionTitle: function () {
          return this.noteRevisionSummary.isEmptyTitle()
            ? ""
            : this.noteRevisionSummary.title;
        },
        lock: function (e, t, n, i) {
          this.noteRevisionSummary.revision == e &&
            this.noteLocked.lock(t, n, i);
        },
        unlock: function (e, t) {
          this.noteRevisionSummary.revision == e && this.noteLocked.unlock(t);
        },
        __class__: bo,
      }),
      ((t["albero.entity.note.NoteTextParts"] = Ht).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteTextParts",
      ]),
      (Ht.isUncheckedCheckBox = function (e) {
        return I.startsWith(e, "- [ ] ");
      }),
      (Ht.isCheckedCheckBox = function (e) {
        return (
          !!I.startsWith(e, "- [x] ") || I.startsWith(e, "- [x] ".toUpperCase())
        );
      }),
      ((t["albero.entity.note.NoteUnlockEvent"] = Gt).__name__ = [
        "albero",
        "entity",
        "note",
        "NoteUnlockEvent",
      ]),
      (Gt.prototype = { __class__: Gt }),
      function (e) {
        ((this.id = D.fromIntOrInt64(e.id)),
          (this.order = e.order),
          (this.imageUrl = e.imageUrl),
          (this.type = Mo.getReactionTypeFromReactionId(this.id)));
      }),
    Co =
      (((t["albero.entity.reaction.AvailableMessageReaction"] = Mo).__name__ = [
        "albero",
        "entity",
        "reaction",
        "AvailableMessageReaction",
      ]),
      (Mo.compare = function (e, t) {
        return e.order - t.order;
      }),
      (Mo.getReactionTypeFromReactionId = function (e) {
        var t = _.make(0, 1);
        return null != e && null != t && e.high == t.high && e.low == t.low
          ? Co.ThumbsUp
          : ((t = _.make(0, 2)),
            null != e && null != t && e.high == t.high && e.low == t.low
              ? Co.Check
              : ((t = _.make(0, 3)),
                null != e && null != t && e.high == t.high && e.low == t.low
                  ? Co.Heart
                  : ((t = _.make(0, 4)),
                    null != e && null != t && e.high == t.high && e.low == t.low
                      ? Co.Thank
                      : ((t = _.make(0, 5)),
                        null != e &&
                        null != t &&
                        e.high == t.high &&
                        e.low == t.low
                          ? Co.Eyes
                          : ((t = _.make(0, 6)),
                            null != e &&
                            null != t &&
                            e.high == t.high &&
                            e.low == t.low
                              ? Co.Weep
                              : Co.Custom(e))))));
      }),
      (Mo.getReactionIdFromReactionType = function (e) {
        switch (e[1]) {
          case 0:
            return _.make(0, 2);
          case 1:
            return _.make(0, 1);
          case 2:
            return _.make(0, 3);
          case 3:
            return _.make(0, 4);
          case 4:
            return _.make(0, 6);
          case 5:
            return _.make(0, 5);
          case 6:
            return e[2];
        }
      }),
      (Mo.prototype = { __class__: Mo }),
      (t["albero.entity.reaction.ReactionType"] = {
        __ename__: ["albero", "entity", "reaction", "ReactionType"],
        __constructs__: [
          "Check",
          "ThumbsUp",
          "Heart",
          "Thank",
          "Weep",
          "Eyes",
          "Custom",
        ],
      })),
    ko =
      ((Co.Check = ["Check", 0]),
      (Co.Check.toString = a),
      ((Co.Check.__enum__ = Co).ThumbsUp = ["ThumbsUp", 1]),
      (Co.ThumbsUp.toString = a),
      ((Co.ThumbsUp.__enum__ = Co).Heart = ["Heart", 2]),
      (Co.Heart.toString = a),
      ((Co.Heart.__enum__ = Co).Thank = ["Thank", 3]),
      (Co.Thank.toString = a),
      ((Co.Thank.__enum__ = Co).Weep = ["Weep", 4]),
      (Co.Weep.toString = a),
      ((Co.Weep.__enum__ = Co).Eyes = ["Eyes", 5]),
      (Co.Eyes.toString = a),
      ((Co.Eyes.__enum__ = Co).Custom = function (e) {
        e = ["Custom", 6, e];
        return ((e.__enum__ = Co), (e.toString = a), e);
      }),
      (Co.__empty_constructs__ = [
        Co.Check,
        Co.ThumbsUp,
        Co.Heart,
        Co.Thank,
        Co.Weep,
        Co.Eyes,
      ]),
      function () {
        ((this.availableMessageReactionsMap = new B()),
          (this.summaries = new B()),
          (this.usersMap = new B()),
          (this.loadingMap = new B()));
      }),
    Do =
      (((t["albero.entity.reaction.MessageReactionStore"] = ko).__name__ = [
        "albero",
        "entity",
        "reaction",
        "MessageReactionStore",
      ]),
      (ko.prototype = {
        updateAvailableMessageReactions: function (e, t) {
          t.sort(Mo.compare);
          var e = "_" + e.high + "_" + e.low,
            n = this.availableMessageReactionsMap;
          null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
        },
        getAvailableMessageReactions: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.availableMessageReactionsMap,
            t = null != Y[e] ? t.getReserved(e) : t.h[e];
          return null == t ? [] : t;
        },
        isAvailableMessageReactionsLoaded: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.availableMessageReactionsMap;
          return null != Y[e] ? t.existsReserved(e) : t.h.hasOwnProperty(e);
        },
        setMessageReactionSummary: function (e) {
          var t = e.messageId,
            t = "_" + t.high + "_" + t.low,
            n = this.summaries;
          null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e);
        },
        removeMessageReactionSummary: function (e) {
          this.summaries.remove("_" + e.high + "_" + e.low);
        },
        getMessageReactionSummary: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.summaries;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        setMessageReactionUsers: function (e) {
          var t = e.messageId,
            t = "_" + t.high + "_" + t.low,
            n = this.usersMap;
          null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e);
        },
        removeMessageReactionUsers: function (e) {
          this.usersMap.remove("_" + e.high + "_" + e.low);
        },
        getMessageReactionUsers: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.usersMap;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        addLoadingMessageReactionSmmary: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.loadingMap;
          null != Y[e] ? t.setReserved(e, !0) : (t.h[e] = !0);
        },
        removeLoadingMessageReactionSmmary: function (e) {
          this.loadingMap.remove("_" + e.high + "_" + e.low);
        },
        isLoadingMessageReactionSmmary: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.loadingMap;
          return null != Y[e] ? t.existsReserved(e) : t.h.hasOwnProperty(e);
        },
        __class__: ko,
      }),
      ((t["albero.entity.reaction.MessageReactionSummary"] = jt).__name__ = [
        "albero",
        "entity",
        "reaction",
        "MessageReactionSummary",
      ]),
      (jt.prototype = {
        getReactionCount: function (e) {
          var t = this.reactionCounts,
            e = Va.toString(e),
            t = null != Y[e] ? t.getReserved(e) : t.h[e];
          return null == t ? 0 : t;
        },
        setReactionCount: function (e, t) {
          var n = this.reactionCounts,
            e = Va.toString(e);
          null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
        },
        updateMyReaction: function (e) {
          (null != this.myReactionId &&
            this.setReactionCount(
              this.myReactionId,
              this.getReactionCount(this.myReactionId) - 1,
            ),
            null != e && this.setReactionCount(e, this.getReactionCount(e) + 1),
            (this.myReactionId = e));
        },
        isMyReactionId: function (e) {
          var t = this.myReactionId;
          return null != e && null != t && e.high == t.high && e.low == t.low;
        },
        hasMyReaction: function () {
          return null != this.myReactionId;
        },
        __class__: jt,
      }),
      ((t["albero.entity.reaction.MessageReactionUsers"] = zt).__name__ = [
        "albero",
        "entity",
        "reaction",
        "MessageReactionUsers",
      ]),
      (zt.prototype = {
        getReactionUserIds: function (e) {
          var t = this.reactionUserIds,
            e = Va.toString(e),
            t = null != Y[e] ? t.getReserved(e) : t.h[e];
          return null == t ? [] : t;
        },
        isReactionUser: function (e) {
          for (
            var t = this.reactionUserIds, n = new is(t, t.arrayKeys());
            n.hasNext();
          ) {
            var i = n.next();
            if (D.contains(i, e)) return !0;
          }
          return !1;
        },
        __class__: zt,
      }),
      ((t["albero.js.KatakanaToRoman"] = Yt).__name__ = [
        "albero",
        "js",
        "KatakanaToRoman",
      ]),
      (Yt.assureConvertTable = function () {
        return (
          null == Yt.convertTable &&
            ((Yt.convertTable = new B()),
            k.iter(
              [
                ["ア", "A"],
                ["イ", "I"],
                ["ウ", "U"],
                ["エ", "E"],
                ["オ", "O"],
                ["カ", "KA"],
                ["キ", "KI"],
                ["ク", "KU"],
                ["ケ", "KE"],
                ["コ", "KO"],
                ["キャ", "KYA"],
                ["キュ", "KYU"],
                ["キョ", "KYO"],
                ["サ", "SA"],
                ["シ", "SI"],
                ["ス", "SU"],
                ["セ", "SE"],
                ["ソ", "SO"],
                ["シャ", "SHA"],
                ["シュ", "SHU"],
                ["ショ", "SHO"],
                ["タ", "TA"],
                ["チ", "TI"],
                ["ツ", "TU"],
                ["テ", "TE"],
                ["ト", "TO"],
                ["チャ", "THA"],
                ["チュ", "THU"],
                ["チョ", "THO"],
                ["ナ", "NA"],
                ["ニ", "NI"],
                ["ヌ", "NU"],
                ["ネ", "NE"],
                ["ノ", "NO"],
                ["ニャ", "NYA"],
                ["ニュ", "NYU"],
                ["ニョ", "NYO"],
                ["ハ", "HA"],
                ["ヒ", "HI"],
                ["フ", "FU"],
                ["ヘ", "HE"],
                ["ホ", "HO"],
                ["ヒャ", "HYA"],
                ["ヒュ", "HYU"],
                ["ヒョ", "HYO"],
                ["マ", "MA"],
                ["ミ", "MI"],
                ["ム", "MU"],
                ["メ", "ME"],
                ["モ", "MO"],
                ["ミャ", "MYA"],
                ["ミュ", "MYU"],
                ["ミョ", "MYO"],
                ["ヤ", "YA"],
                ["ユ", "YU"],
                ["ヨ", "YO"],
                ["ラ", "RA"],
                ["リ", "RI"],
                ["ル", "RU"],
                ["レ", "RE"],
                ["ロ", "RO"],
                ["リャ", "RYA"],
                ["リュ", "RYU"],
                ["リョ", "RYO"],
                ["ワ", "WA"],
                ["ヲ", "O"],
                ["ン", "N"],
                ["ガ", "GA"],
                ["ギ", "GI"],
                ["グ", "GU"],
                ["ゲ", "GE"],
                ["ゴ", "GO"],
                ["ギャ", "GYA"],
                ["ギュ", "GYU"],
                ["ギョ", "GYO"],
                ["ザ", "ZA"],
                ["ジ", "JI"],
                ["ズ", "ZU"],
                ["ゼ", "ZE"],
                ["ゾ", "ZO"],
                ["ジャ", "ZYA"],
                ["ジュ", "ZYU"],
                ["ジョ", "ZYO"],
                ["ダ", "DA"],
                ["ヂ", "JI"],
                ["ヅ", "ZU"],
                ["デ", "DE"],
                ["ド", "DO"],
                ["バ", "BA"],
                ["ビ", "BI"],
                ["ブ", "BU"],
                ["ベ", "BE"],
                ["ボ", "BO"],
                ["ビャ", "BYA"],
                ["ビュ", "BYU"],
                ["ビョ", "BYO"],
                ["パ", "PA"],
                ["ピ", "PI"],
                ["プ", "PU"],
                ["ペ", "PE"],
                ["ポ", "PO"],
                ["ピャ", "PYA"],
                ["ピュ", "PYU"],
                ["ピョ", "PYO"],
                ["ー", ""],
                ["ッ", "ッ"],
              ],
              function (e) {
                var t = Yt.convertTable,
                  n = e[0],
                  e = e[1];
                null != Y[n] ? t.setReserved(n, e) : (t.h[n] = e);
              },
            )),
          Yt.convertTable
        );
      }),
      (Yt.matchNext = function (e, t) {
        var n = Yt.assureConvertTable();
        if (t + 1 < e.length) {
          var i = e.substring(t, t + 2),
            i = null != Y[i] ? n.getReserved(i) : n.h[i];
          if (null != i) return Do.MatchedToTwoCharsKey(i);
        }
        if (t < e.length) {
          ((i = e.substring(t, t + 1)),
            (n = null != Y[i] ? n.getReserved(i) : n.h[i]));
          if (null != n) return Do.MatchedToOneCharKey(n);
        }
        i = t < e.length ? e.substring(t, t + 1) : "";
        return Do.NotMatched(i);
      }),
      (Yt.prototype = {
        convert: function (e) {
          for (var t = 0, n = ""; t < e.length; ) {
            var i = Yt.matchNext(e, t);
            switch (i[1]) {
              case 0:
                ((t += 2), (n += i[2]));
                break;
              case 1:
                (++t, (n += i[2]));
                break;
              case 2:
                (++t, (n += i[2]));
            }
          }
          return (n = (n = n.replace(Yt.TU_REG.r, "$1$1")).replace(
            Yt.XTU_REG.r,
            "XTU",
          ));
        },
        __class__: Yt,
      }),
      (t["albero.js.ResultOfMatchNext"] = {
        __ename__: ["albero", "js", "ResultOfMatchNext"],
        __constructs__: [
          "MatchedToTwoCharsKey",
          "MatchedToOneCharKey",
          "NotMatched",
        ],
      })),
    Ro =
      ((Do.MatchedToTwoCharsKey = function (e) {
        e = ["MatchedToTwoCharsKey", 0, e];
        return ((e.__enum__ = Do), (e.toString = a), e);
      }),
      (Do.MatchedToOneCharKey = function (e) {
        e = ["MatchedToOneCharKey", 1, e];
        return ((e.__enum__ = Do), (e.toString = a), e);
      }),
      (Do.NotMatched = function (e) {
        e = ["NotMatched", 2, e];
        return ((e.__enum__ = Do), (e.toString = a), e);
      }),
      (Do.__empty_constructs__ = []),
      function (e) {
        ((this.isDirty = !1), (this.key = e));
      }),
    c =
      (((t["albero.js.LocalStorage"] = Ro).__name__ = [
        "albero",
        "js",
        "LocalStorage",
      ]),
      (Ro.getLocalStorage = function () {
        return new (require("node-localstorage").LocalStorage)(
          null != qs.storagePath ? qs.storagePath : "./storage.local",
          null != qs.storageQuota ? P.parseInt(qs.storageQuota) : null,
        );
      }),
      (Ro.prototype = {
        load: function () {
          var e = this.val();
          if (null == e || "" == e) return null;
          var t = null;
          try {
            ((t = C.run(e)), (this.isDirty = !1));
          } catch (e) {
            (e instanceof H && (e = e.val), F.w(e));
          }
          return t;
        },
        save: function (e) {
          try {
            (this.val(Qa.run(e)), (this.isDirty = !1));
          } catch (e) {
            (e instanceof H && (e = e.val), F.w(e));
          }
        },
        setDirtyFlag: function () {
          if (!this.isDirty)
            try {
              (this.val(""), (this.isDirty = !0));
            } catch (e) {
              (e instanceof H && (e = e.val), F.w(e));
            }
        },
        val: function (e) {
          var t = Ro.getLocalStorage();
          return (
            null != t &&
              (null == e
                ? (e = t.getItem(this.key))
                : "" == e
                  ? t.removeItem(this.key)
                  : t.setItem(this.key, e)),
            e
          );
        },
        __class__: Ro,
      }),
      function () {}),
    F =
      (((t["albero.js.Localizer"] = c).__name__ = [
        "albero",
        "js",
        "Localizer",
      ]),
      (c.setLocalizeForMock = function (e, t) {
        null == c.mock && (c.mock = new B());
        var n = c.mock;
        null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
      }),
      (c.setLocalizeValueForMock = function (e, t) {
        c.setLocalizeForMock(e, function (e) {
          return t;
        });
      }),
      (c.localizeForMock = function (e, t) {
        var n;
        return null == c.mock ||
          ((n = c.mock), null == (n = null != Y[e] ? n.getReserved(e) : n.h[e]))
          ? null
          : n(t);
      }),
      (c.runOrReserve = function (e) {
        e();
      }),
      (c.localize = function (e, t) {
        return "";
      }),
      (c.asyncPrepare = function () {
        return Promise.resolve();
      }),
      (y.AlberoLog = function () {})),
    Oo =
      (((t.AlberoLog = F).__name__ = ["AlberoLog"]),
      (F.getLogLevel = function (e) {
        switch ((e = null != e ? e.toUpperCase() : "INFO")) {
          case "ALERT":
            return 1;
          case "CRITICAL":
            return 2;
          case "EMERGENCY":
            return 0;
          case "ERROR":
            return 3;
          case "INFO":
            return 6;
          case "NOTICE":
            return 5;
          case "WARNING":
            return 4;
          default:
            return 7;
        }
      }),
      (F._nop = function () {}),
      (F.dateStr = function () {
        return "[" + U.dateStr(new Date()) + "] ";
      }),
      (F.setInternalLogger = function (e) {
        F.logger = e;
      }),
      (F.v = function (e, t, n, i, r) {
        (null == r && (r = ""),
          null == i && (i = ""),
          null == n && (n = ""),
          null == t && (t = ""),
          null != F.logger
            ? 7 <= F.level && F.logger.verbose([e, t, n, i, r])
            : F._d(e, t, n, i, r));
      }),
      (F.d = function (e, t, n, i, r) {
        (null == r && (r = ""),
          null == i && (i = ""),
          null == n && (n = ""),
          null == t && (t = ""),
          null != F.logger
            ? 7 <= F.level && F.logger.debug([e, t, n, i, r])
            : F._d(e, t, n, i, r));
      }),
      (F.i = function (e, t, n, i, r) {
        (null == r && (r = ""),
          null == i && (i = ""),
          null == n && (n = ""),
          null == t && (t = ""),
          null != F.logger
            ? 6 <= F.level && F.logger.info([e, t, n, i, r])
            : F._i(e, t, n, i, r));
      }),
      (F.w = function (e, t, n, i, r) {
        (null == r && (r = ""),
          null == i && (i = ""),
          null == n && (n = ""),
          null == t && (t = ""),
          null != F.logger
            ? 4 <= F.level && F.logger.warn([e, t, n, i, r])
            : F._w(e, t, n, i, r));
      }),
      (F.e = function (e, t, n, i, r) {
        (null == r && (r = ""),
          null == i && (i = ""),
          null == n && (n = ""),
          null == t && (t = ""),
          null != F.logger
            ? 3 <= F.level && F.logger.error([e, t, n, i, r])
            : F._e(e, t, n, i, r));
      }),
      ((t["albero.js.NFKCTextSearch"] = Wt).__name__ = [
        "albero",
        "js",
        "NFKCTextSearch",
      ]),
      (Wt.prototype = {
        search: function (e) {
          for (
            var t = new Cn(),
              n = [],
              i =
                (this.ignoreCase && (e = e.toUpperCase()),
                this.normalize(e, t, n),
                0),
              r = [],
              o = t.b;
            -1 != (a = o.indexOf(this.normalizedPattern, i));
          ) {
            for (
              var i = a + this.normalizedPattern.length,
                a = n[a],
                s = n[i - 1],
                l = null,
                u = i,
                c = n.length;
              u < c && s == (l = n[u++]);
            )
              l = null;
            (null == l && (l = e.length), r.push(new Oo(a, l)));
          }
          return r;
        },
        normalize: function (e, t, n) {
          for (var i = e.length, r = new Cn(), o = 0, a = 0, s = i; a < s; ) {
            var l = a++,
              u = r.b.length,
              c = r.b,
              _ = c + e.charAt(l),
              h = this.normalizeFunc(_);
            if (_ != h)
              if (I.startsWith(h, c)) {
                for (
                  var _ = c.substring(0, u),
                    c =
                      ((t.b += P.string(_)),
                      (r = new Cn()),
                      h.substring(u, h.length)),
                    d = ((r.b += P.string(c)), 0),
                    f = u;
                  d < f;
                ) {
                  d++;
                  n.push(o);
                }
                o = l;
              } else (r = new Cn()).b += null == h ? "null" : "" + h;
            else {
              t.b += P.string(r);
              for (
                var r = new Cn(),
                  _ = e.charAt(l),
                  m = ((r.b += P.string(_)), 0),
                  p = u;
                m < p;
              ) {
                m++;
                n.push(o);
              }
              o = l;
            }
          }
          t.b += P.string(r.b);
          for (var g = 0, y = r.b.length; g < y; ) {
            g++;
            n.push(o);
          }
        },
        __class__: Wt,
      }),
      function (e, t) {
        ((this.start = e), (this.end = t));
      }),
    Fo =
      (((t["albero.js.HitSpan"] = Oo).__name__ = ["albero", "js", "HitSpan"]),
      (Oo.prototype = { __class__: Oo }),
      ((t["albero.js.PromiseHelper"] = Kt).__name__ = [
        "albero",
        "js",
        "PromiseHelper",
      ]),
      (Kt.resolveVoid = function (e) {
        e();
      }),
      (Kt.delay = function (n) {
        return new Promise(function (e, t) {
          Xa.delay(function () {
            Kt.resolveVoid(e);
          }, n);
        });
      }),
      function () {}),
    xo =
      (((t["albero.js.ProxyHelper"] = Fo).__name__ = [
        "albero",
        "js",
        "ProxyHelper",
      ]),
      (Fo.createAgent = function (e) {
        var e = ps.parse(e),
          t = { host: e.hostname },
          n =
            (null != e.port && (t.port = P.parseInt(e.port)),
            null != e.auth && (t.proxyAuth = e.auth),
            require("tunnel-agent"));
        return ("https:" == e.protocol ? n.httpsOverHttps : n.httpsOverHttp)({
          proxy: t,
        });
      }),
      ((t["albero.js.SettingsResolver"] = Vt).__name__ = [
        "albero",
        "js",
        "SettingsResolver",
      ]),
      (Vt.get = function (e, t) {
        return null != e ? e : t;
      }),
      (Vt.getAcceptableEventTimeDiff = function () {
        return qa.parseString(Vt.get(qs.acceptableEventTimeDiff, "10000"));
      }),
      function () {}),
    Uo =
      (((t["albero.js.TextCanonicalizer"] = xo).__name__ = [
        "albero",
        "js",
        "TextCanonicalizer",
      ]),
      (xo.canonicalize = function (e) {
        return xo.basicLatinAlphabetToUpperCase(
          xo.hiraganaToKatakana(xo.normalize(xo.eraseInvisible(e))),
        );
      }),
      (xo.canonicalizeForRoman = function (e) {
        return xo.basicLatinAlphabetToUpperCase(
          xo.katakanaToRoman(
            xo.hiraganaToKatakana(xo.normalize(xo.eraseInvisible(e))),
          ),
        );
      }),
      (xo.basicLatinAlphabetToUpperCase = function (e) {
        return e.toUpperCase();
      }),
      (xo.hiraganaToKatakana = function (e) {
        for (var t = 0, n = "", i = 0, r = e.length; i < r; ) {
          var o = i++,
            a = U.cca(e, o);
          a >= xo.HIRAGANA_SMALL_A &&
            a <= xo.HIRAGANA_NN &&
            ((n =
              (n += e.substring(t, o)) +
              String.fromCharCode(
                a + (xo.KATAKANA_SMALL_A - xo.HIRAGANA_SMALL_A),
              )),
            (t = 1 + o));
        }
        return (n += e.substring(t, e.length));
      }),
      (xo.katakanaToRoman = function (e) {
        return new Yt().convert(e);
      }),
      (xo.normalize = function (e) {
        return new Uo(null).normalize(e);
      }),
      (xo.normalizeForFile = function (e) {
        return new Uo(null).normalizeForFile(e);
      }),
      (xo.eraseInvisible = function (e) {
        var t = new RegExp("[\\u0000-\\u0020　]", "g".split("u").join(""));
        return e.replace(t, "");
      }),
      function (e) {
        ((this.beforeTask = null), (this.beforeTask = e));
      }),
    Lo =
      (((t["albero.js.Normalizer"] = Uo).__name__ = [
        "albero",
        "js",
        "Normalizer",
      ]),
      (Uo.prototype = {
        normalize: function (e) {
          try {
            return (
              null != this.beforeTask && this.beforeTask(),
              e.normalize("NFKC")
            );
          } catch (e) {
            return (
              e instanceof H && (e = e.val),
              F.e("invalid character:", e),
              c.localize("TextCanonicalizer.invalid")
            );
          }
        },
        normalizeForFile: function (e) {
          try {
            return (
              null != this.beforeTask && this.beforeTask(),
              e.normalize("NFC")
            );
          } catch (e) {
            return (
              e instanceof H && (e = e.val),
              F.e("invalid character:", e),
              c.localize("TextCanonicalizer.invalid")
            );
          }
        },
        __class__: Uo,
      }),
      ((t["albero.js.WebSocket"] = qt).__name__ = [
        "albero",
        "js",
        "WebSocket",
      ]),
      (qt.prototype = { __class__: qt }),
      ((t["albero.js.WebSocketFactory"] = Qt).__name__ = [
        "albero",
        "js",
        "WebSocketFactory",
      ]),
      (Qt.newInstance = function (e, t) {
        return new Lo(e, t);
      }),
      function (e, t) {
        var n = this,
          t =
            ((this.onopen = t.onopen),
            (this.onmessage = t.onmessage),
            (this.onerror = null),
            (this.onclose = t.onclose),
            (this.onpong = t.onpong),
            qs.wsConfig),
          i = require("websocket").client,
          i =
            ((this.ws = new i(t)),
            this.ws.on("connectFailed", z(this, this.onError)),
            this.ws.on("connect", function (e) {
              ((n.connection = e).on("error", z(n, n.onError)),
                e.on("close", z(n, n.onConnectionClose)),
                e.on("message", z(n, n.onMessage)),
                e.on("pong", z(n, n.onPong)),
                n.onOpen(null));
            }),
            null);
        (null != qs.proxyURL && (i = { agent: Fo.createAgent(qs.proxyURL) }),
          this.ws.connect(e, null, null, null, i));
      }),
    Po =
      (((t["albero.js.WebSocketForNodeJs"] = Lo).__name__ = [
        "albero",
        "js",
        "WebSocketForNodeJs",
      ]),
      (Lo.__interfaces__ = [qt]),
      (Lo.prototype = {
        onOpen: function (e) {
          (F.i("WebSocket opened."), null != this.onopen && this.onopen());
        },
        onMessage: function (e) {
          var t;
          null != this.onmessage &&
            ((t = null),
            (t = this.getBinaryDataForHubot(e)),
            this.onmessage(t));
        },
        getBinaryDataForHubot: function (e) {
          return rs.ofData(e.binaryData);
        },
        onPong: function (e) {
          null != this.onpong && this.onpong(e);
        },
        onError: function (e) {
          (F.e("WebSocket error. event:", e),
            null != this.onerror && this.onerror());
        },
        onClose: function (e) {
          (this.removeAllEventListenersForHubot(),
            F.i(
              "WebSocket closed. " +
                P.string(e.code) +
                " " +
                P.string(e.reason) +
                " " +
                P.string(e.wasClean),
            ),
            null != this.onclose && this.onclose(e.code, e.reason, e.wasClean));
        },
        removeAllEventListenersForHubot: function () {
          (null != this.ws && (this.ws.removeAllListeners(), (this.ws = null)),
            null != this.connection &&
              (this.connection.removeAllListeners(), (this.connection = null)));
        },
        onConnectionClose: function (e, t) {
          this.onClose({ code: e, reason: t });
        },
        close: function () {
          var e;
          ((this.onopen = null),
            (this.onmessage = null),
            (this.onerror = null),
            (this.onclose = null),
            (this.onpong = null),
            this.isClosed() ||
              ((e = this.connection),
              this.onConnectionClose(1e3, "Normal connection closure"),
              e.close()));
        },
        send: function (e) {
          this.isClosed() || this.connection.sendBytes(new gs(e.b.bufferValue));
        },
        ping: function (e) {
          this.connection.ping(e);
        },
        isClosed: function () {
          return (
            null == this.ws ||
            null == this.connection ||
            !this.connection.connected
          );
        },
        __class__: Lo,
      }),
      ((t["puremvc.interfaces.IProxy"] = Xt).__name__ = [
        "puremvc",
        "interfaces",
        "IProxy",
      ]),
      (Xt.prototype = { __class__: Xt }),
      ((t["albero.proxy.AccessTokenResolverProxy"] = Jt).__name__ = [
        "albero",
        "proxy",
        "AccessTokenResolverProxy",
      ]),
      (Jt.__interfaces__ = [Xt]),
      (Jt.prototype = { __class__: Jt }),
      function () {}),
    l =
      (((t["albero.proxy.AccessTokenResolverProxyFactory"] = Po).__name__ = [
        "albero",
        "proxy",
        "AccessTokenResolverProxyFactory",
      ]),
      (Po.newInstance = function () {
        return new Bo("accessTokenResolver");
      }),
      function (e, t) {
        (Bn.call(this),
          (this.proxyName = null != e ? e : l.NAME),
          null != t && this.setData(t));
      }),
    Bo =
      (((t["puremvc.patterns.proxy.Proxy"] = l).__name__ = [
        "puremvc",
        "patterns",
        "proxy",
        "Proxy",
      ]),
      (l.__interfaces__ = [Xt]),
      (l.__super__ = Bn),
      (l.prototype = e(Bn.prototype, {
        getProxyName: function () {
          return this.proxyName;
        },
        setData: function (e) {
          this.data = e;
        },
        getData: function () {
          return this.data;
        },
        onRegister: function () {},
        onRemove: function () {},
        __class__: l,
      })),
      function (e) {
        l.call(this, e);
      }),
    Ho =
      (((t["albero.proxy.AccessTokenResolverProxyForDirectJsHubot"] =
        Bo).__name__ = [
        "albero",
        "proxy",
        "AccessTokenResolverProxyForDirectJsHubot",
      ]),
      (Bo.__interfaces__ = [Jt]),
      (Bo.__super__ = l),
      (Bo.prototype = e(l.prototype, {
        asyncGetAccessToken: function (e, t) {
          var n = this.settings.getAccessToken();
          null != n
            ? t(n)
            : null != (n = null != qs.account ? qs.account.split(":") : null) &&
                2 == n.length
              ? this.api.createAccessToken(n[0], n[1], t)
              : null != (e = null == e ? this.accountLoader.load() : e) &&
                (null != e.accessToken
                  ? (this.settings.setAccessToken(e.accessToken),
                    t(e.accessToken))
                  : this.api.createAccessToken(e.email, e.pass, t));
        },
        __class__: Bo,
      })),
      ((t["albero.proxy.AccountLoaderProxy"] = Zt).__name__ = [
        "albero",
        "proxy",
        "AccountLoaderProxy",
      ]),
      (Zt.__interfaces__ = [Xt]),
      (Zt.prototype = { __class__: Zt }),
      function () {}),
    Go =
      (((t["albero.proxy.AccountLoaderProxyFactory"] = Ho).__name__ = [
        "albero",
        "proxy",
        "AccountLoaderProxyFactory",
      ]),
      (Ho.newInstance = function () {
        var e = x.resolveClass("albero.debug.proxy.AccountLoaderProxyImpl");
        return null != e
          ? x.createInstance(e, ["accountLoader"])
          : new Go("accountLoader");
      }),
      function (e) {
        l.call(this, e);
      }),
    jo =
      (((t["albero.proxy.AccountLoaderProxyForHubot"] = Go).__name__ = [
        "albero",
        "proxy",
        "AccountLoaderProxyForHubot",
      ]),
      (Go.__interfaces__ = [Zt]),
      (Go.__super__ = l),
      (Go.prototype = e(l.prototype, {
        load: function () {
          var i = this,
            t = require("read");
          return (
            t({ prompt: "Email: " }, function (e, n) {
              t({ prompt: "Password: ", silent: !0 }, function (e, t) {
                i.sendNotification("SignIn", new oe(n, t));
              });
            }),
            null
          );
        },
        __class__: Go,
      })),
      function () {
        l.call(this, "broadcast");
      }),
    zo =
      (((t["albero.proxy.AlberoBroadcastProxy"] = jo).__name__ = [
        "albero",
        "proxy",
        "AlberoBroadcastProxy",
      ]),
      (jo.__super__ = l),
      (jo.prototype = e(l.prototype, {
        handleNotification: function (e, t) {
          var n = this;
          switch (
            (F.v("Receive request from server. name:", e, " body:", t), e)
          ) {
            case "notify_add_account_control_request":
              var i = new le(t);
              (this.dataStore.setAccountControlRequest(i),
                this.sendNotification(e, i));
              break;
            case "notify_add_acquaintance":
              var i = new ve(t[1]);
              (this.dataStore.addAcquaintance(i), this.sendNotification(e, i));
              break;
            case "notify_add_acquaintances":
              i = t[1].map(function (e) {
                return new ve(e);
              });
              (k.iter(i, z((j = this.dataStore), j.addAcquaintance)),
                this.api.sendAddAcquaintanceNotifications(i),
                this.sendNotification(e, i));
              break;
            case "notify_add_domain_invite":
              i = new Se(t);
              (this.dataStore.setDomainInvite(i), this.sendNotification(e, i));
              break;
            case "notify_add_favorite_message":
              i = new kt(t);
              (this.dataStore.addFavoriteMessage(i),
                this.sendNotification(e, i));
              break;
            case "notify_add_favorite_talk":
              i = new Ne(t);
              (this.updateTalkStatusByFavoriteTalkEvent(i),
                this.sendNotification(e, i));
              break;
            case "notify_add_friend":
              i = new ve(t[1]);
              (this.dataStore.addFriend(i), this.sendNotification(e, i));
              break;
            case "notify_add_talkers":
              i = new qr(t);
              (null == this.dataStore.getTalk(i.id) &&
                this.sendNotification("notify_add_talkers_including_me", i),
                this.dataStore.setTalk(i),
                this.sendNotification(e, i));
              break;
            case "notify_close_conference":
              i = new he(t);
              (this.dataStore.removeConference(i),
                this.updateTalkStatusByConferenceClosedEvent(i),
                this.sendNotification(e, i));
              break;
            case "notify_conference_participant_join":
              var i = t[3],
                r = t[4],
                o = this.dataStore;
              (null != o.me &&
                null != (o = o.me.id) &&
                null != r &&
                o.high == r.high &&
                o.low == r.low &&
                this.dataStore.addReactedConfereceId(i),
                this.sendNotification(e, t));
              break;
            case "notify_conference_participant_reject":
              var o = t[3];
              (this.dataStore.addReactedConfereceId(o),
                this.sendNotification(e, t));
              break;
            case "notify_create_announcement":
              var r = new Ri(t);
              (this.dataStore.setAnnouncement(r),
                this.keywordWatcher.onAnnouncementCreated(r),
                this.sendNotification(
                  "notify_update_announcement_status",
                  this.newAnnouncementStatus(r),
                ),
                this.sendNotification(e, r));
              break;
            case "notify_create_message":
              var a = new O(t),
                i =
                  (this.dataStore.setMessage(a),
                  this.keywordWatcher.onMessageCreated(a),
                  this.newTalkStatusByMessage(a));
              switch (
                (null != i &&
                  this.sendNotification("notify_update_local_talk_status", i),
                a.type[1])
              ) {
                case 0:
                  var s = a.content,
                    l = null;
                  switch (s.type) {
                    case "delete_talker":
                      l = s.deleted_user_id;
                      break;
                    case "hide_pair_talk":
                      l = s.user_id;
                  }
                  if (null == l) return void this.sendNotification(e, a);
                  for (
                    var u = a.talkId,
                      c = this.dataStore.getQuestions(u, l),
                      _ = 0;
                    _ < c.length;
                  ) {
                    var h = c[_];
                    (++_,
                      (h.closed = !0),
                      this.sendNotification("notify_update_question", h));
                  }
                  this.sendNotification(e, a);
                  break;
                case 4:
                  (null != a.content.file_id &&
                    ((u = this.dataStore.getMessageFileInfos(a.id)),
                    k.iter(u, function (e) {
                      (n.dataStore.setTalkFileInfo(e, !1),
                        n.sendNotification("notify_create_attachment", e));
                    })),
                    this.sendNotification(e, a));
                  break;
                case 5:
                  u = this.dataStore.getMessageFileInfos(a.id);
                  (k.iter(u, function (e) {
                    (n.dataStore.setTalkFileInfo(e, !1),
                      n.sendNotification("notify_create_attachment", e));
                  }),
                    this.sendNotification(e, a));
                  break;
                case 13:
                case 15:
                case 17:
                  var u = this.dataStore.getTalk(a.talkId);
                  (null != u &&
                    ((u = this.dataStore.setQuestionIfLatest(
                      Cr.fromMessageAndTalk(a, u),
                    )),
                    this.sendNotification("notify_update_question", u)),
                    this.sendNotification(e, a));
                  break;
                case 14:
                case 16:
                case 18:
                case 19:
                case 20:
                case 21:
                  var d,
                    u = a.content.in_reply_to,
                    f = this.dataStore.getTalk(a.talkId),
                    m = this.dataStore.getQuestion(u);
                  (null != f &&
                    null != m &&
                    ((f = this.dataStore),
                    (d = a.userId),
                    (f =
                      null != f.me &&
                      null != (f = f.me.id) &&
                      null != d &&
                      f.high == d.high &&
                      f.low == d.low),
                    m.updateByMessage(a, f),
                    this.sendNotification("notify_update_question", m),
                    a.isClosedActionStampMessage()) &&
                    this.dataStore.closeAction(u),
                    this.sendNotification(e, a));
                  break;
                default:
                  this.sendNotification(e, a);
              }
              break;
            case "notify_create_note":
              o = new bo(t);
              (this.dataStore.setNoteRevisionSummary(
                o.noteId,
                o.noteRevisionSummary,
              ),
                this.sendNotification(e, o));
              break;
            case "notify_create_group_talk":
            case "notify_create_pair_talk":
              var r = new qr(t);
              (this.dataStore.setTalk(r),
                this.newTalkStatusByTalk(r),
                this.sendNotification(e, r));
              break;
            case "notify_delete_account_control_request":
              i = D.fromIntOrInt64(t);
              (this.dataStore.removeAccountControlRequest(i),
                this.sendNotification(e, i));
              break;
            case "notify_delete_acquaintance":
              (this.dataStore.removeAcquaintance(t[0], t[1]),
                this.sendNotification(e, t));
              break;
            case "notify_delete_acquaintances":
              for (var p = t[0], g = t[1], y = 0; y < g.length; ) {
                var S = g[y];
                (++y,
                  this.dataStore.removeAcquaintance(p, S),
                  this.sendNotification("notify_delete_acquaintance", [p, S]));
              }
              this.sendNotification(e, t);
              break;
            case "notify_delete_announcement":
              var o = t[0],
                r = t[1],
                i = this.dataStore.getAnnouncement(r),
                i =
                  (null != i && i.markAsDeleted(),
                  this.dataStore.getAnnouncementStatus(o));
              (null != i &&
                i.updateByAnnouncementDeletion(r) &&
                this.dataStore.setAnnouncementStatus(i),
                this.sendNotification(e, t));
              break;
            case "notify_delete_attachment":
              o = new Ae(t);
              (this.dataStore.deleteFile(o), this.sendNotification(e, o));
              break;
            case "notify_delete_domain_invite":
              r = t;
              (this.dataStore.removeDomainInvite(r),
                this.sendNotification(e, r));
              break;
            case "notify_delete_favorite_message":
              i = new Dt(t);
              (this.dataStore.deleteFavoriteMessage(i),
                this.sendNotification(e, i));
              break;
            case "notify_delete_favorite_talk":
              o = new Ne(t);
              (this.updateTalkStatusByFavoriteTalkEvent(o),
                this.sendNotification(e, o));
              break;
            case "notify_delete_friend":
              (this.dataStore.removeFriend(t[0], t[1]),
                this.sendNotification(e, t));
              break;
            case "notify_delete_message":
              ((r = new Re(t)),
                (i = this.dataStore.getMessage(r.messageId)),
                (o =
                  (null != i && i.markAsDeleted(),
                  this.updateTalkStatusOnMessageDelete(r))));
              (null != o &&
                this.sendNotification("notify_update_local_talk_status", o),
                this.sendNotification(e, r));
              break;
            case "notify_delete_note":
              ((i = new So(t)),
                (o =
                  (this.dataStore.setNote(i, null),
                  this.dataStore.setNoteRevisionSummary(i, null),
                  this.dataStore.getNoteLocalEdit(i))));
              (null != o && o.deleted(), this.sendNotification(e, i));
              break;
            case "notify_delete_scheduled_message":
              r = this.dataStore.removeScheduledMessage(t);
              this.sendNotification(e, r);
              break;
            case "notify_delete_stampset":
              o = t;
              (this.dataStore.removeStampsetInfo(o),
                this.sendNotification(e, o));
              break;
            case "notify_delete_talk":
              i = t;
              (this.dataStore.removeTalk(i),
                this.dataStore.removeTalkStatus(i),
                this.dataStore.removeConferencesOnLeaveTalk(i),
                this.sendNotification("brand_badge_changed"),
                this.sendNotification(e, t));
              break;
            case "notify_disable_push_notification":
              r = new Je(t, !1);
              (this.updateTalkStatusByTalkPushNotificationEvent(r),
                this.sendNotification(e, r));
              break;
            case "notify_enable_push_notification":
              o = new Je(t, !0);
              (this.updateTalkStatusByTalkPushNotificationEvent(o),
                this.sendNotification(e, o));
              break;
            case "notify_flow_notification_badge":
              i = new be(t);
              (this.dataStore.setFlowNotificationBadge(i.domainId, i),
                this.sendNotification(e, i));
              break;
            case "notify_join_account_control_group":
              r = new ae(t);
              (this.dataStore.setAccountControlGroup(r),
                this.dataStore.removeAccountControlRequests(),
                this.sendNotification("brand_badge_changed"),
                this.sendNotification(e, r));
              break;
            case "notify_join_domain":
              var v = this.dataStore.setDomainIfLatest(new ye(t));
              (ye.isAllowReadAnnouncement(v) &&
                Xa.delay(function () {
                  n.api.getAnnouncementStatus(v.id);
                }, 500),
                this.sendNotification(e, [null, v]));
              break;
            case "notify_leave_account_control_group":
              o = t;
              (this.dataStore.removeAccountControlGroup(o),
                this.sendNotification(e, o));
              break;
            case "notify_leave_domain":
              i = t;
              (this.dataStore.removeDomain(i),
                this.dataStore.removeConferencesOnLeaveDomain(i),
                this.sendNotification(e, i));
              break;
            case "notify_lock_note":
              ((r = new Ft(t)), (o = this.dataStore.getNote(r.noteId)));
              (null != o &&
                o.lock(r.revision, r.userId, r.deviceId, r.expiredAt),
                this.sendNotification(e, r));
              break;
            case "notify_open_conference":
              i = new he(t);
              (this.dataStore.addConference(i), this.sendNotification(e, i));
              break;
            case "notify_processed_scheduled_message":
              ((o = t.scheduled_message_id),
                (r = t.result_code),
                (i = this.dataStore.removeScheduledMessage(o)));
              (0 != r &&
                this.sendNotification("error_occurred", {
                  localErrorType: R.SendingScheduledMessageFailed,
                }),
                this.sendNotification(e, i));
              break;
            case "notify_schedule_message":
              ((o = t[0]),
                (r = t[1]),
                (i = this.dataStore.getScheduledMessage(r)));
              (null == i &&
                ((i = fo.fromNotifyScheduleMessage({ id: r, talkId: o })),
                this.dataStore.setScheduledMessage(i)),
                this.sendNotification(e, i));
              break;
            case "notify_unlock_note":
              ((r = new Gt(t)), (o = this.dataStore.getNote(r.noteId)));
              (null != o && o.unlock(r.revision, r.userId),
                this.sendNotification(e, r));
              break;
            case "notify_update_account_control_group_partially":
              i = new se(t);
              (this.dataStore.updateAccountControlGroup(i),
                this.sendNotification(e, i));
              break;
            case "notify_update_announcement_status":
              ((o = new ce(t)),
                (r = this.dataStore.getAnnouncementStatus(o.domainId)));
              ((r =
                null == r
                  ? this.assureAnnouncementStatus(o.domainId)
                  : r).updateReadWithStatusUpdate(o),
                this.dataStore.setAnnouncementStatus(r),
                this.sendNotification(e, o));
              break;
            case "notify_update_department_tree":
              i = t;
              (this.dataStore.clearDomainDepartment(i),
                this.api.getDepartmentTree(i),
                this.sendNotification(e, i));
              break;
            case "notify_update_department_users":
              r = t;
              (this.dataStore.clearDomainDepartmentUsers(r) &&
                this.sendNotification("department_user_count_cleared", r),
                this.api.getUsers(
                  r,
                  this.dataStore.getFriendsOrAcquaintances(r).map(function (e) {
                    return e.id;
                  }),
                ),
                this.api.getMe(),
                this.sendNotification(e, r));
              break;
            case "notify_update_domain":
              var o = new ye(t),
                i = this.dataStore.getDomain(o.id),
                r = this.dataStore.setDomainIfLatest(o),
                o =
                  (null != i &&
                    ye.isAllowReadAnnouncement(i) &&
                    !ye.isAllowReadAnnouncement(r) &&
                    (this.dataStore.removeAnnouncementStatus(i.id),
                    ((o = new ce()).domainId = r.id),
                    this.sendNotification(
                      "notify_update_announcement_status",
                      o,
                    )),
                  !ye.isAllowReadAnnouncement(i) &&
                    ye.isAllowReadAnnouncement(r) &&
                    this.api.getAnnouncementStatus(r.id),
                  ye.isChangedContractTreeEnabled(i, r) &&
                    !ye.isDepartmentEnabled(r) &&
                    ((o = r.id),
                    this.dataStore.clearDomainDepartment(o),
                    this.dataStore.clearUsersDepartments(o)),
                  r.directAppsSetting.enabled.filter(function (e) {
                    return null == n.dataStore.getDirectApp(e);
                  }));
              (0 != o.length &&
                this.api.getDirectApps(o).catch(function (e) {
                  F.v("getDirectApps error:", e);
                }),
                this.sendNotification(e, [i, r]));
              break;
            case "notify_update_domain_builtin_stampsetting":
              ((o = t[0]), (i = this.dataStore.getDomain(o)));
              null != i &&
                ((r = new Li(t[1])),
                i.updateBuiltinStampSetting(r),
                this.dataStore.setDomainIfLatest(i),
                this.sendNotification(e, o));
              break;
            case "notify_update_domain_stampsetting":
              ((r = t[0]), (i = this.dataStore.getDomain(r)));
              null != i &&
                ((o = i.stampsetSetting),
                i.updateStampsetSetting(new Kr(t[1])),
                this.dataStore.setDomainIfLatest(i),
                (i = i.stampsetSetting),
                (r = new nt(r, o, i)).isValid()) &&
                this.sendNotification(e, r);
              break;
            case "notify_update_domain_users":
              for (var T = t, I = 0; I < T.length; ) {
                var E = T[I],
                  w = (++I, this.dataStore.getDomain(E));
                null != w &&
                  w.setting.contactTreeEnabled &&
                  this.dataStore.clearDomainDepartmentUsers(E) &&
                  this.sendNotification("department_user_count_cleared", E);
              }
              this.sendNotification(e, t);
              break;
            case "notify_delete_talker":
            case "notify_update_group_talk":
              o = new qr(t);
              (this.dataStore.setTalk(o), this.sendNotification(e, o));
              break;
            case "notify_update_message_reactions":
              (k.iter(t, function (e) {
                e = e[1];
                (n.dataStore.removeMessageReactionSummary(e),
                  n.dataStore.removeMessageReactionUsers(e));
              }),
                this.sendNotification(e, t));
              break;
            case "notify_update_note_partially":
              null == t
                ? F.e("notePartiallyUpdate is null.")
                : null == t.note_id
                  ? F.e("noteId is null.")
                  : ((i = new So(t.note_id)),
                    null != (r = this.dataStore.getNote(i)) &&
                      null != t.note_revision_summary &&
                      r.compareRevision(t.note_revision_summary.revision) <=
                        -1 &&
                      this.dataStore.setNote(i, null),
                    null == t.note_revision_summary && null == t.setting
                      ? F.e("noteSetting and noteRevisionSummary are null.")
                      : (null != t.setting &&
                          this.sendNotification(
                            "notify_update_note_for_setting",
                            new Pt(t),
                          ),
                        null != t.note_revision_summary &&
                          ((o = new Bt(t)),
                          null ==
                            (r = this.dataStore.getNoteRevisionSummary(i)) ||
                            r.compareRevision(o.getRevision()) < 0) &&
                          (null != (r = this.dataStore.getNoteLocalEdit(i)) &&
                            r.conflict(),
                          this.dataStore.setNoteRevisionSummary(
                            i,
                            o.noteRevisionSummary,
                          ),
                          this.sendNotification(
                            "notify_update_note_for_revision",
                            new Bt(t),
                          )),
                        this.sendNotification(e, new Lt(t))));
              break;
            case "notify_update_read_statuses":
              var N = new Fe(t),
                r =
                  (k.iter(N.messageIds, function (e) {
                    e = n.dataStore.getMessage(e);
                    null != e && e.addReadUsersUntillMaxCount(N.readUserIds);
                  }),
                  this.updateTalkStatus(N));
              (null != r &&
                this.sendNotification("notify_update_local_talk_status", r),
                this.sendNotification(e, N));
              break;
            case "notify_update_stampset":
              i = new Ve(t);
              (this.dataStore.updateStampsetInfo(i),
                this.sendNotification(e, i));
              break;
            case "notify_update_talk_status":
              ((o = new et(t)), (r = this.dataStore.getTalkStatus(o.talkId)));
              (null != r &&
                r.update(o) &&
                (this.dataStore.setTalkStatus(r),
                this.sendNotification("notify_update_local_talk_status", r)),
                this.sendNotification(e, o));
              break;
            case "notify_update_user":
              if (t instanceof Array && null == t.__enum__) {
                i = new ve(t[1]);
                this.dataStore.setUserIfLatest(null, i);
              } else {
                for (
                  var r = new fr(t),
                    A = (F.d("Current user updated. user:", r), this.dataStore),
                    b = ((A.me = r), 0),
                    M = A.getDomains();
                  b < M.length;
                ) {
                  var C = M[b];
                  (++b,
                    A.sendNotification(
                      "notify_update_user",
                      A.me.toDomainUser(C.id),
                    ));
                }
                (A.sendNotification("current_user_changed", r),
                  A.storageForUsers.setDirtyFlag());
              }
              break;
            default:
              this.sendNotification(e, t);
          }
        },
        updateTalkStatusByFavoriteTalkEvent: function (e) {
          var t = this.dataStore.getTalkStatus(e.talkId);
          if (null == t) {
            var n = this.dataStore.getTalk(e.talkId);
            if (null == n)
              return void F.e("talk not found on change favorite talk.");
            (((t = new Ze()).id = e.talkId),
              (t.talkOrderingTimestamp = n.updatedAt));
          }
          (null == t.favoriteVersion ||
            e.favoriteVersion > t.favoriteVersion) &&
            ((t.favoriteVersion = e.favoriteVersion),
            (t.orderInFavorites = e.orderInFavorites),
            this.dataStore.setTalkStatus(t));
        },
        updateTalkStatusByConferenceClosedEvent: function (e) {
          var t,
            n = this.dataStore.getTalkStatus(e.talkId);
          null != n &&
            null != n.maxMessage &&
            ((t = n.maxMessage.id), (e = e.messageId), null != t) &&
            null != e &&
            t.high == e.high &&
            t.low == e.low &&
            ((n.maxMessage.content.closed = !0),
            this.dataStore.setTalkStatus(n));
        },
        updateTalkStatusByTalkPushNotificationEvent: function (e) {
          var t = this.dataStore.getTalkStatus(e.talkId);
          if (null == t) {
            var n = this.dataStore.getTalk(e.talkId);
            if (null == n)
              return void F.e("talk not found on change favorite talk.");
            (((t = new Ze()).id = e.talkId),
              (t.talkOrderingTimestamp = n.updatedAt));
          }
          t.updateByTalkPushNotificationEvent(e) &&
            this.dataStore.setTalkStatus(t);
        },
        newTalkStatusByTalk: function (e) {
          var t = this.dataStore.getTalkStatus(e.id);
          return (
            null == t && ((t = new Ze()).id = e.id),
            null == t.talkOrderingTimestamp &&
              (t.talkOrderingTimestamp = e.updatedAt),
            this.dataStore.setTalkStatus(t),
            t
          );
        },
        newTalkStatusByMessage: function (e) {
          var t = this.dataStore.getTalkStatus(e.talkId),
            n = (null == t && ((t = new Ze()).id = e.talkId), this.dataStore),
            i = e.userId;
          return (
            (n =
              null != n.me &&
              null != (n = n.me.id) &&
              null != i &&
              n.high == i.high &&
              n.low == i.low),
            t.updateByMessage(e, n),
            this.dataStore.setTalkStatus(t),
            t
          );
        },
        updateTalkStatusOnMessageDelete: function (e) {
          var t = this.dataStore.getTalkStatus(e.talkId);
          return (
            null != t &&
              t.updateByMessageDeletion(e) &&
              this.dataStore.setTalkStatus(t),
            t
          );
        },
        updateTalkStatus: function (e) {
          var t,
            n = this;
          return k.exists(e.readUserIds, function (e) {
            var t = n.dataStore;
            return (
              null != t.me &&
              null != (t = t.me.id) &&
              null != e &&
              t.high == e.high &&
              t.low == e.low
            );
          })
            ? null == (t = this.dataStore.getTalkStatus(e.talkId))
              ? (F.e("talkStatus is not found."), null)
              : (t.updateByMessageReadStatusesUpdate(e) &&
                  this.dataStore.setTalkStatus(t),
                t)
            : null;
        },
        newAnnouncementStatus: function (e) {
          var t = this.dataStore.getAnnouncementStatus(e.domainId);
          return (
            null == t && ((t = new ue()).domainId = e.domainId),
            t.updateByAnnouncement(e),
            this.dataStore.setAnnouncementStatus(t),
            t
          );
        },
        assureAnnouncementStatus: function (e) {
          var t = this.dataStore.getAnnouncementStatus(e);
          return (
            null == t &&
              ((t = ue.fromDomainId(e)),
              this.dataStore.setAnnouncementStatus(t)),
            t
          );
        },
        __class__: jo,
      })),
      function () {
        l.call(this, "api");
      }),
    Yo =
      (((t["albero.proxy.AlberoServiceProxy"] = zo).__name__ = [
        "albero",
        "proxy",
        "AlberoServiceProxy",
      ]),
      (zo.__super__ = l),
      (zo.prototype = e(l.prototype, {
        getOSString: function () {
          return "bot";
        },
        createAccessToken: function (e, t, n) {
          function i(e) {
            (F.v("access token:", e), o.settings.setAccessToken(e), n(e));
          }
          function r(e) {
            o.sendNotification("Url", bi.FORWARD(g.error));
          }
          var o = this,
            a = e.split("$"),
            s = this.settings.getOs(),
            l = this.settings.getIDFV();
          (-1 == e.indexOf("@") && 2 == a.length
            ? this.wsApi.createAccessTokenById(a[1], a[0], t, l, s)
            : this.wsApi.createAccessToken(e, t, l, s)
          ).then(i, r);
        },
        authorizeDevice: function (e, t) {
          var n = this,
            i = this.settings.getIDFV();
          this.wsApi.authorizeDevice(e, i).then(
            function (e) {
              t();
            },
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        createSession: function (e, n) {
          var i = this,
            t = this.getOSString();
          (this.apiCaller.onSessionClear(),
            this.wsApi.createSession(e, "1.128", t).then(
              function (e) {
                var t = W.nowAsInt64();
                e.passwordExpiration.isExpired(t)
                  ? i.sendNotification("password_expiration_overed")
                  : (e.passwordExpiration.needWarning(
                      t,
                      i.settings.getPasswordWarningSkipUntil(),
                    ) &&
                      (i.settings.setPasswordWarningSkipUntil(
                        W.oneDayAfterAsInt64(),
                      ),
                      i.sendNotification("password_expiration_warned")),
                    i.apiCaller.onSessionCreated(),
                    n(e),
                    i.apiCaller.fireCallbacksWaitingForSession());
              },
              function (e) {
                if (null != e) {
                  if (401 == e.code && "expired password" == e.message)
                    return void i.sendNotification(
                      "password_expiration_overed",
                    );
                  if (401 == e.code && "deleted account" == e.message)
                    return void i.sendNotification("SignOut");
                  if (401 == e.code && "unauthorized device" == e.message)
                    return void i.sendNotification("Device", Kn.AUTHORIZE);
                  (((400 == e.code && "older version specified" == e.message) ||
                    (401 == e.code && "bad token" == e.message) ||
                    (401 == e.code && "invalid token" == e.message)) &&
                    i.settings.clearAccessToken(),
                    i.apiCaller.handleServerErrorDefault(e));
                }
                i.sendNotification("Url", bi.FORWARD(g.error));
              },
            ));
        },
        startNotification: function () {
          var t = this;
          this.wsApi.startNotification().then(
            function (e) {
              e ||
                (t.dataStore.clear(!0),
                t.sendNotification("start_notification_failed"));
            },
            function (e) {
              t.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        resetNotification: function (e) {
          var t = this;
          this.wsApi.resetNotification().then(e, function (e) {
            t.apiCaller.handleServerErrorDefault(e);
          });
        },
        sendDomainNotifications: function (e) {
          for (var t = 0; t < e.length; ) {
            var n = e[t];
            (++t, this.sendNotification("notify_update_domain", [null, n]));
          }
        },
        sendAnnouncementStatusNotification: function (e) {
          var t = new ce();
          ((t.domainId = e.domainId),
            this.sendNotification("notify_update_announcement_status", t));
        },
        sendAddFriendNotifications: function (e) {
          for (var t = 0; t < e.length; ) {
            var n = e[t];
            (++t, this.sendNotification("notify_add_friend", n));
          }
        },
        sendAddAcquaintanceNotifications: function (e) {
          for (var t = 0; t < e.length; ) {
            var n = e[t];
            (++t, this.sendNotification("notify_add_acquaintance", n));
          }
        },
        sendTalkNotification: function (e) {
          var t =
            e.type == Xr.PairTalk
              ? "notify_create_pair_talk"
              : "notify_create_group_talk";
          this.sendNotification(t, e);
        },
        sendTalkNotifications: function (e) {
          k.iter(e, z(this, this.sendTalkNotification));
        },
        sendTalkStatusNotifications: function (e) {
          for (var t = 0; t < e.length; ) {
            var n = e[t];
            (++t, this.sendNotification("notify_update_local_talk_status", n));
          }
        },
        sendAddDomainInviteNotifications: function (e) {
          for (var t = 0; t < e.length; ) {
            var n = e[t];
            (++t, this.sendNotification("notify_add_domain_invite", n));
          }
        },
        sendAddAccountControlRequestNotifications: function (e) {
          for (var t = 0; t < e.length; ) {
            var n = e[t];
            (++t,
              this.sendNotification("notify_add_account_control_request", n));
          }
        },
        sendJoinAccountControlGroupNotification: function (e) {
          this.sendNotification("notify_join_account_control_group", e);
        },
        deleteSession: function () {
          (this.settings.clearSelectedStampsetType(),
            this.settings.clearDomainSelection(),
            this.settings.clearAccessToken(),
            this.apiCaller.restart());
        },
        updateUser: function (t, e, n, i, r) {
          function o(e) {
            s.sendNotification("update_user_errored", s.dataStore.me);
          }
          function a(e) {
            s.wsApi.updateUser(t, e, i, r).then(
              function (e) {
                for (
                  var t = s.dataStore, n = ((t.me = e), 0), i = t.getDomains();
                  n < i.length;
                ) {
                  var r = i[n];
                  (++n,
                    t.sendNotification(
                      "notify_update_user",
                      t.me.toDomainUser(r.id),
                    ));
                }
                (t.sendNotification("current_user_changed", e),
                  t.storageForUsers.setDirtyFlag(),
                  s.sendNotification("update_user_responsed", s.dataStore.me));
              },
              function (e) {
                (o(), s.apiCaller.handleServerErrorDefault(e));
              },
            );
          }
          var s = this;
          null != n
            ? a(n)
            : null != e
              ? this.uploadFile(
                  e,
                  null,
                  u.PROFILE_IMAGE,
                  function (e) {
                    a(e.get_url);
                  },
                  o,
                )
              : a();
        },
        updateProfile: function (e) {
          var o = this;
          this.wsApi.updateProfile(e).then(
            function (e) {
              for (
                var t = o.dataStore, n = ((t.me = e), 0), i = t.getDomains();
                n < i.length;
              ) {
                var r = i[n];
                (++n,
                  t.sendNotification(
                    "notify_update_user",
                    t.me.toDomainUser(r.id),
                  ));
              }
              (t.sendNotification("current_user_changed", e),
                t.storageForUsers.setDirtyFlag(),
                o.sendNotification("update_profile_responsed", o.dataStore.me));
            },
            function (e) {
              (o.sendNotification("update_profile_errored", o.dataStore.me),
                o.apiCaller.handleServerErrorDefault(e));
            },
          );
        },
        addFriend: function (e) {
          var t = this,
            n = this.settings.getSelectedDomainId();
          this.wsApi.addFriend(n, e.id).then(
            function (e) {},
            function (e) {
              t.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        deleteFriend: function (e) {
          var t = this,
            n = this.settings.getSelectedDomainId();
          this.wsApi.deleteFriend(n, e.id).then(
            function (e) {},
            function (e) {
              t.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getFriends: function (i) {
          var r = this;
          this.wsApi.getFriends().then(
            function (e) {
              for (var t = 0; t < e.length; ) {
                var n = e[t];
                (++t,
                  k.iter(n.domainUsers, z((j = r.dataStore), j.addFriend)),
                  r.sendAddFriendNotifications(n.domainUsers));
              }
              null != i && i();
            },
            function (e) {
              r.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getAcquaintances: function (i) {
          var r = this;
          this.wsApi.getAcquaintances().then(
            function (e) {
              for (var t = 0; t < e.length; ) {
                var n = e[t];
                (++t,
                  k.iter(
                    n.domainUsers,
                    z((j = r.dataStore), j.addAcquaintance),
                  ),
                  r.sendAddAcquaintanceNotifications(n.domainUsers));
              }
              null != i && i();
            },
            function (e) {
              r.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getAllUsers: function (e, t) {
          var n = this;
          null == t && (t = this.settings.getSelectedDomainId());
          this.wsApi.getAllUsers(t, [2, 4, 8], 40, e).then(
            function (e) {
              (n.dataStore.setUsersIfLatest(t, e.contents),
                n.sendNotification("get_domain_users_responsed", [
                  t,
                  e.marker,
                  e.nextMarker,
                  e.contents,
                ]));
            },
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        searchDomainUsers: function (t, e) {
          var n = this,
            i = this.settings.getSelectedDomainId();
          this.wsApi.searchDomainUsers(i, t, [2, 4, 8], 40, e).then(
            function (e) {
              (n.dataStore.setUsersIfLatest(i, e.contents),
                n.sendNotification("get_domain_users_responsed", [
                  i,
                  e.marker,
                  e.nextMarker,
                  e.contents,
                  t,
                ]));
            },
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getUsers: function (t, n) {
          var i = this,
            e = n.slice(0, 100);
          0 != e.length &&
            this.wsApi.getUsers(t, e).then(
              function (e) {
                (i.dataStore.setUsersIfLatest(t, e),
                  i.sendNotification("get_users_responsed", [t, e]));
                e = n.slice(100);
                i.getUsers(t, e);
              },
              function (e) {
                i.apiCaller.handleServerErrorDefault(e);
              },
            );
        },
        getProfile: function (e, t) {
          var n = this;
          this.wsApi.getProfile(e, t).then(
            function (e) {
              n.sendNotification("get_profile_responsed", e);
            },
            function (e) {
              403 == e.code && "invalid user" == e.message
                ? null != e.detail &&
                  null != e.detail.invalid_user_ids &&
                  n.sendNotification(
                    "get_profile_errored",
                    e.detail.invalid_user_ids,
                  )
                : n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getUserIdentifiers: function (t, n, i) {
          var e,
            r = this;
          0 == n.length
            ? null != i && i()
            : ((e = n.slice(0, 100)),
              this._getUserIdentifiers(t, e, function () {
                var e = n.slice(100);
                r.getUserIdentifiers(t, e, i);
              }));
        },
        _getUserIdentifiers: function (a, s, l) {
          var u = this;
          this.wsApi.getUserIdentifiers(a, s).then(
            function (e) {
              for (
                var e = e.contents,
                  n = new B(),
                  t =
                    (k.iter(e, function (e) {
                      var t = e.userId,
                        t = "_" + t.high + "_" + t.low;
                      null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e);
                    }),
                    0),
                  i = s.length;
                t < i;
              ) {
                var r,
                  o = t++,
                  o = u.dataStore.getUser(a, s[o]);
                null != o &&
                  ((r = "_" + (r = o.id).high + "_" + r.low),
                  null != (r = null != Y[r] ? n.getReserved(r) : n.h[r]) &&
                    o.setUserIdentifier(r),
                  (o.__uid_caching = !0));
              }
              null != l && l();
            },
            function (e) {
              for (var t = 0, n = s.length; t < n; ) {
                var i = t++,
                  i = u.dataStore.getUser(a, s[i]);
                null != i && (i.__uid_caching = null);
              }
              null != l && l();
            },
          );
        },
        getAllUserIdentifiers: function (n) {
          for (
            var i,
              r,
              o = this,
              a = null,
              e = this.dataStore.getUsers(),
              t = (null == e && (e = []), 0);
            t < e.length;
          ) {
            var s,
              l,
              u = e[t];
            (++t,
              null == u.__uid_caching &&
                ((u.__uid_caching = !1),
                (s = "_" + (s = u.domainId).high + "_" + s.low),
                null == a && (a = new B()),
                null == (l = null != Y[s] ? a.getReserved(s) : a.h[s]) &&
                  ((l = []), null != Y[s] ? a.setReserved(s, l) : (a.h[s] = l)),
                l.push(u.id)));
          }
          (null == a
            ? n
            : ((i = a.keys()),
              (r = null),
              (r = function () {
                var e, t;
                i.hasNext()
                  ? ((t = i.next()),
                    (e = D.makeFromIdStr(t)),
                    (t = null != Y[t] ? a.getReserved(t) : a.h[t]),
                    o.getUserIdentifiers(e, t, r))
                  : n();
              })))();
        },
        getAllUserEmails: function (e) {
          (F.w(
            "'getAllUserEmails' is deprecated, please use 'getAllUserIdentifiers' instead",
          ),
            this.getAllUserIdentifiers(e));
        },
        getDepartmentTree: function (t) {
          var n = this;
          this.wsApi.getDepartmentTree(t).then(
            function (e) {
              k.iter(e.departments, function (e) {
                (e.isRoot() && n.dataStore.setRootDepartmentId(t, e.id),
                  n.dataStore.setDepartment(e));
              });
              e = n.settings.getSelectedDepartmentId(t);
              (null != e &&
                null == n.dataStore.getDepartment(e) &&
                n.settings.setSelectedDepartmentId(t, null),
                n.sendNotification("get_department_tree_responsed", [t]));
            },
            function (e) {
              403 == e.code && "forbidden" == e.message
                ? n.sendNotification("get_department_tree_canceled", [t])
                : n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getDepartmentUsers: function (e, t) {
          this._getDepartmentUsers(e, t, null);
        },
        getDepartmentAllUsers: function (a) {
          var s,
            l,
            u = this,
            e = this.dataStore.getDepartment(a);
          null != e && null != e.userIds
            ? this.sendNotification("department_user_ids_prepared", e.userIds)
            : ((s = new B()),
              (l = function (e, t, n) {
                if (
                  (k.iter(n, function (e) {
                    var t = e.id,
                      t = "_" + t.high + "_" + t.low,
                      e = e.id;
                    null != Y[t] ? s.setReserved(t, e) : (s.h[t] = e);
                  }),
                  null == t)
                ) {
                  for (
                    var i = [], n = s.arrayKeys(), r = new is(s, n);
                    r.hasNext();
                  ) {
                    var o = r.next();
                    i.push(o);
                  }
                  (F.v(
                    "department all users prepared. [departmentId, users_count]:",
                    a,
                    i.length,
                  ),
                    u.dataStore.setDepartmentUsers(a, i),
                    u.sendNotification("department_user_ids_prepared", i));
                } else u._getDepartmentUsers(a, t, l);
              }),
              this._getDepartmentUsers(a, null, l));
        },
        _getDepartmentUsers: function (t, e, n) {
          var i = this,
            r = this.settings.getSelectedDomainId(),
            o = this.dataStore.getDepartment(t),
            o = null != o && null != o.childrenIds;
          this.wsApi.getDepartmentUsers(r, t, o, 100, e).then(
            function (e) {
              (i.dataStore.setUsersIfLatest(r, e.contents),
                i.sendNotification("get_department_users_responsed", [
                  r,
                  e.marker,
                  e.nextMarker,
                  e.contents,
                  t,
                ]),
                null != n && n(e.marker, e.nextMarker, e.contents));
            },
            function (e) {
              403 == e.code && "forbidden" == e.message
                ? i.sendNotification("get_department_users_canceled", [r])
                : i.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getMe: function () {
          var o = this;
          this.wsApi.getMe().then(
            function (e) {
              for (
                var t = o.dataStore, n = ((t.me = e), 0), i = t.getDomains();
                n < i.length;
              ) {
                var r = i[n];
                (++n,
                  t.sendNotification(
                    "notify_update_user",
                    t.me.toDomainUser(r.id),
                  ));
              }
              (t.sendNotification("current_user_changed", e),
                t.storageForUsers.setDirtyFlag(),
                o.sendNotification("get_me_responsed"));
            },
            function (e) {
              o.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getDepartmentUserCount: function (e, t) {
          t = t.filter(z((j = this.dataStore), j.isUserCountLoadingRequired));
          this._getDepartmentUserCount(e, t);
        },
        _getDepartmentUserCount: function (t, n) {
          var i = this,
            r = n.slice(0, 100);
          0 != r.length &&
            (k.iter(r, function (e) {
              i.dataStore.setUserCountLoading(e, !0);
            }),
            this.wsApi.getDepartmentUserCount(t, r).then(
              function (e) {
                k.iter(
                  e.departments,
                  z((j = i.dataStore), j.setDepartmentUserCount),
                );
                e = n.slice(100);
                (i.sendNotification("get_department_user_count_responsed", {
                  completed: r,
                  remaining: e,
                }),
                  i._getDepartmentUserCount(t, e));
              },
              function (e) {
                (k.iter(r, function (e) {
                  i.dataStore.setUserCountLoading(e, !1);
                }),
                  i.apiCaller.handleServerErrorDefault(e));
              },
            ));
        },
        getDomains: function (t) {
          var n = this;
          this.wsApi.getDomains().then(
            function (e) {
              k.iter(e, z((j = n.dataStore), j.setDomainIfLatest));
              e = n.dataStore.getDomains();
              (n.sendDomainNotifications(e), null != t && t());
            },
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        leaveDomain: function (e) {
          var t = this;
          this.wsApi.leaveDomain(e.id).then(
            function (e) {
              t.dataStore.removeDomain(e);
            },
            function (e) {
              t.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getDomainInvites: function (t) {
          var e =
            null == t
              ? function (e) {}
              : function (e) {
                  t();
                };
          this._getDomainInvites(e);
        },
        _getDomainInvites: function (t) {
          var n = this;
          this.wsApi.getDomainInvites().then(
            function (e) {
              (k.iter(e, z((j = n.dataStore), j.setDomainInvite)),
                n.sendAddDomainInviteNotifications(e),
                t(e));
            },
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        acceptDomainInvite: function (e) {
          var t = this;
          this.wsApi.acceptDomainInvite(e).then(
            function (e) {},
            function (e) {
              t.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        deleteDomainInvite: function (e) {
          var t = this;
          this.wsApi.deleteDomainInvite(e).then(
            function (e) {},
            function (e) {
              t.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getAccountControlRequests: function (t) {
          var n = this;
          this.wsApi.getAccountControlRequests().then(
            function (e) {
              (k.iter(e, z((j = n.dataStore), j.setAccountControlRequest)),
                n.sendAddAccountControlRequestNotifications(e),
                null != t && t());
            },
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        acceptAccountControlRequest: function (e) {
          var t = this;
          null != e &&
            null != (e = this.dataStore.getAccountControlRequest(e)) &&
            this.wsApi.acceptAccountControlRequest(e).then(
              function (e) {},
              function (e) {
                t.apiCaller.handleServerErrorDefault(e);
              },
            );
        },
        rejectAccountControlRequest: function (e) {
          var t = this;
          null != e &&
            null != (e = this.dataStore.getAccountControlRequest(e)) &&
            this.wsApi.rejectAccountControlRequest(e).then(
              function (e) {},
              function (e) {
                t.apiCaller.handleServerErrorDefault(e);
              },
            );
        },
        getJoinedAccountControlGroup: function (t) {
          var n = this;
          this.wsApi.getJoinedAccountControlGroup().then(
            function (e) {
              (k.iter(e, z((j = n.dataStore), j.setAccountControlGroup)),
                1 == e.length &&
                  n.sendJoinAccountControlGroupNotification(e[0]),
                null != t && t());
            },
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getTalks: function (t) {
          var n = this;
          this.wsApi
            .getTalks()
            .then(function (e) {
              return (
                k.iter(e, z((j = n.dataStore), j.setTalk)),
                n.wsApi.getTalkStatuses()
              );
            })
            .then(
              function (e) {
                k.iter(e, z((j = n.dataStore), j.setTalkStatus));
                ((e = n.dataStore.getTalks()),
                  n.sendTalkNotifications(e),
                  (e = n.dataStore.getTalkStatuses()));
                (n.sendTalkStatusNotifications(e), null != t && t());
              },
              function (e) {
                n.apiCaller.handleServerErrorDefault(e);
              },
            );
        },
        getReadStatus: function (e, t) {
          var n = this;
          this.wsApi.getReadStatus(e, t).then(
            function (e) {
              n.sendNotification("notify_get_message_status", e);
            },
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        createGroupTalk: function (e, t, n, i) {
          var r = this;
          (null == e && (e = this.settings.getSelectedDomainId()),
            0 != t.length &&
              this.wsApi.createGroupTalk(e, t, n).then(
                function (e) {
                  (r.dataStore.setTalk(e),
                    r.sendNotification("create_group_talk_complete", e),
                    null != i &&
                      r.sendNotification("SelectTalk", gi.SelectTalk(i, e)));
                },
                function (e) {
                  var t;
                  (r.sendNotification("create_group_talk_fail"),
                    null != e &&
                    500 == e.code &&
                    null != e.detail &&
                    "failure in adding some users" == e.message
                      ? ((t = _a.toLocalError(R.AddingSomeUsersFailure, e)),
                        r.sendNotification("error_occurred", t))
                      : null != e &&
                          502 == e.code &&
                          null != e.detail &&
                          "talks" == e.detail.limit_target
                        ? ((t = e.detail.limit_max),
                          (t = _a.toLocalError(R.TalksCountOvered(t), e)),
                          r.sendNotification("error_occurred", t))
                        : null != e &&
                            502 == e.code &&
                            null != e.detail &&
                            "talkers" == e.detail.limit_target
                          ? ((t = e.detail.limit_max),
                            (t = _a.toLocalError(R.TalkersCountOvered(t), e)),
                            r.sendNotification("error_occurred", t))
                          : r.apiCaller.handleServerErrorDefault(e));
                },
              ));
        },
        createPairTalk: function (e, t, n) {
          var i = this;
          (null == e && (e = this.settings.getSelectedDomainId()),
            this.wsApi.createPairTalk(e, t).then(
              function (e) {
                (i.dataStore.setTalk(e),
                  i.sendNotification("create_pair_talk_complete", e),
                  null != n &&
                    i.sendNotification("SelectTalk", gi.SelectTalk(n, e)));
              },
              function (e) {
                var t;
                (i.sendNotification("create_pair_talk_fail"),
                  null != e &&
                  502 == e.code &&
                  null != e.detail &&
                  "talks" == e.detail.limit_target
                    ? ((t = e.detail.limit_max),
                      (t = _a.toLocalError(R.TalksCountOvered(t), e)),
                      i.sendNotification("error_occurred", t))
                    : i.apiCaller.handleServerErrorDefault(e));
              },
            ));
        },
        updateGroupTalk: function (e, t, n, i) {
          var r = this;
          this.wsApi.updateGroupTalk(e.id, t, n, i).then(
            function (e) {
              (r.dataStore.setTalk(e),
                r.sendNotification("update_group_talk_complete", e));
            },
            function (e) {
              (r.sendNotification("update_group_talk_fail"),
                r.apiCaller.handleServerErrorDefault(e));
            },
          );
        },
        addTalkers: function (e, t, n) {
          var i,
            r,
            o = this;
          e.type == Xr.PairTalk
            ? ((i = e.userIds.concat(t)),
              D.remove(i, this.dataStore.me.id),
              (r = this.settings.getSelectedDomainId()),
              this.createGroupTalk(r, i, Qr.createDefault(), n))
            : 0 != t.length &&
              this.wsApi.addTalkers(e.id, t).then(
                function (e) {
                  (o.dataStore.setTalk(e),
                    o.sendNotification("notify_add_talkers", e),
                    o.sendNotification("add_talkers_succeeded", e));
                },
                function (e) {
                  var t;
                  (o.sendNotification("add_talkers_failed", e),
                    null != e &&
                    500 == e.code &&
                    null != e.detail &&
                    "failure in adding some users" == e.message
                      ? ((t = _a.toLocalError(R.AddingSomeUsersFailure, e)),
                        o.sendNotification("error_occurred", t))
                      : null != e &&
                          502 == e.code &&
                          null != e.detail &&
                          "talkers" == e.detail.limit_target
                        ? ((t = e.detail.limit_max),
                          (t = _a.toLocalError(R.TalkersCountOvered(t), e)),
                          o.sendNotification("error_occurred", t))
                        : o.apiCaller.handleServerErrorDefault(e));
                },
              );
        },
        deleteTalker: function (t, n) {
          var i,
            r = this,
            o = null,
            a = null,
            e = this.dataStore;
          ((i =
            null != e.me &&
            null != (e = e.me.id) &&
            null != n &&
            e.high == n.high &&
            e.low == n.low) &&
            ((o = t),
            (a = this.dataStore.getTalkStatus(t.id)),
            this.dataStore.removeTalk(t.id),
            this.dataStore.removeTalkStatus(t.id)),
            this.wsApi.deleteTalker(t.id, n).then(
              function (e) {
                (D.remove(t.userIds, n),
                  i || r.dataStore.setTalk(t),
                  i && r.dataStore.removeConferencesOnLeaveTalk(t.id),
                  r.sendNotification("notify_delete_talker", t));
              },
              function (e) {
                (null != o && r.dataStore.setTalk(o),
                  null != a && r.dataStore.setTalkStatus(a),
                  r.sendNotification("error_occurred", e));
              },
            ));
        },
        getMessages: function (t, e, n, i) {
          var r = this;
          (null == n && (n = gr.desc),
            this.wsApi
              .getMessages(
                t,
                20,
                (e = null == e ? { sinceId: null, maxId: null } : e),
                n,
              )
              .then(
                function (e) {
                  (k.iter(e, z((j = r.dataStore), j.setMessage)),
                    r.sendNotification("notify_get_messages", {
                      talkId: t,
                      messages: e,
                      sortOrder: n,
                      callerKey: i,
                    }));
                },
                function (e) {
                  r.apiCaller.handleServerErrorDefault(e);
                },
              ));
        },
        createMessage: function (e, t, n, i) {
          var r,
            o = this.dataFactory.newDummyMessage(e, t, n);
          null == this.dataStore.getTalk(e)
            ? (F.w(
                "WARNING method: create_message",
                (r = { code: 400, message: "invalid talk_id (deleted talk)" }),
              ),
              this.sendNotification("create_message_fail", [r, o], i))
            : (this.sendNotification("create_message_start", o),
              this._createMessage(e, t, n, o, i).catch(function (e) {}));
        },
        _createMessage: function (e, t, n, i, r) {
          var o = this;
          return this.wsApi.createMessage(e, t, n).then(
            function (e) {
              return (
                o.dataStore.removeMessage(i.id),
                o.dataStore.setMessage(e),
                o.dataStore.setFromThisDevice(e),
                o.sendNotification("create_message_complete", [e, i.id], r),
                e
              );
            },
            function (e) {
              var t;
              return (
                o.sendNotification("create_message_fail", [e, i], r),
                null != e && 409 == e.code
                  ? ((t = _a.toLocalError(R.SendingMessageConflicted, e)),
                    o.sendNotification("error_occurred", t))
                  : o.sendNotification("error_occurred", e),
                Promise.reject(e)
              );
            },
          );
        },
        forwardMessages: function (t, n, i) {
          var e,
            r,
            o,
            a = this;
          (i = null == i ? 0 : i) >= n.length ||
            ((e = n[i]),
            (r = t.id),
            (o = this.dataFactory.newDummyMessage(r, e.type, e.content)),
            this.sendNotification("create_message_start", o),
            this.wsApi.createMessage(r, e.type, e.content).then(
              function (e) {
                (a.dataStore.setMessage(e),
                  a.sendNotification("create_message_complete", [e, o.id]),
                  a.forwardMessages(t, n, i + 1),
                  a.dataStore.removeMessage(o.id));
              },
              function (e) {
                var t;
                (a.sendNotification("create_message_fail", [e, o]),
                  null != e && 409 == e.code
                    ? ((t = _a.toLocalError(R.SendingMessageConflicted, e)),
                      a.sendNotification("error_occurred", t))
                    : a.sendNotification("error_occurred", e));
              },
            ));
        },
        deleteMessage: function (e, t) {
          var n = this;
          this.wsApi.deleteMessage(e, t).then(
            function (e) {},
            function (e) {
              var t;
              null != e &&
                ((403 != e.code || "frozen domain" != e.message) &&
                409 == e.code &&
                "conflict" == e.message
                  ? ((t = _a.toLocalError(R.DeletingMessageConflicted, e)),
                    n.sendNotification("error_occurred", t))
                  : n.sendNotification("error_occurred", e));
            },
          );
        },
        updateReadStatuses: function (n, i) {
          var r = this;
          this.wsApi.updateReadStatuses(n, i).then(
            function (e) {
              var t = r.dataStore.getTalkStatus(n);
              null != t &&
                (t.updateByReadingMessagesFinished(i),
                r.sendNotification("notify_update_local_talk_status", t));
            },
            function (e) {
              var t = r.dataStore.getTalkStatus(n);
              (null != t &&
                (t.updateByReadingMessagesCanceled(i),
                r.dataStore.setTalkStatus(t),
                r.sendNotification("notify_update_local_talk_status", t)),
                r.apiCaller.handleServerErrorDefault(e));
            },
          );
        },
        searchMessagesAroundDateTime: function (e, t) {
          return this.wsApi.searchMessagesAroundDateTime(e, t);
        },
        uploadV2: function (e, t, n, i) {
          var r = this,
            o = Mt.toFileInfoDynamic(n),
            a = this.dataFactory.newDummyMessage(t, d.file, o);
          (this.sendNotification("create_message_start", a),
            this.prepareMessageFileAsync(n, e, !1).then(
              function (e) {
                r._createMessage(t, d.file, e, a, i).catch(function (e) {});
              },
              function (e) {
                r.sendNotification("create_message_fail", [e, a], i);
              },
            ));
        },
        uploadForHubot: function (e, t, n) {
          var i,
            r = this.dataStore.getTalk(e),
            t = { file: t, thumb: null };
          null == r
            ? ((i = Mt.toFileInfoDynamic(t)),
              (e = this.dataFactory.newDummyMessage(e, d.file, i)),
              F.w(
                "WARNING method: upload",
                (i = { code: 400, message: "invalid talk_id (deleted talk)" }),
              ),
              this.sendNotification("create_message_fail", [i, e], n))
            : this.uploadV2(r.domainId, r.id, t, n);
        },
        uploadMultiV2: function (e, t, n, i, r) {
          var o = this,
            a = Mt.toFileInfoDynamicArray(i),
            s = this.dataFactory.newDummyMessage(t, d.textMultipleFile, {
              text: n,
              files: a,
            });
          (this.sendNotification("create_message_start", s),
            this.prepareMessageFilesAsync(e, n, i, !1).then(
              function (e) {
                o._createMessage(t, d.textMultipleFile, e, s, r).catch(
                  function (e) {},
                );
              },
              function (e) {
                o.sendNotification("create_message_fail", [e, s], r);
              },
            ));
        },
        uploadMultiForHubot: function (e, t, n, i) {
          var r,
            o = this.dataStore.getTalk(e),
            n = n.map(function (e) {
              return { file: e, thumb: null };
            });
          null == o
            ? ((r = Mt.toFileInfoDynamicArray(n)),
              (e = this.dataFactory.newDummyMessage(e, d.textMultipleFile, {
                text: t,
                files: r,
              })),
              F.w(
                "WARNING method: uploadMulti",
                (r = { code: 400, message: "invalid talk_id (deleted talk)" }),
              ),
              this.sendNotification("create_message_fail", [r, e], i))
            : this.uploadMultiV2(o.domainId, o.id, t, n, i);
        },
        validateFileName: function (t) {
          var n = I.replace(t.name, "$", "__");
          if (n == t.name) return t;
          try {
            return ((t.name = n), t);
          } catch (e) {
            var i = t.slice(0, t.size, t.type);
            return new File([i], n, { type: t.type });
          }
        },
        prepareMessageFilesAsync: function (n, t, i, r) {
          for (var o = this, e = Promise.resolve([]), a = 0; a < i.length; ) {
            var s = [i[a]];
            (++a,
              (e = e.then(
                (function (e) {
                  return function (t) {
                    return o
                      .prepareMessageFileAsync(e[0], n, r)
                      .then(function (e) {
                        return (t.push(e), t);
                      });
                  };
                })(s),
              )));
          }
          return e.then(function (e) {
            return e.length != i.length
              ? Promise.reject({
                  code: 400,
                  message: "uploaded files count is invalid.",
                })
              : ((e = { files: e }),
                null != t && 0 < t.length && (e.text = t),
                Promise.resolve(e));
          });
        },
        prepareMessageFileAsync: function (e, r, o) {
          var n,
            a,
            s = this;
          return null != e.uploadedFile
            ? Promise.resolve(e.uploadedFile)
            : ((n = e.thumb),
              (a = this.validateFileName((a = e.file))),
              new Promise(function (t, e) {
                null == n
                  ? t(m.None)
                  : s.uploadFile(
                      n.file,
                      r,
                      u.THUMBNAIL,
                      function (e) {
                        (n.setAuth(e), t(m.Some(n)));
                      },
                      function (e) {
                        t(m.None);
                      },
                    );
              }).then(function (i) {
                return new Promise(function (n, e) {
                  var t = o ? u.SCHEDULE_MESSAGE : u.MESSAGE;
                  s.uploadFile(
                    a,
                    r,
                    t,
                    function (e) {
                      var t = xo.normalizeForFile(a.name),
                        t = K.createFileInfoDynamic(a, t, i, e);
                      n(t);
                    },
                    e,
                  );
                });
              }));
        },
        prepareNoteFileAsync: function (r, n, o) {
          var a = this;
          return (
            (r = this.validateFileName(r)),
            new Promise(function (t, e) {
              null == n
                ? t(m.None)
                : a.uploadFile(
                    n.file,
                    o,
                    u.NOTE_THUMBNAIL,
                    function (e) {
                      (n.setAuth(e), t(m.Some(n)));
                    },
                    function (e) {
                      t(m.None);
                    },
                  );
            }).then(function (i) {
              return new Promise(function (n, e) {
                a.uploadFile(
                  r,
                  o,
                  u.NOTE_ATTACHMENT,
                  function (e) {
                    var t = xo.normalizeForFile(r.name),
                      t = K.createFileInfoDynamic(r, t, i, e);
                    n(t);
                  },
                  e,
                );
              });
            })
          );
        },
        uploadFile: function (e, t, i, n, r) {
          var o = this,
            a = e,
            s = xo.normalizeForFile(a.name),
            l = a.type,
            a = a.size;
          this.wsApi.createUploadAuth(s, l, a, t, i).then(
            function (t) {
              o.fileService.upload(t, l, e).then(
                function (e) {
                  n(t.rawValue);
                },
                function (e) {
                  F.e("failed to upload: " + P.string(e));
                  var t,
                    e = new ca({ code: null, detail: null, message: e });
                  (i != u.THUMBNAIL &&
                    i != u.NOTE_THUMBNAIL &&
                    i != u.NOTE_ATTACHMENT &&
                    ((t =
                      i == u.SCHEDULE_MESSAGE
                        ? R.ScheduleMessageFailed(dr.GeneralFileError)
                        : R.GeneralFileError),
                    (t = _a.toLocalError(t, e)),
                    o.sendNotification("error_occurred", t)),
                    r(e));
                },
              );
            },
            function (e) {
              var t, n;
              (i != u.THUMBNAIL &&
                i != u.NOTE_THUMBNAIL &&
                i != u.NOTE_ATTACHMENT &&
                null != e &&
                (502 == e.code && null != e.detail
                  ? ((n = D.fromIntOrInt64(e.detail.limit_max)),
                    "file_size" == e.detail.limit_target
                      ? ((t =
                          i == u.SCHEDULE_MESSAGE
                            ? R.ScheduleMessageFailed(dr.UploadSizeOvered(n))
                            : R.UploadSizeOvered(n)),
                        (t = _a.toLocalError(t, e)),
                        o.sendNotification("error_occurred", t))
                      : "storage_size" == e.detail.limit_target &&
                        ((t =
                          i == u.SCHEDULE_MESSAGE
                            ? R.ScheduleMessageFailed(dr.StorageSizeOvered(n))
                            : R.StorageSizeOvered(n)),
                        (n = _a.toLocalError(t, e)),
                        o.sendNotification("error_occurred", n)))
                  : o.sendNotification("error_occurred", e)),
                r(e));
            },
          );
        },
        createDownloadAuth: function (e, t, n, i, r) {
          this.wsApi.createDownloadAuth(e, t, n).then(
            function (e) {
              i(e);
            },
            function (e) {
              r(e);
            },
          );
        },
        deleteAttachment: function (e, t) {
          var n = this;
          this.wsApi.deleteAttachment(e, t).then(
            function (e) {},
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getAttachments: function (t, e) {
          var n = this;
          this.wsApi
            .getAttachments(
              t,
              20,
              (e = null == e ? { sinceId: null, maxId: null } : e).sinceId,
              e.maxId,
            )
            .then(
              function (e) {
                (k.iter(e, function (e) {
                  n.dataStore.setTalkFileInfo(e, !0);
                }),
                  n.sendNotification("get_file_responsed", {
                    talkId: t,
                    files: e,
                  }));
              },
              function (e) {
                n.apiCaller.handleServerErrorDefault(e);
              },
            );
        },
        createAnnouncement: function (e, t, n, i) {
          var r,
            o = this;
          (null == e && null == (e = this.settings.getSelectedDomainId())) ||
            ((r = this.dataFactory.newDummyMessage(null, t, n)),
            this.sendNotification("create_announcement_start", r),
            this.wsApi.createAnnouncement(e, t, n).then(
              function (e) {
                o.sendNotification(
                  "create_announcement_complete",
                  [e, r.id],
                  i,
                );
              },
              function (e) {
                (o.sendNotification("create_announcement_fail", [e, r], i),
                  o.sendNotification("error_occurred", e));
              },
            ));
        },
        getAnnouncements: function (e, t, n) {
          var i = this,
            r = this.settings.getSelectedDomainId();
          null != r &&
            this.wsApi
              .getAnnouncements(
                r,
                20,
                (e = null == e ? { sinceId: null, maxId: null } : e),
                t,
              )
              .then(
                function (e) {
                  (k.iter(e, z((j = i.dataStore), j.setAnnouncement)),
                    i.sendNotification("notify_get_announcements", {
                      domainId: r,
                      announcements: e,
                      sortOrder: t,
                      callerKey: n,
                    }));
                },
                function (e) {
                  i.apiCaller.handleServerErrorDefault(e);
                },
              );
        },
        getAnnouncementStatuses: function (t) {
          var n = this;
          this.wsApi.getAnnouncementStatuses().then(
            function (e) {
              (k.iter(e, function (e) {
                (n.dataStore.setAnnouncementStatus(e),
                  n.sendAnnouncementStatusNotification(e));
              }),
                null != t && t());
            },
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getAnnouncementStatus: function (e, t) {
          var n = this;
          this.wsApi.getAnnouncementStatus(e).then(
            function (e) {
              (n.dataStore.setAnnouncementStatus(e),
                n.sendAnnouncementStatusNotification(e),
                null != t && t());
            },
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        updateAnnouncementReadStatus: function (n, i) {
          var r = this;
          this.wsApi.updateAnnouncementReadStatus(n, i).then(
            function (e) {
              var t = r.dataStore.getAnnouncementStatus(n);
              null != t &&
                (t.updateByReadingAnnouncementsFinished(i),
                r.sendAnnouncementStatusNotification(t));
            },
            function (e) {
              var t = r.dataStore.getAnnouncementStatus(n);
              (null != t &&
                (t.updateByReadingAnnouncementsCanceled(i),
                r.dataStore.setAnnouncementStatus(t),
                r.sendAnnouncementStatusNotification(t)),
                r.apiCaller.handleServerErrorDefault(e));
            },
          );
        },
        getQuestions: function (e, t, n, i) {
          var r = this,
            o = this.settings.getSelectedDomainId(),
            a = null,
            e =
              (null != e && ((o = e.domainId), (a = e.id)),
              h.orNull(Cr.filterToOpBool(n)));
          this.wsApi
            .getQuestions(
              o,
              a,
              t,
              e,
              20,
              (i = null == i ? { sinceId: null, maxId: null } : i),
            )
            .then(
              function (e) {
                e = e.map(z((j = r.dataStore), j.setQuestionIfLatest));
                r.sendNotification("get_questions_responsed", {
                  domainId: o,
                  talkId: a,
                  questions: e,
                  fromType: t,
                  filter: n,
                });
              },
              function (e) {
                r.apiCaller.handleServerErrorDefault(e);
              },
            );
        },
        getQuestion: function (e, t) {
          var n = this;
          this.wsApi.getQuestion(e).then(
            function (e) {
              e = n.dataStore.setQuestionIfLatest(e);
              (n.sendNotification("notify_update_question", e),
                null != t && t(e));
            },
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        searchMessages: function (n, i, r, o) {
          var a = this;
          (this.searchService.prepareForSearching(n, r),
            this.wsApi
              .searchMessages(n.domainId, n.talkId, n.getQueryText(), i, r)
              .then(
                function (e) {
                  (a.searchService.isRecentParams(n) &&
                    a.searchService.isSearching() &&
                    a.searchService.saveSearchMessagesResult(e),
                    null != o && o());
                },
                function (e) {
                  var t;
                  (a.searchService.notifySearchMessagesFail(),
                    400 == e.code && "search too many talks" == e.message
                      ? ((t = _a.toLocalError(R.SearchTooManyTalks, e)),
                        a.sendNotification("error_occurred", t))
                      : 400 == e.code && null != r
                        ? a.searchMessages(n, i, null, o)
                        : a.sendNotification("error_occurred", e));
                },
              ));
        },
        searchAttachments: function (n, i, r, o) {
          var a = this;
          (this.searchService.prepareForSearching(n, r),
            this.wsApi
              .searchAttachments(n.domainId, n.talkId, n.getQueryText(), i, r)
              .then(
                function (e) {
                  (a.searchService.isRecentParams(n) &&
                    a.searchService.isSearching() &&
                    a.searchService.saveSearchAttachmentsResult(e),
                    null != o && o());
                },
                function (e) {
                  var t;
                  (a.searchService.notifySearchAttachmentsFail(),
                    400 == e.code && "search too many talks" == e.message
                      ? ((t = _a.toLocalError(R.SearchTooManyTalks, e)),
                        a.sendNotification("error_occurred", t))
                      : 400 == e.code && null != r
                        ? a.searchAttachments(n, i, null, o)
                        : a.sendNotification("error_occurred", e));
                },
              ));
        },
        startUpdateLastUsedAtIfNeed: function () {
          this.lastUsedAtUpdater.startPolling(z(this, this.updateLastUsedAt));
        },
        updateLastUsedAt: function () {
          var t = this;
          this.wsApi.updateLastUsedAt().then(
            function (e) {
              t.lastUsedAtUpdater.update();
            },
            function (e) {
              (t.lastUsedAtUpdater.cancel(),
                t.apiCaller.handleServerErrorDefault(e));
            },
          );
        },
        addFavoriteTalk: function (e, t) {
          var n = this;
          this.wsApi.addFavoriteTalk(e, t).catch(function (e) {
            n.apiCaller.handleServerErrorDefault(e);
          });
        },
        deleteFavoriteTalk: function (e, t) {
          var n = this;
          this.wsApi.deleteFavoriteTalk(e, t).catch(function (e) {
            n.apiCaller.handleServerErrorDefault(e);
          });
        },
        disablePushNotification: function (e, t) {
          var n = this;
          this.wsApi.disablePushNotification(e, t).catch(function (e) {
            n.apiCaller.handleServerErrorDefault(e);
          });
        },
        enablePushNotification: function (e, t, n) {
          var i = this;
          this.wsApi.enablePushNotification(e, t, n).catch(function (e) {
            i.apiCaller.handleServerErrorDefault(e);
          });
        },
        getSolutions: function (e, t) {
          var n = this;
          (k.iter(t, z((j = this.dataStore), j.addLoadingSolutionId)),
            this.wsApi.getSolutions(e, t).then(
              function (e) {
                (k.iter(e, z((j = n.dataStore), j.setSolution)),
                  k.iter(t, z((j = n.dataStore), j.removeLoadingSolutionId)),
                  n.sendNotification("solutions_loaded"));
              },
              function (e) {
                (k.iter(t, z((j = n.dataStore), j.removeLoadingSolutionId)),
                  n.apiCaller.handleServerErrorDefault(e));
              },
            ));
        },
        getPresences: function (e, t) {
          var n = this;
          return this.wsApi.getPresences(e, t).then(
            function (e) {
              return (
                k.iter(e, z((j = n.dataStore), j.setUserPresence)),
                n.sendNotification("presences_updated", t),
                Promise.resolve(t)
              );
            },
            function (e) {
              return (
                n.apiCaller.handleServerErrorDefault(e),
                Promise.reject(e)
              );
            },
          );
        },
        joinConference: function (t, e) {
          var n = this;
          this.wsApi.joinConference(t, e).then(
            function (e) {
              n.sendNotification("join_conference_responsed", {
                talkId: t,
                joinConference: e,
              });
            },
            function (e) {
              n.sendNotification("join_conference_canceled", e);
            },
          );
        },
        rejectConference: function (e) {
          var t = this;
          this.wsApi.rejectConference(e).catch(function (e) {
            404 == e.code
              ? F.v("conference not found.")
              : 409 == e.code
                ? F.v("joined.")
                : t.sendNotification("error_occurred", e);
          });
        },
        leaveConference: function (e) {
          var t = this;
          this.wsApi.leaveConference(e).catch(function (e) {
            404 == e.code
              ? F.v("conference not found.")
              : t.sendNotification("error_occurred", e);
          });
        },
        getConferences: function () {
          var t = this;
          this.wsApi.getConferences().then(
            function (e) {
              (k.iter(e, function (e) {
                (t.dataStore.addConference(e),
                  t.sendNotification("notify_open_conference", e));
              }),
                t.dataStore.filterReactedConferenceIds(
                  e
                    .filter(function (e) {
                      return !t.dataStore.isReactedConferenceId(e.id);
                    })
                    .map(function (e) {
                      return e.id;
                    }),
                ));
            },
            function (e) {
              t.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getConferenceParticipants: function (n) {
          var i = this;
          this.wsApi.getConferenceParticipants(n).then(
            function (e) {
              var t = i.dataStore.getConference(n);
              null != t &&
                ((t.participants = e),
                i.sendNotification("get_conference_participants_responsed", t));
            },
            function (e) {
              i.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getNoteStatuses: function (e, t) {
          var n = this.settings.getSelectedDomainId();
          this.apiNote.getNoteStatuses(n, e, t).catch(function (e) {
            F.v("getNoteStatuses error:", e);
          });
        },
        getNote: function (e, t) {
          var n = null != t ? t : function (e, t) {};
          this.apiNote
            .getNote(e)
            .then(function (e) {
              n(null, e);
            })
            .catch(function (e) {
              (F.v("getNote error:", e), n(e, null));
            });
        },
        createNote: function (e, t, n, i, r) {
          (F.w("'createNote' is deprecated, please use 'res.send' instead."),
            this.apiNote.createNote(e, t, n, i, r).catch(function (e) {
              F.v("createNote error:", e);
            }));
        },
        createNoteV2: function (o, a, s, e) {
          var l = this,
            t = this.dataStore.getTalk(a.getTalkId());
          if (null != t) {
            for (
              var r = t.domainId,
                u = a.getNoteId(),
                n = (this.dataStore.setNoteLocalEdit(u, a), Promise.resolve()),
                i = 0,
                c = a.getUnuploadedFileInfos();
              i < c.length;
            ) {
              var _ = [c[i]];
              (++i,
                (n = n.then(
                  (function (i) {
                    return function (e) {
                      return l
                        .prepareNoteFileAsync(i[0].file, i[0].localThumbInfo, r)
                        .then(
                          ((t = i),
                          function (e) {
                            t[0].setUploadResult(sr.SUCCESS(e));
                          }),
                        )
                        .catch(
                          ((n = i),
                          function (e) {
                            n[0].setUploadResult(sr.FAILUER);
                            var t = De.convertFileErrorForNote(u, e);
                            return (
                              null != t &&
                                (a.setState(f.FAILED_BY_FILE(n[0].id, t)),
                                l.dataStore.setNoteLocalEdit(u, a),
                                (t = _a.toLocalError(t, e)),
                                l.sendNotification("error_occurred", t)),
                              Promise.reject(e)
                            );
                          }),
                        );
                      var n, t;
                    };
                  })(_),
                )));
            }
            var h = null != e ? e : function (e, t) {};
            n.then(function (e) {
              var t = l.apiNote,
                n = a.getTitle(),
                i = a.getContentType(),
                r = a.getContent();
              return t.createNote(o, n, i, r, s, a);
            })
              .then(function (e) {
                h(null, e);
              })
              .catch(function (e) {
                (F.v("createNoteV2 error:", e), h(e, null));
              });
          }
        },
        updateNoteSetting: function (e, t, n) {
          this.apiNote.updateNoteSetting(e, t, n).catch(function (e) {
            F.v("updateNoteSetting error:", e);
          });
        },
        updateNote: function (e, t, n, i, r, o) {
          this.apiNote.updateNote(e, t, n, i, r, o).catch(function (e) {
            F.v("updateNote error:", e);
          });
        },
        updateNoteV2: function (o, a, s, l, e) {
          var u = this,
            t = this.dataStore.getTalk(s.getTalkId());
          if (null == t) return Promise.resolve();
          for (
            var r = t.domainId,
              n = (this.dataStore.setNoteLocalEdit(o, s), Promise.resolve()),
              i = 0,
              c = s.getUnuploadedFileInfos();
            i < c.length;
          ) {
            var _ = [c[i]];
            (++i,
              (n = n.then(
                (function (i) {
                  return function (e) {
                    return u
                      .prepareNoteFileAsync(i[0].file, i[0].localThumbInfo, r)
                      .then(
                        ((t = i),
                        function (e) {
                          t[0].setUploadResult(sr.SUCCESS(e));
                        }),
                      )
                      .catch(
                        ((n = i),
                        function (e) {
                          n[0].setUploadResult(sr.FAILUER);
                          var t = De.convertFileErrorForNote(o, e);
                          return (
                            null != t &&
                              (s.setState(f.FAILED_BY_FILE(n[0].id, t)),
                              u.dataStore.setNoteLocalEdit(o, s),
                              (t = _a.toLocalError(t, e)),
                              u.sendNotification("error_occurred", t)),
                            Promise.reject(e)
                          );
                        }),
                      );
                    var n, t;
                  };
                })(_),
              )));
          }
          var h = null != e ? e : function (e, t) {};
          return n
            .then(function (e) {
              var t = u.apiNote,
                n = s.getTitle(),
                i = s.getContentType(),
                r = s.getContent();
              return t.updateNote(o, a, n, i, r, l, s);
            })
            .then(function (e) {
              h(null, e);
            })
            .catch(function (e) {
              return (
                F.v("updateNote error:", e),
                h(e, null),
                409 == e.code && "locked by another user" == e.message
                  ? Promise.reject(e)
                  : u.unlockNote(o, a).then(function (e) {})
              );
            })
            .catch(function (e) {});
        },
        deleteNote: function (e, t) {
          var n = null != t ? t : function (e, t) {};
          this.apiNote
            .deleteNote(e)
            .then(function (e) {
              n(null, e);
            })
            .catch(function (e) {
              (F.v("deleteNote error:", e), n(e, null));
            });
        },
        lockNote: function (e, t) {
          var n = this.settings.getSelectedDomainId();
          return this.apiNote.lockNote(n, e, t);
        },
        unlockNote: function (e, t) {
          var n = this.settings.getSelectedDomainId();
          return this.apiNote.unlockNote(n, e, t);
        },
        addFavoriteMessage: function (e, t, n) {
          var i = this;
          this.wsApi.addFavoriteMessage(e, t, n).then(
            function (e) {
              i.sendNotification("add_favorite_message_completed", e);
            },
            function (e) {
              (400 == e.code &&
                "invalid message" == e.message &&
                F.v("invalid message."),
                400 == e.code &&
                  "invalid message (not favoritable message)" == e.message &&
                  F.v("invalid message (not favoritable message)."),
                403 == e.code && "forbidden" == e.message && F.v("forbidden."),
                404 == e.code && F.v("message is not found."),
                409 == e.code && "conflict" == e.message
                  ? F.v("conflict")
                  : i.sendNotification("error_occurred", e));
            },
          );
        },
        deleteFavoriteMessage: function (e, t) {
          var n = this;
          this.wsApi.deleteFavoriteMessage(e, t).then(
            function (e) {
              n.sendNotification("delete_favorite_message_completed", e);
            },
            function (e) {
              (400 == e.code && F.v("invalid parameter."),
                404 == e.code && F.v("message is not found."),
                n.sendNotification("error_occurred", e));
            },
          );
        },
        getFavoriteMessages: function (e, t, n) {
          var i = this;
          this.wsApi.getFavoriteMessages(e, t, 20, n).then(
            function (e) {
              (k.iter(e.messages, z((j = i.dataStore), j.setMessage)),
                i.sendNotification("get_favorite_messages_completed", e));
            },
            function (e) {
              (400 == e.code && F.v("invalid parameter."),
                i.sendNotification("error_occurred", e));
            },
          );
        },
        scheduleFileMessage: function (t, e, n, i) {
          var r = this,
            o = Mt.toFileInfoDynamic(i),
            a = this.dataFactory.newDummyMessage(n, d.file, o, t);
          (this.sendNotification("schedule_message_started", a),
            this.prepareMessageFileAsync(i, e, !0)
              .then(function (e) {
                return r._scheduleMessage(t, n, d.file, e);
              })
              .then(
                function (e) {
                  r.sendNotification("schedule_message_completed", [e, a.id]);
                },
                function (e) {
                  r.sendNotification("schedule_message_faileed", [e, a]);
                },
              ));
        },
        scheduleFilesMessage: function (t, e, n, i, r) {
          var o = this,
            a = Mt.toFileInfoDynamicArray(r),
            s = this.dataFactory.newDummyMessage(
              n,
              d.textMultipleFile,
              { text: i, files: a },
              t,
            );
          (this.sendNotification("schedule_message_started", s),
            this.prepareMessageFilesAsync(e, i, r, !0)
              .then(function (e) {
                return o._scheduleMessage(t, n, d.textMultipleFile, e);
              })
              .then(
                function (e) {
                  o.sendNotification("schedule_message_completed", [e, s.id]);
                },
                function (e) {
                  o.sendNotification("schedule_message_faileed", [e, s]);
                },
              ));
        },
        scheduleMessage: function (e, t, n, i) {
          var r = this,
            o = this.dataFactory.newDummyMessage(t, n, i, e);
          (this.sendNotification("schedule_message_started", o),
            this._scheduleMessage(e, t, n, i).then(
              function (e) {
                r.sendNotification("schedule_message_completed", [e, o.id]);
              },
              function (e) {
                r.sendNotification("schedule_message_faileed", [e, o]);
              },
            ));
        },
        _scheduleMessage: function (e, t, n, i) {
          var r = this;
          return this.wsApi.scheduleMessage(e, t, n, i).then(
            function (e) {
              return (r.dataStore.setScheduledMessage(e), e);
            },
            function (e) {
              var t;
              return (
                400 == e.code
                  ? ((t = R.ScheduleMessageFailed(dr.InvalidCreateAt)),
                    (t = _a.toLocalError(t, e)),
                    r.sendNotification("error_occurred", t))
                  : 403 == e.code
                    ? ((t = R.ScheduleMessageFailed(dr.Unknown)),
                      (t = _a.toLocalError(t, e)),
                      r.sendNotification("error_occurred", t))
                    : 409 == e.code
                      ? ((t = R.ScheduleMessageFailed(dr.Conflict)),
                        (t = _a.toLocalError(t, e)),
                        r.sendNotification("error_occurred", t))
                      : 502 == e.code
                        ? ((t = R.ScheduleMessageFailed(dr.LimitReached)),
                          (t = _a.toLocalError(t, e)),
                          r.sendNotification("error_occurred", t))
                        : r.apiCaller.handleServerErrorDefault(e),
                Promise.reject(e)
              );
            },
          );
        },
        deleteScheduledMessage: function (e) {
          var t = this;
          this._deleteScheduledMessage(e).then(
            function (e) {
              t.sendNotification("delete_scheduled_messages_completed", e);
            },
            function (e) {},
          );
        },
        _deleteScheduledMessage: function (t) {
          var n = this;
          return this.wsApi.deleteScheduledMessage(t).then(
            function (e) {
              return n.dataStore.removeScheduledMessage(t);
            },
            function (e) {
              return (
                n.apiCaller.handleServerErrorDefault(e),
                Promise.reject(e)
              );
            },
          );
        },
        rescheduleFileMessage: function (n, t, e, i, r) {
          var o = this,
            a = Mt.toFileInfoDynamic(r),
            s = this.dataFactory.newDummyScheduledMessageEdit(i, d.file, a, t),
            l =
              (this.dataStore.setScheduledMessageEdit(n, s),
              this.sendNotification("reschedule_message_started", [n, s]),
              this.dataStore.getScheduledMessage(n));
          this.prepareMessageFileAsync(r, e, !0)
            .then(function (e) {
              return o._rescheduleMessage(n, t, i, d.file, e);
            })
            .then(
              function (e) {
                o.sendNotification("reschedule_message_completed", [l, e]);
              },
              function (e) {
                var t = s.toEditWithError();
                (o.dataStore.setScheduledMessageEdit(n, t),
                  o.sendNotification("reschedule_message_failed", [n, t, e]));
              },
            );
        },
        rescheduleFilesMessage: function (n, t, e, i, r, o) {
          var a = this,
            s = Mt.toFileInfoDynamicArray(o),
            l = this.dataFactory.newDummyScheduledMessageEdit(
              i,
              d.textMultipleFile,
              { text: r, files: s },
              t,
            ),
            u =
              (this.dataStore.setScheduledMessageEdit(n, l),
              this.sendNotification("reschedule_message_started", [n, l]),
              this.dataStore.getScheduledMessage(n));
          this.prepareMessageFilesAsync(e, r, o, !0)
            .then(function (e) {
              return a._rescheduleMessage(n, t, i, d.textMultipleFile, e);
            })
            .then(
              function (e) {
                a.sendNotification("reschedule_message_completed", [u, l, e]);
              },
              function (e) {
                var t = l.toEditWithError();
                (a.dataStore.setScheduledMessageEdit(n, t),
                  a.sendNotification("reschedule_message_failed", [n, t, e]));
              },
            );
        },
        rescheduleMessage: function (n, e, t, i, r) {
          var o = this,
            a = this.dataFactory.newDummyScheduledMessageEdit(t, i, r, e),
            s =
              (this.dataStore.setScheduledMessageEdit(n, a),
              this.sendNotification("reschedule_message_started", [n, a]),
              this.dataStore.getScheduledMessage(n));
          this._rescheduleMessage(n, e, t, i, r).then(
            function (e) {
              o.sendNotification("reschedule_message_completed", [s, e]);
            },
            function (e) {
              var t = a.toEditWithError();
              (o.dataStore.setScheduledMessageEdit(n, t),
                o.sendNotification("reschedule_message_failed", [n, t, e]));
            },
          );
        },
        _rescheduleMessage: function (t, e, n, i, r) {
          var o = this;
          return this.wsApi.rescheduleMessage(t, e, n, i, r).then(
            function (e) {
              return (
                o.dataStore.removeScheduledMessage(t),
                o.dataStore.setScheduledMessage(e),
                e
              );
            },
            function (e) {
              var t;
              return (
                400 == e.code
                  ? ((t = R.ScheduleMessageFailed(dr.InvalidCreateAt)),
                    (t = _a.toLocalError(t, e)),
                    o.sendNotification("error_occurred", t))
                  : 403 == e.code
                    ? ((t = R.ScheduleMessageFailed(dr.Unknown)),
                      (t = _a.toLocalError(t, e)),
                      o.sendNotification("error_occurred", t))
                    : 409 == e.code
                      ? ((t = R.ScheduleMessageFailed(dr.Conflict)),
                        (t = _a.toLocalError(t, e)),
                        o.sendNotification("error_occurred", t))
                      : 502 == e.code
                        ? ((t = R.ScheduleMessageFailed(dr.LimitReached)),
                          (t = _a.toLocalError(t, e)),
                          o.sendNotification("error_occurred", t))
                        : o.apiCaller.handleServerErrorDefault(e),
                Promise.reject(e)
              );
            },
          );
        },
        sendScheduledMessageNow: function (t, e) {
          var n,
            i = this,
            r = t.talkId,
            o = t.type,
            a = t.content,
            s = this.dataFactory.newDummyMessage(r, o, a);
          null == this.dataStore.getTalk(r)
            ? (F.w(
                "WARNING method: create_message",
                (n = { code: 400, message: "invalid talk_id (deleted talk)" }),
              ),
              this.sendNotification("create_message_fail", [n, s], e))
            : (this.sendNotification("create_message_start", s),
              this._createMessage(r, o, a, s, e).then(
                function (e) {
                  i.deleteScheduledMessage(t.id);
                },
                function (e) {},
              ));
        },
        getScheduledMessages: function (t, e, n) {
          var i = this;
          this.wsApi.getScheduledMessages(t, e, 10, null).then(
            function (e) {
              (n && i.dataStore.removeScheduledMessagesInDomain(t),
                k.iter(e.contents, z((j = i.dataStore), j.setScheduledMessage)),
                i.sendNotification("get_scheduled_messages_completed", e));
            },
            function (e) {
              (i.sendNotification("get_scheduled_messages_failed"),
                i.apiCaller.handleServerErrorDefault(e));
            },
          );
        },
        getStampsets: function (e) {
          var t = this;
          return this.wsApi.getStampsets(e).then(function (e) {
            return (t.dataStore.addStampsets(e), e);
          });
        },
        setMessageReaction: function (t, n, i) {
          var r = this,
            o = this.dataStore.getMessageReactionSummary(n);
          this.wsApi.setMessageReaction(t, n, i).then(
            function (e) {
              (null != o &&
                (o.updateMyReaction(i),
                r.dataStore.removeMessageReactionUsers(n)),
                r.sendNotification("set_message_reaction_completed", [t, n]));
            },
            function (e) {
              r.sendNotification("set_message_reaction_failed");
              var t = o.hasMyReaction()
                  ? R.FailedToChangeReaction
                  : R.FailedToAddReaction,
                t = _a.toLocalError(t, e);
              r.sendNotification("error_occurred", t);
            },
          );
        },
        resetMessageReaction: function (n, i) {
          var r = this;
          this.wsApi.resetMessageReaction(n, i).then(
            function (e) {
              var t = r.dataStore.getMessageReactionSummary(i);
              (null != t &&
                (t.updateMyReaction(null),
                r.dataStore.removeMessageReactionUsers(i)),
                r.sendNotification("reset_message_reaction_completed", [n, i]));
            },
            function (e) {
              r.sendNotification("reset_message_reaction_failed");
              e = _a.toLocalError(R.FailedToCancelReaction, e);
              r.sendNotification("error_occurred", e);
            },
          );
        },
        getMessageReactionSummaries: function (t, n) {
          var i = this;
          (k.iter(
            n,
            z((j = this.dataStore), j.addLoadingMessageReactionSmmary),
          ),
            this.wsApi.getMessageReactionSummaries(t, n).then(
              function (e) {
                (k.iter(
                  n,
                  z((j = i.dataStore), j.removeLoadingMessageReactionSmmary),
                ),
                  k.iter(e, function (e) {
                    i.dataStore.setMessageReactionSummary(e);
                  }),
                  i.sendNotification(
                    "get_message_reaction_summaries_completed",
                    [t, n],
                  ));
              },
              function (e) {
                (k.iter(
                  n,
                  z((j = i.dataStore), j.removeLoadingMessageReactionSmmary),
                ),
                  i.sendNotification("get_message_reaction_summaries_failed"));
              },
            ));
        },
        getMessageReactionUsers: function (e, t) {
          var n = this;
          return this.wsApi.getMessageReactionUsers(e, t).then(
            function (e) {
              return (
                n.dataStore.setMessageReactionUsers(e),
                Promise.resolve(e)
              );
            },
            function (e) {
              return Promise.reject(e);
            },
          );
        },
        getAvailableMessageReactions: function (t) {
          var n = this;
          this.wsApi.getAvailableMessageReactions(t).then(
            function (e) {
              n.dataStore.updateAvailableMessageReactions(t, e);
            },
            function (e) {
              n.apiCaller.handleServerErrorDefault(e);
            },
          );
        },
        getFlowNotificationBadges: function (e) {
          var t = this;
          return this._getFlowNotificationBadges(e).then(
            function (e) {
              t.sendNotification("get_flow_notification_badges_completed");
            },
            function (e) {
              (t.apiCaller.handleServerErrorDefault(e),
                t.sendNotification("get_flow_notification_badges_failed"));
            },
          );
        },
        _getFlowNotificationBadges: function (e) {
          var i,
            t,
            r = this;
          return 0 == e.length
            ? Promise.resolve()
            : ((i = e.slice(0, 100)),
              (t = e.slice(100)),
              new Promise(function (t, n) {
                r.wsApi.getFlowNotificationBadges(i).then(
                  function (e) {
                    (k.iter(e, function (e) {
                      r.dataStore.setFlowNotificationBadge(e.domainId, e);
                    }),
                      Kt.resolveVoid(t));
                  },
                  function (e) {
                    n(e);
                  },
                );
              }).then(function (e) {
                return r._getFlowNotificationBadges(t);
              }));
        },
        getDirectApps: function (e) {
          var t = this;
          return 0 == e.length
            ? Promise.resolve([])
            : this.wsApi.getDirectApps(e).then(
                function (e) {
                  return (
                    k.iter(e, function (e) {
                      t.dataStore.setDirectApp(e.id, e);
                    }),
                    t.sendNotification("get_direct_apps_completed"),
                    e
                  );
                },
                function (e) {
                  return (
                    t.sendNotification("get_direct_apps_failed"),
                    t.sendNotification("error_occurred", e),
                    Promise.reject(e)
                  );
                },
              );
        },
        getFilePreview: function (e) {
          return this.wsApi.getFilePreview(e);
        },
        createFilePreview: function (t) {
          var n = this;
          return this.wsApi.createFilePreview(t).then(function (e) {
            return (n.dataStore.setFilePreview(t, e), e);
          });
        },
        __class__: zo,
      })),
      function () {
        ((this.callbacksWhenCreateSessionSucceed = []),
          (this.createSessionSucceed = !1),
          l.call(this, "apiCaller"));
      }),
    Wo =
      (((t["albero.proxy.ApiCallerProxy"] = Yo).__name__ = [
        "albero",
        "proxy",
        "ApiCallerProxy",
      ]),
      (Yo.delayForReplicationLag = function () {
        return Kt.delay(500);
      }),
      (Yo.__super__ = l),
      (Yo.prototype = e(l.prototype, {
        onSessionClear: function () {
          this.createSessionSucceed = !1;
        },
        onSessionCreated: function () {
          this.createSessionSucceed = !0;
        },
        fireCallbacksWaitingForSession: function () {
          (k.iter(this.callbacksWhenCreateSessionSucceed, function (e) {
            e();
          }),
            this.callbacksWhenCreateSessionSucceed.splice(
              0,
              this.callbacksWhenCreateSessionSucceed.length,
            ));
        },
        callImmediately: function (n, i) {
          var r = this;
          return new Promise(function (e, t) {
            r.rpc.call(n, i, e, t);
          });
        },
        callImmediatelyReliable: function (n, i, r, o) {
          null == o && (o = 3);
          var a = this;
          return this.callImmediately(n, i).catch(function (e) {
            var t = a.getRetryDelay(e, o, r);
            return h.isEmpty(t)
              ? Promise.reject(e)
              : Kt.delay(h.get(t)).then(function (e) {
                  return a.callImmediatelyReliable(n, i, r, o - 1);
                });
          });
        },
        call: function (l, u) {
          var c = this;
          return new Promise(function (e, t) {
            function n() {
              i(r, o, a, s);
            }
            var i = z((j = c.rpc), j.call),
              r = l,
              o = u,
              a = e,
              s = t;
            c.createSessionSucceed
              ? n()
              : c.callbacksWhenCreateSessionSucceed.push(n);
          });
        },
        callApiReliable: function (n, i, r, o) {
          null == o && (o = 3);
          var a = this;
          return this.call(n, i).catch(function (e) {
            var t = a.getRetryDelay(e, o, r);
            return h.isEmpty(t)
              ? Promise.reject(e)
              : Kt.delay(h.get(t)).then(function (e) {
                  return a.callApiReliable(n, i, r, o - 1);
                });
          });
        },
        getRetryDelay: function (e, t, n) {
          return 0 == t
            ? m.None
            : null == n
              ? 404 == e.code
                ? m.Some(this.retryTimeForReplicationLag())
                : m.None
              : n(e);
        },
        handleServerErrorDefault: function (e) {
          this.rpc.onServerError(e);
        },
        restart: function () {
          this.rpc.restart();
        },
        retryTimeForReplicationLag: function () {
          return 1e3;
        },
        __class__: Yo,
      })),
      function () {
        l.call(this, "apiNote");
      }),
    Ko =
      (((t["albero.proxy.ApiNoteProxy"] = Wo).__name__ = [
        "albero",
        "proxy",
        "ApiNoteProxy",
      ]),
      (Wo.__super__ = l),
      (Wo.prototype = e(l.prototype, {
        getNoteStatuses: function (e, r, t) {
          var o = this;
          return this.wsApi.getNoteStatuses(e, r, 60, t).then(
            function (e) {
              if (e.isNotEmptyNoteStatuses())
                for (var t = 0, n = e.noteStatuses; t < n.length; ) {
                  var i = n[t];
                  (++t,
                    o.dataStore.setNoteRevisionSummary(
                      i.noteId,
                      i.getNoteRevisionSummary(),
                    ));
                }
              return (
                o.sendNotification("get_note_statuses_loaded", {
                  talkId: r,
                  result: e,
                }),
                e
              );
            },
            function (e) {
              return (
                o.sendNotification("error_occurred", e),
                Promise.reject(e)
              );
            },
          );
        },
        getNote: function (t) {
          var n = this;
          return this.wsApi.getNote(t).then(
            function (e) {
              return (
                n.dataStore.setNote(e.getNoteId(), e.note),
                n.sendNotification("get_note_loaded", e),
                e
              );
            },
            function (e) {
              return (
                404 == e.code
                  ? (n.dataStore.setNote(t, null),
                    n.sendNotification("get_note_failed_by_note_not_found", t))
                  : n.sendNotification("error_occurred", e),
                Promise.reject(e)
              );
            },
          );
        },
        createNote: function (e, t, n, i, r, o) {
          var a = this;
          return this.wsApi.createNote(e, t, n, i, r).then(
            function (e) {
              var t;
              return (
                null != o && a.dataStore.setNoteLocalEdit(o.getNoteId(), null),
                (t =
                  null != o && o.hasEmitterKey()
                    ? e.updateWithEmitterKey(o.getEmitterKey())
                    : e),
                a.sendNotification("create_note_completed", t),
                e
              );
            },
            function (e) {
              var t, n;
              return (
                null != o &&
                  (o.setState(f.FAILED_BY_UNKNOWN),
                  a.dataStore.setNoteLocalEdit(o.getNoteId(), o)),
                null != o &&
                  o.hasEmitterKey() &&
                  ((t = a.dataStore.me.id),
                  (n = o.getEmitterKey()),
                  a.sendNotification("create_note_failed", new _t(t, e, n))),
                null != o &&
                  0 < o.getFileInfos().length &&
                  400 == e.code &&
                  "invalid file_id" == e.message &&
                  (o.setState(f.FAILED_BY_NOT_FOUND_FORWARDING),
                  a.dataStore.setNoteLocalEdit(o.getNoteId(), o)),
                null != o && 400 == e.code && "invalid talk_id" == e.message
                  ? ((t = De.getNoteForwardingForbiddenError()),
                    (n = _a.toLocalError(t, e)),
                    a.sendNotification("error_occurred", n))
                  : null != o &&
                      400 == e.code &&
                      "rich_text too long" == e.message
                    ? ((t = De.getNoteTextTooLongError()),
                      o.setState(f.FAILED_BY_TEXT_TOO_LONG),
                      a.dataStore.setNoteLocalEdit(o.getNoteId(), o),
                      (n = _a.toLocalError(t, e)),
                      a.sendNotification("error_occurred", n))
                    : null != o &&
                        400 == e.code &&
                        I.startsWith(
                          e.message,
                          "rich_text length must be less than or equal to ",
                        )
                      ? ((t = De.getNoteTextTooComplexError()),
                        o.setState(f.FAILED_BY_TEXT_TOO_COMPLEX),
                        a.dataStore.setNoteLocalEdit(o.getNoteId(), o),
                        (n = _a.toLocalError(t, e)),
                        a.sendNotification("error_occurred", n))
                      : a.sendNotification("error_occurred", e),
                Promise.reject(e)
              );
            },
          );
        },
        updateNoteSetting: function (t, e, n) {
          var i = this;
          return this.wsApi.updateNoteSetting(t, e, n).then(
            function (e) {
              return (
                i.sendNotification("update_note_setting_completed", e),
                e
              );
            },
            function (e) {
              return (
                404 == e.code
                  ? (i.dataStore.setNote(t, null),
                    i.sendNotification(
                      "update_note_setting_failed_by_not_found",
                      t,
                    ))
                  : 409 == e.code && "conflict" == e.message
                    ? (i.dataStore.setNote(t, null),
                      i.sendNotification(
                        "update_note_setting_failed_by_conflict",
                        t,
                      ))
                    : 409 == e.code &&
                        "locked by another user" == e.message &&
                        null != e.detail &&
                        null != e.detail.user_id &&
                        null != e.detail.device_id
                      ? i.sendNotification(
                          "update_note_setting_failed_by_editing",
                          t,
                        )
                      : i.sendNotification("error_occurred", e),
                Promise.reject(e)
              );
            },
          );
        },
        updateNote: function (n, e, t, i, r, o, a) {
          var s = this;
          return this.wsApi.updateNote(n, e, t, i, r, o).then(
            function (e) {
              return (
                null != a && s.dataStore.setNoteLocalEdit(n, null),
                s.dataStore.setNote(n, e.note),
                s.sendNotification("update_note_completed", e),
                e
              );
            },
            function (e) {
              var t;
              return (
                404 == e.code
                  ? null != a
                    ? (a.setState(f.FAILED_BY_NOT_FOUND),
                      s.dataStore.setNoteLocalEdit(n, a))
                    : s.dataStore.setNote(n, null)
                  : 409 == e.code && "conflict" == e.message
                    ? null != a
                      ? (a.setState(f.FAILED_BY_CONFLICT),
                        s.dataStore.setNoteLocalEdit(n, a))
                      : s.dataStore.setNote(n, null)
                    : 409 == e.code &&
                        "locked by another user" == e.message &&
                        null != e.detail &&
                        null != e.detail.user_id &&
                        null != e.detail.device_id
                      ? null != a &&
                        (a.setState(f.FAILED_BY_CONFLICT),
                        s.dataStore.setNoteLocalEdit(n, a))
                      : null != a &&
                          400 == e.code &&
                          "rich_text too long" == e.message
                        ? ((t = De.getNoteTextTooLongError()),
                          a.setState(f.FAILED_BY_TEXT_TOO_LONG),
                          s.dataStore.setNoteLocalEdit(n, a),
                          (t = _a.toLocalError(t, e)),
                          s.sendNotification("error_occurred", t))
                        : null != a &&
                            400 == e.code &&
                            I.startsWith(
                              e.message,
                              "rich_text length must be less than or equal to ",
                            )
                          ? ((t = De.getNoteTextTooComplexError()),
                            a.setState(f.FAILED_BY_TEXT_TOO_COMPLEX),
                            s.dataStore.setNoteLocalEdit(n, a),
                            (t = _a.toLocalError(t, e)),
                            s.sendNotification("error_occurred", t))
                          : (null != a &&
                              (a.setState(f.FAILED_BY_UNKNOWN),
                              s.dataStore.setNoteLocalEdit(n, a)),
                            s.sendNotification("error_occurred", e)),
                Promise.reject(e)
              );
            },
          );
        },
        deleteNote: function (t) {
          var n = this;
          return this.wsApi.deleteNote(t).then(
            function (e) {
              return (n.sendNotification("delete_note_completed", e), e);
            },
            function (e) {
              return (
                404 == e.code && n.dataStore.setNote(t, null),
                n.sendNotification("error_occurred", e),
                Promise.reject(e)
              );
            },
          );
        },
        lockNote: function (n, r, o) {
          var a = this;
          return this.wsApi
            .lockNote(r, o)
            .then(function (e) {
              var t,
                n,
                i = a.dataStore.getNote(r);
              return (
                null != i &&
                  ((t = a.dataStore.me.id),
                  (n = a.settings.getDeviceId()),
                  i.lock(o, t, n, e)),
                new Ut(To.COMPLETED(new xt(e)))
              );
            })
            .catch(function (e) {
              if (400 == e.code) F.v("invalid parameter.");
              else if (403 == e.code) F.v("forbidden.");
              else {
                if (404 == e.code)
                  return (
                    F.v("note is not found."),
                    a.dataStore.setNote(r, null),
                    new Ut(To.FAILED_BY_NOT_FOUND)
                  );
                var t;
                if (
                  409 == e.code &&
                  "locked by another user" == e.message &&
                  null != e.detail &&
                  null != e.detail.user_id &&
                  null != e.detail.device_id
                )
                  return (
                    (t = a.dataStore.getUser(n, e.detail.user_id)),
                    new Ut(To.FAILED_BY_LOCKED(t))
                  );
                if (409 == e.code && "conflict" == e.message)
                  return (F.v("conflict"), new Ut(To.FAILED_BY_CONFLICT));
              }
              return (
                a.sendNotification("error_occurred", e),
                new Ut(To.FAILED_BY_UNKNOWN)
              );
            });
        },
        unlockNote: function (n, i, r) {
          var o = this;
          return this.wsApi
            .unlockNote(i, r)
            .then(function (e) {
              var t = o.dataStore.getNote(i);
              return (
                null != t && t.unlock(r, o.dataStore.me.id),
                new Ut(To.RELEASED)
              );
            })
            .catch(function (e) {
              if (400 == e.code) F.v("invalid parameter.");
              else if (403 == e.code) F.v("forbidden.");
              else {
                if (404 == e.code)
                  return (
                    F.v("note is not found."),
                    o.dataStore.setNote(i, null),
                    new Ut(To.FAILED_BY_NOT_FOUND)
                  );
                if (409 == e.code && "conflict" == e.message)
                  return new Ut(To.FAILED_BY_CONFLICT);
                var t;
                if (
                  409 == e.code &&
                  "locked by another user" == e.message &&
                  null != e.detail &&
                  null != e.detail.user_id &&
                  null != e.detail.device_id
                )
                  return (
                    (t = o.dataStore.getUser(n, e.detail.user_id)),
                    new Ut(To.FAILED_BY_LOCKED(t))
                  );
              }
              return (
                o.sendNotification("error_occurred", e),
                new Ut(To.FAILED_BY_UNKNOWN)
              );
            });
        },
        __class__: Wo,
      })),
      function () {
        l.call(this, "appState");
      }),
    Vo =
      (((t["albero.proxy.AppStateProxy"] = Ko).__name__ = [
        "albero",
        "proxy",
        "AppStateProxy",
      ]),
      (Ko.__super__ = l),
      (Ko.prototype = e(l.prototype, {
        start: function () {
          (this.updateLastActivityAt(), this.checkInactiveInterval());
        },
        activateAppIfNeed: function () {
          (this.updateLastActivityAt(),
            this.appState == Dn.Inactive && this.setAppState(Dn.Active));
        },
        setAppState: function (e) {
          (F.v("APP_STATE_CHANGED: ", e),
            (this.appState = e),
            this.sendNotification("app_state_changed", this.appState));
        },
        updateLastActivityAt: function () {
          this.lastActivityAt = new Date();
        },
        checkInactiveInterval: function () {
          (this.checkInactive(),
            (new Xa(500).run = z(this, this.checkInactive)));
        },
        checkInactive: function () {
          this.appState == Dn.Inactive ||
            new Date().getTime() - this.lastActivityAt.getTime() < 2e3 ||
            this.setAppState(Dn.Inactive);
        },
        __class__: Ko,
      })),
      function () {
        var e = new E(-1, -1),
          e = ((this.dummyFileId = e), new E(-1, -1));
        ((this.dummyMsgId = e), l.call(this, "dataFactory"));
      }),
    qo =
      (((t["albero.proxy.DataFactoryProxy"] = Vo).__name__ = [
        "albero",
        "proxy",
        "DataFactoryProxy",
      ]),
      (Vo.__super__ = l),
      (Vo.prototype = e(l.prototype, {
        onRegister: function () {},
        newDummyMessage: function (e, t, n, i) {
          var r = this,
            o = new O(),
            e =
              ((o.id = this.dummyMsgId),
              (o.userId = this.dataStore.me.id),
              (o.talkId = e),
              (o.type = t),
              (o.content = n),
              (o.scheduleAt = i),
              this.dummyMsgId),
            a =
              ((e =
                null == e
                  ? null
                  : ((t = new E(0, 1)),
                    (n = (e.high - t.high) | 0),
                    (i = (e.low - t.low) | 0),
                    w.ucompare(e.low, t.low) < 0 && (n--, (n |= 0)),
                    new E(n, i))),
              (this.dummyMsgId = e),
              []);
          switch (o.type[1]) {
            case 4:
              a = [this.createDummyFileInfo(o, o.content)];
              break;
            case 5:
              a = G.__cast(o.content.files, Array).map(function (e) {
                return r.createDummyFileInfo(o, e);
              });
          }
          return (this.dataStore.setMessageWithFileInfos(o, a), o);
        },
        createDummyFileInfo: function (e, t) {
          var n,
            i,
            r,
            e = ar.fromMessageAndFileWithFileId(e, t, this.dummyFileId),
            t = this.dummyFileId;
          return (
            (t =
              null == t
                ? null
                : ((n = new E(0, 1)),
                  (i = (t.high - n.high) | 0),
                  (r = (t.low - n.low) | 0),
                  w.ucompare(t.low, n.low) < 0 && (i--, (i |= 0)),
                  new E(i, r))),
            (this.dummyFileId = t),
            e
          );
        },
        newDummyScheduledMessageEdit: function (e, t, n, i) {
          var r = [];
          switch (t[1]) {
            case 4:
              r = [this.createDummyFileInfoForScheduledMessage(n)];
              break;
            case 5:
              r = G.__cast(n.files, Array).map(
                z(this, this.createDummyFileInfoForScheduledMessage),
              );
          }
          return new Ot(t, n, r, i);
        },
        createDummyFileInfoForScheduledMessage: function (e) {
          var t,
            n,
            i,
            r = null == e.file_id ? this.dummyFileId : e.file_id,
            e = ar.fromScheduledMessageFileContentWithFileId(e, r),
            r = this.dummyFileId;
          return (
            (r =
              null == r
                ? null
                : ((t = new E(0, 1)),
                  (n = (r.high - t.high) | 0),
                  (i = (r.low - t.low) | 0),
                  w.ucompare(r.low, t.low) < 0 && (n--, (n |= 0)),
                  new E(n, i))),
            (this.dummyFileId = r),
            e
          );
        },
        __class__: Vo,
      })),
      function () {
        ((this.departmentStore = new Yi()),
          (this.stampStore = new Yr()),
          (this.questionStore = new Or()),
          (this.fileInfoStore = new lr()),
          (this.scheduledMessageStore = new mo()),
          (this.messageReactionStore = new ko()),
          (this.announcementStore = new Fi()),
          (this.messageStore = new pr()),
          (this.conferenceStore = new Pi()),
          (this.solutionStore = new zr()),
          (this.directAppStore = new Vi()),
          (this.flowNotificationBadgeStore = new ur()),
          l.call(this, "dataStore"));
      }),
    Qo =
      (((t["albero.proxy.DataStoreProxy"] = qo).__name__ = [
        "albero",
        "proxy",
        "DataStoreProxy",
      ]),
      (qo.__super__ = l),
      (qo.prototype = e(l.prototype, {
        onRegister: function () {
          this.init();
        },
        setMe: function (e) {
          this.me = e;
          for (var t = 0, n = this.getDomains(); t < n.length; ) {
            var i = n[t];
            (++t,
              this.sendNotification(
                "notify_update_user",
                this.me.toDomainUser(i.id),
              ));
          }
          (this.sendNotification("current_user_changed", e),
            this.storageForUsers.setDirtyFlag());
        },
        getMe: function () {
          return this.me;
        },
        isMe: function (e) {
          var t;
          return (
            null != this.me &&
            null != (t = this.me.id) &&
            null != e &&
            t.high == e.high &&
            t.low == e.low
          );
        },
        getFriendsOrAcquaintances: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.users,
            t = null != Y[e] ? t.getReserved(e) : t.h[e],
            n = [];
          if (null != t)
            for (var i = new is(t, t.arrayKeys()); i.hasNext(); ) {
              var r = i.next();
              (0 != r.type && 1 != r.type) || n.push(r.user);
            }
          return n;
        },
        getFriends: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.users,
            t = null != Y[e] ? t.getReserved(e) : t.h[e],
            n = [];
          if (null != t)
            for (var i = new is(t, t.arrayKeys()); i.hasNext(); ) {
              var r = i.next();
              0 == r.type && n.push(r.user);
            }
          return n;
        },
        getAcquaintances: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.users,
            t = null != Y[e] ? t.getReserved(e) : t.h[e],
            n = [];
          if (null != t)
            for (var i = new is(t, t.arrayKeys()); i.hasNext(); ) {
              var r = i.next();
              1 == r.type && n.push(r.user);
            }
          return n;
        },
        isFriend: function (e, t) {
          var e = "_" + e.high + "_" + e.low,
            n = this.users,
            n = null != Y[e] ? n.getReserved(e) : n.h[e];
          return (
            null != n &&
            ((e = "_" + t.high + "_" + t.low),
            null != (t = null != Y[e] ? n.getReserved(e) : n.h[e])) &&
            0 == t.type
          );
        },
        addFriend: function (e) {
          var t = this.ensureDomainUserMap(e.domainId),
            n = e.id,
            n = "_" + n.high + "_" + n.low,
            i = { type: 0, user: e };
          return (
            null != Y[n] ? t.setReserved(n, i) : (t.h[n] = i),
            this.storageForUsers.setDirtyFlag(),
            e
          );
        },
        addAcquaintance: function (e) {
          var t = this.ensureDomainUserMap(e.domainId),
            n = e.id,
            n = "_" + n.high + "_" + n.low,
            i = { type: 1, user: e };
          return (
            null != Y[n] ? t.setReserved(n, i) : (t.h[n] = i),
            this.storageForUsers.setDirtyFlag(),
            e
          );
        },
        ensureDomainUserMap: function (e) {
          var t,
            e = "_" + e.high + "_" + e.low,
            n = this.users,
            n = null != Y[e] ? n.getReserved(e) : n.h[e];
          return (
            null == n &&
              ((n = new B()),
              (t = this.users),
              null != Y[e] ? t.setReserved(e, n) : (t.h[e] = n)),
            n
          );
        },
        setUserIfLatest: function (e, t) {
          null == e && (e = this.ensureDomainUserMap(t.domainId));
          var n,
            i,
            r,
            o = t.id,
            o = "_" + o.high + "_" + o.low,
            a = null != Y[o] ? e.getReserved(o) : e.h[o];
          (null != a
            ? ((r = a.user.updatedAt),
              0 == (n = ((i = t.updatedAt).high - r.high) | 0) &&
                (n = w.ucompare(i.low, r.low)),
              (0 <
                (i =
                  i.high < 0 ? (r.high < 0 ? n : -1) : 0 <= r.high ? n : 1) ||
                (2 == a.type && 0 != i)) &&
                this.sendNotification("notify_update_user", (a.user = t)))
            : ((r = { type: 2, user: t }),
              null != Y[o] ? e.setReserved(o, r) : (e.h[o] = r)),
            this.storageForUsers.setDirtyFlag());
        },
        setUserPresence: function (e) {
          null == this.userPresences && (this.userPresences = new B());
          var t = e.userId,
            t = "_" + t.high + "_" + t.low,
            n = this.userPresences;
          null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e);
        },
        getUserPresence: function (e) {
          var t;
          return null == this.userPresences
            ? null
            : ((e = "_" + e.high + "_" + e.low),
              (t = this.userPresences),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        isDepartmentsChanged: function (e, t) {
          return !D.eqArray(e.departments, t.departments);
        },
        setUsersIfLatest: function (e, t) {
          var e = this.ensureDomainUserMap(e),
            n = z(this, this.setUserIfLatest),
            i = e;
          k.iter(t, function (e) {
            n(i, e);
          });
        },
        clearDomainUsers: function (e) {
          this.users.remove("_" + e.high + "_" + e.low) &&
            this.storageForUsers.setDirtyFlag();
        },
        clearUsersDepartments: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.users,
            t = null != Y[e] ? t.getReserved(e) : t.h[e];
          if (null != t)
            for (var n = new is(t, t.arrayKeys()); n.hasNext(); ) {
              var i = n.next();
              ((i.user.departments = null),
                this.sendNotification("notify_update_user", i.user));
            }
          this.storageForUsers.setDirtyFlag();
        },
        removeFriend: function (e, t) {
          var e = "_" + e.high + "_" + e.low,
            n = this.users,
            n = null != Y[e] ? n.getReserved(e) : n.h[e];
          (null != n &&
            ((e = "_" + t.high + "_" + t.low),
            null != (t = null != Y[e] ? n.getReserved(e) : n.h[e])) &&
            (t.type = 2),
            this.storageForUsers.setDirtyFlag());
        },
        removeAcquaintance: function (e, t) {
          this.removeFriend(e, t);
        },
        getUser: function (e, t) {
          if (null != t) {
            var n = new E(0, 0);
            if (t.high != n.high || t.low != n.low) {
              if (
                null != this.me &&
                null != (n = this.me.id) &&
                null != t &&
                n.high == t.high &&
                n.low == t.low
              )
                return this.me.toDomainUser(e);
              ((n = "_" + e.high + "_" + e.low),
                (e = this.users),
                (e = null != Y[n] ? e.getReserved(n) : e.h[n]));
              if (null != e) {
                ((n = "_" + t.high + "_" + t.low),
                  (t = null != Y[n] ? e.getReserved(n) : e.h[n]));
                if (null != t) return t.user;
              }
            }
          }
          return null;
        },
        getUserStatus: function (e, t) {
          e = this.getUser(e, t);
          return null == e ? "" : e.status;
        },
        getUsers: function (e, t) {
          function n(n, e) {
            var t = "_" + n.high + "_" + n.low,
              i = s.users,
              r = null != Y[t] ? i.getReserved(t) : i.h[t];
            if (null == r) return [];
            if (null != e)
              return e.map(function (e) {
                var t = new E(0, 0);
                return e.high == t.high && e.low == t.low
                  ? null
                  : null != (t = s.me.id) &&
                      null != e &&
                      t.high == e.high &&
                      t.low == e.low
                    ? s.me.toDomainUser(n)
                    : null !=
                        (e =
                          null != Y[(t = "_" + e.high + "_" + e.low)]
                            ? r.getReserved(t)
                            : r.h[t])
                      ? e.user
                      : null;
              });
            for (var o = [], i = r.arrayKeys(), a = new is(r, i); a.hasNext(); )
              o.push(a.next().user);
            return o;
          }
          var s = this;
          if (null != e) return n(e, t);
          for (
            var e = this.domains, i = new is(e, e.arrayKeys()), r = null;
            i.hasNext();
          )
            r = null == r ? n(i.next().id, t) : r.concat(n(i.next().id, t));
          return null == r ? [] : r;
        },
        getTalk: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.talks;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        getValidPairTalk: function (n, i) {
          var r = this;
          return null == this.me
            ? null
            : k.find(this.getTalks(), function (e) {
                var t;
                return (
                  !!x.enumEq(e.type, Xr.PairTalk) &&
                  null != (t = e.domainId) &&
                  null != i &&
                  t.high == i.high &&
                  t.low == i.low &&
                  !!k.exists(e.userIds, function (e) {
                    return (
                      null != e &&
                      null != n &&
                      e.high == n.high &&
                      e.low == n.low
                    );
                  }) &&
                  !!k.exists(e.userIds, function (e) {
                    var t = r.me.id;
                    return (
                      null != e &&
                      null != t &&
                      e.high == t.high &&
                      e.low == t.low
                    );
                  })
                );
              });
        },
        setTalk: function (e) {
          var t = e.id,
            t = "_" + t.high + "_" + t.low,
            n = this.talks;
          if (
            (null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e),
            null != e.leftUsers)
          )
            for (
              var i = this.ensureDomainUserMap(e.domainId),
                r = 0,
                o = e.leftUsers;
              r < o.length;
            ) {
              var a,
                s,
                l,
                u = o[r],
                c = (++r, u.id),
                c = "_" + c.high + "_" + c.low,
                _ = null != Y[c] ? i.getReserved(c) : i.h[c];
              null != _
                ? 2 == _.type &&
                  ((a = _.user.updatedAt),
                  0 == (s = ((l = u.updatedAt).high - a.high) | 0) &&
                    (s = w.ucompare(l.low, a.low)),
                  0 <
                    (l.high < 0
                      ? a.high < 0
                        ? s
                        : -1
                      : 0 <= a.high
                        ? s
                        : 1)) &&
                  ((_.user = u),
                  this.storageForUsers.setDirtyFlag(),
                  this.sendNotification("notify_update_user", u))
                : ((l = { type: 2, user: u }),
                  null != Y[c] ? i.setReserved(c, l) : (i.h[c] = l),
                  this.storageForUsers.setDirtyFlag(),
                  this.sendNotification("notify_update_user", u));
            }
          return (this.storageForTalks.setDirtyFlag(), e);
        },
        getTalks: function () {
          for (var e = [], t = this.talks.keys(); t.hasNext(); ) {
            var n = t.next(),
              i = this.talks;
            e.push(null != Y[n] ? i.getReserved(n) : i.h[n]);
          }
          return e;
        },
        setTalks: function (e) {
          k.iter(e, z(this, this.setTalk));
        },
        removeTalk: function (e) {
          var t, n;
          (null != this.domainUnreadCounts &&
            ((n = "_" + e.high + "_" + e.low),
            (t = this.talks),
            null != (t = null != Y[n] ? t.getReserved(n) : t.h[n])) &&
            ((n = t.domainId),
            this.domainUnreadCounts.remove("_" + n.high + "_" + n.low)),
            this.questionStore.removeTalkQuestions(e),
            this.talks.remove("_" + e.high + "_" + e.low),
            this.storageForTalks.setDirtyFlag());
        },
        getTalkStatuses: function () {
          for (var e = [], t = this.talkStatuses.keys(); t.hasNext(); ) {
            var n = t.next(),
              i = this.talks;
            (null != Y[n] ? i.existsReserved(n) : i.h.hasOwnProperty(n)) &&
              ((i = this.talkStatuses),
              e.push(null != Y[n] ? i.getReserved(n) : i.h[n]));
          }
          return e;
        },
        getTalkStatus: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.talkStatuses;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        setTalkStatus: function (e) {
          var t = e.id,
            t = "_" + t.high + "_" + t.low,
            n =
              (null != this.domainUnreadCounts &&
                ((n = this.talks),
                null != (n = null != Y[t] ? n.getReserved(t) : n.h[t])) &&
                ((n = n.domainId),
                this.domainUnreadCounts.remove("_" + n.high + "_" + n.low)),
              this.talkStatuses);
          return (
            null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e),
            this.sendNotification("brand_badge_changed"),
            this.storageForTalkStatuses.setDirtyFlag(),
            e
          );
        },
        removeTalkStatus: function (e) {
          var t,
            e = "_" + e.high + "_" + e.low;
          (null != this.domainUnreadCounts &&
            ((t = this.talks),
            null != (t = null != Y[e] ? t.getReserved(e) : t.h[e])) &&
            ((t = t.domainId),
            this.domainUnreadCounts.remove("_" + t.high + "_" + t.low)),
            this.talkStatuses.remove(e),
            this.storageForTalkStatuses.setDirtyFlag());
        },
        getDomains: function () {
          for (
            var e = [], t = this.domains, n = new is(t, t.arrayKeys());
            n.hasNext();
          ) {
            var i = n.next();
            e.push(i);
          }
          return e;
        },
        getDomain: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.domains;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        setDomainIfLatest: function (e) {
          var t,
            n,
            i = e.id,
            i = "_" + i.high + "_" + i.low,
            r = this.domains,
            r = null != Y[i] ? r.getReserved(i) : r.h[i],
            o =
              null == r ||
              ((n = r.updatedAt),
              0 == (t = ((o = e.updatedAt).high - n.high) | 0) &&
                (t = w.ucompare(o.low, n.low)),
              0 <= (o.high < 0 ? (n.high < 0 ? t : -1) : 0 <= n.high ? t : 1));
          return o
            ? ((n = this.domains),
              null != Y[i] ? n.setReserved(i, e) : (n.h[i] = e),
              this.storage.setDirtyFlag(),
              e)
            : r;
        },
        removeDomain: function (t) {
          (this.domains.remove("_" + t.high + "_" + t.low),
            this.storage.setDirtyFlag());
          for (
            var e = this.getTalks().filter(function (e) {
                e = e.domainId;
                return (
                  null != e && null != t && e.high == t.high && e.low == t.low
                );
              }),
              n = 0;
            n < e.length;
          ) {
            var i = e[n];
            (++n, this.removeTalk(i.id), this.removeTalkStatus(i.id));
          }
        },
        getUnreadCount: function () {
          for (var e = 0, t = this.domains.keys(); t.hasNext(); ) {
            var n = t.next();
            e += this._getDomainUnreadCount(n);
          }
          return e;
        },
        _getDomainUnreadCount: function (e) {
          if (
            null != this.domainUnreadCounts &&
            ((n = this.domainUnreadCounts),
            null != Y[e] ? n.existsReserved(e) : n.h.hasOwnProperty(e))
          )
            return (
              (n = this.domainUnreadCounts),
              null != Y[e] ? n.getReserved(e) : n.h[e]
            );
          null == this.domainUnreadCounts &&
            (this.domainUnreadCounts = new B());
          for (
            var t = 0, n = this.talkStatuses, i = new is(n, n.arrayKeys());
            i.hasNext();
          ) {
            var r,
              o = i.next(),
              a = o.calcUnreadCount();
            0 != a &&
              ((o = "_" + (o = o.id).high + "_" + o.low),
              (r = this.talks),
              null != (r = null != Y[o] ? r.getReserved(o) : r.h[o]) &&
                "_" + (o = r.domainId).high + "_" + o.low == e) &&
              (t += a);
          }
          var n = this.announcementStatuses,
            n = null != Y[e] ? n.getReserved(e) : n.h[e],
            s =
              (null != n &&
                ((s = this.domains),
                null != (s = null != Y[e] ? s.getReserved(e) : s.h[e])) &&
                s.role.allowReadAnnouncements &&
                (t += n.calcUnreadCount()),
              this.domainUnreadCounts);
          return (null != Y[e] ? s.setReserved(e, t) : (s.h[e] = t), t);
        },
        getDomainUnreadCount: function (e) {
          return null == e
            ? 0
            : this._getDomainUnreadCount("_" + e.high + "_" + e.low);
        },
        getDomainInvites: function () {
          for (var e = [], t = this.domainInvites.keys(); t.hasNext(); ) {
            var n = t.next(),
              i = this.domainInvites;
            e.push(null != Y[n] ? i.getReserved(n) : i.h[n]);
          }
          return e;
        },
        getDomainInvitesWithoutWaiting: function () {
          var t = this;
          return this.getDomainInvites().filter(function (e) {
            return (
              null == e.accountControlRequestId ||
              null == t.getAccountControlRequest(e.accountControlRequestId)
            );
          });
        },
        getDomainInvitesCount: function () {
          return (
            null == this.accountControlGroup
              ? this.getDomainInvitesWithoutWaiting()
              : this.getDomainInvites()
          ).length;
        },
        getDomainInvite: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.domainInvites;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        setDomainInvite: function (e) {
          var t = e.id,
            t = "_" + t.high + "_" + t.low,
            n = this.domainInvites;
          return (
            null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e),
            this.sendNotification("brand_badge_changed"),
            this.storage.setDirtyFlag(),
            e
          );
        },
        removeDomainInvite: function (e) {
          (this.domainInvites.remove("_" + e.high + "_" + e.low),
            this.sendNotification("brand_badge_changed"),
            this.storage.setDirtyFlag());
        },
        clearFlowNotificationBadge: function () {
          this.flowNotificationBadgeStore.clear();
        },
        setFlowNotificationBadge: function (e, t) {
          var n = this.flowNotificationBadgeStore.getBadge(e);
          (null == n || n.version < t.version) &&
            this.flowNotificationBadgeStore.setBadge(e, t);
        },
        removeFlowNotificationBadge: function (e) {
          this.flowNotificationBadgeStore.removeBadge(e);
        },
        isFlowNotificationBadgeShownForDomains: function (e) {
          if (null != this.domains)
            for (
              var t = this.domains, n = new is(t, t.arrayKeys());
              n.hasNext();
            ) {
              var i = n.next(),
                r = i.id;
              if (
                (null == r ||
                  null == e ||
                  r.high != e.high ||
                  r.low != e.low) &&
                this.isFlowNotificationBadgeShownForOneDomain(i.id)
              )
                return !0;
            }
          return !1;
        },
        isFlowNotificationBadgeShownForOneDomain: function (e) {
          return (
            !!this.flowNotificationBadgeStore.isBadgeShownForOneDomain(e) &&
            this.isDirectAppExists(e)
          );
        },
        isDirectAppExists: function (e) {
          return (
            0 != this.getDomainDirectApps(e).getSortedLinkFeatures().length
          );
        },
        getFlowNotificationBadge: function (e) {
          return this.flowNotificationBadgeStore.getBadge(e);
        },
        getQuestions: function (e, t) {
          return this.questionStore.getQuestions(e, t);
        },
        getQuestion: function (e) {
          return this.questionStore.getQuestion(e);
        },
        setQuestionIfLatest: function (e) {
          var t,
            n,
            i = this.getQuestion(e.id),
            r =
              null != i &&
              ((r = i.updatedAt),
              (t = e.updatedAt),
              0 == (n = (r.high - t.high) | 0) &&
                (n = w.ucompare(r.low, t.low)),
              0 <= (r.high < 0 ? (t.high < 0 ? n : -1) : 0 <= t.high ? n : 1));
          return r ? i : (this.questionStore.setQuestion(e), e);
        },
        getAnnouncementStatus: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.announcementStatuses;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        setAnnouncementStatus: function (e) {
          var t = e.domainId,
            t = "_" + t.high + "_" + t.low,
            n = this.announcementStatuses;
          return (
            null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e),
            null != this.domainUnreadCounts &&
              this.domainUnreadCounts.remove(t),
            this.sendNotification("brand_badge_changed"),
            this.storage.setDirtyFlag(),
            e
          );
        },
        removeAnnouncementStatus: function (e) {
          e = "_" + e.high + "_" + e.low;
          (this.announcementStatuses.remove(e),
            null != this.domainUnreadCounts &&
              this.domainUnreadCounts.remove(e),
            this.sendNotification("brand_badge_changed"),
            this.storage.setDirtyFlag());
        },
        getAccountControlRequests: function () {
          for (
            var e = [],
              t = this.accountControlRequests,
              n = new is(t, t.arrayKeys());
            n.hasNext();
          ) {
            var i = n.next();
            e.push(i);
          }
          return e;
        },
        getAccountControlRequest: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.accountControlRequests;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        setAccountControlRequest: function (e) {
          var t = e.id,
            t = "_" + t.high + "_" + t.low,
            n = this.accountControlRequests,
            n = null != Y[t] ? n.getReserved(t) : n.h[t];
          return null != n && n.version > e.version
            ? n
            : ((n = this.accountControlRequests),
              null != Y[t] ? n.setReserved(t, e) : (n.h[t] = e),
              this.storage.setDirtyFlag(),
              e);
        },
        removeAccountControlRequest: function (e) {
          e = "_" + e.high + "_" + e.low;
          (this.accountControlRequests.remove(e), this.storage.setDirtyFlag());
        },
        removeAccountControlRequests: function () {
          ((this.accountControlRequests = new B()),
            this.storage.setDirtyFlag());
        },
        getAccountControlGroup: function () {
          return this.accountControlGroup;
        },
        setAccountControlGroup: function (e) {
          return (
            (this.accountControlGroup = e),
            this.storage.setDirtyFlag(),
            this.accountControlGroup
          );
        },
        updateAccountControlGroup: function (e) {
          var t, n, i;
          null != this.accountControlGroup &&
            ((n = this.accountControlGroup.id), (t = e.id), null != n) &&
            null != t &&
            n.high == t.high &&
            n.low == t.low &&
            (null != (n = e.group) &&
              this.accountControlGroup.group.version < n.version &&
              ((this.accountControlGroup.group = n),
              this.storage.setDirtyFlag()),
            null != (i = e.profilePolicy)) &&
            this.accountControlGroup.profilePolicy.version < i.version &&
            ((this.accountControlGroup.profilePolicy = i),
            this.storage.setDirtyFlag());
        },
        removeAccountControlGroup: function (e) {
          var t;
          null != this.accountControlGroup &&
            null != (t = this.accountControlGroup.id) &&
            null != e &&
            t.high == e.high &&
            t.low == e.low &&
            ((this.accountControlGroup = null), this.storage.setDirtyFlag());
        },
        addReactedConfereceId: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.reactedConferenceIds;
          (null != Y[e] ? t.setReserved(e, !0) : (t.h[e] = !0),
            this.storage.setDirtyFlag());
        },
        removeReactedConfereceId: function (e) {
          (this.reactedConferenceIds.remove("_" + e.high + "_" + e.low),
            this.storage.setDirtyFlag());
        },
        isReactedConferenceId: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.reactedConferenceIds;
          return null != Y[e] ? t.existsReserved(e) : t.h.hasOwnProperty(e);
        },
        filterReactedConferenceIds: function (e) {
          var n = this,
            t = new B();
          (k.iter(
            e
              .map(function (e) {
                return "_" + e.high + "_" + e.low;
              })
              .filter(function (e) {
                var t = n.reactedConferenceIds;
                return null != Y[e]
                  ? t.existsReserved(e)
                  : t.h.hasOwnProperty(e);
              }),
            function (e) {
              null != Y[e] ? t.setReserved(e, !0) : (t.h[e] = !0);
            },
          ),
            (this.reactedConferenceIds = t));
        },
        getNote: function (e) {
          var t;
          return null == this.notes
            ? null
            : ((t = this.notes),
              (e = e.toString()),
              (t = t),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        setNote: function (e, t) {
          var n;
          return (
            null == this.notes && (this.notes = new B()),
            null == t
              ? (this.fileInfoStore.setNoteFileInfos(e, null),
                this.notes.remove(e.toString()))
              : (this.fileInfoStore.setNoteFileInfos(
                  e,
                  t.noteRevision.contentFiles,
                ),
                (n = this.notes),
                (e = e.toString()),
                (n = n),
                null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t)),
            t
          );
        },
        setNoteRevisionSummary: function (e, t) {
          var n, i;
          (null == this.noteRevisionSummaries &&
            (this.noteRevisionSummaries = new B()),
            null == t
              ? (this.noteRevisionSummaries.remove(e.toString()),
                this.fileInfoStore.setNoteFileInfos(e, null))
              : ((i = this.noteRevisionSummaries),
                (n = e.toString()),
                null != Y[n] ? i.setReserved(n, t) : (i.h[n] = t),
                this.fileInfoStore.setNoteFileInfos(e, t.contentFiles)));
        },
        getNoteRevisionSummary: function (e) {
          var t;
          return null == this.noteRevisionSummaries
            ? null
            : ((t = this.noteRevisionSummaries),
              (e = e.toString()),
              (t = t),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        setNoteLocalEdit: function (e, t) {
          if (
            (null == this.noteLocalEdits && (this.noteLocalEdits = new B()),
            null == t)
          )
            return (
              (i = this.noteLocalEdits),
              (n = e.toString()),
              (i = i),
              null == (i = null != Y[n] ? i.getReserved(n) : i.h[n])
                ? void 0
                : (this.noteLocalEdits.remove(e.toString()),
                  this.fileInfoStore.setNoteFileInfosForEdit(e, null),
                  void this.sendNotification("clear_note_local_edit", i))
            );
          var n = this.noteLocalEdits,
            i = e.toString();
          (null != Y[i] ? n.setReserved(i, t) : (n.h[i] = t),
            this.fileInfoStore.setNoteFileInfosForEdit(e, t.getFileInfos()),
            this.sendNotification("update_note_local_edit", t));
        },
        getNoteLocalEdit: function (e) {
          var t;
          return null == this.noteLocalEdits
            ? null
            : ((t = this.noteLocalEdits),
              (e = e.toString()),
              (t = t),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        getNoteLocalEdits: function (e) {
          var t = [];
          if (null != this.noteLocalEdits)
            for (
              var n = this.noteLocalEdits, i = new is(n, n.arrayKeys());
              i.hasNext();
            ) {
              var r = i.next(),
                o = r.getTalkId();
              null != o &&
                null != e &&
                o.high == e.high &&
                o.low == e.low &&
                t.push(r);
            }
          return t;
        },
        setDirectApp: function (e, t) {
          this.directAppStore.setDirectApp(e, t);
        },
        removeDirectApp: function (e) {
          this.directAppStore.removeDirectApp(e);
        },
        getDirectApp: function (e) {
          return this.directAppStore.getDirectApp(e);
        },
        getDomainDirectApps: function (e) {
          ((e = this.getDomain(e)),
            (e = null == e ? [] : e.directAppsSetting.enabled));
          return this.directAppStore.getDirectApps(e);
        },
        getAllDirectApps: function () {
          return this.directAppStore.getAllDirectApps();
        },
        setMessageWithFileInfos: function (e, t) {
          (0 != t.length && this.fileInfoStore.setMessageFileInfos(e.id, t),
            this.messageStore.setMessage(e));
        },
        setMessage: function (e) {
          var t = [];
          switch (e.type[1]) {
            case 4:
              t = [ar.fromMessageAndFile(e, e.content)];
              break;
            case 5:
              var n = e,
                t = Un.asArray(e.content.files).map(function (e) {
                  return ar.fromMessageAndFile(n, e);
                });
          }
          this.setMessageWithFileInfos(e, t);
        },
        getMessage: function (e) {
          return this.messageStore.getMessage(e);
        },
        removeMessage: function (e) {
          (this.fileInfoStore.removeMessageFileInfos(e),
            this.messageStore.removeMessage(e));
        },
        setFromThisDevice: function (e) {
          this.messageStore.setFromThisDevice(e);
        },
        isFromThisDevice: function (e) {
          return this.messageStore.isFromThisDevice(e);
        },
        closeAction: function (e) {
          this.messageStore.closeAction(e);
        },
        deleteFile: function (e) {
          (this.messageStore.deleteFile(e), this.fileInfoStore.deleteFile(e));
        },
        deleteFavoriteMessage: function (e) {
          this.messageStore.deleteFavoriteMessage(e);
        },
        addFavoriteMessage: function (e) {
          this.messageStore.addFavoriteMessage(e);
        },
        getScheduledMessage: function (e) {
          return this.scheduledMessageStore.getScheduledMessage(e);
        },
        getScheduledMessagesInTalk: function (t) {
          return this.scheduledMessageStore
            .getScheduledMessages()
            .filter(function (e) {
              e = e.talkId;
              return (
                null != e && null != t && e.high == t.high && e.low == t.low
              );
            });
        },
        getScheduledMessagesInDomain: function (t) {
          var n = this;
          return this.scheduledMessageStore
            .getScheduledMessages()
            .filter(function (e) {
              var e = n.getTalk(e.talkId);
              return (
                null != e &&
                null != (e = e.domainId) &&
                null != t &&
                e.high == t.high &&
                e.low == t.low
              );
            });
        },
        setScheduledMessage: function (e) {
          var t = [];
          switch (e.type[1]) {
            case 4:
              t = [ar.fromScheduledMessageFileContent(e.content)];
              break;
            case 5:
              t = Un.asArray(e.content.files).map(
                ar.fromScheduledMessageFileContent,
              );
          }
          (0 != t.length &&
            this.fileInfoStore.setScheduledMessageFileInfos(e.id, t),
            this.scheduledMessageStore.setScheduledMessage(e));
        },
        removeScheduledMessage: function (e) {
          return (
            this.fileInfoStore.removeScheduledMessageFileInfos(e),
            this.removeScheduledMessageEdit(e),
            this.scheduledMessageStore.removeScheduledMessage(e)
          );
        },
        removeScheduledMessagesInDomain: function (n) {
          var i = this;
          k.iter(
            this.scheduledMessageStore.getScheduledMessages(),
            function (e) {
              var t = i.getTalk(e.talkId);
              null != t &&
                null != (t = t.domainId) &&
                null != n &&
                t.high == n.high &&
                t.low == n.low &&
                i.removeScheduledMessage(e.id);
            },
          );
        },
        getScheduledMessageEdit: function (e) {
          return this.scheduledMessageStore.getScheduledMessageEdit(e);
        },
        setScheduledMessageEdit: function (e, t) {
          this.scheduledMessageStore.setScheduledMessageEdit(e, t);
        },
        removeScheduledMessageEdit: function (e) {
          return this.scheduledMessageStore.removeScheduledMessageEdit(e);
        },
        getAnnouncement: function (e) {
          return this.announcementStore.getAnnouncement(e);
        },
        setAnnouncement: function (e) {
          this.announcementStore.setAnnouncement(e);
        },
        removeAnnouncement: function (e) {
          this.announcementStore.removeAnnouncement(e);
        },
        updateAvailableMessageReactions: function (e, t) {
          this.messageReactionStore.updateAvailableMessageReactions(e, t);
        },
        getAvailableMessageReactions: function (e) {
          return this.messageReactionStore.getAvailableMessageReactions(e);
        },
        isAvailableMessageReactionsLoaded: function (e) {
          return this.messageReactionStore.isAvailableMessageReactionsLoaded(e);
        },
        setMessageReactionSummary: function (e) {
          this.messageReactionStore.setMessageReactionSummary(e);
        },
        removeMessageReactionSummary: function (e) {
          this.messageReactionStore.removeMessageReactionSummary(e);
        },
        getMessageReactionSummary: function (e) {
          return this.messageReactionStore.getMessageReactionSummary(e);
        },
        removeMessageReactionUsers: function (e) {
          this.messageReactionStore.removeMessageReactionUsers(e);
        },
        getMessageReactionUsers: function (e) {
          return this.messageReactionStore.getMessageReactionUsers(e);
        },
        setMessageReactionUsers: function (e) {
          this.messageReactionStore.setMessageReactionUsers(e);
        },
        addLoadingMessageReactionSmmary: function (e) {
          this.messageReactionStore.addLoadingMessageReactionSmmary(e);
        },
        removeLoadingMessageReactionSmmary: function (e) {
          this.messageReactionStore.removeLoadingMessageReactionSmmary(e);
        },
        isLoadingMessageReactionSmmary: function (e) {
          return this.messageReactionStore.isLoadingMessageReactionSmmary(e);
        },
        setSolution: function (e) {
          this.solutionStore.setSolution(e);
        },
        getSolution: function (e) {
          return this.solutionStore.getSolution(e);
        },
        addLoadingSolutionId: function (e) {
          this.solutionStore.addLoadingSolutionId(e);
        },
        removeLoadingSolutionId: function (e) {
          this.solutionStore.removeLoadingSolutionId(e);
        },
        isLoadedSolutionId: function (e) {
          return this.solutionStore.isLoadedSolutionId(e);
        },
        isLoadingSolutionId: function (e) {
          return this.solutionStore.isLoadedSolutionId(e);
        },
        addConference: function (e) {
          this.conferenceStore.addConference(e);
        },
        removeConference: function (e) {
          (this.removeReactedConfereceId(e.id),
            this.conferenceStore.removeConference(e),
            this.messageStore.closeConference(e));
        },
        removeConferencesOnLeaveDomain: function (t) {
          var n = this;
          k.iter(
            this.conferenceStore.getConferences().filter(function (e) {
              e = e.domainId;
              return (
                null != e && null != t && e.high == t.high && e.low == t.low
              );
            }),
            function (e) {
              (n.removeConference(e),
                n.sendNotification("notify_close_conference", e));
            },
          );
        },
        removeConferencesOnLeaveTalk: function (t) {
          var n = this;
          k.iter(
            this.conferenceStore.getConferences().filter(function (e) {
              e = e.talkId;
              return (
                null != e && null != t && e.high == t.high && e.low == t.low
              );
            }),
            function (e) {
              (n.removeConference(e),
                n.sendNotification("notify_close_conference", e));
            },
          );
        },
        getConference: function (e) {
          return this.conferenceStore.getConference(e);
        },
        hasConferenceInTalk: function (e) {
          return this.conferenceStore.hasConferenceInTalk(e);
        },
        getConferenceIdInTalk: function (e) {
          return this.conferenceStore.getConferenceIdInTalk(e);
        },
        hasConferenceInDomain: function (t) {
          return (
            null !=
            k.find(
              this.conferenceStore
                .getConferenceTalkIds()
                .map(z(this, this.getTalk)),
              function (e) {
                return (
                  null != e &&
                  null != (e = e.domainId) &&
                  null != t &&
                  e.high == t.high &&
                  e.low == t.low
                );
              },
            )
          );
        },
        hasConferenceInAllDomains: function (t) {
          return null == t
            ? this.conferenceStore.hasConference()
            : null !=
                k.find(
                  this.conferenceStore
                    .getConferenceTalkIds()
                    .map(z(this, this.getTalk)),
                  function (e) {
                    return (
                      null != e &&
                      !(
                        null != (e = e.domainId) &&
                        null != t &&
                        e.high == t.high &&
                        e.low == t.low
                      )
                    );
                  },
                );
        },
        getFileInfo: function (e) {
          return this.fileInfoStore.getFileInfo(e);
        },
        getTalkFileInfo: function (e, t, n) {
          return this.fileInfoStore.getTalkFileInfo(e, t, n);
        },
        setTalkFileInfo: function (e, t) {
          this.fileInfoStore.setTalkFileInfo(e, t);
        },
        getFilePreview: function (e) {
          return h.orNull(this.fileInfoStore.getFilePreview(e));
        },
        setFilePreview: function (e, t) {
          this.fileInfoStore.setFilePreview(e, t);
        },
        getMessageFileInfos: function (e) {
          return this.fileInfoStore.getMessageFileInfos(e);
        },
        getScheduledMessageFileInfos: function (e) {
          return this.fileInfoStore.getScheduledMessageFileInfos(e);
        },
        isAllFileDeletedOnMessage: function (e) {
          return this.fileInfoStore.isAllFileDeletedOnMessage(e);
        },
        addStagedFileInfo: function (e, t, n, i) {
          return this.fileInfoStore.addStagedFileInfo(e, t, n, i);
        },
        replaceStagedFileInfo: function (e, t, n, i) {
          return this.fileInfoStore.replaceStagedFileInfo(e, t, n, i);
        },
        moveStagedFile: function (e, t, n) {
          this.fileInfoStore.moveStagedFile(e, t, n);
        },
        ensureTalkFileInfos: function (e) {
          return this.fileInfoStore.ensureTalkFileInfos(e);
        },
        ensureStagedFileInfos: function (e) {
          return this.fileInfoStore.ensureStagedFileInfos(e);
        },
        getStagedFilesCount: function (e) {
          e = this.fileInfoStore.ensureStagedFileIdStrs(e);
          return null == e ? 0 : e.length;
        },
        removeStagedFileInfos: function (e) {
          this.fileInfoStore.removeStagedFileInfos(e);
        },
        removeStagedFileInfo: function (e, t) {
          this.fileInfoStore.removeStagedFileInfo(e, t);
        },
        restoreStagedFile: function (e, t) {
          this.fileInfoStore.restoreStagedFile(e, t);
        },
        replaceStagedFileInfoFromFileId: function (e, t, n) {
          this.fileInfoStore.replaceStagedFileInfoFromFileId(e, t, n);
        },
        addStampsets: function (e) {
          this.stampStore.addStampsets(e);
        },
        updateStampsetInfo: function (t) {
          (k.iter(this.getDomains(), function (e) {
            e.updateStampsetInfo(t);
          }),
            this.stampStore.updateStampsetInfo(t));
        },
        removeStampsetInfo: function (t) {
          (k.iter(this.getDomains(), function (e) {
            e.deleteStampsetInfo(t);
          }),
            this.stampStore.removeStampsetInfo(t));
        },
        addStampsetInfos: function (e) {
          this.stampStore.addStampsetInfos(e);
        },
        getStampsetByStampsetId: function (e) {
          return this.stampStore.getStampsetByStampsetId(e);
        },
        getStampsetsByTabIds: function (e) {
          return this.stampStore.getStampsetsByTabIds(e);
        },
        getStampByStampKey: function (e) {
          return this.stampStore.getStampByStampKey(e);
        },
        getOriginalStamp: function (e) {
          return this.stampStore.getOriginalStamp(e);
        },
        applyHistories: function (e, t) {
          this.stampStore.applyHistories(e, t);
        },
        getDepartment: function (e) {
          return this.departmentStore.getDepartment(e);
        },
        setDepartment: function (e) {
          this.departmentStore.setDepartment(e);
        },
        clearDomainDepartment: function (e) {
          this.departmentStore.clearDomainDepartment(e);
        },
        getRootDepartmentId: function (e) {
          return this.departmentStore.getRootDepartmentId(e);
        },
        setRootDepartmentId: function (e, t) {
          this.departmentStore.setRootDepartmentId(e, t);
        },
        setDepartmentUserCount: function (e) {
          this.departmentStore.setDepartmentUserCount(e);
        },
        setUserCountLoading: function (e, t) {
          this.departmentStore.setUserCountLoading(e, t);
        },
        isUserCountLoadingRequired: function (e) {
          return this.departmentStore.isUserCountLoadingRequired(e);
        },
        clearDomainDepartmentUsers: function (e) {
          return this.departmentStore.clearDomainDepartmentUsers(e);
        },
        setDepartmentUsers: function (e, t) {
          this.departmentStore.setDepartmentUsers(e, t);
        },
        isInDepartment: function (e) {
          return this.departmentStore.isInDepartment(e);
        },
        getDepartmentPath: function (e) {
          return this.departmentStore.getDepartmentPath(e);
        },
        getDepartmentPathWithoutRoot: function (e) {
          return this.departmentStore.getDepartmentPathWithoutRoot(e);
        },
        getDepartmentPathSlashSeparatedIfNeeded: function (e) {
          return this.departmentStore.getDepartmentPathSlashSeparatedIfNeeded(
            e,
          );
        },
        clear: function (e) {
          ((this.users = new B()),
            (this.talks = new B()),
            (this.talkStatuses = new B()),
            (this.domains = new B()),
            (this.domainInvites = new B()),
            (this.announcementStatuses = new B()),
            (this.accountControlRequests = new B()),
            (this.accountControlGroup = null),
            (this.reactedConferenceIds = new B()),
            e && (this.me = null),
            this.storage.setDirtyFlag(),
            this.storageForTalks.setDirtyFlag(),
            this.storageForTalkStatuses.setDirtyFlag(),
            this.storageForUsers.setDirtyFlag());
        },
        init: function () {
          ((this.storage = new Ro("dataStore")),
            (this.storageForTalks = new Ro("dataStoreTalks")),
            (this.storageForTalkStatuses = new Ro("dataStoreTalkStatuses")),
            (this.storageForUsers = new Ro("dataStoreUsers")));
          var e = !1,
            t = this.storage.load(),
            t =
              (null != t &&
                "1.128_0" == t.shift() &&
                ((this.domains = t.shift()),
                (this.domainInvites = t.shift()),
                (this.announcementStatuses = t.shift()),
                null != this.announcementStatuses &&
                  k.iter(this.announcementStatuses, function (e) {
                    e.clearReadingAnnouncements();
                  }),
                (this.accountControlRequests = t.shift()),
                (this.accountControlGroup = t.shift()),
                (this.reactedConferenceIds = t.shift()),
                (e = !0)),
              !1),
            n = this.storageForTalks.load(),
            n =
              (null != n &&
                "1.128_0" == n.shift() &&
                ((this.talks = n.shift()), (t = !0)),
              !1),
            i = this.storageForTalkStatuses.load(),
            i =
              (null != i &&
                "1.128_0" == i.shift() &&
                ((this.talkStatuses = i.shift()),
                null != this.talkStatuses &&
                  k.iter(this.talkStatuses, function (e) {
                    e.clearReadingMessages();
                  }),
                (n =
                  null != this.talkStatuses &&
                  null != this.talks &&
                  k.count(this.talkStatuses) <= k.count(this.talks))),
              !1),
            r = this.storageForUsers.load();
          (null != r &&
            "1.128_0" == r.shift() &&
            ((this.me = r.shift()), (this.users = r.shift()), (i = !0)),
            e && t && n && i
              ? F.d("dataStore is loaded.")
              : (this.clear(!0), F.d("dataStore is initialized.")));
        },
        saveIfNeeded: function () {
          (this.storage.isDirty &&
            (this.storage.save([
              "1.128_0",
              this.domains,
              this.domainInvites,
              this.announcementStatuses,
              this.accountControlRequests,
              this.accountControlGroup,
              this.reactedConferenceIds,
            ]),
            F.d("dataStore is saved.")),
            this.storageForTalks.isDirty &&
              (this.storageForTalks.save(["1.128_0", this.talks]),
              F.d("dataStoreTalks is saved.")),
            this.storageForTalkStatuses.isDirty &&
              (this.storageForTalkStatuses.save(["1.128_0", this.talkStatuses]),
              F.d("dataStoreTalkStatuses is saved.")),
            this.storageForUsers.isDirty &&
              (this.storageForUsers.save(["1.128_0", this.me, this.users]),
              F.d("dataStoreUsers is saved.")));
        },
        __class__: qo,
      })),
      function () {
        l.call(this, "features");
      }),
    Xo =
      (((t["albero.proxy.FeaturesProxy"] = Qo).__name__ = [
        "albero",
        "proxy",
        "FeaturesProxy",
      ]),
      (Qo.__super__ = l),
      (Qo.prototype = e(l.prototype, {
        canSendMessageOnMultiView: function (e) {
          return (
            !!this.settings.isOnPremises() ||
            this.selectedDomainHasSolution(e, jr.MultiView)
          );
        },
        canDetectKeyword: function (e) {
          return (
            !!this.settings.isOnPremises() ||
            this.selectedDomainHasSolution(e, jr.KeywordWatching)
          );
        },
        selectedDomainHasSolution: function (e, t) {
          return (
            null != e &&
            ((e =
              null == (e = this.dataStore.getDomain(e))
                ? []
                : e.getSolutionIds()),
            k.exists(e, function (e) {
              return ze.getSolutionType(e) == t;
            }))
          );
        },
        canConference: function (e) {
          return (
            null != e &&
            null != (e = this.dataStore.getDomain(e)) &&
            null != e.contract &&
            e.contract.isConferenceEnalbed()
          );
        },
        __class__: Qo,
      })),
      function () {}),
    Jo =
      (((t["albero.proxy.FileServiceProxy"] = Xo).__name__ = [
        "albero",
        "proxy",
        "FileServiceProxy",
      ]),
      (Xo.__interfaces__ = [Xt]),
      (Xo.prototype = { __class__: Xo }),
      ((t["albero.proxy.UploadAuth"] = $t).__name__ = [
        "albero",
        "proxy",
        "UploadAuth",
      ]),
      ($t.prototype = { __class__: $t }),
      function () {}),
    Zo =
      (((t["albero.proxy.FileServiceProxyFactory"] = Jo).__name__ = [
        "albero",
        "proxy",
        "FileServiceProxyFactory",
      ]),
      (Jo.newInstance = function () {
        return new Zo("fileService");
      }),
      function (e) {
        l.call(this, e);
      }),
    $o =
      (((t["albero.proxy.FileServiceProxyImplForHubot"] = Zo).__name__ = [
        "albero",
        "proxy",
        "FileServiceProxyImplForHubot",
      ]),
      (Zo.__interfaces__ = [Xo]),
      (Zo.getContentTypeByFileName = function (e, t) {
        return new M().resolve({ name: e, type: t });
      }),
      (Zo.__super__ = l),
      (Zo.prototype = e(l.prototype, {
        downloadUrl: function (e, t) {
          var n, i;
          return (
            null == t && (t = !1),
            null == e || 0 == e.length
              ? "data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=="
              : I.startsWith(e, this.settings.getValidHost())
                ? ((n = this.settings.getAccessToken()),
                  -1 != (i = e).indexOf("?") ? (i += "&") : (i += "?"),
                  (i += "Authorization=ALB%20" + n),
                  t && (i += "&download=true"),
                  i)
                : e
          );
        },
        download: function (e, r, o) {
          var a = this;
          ((e = this.downloadUrl(e)),
            I.startsWith(e, "data:")
              ? o(null, new Error("download URL was not resolved"))
              : ((e = ps.parse(e)),
                null != qs.proxyURL && (e.agent = Fo.createAgent(qs.proxyURL)),
                (e = ds.request(e, function (e) {
                  var t,
                    n,
                    i = e.headers.location;
                  null != i
                    ? (F.d("redirect to " + i), a.download(i, r, o))
                    : 2 != Math.floor(e.statusCode / 100)
                      ? (F.d("failed to download. status:" + e.statusCode),
                        (t = ""),
                        e.setEncoding("utf8"),
                        e.on("data", function (e) {
                          t += e;
                        }),
                        e.on("end", function () {
                          new $o(t).shouldRetryDownload()
                            ? o(
                                null,
                                new Error(
                                  "The download URL is still being prepared. Please try again later.",
                                ),
                              )
                            : o(
                                null,
                                new Error(
                                  "HTTP request failed (" +
                                    e.statusCode +
                                    " " +
                                    e.statusMessage +
                                    ")",
                                ),
                              );
                        }),
                        e.on("error", function (e) {
                          o(null, e);
                        }))
                      : ((n = hs.createWriteStream(r)),
                        e.on("data", function (e) {
                          n.write(e);
                        }),
                        e.on("end", function () {
                          (n.end(),
                            n.on("finish", function () {
                              o(r, null);
                            }));
                        }),
                        e.on("error", function (e) {
                          o(null, e);
                        }),
                        n.on("error", function (e) {
                          o(null, e);
                        }));
                })).on("error", function (e) {
                  (F.d("failed to download. request error:" + e.message),
                    o(null, e));
                }),
                e.end()));
        },
        uploadForNodeJS: function (e, t, r) {
          var n = r,
            o = ps.parse(e.put_url),
            n =
              ((o.method = "PUT"),
              (o.headers = {}),
              (o.headers["Content-Length"] = "" + n.size),
              e.post_form);
          return (
            (o.headers["Content-Type"] = n["Content-Type"]),
            (o.headers["Content-Disposition"] = n["Content-Disposition"]),
            null != qs.proxyURL && (o.agent = Fo.createAgent(qs.proxyURL)),
            new Promise(function (n, i) {
              var t = ds.request(o, function (e) {
                  var t;
                  2 != Math.floor(e.statusCode / 100)
                    ? ((t = ""),
                      e.on("data", function (e) {
                        t += e;
                      }),
                      e.on("end", function () {
                        i(e.statusCode + ": " + t);
                      }))
                    : Kt.resolveVoid(n);
                }),
                e =
                  (t.on("error", function (e) {
                    i(e.message);
                  }),
                  hs.createReadStream(r.path));
              (e.on("data", function (e) {
                t.write(e);
              }),
                e.on("end", function () {
                  t.end();
                }),
                e.on("error", function () {
                  t.end();
                }));
            })
          );
        },
        upload: function (e, t, n) {
          return this.uploadForNodeJS(e, t, n);
        },
        createDummyFile: function (e, t, n) {
          var i = I.replace(e, "/", "\\").split("\\");
          return 0 == i.length ||
            0 == (t = null != t ? t : i[i.length - 1]).length
            ? null
            : ((i = Zo.getContentTypeByFileName(t, n)),
              { name: t, size: 0 | hs.statSync(e).size, type: i, path: e });
        },
        __class__: Zo,
      })),
      function (t) {
        try {
          var e = JSON.parse(t);
          ((this.code = e.code),
            (this.message = e.message),
            (this.detail = e.detail));
        } catch (e) {
          F.d("Failed to parse download error: " + t);
        }
      }),
    ea =
      (((t["albero.proxy.DownloadErrorResponseBody"] = $o).__name__ = [
        "albero",
        "proxy",
        "DownloadErrorResponseBody",
      ]),
      ($o.prototype = {
        shouldRetryDownload: function () {
          var e;
          return (
            !!(
              "PendingFile" == this.code &&
              null != this.detail &&
              "number" == typeof (e = this.detail.reason) &&
              (0 | e) === e
            ) && this.detail.reason < 1e3
          );
        },
        __class__: $o,
      }),
      function () {
        l.call(this, "keywordWatcher");
      }),
    ta =
      (((t["albero.proxy.KeywordWatcherProxy"] = ea).__name__ = [
        "albero",
        "proxy",
        "KeywordWatcherProxy",
      ]),
      (ea.__super__ = l),
      (ea.prototype = e(l.prototype, {
        onMessageCreated: function (e) {
          if (null != e) {
            var t,
              n,
              i = h.map(
                h.option(this.dataStore.getTalk(e.talkId)),
                function (e) {
                  return e.domainId;
                },
              );
            switch (i[1]) {
              case 0:
                var r = i[2];
                break;
              case 1:
                r = this.settings.getSelectedDomainId();
            }
            this.features.canDetectKeyword(r) &&
              !(
                !this.settings.getBrowserSettings()
                  .keywordWatchingSelfMessage &&
                ((n = this.dataStore), (t = e.userId), null != n.me) &&
                null != (n = n.me.id) &&
                null != t &&
                n.high == t.high &&
                n.low == t.low
              ) &&
              ((n = e.messageStringForKeywordDetection(
                this.settings.getBrowserSettings().keywordWatchingActionReply,
              )),
              this.detectKeyword(n, e.type, e.id, r, e.talkId));
          }
        },
        onAnnouncementCreated: function (e) {
          var t;
          null != e &&
            this.features.canDetectKeyword(e.domainId) &&
            ((t = O.getDisplayTextWithoutEscape(e.type, e.content)),
            this.detectKeyword(t, e.type, e.id, e.domainId, null));
        },
        detectKeyword: function (s, l, u, c, _) {
          var h = this;
          this.reserveOnNormalizedKeywordPrepared(function (e) {
            for (var t = I.htmlEscape(s, !0), n = e.keys(); n.hasNext(); ) {
              var i = n.next(),
                r = e.get(i);
              if (null != r && 0 != r.length && ra.match(l, i))
                for (var o = 0; o < r.length; ) {
                  var a = r[o];
                  if ((++o, 0 != a.search(t).length))
                    return void h.addkeywordDetection(c, _, u);
                }
            }
          });
        },
        reserveOnNormalizedKeywordPrepared: function (e) {
          var t = this.settings.getBrowserSettings().keywordWatchingText;
          if (this.cachedKeyword != t) {
            ((this.cachedKeyword = t),
              (this.cachedNormalizedKeywordsMap = new es()));
            for (var n = 0, i = ia.__empty_constructs__; n < i.length; ) {
              var r = i[n];
              (++n, this.cachedNormalizedKeywordsMap.set(r, []));
            }
            for (var o = 0, a = ta.parse(t); o < a.length; ) {
              var s = a[o];
              (++o,
                this.cachedNormalizedKeywordsMap
                  .get(s.type)
                  .push(new Wt(s.word, !0, xo.normalize)));
            }
          }
          e(this.cachedNormalizedKeywordsMap);
        },
        addkeywordDetection: function (e, t, n) {
          null == this.detectedDomainTalkIdStrings &&
            (this.detectedDomainTalkIdStrings = new B());
          var i,
            r = this.getKey(e, t);
          null != r &&
            (this.hasKeywordDetection(e, t)
              ? ((i = this.detectedDomainTalkIdStrings),
                null != Y[r] ? i.setReserved(r, n) : (i.h[r] = n))
              : ((i = this.detectedDomainTalkIdStrings),
                null != Y[r] ? i.setReserved(r, n) : (i.h[r] = n),
                (i =
                  this.settings.getBrowserSettings().keywordWatchingEmphasis),
                this.sendNotification("keyword_detaction_updated", {
                  domainId: e,
                  talkId: t,
                  detected: !0,
                  emphasis: i,
                })));
        },
        readKeywordDetection: function (e, t, n) {
          var i, r;
          !this.settings.getBrowserSettings().keywordWatchingEmphasis &&
            ((i = this.getDetectId(e, t)),
            F.v("readKeywordDetection", i, n),
            null != i &&
              (0 == (r = (n.high - i.high) | 0) &&
                (r = w.ucompare(n.low, i.low)),
              0 <=
                (n.high < 0 ? (i.high < 0 ? r : -1) : 0 <= i.high ? r : 1))) &&
            this.removeKeywordDetection(e, t);
        },
        removeKeywordDetection: function (e, t) {
          var n;
          null != this.detectedDomainTalkIdStrings &&
            null != (n = this.getKey(e, t)) &&
            (this.detectedDomainTalkIdStrings.remove(n),
            (n = this.settings.getBrowserSettings().keywordWatchingEmphasis),
            this.sendNotification("keyword_detaction_updated", {
              domainId: e,
              talkId: t,
              detected: !1,
              emphasis: n,
            }));
        },
        removeKeywordDetectionAll: function () {
          if (null != this.detectedDomainTalkIdStrings)
            for (
              var e =
                  this.settings.getBrowserSettings().keywordWatchingEmphasis,
                t = this.detectedDomainTalkIdStrings.keys();
              t.hasNext();
            ) {
              var n,
                i,
                r = t.next();
              (I.startsWith(r, "talk_")
                ? ((n = r.substring("talk_".length)),
                  (n = D.makeFromIdStr(n)),
                  this.sendNotification("keyword_detaction_updated", {
                    domainId: null,
                    talkId: n,
                    detected: !1,
                    emphasis: e,
                  }))
                : I.startsWith(r, "announcement_") &&
                  ((n = r.substring("announcement_".length)),
                  (i = D.makeFromIdStr(n)),
                  this.sendNotification("keyword_detaction_updated", {
                    domainId: i,
                    talkId: null,
                    detected: !1,
                    emphasis: e,
                  })),
                this.detectedDomainTalkIdStrings.remove(r));
            }
        },
        getDetectId: function (e, t) {
          return null == this.detectedDomainTalkIdStrings ||
            null == (e = this.getKey(e, t))
            ? null
            : ((t = this.detectedDomainTalkIdStrings),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        hasKeywordDetection: function (e, t) {
          return null != this.getDetectId(e, t);
        },
        getKey: function (e, t) {
          return null == e && null == t
            ? null
            : null != t
              ? "talk__" + t.high + "_" + t.low
              : "announcement__" + e.high + "_" + e.low;
        },
        __class__: ea,
      })),
      function () {}),
    na =
      (((t["albero.proxy.KeywordParser"] = ta).__name__ = [
        "albero",
        "proxy",
        "KeywordParser",
      ]),
      (ta.parse = function (e) {
        ((e = I.htmlEscape(e, !0)),
          (e = new v("&quot;(.*?)&quot;", "g").map(e, function (e) {
            return ta.encodeSpace(e.matched(1));
          })));
        return new v("[ 　]", "g")
          .split(e)
          .filter(Ln.isNotEmpty)
          .map(ta.decodeSpace)
          .map(ta.createParsedKeyword);
      }),
      (ta.encodeSpace = function (e) {
        return I.replace(
          I.replace(e, " ", "&SPACE"),
          "　",
          "&FULL_PITCH_SPACE",
        );
      }),
      (ta.decodeSpace = function (e) {
        return I.replace(
          I.replace(e, "&SPACE", " "),
          "&FULL_PITCH_SPACE",
          "　",
        );
      }),
      (ta.createParsedKeyword = function (e) {
        var t = new v("\\[type:([A-Za-z]+)\\](.*)", "g");
        if (t.match(e)) {
          var n = t.matched(1),
            t = t.matched(2),
            n = ra.getKeywordTypeByMessageTypeString(n);
          if (n != ia.normal) return new na(n, t);
        }
        return new na(ia.normal, e);
      }),
      function (e, t) {
        ((this.type = e), (this.word = t));
      }),
    ia =
      (((t["albero.proxy.ParsedKeyword"] = na).__name__ = [
        "albero",
        "proxy",
        "ParsedKeyword",
      ]),
      (na.prototype = { __class__: na }),
      (t["albero.proxy.KeywordType"] = {
        __ename__: ["albero", "proxy", "KeywordType"],
        __constructs__: [
          "normal",
          "selectOneReplyOnly",
          "yesOrNoReplyOnly",
          "todoDoneOnly",
        ],
      })),
    ra =
      ((ia.normal = ["normal", 0]),
      (ia.normal.toString = a),
      ((ia.normal.__enum__ = ia).selectOneReplyOnly = [
        "selectOneReplyOnly",
        1,
      ]),
      (ia.selectOneReplyOnly.toString = a),
      ((ia.selectOneReplyOnly.__enum__ = ia).yesOrNoReplyOnly = [
        "yesOrNoReplyOnly",
        2,
      ]),
      (ia.yesOrNoReplyOnly.toString = a),
      ((ia.yesOrNoReplyOnly.__enum__ = ia).todoDoneOnly = ["todoDoneOnly", 3]),
      (ia.todoDoneOnly.toString = a),
      ((ia.todoDoneOnly.__enum__ = ia).__empty_constructs__ = [
        ia.normal,
        ia.selectOneReplyOnly,
        ia.yesOrNoReplyOnly,
        ia.todoDoneOnly,
      ]),
      function () {}),
    oa =
      (((t["albero.proxy.KeywordTypeHelper"] = ra).__name__ = [
        "albero",
        "proxy",
        "KeywordTypeHelper",
      ]),
      (ra.getKeywordType = function (e) {
        return null != e
          ? ra.getKeywordTypeByMessageTypeString(O.typeString(e))
          : ia.normal;
      }),
      (ra.getKeywordTypeByMessageTypeString = function (e) {
        if (null != e)
          switch (e) {
            case "selectOneReply":
              return ia.selectOneReplyOnly;
            case "todoDone":
              return ia.todoDoneOnly;
            case "yesOrNoReply":
              return ia.yesOrNoReplyOnly;
          }
        return ia.normal;
      }),
      (ra.match = function (e, t) {
        switch (t[1]) {
          case 0:
            return !0;
          case 1:
            return e == d.selectOneReply;
          case 2:
            return e == d.yesOrNoReply;
          case 3:
            return e == d.todoDone;
        }
      }),
      (ra.createFilterClassStringFromMessageType = function (e) {
        return ra.createFilterClassString(ra.getKeywordType(e));
      }),
      (ra.createFilterClassString = function (e) {
        return "keyword_type_" + P.string(e);
      }),
      function () {
        ((this.isCallingUpdateLastUsedAt = !1),
          l.call(this, "lastUsedAtUpdater"));
      }),
    aa =
      (((t["albero.proxy.LastUsedAtUpdaterProxy"] = oa).__name__ = [
        "albero",
        "proxy",
        "LastUsedAtUpdaterProxy",
      ]),
      (oa.__super__ = l),
      (oa.prototype = e(l.prototype, {
        startPolling: function (e) {
          var t,
            n = this;
          null == this.pollingTimer &&
            (0 <
            (t =
              null == (t = this.getLastUsedExpiredAt())
                ? 0
                : 4294967296 * t.high +
                  (0 <= (t = t.low) ? t : t + 4294967296) -
                  new Date().getTime())
              ? Xa.delay(function () {
                  n.startPolling(e);
                }, 0 | t)
              : ((this.isCallingUpdateLastUsedAt = !0),
                e(),
                (this.pollingTimer = new Xa(
                  this.settings.getConfiguration().updateLastUsedAtInterval,
                )),
                (this.pollingTimer.run = function () {
                  n.isCallingUpdateLastUsedAt ||
                    ((n.isCallingUpdateLastUsedAt = !0), e());
                })));
        },
        stopPolling: function () {
          null != this.pollingTimer &&
            (this.pollingTimer.stop(), (this.pollingTimer = null));
        },
        getLastUsedExpiredAt: function () {
          return null != this.lastUsedExpiredAtCache
            ? this.lastUsedExpiredAtCache
            : this.settings.getLastUsedExpiredAt();
        },
        update: function () {
          ((this.lastUsedExpiredAtCache = W.afterAsInt64(
            this.settings.getConfiguration().updateLastUsedAtInterval,
          )),
            this.settings.setLastUsedExpiredAt(this.lastUsedExpiredAtCache),
            (this.isCallingUpdateLastUsedAt = !1));
        },
        cancel: function () {
          this.isCallingUpdateLastUsedAt = !1;
        },
        __class__: oa,
      })),
      function () {
        l.call(this, "limitations");
      }),
    sa =
      (((t["albero.proxy.LimitationsProxy"] = aa).__name__ = [
        "albero",
        "proxy",
        "LimitationsProxy",
      ]),
      (aa.__super__ = l),
      (aa.prototype = e(l.prototype, {
        getMaxTalkers: function () {
          var t = this,
            e = h.flatMap(
              h.flatMap(
                h.option(this.settings.getSelectedDomainId()),
                function (e) {
                  return h.option(t.dataStore.getDomain(e));
                },
              ),
              function (e) {
                return e.getMaxTalkers();
              },
            );
          switch (e[1]) {
            case 0:
              return e[2];
            case 1:
              return this.settings.getConfiguration().maxTalkers;
          }
        },
        __class__: aa,
      })),
      function () {
        ((this.connectionStatus = Rn.Ok),
          l.call(this, "rpc"),
          (this.responseHandlers = new ts()),
          (this.connectionKeeper = new la(z(this, this.ping))));
      }),
    la =
      (((t["albero.proxy.MsgPackRpcProxy"] = sa).__name__ = [
        "albero",
        "proxy",
        "MsgPackRpcProxy",
      ]),
      (sa.__super__ = l),
      (sa.prototype = e(l.prototype, {
        initWebSocket: function () {
          null == this.ws &&
            ((this.ws = Qt.newInstance(qs.endpoint, {
              onopen: z(this, this.onOpen),
              onmessage: z(this, this.onMessage),
              onclose: z(this, this.onClose),
              onpong: z(this, this.onPong),
            })),
            (sa.pingAt = null));
        },
        finishWebSocket: function () {
          if (null == this.ws) return null;
          this.ws.close();
          var e = this.ws;
          return (
            (this.ws = null),
            (sa.pingAt = null),
            this.lastUsedAtUpdater.stopPolling(),
            e
          );
        },
        onRegister: function () {
          (this.initWebSocket(), this.connectionKeeper.start());
        },
        onRemove: function () {
          (this.connectionKeeper.stop(), this.finishWebSocket());
        },
        onOpen: function () {
          ((this.connectionStatus = Rn.Ok),
            this.connectionKeeper.setConnected(!0),
            this.sendNotification("SignIn"));
        },
        onPong: function (e) {
          ((sa.pingAt = null), F.v("PONG " + P.string(e)));
        },
        onMessage: function (e) {
          var t,
            n,
            i,
            r,
            e = new ys(e, !0, new ha()).o,
            o = G.__cast(e[0], js);
          1 == o && 4 == e.length
            ? ((t = G.__cast(e[1], js)),
              (r = e[2]),
              (n = e[3]),
              null == (i = this.responseHandlers.h[t])
                ? F.e(
                    "No ResponseHandler prepared. msgId:%s error:%s result:",
                    t,
                    r,
                    n,
                  )
                : (F.v("response received. method:", i.method, " data:", e),
                  null == r
                    ? i.onSuccess(n)
                    : (F.e(
                        "Receive Error Response. method:",
                        i.method,
                        " error:",
                        r,
                      ),
                      i.onError(r)),
                  this.responseHandlers.remove(t)))
            : 0 == o &&
              4 == e.length &&
              ((n = G.__cast(e[1], js)),
              (i = G.__cast(e[2], String)),
              (r = G.__cast(e[3], Array)),
              F.v("request received. method:", i),
              0 == r.length
                ? F.e("empty params.")
                : (this.broadcast.handleNotification(i, r[0]),
                  null == this.ws
                    ? F.e("websocket was finished.")
                    : (F.v("response sent. method:", i),
                      this.ws.send(new Ss([1, n, null, !0]).o.getBytes()))));
        },
        onClose: function (e, t, n) {
          (F.i(
            "onClose. code:" +
              e +
              ", reason:" +
              t +
              ", wasClean:" +
              (null == n ? "null" : "" + n),
          ),
            (1001 == e && n) ||
              ((this.connectionStatus =
                (1e3 != e && 1005 != e) || "concurrent access" != t
                  ? (1e3 != e && 1005 != e) || "forcibly closed" != t
                    ? Rn.Error
                    : Rn.ForcibliyClosedError
                  : Rn.ConcurrentAccessError),
              this.connectionStatus == Rn.ForcibliyClosedError
                ? this.sendNotification("SignOut")
                : this.sendNotification("Url", bi.FORWARD(g.error))),
            this.finishWebSocket(),
            this.connectionKeeper.setConnected(!1));
        },
        restart: function (e) {
          null == e && (e = 500);
          for (
            var t = this,
              n = this.finishWebSocket(),
              i = this.responseHandlers.keys();
            i.hasNext();
          ) {
            var r = i.next();
            this.responseHandlers.remove(r);
          }
          var o = null,
            o = function () {
              null == n || n.isClosed() ? t.initWebSocket() : Xa.delay(o, 100);
            };
          Xa.delay(o, e);
        },
        call: function (t, e, n, i) {
          var r,
            o,
            a,
            s,
            l = this;
          null == this.ws
            ? F.e("disconnected. data:", this.data)
            : (null == e && (e = []),
              null == n && (n = function (e) {}),
              null == i && (i = z(this, this.onServerError)),
              (r = sa.lastMsgId++),
              (o = new ua(t, n, i)),
              (this.responseHandlers.h[r] = o),
              (s = new Ss((a = [0, r, t, e])).o.getBytes()),
              this.sendAsPromise(s)
                .catch(function (e) {
                  return (
                    F.v("try to send again 1 second later. data:", a),
                    Kt.delay(1e3).then(function (e) {
                      return l.sendAsPromise(s);
                    })
                  );
                })
                .then(
                  function (e) {
                    F.v("send request. data:", a);
                  },
                  function (e) {
                    (o.onError(
                      new ca({
                        message: "failed to execute send on websocket",
                        detail: t,
                      }),
                    ),
                      l.responseHandlers.remove(r));
                  },
                ));
        },
        sendAsPromise: function (e) {
          try {
            return (this.ws.send(e), Promise.resolve());
          } catch (e) {
            return (e instanceof H && (e = e.val), Promise.reject(e));
          }
        },
        ping: function () {
          switch (this.connectionStatus[1]) {
            case 2:
            case 3:
              return;
          }
          null == this.ws || this.ws.isClosed() || 0 < sa.pingAt
            ? (this.restart(),
              F.v(
                "restart connection. " +
                  sa.pingAt +
                  " " +
                  P.string(null == this.ws || this.ws.isClosed()),
              ))
            : (F.v("before pingAt " + sa.pingAt),
              this.connectionKeeper.isConnected() &&
                (this.ws.ping("PING"),
                (sa.pingAt = new Date().getTime()),
                F.v("send ping")));
        },
        onServerError: function (e) {
          var t;
          (this.sendNotification("error_occurred", e),
            503 == e.code &&
              "service temporarily unavailable" == e.message &&
              ((t = e.detail.retry_after), this.restart(1e3 * t)),
            "invalid session" == e.message && this.restart());
        },
        __class__: sa,
      })),
      function (e) {
        ((this.connected = !1), (this.ping = e));
      }),
    ua =
      (((t["albero.proxy._MsgPackRpcProxy.ConnectionKeeper"] = la).__name__ = [
        "albero",
        "proxy",
        "_MsgPackRpcProxy",
        "ConnectionKeeper",
      ]),
      (la.prototype = {
        start: function () {},
        stop: function () {
          this.deleteTimer();
        },
        setConnected: function (e) {
          (this.connected == e && null != this.timer) ||
            ((this.connected = e), this.resetTimer());
        },
        isConnected: function () {
          return this.connected;
        },
        resetTimer: function () {
          var i,
            r = this;
          (this.deleteTimer(),
            this.connected
              ? ((this.timer = new Xa(45e3)), (this.timer.run = this.ping))
              : ((i = null),
                (i = function (t) {
                  var n = 0 | Math.min(2 * t, 45e3);
                  return function () {
                    r.ping();
                    var e = i(n);
                    r.timer = Xa.delay(e, t);
                  };
                })(3e3)()));
        },
        deleteTimer: function () {
          null != this.timer && (this.timer.stop(), (this.timer = null));
        },
        __class__: la,
      }),
      function (e, t, n) {
        ((this.method = e), (this.onSuccess = t), (this.onError = n));
      }),
    ca =
      (((t["albero.proxy._MsgPackRpcProxy.ResponseHandler"] = ua).__name__ = [
        "albero",
        "proxy",
        "_MsgPackRpcProxy",
        "ResponseHandler",
      ]),
      (ua.prototype = { __class__: ua }),
      function (e) {
        ((this.code = e.code),
          (this.message = e.message),
          (this.detail = e.detail),
          (this.localErrorType = e.localErrorType));
      }),
    _a =
      (((t["albero.proxy.Error"] = ca).__name__ = ["albero", "proxy", "Error"]),
      (ca.empty = function () {
        return new ca({});
      }),
      (ca.prototype = { __class__: ca }),
      function () {}),
    ha =
      (((t["albero.proxy.ErrorConverter"] = _a).__name__ = [
        "albero",
        "proxy",
        "ErrorConverter",
      ]),
      (_a.toLocalError = function (e, t) {
        return new ca({
          code: t.code,
          message: t.message,
          detail: t.detail,
          localErrorType: e,
        });
      }),
      ((t["msgpack.DecodeStrategy"] = en).__name__ = [
        "msgpack",
        "DecodeStrategy",
      ]),
      (en.prototype = { __class__: en }),
      function () {}),
    da =
      (((t["albero.proxy.AlberoDecodeStrategy"] = ha).__name__ = [
        "albero",
        "proxy",
        "AlberoDecodeStrategy",
      ]),
      (ha.__interfaces__ = [en]),
      (ha.prototype = {
        parseBytes: function (e) {
          return "rich_text_compressed" == this.decodingMapKey
            ? e
            : e.toString();
        },
        startDecodingMapValue: function (e) {
          this.decodingMapKey = e;
        },
        endDecodingMapValue: function () {
          this.decodingMapKey = null;
        },
        __class__: ha,
      }),
      function () {
        (l.call(this, "readStatusUpdater"),
          (this.updateReadStatusesTimers = new B()),
          (this.updateReadAnnouncementStatusesTimers = new B()));
      }),
    fa =
      (((t["albero.proxy.ReadStatusUpdaterProxy"] = da).__name__ = [
        "albero",
        "proxy",
        "ReadStatusUpdaterProxy",
      ]),
      (da.__super__ = l),
      (da.prototype = e(l.prototype, {
        updateReadStatuses: function (e, t) {
          var n,
            i,
            r = this,
            o = this.dataStore.getTalkStatus(e);
          null != o &&
            o.updateByReadingMessagesStarted(t, this.dataStore.me.id) &&
            (this.dataStore.setTalkStatus(o),
            (n = "_" + e.high + "_" + e.low),
            (t = this.updateReadStatusesTimers),
            null != (t = null != Y[n] ? t.getReserved(n) : t.h[n]) && t.stop(),
            (i = o.calcMaxReadMessageId()),
            (t = Xa.delay(function () {
              (r.updateReadStatusesTimers.remove(n),
                null != r.dataStore.getTalkStatus(e) &&
                  r.api.updateReadStatuses(e, i));
            }, 1e3)),
            (o = this.updateReadStatusesTimers),
            null != Y[n] ? o.setReserved(n, t) : (o.h[n] = t));
        },
        updateAnnouncementReadStatus: function (e, t) {
          var n,
            i,
            r = this,
            o = this.dataStore.getAnnouncementStatus(e);
          null != o &&
            o.updateByReadingAnnouncementsStarted(t) &&
            (this.dataStore.setAnnouncementStatus(o),
            (n = "_" + e.high + "_" + e.low),
            (t = this.updateReadAnnouncementStatusesTimers),
            null != (t = null != Y[n] ? t.getReserved(n) : t.h[n]) && t.stop(),
            (i = o.calcMaxReadAnnouncementId()),
            (t = Xa.delay(function () {
              (r.updateReadAnnouncementStatusesTimers.remove(n),
                r.api.updateAnnouncementReadStatus(e, i));
            }, 1e3)),
            (o = this.updateReadAnnouncementStatusesTimers),
            null != Y[n] ? o.setReserved(n, t) : (o.h[n] = t));
        },
        __class__: da,
      })),
      function () {
        l.call(this, "routing");
      }),
    ma =
      (((t["albero.proxy.RoutingProxy"] = fa).__name__ = [
        "albero",
        "proxy",
        "RoutingProxy",
      ]),
      (fa.__super__ = l),
      (fa.prototype = e(l.prototype, {
        onRegister: function () {},
        init: function () {
          null == this.router &&
            ((this.router = new ma(this, this.settings, this.dataStore)),
            this.sendNotification("first_routing_will_start"));
        },
        forward: function (e) {
          (this.init(), this.router.forward(e));
        },
        redirect: function (e, t) {
          (this.init(), this.router.redirect(e, t));
        },
        back: function () {
          null != this.router && this.router.back();
        },
        stop: function () {
          null != this.router && this.router.stop();
        },
        __class__: fa,
      })),
      function (e, t, n) {
        ((this.proxy = e),
          (this.settings = t),
          (this.dataStore = n),
          (this.urlConverterDelegate = new pa(t, n)));
      }),
    pa =
      (((t["albero.proxy._RoutingProxy.LocalRouter"] = ma).__name__ = [
        "albero",
        "proxy",
        "_RoutingProxy",
        "LocalRouter",
      ]),
      (ma.prototype = {
        notify: function (e) {
          var t = this.getDomainId(e);
          (this.settings.setSelectedDomainId(t),
            this.proxy.sendNotification("current_page_changed", e));
        },
        forward: function (e) {
          0 != e[1] && this.notify(e);
        },
        redirect: function (e, t) {},
        redirectWithHash: function () {},
        back: function () {
          this.notify(this.prev);
        },
        stop: function () {},
        getDomainId: function (e) {
          switch (e[1]) {
            case 2:
              return e[2];
            case 3:
              return e[2];
            case 4:
              return e[2];
            case 5:
              return e[2];
            case 6:
              return e[2];
            case 7:
              return e[2];
            case 8:
              return e[2];
            case 12:
              return e[2];
            case 13:
              return e[2];
            default:
              return null;
          }
        },
        parseFragment: function (e) {
          return te.createUrls(e, this.urlConverterDelegate);
        },
        toFragment: function (e) {
          return te.createFragment(e, this.urlConverterDelegate);
        },
        __class__: ma,
      }),
      function (e, t) {
        ((this.settings = e), (this.dataStore = t));
      }),
    ga =
      (((t["albero.proxy.UrlConverterDelegate"] = pa).__name__ = [
        "albero",
        "proxy",
        "UrlConverterDelegate",
      ]),
      (pa.__interfaces__ = [ne]),
      (pa.prototype = {
        isMultiViewMode: function () {
          return this.settings.isMultiViewMode();
        },
        getLastSelectedDomainId: function () {
          return h.option(this.settings.getLastSelectedDomainId());
        },
        getFallbackDomainId: function () {
          var e;
          switch (
            h.flatMap(
              this.getLastSelectedDomainId(),
              z(this, this.getDomain),
            )[1]
          ) {
            case 0:
              e = h.flatMap(
                this.getLastSelectedDomainId(),
                z(this, this.getDomain),
              );
              break;
            case 1:
              e = this.getOldestDomain();
          }
          return h.map(e, function (e) {
            return e.id;
          });
        },
        getDomain: function (e) {
          return h.option(this.dataStore.getDomain(e));
        },
        getOldestDomain: function () {
          return new Te(this.dataStore.getDomains()).getOldestDomain();
        },
        __class__: pa,
      }),
      function () {
        l.call(this, "searchService");
      }),
    ya =
      (((t["albero.proxy.SearchServiceProxy"] = ga).__name__ = [
        "albero",
        "proxy",
        "SearchServiceProxy",
      ]),
      (ga.__super__ = l),
      (ga.prototype = e(l.prototype, {
        isSearching: function () {
          return this.searching;
        },
        clearSearching: function () {
          this.searching = !1;
        },
        isRecentParams: function (e) {
          return (
            null != this.recentParams && null != e && this.recentParams.eq(e)
          );
        },
        clearRecentParams: function () {
          this.recentParams = null;
        },
        clearEditingParams: function () {
          this.editingParams = null;
        },
        saveParams: function (e) {
          ((this.editingParams = e), (this.recentParams = e.copy()));
        },
        updateTalkIdParam: function (e) {
          (this.assureParams(), (this.editingParams.talkId = e));
        },
        updateSearchTypeParam: function (e) {
          (this.assureParams(), (this.editingParams.searchType = e));
        },
        updateFromUserIdParam: function (e) {
          (this.assureParams(), (this.editingParams.fromUserId = e));
        },
        updateKeywordParam: function (e) {
          (this.assureParams(), (this.editingParams.keyword = e));
        },
        updateDomainIdParam: function (e) {
          (this.assureParams(), (this.editingParams.domainId = e));
        },
        updateSinceTextParam: function (e) {
          (this.assureParams(), (this.editingParams.sinceText = e));
        },
        updateUntilTextParam: function (e) {
          (this.assureParams(), (this.editingParams.untilText = e));
        },
        assureParams: function () {
          var e;
          null == this.editingParams &&
            ((e = {
              domain_id_str:
                "_" +
                (e = this.settings.getSelectedDomainId()).high +
                "_" +
                e.low,
            }),
            (this.editingParams = new He(e)));
        },
        prepareForSearching: function (e, t) {
          ((this.searching = !0),
            null == t &&
              this.sendNotification("notify_search_prepare", {
                searching: !0,
                params: e,
              }));
        },
        notifySearchMessagesFail: function () {
          ((this.searching = !1),
            this.sendNotification("notify_search_messages_fail", {}));
        },
        notifySearchAttachmentsFail: function () {
          ((this.searching = !1),
            this.sendNotification("notify_search_attachments_fail", {}));
        },
        saveSearchMessagesResult: function (e) {
          (null == this.contextMap && (this.contextMap = new B()),
            (this.messagesNextMarker = e.nextMarker),
            null == e.marker && (this.hitMessageIdSet = new B()));
          for (var t = 0, n = e.contents; t < n.length; ) {
            var i = n[t],
              r = (++t, i.messages[0].id),
              r = "_" + r.high + "_" + r.low,
              o = this.contextMap;
            (null != Y[r] ? o.setReserved(r, i) : (o.h[r] = i),
              (this.hitMessageIdSet = k.fold(
                i.messages,
                function (e, t) {
                  return (
                    e.hit &&
                      ((e = "_" + (e = e.id).high + "_" + e.low),
                      null != Y[e] ? t.setReserved(e, 1) : (t.h[e] = 1)),
                    t
                  );
                },
                this.hitMessageIdSet,
              )));
          }
          ((this.searching = !1),
            this.sendNotification("notify_search_messages", e));
        },
        saveSearchAttachmentsResult: function (e) {
          (null == this.attachmentsMap && (this.attachmentsMap = new B()),
            (this.attachmentsNextMarker = e.nextMarker));
          for (var t = 0, n = e.contents; t < n.length; ) {
            var i = n[t],
              r = (++t, i.attachmentId),
              r = "_" + r.high + "_" + r.low,
              o = this.attachmentsMap;
            null != Y[r] ? o.setReserved(r, i) : (o.h[r] = i);
          }
          ((this.searching = !1),
            this.sendNotification("notify_search_attachments", e));
        },
        getHitContext: function (e) {
          var t = this.contextMap;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        getFileInfo: function (e) {
          var t = this.attachmentsMap;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        isHitMessage: function (e) {
          var e = "_" + e.high + "_" + e.low,
            t = this.hitMessageIdSet;
          return null != Y[e] ? t.existsReserved(e) : t.h.hasOwnProperty(e);
        },
        reserveOnNormalizedKeywordPrepared: function (e) {
          var t,
            n = this;
          null != this.recentParams &&
            null != this.recentParams.keyword &&
            ((t = null == this.recentParams ? "" : this.recentParams.keyword),
            this.cachedKeyword != t &&
              ((this.cachedKeyword = t),
              null == this.cachedNormalizedKeywords
                ? (this.cachedNormalizedKeywords = [])
                : this.cachedNormalizedKeywords.splice(
                    0,
                    this.cachedNormalizedKeywords.length,
                  ),
              (t = new v("[ 　]", "g").split(I.htmlEscape(t, !0))),
              k.iter(t.filter(Ln.isNotEmpty), function (e) {
                n.cachedNormalizedKeywords.push(new Wt(e, !0, xo.normalize));
              })),
            e(this.cachedNormalizedKeywords));
        },
        __class__: ga,
      })),
      function () {
        (l.call(this, "session"), (this.createSessionCount = 0));
      }),
    Sa =
      (((t["albero.proxy.SessionProxy"] = ya).__name__ = [
        "albero",
        "proxy",
        "SessionProxy",
      ]),
      (ya.__super__ = l),
      (ya.prototype = e(l.prototype, {
        start: function (e) {
          this.api.createSession(e, z(this, this.onSessionCreated));
        },
        onSessionCreated: function (e) {
          for (
            var t,
              n,
              i = this,
              r =
                ((n =
                  !!e.notification &&
                  ((o = this.dataStore), (t = e.user.id), null != o.me) &&
                  null != (o = o.me.id) &&
                  null != t &&
                  o.high == t.high &&
                  o.low == t.low),
                this.dataStore),
              o = e.user,
              a = ((r.me = o), 0),
              s = r.getDomains();
            a < s.length;
          ) {
            var l = s[a];
            (++a,
              r.sendNotification(
                "notify_update_user",
                r.me.toDomainUser(l.id),
              ));
          }
          (r.sendNotification("current_user_changed", o),
            r.storageForUsers.setDirtyFlag(),
            this.settings.setConfiguration(e.configuration),
            this.settings.setDeviceId(e.device.device_id),
            this.sendNotification("data_recovering"),
            Promise.all([c.asyncPrepare()]).then(function (e) {
              (n
                ? i.prepareNecessaryDataFromCache()
                : (i.dataStore.clear(!1), i.prepareNecessaryDataFromServer()),
                i.createSessionCount++);
            }));
        },
        prepareNecessaryDataFromServer: function () {
          function n(n) {
            return new Promise(function (e, t) {
              n(function () {
                Kt.resolveVoid(e);
              });
            });
          }
          var i = this;
          this.api.resetNotification(function (e) {
            Promise.all([
              n(z((j = i.api), j.getDomains)).then(function (e) {
                var t = i.dataStore.getDomains();
                return Promise.all([
                  0 < t.length
                    ? n(z((j = i.api), j.getAnnouncementStatuses))
                    : Promise.resolve(),
                  Promise.all([
                    n(z((j = i.api), j.getFriends)),
                    n(z((j = i.api), j.getAcquaintances)),
                  ]).then(function (e) {
                    return n(z((j = i.api), j.getTalks));
                  }),
                ]);
              }),
              n(z((j = i.api), j.getAccountControlRequests)).then(function (e) {
                return n(z((j = i.api), j.getDomainInvites));
              }),
              n(z((j = i.api), j.getJoinedAccountControlGroup)),
            ]).then(function (e) {
              i.api.getAllUserIdentifiers(z(i, i.onNecessaryDataPrepared));
            });
          });
        },
        prepareNecessaryDataFromCache: function () {
          for (
            var e = Un.sortAndReturn(
                this.dataStore.getDomains(),
                function (e, t) {
                  var t = t.id,
                    e = e.id,
                    n = (t.high - e.high) | 0;
                  return (
                    0 == n && (n = w.ucompare(t.low, e.low)),
                    t.high < 0 ? (e.high < 0 ? n : -1) : 0 <= e.high ? n : 1
                  );
                },
              ),
              t = (this.api.sendDomainNotifications(e), 0);
            t < e.length;
          ) {
            var n = e[t],
              i = (++t, this.dataStore.getAnnouncementStatus(n.id));
            (null == i &&
              ((i = ue.fromDomainId(n.id)),
              this.dataStore.setAnnouncementStatus(i)),
              this.api.sendAnnouncementStatusNotification(i),
              this.api.sendAddFriendNotifications(
                this.dataStore.getFriends(n.id),
              ),
              this.api.sendAddAcquaintanceNotifications(
                this.dataStore.getAcquaintances(n.id),
              ));
          }
          (this.api.sendTalkNotifications(this.dataStore.getTalks()),
            this.api.sendTalkStatusNotifications(
              this.dataStore.getTalkStatuses(),
            ),
            this.api.sendAddDomainInviteNotifications(
              this.dataStore.getDomainInvites(),
            ),
            this.api.sendAddAccountControlRequestNotifications(
              this.dataStore.getAccountControlRequests(),
            ));
          var r = this.dataStore.getAccountControlGroup();
          (null != r && this.api.sendJoinAccountControlGroupNotification(r),
            this.onNecessaryDataPrepared());
        },
        onNecessaryDataPrepared: function () {
          var e,
            t = this.settings.getConfiguration().botExpiredVersion,
            t =
              (F.d("bot_expired_version", t),
              null != t &&
                ((e = require("../package.json").version),
                F.d("current", e),
                null != e) &&
                this.isVersionExpired(e, t) &&
                process.stderr.write(
                  "-----------------------------------------\nCurrent version is expired! (current: " +
                    e +
                    ")\nRun 'npm install direct-js' to update\n-----------------------------------------\n",
                ),
              this.sendNotification("data_recovered"),
              this.api.startNotification(),
              this.api.startUpdateLastUsedAtIfNeed(),
              new Te(this.dataStore.getDomains()).getDirectAppsIds());
          this.api.getDirectApps(t).catch(function (e) {
            F.v("getNoteStatuses error:", e);
          });
        },
        isVersionExpired: function (e, t) {
          function n(e) {
            return e
              .split(".")
              .map(function (e) {
                return P.parseInt(e);
              })
              .filter(function (e) {
                return null != e;
              });
          }
          var i,
            r = n(e),
            o = n(t);
          if (3 != r.length || 3 != o.length) return !1;
          for (i = 0, 0; i < 3; ) {
            var a = i++;
            if (o[a] > r[a]) return !0;
            if (o[a] < r[a]) return !1;
          }
          return !0;
        },
        __class__: ya,
      })),
      function () {
        ((this.deviceId = null),
          (this.selectedDomainIdOnTalkSelectionChanged = null),
          (this.selectedDomainId = null),
          (this.remember = !0),
          l.call(this, "settings"));
      }),
    va =
      (((t["albero.proxy.SettingsProxy"] = Sa).__name__ = [
        "albero",
        "proxy",
        "SettingsProxy",
      ]),
      (Sa.__interfaces__ = [Xe]),
      (Sa.__super__ = l),
      (Sa.prototype = e(l.prototype, {
        onRegister: function () {
          this.loadBrowserSettings();
        },
        getValidHost: function () {
          return "https://" + qs.host;
        },
        setAccessTokenRemember: function (e) {
          this.remember = e;
        },
        setAccessToken: function (e) {
          this.remember &&
            ((this.accessToken = e),
            this.sendNotification("access_token_changed", e));
        },
        getAccessToken: function () {
          return (
            null == this.accessToken && (this.accessToken = qs.accessToken),
            this.accessToken
          );
        },
        clearAccessToken: function () {
          this.accessToken = null;
        },
        getOs: function () {
          return "bot";
        },
        getIDFV: function () {
          var e = o.getItem("idfv");
          return (
            (null != e && "" != e) ||
              ((e = Ln.randomString()), o.setItem("idfv", e)),
            e
          );
        },
        setConfiguration: function (e) {
          ((this.configuration = e),
            this.sendNotification("configuration_changed", e));
        },
        getConfiguration: function () {
          return this.configuration;
        },
        setDeviceId: function (e) {
          this.deviceId = e;
        },
        getDeviceId: function () {
          return this.deviceId;
        },
        setSelectedDomainId: function (e) {
          var t = this.selectedDomainId;
          (null != t && null != e && t.high == e.high && t.low == e.low) ||
            (null != e
              ? (o.setItem("selected_domain_id_h", P.string(e.high)),
                o.setItem("selected_domain_id_l", P.string(e.low)))
              : (o.removeItem("selected_domain_id_h"),
                o.removeItem("selected_domain_id_l")),
            (this.selectedDomainId = e),
            this.sendNotification("domain_selection_changed", e));
        },
        getLastSelectedDomainId: function () {
          var e = o.getItemAsInt("selected_domain_id_h"),
            t = o.getItemAsInt("selected_domain_id_l");
          return null != e && null != t ? _.make(e, t) : null;
        },
        getSelectedDomainId: function () {
          return this.selectedDomainId;
        },
        getSelectedDomainIdOnTalkSelectionChanged: function () {
          return this.selectedDomainIdOnTalkSelectionChanged;
        },
        clearDomainSelection: function () {
          ((this.selectedDomainId = null),
            o.removeItem("selected_domain_id_h"),
            o.removeItem("selected_domain_id_l"));
        },
        clearSelectedTalk: function (e) {
          null == e && (e = yr.SingleTalkPane);
          var t = this.selectedDomainId;
          (this.assureTalkSelectionMap("_" + t.high + "_" + t.low).remove(e),
            (this.selectedDomainIdOnTalkSelectionChanged =
              this.selectedDomainId),
            this.sendNotification(
              "talk_selection_changed",
              new Zr(e, $r.NotSelected),
            ));
        },
        clearSelectedMessage: function (e, t, n) {
          var i = "_" + t.high + "_" + t.low;
          return (
            null != i &&
            ((i = this.assureTalkSelectionMap(i)),
            (n = $r.TalkSelected(n, m.None)),
            i.set(e, n),
            (i = this.selectedDomainId),
            null != t) &&
            null != i &&
            t.high == i.high &&
            t.low == i.low
          );
        },
        selectTalk: function (e, t, n, i) {
          var r = "_" + t.high + "_" + t.low;
          return (
            null != r &&
            ((r = this.assureTalkSelectionMap(r)),
            (n = $r.TalkSelected(n, i)),
            r.set(e, n),
            (i = this.selectedDomainId),
            null != t) &&
            null != i &&
            t.high == i.high &&
            t.low == i.low &&
            ((this.selectedDomainIdOnTalkSelectionChanged =
              this.selectedDomainId),
            this.sendNotification("talk_selection_changed", new Zr(e, n)),
            !0)
          );
        },
        selectAnnouncement: function (e) {
          var t;
          return (
            null != this.selectedDomainId &&
            ((t = "_" + (t = this.selectedDomainId).high + "_" + t.low),
            this.assureTalkSelectionMap(t).set(e, $r.AnnouncementsSelected),
            (this.selectedDomainIdOnTalkSelectionChanged =
              this.selectedDomainId),
            this.sendNotification(
              "talk_selection_changed",
              new Zr(e, $r.AnnouncementsSelected),
            ),
            !0)
          );
        },
        hasSelectedTalkOrAnnouncements: function (e) {
          null == e && (e = yr.SingleTalkPane);
          var t = this.getTalkSelectionMap(this.selectedDomainId);
          return null != t && t.get(e) != $r.NotSelected;
        },
        hasSelectedTalk: function (e) {
          null == e && (e = yr.SingleTalkPane);
          var t = this.getTalkSelectionMap(this.selectedDomainId);
          return null != t && null != (t = t.get(e)) && 0 != t[1];
        },
        getTalkSelectionMap: function (e) {
          var t;
          return null == this.paneTalkSelections || null == e
            ? null
            : ((e = "_" + e.high + "_" + e.low),
              (t = this.paneTalkSelections),
              null != Y[e] ? t.getReserved(e) : t.h[e]);
        },
        assureTalkSelectionMap: function (e) {
          null == this.paneTalkSelections &&
            (this.paneTalkSelections = new B());
          var t,
            n = this.paneTalkSelections,
            n = null != Y[e] ? n.getReserved(e) : n.h[e];
          return (
            null == n &&
              ((n = new es()),
              (t = this.paneTalkSelections),
              null != Y[e] ? t.setReserved(e, n) : (t.h[e] = n)),
            n
          );
        },
        getSelectedTalkId: function () {
          var e = this.getTalkSelection(
            this.selectedDomainId,
            yr.SingleTalkPane,
          );
          return 1 == e[1] ? e[2] : null;
        },
        isSelectedTalkOrAnnouncements: function (e, t) {
          return eo.isSelected(
            this.getTalkSelection(this.selectedDomainId, e),
            t,
          );
        },
        getTalkSelection: function (e, t) {
          return null == this.paneTalkSelections ||
            null == (e = this.getTalkSelectionMap(e)) ||
            null == (e = e.get(t))
            ? $r.NotSelected
            : e;
        },
        setSelectedStampsetType: function (e) {
          o.setItem("selected_stamp_tab_id", e.toTabId());
        },
        getSelectedStampsetType: function () {
          return Vr.fromTabId(o.getItem("selected_stamp_tab_id"));
        },
        clearSelectedStampsetType: function () {
          o.removeItem("selected_stamp_tab_id");
        },
        setInputTextForTalkId: function (e, t) {
          null == e || 0 == e.length
            ? this.clearInputTextForTalkId(t)
            : (this.loadInputTextForAll(),
              (this.inputTexts["_" + t.high + "_" + t.low] = e));
        },
        getInputTextForTalkId: function (e) {
          this.loadInputTextForAll();
          e = this.inputTexts["_" + e.high + "_" + e.low];
          return null != e ? e : "";
        },
        clearInputTextForTalkId: function (e) {
          (this.loadInputTextForAll(),
            delete this.inputTexts["_" + e.high + "_" + e.low]);
        },
        loadInputTextForAll: function () {
          null == this.inputTexts &&
            ((this.inputTexts = o.getItemAsJson("input_text")),
            null == this.inputTexts) &&
            (this.inputTexts = {});
        },
        saveInputTextForAll: function () {
          var e;
          null != this.inputTexts &&
            ((e = JSON.stringify(this.inputTexts)), o.setItem("input_text", e));
        },
        clearInputTextForAll: function () {
          (o.removeItem("input_text"), (this.inputTexts = null));
        },
        setCopyProfileToAllDomains: function (e) {
          e
            ? o.setItem("copy_profile_to_all_domains", "true")
            : o.removeItem("copy_profile_to_all_domains");
        },
        isCopyProfileToAllDomains: function () {
          return null != o.getItem("copy_profile_to_all_domains");
        },
        setSearchHistoriesForSelectedDomain: function (e, t) {
          var n, i;
          null != this.selectedDomainId &&
            ((n = "_" + (n = this.selectedDomainId).high + "_" + n.low),
            (i = this.browserSettings.createCopy()).setSearchHistories(
              "_" + e.high + "_" + e.low,
              n,
              t,
            ),
            this.applyUserDataSettings(i));
        },
        getOldSearchHistoriesForSelectedDomain: function () {
          if (null == this.selectedDomainId) return null;
          this.loadSearchHistoryForAll();
          var e = this.selectedDomainId;
          return this.searchHistories["_" + e.high + "_" + e.low];
        },
        getSearchHistoriesForSelectedDomain: function (e) {
          var t;
          return null == this.selectedDomainId
            ? null
            : ((t = this.selectedDomainId),
              this.browserSettings.getSearchHistories(
                "_" + e.high + "_" + e.low,
                "_" + t.high + "_" + t.low,
              ));
        },
        loadSearchHistoryForAll: function () {
          null == this.searchHistories &&
            ((this.searchHistories = o.getItemAsJson("search_histories")),
            null == this.searchHistories) &&
            (this.searchHistories = {});
        },
        setSelectedDepartmentId: function (e, t) {
          var n;
          null != e &&
            (this.loadSelectedDepartmentIdsForAll(),
            null == t
              ? this.selectedDepartmentIds.remove("_" + e.high + "_" + e.low)
              : ((e = "_" + e.high + "_" + e.low),
                (n = this.selectedDepartmentIds),
                null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t)),
            this.saveSelectedDepartmentIdsForAll());
        },
        getSelectedDepartmentId: function (e) {
          if (null == e) return null;
          this.loadSelectedDepartmentIdsForAll();
          var e = "_" + e.high + "_" + e.low,
            t = this.selectedDepartmentIds;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        loadSelectedDepartmentIdsForAll: function () {
          null == this.selectedDepartmentIds &&
            ((this.selectedDepartmentIds = o.getItemWithUnserialize(
              "selected_department_ids",
            )),
            null == this.selectedDepartmentIds) &&
            (this.selectedDepartmentIds = new B());
        },
        saveSelectedDepartmentIdsForAll: function () {
          null != this.selectedDepartmentIds &&
            o.setItemWithSerialize(
              "selected_department_ids",
              this.selectedDepartmentIds,
            );
        },
        setRightPanelOpened: function (e) {
          var t = this.browserSettings.createCopy();
          ((t.talkRightPanelOpened = e), this.applyTalkSettings(t));
        },
        isRightPanelOpened: function () {
          return this.browserSettings.talkRightPanelOpened;
        },
        setRightPanelWidth: function (e) {
          var t = this.browserSettings.createCopy();
          ((t.talkRightPanelWidth = e), this.applyTalkSettings(t));
        },
        getRightPanelWidth: function () {
          return this.browserSettings.talkRightPanelWidth;
        },
        setMultiViewMode: function (e) {
          e
            ? o.setItem("multi_view_mode", "true")
            : o.removeItem("multi_view_mode");
        },
        isMultiViewMode: function () {
          return null != o.getItem("multi_view_mode");
        },
        setPasswordWarningSkipUntil: function (e) {
          null == e
            ? o.removeItem("password_warning_skip_until")
            : o.setItem(
                "password_warning_skip_until",
                "_" + e.high + "_" + e.low,
              );
        },
        getPasswordWarningSkipUntil: function () {
          var e = o.getItem("password_warning_skip_until");
          return null == e ? null : D.makeFromIdStr(e);
        },
        setLastUsedExpiredAt: function (e) {
          null == e
            ? o.removeItem("last_used_expired_at")
            : o.setItem("last_used_expired_at", "_" + e.high + "_" + e.low);
        },
        getLastUsedExpiredAt: function () {
          var e = o.getItem("last_used_expired_at");
          return null == e ? null : D.makeFromIdStr(e);
        },
        setOnPremise: function (e) {
          this.onPremises = e;
        },
        isOnPremises: function () {
          return this.onPremises;
        },
        updateThumbnailExpansionEnabled: function (e, t) {
          var n = this.browserSettings.createCopy();
          (n.setThumbnailExpansionEnabled("_" + e.high + "_" + e.low, t),
            this.applyUserDataSettings(n));
        },
        isThumbnailExpansionEnabled: function (e) {
          return this.browserSettings.isThumbnailExpansionEnabled(
            "_" + e.high + "_" + e.low,
          );
        },
        getBrowserSettings: function () {
          return this.browserSettings;
        },
        applyNotificationSettings: function (e) {
          return (
            e.copyNotificationSettingsTo(this.browserSettings),
            this.saveBrowserSettings()
          );
        },
        applyKeywordWatchingSettings: function (e) {
          return (
            e.copyKeywordSettingsTo(this.browserSettings),
            this.saveBrowserSettings()
          );
        },
        applyTalkSettings: function (e) {
          return (
            e.copyTalkSettingsTo(this.browserSettings),
            this.saveBrowserSettings()
          );
        },
        applyUserDataSettings: function (e) {
          return (
            e.copyUserDataContainerTo(this.browserSettings),
            this.saveBrowserSettings()
          );
        },
        applyLanguageSettings: function (e) {
          return (
            e.copyLanguageSettingsTo(this.browserSettings),
            this.saveBrowserSettings()
          );
        },
        applyConferenceSettings: function (e) {
          return (
            e.copyConferenceSettingsTo(this.browserSettings),
            this.saveBrowserSettings()
          );
        },
        applyDomainsSortModeSettings: function (e) {
          return (
            e.copyDomainsSortModeSettingsTo(this.browserSettings),
            this.saveBrowserSettings()
          );
        },
        applyAppLeftPanelOpenedSettings: function (e) {
          return (
            e.copyAppLeftPanelOpenedSettingsTo(this.browserSettings),
            this.saveBrowserSettings()
          );
        },
        applyAppLeftPanelWidthSettings: function (e) {
          return (
            e.copyAppLeftPanelWidthSettingsTo(this.browserSettings),
            this.saveBrowserSettings()
          );
        },
        applyScheduledMessageDefaultTimeSettings: function (e) {
          return (
            e.copyScheduledMessageDefaultTimeStrSettingsTo(
              this.browserSettings,
            ),
            this.saveBrowserSettings()
          );
        },
        getAppLeftPanelWidth: function () {
          return this.browserSettings.appLeftPanelWidth;
        },
        loadBrowserSettings: function () {
          this.browserSettings = _e.fromJson(o.getItemAsJson(_e.NAME));
        },
        saveBrowserSettings: function () {
          var e = o.getItemAsJson(_e.NAME);
          return (
            this.browserSettings.copyTo((e = null == e ? {} : e)),
            o.setItem(_e.NAME, JSON.stringify(e))
          );
        },
        isLegacyNoteEnabled: function () {
          return this.browserSettings.legacyNoteEnabled;
        },
        isLeaveTalksEnabled: function () {
          return this.browserSettings.leaveTalksEnabled;
        },
        __class__: Sa,
      })),
      function () {
        l.call(this, "talksService");
      }),
    Ta =
      (((t["albero.proxy.TalksServiceProxy"] = va).__name__ = [
        "albero",
        "proxy",
        "TalksServiceProxy",
      ]),
      (va.__super__ = l),
      (va.prototype = e(l.prototype, {
        getDomainTalksOrderedByTimestampDesc: function (e) {
          return Un.reverseAndReturn(
            this.getDomainTalksOrderedByTimestampAsc(e),
          );
        },
        getDomainTalksOrderedByTimestampAsc: function (e) {
          return Un.sortAndReturn(
            this.getDomainTalks(e).map(z(this, this.createTalkPack)),
            Qe.compareTalkOrderingTimestamp,
          ).map(z(this, this.getTalkFromTalkPack));
        },
        getDomainTalksCount: function (e) {
          return this.getDomainTalks(e).length;
        },
        hasConferenceInTalk: function (e) {
          return null != e && this.dataStore.hasConferenceInTalk(e.id);
        },
        isFavoriteTalk: function (e) {
          e = this.dataStore.getTalkStatus(e.id);
          return null != e && null != e.orderInFavorites;
        },
        hasUnread: function (e) {
          e = this.dataStore.getTalkStatus(e.id);
          return null != e && 0 != e.calcUnreadCount();
        },
        getTalkFromTalkPack: function (e) {
          return e.talk;
        },
        getTalkIdFromTalk: function (e) {
          return e.id;
        },
        getDomainTalks: function (t) {
          return null == t
            ? []
            : this.dataStore.getTalks().filter(function (e) {
                e = e.domainId;
                return (
                  null != e && null != t && e.high == t.high && e.low == t.low
                );
              });
        },
        createTalkPack: function (e) {
          return new Qe(e, this.dataStore.getTalkStatus(e.id));
        },
        __class__: va,
      })),
      function () {
        l.call(this, "userPresences");
      }),
    Ia =
      (((t["albero.proxy.UserPresencesProxy"] = Ta).__name__ = [
        "albero",
        "proxy",
        "UserPresencesProxy",
      ]),
      (Ta.__super__ = l),
      (Ta.prototype = e(l.prototype, {
        getUserPresenceState: function (e, t) {
          var n = h.map(this.getUserPresence(e), function (e) {
            return e.toState(t);
          });
          switch (n[1]) {
            case 0:
              return n[2];
            case 1:
              return no.GT_60;
          }
        },
        getUserPresence: function (e) {
          return h.option(this.dataStore.getUserPresence(e));
        },
        getLastUsedAtString: function (e) {
          var t = h.map(
            h.flatMap(this.getUserPresence(e), function (e) {
              return e.lastUsedAtOpt;
            }),
            function (e) {
              return W.datetimeString(e);
            },
          );
          switch (t[1]) {
            case 0:
              return t[2];
            case 1:
              return " - ";
          }
        },
        __class__: Ta,
      })),
      function () {
        l.call(this, "wsApi");
      }),
    Ea =
      (((t["albero.proxy.WsApiProxy"] = Ia).__name__ = [
        "albero",
        "proxy",
        "WsApiProxy",
      ]),
      (Ia.__super__ = l),
      (Ia.prototype = e(l.prototype, {
        createAccessToken: function (e, t, n, i) {
          return this.apiCaller.callImmediately("create_access_token", [
            e,
            t,
            n,
            i,
            "",
          ]);
        },
        createAccessTokenById: function (e, t, n, i, r) {
          return this.apiCaller.callImmediately("create_access_token_by_id", [
            e,
            t,
            n,
            i,
            r,
            "",
          ]);
        },
        authorizeDevice: function (e, t) {
          return this.apiCaller.callImmediately("authorize_device", [e, t]);
        },
        createSession: function (e, t, n) {
          var i = this;
          return this.apiCaller
            .callImmediatelyReliable(
              "create_session",
              null != n ? [e, t, n] : [e, t],
              function (e) {
                return i.shouldRetryCreateSession(e)
                  ? m.Some(i.apiCaller.retryTimeForReplicationLag())
                  : m.None;
              },
            )
            .then(function (e) {
              return new je(e);
            });
        },
        shouldRetryCreateSession: function (e) {
          return 401 == e.code && "invalid token" == e.message;
        },
        startNotification: function () {
          return this.apiCaller.call("start_notification", []);
        },
        resetNotification: function () {
          return this.apiCaller.call("reset_notification", []);
        },
        updateLastUsedAt: function () {
          return this.apiCaller.call("update_last_used_at", []);
        },
        addFriend: function (e, t) {
          return this.apiCaller.call("add_friend", [e, t]).then(function (e) {
            return { domainId: e[0], domainUser: e[1] };
          });
        },
        deleteFriend: function (e, t) {
          return this.apiCaller
            .call("delete_friend", [e, t])
            .then(function (e) {
              return { domainId: e[0], userId: e[1] };
            });
        },
        getFriends: function () {
          return this.apiCaller.call("get_friends", []).then(function (e) {
            return e.map(function (e) {
              return {
                domainId: e[0],
                domainUsers: e[1].map(function (e) {
                  return new ve(e);
                }),
              };
            });
          });
        },
        getAcquaintances: function () {
          return this.apiCaller
            .call("get_acquaintances", [])
            .then(function (e) {
              return e.map(function (e) {
                return {
                  domainId: e[0],
                  domainUsers: e[1].map(function (e) {
                    return new ve(e);
                  }),
                };
              });
            });
        },
        getAllUsers: function (e, t, n, i) {
          t = k.fold(
            t,
            function (e, t) {
              return e + t;
            },
            0,
          );
          return this.apiCaller
            .call("get_domain_users", [e, t, n, i])
            .then(function (e) {
              return {
                marker: e.marker,
                nextMarker: e.next_marker,
                contents: e.contents.map(function (e) {
                  return new ve(e);
                }),
              };
            });
        },
        searchDomainUsers: function (e, t, n, i, r) {
          n = k.fold(
            n,
            function (e, t) {
              return e + t;
            },
            0,
          );
          return this.apiCaller
            .call("search_domain_users", [e, t, n, i, r])
            .then(function (e) {
              return {
                marker: e.marker,
                nextMarker: e.next_marker,
                contents: e.contents.map(function (e) {
                  return new ve(e);
                }),
              };
            });
        },
        getUsers: function (e, t) {
          return this.apiCaller.call("get_users", [e, t]).then(function (e) {
            return e.map(function (e) {
              return new ve(e);
            });
          });
        },
        getUserIdentifiers: function (e, t) {
          return this.apiCaller
            .call("get_user_identifiers", [e, t])
            .then(function (e) {
              var t = e.contents.map(function (e) {
                return new rt(e);
              });
              return { domainId: e.domainId, contents: t };
            });
        },
        getProfile: function (e, t) {
          return this.apiCaller.call("get_profile", [e, t]).then(function (e) {
            return new Tr(e);
          });
        },
        updateUser: function (e, t, n, i) {
          return this.apiCaller
            .call("update_user", [e, t, n, i])
            .then(function (e) {
              return new fr(e);
            });
        },
        updateProfile: function (e) {
          e = Nr.convertToApiParams(e);
          return this.apiCaller.call("update_profile", [e]).then(function (e) {
            return new fr(e);
          });
        },
        getPresences: function (e, t) {
          return this.apiCaller
            .call("get_presences", [e, t])
            .then(function (e) {
              var t = W.nowAsInt64();
              return e.map(function (e) {
                return new ot(e, t);
              });
            });
        },
        getDomains: function () {
          return this.apiCaller.call("get_domains", []).then(function (e) {
            return e.map(function (e) {
              return new ye(e);
            });
          });
        },
        leaveDomain: function (e) {
          return this.apiCaller.call("leave_domain", [e]);
        },
        getDomainInvites: function () {
          return this.apiCaller
            .call("get_domain_invites", [])
            .then(function (e) {
              return e.map(function (e) {
                return new Se(e);
              });
            });
        },
        acceptDomainInvite: function (e) {
          return this.apiCaller.call("accept_domain_invite", [e]);
        },
        deleteDomainInvite: function (e) {
          return this.apiCaller.call("delete_domain_invite", [e]);
        },
        getDepartmentTree: function (e) {
          return this.apiCaller
            .call("get_department_tree", [e])
            .then(function (e) {
              return new fe(e);
            });
        },
        getDepartmentUsers: function (e, t, n, i, r) {
          return this.apiCaller
            .call("get_department_users", [e, t, n, i, r])
            .then(function (e) {
              return {
                marker: e.marker,
                nextMarker: e.next_marker,
                contents: e.contents.map(function (e) {
                  return new ve(e);
                }),
              };
            });
        },
        getMe: function () {
          return this.apiCaller.call("get_me", []).then(function (e) {
            return new fr(e);
          });
        },
        getDepartmentUserCount: function (e, t) {
          return this.apiCaller
            .call("get_department_user_count", [e, t])
            .then(function (e) {
              return new pe(e);
            });
        },
        createGroupTalk: function (e, t, n, i) {
          return this.apiCaller
            .call("create_group_talk", [
              e,
              t,
              n.allowDisplayPastMessages,
              (i = null == i ? !0 : i),
            ])
            .then(function (e) {
              return new qr(e);
            });
        },
        createPairTalk: function (e, t) {
          return this.apiCaller
            .call("create_pair_talk", [e, t])
            .then(function (e) {
              return new qr(e);
            });
        },
        updateGroupTalk: function (e, t, n, i) {
          return this.apiCaller
            .call("update_group_talk", [e, t, n, i])
            .then(function (e) {
              return new qr(e);
            });
        },
        addTalkers: function (e, t) {
          return this.apiCaller
            .call("add_talkers", [e, t, !0])
            .then(function (e) {
              return new qr(e);
            });
        },
        deleteTalker: function (e, t) {
          return this.apiCaller.call("delete_talker", [e, t]);
        },
        getTalks: function () {
          return this.apiCaller.call("get_talks", []).then(function (e) {
            return e.map(function (e) {
              return new qr(e);
            });
          });
        },
        getTalkStatuses: function () {
          return this.apiCaller
            .call("get_talk_statuses", [])
            .then(function (e) {
              return e.map(function (e) {
                return new Ze(e);
              });
            });
        },
        getReadStatus: function (e, t) {
          return this.apiCaller
            .call("get_read_status", [e, t])
            .then(function (e) {
              return new Oe(e);
            });
        },
        addFavoriteTalk: function (e, t) {
          return this.apiCaller.call("add_favorite_talk", [e, t]);
        },
        deleteFavoriteTalk: function (e, t) {
          return this.apiCaller.call("delete_favorite_talk", [e, t]);
        },
        disablePushNotification: function (e, t) {
          return this.apiCaller.call("disable_push_notification", [e, t]);
        },
        enablePushNotification: function (e, t, n) {
          return this.apiCaller.call("enable_push_notification", [
            e,
            t,
            br.toRawValue(n),
          ]);
        },
        createAnnouncement: function (e, t, n) {
          return this.apiCaller
            .call("create_announcement", [e, O.enumIndex(t), n])
            .then(function (e) {
              return new Ri(e);
            });
        },
        getAnnouncements: function (e, t, n, i) {
          return this.apiCaller
            .call("get_announcements", [
              e,
              t,
              n.sinceId,
              n.maxId,
              xe.orderToInt(i),
            ])
            .then(function (e) {
              return e.map(function (e) {
                return new Ri(e);
              });
            });
        },
        getAnnouncementStatuses: function () {
          return this.apiCaller
            .call("get_announcement_statuses", [])
            .then(function (e) {
              return e.map(function (e) {
                return new ue(e);
              });
            });
        },
        getAnnouncementStatus: function (e) {
          var t = this;
          return this.apiCaller
            .callApiReliable("get_announcement_status", [e], function (e) {
              return 400 == e.code && "invalid domain_id" == e.message
                ? m.Some(t.apiCaller.retryTimeForReplicationLag())
                : m.None;
            })
            .then(function (e) {
              return new ue(e);
            });
        },
        updateAnnouncementReadStatus: function (e, t) {
          return this.apiCaller.call("update_announcement_status", [e, t]);
        },
        getMessages: function (e, t, n, i) {
          return this.apiCaller
            .call("get_messages", [e, t, n.sinceId, n.maxId, xe.orderToInt(i)])
            .then(function (e) {
              return e.map(function (e) {
                return new O(e);
              });
            });
        },
        createMessage: function (e, t, n) {
          var i = this;
          return t == d.file || t == d.textMultipleFile
            ? this.apiCaller
                .callApiReliable(
                  "create_message",
                  [e, O.enumIndex(t), n],
                  function (e) {
                    return i.shouldRetryCreateMessage(e)
                      ? m.Some(i.apiCaller.retryTimeForReplicationLag())
                      : m.None;
                  },
                )
                .then(function (e) {
                  return new O(e);
                })
            : this.apiCaller
                .call("create_message", [e, O.enumIndex(t), n])
                .then(function (e) {
                  return new O(e);
                });
        },
        shouldRetryCreateMessage: function (e) {
          return (
            null != e.code &&
            (500 == e.code
              ? "Internal Error (IllegalStateException)" == e.message
              : 400 == e.code && "invalid file_id" == e.message)
          );
        },
        deleteMessage: function (e, t) {
          return this.apiCaller.call("delete_message", [e, t]);
        },
        scheduleMessage: function (t, n, i, r) {
          var e,
            o,
            a = this;
          return (
            i == d.file || i == d.textMultipleFile
              ? ((e = this.apiCaller),
                (o = O.enumIndex(i)),
                e.callApiReliable(
                  "schedule_message",
                  [t, n, o, r],
                  function (e) {
                    return a.shouldRetryScheduleMessage(e)
                      ? m.Some(a.apiCaller.retryTimeForReplicationLag())
                      : m.None;
                  },
                ))
              : ((e = this.apiCaller),
                (o = O.enumIndex(i)),
                e.call("schedule_message", [t, n, o, r]))
          ).then(function (e) {
            return fo.fromParams({
              id: e,
              talkId: n,
              createMessageAt: t,
              type: i,
              content: r,
            });
          });
        },
        rescheduleMessage: function (e, t, n, i, r) {
          var o,
            a,
            s = this;
          return (
            i == d.file || i == d.textMultipleFile
              ? ((o = this.apiCaller),
                (a = O.enumIndex(i)),
                o.callApiReliable(
                  "reschedule_message",
                  [e, t, n, a, r],
                  function (e) {
                    return s.shouldRetryScheduleMessage(e)
                      ? m.Some(s.apiCaller.retryTimeForReplicationLag())
                      : m.None;
                  },
                ))
              : ((o = this.apiCaller),
                (a = O.enumIndex(i)),
                o.call("reschedule_message", [e, t, n, a, r]))
          ).then(function (e) {
            return fo.fromParams({
              id: e,
              talkId: n,
              createMessageAt: t,
              type: i,
              content: r,
            });
          });
        },
        shouldRetryScheduleMessage: function (e) {
          return (
            null != e.code &&
            (500 == e.code
              ? "Internal Error (IllegalStateException)" == e.message
              : 400 == e.code && "invalid file_id" == e.message)
          );
        },
        deleteScheduledMessage: function (e) {
          return this.apiCaller.call("delete_scheduled_message", [e]);
        },
        getScheduledMessages: function (t, n, e, i) {
          return this.apiCaller
            .call("get_scheduled_messages", [t, n, e, i])
            .then(function (e) {
              return new ut(t, n, e);
            });
        },
        updateReadStatuses: function (e, t) {
          var n = this;
          return this.apiCaller.callApiReliable(
            "update_read_statuses",
            [e, t],
            function (e) {
              return 429 == e.code && "too many requests" == e.message
                ? m.Some(1e3 * n.asInt(e.detail.retry_after))
                : m.None;
            },
          );
        },
        searchMessagesAroundDateTime: function (e, t) {
          return this.apiCaller
            .call("search_messages_around_datetime", [e, t])
            .then(function (e) {
              return new ct(t, e);
            });
        },
        searchMessages: function (e, t, n, i, r) {
          return this.apiCaller
            .call("search_messages", [e, t, n, i, r])
            .then(function (e) {
              return new Be(e);
            });
        },
        searchAttachments: function (e, t, n, i, r) {
          return this.apiCaller
            .call("search_attachments", [e, t, n, i, r])
            .then(function (e) {
              return new Pe(e);
            });
        },
        createUploadAuth: function (e, t, n, i, r) {
          return this.apiCaller
            .call("create_upload_auth", [e, t, n, i, it.getUseTypeInt(r)])
            .then(function (e) {
              return new $t(e);
            });
        },
        createDownloadAuth: function (e, t, n) {
          var i = this;
          return this.apiCaller
            .callApiReliable("create_download_auth", [e, t, n], function (e) {
              return 429 == e.code && "too many requests" == e.message
                ? m.Some(1e3 * i.asInt(e.detail.retry_after))
                : m.None;
            })
            .then(function (e) {
              return new Ie(e);
            });
        },
        getFilePreview: function (e) {
          return this.apiCaller
            .call("get_file_preview", [e])
            .then(function (e) {
              return new Et(e);
            });
        },
        createFilePreview: function (e) {
          return this.apiCaller
            .call("create_file_preview", [e])
            .then(function (e) {
              return new Et(e);
            });
        },
        deleteAttachment: function (e, t) {
          return this.apiCaller.call("delete_attachment", [e, t]);
        },
        getAttachments: function (e, t, n, i) {
          return this.apiCaller
            .call("get_attachments", [e, t, n, i])
            .then(function (e) {
              return e.map(function (e) {
                return new ar(e);
              });
            });
        },
        getQuestions: function (e, t, n, i, r, o) {
          n = Cr.fromTypeToInt(n);
          return this.apiCaller
            .call("get_actions", [e, t, n, i, r, o.sinceId, o.maxId])
            .then(function (e) {
              return e.map(function (e) {
                return new Cr(e);
              });
            });
        },
        getQuestion: function (e) {
          return this.apiCaller
            .callApiReliable("get_action", [e])
            .then(function (e) {
              return new Cr(e);
            });
        },
        addFavoriteMessage: function (e, t, n) {
          return this.apiCaller
            .call("add_favorite_message", [e, t, n])
            .then(function (e) {
              return new at(e);
            });
        },
        deleteFavoriteMessage: function (e, t) {
          return this.apiCaller
            .call("delete_favorite_message", [e, t])
            .then(function (e) {
              return new st(e);
            });
        },
        getFavoriteMessages: function (i, r, e, t) {
          return this.apiCaller
            .call("get_favorite_messages", [i, r, e, t])
            .then(function (e) {
              var t = e.marker,
                n = e.next_marker,
                e = e.contents.map(function (e) {
                  return new O(e);
                });
              return new lt(i, r, t, n, e);
            });
        },
        getNoteStatuses: function (e, t, n, i) {
          return this.apiCaller
            .call("get_note_statuses", [e, t, n, i])
            .then(function (e) {
              return new gt(e);
            });
        },
        getNote: function (e) {
          return this.apiCaller
            .call("get_note", [e.value])
            .then(function (e) {
              return new pt(e);
            })
            .then(function (e) {
              return e.uncompress().catch(function (e) {
                e = _a.toLocalError(
                  R.NoteContentUncompressError(e),
                  ca.empty(),
                );
                return Promise.reject(e);
              });
            });
        },
        createNote: function (i, r, o, t, a) {
          var s = this;
          return new J()
            .run(t)
            .catch(function (e) {
              return (F.e("Failed to compress rich_text:", e), t);
            })
            .then(function (e) {
              var t = yo.getValue(o),
                n = z(s, s.retryTimeForCreateOrUpdateNoteFromError);
              return s.apiCaller.callApiReliable(
                "create_note",
                [i, r, t, e, a],
                n,
              );
            })
            .then(function (e) {
              return new dt(e);
            })
            .then(function (e) {
              return e.uncompress().catch(function (e) {
                e = _a.toLocalError(
                  R.NoteContentUncompressError(e),
                  ca.empty(),
                );
                return Promise.reject(e);
              });
            });
        },
        updateNoteSetting: function (e, t, n) {
          return this.apiCaller
            .call("update_note_setting", [e.value, t, n])
            .then(function (e) {
              return new Tt(e);
            })
            .then(function (e) {
              return e.uncompress().catch(function (e) {
                e = _a.toLocalError(
                  R.NoteContentUncompressError(e),
                  ca.empty(),
                );
                return Promise.reject(e);
              });
            });
        },
        updateNote: function (i, r, o, a, t, s) {
          var l = this;
          return new J()
            .run(t)
            .catch(function (e) {
              return (F.e("Failed to compress rich_text:", e), t);
            })
            .then(function (e) {
              var t = yo.getValue(a),
                n = z(l, l.retryTimeForCreateOrUpdateNoteFromError);
              return l.apiCaller.callApiReliable(
                "update_note",
                [i.value, r, o, t, e, s],
                n,
              );
            })
            .then(function (e) {
              return new St(e);
            })
            .then(function (e) {
              return e.uncompress().catch(function (e) {
                e = _a.toLocalError(
                  R.NoteContentUncompressError(e),
                  ca.empty(),
                );
                return Promise.reject(e);
              });
            });
        },
        deleteNote: function (e) {
          return this.apiCaller
            .call("delete_note", [e.value])
            .then(function (e) {
              return new ft(e);
            });
        },
        lockNote: function (e, t) {
          return this.apiCaller.call("lock_note", [e.value, t]);
        },
        unlockNote: function (e, t) {
          return this.apiCaller.call("unlock_note", [e.value, t]);
        },
        retryTimeForCreateOrUpdateNoteFromError: function (e) {
          return this.shouldRetryCreateOrUpdateNote(e)
            ? m.Some(this.apiCaller.retryTimeForReplicationLag())
            : m.None;
        },
        shouldRetryCreateOrUpdateNote: function (e) {
          return (
            null != e.code &&
            (500 == e.code
              ? "Internal Error (IllegalStateException)" == e.message
              : 400 == e.code && "invalid file_id" == e.message)
          );
        },
        setMessageReaction: function (e, t, n) {
          return this.apiCaller.call("set_message_reaction", [e, t, n]);
        },
        resetMessageReaction: function (e, t) {
          return this.apiCaller.call("reset_message_reaction", [e, t]);
        },
        getMessageReactionSummaries: function (e, t) {
          return this.apiCaller
            .callApiReliable("get_message_reaction_summaries", [e, t])
            .then(function (e) {
              return e.map(function (e) {
                return new jt(e);
              });
            });
        },
        getMessageReactionUsers: function (e, t) {
          return this.apiCaller
            .call("get_message_reaction_users", [e, t])
            .then(function (e) {
              return new zt(e);
            });
        },
        getAvailableMessageReactions: function (e) {
          return this.apiCaller
            .call("get_available_message_reactions", [e])
            .then(function (e) {
              return e.map(function (e) {
                return new Mo(e);
              });
            });
        },
        getSolutions: function (e, t) {
          return this.apiCaller
            .call("get_solutions", [e, t])
            .then(function (e) {
              return e.map(function (e) {
                return new ze(e);
              });
            });
        },
        getConferences: function () {
          return this.apiCaller.call("get_conferences", []).then(function (e) {
            return e.map(function (e) {
              return new he(e);
            });
          });
        },
        joinConference: function (e, t) {
          return this.apiCaller
            .call("join_conference", [e, t])
            .then(function (e) {
              return new ke(e);
            });
        },
        rejectConference: function (e) {
          return this.apiCaller.call("reject_conference", [e]);
        },
        leaveConference: function (e) {
          return this.apiCaller.call("leave_conference", [e]);
        },
        getConferenceParticipants: function (e) {
          return this.apiCaller.call("get_conference_participants", [e]);
        },
        getAccountControlRequests: function () {
          return this.apiCaller
            .call("get_account_control_requests", [])
            .then(function (e) {
              return e.map(function (e) {
                return new le(e);
              });
            });
        },
        getJoinedAccountControlGroup: function () {
          return this.apiCaller
            .call("get_joined_account_control_group", [])
            .then(function (e) {
              return e.map(function (e) {
                return new ae(e);
              });
            });
        },
        acceptAccountControlRequest: function (e) {
          return this.apiCaller.call("accept_account_control_request", [
            e.id,
            e.version,
          ]);
        },
        rejectAccountControlRequest: function (e) {
          return this.apiCaller.call("reject_account_control_request", [
            e.id,
            e.version,
          ]);
        },
        getStampsets: function (e) {
          return this.apiCaller.call("get_stampsets", [e]).then(function (e) {
            return e.map(Wr.originalStampset);
          });
        },
        getFlowNotificationBadges: function (e) {
          return this.apiCaller
            .call("get_flow_notification_badges", [e])
            .then(function (e) {
              return e.map(function (e) {
                return new be(e);
              });
            });
        },
        getDirectApps: function (e) {
          return this.apiCaller.call("get_direct_apps", [e]).then(function (e) {
            return e.map(function (e) {
              return new ge(e);
            });
          });
        },
        asInt: function (e) {
          return G.__cast(e, js);
        },
        __class__: Ia,
      })),
      (t["albero_cli.UpdateNoteAttachment"] = {
        __ename__: ["albero_cli", "UpdateNoteAttachment"],
        __constructs__: ["LocalFileAttachment", "RemoteFileAttachment"],
      })),
    wa =
      ((Ea.LocalFileAttachment = function (e) {
        e = ["LocalFileAttachment", 0, e];
        return ((e.__enum__ = Ea), (e.toString = a), e);
      }),
      (Ea.RemoteFileAttachment = function (e) {
        e = ["RemoteFileAttachment", 1, e];
        return ((e.__enum__ = Ea), (e.toString = a), e);
      }),
      (Ea.__empty_constructs__ = []),
      ((t["albero_cli.NoteUtil"] = p).__name__ = ["albero_cli", "NoteUtil"]),
      (p.fqen = function (e) {
        if (-1 == p.NOTE_EVENT_NAMES.indexOf(e))
          throw new Error("unknown event name: " + e);
        return "direct.notes." + e;
      }),
      (p.validateCreateNoteInput = function (e) {
        if (p.isStringNullOrEmpty(e.note_title))
          return $a.Left(
            "'note_title' must be specified and must not be empty.",
          );
        if (p.isStringNullOrEmpty(e.note_content))
          return $a.Left(
            "'note_content' must be specified and must not be empty.",
          );
        var t = null;
        if (
          e.note_attachments instanceof Array &&
          null == e.note_attachments.__enum__ &&
          p.hasPathField(e.note_attachments)
        )
          t = e.note_attachments;
        else {
          if (null != e.note_attachments)
            return $a.Left(
              "'note_attachments' is an invalid value. Please check: " +
                P.string(e.note_attachments),
            );
          t = [];
        }
        var n = p.validateNoteContentType(e);
        switch (n[1]) {
          case 0:
            return $a.Left(n[2]);
          case 1:
            var i = n[2];
            switch (i[1]) {
              case 0:
                var r = i[2];
                if (r == go.RICH_TEXT) {
                  var o = p.validateXmlContent(e);
                  if (0 === o[1]) return $a.Left(o[2]);
                }
                return $a.Right({
                  title: e.note_title,
                  contentType: r,
                  content: e.note_content,
                  attachments: t,
                });
              case 1:
                return $a.Right({
                  title: e.note_title,
                  content: e.note_content,
                  attachments: t,
                });
            }
        }
      }),
      (p.validateNoteContentType = function (e) {
        return null == e.note_content_type ||
          e.note_content_type == Ma.CONTENT_TYPE_TEXT
          ? $a.Right(m.None)
          : e.note_content_type == Ma.CONTENT_TYPE_XML
            ? $a.Right(m.Some(go.RICH_TEXT))
            : $a.Left(
                "'note_content_type' is an invalid value. Please check: " +
                  P.string(e.note_content_type),
              );
      }),
      (p.validateXmlContent = function (e) {
        var t = p.validateXml(e.note_content);
        return t.ok
          ? $a.Right(e.note_content)
          : $a.Left(
              "There is a problem with 'note_content' field. " + t.message,
            );
      }),
      (p.validateXml = function (e) {
        return e.length > p.MAX_CONTENT_BYTES
          ? {
              ok: !1,
              errorType: "limit",
              message:
                "Content length must be less than or equal to " +
                p.MAX_CONTENT_BYTES +
                " bytes, but the actual length is " +
                e.length +
                " bytes.",
            }
          : Qs.validateXml(e);
      }),
      (p.hasPathField = function (e) {
        return k.foreach(e, function (e) {
          return !p.isStringNullOrEmpty(e.path);
        });
      }),
      (p.validateUpdateNoteInput = function (e, t) {
        if (0 == Object.keys(e).length)
          return $a.Left(
            "'update' function must not be called with no update fields.",
          );
        if (null != e.note_title) {
          if ("string" != typeof e.note_title)
            return $a.Left("'note_title' must be specified as String type.");
          if (0 == e.note_title.length)
            return $a.Left("'note_title' must not be empty.");
        }
        if (null != e.note_content_type)
          return $a.Left(
            "'note_content_type' must not be specified. It cannot be changed.",
          );
        if (null != e.note_content) {
          if ("string" != typeof e.note_content)
            return $a.Left("'note_content' must be specified as String type.");
          if (0 == e.note_content.length)
            return $a.Left("'note_content' must not be empty.");
          if (t == yo.getValue(go.RICH_TEXT)) {
            t = p.validateXmlContent(e);
            if (0 === t[1]) return $a.Left(t[2]);
          }
        }
        t = null;
        if (null != e.note_attachments) {
          if (
            !(
              e.note_attachments instanceof Array &&
              null == e.note_attachments.__enum__ &&
              p.all(e.note_attachments, p.isValidAsAttachment)
            )
          )
            return $a.Left(
              "'note_attachments' is an invalid value: " +
                JSON.stringify(e.note_attachments),
            );
          t = p.toUpdateNoteAttachmentArray(e.note_attachments);
        }
        return $a.Right({
          title: e.note_title,
          content: e.note_content,
          attachments: t,
        });
      }),
      (p.isValidAsAttachment = function (e) {
        return p.isStringNullOrEmpty(e.url)
          ? !p.isStringNullOrEmpty(e.path)
          : !p.isStringNullOrEmpty(e.id) &&
              !p.isStringNullOrEmpty(e.name) &&
              null != e.content_size &&
              !p.isStringNullOrEmpty(e.content_type);
      }),
      (p.toUpdateNoteAttachmentArray = function (e) {
        var n = ja.parseInt64;
        return k.array(
          k.map(e, function (e) {
            var t;
            return p.isStringNullOrEmpty(e.url)
              ? Ea.LocalFileAttachment(e)
              : (((t = null == e ? null : JSON.parse(JSON.stringify(e))).id = n(
                  e.id,
                )),
                Ea.RemoteFileAttachment(t));
          }),
        );
      }),
      (p.validateUpdateNoteOptions = function (e) {
        var t = { notify_talk_members_of_update: !0 };
        if (null != e && null != e.notify_talk_members_of_update) {
          if ("boolean" != typeof e.notify_talk_members_of_update)
            return $a.Left(
              "'notify_talk_members_of_update' must be of Bool type.",
            );
          t.notify_talk_members_of_update = e.notify_talk_members_of_update;
        }
        return $a.Right(t);
      }),
      (p.all = function (e, t) {
        return k.foreach(e, t);
      }),
      (p.isStringNullOrEmpty = function (e) {
        return null == e || 0 == e.length;
      }),
      (p.prototype = { __class__: p }),
      ((t["albero_cli.INotes"] = tn).__name__ = ["albero_cli", "INotes"]),
      (tn.prototype = { __class__: tn }),
      function (e, t) {
        this.delegate = e;
        e = null != t ? t : { max: 24, interval: 12e4 };
        this.ratelimit = new Mn(e.max, e.interval);
      }),
    Na =
      (((t["albero_cli.LimitedNotes"] = wa).__name__ = [
        "albero_cli",
        "LimitedNotes",
      ]),
      (wa.__interfaces__ = [tn]),
      (wa.prototype = {
        limit: function (e) {
          return this.ratelimit.apply(e);
        },
        get: function (e) {
          var t = this;
          return this.limit(function () {
            return t.delegate.get(e);
          });
        },
        update: function (e, t, n) {
          var i = this;
          return this.limit(function () {
            return i.delegate.update(e, t, n);
          });
        },
        delete: function (e) {
          var t = this;
          return this.limit(function () {
            return t.delegate.delete(e);
          });
        },
        validateXml: function (e) {
          return this.delegate.validateXml(e);
        },
        on: function (e, t) {
          return this.delegate.on(e, t);
        },
        off: function (e, t) {
          return this.delegate.off(e, t);
        },
        __class__: wa,
      }),
      function (e, t, n) {
        ((this.utils = e), (this.api = t), (this.fileService = n));
      }),
    Aa =
      (((t["albero_cli.Notes"] = Na).__name__ = ["albero_cli", "Notes"]),
      (Na.__interfaces__ = [tn]),
      (Na.prototype = {
        toNoteId: function (e) {
          return new So(this.utils.parseInt64(e.id));
        },
        handleResponse: function (n, i, r) {
          return function (e, t) {
            null != e ? ((e = Ha.of(e)), i(e)) : ((e = r(t)), n(e));
          };
        },
        get: function (e) {
          var i = this,
            r = this.toNoteId(e);
          return new Promise(function (e, t) {
            var n = i.api,
              e = i.handleResponse(e, t, Wa.of);
            n.getNote(r, e);
          });
        },
        update: function (e, t, n) {
          var i = p.validateUpdateNoteInput(t, e.noteRevision.contentType);
          switch (i[1]) {
            case 0:
              return Promise.reject(Ha.createInvalidParameterError(i[2]));
            case 1:
              var r = i[2],
                o = p.validateUpdateNoteOptions(n);
              switch (o[1]) {
                case 0:
                  return Promise.reject(Ha.createInvalidParameterError(o[2]));
                case 1:
                  return this.updateNote(e, r, o[2]);
              }
          }
        },
        delete: function (e) {
          var i = this,
            r = this.toNoteId(e);
          return new Promise(function (e, t) {
            var n = i.api,
              e = i.handleResponse(e, t, Ya.of);
            n.deleteNote(r, e);
          });
        },
        updateNote: function (e, t, n) {
          var r = this,
            i = e.talkId_i64,
            o = new So(e.id_i64),
            a = this.mergeNoteObject(e, t),
            s = this.createNoteLocalEdit(i, o, a, null),
            l = n.notify_talk_members_of_update;
          return new Promise(function (e, t) {
            var n = r.api,
              i = a.currentRevision,
              e = r.handleResponse(e, t, Ka.of);
            n.updateNoteV2(o, i, s, l, e);
          });
        },
        mergeNoteObject: function (e, t) {
          e = {
            title: e.noteRevision.title,
            contentType: e.noteRevision.contentType,
            content: e.noteRevision.contentText,
            attachments: p.toUpdateNoteAttachmentArray(e.get_attachments()),
            currentRevision: e.noteRevision.revision,
          };
          return (
            null != t.title && (e.title = t.title),
            null != t.content && (e.content = t.content),
            null != t.attachments && (e.attachments = t.attachments),
            e
          );
        },
        createNoteLocalEdit: function (e, t, n, i) {
          var r = this.createNoteAttachmentFileInfoList(n.attachments),
            o = n.contentType == yo.getValue(go.RICH_TEXT);
          return new vo(e, t, n.title, n.content, r, i, null, o);
        },
        createNoteAttachmentFileInfoList: function (e) {
          for (var t = [], n = 0; n < e.length; ) {
            var i = e[n],
              r = (++n, null);
            switch (i[1]) {
              case 0:
                var o = i[2],
                  o = this.fileService.createDummyFile(o.path, o.name, o.type),
                  r = ar.createStagedFileInfoFromDummyFile(null, o, null);
                break;
              case 1:
                r = ar.createUploadedFileInfo(i[2]);
            }
            t.push(r);
          }
          return t;
        },
        validateXml: function (e) {
          return p.validateXml(e);
        },
        on: function (e, t) {
          return (this.utils.on(p.fqen(e), t), this);
        },
        off: function (e, t) {
          return (this.utils.off(p.fqen(e), t), this);
        },
        __class__: Na,
      }),
      ((t["albero_cli.entity.Domain"] = nn).__name__ = [
        "albero_cli",
        "entity",
        "Domain",
      ]),
      (nn.create = function (e, t) {
        return new nn(e, t);
      }),
      (nn.prototype = {
        toCompatibleData: function () {
          var e = JSON.parse(JSON.stringify(this)),
            t = ((e.id_i64 = this.id), this.id);
          return (
            (e.id = "_" + t.high + "_" + t.low),
            (e.talks = this.talks),
            e
          );
        },
        __class__: nn,
      }),
      function (e) {
        ((this.status = e.status),
          (this.displayName = e.displayName),
          (this.canonicalDisplayName = e.canonicalDisplayName),
          (this.phoneticDisplayName = e.phoneticDisplayName),
          (this.canonicalPhoneticDisplayName = e.canonicalPhoneticDisplayName),
          (this.profileImageUrl = e.profileImageUrl),
          (this.updatedAt = e.updatedAt),
          (this.id = e.id));
      }),
    ba =
      (((t["albero_cli.entity.User"] = Aa).__name__ = [
        "albero_cli",
        "entity",
        "User",
      ]),
      (Aa.of = function (e) {
        return new Aa(e);
      }),
      (Aa.prototype = {
        toCompatibleData: function () {
          var e = JSON.parse(JSON.stringify(this)),
            t = ((e.id_i64 = this.id), this.id);
          return (
            (e.id = "_" + t.high + "_" + t.low),
            (e.name = this.displayName),
            (e.profile_url = this.profileImageUrl),
            e
          );
        },
        __class__: Aa,
      }),
      ((t["albero_cli.entity.DomainUser"] = rn).__name__ = [
        "albero_cli",
        "entity",
        "DomainUser",
      ]),
      (rn.of = function (e) {
        return new rn(e);
      }),
      (rn.__super__ = Aa),
      (rn.prototype = e(Aa.prototype, {
        convertProfileContact: function (e) {
          return null == e
            ? null
            : e.map(function (e) {
                return ba.of(e);
              });
        },
        __class__: rn,
      })),
      ((t["albero_cli.entity.Me"] = on).__name__ = [
        "albero_cli",
        "entity",
        "Me",
      ]),
      (on.of = function (e) {
        return new on(e);
      }),
      (on.__super__ = Aa),
      (on.prototype = e(Aa.prototype, { __class__: on })),
      function (e) {
        ((this.profileItemId = e.profileItemId),
          (this.value = e.value),
          (this.canonicalValue = e.canonicalValue));
      }),
    Ma =
      (((t["albero_cli.entity.ProfileItemValue"] = ba).__name__ = [
        "albero_cli",
        "entity",
        "ProfileItemValue",
      ]),
      (ba.of = function (e) {
        return new ba(e);
      }),
      (ba.prototype = { __class__: ba }),
      ((t["albero_cli.entity.RemoteFileInfo"] = an).__name__ = [
        "albero_cli",
        "entity",
        "RemoteFileInfo",
      ]),
      (an.of = function (e) {
        return new an(e);
      }),
      (an.prototype = { __class__: an }),
      ((t["albero_cli.entity.Talk"] = sn).__name__ = [
        "albero_cli",
        "entity",
        "Talk",
      ]),
      (sn.talkType = function (e) {
        switch (e[1]) {
          case 1:
            return 1;
          case 2:
            return 2;
          default:
            return 0;
        }
      }),
      (sn.create = function (e, t, n) {
        return new sn(e, t, n);
      }),
      (sn.prototype = {
        toCompatibleData: function () {
          var e = JSON.parse(JSON.stringify(this)),
            t = ((e.id_i64 = this.id), this.id),
            t =
              ((e.id = "_" + t.high + "_" + t.low),
              (e.topic = this.name),
              (e.type = sn.talkType(this.type)),
              (e.userIds = this.userIds),
              (e.users = this.users),
              (e.domain = this.domain),
              (e.domainId_i64 = this.domainId),
              this.domainId);
          return ((e.domainId = "_" + t.high + "_" + t.low), e);
        },
        __class__: sn,
      }),
      function (e) {
        ((this.id = e.noteId.toString()), (this.id_i64 = e.noteId.value));
        var t = e.talkId;
        ((this.talkId = "_" + t.high + "_" + t.low),
          (this.talkId_i64 = e.talkId),
          (this.createdBy = e.createdBy),
          (this.createdAt = e.createdAt),
          (this.noteRevision = Ca.of(e.noteRevision)),
          Object.defineProperty(this, "title", { get: this.get_title }),
          Object.defineProperty(this, "contentType", {
            get: this.get_contentType,
          }),
          Object.defineProperty(this, "content", { get: this.get_content }),
          Object.defineProperty(this, "attachments", {
            get: this.get_attachments,
          }),
          Object.defineProperty(this, "lastUpdatedBy", {
            get: this.get_lastUpdatedBy,
          }),
          Object.defineProperty(this, "lastUpdatedAt", {
            get: this.get_lastUpdatedAt,
          }));
      }),
    Ca =
      (((t["albero_cli.entity.note.Note"] = Ma).__name__ = [
        "albero_cli",
        "entity",
        "note",
        "Note",
      ]),
      (Ma.toNoteContentTypeString = function (e) {
        return 2 == yo.createNoteContentType(e)[1]
          ? Ma.CONTENT_TYPE_XML
          : Ma.CONTENT_TYPE_TEXT;
      }),
      (Ma.of = function (e) {
        return new Ma(e);
      }),
      (Ma.prototype = {
        get_title: function () {
          return this.noteRevision.title;
        },
        get_contentType: function () {
          return Ma.toNoteContentTypeString(this.noteRevision.contentType);
        },
        get_content: function () {
          return this.noteRevision.contentText;
        },
        get_attachments: function () {
          return this.noteRevision.contentFiles;
        },
        get_lastUpdatedBy: function () {
          return this.noteRevision.createdBy;
        },
        get_lastUpdatedAt: function () {
          return this.noteRevision.createdAt;
        },
        __class__: Ma,
      }),
      ((t["albero_cli.entity.note.NotePartialUpdate"] = ln).__name__ = [
        "albero_cli",
        "entity",
        "note",
        "NotePartialUpdate",
      ]),
      (ln.of = function (e) {
        return new ln(e);
      }),
      (ln.prototype = {
        get_title: function () {
          return this.noteRevisionSummary.title;
        },
        get_contentType: function () {
          return Ma.toNoteContentTypeString(
            this.noteRevisionSummary.contentType,
          );
        },
        get_contentSummary: function () {
          return this.noteRevisionSummary.contentSummary;
        },
        get_attachments: function () {
          return this.noteRevisionSummary.contentFiles;
        },
        get_lastUpdatedBy: function () {
          return this.noteRevisionSummary.createdBy;
        },
        get_lastUpdatedAt: function () {
          return this.noteRevisionSummary.createdAt;
        },
        __class__: ln,
      }),
      function (e) {
        ((this.revision = e.revision),
          (this.title = e.title),
          (this.contentType = yo.getValue(e.contentType)),
          (this.contentText = e.contentText),
          (this.contentFiles = Ca.convertFiles(e.contentFiles)),
          (this.createdBy = e.createdBy),
          (this.createdAt = e.createdAt));
      }),
    ka =
      (((t["albero_cli.entity.note.NoteRevision"] = Ca).__name__ = [
        "albero_cli",
        "entity",
        "note",
        "NoteRevision",
      ]),
      (Ca.convertFiles = function (e) {
        return null == e || e.length <= 0
          ? []
          : e.map(function (e) {
              return an.of(e);
            });
      }),
      (Ca.of = function (e) {
        return new Ca(e);
      }),
      (Ca.prototype = { __class__: Ca }),
      function (e) {
        ((this.revision = e.revision),
          (this.title = e.title),
          (this.contentType = yo.getValue(e.contentType)),
          (this.contentSummary = e.contentSummary),
          (this.contentFiles = Ca.convertFiles(e.contentFiles)),
          (this.createdBy = e.createdBy),
          (this.createdAt = e.createdAt));
      }),
    Da =
      (((t["albero_cli.entity.note.NoteRevisionSummary"] = ka).__name__ = [
        "albero_cli",
        "entity",
        "note",
        "NoteRevisionSummary",
      ]),
      (ka.of = function (e) {
        return new ka(e);
      }),
      (ka.prototype = { __class__: ka }),
      ((t["albero_cli.entity.note.NoteSummary"] = un).__name__ = [
        "albero_cli",
        "entity",
        "note",
        "NoteSummary",
      ]),
      (un.of = function (e) {
        return new un(e);
      }),
      (un.prototype = {
        get_title: function () {
          return this.noteRevisionSummary.title;
        },
        get_contentType: function () {
          return Ma.toNoteContentTypeString(
            this.noteRevisionSummary.contentType,
          );
        },
        get_contentSummary: function () {
          return this.noteRevisionSummary.contentSummary;
        },
        get_attachments: function () {
          return this.noteRevisionSummary.contentFiles;
        },
        get_lastUpdatedBy: function () {
          return this.noteRevisionSummary.createdBy;
        },
        get_lastUpdatedAt: function () {
          return this.noteRevisionSummary.createdAt;
        },
        __class__: un,
      }),
      ((t["puremvc.interfaces.IMediator"] = cn).__name__ = [
        "puremvc",
        "interfaces",
        "IMediator",
      ]),
      (cn.prototype = { __class__: cn }),
      function (e, t) {
        (Bn.call(this),
          (this.mediatorName = null != e ? e : Da.NAME),
          null != t && (this.viewComponent = t));
      }),
    Ra =
      (((t["puremvc.patterns.mediator.Mediator"] = Da).__name__ = [
        "puremvc",
        "patterns",
        "mediator",
        "Mediator",
      ]),
      (Da.__interfaces__ = [cn]),
      (Da.__super__ = Bn),
      (Da.prototype = e(Bn.prototype, {
        getMediatorName: function () {
          return this.mediatorName;
        },
        setViewComponent: function (e) {
          this.viewComponent = e;
        },
        getViewComponent: function () {
          return this.viewComponent;
        },
        listNotificationInterests: function () {
          return [];
        },
        handleNotification: function (e) {},
        onRegister: function () {},
        onRemove: function () {},
        __class__: Da,
      })),
      function () {
        Da.call(this, "commandline", null);
      }),
    Oa =
      (((t["albero_cli.mediator.CommandLineMediator"] = Ra).__name__ = [
        "albero_cli",
        "mediator",
        "CommandLineMediator",
      ]),
      (Ra.__super__ = Da),
      (Ra.prototype = e(Da.prototype, {
        onRegister: function () {
          this.getViewComponent();
          ((this.eventEmitter = vn.getInstance()),
            (this.dataRecovered = !1),
            this.startDataSaver());
        },
        listNotificationInterests: function () {
          return [
            "app_state_changed",
            "access_token_changed",
            "current_user_changed",
            "configuration_changed",
            "domain_selection_changed",
            "talk_selection_changed",
            "talk_list_scroll_to_talk_top_needed",
            "user_selection_needed",
            "user_selection_changed",
            "friend_selection_needed",
            "common_stamp_set_loaded",
            "stamp_selection_started",
            "stamp_selection_ended",
            "stamp_selection_changed",
            "action_selection_changed",
            "current_page_changed",
            "current_page_reassgined",
            "fileinfo_selection_changed",
            "message_fileinfo_selection_changed",
            "staged_fileinfo_selection_changed",
            "staged_fileinfos_added",
            "staged_fileinfos_moved",
            "error_occurred",
            "brand_badge_changed",
            "send_form_top_changed",
            "solutions_loaded",
            "right_pane_opened",
            "right_pane_closed",
            "password_expiration_overed",
            "password_expiration_warned",
            "photo_editor_saved",
            "mc_authenticated_user_received",
            "keyword_watching_updated",
            "keyword_detaction_updated",
            "department_selection_changed_for_members_page",
            "presences_updated",
            "noteinfo_selection_changed",
            "start_note_editing",
            "start_notification_failed",
            "signed_out",
            "data_recovering",
            "data_recovered",
            "notify_update_user",
            "notify_add_friend",
            "notify_add_acquaintance",
            "notify_add_acquaintances",
            "notify_delete_friend",
            "notify_delete_acquaintance",
            "notify_delete_acquaintances",
            "notify_update_domain_users",
            "get_domain_users_responsed",
            "get_users_responsed",
            "get_profile_responsed",
            "get_profile_errored",
            "update_user_responsed",
            "update_user_errored",
            "update_profile_responsed",
            "update_profile_errored",
            "notify_update_department_tree",
            "notify_update_department_users",
            "get_department_tree_responsed",
            "get_department_tree_canceled",
            "get_department_users_responsed",
            "get_department_users_canceled",
            "get_department_user_count_responsed",
            "get_me_responsed",
            "department_user_count_cleared",
            "department_user_ids_prepared",
            "notify_add_domain_invite",
            "notify_accept_domain_invite",
            "notify_delete_domain_invite",
            "notify_join_domain",
            "notify_update_domain",
            "notify_leave_domain",
            "notify_add_domain_members",
            "notify_create_pair_talk",
            "notify_create_group_talk",
            "notify_update_group_talk",
            "notify_add_talkers",
            "notify_add_talkers_including_me",
            "notify_delete_talker",
            "notify_delete_talk",
            "notify_create_message",
            "notify_delete_message",
            "notify_update_read_statuses",
            "notify_update_talk_status",
            "notify_update_local_talk_status",
            "notify_get_messages",
            "notify_get_message_status",
            "notify_add_favorite_talk",
            "notify_delete_favorite_talk",
            "notify_disable_push_notification",
            "notify_enable_push_notification",
            "create_message_start",
            "create_message_complete",
            "create_message_fail",
            "notify_create_announcement",
            "notify_delete_announcement",
            "notify_update_announcement_status",
            "notify_get_announcements",
            "create_announcement_start",
            "create_announcement_complete",
            "create_announcement_fail",
            "notify_update_question",
            "get_questions_responsed",
            "notify_create_attachment",
            "notify_delete_attachment",
            "get_note_statuses_loaded",
            "get_note_loaded",
            "get_note_failed_by_note_not_found",
            "create_note_completed",
            "create_note_failed",
            "update_note_local_edit",
            "clear_note_local_edit",
            "update_note_setting_completed",
            "update_note_setting_failed_by_not_found",
            "update_note_setting_failed_by_conflict",
            "update_note_setting_failed_by_editing",
            "update_note_completed",
            "delete_note_completed",
            "notify_create_note",
            "notify_update_note_partially",
            "notify_delete_note",
            "notify_update_note_for_setting",
            "notify_update_note_for_revision",
            "create_note_button_clicked",
            "notify_lock_note",
            "notify_unlock_note",
            "get_file_responsed",
            "notify_search_messages",
            "notify_search_attachments",
            "notify_search_messages_fail",
            "notify_search_attachments_fail",
            "notify_search_prepare",
            "notify_search_clear",
            "notify_filter_box_text_changed",
            "notify_search_box_popup",
            "notify_add_account_control_request",
            "notify_delete_account_control_request",
            "notify_join_account_control_group",
            "notify_update_account_control_group_partially",
            "notify_leave_account_control_group",
            "prepare_conference_from_message",
            "join_conference_responsed",
            "join_conference_canceled",
            "notify_open_conference",
            "notify_close_conference",
            "notify_conference_participant_join",
            "notify_conference_participant_limit",
            "notify_conference_participant_reject",
            "notify_all_talk_members_rejected_conference",
          ];
        },
        handleNotification: function (e) {
          var t = this,
            n = e.getName();
          switch (n) {
            case "current_page_changed":
              var i = e.getBody();
              (x.enumEq(i, g.error) && (this.dataRecovered = !1),
                this.eventEmitter.emit(e.getName(), e.getBody()));
              break;
            case "data_recovered":
              if (this.dataRecovered) return;
              ((this.dataRecovered = !0), this.eventEmitter.emit(e.getName()));
              break;
            case "data_recovering":
              ((this.dataRecovered = !1), this.eventEmitter.emit(e.getName()));
              break;
            case "error_occurred":
              this.eventEmitter.emit(
                e.getName(),
                new Error("AdapterError"),
                e.getBody(),
              );
              break;
            case "signed_out":
              (F.e("signout"), process.exit(1));
              break;
            case "start_notification_failed":
              (F.e("start_notification failed."), process.exit(1));
              break;
            default:
              this.eventEmitter.emit(e.getName(), e.getBody());
          }
          var r, o, a;
          this.getViewComponent();
          switch (e.getName()) {
            case "access_token_changed":
              (this.dataStore.clear(!0), F.i(e.getBody()), process.exit(0));
              break;
            case "create_announcement_complete":
            case "create_message_complete":
              var s = e.getBody(),
                s = s[0],
                l = this.dataStore,
                u = s.userId;
              null != l.me &&
                null != (l = l.me.id) &&
                null != u &&
                l.high == u.high &&
                l.low == u.low &&
                (this.messageEvent.messageCreated(s, e.getType()),
                this.sendQueue.sendNext(s));
              break;
            case "create_announcement_fail":
            case "create_message_fail":
              var l = e.getBody()[0],
                u = e.getBody()[1],
                s = this.dataStore,
                c = u.userId;
              null != s.me &&
                null != (s = s.me.id) &&
                null != c &&
                s.high == c.high &&
                s.low == c.low &&
                (this.messageEvent.messageCreated(null, e.getType()),
                (null != l && 429 == l.code) || (null != l && 503 == l.code)
                  ? this.sendQueue.retryAfter(1e3 * (0 | l.detail.retry_after))
                  : this.sendQueue.sendNext(u));
              break;
            case "create_note_completed":
              ((s = e.getBody()), (c = this.dataStore), (l = s.note.createdBy));
              null != c.me &&
                null != (u = c.me.id) &&
                null != l &&
                u.high == l.high &&
                u.low == l.low &&
                (this.messageEvent.noteCreated(s.note, s.emitterKey),
                this.sendQueue.sendNextTalkMessage());
              break;
            case "create_note_failed":
              ((c = e.getBody()), (u = this.dataStore), (l = c.callerId));
              null != u.me &&
                null != (s = u.me.id) &&
                null != l &&
                s.high == l.high &&
                s.low == l.low &&
                (this.messageEvent.noteCreated(null, c.emitterKey),
                this.sendQueue.sendNextTalkMessage(),
                F.d("failed to create note: " + P.string(c.error)));
              break;
            case "current_user_changed":
              e.getBody();
              this.sendQueue.restart();
              break;
            case "notify_add_domain_invite":
              e.getBody();
              break;
            case "notify_create_group_talk":
            case "notify_create_pair_talk":
              this.dataRecovered &&
                ((r = e.getBody()),
                Xa.delay(function () {
                  t.emitWithMe(r, "JoinMessage", t.dataStore.me);
                }, 500));
              break;
            case "notify_create_message":
              var _ = e.getBody(),
                u = this.dataStore,
                s = _.userId;
              (null != u.me &&
                null != (l = u.me.id) &&
                null != s &&
                l.high == s.high &&
                l.low == s.low) ||
                (_.isBotMessage && !qs.talkWithBot) ||
                ((c = this.dataStore.getTalkStatus(_.talkId)),
                (u = Vt.getAcceptableEventTimeDiff()),
                null != c && !c.isMessageAcceptable(_, u)) ||
                Xa.delay(function () {
                  null != t.dataStore.getTalk(_.talkId) &&
                    (t.sendNotification("Read", ui.TALK(_.talkId)),
                    t.dispatch(_));
                }, 200);
              break;
            case "notify_create_note":
              l = e.getBody();
              this.eventEmitter.emit(
                p.fqen("created"),
                this.hubotObject.noteSummaryObject(l),
              );
              break;
            case "notify_delete_note":
              s = e.getBody();
              this.eventEmitter.emit(
                p.fqen("deleted"),
                this.hubotObject.noteIdObject(s),
              );
              break;
            case "notify_delete_talk":
              c = e.getBody();
              this.messageEvent.deleteTalk(c);
              break;
            case "notify_get_message_status":
              u = e.getBody();
              this.messageEvent.messageRead(
                u.talkId,
                u.id,
                u.readUserIds,
                u.unreadUserIds,
              );
              break;
            case "notify_leave_domain":
              l = e.getBody();
              this.messageEvent.leaveDomain(l);
              break;
            case "notify_update_group_talk":
              this.dataRecovered &&
                ((o = e.getBody()),
                (a = null != o.name ? o.name : ""),
                Xa.delay(function () {
                  t.emitWithMe(o, "TopicChangeMessage", t.dataStore.me, a);
                }, 500));
              break;
            case "notify_update_local_talk_status":
              s = e.getBody();
              this.messageEvent.messageReadEveryone(
                s.id,
                s.maxEveryoneReadMessageId,
              );
              break;
            case "notify_update_note_partially":
              c = e.getBody();
              this.eventEmitter.emit(
                p.fqen("updated"),
                this.hubotObject.notePartialUpdateObject(c),
              );
              break;
            case "notify_update_read_statuses":
              for (
                var h = e.getBody(), d = 0, f = h.messageIds;
                d < f.length;
              ) {
                var m = f[d];
                (++d,
                  this.messageEvent.messageRead(h.talkId, m, h.readUserIds));
              }
          }
        },
        emitWithUser: function (e, t, n, i) {
          this.emit(e, t, this.hubotObject.userObject(n), i);
        },
        emitWithMe: function (e, t, n, i) {
          this.emit(e, t, this.hubotObject.meObject(n), i);
        },
        emit: function (i, r, o, a) {
          var s = this;
          null != r &&
            null != i &&
            null != o &&
            this.api.getAllUserIdentifiers(function () {
              var e = s.eventEmitter,
                t = i.id,
                t = "_" + t.high + "_" + t.low,
                n = s.hubotObject.talkObjects();
              e.emit(r, { room: t, rooms: n }, o, a);
            });
        },
        leave: function (e) {
          var t = this;
          Xa.delay(function () {
            null != t.dataStore.getTalk(e.id) &&
              t.sendNotification("Talk", Ii.DELETE_FOR_HUBOT(e.id));
          }, 500);
        },
        dispatch: function (o) {
          function e(e, t, n, i) {
            var t = a.dataStore.getUser(s.domainId, t),
              r = o.id;
            a.emitWithUser(s, e, t, {
              id: "_" + r.high + "_" + r.low,
              content: n,
              mention: i,
            });
          }
          var a = this,
            t = o.content,
            s = this.dataStore.getTalk(o.talkId);
          if (o.type == d.system) {
            var n = t.type;
            if ("add_talkers" == n) {
              var i = t.added_user_ids,
                r = this.dataStore.me.id;
              if (D.contains(i, r)) e("JoinMessage", r);
              else
                for (var l = 0; l < i.length; ) {
                  var u = i[l];
                  (++l, e("EnterMessage", u));
                }
            } else
              "delete_talker" == n
                ? (e("LeaveMessage", t.deleted_user_id),
                  t.user_ids.length <= 1 && this.leave(s))
                : "hide_pair_talk" == n &&
                  (e("LeaveMessage", t.user_id), this.leave(s));
          } else {
            ((r = null != qs.name ? qs.name : "Hubot"),
              (n = Oa.convertNonSystemMessageToHubotTextMessage(o, s, r)));
            null != n &&
              e("TextMessage", n.sender, n.text, {
                all: n.mentionAll,
                me: n.mentionMe,
                data: n.mentions,
              });
          }
        },
        startDataSaver: function () {
          var e = this;
          Xa.delay(function () {
            (e.dataRecovered && e.dataStore.saveIfNeeded(), e.startDataSaver());
          }, 5e3);
        },
        __class__: Ra,
      })),
      function () {}),
    Fa =
      (((t["albero_cli.mediator.MessageUtil"] = Oa).__name__ = [
        "albero_cli",
        "mediator",
        "MessageUtil",
      ]),
      (Oa.convertNonSystemMessageToHubotTextMessage = function (e, t, n) {
        var i = [],
          r = null;
        switch (e.type[1]) {
          case 0:
            r = null;
            break;
          case 1:
            ((i = Oa.extractMentionParts(e.content, e.type)),
              (r = An.removeMentionMarkup(e.content)));
            break;
          default:
            for (
              var o = e.content,
                a = null == o ? null : JSON.parse(JSON.stringify(o)),
                s = 0,
                l = L.fields(a);
              s < l.length;
            ) {
              var u,
                c = l[s],
                _ = (++s, L.field(a, c));
              L.isObject(a) && null != _.high && null != _.low
                ? ((u = new E(_.high, _.low)),
                  (a[c] =
                    "stamp_index" == c
                      ? Va.toString(u)
                      : "_" + u.high + "_" + u.low))
                : "stamp_set" == c && (a[c] = P.string(_));
            }
            (e.type != d.textMultipleFile ||
              (null != a.text && 0 != a.text.length) ||
              null == a.files ||
              1 != a.files.length ||
              (a = a.files[0]),
              (i = Oa.extractMentionParts(a, e.type)),
              Oa.removeMentionMarkup(a, e.type),
              (r = JSON.stringify(a)));
        }
        return null != r
          ? (t.type != Xr.PairTalk || I.startsWith(r, n) || (r = n + " " + r),
            Fa.create(e.userId, r, i, e.mention))
          : null;
      }),
      (Oa.extractMentionParts = function (e, t) {
        switch (t[1]) {
          case 1:
            return An.extractMentionParts(e);
          case 2:
          case 5:
          case 12:
            if (e.text) return An.extractMentionParts(e.text);
            break;
          case 17:
          case 18:
          case 21:
            if (e.title) return An.extractMentionParts(e.title);
            break;
          case 13:
          case 14:
          case 15:
          case 16:
          case 19:
          case 20:
            if (e.question) return An.extractMentionParts(e.question);
        }
        return [];
      }),
      (Oa.removeMentionMarkup = function (e, t) {
        switch (t[1]) {
          case 2:
          case 5:
          case 12:
            e.text && (e.text = An.removeMentionMarkup(e.text));
            break;
          case 17:
          case 18:
          case 21:
            e.title && (e.title = An.removeMentionMarkup(e.title));
            break;
          case 13:
          case 14:
          case 15:
          case 16:
          case 19:
          case 20:
            e.question && (e.question = An.removeMentionMarkup(e.question));
        }
      }),
      function (e, t, n, i, r) {
        ((this.sender = e),
          (this.text = t),
          (this.mentionAll = n),
          (this.mentionMe = i),
          (this.mentions = r));
      }),
    xa =
      (((t["albero_cli.mediator.HubotTextMessage"] = Fa).__name__ = [
        "albero_cli",
        "mediator",
        "HubotTextMessage",
      ]),
      (Fa.create = function (e, t, n, i) {
        var r =
            0 <
            n.filter(function (e) {
              return e.isALL();
            }).length,
          n = n
            .filter(function (e) {
              return !e.isALL();
            })
            .map(function (e) {
              return { user: { id: Fa.toIdStr(e.userIdNumStr) } };
            });
        return new Fa(e, t, r, i, n);
      }),
      (Fa.toIdStr = function (e) {
        e = D.parse(e);
        return "_" + e.high + "_" + e.low;
      }),
      (Fa.prototype = { __class__: Fa }),
      function () {
        l.call(this, "hubotObject");
      }),
    Ua =
      (((t["albero_cli.proxy.HubotObjectProxy"] = xa).__name__ = [
        "albero_cli",
        "proxy",
        "HubotObjectProxy",
      ]),
      (xa.__super__ = l),
      (xa.prototype = e(l.prototype, {
        userObject: function (e) {
          return rn.of(e).toCompatibleData();
        },
        meObject: function (e) {
          return on.of(e).toCompatibleData();
        },
        userObjectByIdStr: function (e, t) {
          ((t = t.split("_")),
            (e = this.userObjectsByIds(e, [
              2 < t.length ? _.make(P.parseInt(t[1]), P.parseInt(t[2])) : null,
            ])));
          return 0 < e.length ? e[0] : null;
        },
        userObjectsByIds: function (e, t, n) {
          for (
            var i = [], r = 0, o = this.dataStore.getUsers(e, t);
            r < o.length;
          ) {
            var a,
              s,
              l = o[r];
            (++r,
              null != l &&
                ((a = "u_" + (a = l.id).high + "_" + a.low),
                null ==
                  (s =
                    null != n
                      ? null != Y[a]
                        ? n.getReserved(a)
                        : n.h[a]
                      : null) &&
                  ((s = this.userObject(l)), null != n) &&
                  ((l = s), null != Y[a] ? n.setReserved(a, l) : (n.h[a] = l)),
                null != s) &&
                i.push(s));
          }
          return i;
        },
        userObjects: function (e) {
          for (
            var t = {}, n = new B(), i = 0, r = this.dataStore.getUsers(e);
            i < r.length;
          ) {
            var o,
              a,
              s = r[i],
              l = (++i, s.id),
              l = "_" + l.high + "_" + l.low,
              u =
                null == t[l] ||
                (0 ==
                  (a =
                    ((u = (null != Y[l] ? n.getReserved(l) : n.h[l]).updatedAt)
                      .high -
                      (o = s.updatedAt).high) |
                    0) && (a = w.ucompare(u.low, o.low)),
                (u.high < 0 ? (o.high < 0 ? a : -1) : 0 <= o.high ? a : 1) <=
                  0);
            u &&
              ((t[l] = this.userObject(s)),
              null != Y[l] ? n.setReserved(l, s) : (n.h[l] = s));
          }
          return t;
        },
        talkObject: function (e, t) {
          var n = this.userObjectsByIds(e.domainId, e.userIds, t),
            t = this.domainObjectById(e.domainId, t);
          return sn.create(e, n, t).toCompatibleData();
        },
        talkObjects: function (e) {
          for (
            var t = new B(), n = {}, i = 0, r = this.dataStore.getTalks();
            i < r.length;
          ) {
            var o,
              a = r[i];
            (++i,
              (null != e &&
                ((o = a.domainId),
                !(
                  null != e &&
                  null != o &&
                  e.high == o.high &&
                  e.low == o.low
                ))) ||
                (n["_" + (o = a.id).high + "_" + o.low] = this.talkObject(
                  a,
                  t,
                )));
          }
          return n;
        },
        domainObjectById: function (e, t) {
          var n, i;
          return null == e
            ? null
            : ((n = "d_" + e.high + "_" + e.low),
              null == (i = null != Y[n] ? t.getReserved(n) : t.h[n]) &&
                ((e = i = this.domainObject(this.dataStore.getDomain(e))),
                null != Y[n] ? t.setReserved(n, e) : (t.h[n] = e)),
              i);
        },
        domainObject: function (e, t) {
          var n = null;
          return (
            (t = null == t ? !1 : t) && (n = this.talkObjects(e.id)),
            nn.create(e, n).toCompatibleData()
          );
        },
        domainObjects: function () {
          for (
            var e = {}, t = 0, n = this.dataStore.getDomains();
            t < n.length;
          ) {
            var i = n[t],
              r = (++t, i.id);
            e["_" + r.high + "_" + r.low] = this.domainObject(i);
          }
          return e;
        },
        noteObject: function (e) {
          return Ma.of(e);
        },
        notePartialUpdateObject: function (e) {
          return ln.of(e);
        },
        noteSummaryObject: function (e) {
          return un.of(e);
        },
        noteIdObject: function (e) {
          return { id: e.toString(), id_i64: e.value };
        },
        messageObject: function (e) {
          var t = new B(),
            n = this.dataStore.getTalk(e.talkId),
            i = this.userObjectsByIds(n.domainId, [e.userId], t),
            r = {},
            o = ((r.type = O.enumIndex(e.type)), e.content),
            o =
              ((r.content = null == o ? null : JSON.parse(JSON.stringify(o))),
              e.createdAt),
            o =
              ((r.createdAt = null == o ? null : JSON.parse(JSON.stringify(o))),
              e.id),
            o =
              ((r.id_i64 = null == o ? null : JSON.parse(JSON.stringify(o))),
              e.id),
            o =
              ((r.id = "_" + o.high + "_" + o.low),
              (r.user = 0 < i.length ? i[0] : null),
              e.userId),
            i =
              ((r.userId = "_" + o.high + "_" + o.low),
              (r.talk = this.talkObject(n, t)),
              e.talkId);
          return ((r.talkId = "_" + i.high + "_" + i.low), r);
        },
        __class__: xa,
      })),
      function () {
        (l.call(this, "messageEvent"), (this.emitters = new B()));
      }),
    La =
      (((t["albero_cli.proxy.MessageEventProxy"] = Ua).__name__ = [
        "albero_cli",
        "proxy",
        "MessageEventProxy",
      ]),
      (Ua.__super__ = l),
      (Ua.prototype = e(l.prototype, {
        registEmitter: function (e, t) {
          La.factory = this;
          var n,
            t = La.createInstance(t);
          null != t &&
            ((n = this.emitters),
            null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t));
        },
        noteCreated: function (e, t) {
          var n,
            i = this.emitters,
            i = null != Y[t] ? i.getReserved(t) : i.h[t];
          null == i ||
            (this.emitters.remove(t), null == e) ||
            ((t = e.noteId.toString()),
            (n = this.emitters),
            null != Y[t] ? n.setReserved(t, i) : (n.h[t] = i),
            (i.talk = this.dataStore.getTalk(e.talkId)),
            i.setFirstReader(this.dataStore.me),
            i.onceOnSend(e),
            i.hasEventListener()) ||
            this.emitters.remove(t);
        },
        messageCreated: function (e, t) {
          var n,
            i = this.emitters,
            i = null != Y[t] ? i.getReserved(t) : i.h[t];
          null == i ||
            (this.emitters.remove(t), null == e) ||
            ((t = "_" + (t = e.id).high + "_" + t.low),
            (n = this.emitters),
            null != Y[t] ? n.setReserved(t, i) : (n.h[t] = i),
            (i.talk = this.dataStore.getTalk(e.talkId)),
            i.setFirstReader(this.dataStore.me),
            i.emitOnSend(e),
            i.removeEventListener("send"),
            i.hasEventListener()) ||
            ((n = e.id), this.emitters.remove("_" + n.high + "_" + n.low));
        },
        messageRead: function (t, n, e, i) {
          var r = this,
            o = "_" + n.high + "_" + n.low,
            a = this.emitters,
            a = null != Y[o] ? a.getReserved(o) : a.h[o];
          null != a &&
            (a.emitOnRead(e, i),
            (o = 16 <= a.__readerIds.length),
            a.startTimer(o, function (e) {
              e
                ? r.messageReadEveryone(t, n)
                : r.sendNotification("Read", ui.READ_STATUS(t, n));
            }),
            0 == a.__unreadIds.length) &&
            this.messageReadEveryone(t, n);
        },
        messageReadEveryone: function (e, t) {
          if (null != t)
            for (var n = this.emitters.keys(); n.hasNext(); ) {
              var i,
                r = n.next(),
                o = this.emitters,
                o = null != Y[r] ? o.getReserved(r) : o.h[r];
              null != o.talk &&
                null != (o = o.talk.id) &&
                null != e &&
                o.high == e.high &&
                o.low == e.low &&
                (0 == (i = ((o = D.makeFromIdStr(r)).high - t.high) | 0) &&
                  (i = w.ucompare(o.low, t.low)),
                (o.high < 0 ? (t.high < 0 ? i : -1) : 0 <= t.high ? i : 1) <=
                  0) &&
                this._messageReadEveryone(e, r);
            }
        },
        _messageReadEveryone: function (e, t) {
          var n = this.emitters,
            n = null != Y[t] ? n.getReserved(t) : n.h[t];
          null == n ||
            (0 < n.__unreadIds.length && n.emitOnRead(n.__unreadIds),
            n.stopTimer(),
            n.removeEventListener("read"),
            n.hasEventListener()) ||
            this.emitters.remove(t);
        },
        deleteTalk: function (e) {
          for (var t = this.emitters.keys(); t.hasNext(); ) {
            var n,
              i = t.next(),
              r = this.emitters,
              r = null != Y[i] ? r.getReserved(i) : r.h[i];
            null != r.talk &&
              null != (n = r.talk.id) &&
              null != e &&
              n.high == e.high &&
              n.low == e.low &&
              (r.stopTimer(), this.emitters.remove(i));
          }
        },
        leaveDomain: function (e) {
          for (var t = this.emitters.keys(); t.hasNext(); ) {
            var n,
              i = t.next(),
              r = this.emitters,
              r = null != Y[i] ? r.getReserved(i) : r.h[i];
            null != r.talk &&
              null != (n = r.talk.domainId) &&
              null != e &&
              n.high == e.high &&
              n.low == e.low &&
              (r.stopTimer(), this.emitters.remove(i));
          }
        },
        __class__: Ua,
      })),
      function (e) {
        ((this.context = e),
          (this.listeners = new B()),
          Object.defineProperty(this, "readUsers", { get: this.get_readUsers }),
          Object.defineProperty(this, "unreadUsers", {
            get: this.get_unreadUsers,
          }));
      }),
    Pa =
      (((t["albero_cli.proxy.Emitter"] = La).__name__ = [
        "albero_cli",
        "proxy",
        "Emitter",
      ]),
      (La.createInstance = function (e) {
        for (var t = null, n = L.fields(e), i = 0; i < n.length; ) {
          var r,
            o = n[i];
          (++i,
            I.startsWith(o, "on") &&
              ((r = L.field(e, o)), L.isFunction(r)) &&
              (L.deleteField(e, o),
              (t = null == t ? new La(e) : t).addEventListener(
                U.substr(o, 2, null),
                r,
              )));
        }
        return t;
      }),
      (La.prototype = {
        addEventListener: function (e, t) {
          var n = this.listeners;
          null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
        },
        hasEventListener: function () {
          var e = this.listeners;
          return new is(e, e.arrayKeys()).hasNext();
        },
        removeEventListener: function (e) {
          this.listeners.remove(e);
        },
        emit: function (e, t) {
          var n = this.listeners,
            n = null != Y[e] ? n.getReserved(e) : n.h[e];
          null != n && n.apply(this, t);
        },
        onceOnSend: function (e) {
          ((this.note = this.createNoteObjectForHubot(e)),
            this.emit("send", [this]),
            this.removeEventListener("send"));
        },
        emitOnSend: function (e) {
          ((this.message = this.get_message(e)), this.emit("send", [this, e]));
        },
        emitOnRead: function (e, t) {
          for (
            var n = [], i = (null != t && (this.__unreadIds = t), 0);
            i < e.length;
          ) {
            var r = e[i];
            (++i,
              D.contains(this.__readerIds, r) ||
                (this.__readerIds.push(r), n.push(r)),
              D.remove(this.__unreadIds, r));
          }
          0 < n.length &&
            this.emit("read", [
              this.users(n),
              this.get_readUsers(),
              this.get_unreadUsers(),
            ]);
        },
        setFirstReader: function (e) {
          ((this.__readerIds = [e.id]),
            (this.__unreadIds = this.talk.userIds.slice()),
            D.remove(this.__unreadIds, e.id));
        },
        get_readUsers: function () {
          return this.users(this.__readerIds);
        },
        get_unreadUsers: function () {
          return this.users(this.__unreadIds);
        },
        users: function (e) {
          return null == e
            ? []
            : La.factory.hubotObject.userObjectsByIds(this.talk.domainId, e);
        },
        createNoteObjectForHubot: function (e) {
          return La.factory.hubotObject.noteObject(e);
        },
        get_message: function (e) {
          return La.factory.hubotObject.messageObject(e);
        },
        answer: function (r) {
          var o = this,
            a = O.typeOf(this.message.type),
            e = this.message.id.split("_"),
            t =
              2 < e.length ? _.make(P.parseInt(e[1]), P.parseInt(e[2])) : null;
          switch (a[1]) {
            case 14:
            case 16:
            case 18:
            case 19:
            case 20:
            case 21:
              var n = this.message.content.in_reply_to,
                t = new E(n.high, n.low);
          }
          e = _i.Question(t, function (e) {
            var t = e.responses;
            switch (a[1]) {
              case 13:
              case 14:
              case 19:
                var n = o.users(t[0].userIds),
                  i = o.users(t[1].userIds);
                r(n, i);
                break;
              case 17:
              case 18:
              case 21:
                ((n = o.users(t[0].userIds)), (i = o.users(t[1].userIds)));
                r(n, i);
                break;
              case 15:
              case 16:
              case 20:
                n = t.map(function (e) {
                  return o.users(e.userIds);
                });
                r(n);
            }
          });
          La.factory.sendNotification("ReloadData", e);
        },
        startTimer: function (e, t) {
          var n = this,
            i =
              (null != this.__readerTimer && this.stopTimer(),
              (this.__readerTimerWait = e ? 1 : 1440),
              null),
            i = function () {
              ((n.__readerTimerWait *= 2),
                n.__readerTimerWait < 1440
                  ? (t(!1),
                    (n.__readerTimer = Xa.delay(
                      i,
                      (60 * n.__readerTimerWait * 1e3) | 0,
                    )))
                  : (t(!0), (n.__readerTimer = null)));
            };
          this.__readerTimer = Xa.delay(
            i,
            (60 * this.__readerTimerWait * 1e3) | 0,
          );
        },
        stopTimer: function () {
          null != this.__readerTimer &&
            (this.__readerTimer.stop(), (this.__readerTimer = null));
        },
        __class__: La,
      }),
      function () {
        (l.call(this, "sendQueue"),
          (this.sendCount = 0),
          (this.sendMsgQueue = new Ba(this, 550, new Ro("sendQueueMessage"))),
          (this.sendAnnounceQueue = new Ba(
            this,
            5050,
            new Ro("sendQueueAnnounce"),
          )));
      }),
    Ba =
      (((t["albero_cli.proxy.SendQueueProxy"] = Pa).__name__ = [
        "albero_cli",
        "proxy",
        "SendQueueProxy",
      ]),
      (Pa.__super__ = l),
      (Pa.prototype = e(l.prototype, {
        restart: function () {
          (this.sendMsgQueue.sendNext(), this.sendAnnounceQueue.sendNext());
        },
        retryAfter: function (e) {
          (this.sendMsgQueue.retryAfter(e),
            this.sendAnnounceQueue.retryAfter(e));
        },
        sendNextTalkMessage: function () {
          this.sendMsgQueue.sendNext();
        },
        sendNext: function (e) {
          (null != e.talkId
            ? this.sendMsgQueue
            : this.sendAnnounceQueue
          ).sendNext();
        },
        sendAnnouncement: function (e, t) {
          this.send(e, null, t);
        },
        sendMessage: function (e, t) {
          this.send(null, e, t);
        },
        send: function (e, t, n) {
          if ("string" != typeof n) {
            var i = new E(0, this.sendCount++),
              r =
                (this.messageEvent.registEmitter("_" + i.high + "_" + i.low, n),
                new O());
            ((r.id = i),
              (r.domainId = e),
              (r.talkId = t),
              (r.content = this.parseContent(n)),
              (r.type = this.detectType(r.content)),
              r.type != d.unknown &&
                (r.type == d.file && null != r.content.path
                  ? this.sendFile(t, r.content, "_" + i.high + "_" + i.low)
                  : r.type == d.noteCreated
                    ? this.sendNote(t, r.content, "_" + i.high + "_" + i.low)
                    : ((r.type != d.yesOrNo && r.type != d.selectOne) ||
                        null != r.content.listing ||
                        (r.content.listing = !1),
                      this.applyLimitToMessage(r),
                      this.pushQueue(r))));
          } else
            for (var o = 0, a = N.slice(n, 1024); o < a.length; ) {
              var s = a[o],
                l = (++o, new O());
              ((l.domainId = e),
                (l.talkId = t),
                (l.type = d.text),
                (l.content = s),
                this.pushQueue(l));
            }
        },
        sendNote: function (e, t, n) {
          var i = p.validateCreateNoteInput(t);
          switch (i[1]) {
            case 0:
              F.e("ValidationError: " + i[2]);
              break;
            case 1:
              this.sendMsgQueue.pushQueue(
                "Note",
                si.CREATE_NOTE_FOR_HUBOT(e, i[2], !0, n),
              );
          }
        },
        sendFile: function (e, t, n) {
          var i,
            r = null,
            o = null;
          if ("string" == typeof t) i = t;
          else {
            if (t.path instanceof Array && null == t.path.__enum__)
              return void this.sendFiles(e, t.path, t.name, t.type, t.text, n);
            if (null != t.text)
              return void this.sendFiles(
                e,
                [t.path],
                [t.name],
                [t.type],
                t.text,
                n,
              );
            ((i = t.path), (r = t.name), (o = t.type));
          }
          null != i && hs.existsSync(i)
            ? this.sendMsgQueue.pushQueue(
                "FilePath",
                Xn.UPLOAD_PATH(e, i, r, o, n),
              )
            : F.e("File not found: ", i);
        },
        sendFiles: function (e, t, n, i, r, o) {
          for (var a = 0; a < t.length; ) {
            var s = t[a];
            if ((++a, !hs.existsSync(s)))
              return void F.e("File not found: ", s);
          }
          var l = t.length;
          if (null == n) {
            for (var u = [], c = 0, _ = l; c < _; ) {
              c++;
              u.push(null);
            }
            n = u;
          }
          if (null == i) {
            for (var h = [], d = 0, f = l; d < f; ) {
              d++;
              h.push(null);
            }
            i = h;
          }
          n.length != l || i.length != l
            ? F.e("The length of name or type must be the same as path length.")
            : this.sendMsgQueue.pushQueue(
                "FilePath",
                Xn.UPLOAD_MULTI_PATH(e, r, t, n, i, o),
              );
        },
        pushQueue: function (e) {
          (null == e.talkId
            ? this.sendAnnounceQueue
            : this.sendMsgQueue
          ).pushQueue("Send", e);
        },
        parseContent: function (e) {
          if (null == e) return null;
          for (var t = L.fields(e), n = 0; n < t.length; ) {
            var i,
              r,
              o = t[n],
              a = (++n, L.field(e, o));
            "string" == typeof a
              ? ((r = null),
                "stamp_set" == o
                  ? (r = P.parseInt(a))
                  : "stamp_index" == o
                    ? (r = D.parse(a))
                    : "file_id" == o || "in_reply_to" == o
                      ? (r =
                          2 < (i = a.split("_")).length
                            ? _.make(P.parseInt(i[1]), P.parseInt(i[2]))
                            : null)
                      : "close_yesno" == o ||
                          "close_select" == o ||
                          "close_task" == o
                        ? ((e.close = U.substr(o, 6, null)),
                          (o = "in_reply_to"),
                          (r =
                            2 < (i = a.split("_")).length
                              ? _.make(P.parseInt(i[1]), P.parseInt(i[2]))
                              : null))
                        : ("stampset_id" != o && "stamp_id" != o) ||
                          (r = ja._from(a).toInt64()),
                null != r && (e[o] = r))
              : null != a &&
                L.isObject(e) &&
                null != a.high &&
                null != a.low &&
                ((r = new E(a.high, a.low)), (e[o] = r));
          }
          return 1 == t.length && null != e.text ? e.text : e;
        },
        detectType: function (e) {
          if (null != e) {
            if ("string" == typeof e) return d.text;
            if (null != e.stamp_set) return d.stamp;
            if (null != e.stampset_id) return d.originalStamp;
            if (null != e.lat) return d.geo;
            if (null != e.file_id || null != e.path) return d.file;
            if (null != e.close) {
              var t = e.close;
              if ((L.deleteField(e, "close"), "yesno" == t))
                return d.yesOrNoClosed;
              if ("select" == t) return d.selectOneClosed;
              if ("task" == t) return d.todoClosed;
            } else if (null != e.in_reply_to) {
              if ("boolean" == typeof e.response) return d.yesOrNoReply;
              t = e.response;
              if ("number" == typeof t && (0 | t) === t)
                return d.selectOneReply;
              if (null != e.done) return d.todoDone;
            } else {
              if (null != e.question)
                return null == e.options ? d.yesOrNo : d.selectOne;
              if (null != e.title) return d.todo;
              if (null != e.note_title) return d.noteCreated;
            }
          }
          return d.unknown;
        },
        applyLimitToMessage: function (e) {
          switch (e.type[1]) {
            case 1:
              var t = e.content;
              null != t &&
                1024 < t.length &&
                (e.content = t.substring(0, 1024));
              break;
            case 2:
            case 5:
              t = e.content.text;
              null != t &&
                1024 < t.length &&
                (e.content.text = t.substring(0, 1024));
              break;
            case 13:
            case 15:
              t = e.content.question;
              if (
                (null != t &&
                  1024 < t.length &&
                  (e.content.question = t.substring(0, 1024)),
                e.type == d.selectOne)
              )
                for (
                  var n = e.content.options,
                    i =
                      (9 < n.length &&
                        ((n = n.slice(0, 9)), (e.content.options = n)),
                      0),
                    r = n.length;
                  i < r;
                ) {
                  var o = i++;
                  64 < n[o].length && (n[o] = n[o].substring(0, 64));
                }
              break;
            case 17:
              t = e.content.title;
              null != t &&
                1024 < t.length &&
                (e.content.title = t.substring(0, 1024));
          }
        },
        __class__: Pa,
      })),
      function (e, t, n) {
        ((this.proxy = e),
          (this.sendSpan = t),
          (this.storage = n),
          (this.sendQueue = this.loadQueue()),
          (this.sending = 0 < this.sendQueue.length),
          (this.lastSendNode = null),
          (this.lastSendTime = new Date(0)));
      }),
    Ha =
      (((t["albero_cli.proxy._SendQueueProxy.SendQueue"] = Ba).__name__ = [
        "albero_cli",
        "proxy",
        "_SendQueueProxy",
        "SendQueue",
      ]),
      (Ba.prototype = {
        pushQueue: function (e, t) {
          (this.sendQueue.push({ name: e, param: t }),
            this.saveQueue(),
            this.sending || ((this.sending = !0), this.sendNext()));
        },
        sendNext: function () {
          var e;
          this.sending && 0 != this.sendQueue.length
            ? ((e = new Date().getTime() - this.lastSendTime.getTime()),
              (e = 0 | Math.max(this.sendSpan - e, 100)),
              Xa.delay(z(this, this.sendNotification), e))
            : this.closeQueue();
        },
        retryAfter: function (e) {
          this.sending && null != this.lastSendNode
            ? (this.sendQueue.unshift(this.lastSendNode),
              this.saveQueue(),
              Xa.delay(z(this, this.sendNotification), e))
            : this.closeQueue();
        },
        sendNotification: function () {
          var e = this.sendQueue.shift();
          (this.saveQueue(),
            this.proxy.sendNotification(e.name, e.param),
            (this.lastSendNode = e),
            (this.lastSendTime = new Date()));
        },
        closeQueue: function () {
          ((this.sending = !1), (this.lastSendNode = null));
        },
        saveQueue: function () {
          null != this.storage && this.storage.save(this.sendQueue);
        },
        loadQueue: function () {
          if (null != this.storage) {
            var e = this.storage.load();
            if (null != e) return e;
          }
          return [];
        },
        __class__: Ba,
      }),
      function (e, t, n) {
        ((this.code = e), (this.message = t), (this.detail = n));
      }),
    Ga =
      (((t["albero_cli.value.ApiErrorResponse"] = Ha).__name__ = [
        "albero_cli",
        "value",
        "ApiErrorResponse",
      ]),
      (Ha.of = function (e) {
        switch (e.code) {
          case 400:
            if ("invalid parameter" == e.message)
              return new Ha(Ga.INVALID_PARAMETER);
            break;
          case 403:
            if ("forbidden" == e.message) return new Ha(Ga.FORBIDDEN);
            break;
          case 404:
            if ("not found" == e.message) return new Ha(Ga.NOT_FOUND);
            break;
          case 409:
            switch (e.message) {
              case "conflict":
                return new Ha(Ga.CONFLICT);
              case "locked by another user":
                return new Ha(Ga.LOCKED_BY_ANOTHER_USER, null, {
                  userId: e.detail.user_id,
                });
            }
            break;
          case 412:
            if ("client upgrade required" == e.message)
              return new Ha(Ga.CLIENT_UPGRADE_REQUIRED);
            break;
          case 429:
            return Ha.createTooManyRequestsError(
              e.message,
              e.detail.retry_after,
            );
        }
        return (
          F.w("please contact us: error = " + P.string(e)),
          new Ha(Ga.UNKNOWN)
        );
      }),
      (Ha.createInvalidParameterError = function (e) {
        return new Ha(Ga.INVALID_PARAMETER, e);
      }),
      (Ha.createTooManyRequestsError = function (e, t) {
        return new Ha(Ga.TOO_MANY_REQUESTS, e, { retryAfter: t });
      }),
      (Ha.prototype = { __class__: Ha }),
      function () {}),
    ja =
      (((t["albero_cli.value.ApiErrorCode"] = Ga).__name__ = [
        "albero_cli",
        "value",
        "ApiErrorCode",
      ]),
      function () {}),
    za =
      (((t["albero_cli.value.IDUtil"] = ja).__name__ = [
        "albero_cli",
        "value",
        "IDUtil",
      ]),
      (ja._from = function (e) {
        if (null != e) {
          if ("string" == typeof e) {
            var t = ja.parseInt64(e);
            if (null != t) return new za(t);
          }
          if (null != e.high && null != e.low) {
            t = ja.parseInt64(ja.stringifyInt64(e, !0));
            if (null != t) return new za(t);
          }
        }
        return null;
      }),
      (ja.parseInt64 = function (e) {
        return 0 < e.length && "_" == e.charAt(0)
          ? D.makeFromIdStr(e)
          : D.parse(e);
      }),
      (ja.stringifyInt64 = function (e, t) {
        return (t = null == t ? !1 : t)
          ? "_" + e.high + "_" + e.low
          : Va.toString(e);
      }),
      (ja.prototype = {
        from: function (e) {
          return ja._from(e);
        },
        __class__: ja,
      }),
      function (e) {
        this.value = e;
      }),
    Ya =
      (((t["albero_cli.value.ID"] = za).__name__ = [
        "albero_cli",
        "value",
        "ID",
      ]),
      (za.prototype = {
        toString: function () {
          return Va.toString(this.value);
        },
        toDaabID: function () {
          var e = this.value;
          return "_" + e.high + "_" + e.low;
        },
        toInt64: function () {
          return this.value;
        },
        eq: function (e) {
          var t;
          return (
            this == e ||
            (null != e &&
              ((t = this.value), (e = e.value), null != t) &&
              null != e &&
              t.high == e.high &&
              t.low == e.low)
          );
        },
        __class__: za,
      }),
      function () {}),
    Wa =
      (((t["albero_cli.value.note.DeleteNoteResult"] = Ya).__name__ = [
        "albero_cli",
        "value",
        "note",
        "DeleteNoteResult",
      ]),
      (Ya.of = function (e) {
        return new Ya();
      }),
      (Ya.prototype = { __class__: Ya }),
      function (e) {
        this.note = e;
      }),
    Ka =
      (((t["albero_cli.value.note.GetNoteResult"] = Wa).__name__ = [
        "albero_cli",
        "value",
        "note",
        "GetNoteResult",
      ]),
      (Wa.of = function (e) {
        return new Wa(Ma.of(e.note));
      }),
      (Wa.prototype = { __class__: Wa }),
      function (e) {
        this.note = e;
      }),
    w =
      (((t["albero_cli.value.note.UpdateNoteResult"] = Ka).__name__ = [
        "albero_cli",
        "value",
        "note",
        "UpdateNoteResult",
      ]),
      (Ka.of = function (e) {
        return new Ka(Ma.of(e.note));
      }),
      (Ka.prototype = { __class__: Ka }),
      ((t["haxe.IMap"] = _n).__name__ = ["haxe", "IMap"]),
      (_n.prototype = { __class__: _n }),
      {}),
    Va =
      (((t["haxe._Int32.Int32_Impl_"] = w).__name__ = [
        "haxe",
        "_Int32",
        "Int32_Impl_",
      ]),
      (w.ucompare = function (e, t) {
        return e < 0 ? (t < 0 ? (~t - ~e) | 0 : 1) : t < 0 ? -1 : (e - t) | 0;
      }),
      {}),
    qa =
      (((t["haxe._Int64.Int64_Impl_"] = Va).__name__ = [
        "haxe",
        "_Int64",
        "Int64_Impl_",
      ]),
      (Va.toString = function (e) {
        var t = e,
          e = new E(0, 0);
        if (t.high == e.high && t.low == e.low) return "0";
        var n = "",
          e = !1;
        t.high < 0 && (e = !0);
        for (var i = new E(0, 10); ; ) {
          var r = new E(0, 0);
          if (t.high == r.high && t.low == r.low) break;
          var o,
            a,
            r = Va.divMod(t, i);
          t =
            r.modulus.high < 0
              ? ((a = ~(o = r.modulus).high),
                0 == (o = -o.low) && (a++, (a |= 0)),
                (n = new E(a, o).low + n),
                (o = ~(a = r.quotient).high),
                0 == (a = -a.low) && (o++, (o |= 0)),
                new E(o, a))
              : ((n = r.modulus.low + n), r.quotient);
        }
        return (n = e ? "-" + n : n);
      }),
      (Va.divMod = function (e, t) {
        if (0 == t.high)
          switch (t.low) {
            case 0:
              throw new H("divide by zero");
            case 1:
              return { quotient: new E(e.high, e.low), modulus: new E(0, 0) };
          }
        var n,
          i = e.high < 0 != t.high < 0;
        ((n =
          e.high < 0
            ? ((o = ~e.high), 0 == (r = -e.low) && (o++, (o |= 0)), new E(o, r))
            : new E(e.high, e.low)),
          t.high < 0 &&
            ((o = ~t.high),
            0 == (r = -t.low) && (o++, (o |= 0)),
            (t = new E(o, r))));
        for (var r, o, a = new E(0, 0), s = new E(0, 1); !(t.high < 0); ) {
          var l = w.ucompare(t.high, n.high),
            l = 0 != l ? l : w.ucompare(t.low, n.low);
          t = new E((t.high << 1) | (t.low >>> 31), t.low << 1);
          if (((s = new E((s.high << 1) | (s.low >>> 31), s.low << 1)), 0 <= l))
            break;
        }
        for (;;) {
          var u = new E(0, 0);
          if (s.high == u.high && s.low == u.low) break;
          var c,
            u = w.ucompare(n.high, t.high);
          (0 <= (0 != u ? u : w.ucompare(n.low, t.low)) &&
            ((a = new E(a.high | s.high, a.low | s.low)),
            (u = (n.high - t.high) | 0),
            (c = (n.low - t.low) | 0),
            w.ucompare(n.low, t.low) < 0 && (u--, (u |= 0)),
            (n = new E(u, c))),
            (s = new E(s.high >>> 1, (s.high << 31) | (s.low >>> 1))));
          t = new E(t.high >>> 1, (t.high << 31) | (t.low >>> 1));
        }
        return (
          i &&
            ((o = ~a.high),
            0 == (r = -a.low) && (o++, (o |= 0)),
            (a = new E(o, r))),
          e.high < 0 &&
            ((i = ~n.high),
            0 == (o = -n.low) && (i++, (i |= 0)),
            (n = new E(i, o))),
          { quotient: a, modulus: n }
        );
      }),
      function () {}),
    Qa =
      (((t["haxe.Int64Helper"] = qa).__name__ = ["haxe", "Int64Helper"]),
      (qa.parseString = function (e) {
        for (
          var t = new E(0, 10),
            n = new E(0, 0),
            i = new E(0, 1),
            r = !1,
            o = I.trim(e),
            a =
              ("-" == o.charAt(0) && ((r = !0), (o = o.substring(1, o.length))),
              o.length),
            s = 0,
            l = a;
          s < l;
        ) {
          var u = s++,
            u = U.cca(o, a - 1 - u) - 48;
          if (u < 0 || 9 < u) throw new H("NumberFormatError");
          u = new E(u >> 31, u);
          if (r) {
            var c = 65535 & i.low,
              _ = i.low >>> 16,
              h = 65535 & u.low,
              d = u.low >>> 16,
              f = w._mul(c, h),
              h = w._mul(_, h),
              c = w._mul(c, d),
              _ = (((w._mul(_, d) + (c >>> 16)) | 0) + (h >>> 16)) | 0,
              f =
                ((c <<= 16),
                w.ucompare((d = (f + c) | 0), c) < 0 && (_++, (_ |= 0)),
                (h <<= 16),
                w.ucompare((d = (d + h) | 0), h) < 0 && (_++, (_ |= 0)),
                (_ =
                  (_ + ((w._mul(i.low, u.high) + w._mul(i.high, u.low)) | 0)) |
                  0),
                new E(_, d)),
              c = (n.high - f.high) | 0,
              h = (n.low - f.low) | 0;
            if (
              (w.ucompare(n.low, f.low) < 0 && (c--, (c |= 0)),
              !((n = new E(c, h)).high < 0))
            )
              throw new H("NumberFormatError: Underflow");
          } else {
            var _ = 65535 & i.low,
              d = i.low >>> 16,
              f = 65535 & u.low,
              c = u.low >>> 16,
              h = w._mul(_, f),
              m = w._mul(d, f),
              p = w._mul(_, c),
              g = (((w._mul(d, c) + (p >>> 16)) | 0) + (m >>> 16)) | 0,
              p =
                ((p <<= 16),
                w.ucompare((y = (h + p) | 0), p) < 0 && (g++, (g |= 0)),
                (m <<= 16),
                w.ucompare((y = (y + m) | 0), m) < 0 && (g++, (g |= 0)),
                (g =
                  (g + ((w._mul(i.low, u.high) + w._mul(i.high, u.low)) | 0)) |
                  0),
                new E(g, y)),
              m = (n.high + p.high) | 0,
              u = (n.low + p.low) | 0;
            if (
              (w.ucompare(u, n.low) < 0 && (m++, (m |= 0)),
              (n = new E(m, u)).high < 0)
            )
              throw new H("NumberFormatError: Overflow");
          }
          var g = 65535 & i.low,
            y = i.low >>> 16,
            p = 65535 & t.low,
            m = t.low >>> 16,
            u = w._mul(g, p),
            S = w._mul(y, p),
            v = w._mul(g, m),
            T = (((w._mul(y, m) + (v >>> 16)) | 0) + (S >>> 16)) | 0,
            T =
              ((v <<= 16),
              w.ucompare((u = (u + v) | 0), v) < 0 && (T++, (T |= 0)),
              (S <<= 16),
              w.ucompare((u = (u + S) | 0), S) < 0 && (T++, (T |= 0)),
              (T + ((w._mul(i.low, t.high) + w._mul(i.high, t.low)) | 0)) | 0),
            i = new E(T, u);
        }
        return n;
      }),
      function () {
        ((this.buf = new Cn()),
          (this.cache = []),
          (this.useCache = Qa.USE_CACHE),
          (this.useEnumIndex = Qa.USE_ENUM_INDEX),
          (this.shash = new B()),
          (this.scount = 0));
      }),
    Xa =
      (((t["haxe.Serializer"] = Qa).__name__ = ["haxe", "Serializer"]),
      (Qa.run = function (e) {
        var t = new Qa();
        return (t.serialize(e), t.toString());
      }),
      (Qa.prototype = {
        toString: function () {
          return this.buf.b;
        },
        serializeString: function (e) {
          var t,
            n = this.shash,
            n = null != Y[e] ? n.getReserved(e) : n.h[e];
          null != n
            ? ((this.buf.b += "R"), (this.buf.b += null == n ? "null" : "" + n))
            : ((n = this.shash),
              (t = this.scount++),
              null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t),
              (this.buf.b += "y"),
              (e = encodeURIComponent(e)),
              (this.buf.b += P.string(e.length)),
              (this.buf.b += ":"),
              (this.buf.b += null == e ? "null" : "" + e));
        },
        serializeRef: function (e) {
          for (var t = typeof e, n = 0, i = this.cache.length; n < i; ) {
            var r = n++,
              o = this.cache[r];
            if (typeof o == t && o == e)
              return (
                (this.buf.b += "r"),
                (this.buf.b += null == r ? "null" : "" + r),
                !0
              );
          }
          return (this.cache.push(e), !1);
        },
        serializeFields: function (e) {
          for (var t = 0, n = L.fields(e); t < n.length; ) {
            var i = n[t];
            (++t, this.serializeString(i), this.serialize(L.field(e, i)));
          }
          this.buf.b += "g";
        },
        serialize: function (e) {
          var t = x.typeof(e);
          switch (t[1]) {
            case 0:
              this.buf.b += "n";
              break;
            case 1:
              var n = e;
              0 == n
                ? (this.buf.b += "z")
                : ((this.buf.b += "i"),
                  (this.buf.b += null == n ? "null" : "" + n));
              break;
            case 2:
              n = e;
              isNaN(n)
                ? (this.buf.b += "k")
                : isFinite(n)
                  ? ((this.buf.b += "d"),
                    (this.buf.b += null == n ? "null" : "" + n))
                  : (this.buf.b += n < 0 ? "m" : "p");
              break;
            case 3:
              this.buf.b += e ? "t" : "f";
              break;
            case 4:
              if (G.__instanceof(e, Ks)) {
                n = x.getClassName(e);
                ((this.buf.b += "A"), this.serializeString(n));
              } else if (G.__instanceof(e, Vs))
                ((this.buf.b += "B"), this.serializeString(x.getEnumName(e)));
              else {
                if (this.useCache && this.serializeRef(e)) return;
                ((this.buf.b += "o"), this.serializeFields(e));
              }
              break;
            case 5:
              throw new H("Cannot serialize function");
            case 6:
              var i = t[2];
              if (i == String) this.serializeString(e);
              else if (!this.useCache || !this.serializeRef(e))
                switch (i) {
                  case Array:
                    var r = 0;
                    this.buf.b += "a";
                    for (var o = 0, a = e.length; o < a; ) {
                      var s = o++;
                      null == e[s]
                        ? ++r
                        : (0 < r &&
                            (1 == r
                              ? (this.buf.b += "n")
                              : ((this.buf.b += "u"),
                                (this.buf.b += null == r ? "null" : "" + r)),
                            (r = 0)),
                          this.serialize(e[s]));
                    }
                    (0 < r &&
                      (1 == r
                        ? (this.buf.b += "n")
                        : ((this.buf.b += "u"),
                          (this.buf.b += null == r ? "null" : "" + r))),
                      (this.buf.b += "h"));
                    break;
                  case Date:
                    var l = e;
                    ((this.buf.b += "v"),
                      (this.buf.b += P.string(l.getTime())));
                    break;
                  case En:
                    this.buf.b += "l";
                    for (var u = e.h; null != u; ) {
                      var c = u.item,
                        u = u.next;
                      this.serialize(c);
                    }
                    this.buf.b += "h";
                    break;
                  case ts:
                    this.buf.b += "q";
                    for (var _ = e, h = _.keys(); h.hasNext(); ) {
                      var d = h.next();
                      ((this.buf.b += ":"),
                        (this.buf.b += null == d ? "null" : "" + d),
                        this.serialize(_.h[d]));
                    }
                    this.buf.b += "h";
                    break;
                  case ns:
                    this.buf.b += "M";
                    for (var f = e, m = f.keys(); m.hasNext(); ) {
                      var p = m.next(),
                        g = L.field(p, "__id__");
                      (L.deleteField(p, "__id__"),
                        this.serialize(p),
                        (p.__id__ = g),
                        this.serialize(f.h[p.__id__]));
                    }
                    this.buf.b += "h";
                    break;
                  case B:
                    this.buf.b += "b";
                    for (var y = e, S = y.keys(); S.hasNext(); ) {
                      var v = S.next();
                      (this.serializeString(v),
                        this.serialize(
                          null != Y[v] ? y.getReserved(v) : y.h[v],
                        ));
                    }
                    this.buf.b += "h";
                    break;
                  case rs:
                    var T,
                      I = e,
                      E =
                        ((this.buf.b += "s"),
                        (this.buf.b += P.string(Math.ceil((8 * I.length) / 6))),
                        (this.buf.b += ":"),
                        0),
                      w = I.length - 2;
                    if (null == (N = Qa.BASE64_CODES)) {
                      for (
                        var l = Qa.BASE64.length,
                          N = new Array(l),
                          A = 0,
                          b = Qa.BASE64.length;
                        A < b;
                      ) {
                        var M = A++;
                        N[M] = U.cca(Qa.BASE64, M);
                      }
                      Qa.BASE64_CODES = N;
                    }
                    for (; E < w; ) {
                      var C = I.b[E++],
                        k = I.b[E++],
                        D = I.b[E++];
                      ((this.buf.b += String.fromCharCode(N[C >> 2])),
                        (this.buf.b += String.fromCharCode(
                          N[63 & ((C << 4) | (k >> 4))],
                        )),
                        (this.buf.b += String.fromCharCode(
                          N[63 & ((k << 2) | (D >> 6))],
                        )),
                        (this.buf.b += String.fromCharCode(N[63 & D])));
                    }
                    E == w
                      ? ((l = I.b[E++]),
                        (T = I.b[E++]),
                        (this.buf.b += String.fromCharCode(N[l >> 2])),
                        (this.buf.b += String.fromCharCode(
                          N[63 & ((l << 4) | (T >> 4))],
                        )),
                        (this.buf.b += String.fromCharCode(N[(T << 2) & 63])))
                      : E == 1 + w &&
                        ((l = I.b[E++]),
                        (this.buf.b += String.fromCharCode(N[l >> 2])),
                        (this.buf.b += String.fromCharCode(N[(l << 4) & 63])));
                    break;
                  default:
                    (this.useCache && this.cache.pop(),
                      null != e.hxSerialize
                        ? ((this.buf.b += "C"),
                          this.serializeString(x.getClassName(i)),
                          this.useCache && this.cache.push(e),
                          e.hxSerialize(this),
                          (this.buf.b += "g"))
                        : ((this.buf.b += "c"),
                          this.serializeString(x.getClassName(i)),
                          this.useCache && this.cache.push(e),
                          this.serializeFields(e)));
                }
              break;
            case 7:
              n = t[2];
              if (this.useCache) {
                if (this.serializeRef(e)) return;
                this.cache.pop();
              }
              ((this.buf.b += P.string(this.useEnumIndex ? "j" : "w")),
                this.serializeString(x.getEnumName(n)),
                this.useEnumIndex
                  ? ((this.buf.b += ":"), (this.buf.b += P.string(e[1])))
                  : this.serializeString(e[0]),
                (this.buf.b += ":"));
              for (
                var n = e.length,
                  R = ((this.buf.b += P.string(n - 2)), 2),
                  O = n;
                R < O;
              ) {
                var F = R++;
                this.serialize(e[F]);
              }
              this.useCache && this.cache.push(e);
              break;
            default:
              throw new H("Cannot serialize " + P.string(e));
          }
        },
        __class__: Qa,
      }),
      function (e) {
        var t = this;
        this.id = setInterval(function () {
          t.run();
        }, e);
      }),
    Ja =
      (((t["haxe.Timer"] = Xa).__name__ = ["haxe", "Timer"]),
      (Xa.delay = function (e, t) {
        var n = new Xa(t);
        return (
          (n.run = function () {
            (n.stop(), e());
          }),
          n
        );
      }),
      (Xa.prototype = {
        stop: function () {
          null != this.id && (clearInterval(this.id), (this.id = null));
        },
        run: function () {},
        __class__: Xa,
      }),
      function () {}),
    Za =
      (((t["haxe.ds.BalancedTree"] = Ja).__name__ = [
        "haxe",
        "ds",
        "BalancedTree",
      ]),
      (Ja.prototype = {
        set: function (e, t) {
          this.root = this.setLoop(e, t, this.root);
        },
        get: function (e) {
          for (var t = this.root; null != t; ) {
            var n = this.compare(e, t.key);
            if (0 == n) return t.value;
            t = n < 0 ? t.left : t.right;
          }
          return null;
        },
        remove: function (e) {
          try {
            return ((this.root = this.removeLoop(e, this.root)), !0);
          } catch (e) {
            if ((e instanceof H && (e = e.val), G.__instanceof(e, String)))
              return !1;
            throw e;
          }
        },
        iterator: function () {
          var e = [];
          return (this.iteratorLoop(this.root, e), U.iter(e));
        },
        keys: function () {
          var e = [];
          return (this.keysLoop(this.root, e), U.iter(e));
        },
        setLoop: function (e, t, n) {
          var i;
          return null == n
            ? new Za(null, e, t, null)
            : 0 == (i = this.compare(e, n.key))
              ? new Za(n.left, e, t, n.right, null == n ? 0 : n._height)
              : i < 0
                ? ((i = this.setLoop(e, t, n.left)),
                  this.balance(i, n.key, n.value, n.right))
                : ((i = this.setLoop(e, t, n.right)),
                  this.balance(n.left, n.key, n.value, i));
        },
        removeLoop: function (e, t) {
          if (null == t) throw new H("Not_found");
          var n = this.compare(e, t.key);
          return 0 == n
            ? this.merge(t.left, t.right)
            : n < 0
              ? this.balance(
                  this.removeLoop(e, t.left),
                  t.key,
                  t.value,
                  t.right,
                )
              : this.balance(
                  t.left,
                  t.key,
                  t.value,
                  this.removeLoop(e, t.right),
                );
        },
        iteratorLoop: function (e, t) {
          null != e &&
            (this.iteratorLoop(e.left, t),
            t.push(e.value),
            this.iteratorLoop(e.right, t));
        },
        keysLoop: function (e, t) {
          null != e &&
            (this.keysLoop(e.left, t),
            t.push(e.key),
            this.keysLoop(e.right, t));
        },
        merge: function (e, t) {
          var n;
          return null == e
            ? t
            : null == t
              ? e
              : ((n = this.minBinding(t)),
                this.balance(e, n.key, n.value, this.removeMinBinding(t)));
        },
        minBinding: function (e) {
          if (null == e) throw new H("Not_found");
          return null == e.left ? e : this.minBinding(e.left);
        },
        removeMinBinding: function (e) {
          return null == e.left
            ? e.right
            : this.balance(
                this.removeMinBinding(e.left),
                e.key,
                e.value,
                e.right,
              );
        },
        balance: function (e, t, n, i) {
          var r,
            o = null == e ? 0 : e._height,
            a = null == i ? 0 : i._height;
          return a + 2 < o
            ? (null == (r = e.left) ? 0 : r._height) >=
              (null == (r = e.right) ? 0 : r._height)
              ? new Za(e.left, e.key, e.value, new Za(e.right, t, n, i))
              : new Za(
                  new Za(e.left, e.key, e.value, e.right.left),
                  e.right.key,
                  e.right.value,
                  new Za(e.right.right, t, n, i),
                )
            : o + 2 < a
              ? (null == (r = i.right) ? 0 : r._height) >
                (null == (r = i.left) ? 0 : r._height)
                ? new Za(new Za(e, t, n, i.left), i.key, i.value, i.right)
                : new Za(
                    new Za(e, t, n, i.left.left),
                    i.left.key,
                    i.left.value,
                    new Za(i.left.right, i.key, i.value, i.right),
                  )
              : new Za(e, t, n, i, (a < o ? o : a) + 1);
        },
        compare: function (e, t) {
          return L.compare(e, t);
        },
        __class__: Ja,
      }),
      function (e, t, n, i, r) {
        (null == r && (r = -1),
          (this.left = e),
          (this.key = t),
          (this.value = n),
          (this.right = i),
          -1 == r
            ? ((e = this.left),
              (t = this.right),
              (e =
                (null == e ? 0 : e._height) > (null == t ? 0 : t._height)
                  ? null == (n = this.left)
                    ? 0
                    : n._height
                  : null == (i = this.right)
                    ? 0
                    : i._height),
              (this._height = e + 1))
            : (this._height = r));
      }),
    $a =
      (((t["haxe.ds.TreeNode"] = Za).__name__ = ["haxe", "ds", "TreeNode"]),
      (Za.prototype = { __class__: Za }),
      (t["haxe.ds.Either"] = {
        __ename__: ["haxe", "ds", "Either"],
        __constructs__: ["Left", "Right"],
      })),
    es =
      (($a.Left = function (e) {
        e = ["Left", 0, e];
        return ((e.__enum__ = $a), (e.toString = a), e);
      }),
      ($a.Right = function (e) {
        e = ["Right", 1, e];
        return ((e.__enum__ = $a), (e.toString = a), e);
      }),
      ($a.__empty_constructs__ = []),
      function () {
        Ja.call(this);
      }),
    ts =
      (((t["haxe.ds.EnumValueMap"] = es).__name__ = [
        "haxe",
        "ds",
        "EnumValueMap",
      ]),
      (es.__interfaces__ = [_n]),
      (es.__super__ = Ja),
      (es.prototype = e(Ja.prototype, {
        compare: function (e, t) {
          var n = e[1] - t[1];
          return 0 != n
            ? n
            : ((n = e.slice(2)),
              (e = t.slice(2)),
              0 == n.length && 0 == e.length ? 0 : this.compareArgs(n, e));
        },
        compareArgs: function (e, t) {
          var n = e.length - t.length;
          if (0 != n) return n;
          for (var i = 0, r = e.length; i < r; ) {
            var o = i++,
              o = this.compareArg(e[o], t[o]);
            if (0 != o) return o;
          }
          return 0;
        },
        compareArg: function (e, t) {
          return L.isEnumValue(e) && L.isEnumValue(t)
            ? this.compare(e, t)
            : e instanceof Array &&
                null == e.__enum__ &&
                t instanceof Array &&
                null == t.__enum__
              ? this.compareArgs(e, t)
              : L.compare(e, t);
        },
        __class__: es,
      })),
      function () {
        this.h = {};
      }),
    ns =
      (((t["haxe.ds.IntMap"] = ts).__name__ = ["haxe", "ds", "IntMap"]),
      (ts.__interfaces__ = [_n]),
      (ts.prototype = {
        get: function (e) {
          return this.h[e];
        },
        remove: function (e) {
          return !!this.h.hasOwnProperty(e) && (delete this.h[e], !0);
        },
        keys: function () {
          var e,
            t = [];
          for (e in this.h) this.h.hasOwnProperty(e) && t.push(0 | e);
          return U.iter(t);
        },
        iterator: function () {
          return {
            ref: this.h,
            it: this.keys(),
            hasNext: function () {
              return this.it.hasNext();
            },
            next: function () {
              var e = this.it.next();
              return this.ref[e];
            },
          };
        },
        __class__: ts,
      }),
      function () {
        this.h = { __keys__: {} };
      }),
    m =
      (((t["haxe.ds.ObjectMap"] = ns).__name__ = ["haxe", "ds", "ObjectMap"]),
      (ns.__interfaces__ = [_n]),
      (ns.assignId = function (e) {
        return (e.__id__ = ++ns.count);
      }),
      (ns.getId = function (e) {
        return e.__id__;
      }),
      (ns.prototype = {
        set: function (e, t) {
          var n = e.__id__ || (e.__id__ = ++ns.count);
          ((this.h[n] = t), (this.h.__keys__[n] = e));
        },
        get: function (e) {
          return this.h[e.__id__];
        },
        keys: function () {
          var e,
            t = [];
          for (e in this.h.__keys__)
            this.h.hasOwnProperty(e) && t.push(this.h.__keys__[e]);
          return U.iter(t);
        },
        iterator: function () {
          return {
            ref: this.h,
            it: this.keys(),
            hasNext: function () {
              return this.it.hasNext();
            },
            next: function () {
              var e = this.it.next();
              return this.ref[e.__id__];
            },
          };
        },
        __class__: ns,
      }),
      (t["haxe.ds.Option"] = {
        __ename__: ["haxe", "ds", "Option"],
        __constructs__: ["Some", "None"],
      })),
    is =
      ((m.Some = function (e) {
        e = ["Some", 0, e];
        return ((e.__enum__ = m), (e.toString = a), e);
      }),
      (m.None = ["None", 1]),
      (m.None.toString = a),
      ((m.None.__enum__ = m).__empty_constructs__ = [m.None]),
      function (e, t) {
        ((this.map = e),
          (this.keys = t),
          (this.index = 0),
          (this.count = t.length));
      }),
    B =
      (((t["haxe.ds._StringMap.StringMapIterator"] = is).__name__ = [
        "haxe",
        "ds",
        "_StringMap",
        "StringMapIterator",
      ]),
      (is.prototype = {
        hasNext: function () {
          return this.index < this.count;
        },
        next: function () {
          var e = this.map,
            t = this.keys[this.index++];
          return null != Y[t] ? e.getReserved(t) : e.h[t];
        },
        __class__: is,
      }),
      function () {
        this.h = {};
      }),
    rs =
      (((t["haxe.ds.StringMap"] = B).__name__ = ["haxe", "ds", "StringMap"]),
      (B.__interfaces__ = [_n]),
      (B.prototype = {
        get: function (e) {
          return null != Y[e] ? this.getReserved(e) : this.h[e];
        },
        setReserved: function (e, t) {
          (null == this.rh && (this.rh = {}), (this.rh["$" + e] = t));
        },
        getReserved: function (e) {
          return null == this.rh ? null : this.rh["$" + e];
        },
        existsReserved: function (e) {
          return null != this.rh && this.rh.hasOwnProperty("$" + e);
        },
        remove: function (e) {
          return null != Y[e]
            ? ((e = "$" + e),
              !(
                null == this.rh ||
                !this.rh.hasOwnProperty(e) ||
                (delete this.rh[e], 0)
              ))
            : !!this.h.hasOwnProperty(e) && (delete this.h[e], !0);
        },
        keys: function () {
          return U.iter(this.arrayKeys());
        },
        arrayKeys: function () {
          var e = [];
          for (t in this.h) this.h.hasOwnProperty(t) && e.push(t);
          if (null != this.rh)
            for (var t in this.rh) 36 == t.charCodeAt(0) && e.push(t.substr(1));
          return e;
        },
        iterator: function () {
          return new is(this, this.arrayKeys());
        },
        __class__: B,
      }),
      function (e) {
        ((this.length = e.byteLength),
          (this.b = new Uint8Array(e)),
          ((this.b.bufferValue = e).hxBytes = this),
          (e.bytes = this.b));
      }),
    os =
      (((t["haxe.io.Bytes"] = rs).__name__ = ["haxe", "io", "Bytes"]),
      (rs.alloc = function (e) {
        return new rs(new ArrayBuffer(e));
      }),
      (rs.ofString = function (e) {
        for (var t = [], n = 0; n < e.length; ) {
          var i = e.charCodeAt(n++);
          (i =
            55296 <= i && i <= 56319
              ? ((i - 55232) << 10) | (1023 & e.charCodeAt(n++))
              : i) <= 127
            ? t.push(i)
            : (i <= 2047
                ? t.push(192 | (i >> 6))
                : (i <= 65535
                    ? t.push(224 | (i >> 12))
                    : (t.push(240 | (i >> 18)), t.push(128 | ((i >> 12) & 63))),
                  t.push(128 | ((i >> 6) & 63))),
              t.push(128 | (63 & i)));
        }
        return new rs(new Uint8Array(t).buffer);
      }),
      (rs.ofData = function (e) {
        var t = e.hxBytes;
        return null != t ? t : new rs(e);
      }),
      (rs.fastGet = function (e, t) {
        return e.bytes[t];
      }),
      (rs.prototype = {
        getString: function (e, t) {
          if (e < 0 || t < 0 || e + t > this.length)
            throw new H(us.OutsideBounds);
          for (
            var n = "", i = this.b, r = String.fromCharCode, o = e, a = e + t;
            o < a;
          ) {
            var s = i[o++];
            if (s < 128) {
              if (0 == s) break;
              n += r(s);
            } else
              s < 224
                ? (n += r(((63 & s) << 6) | (127 & i[o++])))
                : s < 240
                  ? (n += r(
                      ((31 & s) << 12) | ((127 & i[o++]) << 6) | (127 & i[o++]),
                    ))
                  : (n =
                      (n += r(
                        55232 +
                          ((s =
                            ((15 & s) << 18) |
                            ((127 & i[o++]) << 12) |
                            ((127 & i[o++]) << 6) |
                            (127 & i[o++])) >>
                            10),
                      )) + r((1023 & s) | 56320));
          }
          return n;
        },
        toString: function () {
          return this.getString(0, this.length);
        },
        __class__: rs,
      }),
      function () {
        this.b = [];
      }),
    as =
      (((t["haxe.io.BytesBuffer"] = os).__name__ = [
        "haxe",
        "io",
        "BytesBuffer",
      ]),
      (os.prototype = {
        getBytes: function () {
          var e = new rs(new Uint8Array(this.b).buffer);
          return ((this.b = null), e);
        },
        __class__: os,
      }),
      ((t["haxe.io.Input"] = hn).__name__ = ["haxe", "io", "Input"]),
      (hn.prototype = {
        readByte: function () {
          throw new H("Not implemented");
        },
        readBytes: function (e, t, n) {
          var i = n,
            r = e.b;
          if (t < 0 || n < 0 || e.length < t + n) throw new H(us.OutsideBounds);
          try {
            for (; 0 < i; ) ((r[t] = this.readByte()), ++t, --i);
          } catch (e) {
            if ((e instanceof H && (e = e.val), !G.__instanceof(e, ls)))
              throw e;
          }
          return n - i;
        },
        set_bigEndian: function (e) {
          return (this.bigEndian = e);
        },
        read: function (e) {
          for (var t = new rs(new ArrayBuffer(e)), n = 0; 0 < e; ) {
            var i = this.readBytes(t, n, e);
            if (0 == i) throw new H(us.Blocked);
            ((n += i), (e -= i));
          }
          return t;
        },
        readFloat: function () {
          return cs.i32ToFloat(this.readInt32());
        },
        readDouble: function () {
          var e = this.readInt32(),
            t = this.readInt32();
          return this.bigEndian ? cs.i64ToDouble(t, e) : cs.i64ToDouble(e, t);
        },
        readInt8: function () {
          var e = this.readByte();
          return 128 <= e ? e - 256 : e;
        },
        readInt16: function () {
          var e = this.readByte(),
            t = this.readByte(),
            e = this.bigEndian ? t | (e << 8) : e | (t << 8);
          return 0 != (32768 & e) ? e - 65536 : e;
        },
        readUInt16: function () {
          var e = this.readByte(),
            t = this.readByte();
          return this.bigEndian ? t | (e << 8) : e | (t << 8);
        },
        readInt32: function () {
          var e = this.readByte(),
            t = this.readByte(),
            n = this.readByte(),
            i = this.readByte();
          return this.bigEndian
            ? i | (n << 8) | (t << 16) | (e << 24)
            : e | (t << 8) | (n << 16) | (i << 24);
        },
        __class__: hn,
      }),
      function (e, t, n) {
        if (
          (null == t && (t = 0),
          null == n && (n = e.length - t),
          t < 0 || n < 0 || t + n > e.length)
        )
          throw new H(us.OutsideBounds);
        ((this.b = e.b), (this.pos = t), (this.len = n), (this.totlen = n));
      }),
    ss =
      (((t["haxe.io.BytesInput"] = as).__name__ = ["haxe", "io", "BytesInput"]),
      (as.__super__ = hn),
      (as.prototype = e(hn.prototype, {
        readByte: function () {
          if (0 == this.len) throw new H(new ls());
          return (this.len--, this.b[this.pos++]);
        },
        readBytes: function (e, t, n) {
          if (t < 0 || n < 0 || t + n > e.length) throw new H(us.OutsideBounds);
          if (0 == this.len && 0 < n) throw new H(new ls());
          this.len < n && (n = this.len);
          for (var i = this.b, r = e.b, o = 0, a = n; o < a; ) {
            var s = o++;
            r[t + s] = i[this.pos + s];
          }
          return ((this.pos += n), (this.len -= n), n);
        },
        __class__: as,
      })),
      ((t["haxe.io.Output"] = dn).__name__ = ["haxe", "io", "Output"]),
      (dn.prototype = {
        writeByte: function (e) {
          throw new H("Not implemented");
        },
        writeBytes: function (e, t, n) {
          if (t < 0 || n < 0 || t + n > e.length) throw new H(us.OutsideBounds);
          for (var i = e.b, r = n; 0 < r; ) (this.writeByte(i[t]), ++t, --r);
          return n;
        },
        set_bigEndian: function (e) {
          return (this.bigEndian = e);
        },
        write: function (e) {
          for (var t = e.length, n = 0; 0 < t; ) {
            var i = this.writeBytes(e, n, t);
            if (0 == i) throw new H(us.Blocked);
            ((n += i), (t -= i));
          }
        },
        writeFloat: function (e) {
          this.writeInt32(cs.floatToI32(e));
        },
        writeDouble: function (e) {
          e = cs.doubleToI64(e);
          this.bigEndian
            ? (this.writeInt32(e.high), this.writeInt32(e.low))
            : (this.writeInt32(e.low), this.writeInt32(e.high));
        },
        writeInt8: function (e) {
          if (e < -128 || 128 <= e) throw new H(us.Overflow);
          this.writeByte(255 & e);
        },
        writeInt16: function (e) {
          if (e < -32768 || 32768 <= e) throw new H(us.Overflow);
          this.writeUInt16(65535 & e);
        },
        writeUInt16: function (e) {
          if (e < 0 || 65536 <= e) throw new H(us.Overflow);
          this.bigEndian
            ? (this.writeByte(e >> 8), this.writeByte(255 & e))
            : (this.writeByte(255 & e), this.writeByte(e >> 8));
        },
        writeInt32: function (e) {
          this.bigEndian
            ? (this.writeByte(e >>> 24),
              this.writeByte((e >> 16) & 255),
              this.writeByte((e >> 8) & 255),
              this.writeByte(255 & e))
            : (this.writeByte(255 & e),
              this.writeByte((e >> 8) & 255),
              this.writeByte((e >> 16) & 255),
              this.writeByte(e >>> 24));
        },
        __class__: dn,
      }),
      function () {
        this.b = new os();
      }),
    ls =
      (((t["haxe.io.BytesOutput"] = ss).__name__ = [
        "haxe",
        "io",
        "BytesOutput",
      ]),
      (ss.__super__ = dn),
      (ss.prototype = e(dn.prototype, {
        writeByte: function (e) {
          this.b.b.push(e);
        },
        writeBytes: function (e, t, n) {
          var i = this.b;
          if (t < 0 || n < 0 || t + n > e.length) throw new H(us.OutsideBounds);
          i.b;
          for (var r = e.b, o = t, a = t + n; o < a; ) {
            var s = o++;
            i.b.push(r[s]);
          }
          return n;
        },
        getBytes: function () {
          return this.b.getBytes();
        },
        __class__: ss,
      })),
      function () {}),
    us =
      (((t["haxe.io.Eof"] = ls).__name__ = ["haxe", "io", "Eof"]),
      (ls.prototype = {
        toString: function () {
          return "Eof";
        },
        __class__: ls,
      }),
      (t["haxe.io.Error"] = {
        __ename__: ["haxe", "io", "Error"],
        __constructs__: ["Blocked", "Overflow", "OutsideBounds", "Custom"],
      })),
    cs =
      ((us.Blocked = ["Blocked", 0]),
      (us.Blocked.toString = a),
      ((us.Blocked.__enum__ = us).Overflow = ["Overflow", 1]),
      (us.Overflow.toString = a),
      ((us.Overflow.__enum__ = us).OutsideBounds = ["OutsideBounds", 2]),
      (us.OutsideBounds.toString = a),
      ((us.OutsideBounds.__enum__ = us).Custom = function (e) {
        e = ["Custom", 3, e];
        return ((e.__enum__ = us), (e.toString = a), e);
      }),
      (us.__empty_constructs__ = [us.Blocked, us.Overflow, us.OutsideBounds]),
      function () {}),
    _s =
      (((t["haxe.io.FPHelper"] = cs).__name__ = ["haxe", "io", "FPHelper"]),
      (cs.i32ToFloat = function (e) {
        var t = (e >>> 23) & 255,
          n = 8388607 & e;
        return 0 == n && 0 == t
          ? 0
          : (1 - ((e >>> 31) << 1)) *
              (1 + Math.pow(2, -23) * n) *
              Math.pow(2, t - 127);
      }),
      (cs.floatToI32 = function (e) {
        var t, n;
        return 0 == e
          ? 0
          : ((n = e < 0 ? -e : e),
            (t = Math.floor(Math.log(n) / 0.6931471805599453)) < -127
              ? (t = -127)
              : 128 < t && (t = 128),
            8388608 == (n = Math.round(8388608 * (n / Math.pow(2, t) - 1))) &&
              t < 128 &&
              ((n = 0), ++t),
            (e < 0 ? -2147483648 : 0) | ((t + 127) << 23) | n);
      }),
      (cs.i64ToDouble = function (e, t) {
        var n = ((t >> 20) & 2047) - 1023,
          e =
            4294967296 * (1048575 & t) +
            2147483648 * (e >>> 31) +
            (2147483647 & e);
        return 0 == e && -1023 == n
          ? 0
          : (1 - ((t >>> 31) << 1)) *
              (1 + Math.pow(2, -52) * e) *
              Math.pow(2, n);
      }),
      (cs.doubleToI64 = function (e) {
        var t,
          n,
          i,
          r = cs.i64tmp;
        return (
          0 == e
            ? ((r.low = 0), (r.high = 0))
            : isFinite(e)
              ? ((n = e < 0 ? -e : e),
                (t = Math.floor(Math.log(n) / 0.6931471805599453)),
                (i =
                  ((n = Math.round(
                    4503599627370496 * (n / Math.pow(2, t) - 1),
                  )) /
                    4294967296) |
                  0),
                (r.low = 0 | n),
                (r.high = (e < 0 ? -2147483648 : 0) | ((t + 1023) << 20) | i))
              : 0 < e
                ? ((r.low = 0), (r.high = 2146435072))
                : ((r.low = 0), (r.high = -1048576)),
          r
        );
      }),
      function () {}),
    H =
      (((t["haxe.rtti.Meta"] = _s).__name__ = ["haxe", "rtti", "Meta"]),
      (_s.getMeta = function (e) {
        return e.__meta__;
      }),
      (_s.getFields = function (e) {
        e = _s.getMeta(e);
        return null == e || null == e.fields ? {} : e.fields;
      }),
      function (e) {
        (Error.call(this),
          (this.val = e),
          (this.message = String(e)),
          Error.captureStackTrace && Error.captureStackTrace(this, H));
      }),
    G =
      (((t["js._Boot.HaxeError"] = H).__name__ = ["js", "_Boot", "HaxeError"]),
      (H.wrap = function (e) {
        return e instanceof Error ? e : new H(e);
      }),
      (H.__super__ = Error),
      (H.prototype = e(Error.prototype, { __class__: H })),
      function () {}),
    hs =
      (((t["js.Boot"] = G).__name__ = ["js", "Boot"]),
      (G.getClass = function (e) {
        var t;
        return e instanceof Array && null == e.__enum__
          ? Array
          : null != (t = e.__class__)
            ? t
            : null != (t = G.__nativeClassName(e))
              ? G.__resolveNativeClass(t)
              : null;
      }),
      (G.__string_rec = function (e, t) {
        if (null == e) return "null";
        if (5 <= t.length) return "<...>";
        var n = typeof e;
        switch (
          (n = "function" == n && (e.__name__ || e.__ename__) ? "object" : n)
        ) {
          case "function":
            return "<function>";
          case "object":
            if (e instanceof Array) {
              if (e.__enum__) {
                if (2 == e.length) return e[0];
                for (
                  var i = e[0] + "(", r = ((t += "\t"), 2), o = e.length;
                  r < o;
                ) {
                  var a = r++;
                  i +=
                    2 != a
                      ? "," + G.__string_rec(e[a], t)
                      : G.__string_rec(e[a], t);
                }
                return i + ")";
              }
              for (var s = "[", l = ((t += "\t"), 0), u = e.length; l < u; ) {
                var c = l++;
                s += (0 < c ? "," : "") + G.__string_rec(e[c], t);
              }
              return (s += "]");
            }
            try {
              _ = e.toString;
            } catch (e) {
              return "???";
            }
            if (null != _ && _ != Object.toString && "function" == typeof _) {
              var _ = e.toString();
              if ("[object Object]" != _) return _;
            }
            var h = null,
              d = "{\n",
              f = ((t += "\t"), null != e.hasOwnProperty);
            for (h in e)
              (f && !e.hasOwnProperty(h)) ||
                ("prototype" != h &&
                  "__class__" != h &&
                  "__super__" != h &&
                  "__interfaces__" != h &&
                  "__properties__" != h &&
                  (2 != d.length && (d += ", \n"),
                  (d += t + h + " : " + G.__string_rec(e[h], t))));
            return (d += "\n" + (t = t.substring(1)) + "}");
          case "string":
            return e;
          default:
            return String(e);
        }
      }),
      (G.__interfLoop = function (e, t) {
        if (null == e) return !1;
        if (e == t) return !0;
        var n = e.__interfaces__;
        if (null != n)
          for (var i = 0, r = n.length; i < r; ) {
            var o = n[i++];
            if (o == t || G.__interfLoop(o, t)) return !0;
          }
        return G.__interfLoop(e.__super__, t);
      }),
      (G.__instanceof = function (e, t) {
        if (null == t) return !1;
        switch (t) {
          case Array:
            return e instanceof Array && null == e.__enum__;
          case Ws:
            return "boolean" == typeof e;
          case zs:
            return !0;
          case Ys:
            return "number" == typeof e;
          case js:
            return "number" == typeof e && (0 | e) === e;
          case String:
            return "string" == typeof e;
          default:
            if (null == e) return !1;
            if ("function" == typeof t) {
              if (e instanceof t) return !0;
              if (G.__interfLoop(G.getClass(e), t)) return !0;
            } else if (
              "object" == typeof t &&
              G.__isNativeObj(t) &&
              e instanceof t
            )
              return !0;
            return t == Ks && null != e.__name__
              ? !0
              : (t == Vs && null != e.__ename__) || e.__enum__ == t;
        }
      }),
      (G.__cast = function (e, t) {
        if (G.__instanceof(e, t)) return e;
        throw new H("Cannot cast " + P.string(e) + " to " + P.string(t));
      }),
      (G.__nativeClassName = function (e) {
        e = G.__toStr.call(e).slice(8, -1);
        return "Object" == e || "Function" == e || "Math" == e || "JSON" == e
          ? null
          : e;
      }),
      (G.__isNativeObj = function (e) {
        return null != G.__nativeClassName(e);
      }),
      (G.__resolveNativeClass = function (e) {
        return S[e];
      }),
      require("fs")),
    ds = require("https"),
    fs = require("os"),
    ms = require("path"),
    ps = require("url"),
    gs = require("buffer").Buffer,
    ys = function (e, t, n) {
      this.decodeStrategy = n;
      n = new as(e);
      (n.set_bigEndian(!0), (this.o = this.decode(n, t)));
    },
    Ss =
      (((t["msgpack.Decoder"] = ys).__name__ = ["msgpack", "Decoder"]),
      (ys.prototype = {
        decode: function (e, t) {
          try {
            var n = e.readByte();
            switch (n) {
              case 192:
                return null;
              case 194:
                return !1;
              case 195:
                return !0;
              case 202:
                return e.readFloat();
              case 203:
                return e.readDouble();
              case 204:
                return e.readByte();
              case 205:
                return e.readUInt16();
              case 206:
                return e.readInt32();
              case 207:
                return _.make(e.readInt32(), e.readInt32());
              case 208:
                return e.readInt8();
              case 209:
                return e.readInt16();
              case 210:
                return e.readInt32();
              case 211:
                return _.make(e.readInt32(), e.readInt32());
              case 218:
              case 219:
                return this.decodeStrategy.parseBytes(
                  e.read(218 == n ? e.readUInt16() : e.readInt32()),
                );
              case 220:
              case 221:
                return this.readArray(
                  e,
                  220 == n ? e.readUInt16() : e.readInt32(),
                  t,
                );
              case 222:
              case 223:
                return this.readMap(
                  e,
                  222 == n ? e.readUInt16() : e.readInt32(),
                  t,
                );
              default:
                if (n < 128) return n;
                if (n < 144) return this.readMap(e, 15 & n, t);
                if (n < 160) return this.readArray(e, 15 & n, t);
                if (n < 192)
                  return this.decodeStrategy.parseBytes(e.read(31 & n));
                if (223 < n) return -256 | n;
            }
          } catch (e) {
            if ((e instanceof H && (e = e.val), !G.__instanceof(e, ls)))
              throw e;
          }
          return null;
        },
        readArray: function (e, t, n) {
          for (var i = [], r = 0, o = t; r < o; ) {
            r++;
            i.push(this.decode(e, n));
          }
          return i;
        },
        readMap: function (e, t, n) {
          if (n) {
            for (var i = {}, r = 0, o = t; r < o; ) {
              r++;
              var a = this.decode(e, n),
                s =
                  (this.decodeStrategy.startDecodingMapValue(a),
                  this.decode(e, n));
              (this.decodeStrategy.endDecodingMapValue(), (i[a] = s));
            }
            return i;
          }
          for (var l = new ns(), u = 0, c = t; u < c; ) {
            u++;
            var _ = this.decode(e, n),
              h =
                (this.decodeStrategy.startDecodingMapValue(_),
                this.decode(e, n));
            (this.decodeStrategy.endDecodingMapValue(), l.set(_, h));
          }
          return l;
        },
        getResult: function () {
          return this.o;
        },
        __class__: ys,
      }),
      function (e) {
        ((this.o = new ss()), this.o.set_bigEndian(!0), this.encode(e));
      }),
    vs =
      (((t["msgpack.Encoder"] = Ss).__name__ = ["msgpack", "Encoder"]),
      (Ss.prototype = {
        encode: function (e) {
          var t = x.typeof(e);
          switch (t[1]) {
            case 0:
              this.o.writeByte(192);
              break;
            case 1:
              var n = e;
              n < -32
                ? n < -32768
                  ? (this.o.writeByte(210), this.o.writeInt32(n))
                  : n < -128
                    ? (this.o.writeByte(209), this.o.writeInt16(n))
                    : (this.o.writeByte(208), this.o.writeInt8(n))
                : n < 128
                  ? this.o.writeByte(255 & n)
                  : n < 256
                    ? (this.o.writeByte(204), this.o.writeByte(n))
                    : n < 65536
                      ? (this.o.writeByte(205), this.o.writeUInt16(n))
                      : (this.o.writeByte(206), this.o.writeInt32(n));
              break;
            case 2:
              var n = e,
                i = Math.abs(n);
              1401298464324817e-60 < i && i < 34028234663852886e22
                ? (this.o.writeByte(202), this.o.writeFloat(n))
                : (this.o.writeByte(203), this.o.writeDouble(n));
              break;
            case 3:
              this.o.writeByte(e ? 195 : 194);
              break;
            case 4:
              for (
                var r = L.fields(e),
                  i = k.count(r),
                  o =
                    (i < 16
                      ? this.o.writeByte(128 | i)
                      : i < 65536
                        ? (this.o.writeByte(222), this.o.writeUInt16(i))
                        : (this.o.writeByte(223), this.o.writeInt32(i)),
                    0);
                o < r.length;
              ) {
                var a = r[o],
                  s = (++o, rs.ofString(a)),
                  l = s.length;
                (l < 32
                  ? this.o.writeByte(160 | l)
                  : l < 65536
                    ? (this.o.writeByte(218), this.o.writeUInt16(l))
                    : (this.o.writeByte(219), this.o.writeInt32(l)),
                  this.o.write(s),
                  this.encode(L.field(e, a)));
              }
              break;
            case 5:
              throw new H("Error: Function not supported");
            case 6:
              switch (x.getClassName(t[2])) {
                case "Array":
                  for (
                    var u = e,
                      c = u.length,
                      _ =
                        (c < 16
                          ? this.o.writeByte(144 | c)
                          : c < 65536
                            ? (this.o.writeByte(220), this.o.writeUInt16(c))
                            : (this.o.writeByte(221), this.o.writeInt32(c)),
                        0);
                    _ < u.length;
                  ) {
                    var h = u[_];
                    (++_, this.encode(h));
                  }
                  break;
                case "Hash":
                  for (
                    var d = e,
                      c = k.count(d),
                      f =
                        (c < 16
                          ? this.o.writeByte(128 | c)
                          : c < 65536
                            ? (this.o.writeByte(222), this.o.writeUInt16(c))
                            : (this.o.writeByte(223), this.o.writeInt32(c)),
                        d.keys());
                    f.hasNext();
                  ) {
                    var m = f.next(),
                      p = rs.ofString(m),
                      g = p.length;
                    (g < 32
                      ? this.o.writeByte(160 | g)
                      : g < 65536
                        ? (this.o.writeByte(218), this.o.writeUInt16(g))
                        : (this.o.writeByte(219), this.o.writeInt32(g)),
                      this.o.write(p),
                      this.encode(null != Y[m] ? d.getReserved(m) : d.h[m]));
                  }
                  break;
                case "String":
                  var c = rs.ofString(e),
                    y = c.length;
                  (y < 32
                    ? this.o.writeByte(160 | y)
                    : y < 65536
                      ? (this.o.writeByte(218), this.o.writeUInt16(y))
                      : (this.o.writeByte(219), this.o.writeInt32(y)),
                    this.o.write(c));
                  break;
                case "haxe._Int64.___Int64":
                  y = e;
                  (this.o.writeByte(y.high < 0 ? 211 : 207),
                    this.o.writeInt32(y.high),
                    this.o.writeInt32(y.low));
                  break;
                case "haxe.io.Bytes":
                  ((c = e), (y = c.length));
                  (y < 32
                    ? this.o.writeByte(160 | y)
                    : y < 65536
                      ? (this.o.writeByte(218), this.o.writeUInt16(y))
                      : (this.o.writeByte(219), this.o.writeInt32(y)),
                    this.o.write(c));
              }
              break;
            case 7:
              throw new H("Error: Enum not supported");
            case 8:
              throw new H("Error: Unknown Data Type");
          }
        },
        writeInt: function (e) {
          e < -32
            ? e < -32768
              ? (this.o.writeByte(210), this.o.writeInt32(e))
              : e < -128
                ? (this.o.writeByte(209), this.o.writeInt16(e))
                : (this.o.writeByte(208), this.o.writeInt8(e))
            : e < 128
              ? this.o.writeByte(255 & e)
              : e < 256
                ? (this.o.writeByte(204), this.o.writeByte(e))
                : e < 65536
                  ? (this.o.writeByte(205), this.o.writeUInt16(e))
                  : (this.o.writeByte(206), this.o.writeInt32(e));
        },
        writeInt64: function (e) {
          (this.o.writeByte(e.high < 0 ? 211 : 207),
            this.o.writeInt32(e.high),
            this.o.writeInt32(e.low));
        },
        writeFloat: function (e) {
          var t = Math.abs(e);
          1401298464324817e-60 < t && t < 34028234663852886e22
            ? (this.o.writeByte(202), this.o.writeFloat(e))
            : (this.o.writeByte(203), this.o.writeDouble(e));
        },
        writeRaw: function (e) {
          var t = e.length;
          (t < 32
            ? this.o.writeByte(160 | t)
            : t < 65536
              ? (this.o.writeByte(218), this.o.writeUInt16(t))
              : (this.o.writeByte(219), this.o.writeInt32(t)),
            this.o.write(e));
        },
        writeArray: function (e) {
          for (
            var t = e.length,
              n =
                (t < 16
                  ? this.o.writeByte(144 | t)
                  : t < 65536
                    ? (this.o.writeByte(220), this.o.writeUInt16(t))
                    : (this.o.writeByte(221), this.o.writeInt32(t)),
                0);
            n < e.length;
          ) {
            var i = e[n];
            (++n, this.encode(i));
          }
        },
        writeMapLength: function (e) {
          e < 16
            ? this.o.writeByte(128 | e)
            : e < 65536
              ? (this.o.writeByte(222), this.o.writeUInt16(e))
              : (this.o.writeByte(223), this.o.writeInt32(e));
        },
        writeHashMap: function (e) {
          for (
            var t = k.count(e),
              n =
                (t < 16
                  ? this.o.writeByte(128 | t)
                  : t < 65536
                    ? (this.o.writeByte(222), this.o.writeUInt16(t))
                    : (this.o.writeByte(223), this.o.writeInt32(t)),
                e.keys());
            n.hasNext();
          ) {
            var i = n.next(),
              r = rs.ofString(i),
              o = r.length;
            (o < 32
              ? this.o.writeByte(160 | o)
              : o < 65536
                ? (this.o.writeByte(218), this.o.writeUInt16(o))
                : (this.o.writeByte(219), this.o.writeInt32(o)),
              this.o.write(r),
              this.encode(null != Y[i] ? e.getReserved(i) : e.h[i]));
          }
        },
        writeObjectMap: function (e) {
          for (
            var t = L.fields(e),
              n = k.count(t),
              i =
                (n < 16
                  ? this.o.writeByte(128 | n)
                  : n < 65536
                    ? (this.o.writeByte(222), this.o.writeUInt16(n))
                    : (this.o.writeByte(223), this.o.writeInt32(n)),
                0);
            i < t.length;
          ) {
            var r = t[i],
              o = (++i, rs.ofString(r)),
              a = o.length;
            (a < 32
              ? this.o.writeByte(160 | a)
              : a < 65536
                ? (this.o.writeByte(218), this.o.writeUInt16(a))
                : (this.o.writeByte(219), this.o.writeInt32(a)),
              this.o.write(o),
              this.encode(L.field(e, r)));
          }
        },
        getBytes: function () {
          return this.o.getBytes();
        },
        __class__: Ss,
      }),
      ((t["msgpack.MsgPack"] = fn).__name__ = ["msgpack", "MsgPack"]),
      (fn.encode = function (e) {
        return new Ss(e).o.getBytes();
      }),
      (fn.decode = function (e, t, n) {
        return new ys(e, t, n).o;
      }),
      ((t["puremvc.interfaces.IController"] = mn).__name__ = [
        "puremvc",
        "interfaces",
        "IController",
      ]),
      (mn.prototype = { __class__: mn }),
      function () {
        (((vs.instance = this).commandMap = new B()),
          this.initializeController());
      }),
    Ts =
      (((t["puremvc.core.Controller"] = vs).__name__ = [
        "puremvc",
        "core",
        "Controller",
      ]),
      (vs.__interfaces__ = [mn]),
      (vs.getInstance = function () {
        return (null == vs.instance && (vs.instance = new vs()), vs.instance);
      }),
      (vs.prototype = {
        initializeController: function () {
          this.view = Is.getInstance();
        },
        executeCommand: function (e) {
          var t = this.commandMap,
            n = e.getName(),
            t = null != Y[n] ? t.getReserved(n) : t.h[n];
          null != t && x.createInstance(t, []).execute(e);
        },
        registerCommand: function (e, t) {
          var n = this.commandMap,
            n =
              ((null != Y[e] ? n.existsReserved(e) : n.h.hasOwnProperty(e)) ||
                this.view.registerObserver(
                  e,
                  new ws(z(this, this.executeCommand), this),
                ),
              this.commandMap);
          null != Y[e] ? n.setReserved(e, t) : (n.h[e] = t);
        },
        hasCommand: function (e) {
          var t = this.commandMap;
          return null != Y[e] ? t.existsReserved(e) : t.h.hasOwnProperty(e);
        },
        removeCommand: function (e) {
          this.hasCommand(e) &&
            (this.view.removeObserver(e, this), this.commandMap.remove(e));
        },
        __class__: vs,
      }),
      ((t["puremvc.interfaces.IModel"] = pn).__name__ = [
        "puremvc",
        "interfaces",
        "IModel",
      ]),
      (pn.prototype = { __class__: pn }),
      function () {
        (((Ts.instance = this).proxyMap = new B()), this.initializeModel());
      }),
    Is =
      (((t["puremvc.core.Model"] = Ts).__name__ = ["puremvc", "core", "Model"]),
      (Ts.__interfaces__ = [pn]),
      (Ts.getInstance = function () {
        return (null == Ts.instance && (Ts.instance = new Ts()), Ts.instance);
      }),
      (Ts.prototype = {
        initializeModel: function () {},
        registerProxy: function (e) {
          var t = this.proxyMap,
            n = e.getProxyName();
          (null != Y[n] ? t.setReserved(n, e) : (t.h[n] = e), e.onRegister());
        },
        retrieveProxy: function (e) {
          var t = this.proxyMap;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        hasProxy: function (e) {
          var t = this.proxyMap;
          return null != Y[e] ? t.existsReserved(e) : t.h.hasOwnProperty(e);
        },
        removeProxy: function (e) {
          var t = this.proxyMap,
            t = null != Y[e] ? t.getReserved(e) : t.h[e];
          return (null != t && (this.proxyMap.remove(e), t.onRemove()), t);
        },
        __class__: Ts,
      }),
      ((t["puremvc.interfaces.IView"] = gn).__name__ = [
        "puremvc",
        "interfaces",
        "IView",
      ]),
      (gn.prototype = { __class__: gn }),
      function () {
        (((Is.instance = this).mediatorMap = new B()),
          (this.observerMap = new B()),
          this.initializeView());
      }),
    Es =
      (((t["puremvc.core.View"] = Is).__name__ = ["puremvc", "core", "View"]),
      (Is.__interfaces__ = [gn]),
      (Is.getInstance = function () {
        return (null == Is.instance && (Is.instance = new Is()), Is.instance);
      }),
      (Is.prototype = {
        initializeView: function () {},
        registerObserver: function (e, t) {
          var n,
            i = this.observerMap,
            i =
              ((null != Y[e] ? i.existsReserved(e) : i.h.hasOwnProperty(e)) ||
                ((i = this.observerMap),
                (n = new En()),
                null != Y[e] ? i.setReserved(e, n) : (i.h[e] = n)),
              this.observerMap);
          (null != Y[e] ? i.getReserved(e) : i.h[e]).add(t);
        },
        notifyObservers: function (e) {
          var t = this.observerMap,
            n = e.getName();
          if (null != Y[n] ? t.existsReserved(n) : t.h.hasOwnProperty(n)) {
            for (
              var t = this.observerMap,
                n = e.getName(),
                t = null != Y[n] ? t.getReserved(n) : t.h[n],
                i = new En(),
                r = new Nn(t.h);
              r.hasNext();
            ) {
              var o = r.next();
              i.add(o);
            }
            for (var a = new Nn(i.h); a.hasNext(); ) a.next().notifyObserver(e);
          }
        },
        removeObserver: function (e, t) {
          for (
            var n = this.observerMap,
              i = null != Y[e] ? n.getReserved(e) : n.h[e],
              r = i.h;
            null != r;
          ) {
            var o = r.item,
              r = r.next;
            if (1 == o.compareNotifyContext(t)) {
              i.remove(o);
              break;
            }
          }
          i.isEmpty() && this.observerMap.remove(e);
        },
        registerMediator: function (e) {
          var t = this.mediatorMap,
            n = e.getMediatorName();
          if (null != Y[n] ? !t.existsReserved(n) : !t.h.hasOwnProperty(n)) {
            var t = this.mediatorMap,
              n = e.getMediatorName(),
              i =
                (null != Y[n] ? t.setReserved(n, e) : (t.h[n] = e),
                e.listNotificationInterests());
            if (0 < i.length)
              for (
                var r = new ws(z(e, e.handleNotification), e),
                  o = 0,
                  a = i.length;
                o < a;
              ) {
                var s = o++;
                this.registerObserver(i[s], r);
              }
            e.onRegister();
          }
        },
        retrieveMediator: function (e) {
          var t = this.mediatorMap;
          return null != Y[e] ? t.getReserved(e) : t.h[e];
        },
        removeMediator: function (e) {
          var t = this.mediatorMap,
            n = null != Y[e] ? t.getReserved(e) : t.h[e];
          if (null != n) {
            for (
              var i = n.listNotificationInterests(), r = 0, o = i.length;
              r < o;
            ) {
              var a = r++;
              this.removeObserver(i[a], n);
            }
            (this.mediatorMap.remove(e), n.onRemove());
          }
          return n;
        },
        hasMediator: function (e) {
          var t = this.mediatorMap;
          return null != Y[e] ? t.existsReserved(e) : t.h.hasOwnProperty(e);
        },
        __class__: Is,
      }),
      ((t["puremvc.interfaces.INotification"] = yn).__name__ = [
        "puremvc",
        "interfaces",
        "INotification",
      ]),
      (yn.prototype = { __class__: yn }),
      ((t["puremvc.interfaces.IObserver"] = Sn).__name__ = [
        "puremvc",
        "interfaces",
        "IObserver",
      ]),
      (Sn.prototype = { __class__: Sn }),
      function (e, t, n) {
        ((this.name = e),
          null != t && (this.body = t),
          null != n && (this.type = n));
      }),
    ws =
      (((t["puremvc.patterns.observer.Notification"] = Es).__name__ = [
        "puremvc",
        "patterns",
        "observer",
        "Notification",
      ]),
      (Es.__interfaces__ = [yn]),
      (Es.prototype = {
        getName: function () {
          return this.name;
        },
        setBody: function (e) {
          this.body = e;
        },
        getBody: function () {
          return this.body;
        },
        setType: function (e) {
          this.type = e;
        },
        getType: function () {
          return this.type;
        },
        toString: function () {
          var e = "Notification Name: " + this.getName();
          return (
            (e +=
              "\nBody:" + (null == this.body ? "null" : this.body.toString())) +
            ("\nType:" + (null == this.type ? "null" : this.type))
          );
        },
        __class__: Es,
      }),
      function (e, t) {
        (this.setNotifyMethod(e), this.setNotifyContext(t));
      });
  function Ns(e) {
    return e instanceof Array
      ? function () {
          return U.iter(e);
        }
      : "function" == typeof e.iterator
        ? z(e, e.iterator)
        : e.iterator;
  }
  (((t["puremvc.patterns.observer.Observer"] = ws).__name__ = [
    "puremvc",
    "patterns",
    "observer",
    "Observer",
  ]),
    (ws.__interfaces__ = [Sn]),
    (ws.prototype = {
      setNotifyMethod: function (e) {
        this.notify = e;
      },
      setNotifyContext: function (e) {
        this.context = e;
      },
      getNotifyMethod: function () {
        return this.notify;
      },
      getNotifyContext: function () {
        return this.context;
      },
      notifyObserver: function (e) {
        this.getNotifyMethod()(e);
      },
      compareNotifyContext: function (e) {
        return e == this.context;
      },
      __class__: ws,
    }));
  var j,
    As = 0;
  function z(e, t) {
    var n;
    return null == t
      ? null
      : (null == t.__id__ && (t.__id__ = As++),
        null == e.hx__closures__
          ? (e.hx__closures__ = {})
          : (n = e.hx__closures__[t.__id__]),
        null == n &&
          (((n = function () {
            return n.method.apply(n.scope, arguments);
          }).scope = e),
          (n.method = t),
          (e.hx__closures__[t.__id__] = n)),
        n);
  }
  t.Math = Math;
  var bs,
    Ms,
    Cs,
    ks,
    Ds,
    Rs,
    Os,
    Fs,
    xs,
    Us,
    Ls,
    Ps,
    Bs,
    Hs,
    $n = {},
    Gs =
      (($n.exports =
        ((bs = Object.defineProperty),
        (Ms = Object.getOwnPropertyDescriptor),
        (Cs = Object.getOwnPropertyNames),
        (ks = Object.prototype.hasOwnProperty),
        ((e, t) => {
          for (var n in t) bs(e, n, { get: t[n], enumerable: !0 });
        })((y = {}), {
          codePointCount: () => Rs,
          extractMentionParts: () => Ps,
          findMentionSuggestionHint: () => Hs,
          hasMentionPartForMeOrAll: () => Fs,
          invalidateManualMentionMarkup: () => Us,
          parseMessageMention: () => Ls,
          removeMentionMarkup: () => xs,
        }),
        (Ds = (e) => Array.from(e)),
        (Rs = (e) => Ds(e).length),
        (Os = /{@:(ALL|\d+),(\d+)}@/g),
        (Fs = (e) =>
          Ps(e).some(({ userIdNumStr: e }) => null === e || "ALL" === e)),
        (xs = (t) =>
          Ls(t)
            .map((e) =>
              "mention" === e.type
                ? e.mentionText
                : "normal" === e.type
                  ? t.substring(e.start, e.end)
                  : "",
            )
            .join("")),
        (Us = (e) => e.replace(Os, (e) => e.substring(0, e.length - 1) + " ")),
        (Ls = (e) => {
          let n = 0;
          var t = Ps(e).flatMap((e) => {
            var t = [];
            return (
              n !== e.start &&
                t.push({ type: "normal", start: n, end: e.start }),
              t.push(e),
              (n = e.end),
              t
            );
          });
          return (
            e.length !== n &&
              t.push({ type: "normal", start: n, end: e.length }),
            t
          );
        }),
        (Ps = (o) => {
          let a = 0;
          return Array.from(o.matchAll(Os)).flatMap((e) => {
            var t,
              n,
              i,
              r = null != (r = e.index) ? r : 0;
            return r < a ||
              ((t = r + e[0].length), (n = e[1]), (e = parseInt(e[2])) < 2) ||
              ((i = o.substring(t, t + 2 * (e = e - 1))),
              (i = Ds(i).slice(0, e)).length !== e)
              ? []
              : ((r = {
                  type: "mention",
                  start: r,
                  end: t + (e = i.join("")).length,
                  userIdNumStr: n,
                  mentionText: "@" + e,
                }),
                (a = r.end),
                r);
          });
        }),
        (Bs = /^(.*?)([@＠]?[@＠])([^@＠\s]*)$/),
        (Hs = (e) => {
          e = e.match(Bs);
          return null === e
            ? null
            : { prev: e[1], atMarks: e[2], keyword: e[3] };
        }),
        (y = y),
        ((t, n, i, r) => {
          if ((n && "object" == typeof n) || "function" == typeof n)
            for (let e of Cs(n))
              ks.call(t, e) ||
                e === i ||
                bs(t, e, {
                  get: () => n[e],
                  enumerable: !(r = Ms(n, e)) || r.enumerable,
                });
          return t;
        })(bs({}, "__esModule", { value: !0 }), y))),
      $n.exports),
    js =
      (((String.prototype.__class__ = t.String = String).__name__ = ["String"]),
      ((t.Array = Array).__name__ = ["Array"]),
      ((Date.prototype.__class__ = t.Date = Date).__name__ = ["Date"]),
      (t.Int = { __name__: ["Int"] })),
    zs = (t.Dynamic = { __name__: ["Dynamic"] }),
    Ys = (t.Float = Number),
    Ws = ((Ys.__name__ = ["Float"]), (t.Bool = Boolean)),
    Ks = ((Ws.__ename__ = ["Bool"]), (t.Class = { __name__: ["Class"] })),
    Vs = {},
    qs = {},
    y = {},
    Qs =
      ((y.exports = (() => {
        var e = require("fs");
        let r = require("libxmljs");
        var t = require("path");
        let n = (e) => e.namespace()?.href() ?? "",
          a = (e) => `{${n(e)}}` + e.name();
        ((t = t.resolve(__dirname, "./note-richtext.v1.xsd")),
          (e = r.parseXml(e.readFileSync(t, { encoding: "utf-8" }), {
            baseUrl: t,
          })));
        let i = new Map();
        function o(e) {
          var t = Number(e.root().getAttribute("version")?.value()),
            n = i.get(t);
          return n
            ? e.validate(n)
              ? { ok: !0 }
              : {
                  ok: !1,
                  errorType: "xsd",
                  message: e.validationErrors
                    .map((e) => e.message.trim())
                    .join(" "),
                }
            : {
                ok: !1,
                errorType: "version",
                message:
                  `The note version ${t} is not supported by this package. Please use one of the following values: ` +
                  Array.from(i.keys()).join(","),
              };
        }
        i.set(1, e);
        let s = new Map();
        function l(e) {
          try {
            return (
              (function n(e, i, r) {
                for (var o of e.childNodes())
                  if ("element" === o.type()) {
                    if (!i.includes(o.name()))
                      throw new Error(
                        `Element '${a(o)}': This element is not allowed at '${o.path()}'.`,
                      );
                    {
                      let t = r.get(o.name()) ?? [],
                        e = [...i].filter((e) => t.includes(e)) ?? [];
                      n(o, e, r);
                    }
                  }
              })(e.root(), s.get(e.root().name()), s),
              { ok: !0 }
            );
          } catch (e) {
            return { ok: !1, errorType: "descendant", message: e.message };
          }
        }
        (s.set("note", [
          "textsize",
          "color",
          "b",
          "link",
          "u",
          "strike",
          "checkbox",
        ]),
          s.set("textsize", ["color", "b", "link", "u", "strike"]),
          s.set("color", ["textsize", "b", "link", "u", "strike"]),
          s.set("b", ["textsize", "color", "link", "u", "strike"]),
          s.set("link", ["textsize", "b", "strike"]),
          s.set("u", ["textsize", "color", "b", "link", "strike"]),
          s.set("strike", ["textsize", "color", "b", "link", "u"]),
          s.set("checkbox", ["color", "b", "link", "u", "strike"]));
        let u = (e) => "element" === e.type(),
          c = (e) => "element" === e.type() && "checkbox" === e.name(),
          _ = (e) => "element" === e.type() && "link" === e.name(),
          h = (e) =>
            "element" === e.type() &&
            ["textsize", "color", "b", "link", "u", "strike"].includes(
              e.name(),
            );
        function d(e, t, n, i) {
          for (var r of e.childNodes())
            t(r) && (n(r) && i.push(r), d(r, t, n, i));
        }
        function f(e) {
          var t = [];
          d(e.root(), u, _, t);
          try {
            for (var n of t) {
              var i = n.text();
              if (null == i || "" === i)
                throw new Error(
                  `Element '${a(n)}': This element must not be empty. (element path: ${n.path()})`,
                );
              if (/\r|\n/.test(i))
                throw new Error(
                  `Element '${a(n)}': This element must not contain a new line. (element path: ${n.path()})`,
                );
              var r = n
                .attrs()
                .filter((e) => "url" === e.name())[0]
                .value();
              if (!/^[\x20-\x7E]+$/.test(r))
                throw new Error(
                  `Element '${a(n)}': The attribute 'url' must be URL-encoded or must be converted by Punycode if it's an IDN. (element path: ${n.path()})`,
                );
            }
            return { ok: !0 };
          } catch (e) {
            return { ok: !1, errorType: "link", message: e.message };
          }
        }
        function m(e) {
          var t = [];
          d(e.root(), u, c, t);
          try {
            for (var n of t)
              if (/\r|\n/.test(n.text()))
                throw new Error(
                  `Element '${a(n)}': This element must not contain a new line. (element path: ${n.path()})`,
                );
            return { ok: !0 };
          } catch (e) {
            return { ok: !1, errorType: "checkbox", message: e.message };
          }
        }
        function p(e) {
          var t = [];
          d(e.root(), u, h, t);
          try {
            for (var n of t) {
              var i = n.text();
              if (null == i || "" === i)
                throw new Error(
                  `Element '${a(n)}': This element must not be empty. (element path: ${n.path()})`,
                );
            }
            return { ok: !0 };
          } catch (e) {
            return { ok: !1, errorType: "text-decorator", message: e.message };
          }
        }
        return {
          validateXml: function (e) {
            let t;
            try {
              t = r.parseXml(e).doc();
            } catch (e) {
              return { ok: !1, errorType: "parse", message: e.message };
            }
            let n = { ok: !0 };
            var i;
            for (i of [o, l, f, m, p]) if (!(n = i(t)).ok) break;
            return n;
          },
        };
      })()),
      y.exports),
    Y = {};
  ((i.APP_STATE_CHANGED = "app_state_changed"),
    (i.ACCESS_TOKEN_CHANGED = "access_token_changed"),
    (i.CURRENT_USER_CHANGED = "current_user_changed"),
    (i.CONFIGURATION_CHANGED = "configuration_changed"),
    (i.DOMAIN_SELECTION_CHANGED = "domain_selection_changed"),
    (i.TALK_SELECTION_CHANGED = "talk_selection_changed"),
    (i.TALK_LIST_SCROLL_TO_TALK_TOP_NEEDED =
      "talk_list_scroll_to_talk_top_needed"),
    (i.USER_SELECTION_NEEDED = "user_selection_needed"),
    (i.USER_SELECTION_CHANGED = "user_selection_changed"),
    (i.FRIEND_SELECTION_NEEDED = "friend_selection_needed"),
    (i.COMMON_STAMP_SET_LOADED = "common_stamp_set_loaded"),
    (i.STAMP_SELECTION_STARTED = "stamp_selection_started"),
    (i.STAMP_SELECTION_ENDED = "stamp_selection_ended"),
    (i.STAMP_SELECTION_CHANGED = "stamp_selection_changed"),
    (i.ACTION_SELECTION_CHANGED = "action_selection_changed"),
    (i.CURRENT_PAGE_CHANGED = "current_page_changed"),
    (i.CURRENT_PAGE_REASSIGNED = "current_page_reassgined"),
    (i.FILEINFO_SELECTION_CHANGED = "fileinfo_selection_changed"),
    (i.NOTE_FILEINFO_SELECTION_CHANGED = "note_fileinfo_selection_changed"),
    (i.MESSAGE_FILEINFO_SELECTION_CHANGED =
      "message_fileinfo_selection_changed"),
    (i.STAGED_FILEINFO_SELECTION_CHANGED = "staged_fileinfo_selection_changed"),
    (i.STAGED_FILEINFOS_ADDED = "staged_fileinfos_added"),
    (i.STAGED_FILEINFOS_MOVED = "staged_fileinfos_moved"),
    (i.CLOSE_FILEPREVIEW_MODAL = "close_filepreview_modal"),
    (i.ERROR_OCCURRED = "error_occurred"),
    (i.UNREAD_COUNT_CHANGED = "brand_badge_changed"),
    (i.SEND_FORM_TOP_CHANGED = "send_form_top_changed"),
    (i.SOLUTIONS_LOADED = "solutions_loaded"),
    (i.RIGHT_PANE_OPENED = "right_pane_opened"),
    (i.RIGHT_PANE_CLOSED = "right_pane_closed"),
    (i.PASSWORD_EXPIRATION_OVERED = "password_expiration_overed"),
    (i.PASSWORD_EXPIRATION_WARNED = "password_expiration_warned"),
    (i.PHOTO_EDITOR_SAVED = "photo_editor_saved"),
    (i.ICON_EDITOR_SAVED = "icon_editor_saved"),
    (i.MC_AUTHENTICATED_USER_RECEIVED = "mc_authenticated_user_received"),
    (i.KEYWORD_WATCHING_UPDATED = "keyword_watching_updated"),
    (i.KEYWORD_DETECTION_UPDATED = "keyword_detaction_updated"),
    (i.DEPARTMENT_SELECTION_CHANGED_FOR_MEMBERS_PAGE =
      "department_selection_changed_for_members_page"),
    (i.PRESENCES_UPDATED = "presences_updated"),
    (i.NOTEINFO_SELECTION_CHANGED = "noteinfo_selection_changed"),
    (i.START_NOTE_EDITING = "start_note_editing"),
    (i.OPEN_EXISTING_PAIR_TALK_COMPLETED = "open_existing_pair_talk_completed"),
    (i.FIRST_ROUTING_WILL_START = "first_routing_will_start"),
    (i.SIGNED_OUT = "signed_out"),
    (i.ROUTUNG_TALK_FALLBACK = "routing_talk_fallback"),
    (i.ROUTUNG_DOMAIN_FALLBACK = "routing_domain_fallback"),
    (i.TALKS_SORT_MODE_CHANGED = "talks_sort_mode_changed"),
    (i.RESERVE_SCHEDULED_MESSAGE_EDITING = "reserve_scheduled_message_editing"),
    (i.DATA_RECOVERING = "data_recovering"),
    (i.DATA_RECOVERED = "data_recovered"),
    (i.NOTIFY_UPDATE_USER = "notify_update_user"),
    (i.NOTIFY_ADD_FRIEND = "notify_add_friend"),
    (i.NOTIFY_ADD_ACQUAINTANCE = "notify_add_acquaintance"),
    (i.NOTIFY_ADD_ACQUAINTANCES = "notify_add_acquaintances"),
    (i.NOTIFY_DELETE_FRIEND = "notify_delete_friend"),
    (i.NOTIFY_DELETE_ACQUAINTANCE = "notify_delete_acquaintance"),
    (i.NOTIFY_DELETE_ACQUAINTANCES = "notify_delete_acquaintances"),
    (i.NOTIFY_UPDATE_DOMAIN_USERS = "notify_update_domain_users"),
    (i.GET_DOMAIN_USERS_RESPONSED = "get_domain_users_responsed"),
    (i.GET_USERS_RESPONSED = "get_users_responsed"),
    (i.GET_PROFILE_RESPONSED = "get_profile_responsed"),
    (i.GET_PROFILE_ERRORED = "get_profile_errored"),
    (i.UPDATE_USER_RESPONSED = "update_user_responsed"),
    (i.UPDATE_USER_ERRORED = "update_user_errored"),
    (i.UPDATE_PROFILE_RESPONSED = "update_profile_responsed"),
    (i.UPDATE_PROFILE_ERRORED = "update_profile_errored"),
    (i.NOTIFY_UPDATE_DEPARTMENT_TREE = "notify_update_department_tree"),
    (i.NOTIFY_UPDATE_DEPARTMENT_USERS = "notify_update_department_users"),
    (i.GET_DEPARTMENT_TREE_RESPONSED = "get_department_tree_responsed"),
    (i.GET_DEPARTMENT_TREE_CANCELED = "get_department_tree_canceled"),
    (i.GET_DEPARTMENT_USERS_RESPONSED = "get_department_users_responsed"),
    (i.GET_DEPARTMENT_USERS_CANCELED = "get_department_users_canceled"),
    (i.GET_DEPARTMENT_USER_COUNT_RESPONSED =
      "get_department_user_count_responsed"),
    (i.GET_ME_RESPONSED = "get_me_responsed"),
    (i.DEPARTMENT_USER_COUNT_CLEARED = "department_user_count_cleared"),
    (i.DEPARTMENT_USER_IDS_PREPARED = "department_user_ids_prepared"),
    (i.NOTIFY_ADD_DOMAIN_INVITE = "notify_add_domain_invite"),
    (i.NOTIFY_ACCEPT_DOMAIN_INVITE = "notify_accept_domain_invite"),
    (i.NOTIFY_DELETE_DOMAIN_INVITE = "notify_delete_domain_invite"),
    (i.NOTIFY_JOIN_DOMAIN = "notify_join_domain"),
    (i.NOTIFY_UPDATE_DOMAIN = "notify_update_domain"),
    (i.NOTIFY_LEAVE_DOMAIN = "notify_leave_domain"),
    (i.NOTIFY_ADD_DOMAIN_MEMBERS = "notify_add_domain_members"),
    (i.NOTIFY_CREATE_PAIR_TALK = "notify_create_pair_talk"),
    (i.CREATE_PAIR_TALK_COMPLETE = "create_pair_talk_complete"),
    (i.CREATE_PAIR_TALK_FAIL = "create_pair_talk_fail"),
    (i.NOTIFY_CREATE_GROUP_TALK = "notify_create_group_talk"),
    (i.CREATE_GROUP_TALK_COMPLETE = "create_group_talk_complete"),
    (i.CREATE_GROUP_TALK_FAIL = "create_group_talk_fail"),
    (i.NOTIFY_UPDATE_GROUP_TALK = "notify_update_group_talk"),
    (i.UPDATE_GROUP_TALK_COMPLETE = "update_group_talk_complete"),
    (i.UPDATE_GROUP_TALK_FAIL = "update_group_talk_fail"),
    (i.NOTIFY_ADD_TALKERS = "notify_add_talkers"),
    (i.NOTIFY_ADD_TALKERS_INCLUDING_ME = "notify_add_talkers_including_me"),
    (i.NOTIFY_DELETE_TALKER = "notify_delete_talker"),
    (i.NOTIFY_DELETE_TALK = "notify_delete_talk"),
    (i.NOTIFY_UPDATE_READ_STATUSES = "notify_update_read_statuses"),
    (i.NOTIFY_UPDATE_TALK_STATUS = "notify_update_talk_status"),
    (i.NOTIFY_UPDATE_LOCAL_TALK_STATUS = "notify_update_local_talk_status"),
    (i.NOTIFY_ADD_FAVORITE_TALK = "notify_add_favorite_talk"),
    (i.NOTIFY_DELETE_FAVORITE_TALK = "notify_delete_favorite_talk"),
    (i.NOTIFY_DISABLE_PUSH_NOTIFICATION = "notify_disable_push_notification"),
    (i.NOTIFY_ENABLE_PUSH_NOTIFICATION = "notify_enable_push_notification"),
    (i.ADD_TALKERS_SUCCEEDED = "add_talkers_succeeded"),
    (i.ADD_TALKERS_FAILED = "add_talkers_failed"),
    (i.NOTIFY_CREATE_MESSAGE = "notify_create_message"),
    (i.NOTIFY_DELETE_MESSAGE = "notify_delete_message"),
    (i.NOTIFY_GET_MESSAGES = "notify_get_messages"),
    (i.NOTIFY_GET_MESSAGE_READ_STATUS = "notify_get_message_status"),
    (i.CREATE_MESSAGE_START = "create_message_start"),
    (i.CREATE_MESSAGE_COMPLETE = "create_message_complete"),
    (i.CREATE_MESSAGE_FAIL = "create_message_fail"),
    (i.ADD_FAVORITE_MESSAGE_COMPLETED = "add_favorite_message_completed"),
    (i.DELETE_FAVORITE_MESSAGE_COMPLETED = "delete_favorite_message_completed"),
    (i.GET_FAVORITE_MESSAGES_COMPLETED = "get_favorite_messages_completed"),
    (i.NOTIFY_ADD_FAVORITE_MESSAGE = "notify_add_favorite_message"),
    (i.NOTIFY_DELETE_FAVORITE_MESSAGE = "notify_delete_favorite_message"),
    (i.NOTIFY_SCHEDULE_MESSAGE = "notify_schedule_message"),
    (i.NOTIFY_PROCESSED_SCHEDULED_MESSAGE =
      "notify_processed_scheduled_message"),
    (i.NOTIFY_DELETE_SCHEDULED_MESSAGE = "notify_delete_scheduled_message"),
    (i.GET_SCHEDULED_MESSAGES_COMPLETED = "get_scheduled_messages_completed"),
    (i.GET_SCHEDULED_MESSAGES_FAILED = "get_scheduled_messages_failed"),
    (i.DELETE_SCHEDULED_MESSAGES_COMPLETED =
      "delete_scheduled_messages_completed"),
    (i.SCHEDULE_MESSAGE_STARTED = "schedule_message_started"),
    (i.SCHEDULE_MESSAGE_COMPLETED = "schedule_message_completed"),
    (i.SCHEDULE_MESSAGE_FAILED = "schedule_message_faileed"),
    (i.RESCHEDULE_MESSAGE_STARTED = "reschedule_message_started"),
    (i.RESCHEDULE_MESSAGE_COMPLETED = "reschedule_message_completed"),
    (i.RESCHEDULE_MESSAGE_FAILED = "reschedule_message_failed"),
    (i.NOTIFY_UPDATE_MESSAGE_REACTIONS = "notify_update_message_reactions"),
    (i.GET_MESSAGE_REACTION_SUMMARIES_COMPLETED =
      "get_message_reaction_summaries_completed"),
    (i.GET_MESSAGE_REACTION_SUMMARIES_FAILED =
      "get_message_reaction_summaries_failed"),
    (i.SET_MESSAGE_REACTION_COMPLETED = "set_message_reaction_completed"),
    (i.SET_MESSAGE_REACTION_FAILED = "set_message_reaction_failed"),
    (i.RESET_MESSAGE_REACTION_COMPLETED = "reset_message_reaction_completed"),
    (i.RESET_MESSAGE_REACTION_FAILED = "reset_message_reaction_failed"),
    (i.NOTIFY_CREATE_ANNOUNCEMENT = "notify_create_announcement"),
    (i.NOTIFY_DELETE_ANNOUNCEMENT = "notify_delete_announcement"),
    (i.NOTIFY_UPDATE_ANNOUNCEMENT_STATUS = "notify_update_announcement_status"),
    (i.NOTIFY_GET_ANNOUNCEMENTS = "notify_get_announcements"),
    (i.CREATE_ANNOUNCEMENT_START = "create_announcement_start"),
    (i.CREATE_ANNOUNCEMENT_COMPLETE = "create_announcement_complete"),
    (i.CREATE_ANNOUNCEMENT_FAIL = "create_announcement_fail"),
    (i.NOTIFY_UPDATE_QUESTION = "notify_update_question"),
    (i.NOTIFY_GET_QUESTIONS = "get_questions_responsed"),
    (i.NOTIFY_CREATE_ATTACHMENT = "notify_create_attachment"),
    (i.NOTIFY_DELETE_ATTACHMENT = "notify_delete_attachment"),
    (i.GET_FILE_RESPONSED = "get_file_responsed"),
    (i.GET_FILE_PREVIEW = "get_file_preview"),
    (i.CREATE_FILE_PREVIEW = "create_file_preview"),
    (i.GET_NOTE_STATUSES_LOADED = "get_note_statuses_loaded"),
    (i.GET_NOTE_LOADED = "get_note_loaded"),
    (i.GET_NOTE_FAILED_BY_NOTE_NOT_FOUND = "get_note_failed_by_note_not_found"),
    (i.CREATE_NOTE_COMPLETED = "create_note_completed"),
    (i.CREATE_NOTE_FAILED = "create_note_failed"),
    (i.UPDATE_NOTE_LOCAL_EDIT = "update_note_local_edit"),
    (i.CLEAR_NOTE_LOCAL_EDIT = "clear_note_local_edit"),
    (i.UPDATE_NOTE_SETTING_COMPLETED = "update_note_setting_completed"),
    (i.UPDATE_NOTE_SETTING_FAILED_BY_NOT_FOUND =
      "update_note_setting_failed_by_not_found"),
    (i.UPDATE_NOTE_SETTING_FAILED_BY_CONFLICT =
      "update_note_setting_failed_by_conflict"),
    (i.UPDATE_NOTE_SETTING_FAILED_BY_EDITING =
      "update_note_setting_failed_by_editing"),
    (i.UPDATE_NOTE_COMPLETED = "update_note_completed"),
    (i.DELETE_NOTE_COMPLETED = "delete_note_completed"),
    (i.NOTIFY_CREATE_NOTE = "notify_create_note"),
    (i.NOTIFY_UPDATE_NOTE_PARTIALLY = "notify_update_note_partially"),
    (i.NOTIFY_DELETE_NOTE = "notify_delete_note"),
    (i.NOTIFY_UPDATE_NOTE_FOR_SETTING = "notify_update_note_for_setting"),
    (i.NOTIFY_UPDATE_NOTE_FOR_SUMMARY = "notify_update_note_for_revision"),
    (i.CREATE_NOTE_BUTTON_CLICKED = "create_note_button_clicked"),
    (i.NOTIFY_LOCK_NOTE = "notify_lock_note"),
    (i.NOTIFY_UNLOCK_NOTE = "notify_unlock_note"),
    (i.NOTIFY_SEARCH_MESSAGES = "notify_search_messages"),
    (i.NOTIFY_SEARCH_ATTACHMENTS = "notify_search_attachments"),
    (i.NOTIFY_SEARCH_MESSAGES_FAIL = "notify_search_messages_fail"),
    (i.NOTIFY_SEARCH_ATTACHMENTS_FAIL = "notify_search_attachments_fail"),
    (i.NOTIFY_SEARCH_PREPARE = "notify_search_prepare"),
    (i.NOTIFY_SEARCH_CLEAR = "notify_search_clear"),
    (i.NOTIFY_FILTER_BOX_TEXT_CHANGED = "notify_filter_box_text_changed"),
    (i.NOTIFY_SEARCH_BOX_POPUP = "notify_search_box_popup"),
    (i.NOTIFY_ADD_ACCOUNT_CONTROL_REQUEST =
      "notify_add_account_control_request"),
    (i.NOTIFY_DELETE_ACCOUNT_CONTROL_REQUEST =
      "notify_delete_account_control_request"),
    (i.NOTIFY_JOIN_ACCOUNT_CONTROL_GROUP = "notify_join_account_control_group"),
    (i.NOTIFY_UPDATE_ACCOUNT_CONTROL_GROUP_PARTIALLY =
      "notify_update_account_control_group_partially"),
    (i.NOTIFY_LEAVE_ACCOUNT_CONTROL_GROUP =
      "notify_leave_account_control_group"),
    (i.PREPARE_CONFERENCE_FROM_MESSAGE = "prepare_conference_from_message"),
    (i.JOIN_CONFERENCE_RESPONSED = "join_conference_responsed"),
    (i.JOIN_CONFERENCE_CANCELED = "join_conference_canceled"),
    (i.GET_CONFERENCE_PARTICIPANTS_RESPONSED =
      "get_conference_participants_responsed"),
    (i.NOTIFY_OPEN_CONFERENCE = "notify_open_conference"),
    (i.NOTIFY_CLOSE_CONFERENCE = "notify_close_conference"),
    (i.NOTIFY_CONFERENCE_PARTICIPANT_JOIN =
      "notify_conference_participant_join"),
    (i.NOTIFY_CONFERENCE_PARTICIPANT_LIMIT =
      "notify_conference_participant_limit"),
    (i.NOTIFY_CONFERENCE_PARTICIPANT_REJECT =
      "notify_conference_participant_reject"),
    (i.NOTIFY_ALL_TALK_MEMBERS_REJECTED_CONFERENCE =
      "notify_all_talk_members_rejected_conference"),
    (i.NOTIFY_UPDATE_DOMAIN_STAMP_SETTING =
      "notify_update_domain_stampsetting"),
    (i.NOTIFY_UPDATE_STAMPSET = "notify_update_stampset"),
    (i.NOTIFY_DELETE_STAMPSET = "notify_delete_stampset"),
    (i.NOTIFY_UPDATE_DOMAIN_BUILTIN_STAMP_SETTING =
      "notify_update_domain_builtin_stampsetting"),
    (i.NOTIFY_FLOW_NOTIFICATION_BADGE = "notify_flow_notification_badge"),
    (i.GET_FLOW_NOTIFICATION_BADGES_COMPLETED =
      "get_flow_notification_badges_completed"),
    (i.GET_FLOW_NOTIFICATION_BADGES_FAILED =
      "get_flow_notification_badges_failed"),
    (i.GET_DIRECT_APPS_COMPLETED = "get_direct_apps_completed"),
    (i.GET_DIRECT_APPS_FAILED = "get_direct_apps_failed"),
    (i.START_NOTIFICATION_FAILED = "start_notification_failed"),
    (Fn.DEFAULT_RESOLVER = new On()),
    (Fn.BASE64 =
      "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789%:"),
    (_.accessMode = xn.Disabled),
    (te.ANNOUNCEMENTS_KEY = "-1"),
    (Gn.__meta__ = { fields: { api: { inject: null } } }),
    (Gn.NAME = "AccountControlRequest"),
    (zn.__meta__ = {
      fields: { api: { inject: null }, dataStore: { inject: null } },
    }),
    (zn.NAME = "Conference"),
    (Wn.__meta__ = { fields: { api: { inject: null } } }),
    (Wn.NAME = "Device"),
    (Vn.__meta__ = { fields: { api: { inject: null } } }),
    (Vn.NAME = "Domain"),
    (Qn.__meta__ = {
      fields: {
        api: { inject: null },
        dataStore: { inject: null },
        fileService: { inject: null },
      },
    }),
    (Qn.NAME = "FilePath"),
    (Jn.__meta__ = { fields: { api: { inject: null } } }),
    (Jn.NAME = "FilePreview"),
    (Zn.__meta__ = {
      fields: { api: { inject: null }, dataStore: { inject: null } },
    }),
    (Zn.NAME = "LoadStampset"),
    (ti.__meta__ = { fields: { api: { inject: null } } }),
    (ti.NAME = "ManageFriends"),
    (ni.__meta__ = { fields: { api: { inject: null } } }),
    (ni.NAME = "Message"),
    (ri.__meta__ = {
      fields: { dataStore: { inject: null }, api: { inject: null } },
    }),
    (ri.NAME = "MessageReaction"),
    (ai.__meta__ = {
      fields: {
        api: { inject: null },
        dataStore: { inject: null },
        fileService: { inject: null },
      },
    }),
    (ai.NAME = "Note"),
    (li.__meta__ = {
      fields: {
        api: { inject: null },
        keywordWatcher: { inject: null },
        readStatusUpdater: { inject: null },
      },
    }),
    (li.NAME = "Read"),
    (ci.__meta__ = { fields: { api: { inject: null } } }),
    (ci.NAME = "ReloadData"),
    (hi.__meta__ = {
      fields: { dataStore: { inject: null }, api: { inject: null } },
    }),
    (hi.NAME = "ScheduleMessage"),
    (fi.__meta__ = {
      fields: { api: { inject: null }, searchService: { inject: null } },
    }),
    (fi.NAME = "Search"),
    (pi.__meta__ = {
      fields: { dataStore: { inject: null }, settings: { inject: null } },
    }),
    (pi.NAME = "SelectTalk"),
    (yi.__meta__ = { fields: { api: { inject: null } } }),
    (yi.NAME = "Send"),
    (Si.__meta__ = {
      fields: {
        accessTokenResolver: { inject: null },
        session: { inject: null },
      },
    }),
    (Si.NAME = "SignIn"),
    (vi.__meta__ = {
      fields: {
        api: { inject: null },
        settings: { inject: null },
        dataStore: { inject: null },
      },
    }),
    (vi.NAME = "SignOut"),
    (Ti.__meta__ = {
      fields: { api: { inject: null }, dataStore: { inject: null } },
    }),
    (Ti.NAME = "Talk"),
    (Ei.__meta__ = {
      fields: { api: { inject: null }, dataStore: { inject: null } },
    }),
    (Ei.NAME = "UpdateProfile"),
    (wi.__meta__ = { fields: { api: { inject: null } } }),
    (wi.NAME = "UpdateUser"),
    (Ni.__meta__ = {
      fields: {
        userPresences: { inject: null },
        api: { inject: null },
        settings: { inject: null },
      },
    }),
    (Ni.NAME = "UpdateUserPresences"),
    (Ni.MAX_USERS_PER_API = 100),
    (Ai.__meta__ = { fields: { routing: { inject: null } } }),
    (Ai.NAME = "Url"),
    (_e.TalkAutoScrollDelay = 1),
    (_e.TalkAutoScrollImmediately = 2),
    (_e.TalkAutoScrollNone = 3),
    (_e.NotificationPopupTime5Sec = 5),
    (_e.TalksSortModeNewMessagesFirst = "new_messages_first"),
    (_e.TalksSortModeUnreadMessagesFirst = "unread_messages_first"),
    (_e.TalksSortModeOldMessagesFirst = "old_messages_first"),
    (_e.DomainsSortModeNameOrderAsc = "name_order_asc"),
    (_e.OLD_KEY_RIGHT_PANEL_OPENED = "right_panel_opened"),
    (_e.OLD_KEY_RIGHT_PANEL_WIDTH = "right_panel_width"),
    (_e.OLD_KEY_APP_LEFT_PANEL_OPENED = "app_left_panel_opened"),
    (_e.NAME = "browser_settings"),
    (Ui.MAX_SIZE = 72),
    (he.INCOMING_TIME = 6e4),
    (Bi.DEFAULT_PRESENCE_EXPIRATION = 3e5),
    (Bi.DEFAULT_UPDATE_LAST_USED_AT_INTERVAL = 6e5),
    (Gi.DEFAULT_TTL = 3600),
    (Gi.DEFAULT_MAX_PARTICIPANTS = 8),
    (ar.stagedFileInfosMaxId = 0),
    (Me.TIMESTAMP_SHIFT = 22),
    (O.MAX_READ_USER_IDS_COUNT = 16),
    (He.QUERY_DATE_TODAY = "t"),
    (He.QUERY_DATE_YESTERDAY = "-1d"),
    (He.QUERY_DATE_THREE_DAYS_BEFORE = "-3d"),
    (Wr.HISTORY_TAB_NAME = "stamp-history"),
    (ot.SIXTY_MINUTES = new E(0, 36e5)),
    (ot.FIFTEEN_MINUTES = new E(0, 9e5)),
    (So.PREFIX = "note-"),
    (So.nextDummyId = new E(-1, -1)),
    (xt.CONTINUE_MARGIN_RATE = 0.1),
    (Ht.CHEKBOX_UNCHECKED = "- [ ] "),
    (Ht.CHEKBOX_CHECKED = "- [x] "),
    (Yt.TU_REG = new v("ッ([BCDFGHIJKLMNOPQRSTUVWYZ])", "gm")),
    (Yt.XTU_REG = new v("ッ", "gm")),
    (F.EMERGENCY = 0),
    (F.ALERT = 1),
    (F.CRITICAL = 2),
    (F.ERROR = 3),
    (F.WARNING = 4),
    (F.NOTICE = 5),
    (F.INFO = 6),
    (F.DEBUG = 7),
    (F.level = F.getLogLevel(process.env.HUBOT_LOG_LEVEL)),
    (F._d =
      7 <= F.level && null != console ? console.log.bind(console) : F._nop),
    (F._i =
      6 <= F.level && null != console ? console.info.bind(console) : F._nop),
    (F._w =
      4 <= F.level && null != console ? console.warn.bind(console) : F._nop),
    (F._e =
      3 <= F.level && null != console ? console.error.bind(console) : F._nop),
    (xo.HIRAGANA_SMALL_A = 12353),
    (xo.HIRAGANA_NN = 12435),
    (xo.KATAKANA_SMALL_A = 12449),
    (Po.NAME = "accessTokenResolver"),
    (l.NAME = "Proxy"),
    (Bo.__meta__ = {
      fields: {
        settings: { inject: null },
        accountLoader: { inject: null },
        api: { inject: null },
      },
    }),
    (Ho.NAME = "accountLoader"),
    (jo.__meta__ = {
      fields: {
        api: { inject: null },
        dataStore: { inject: null },
        settings: { inject: null },
        keywordWatcher: { inject: null },
      },
    }),
    (jo.NAME = "broadcast"),
    (zo.__meta__ = {
      fields: {
        wsApi: { inject: null },
        apiCaller: { inject: null },
        apiNote: { inject: null },
        settings: { inject: null },
        dataStore: { inject: null },
        dataFactory: { inject: null },
        fileService: { inject: null },
        searchService: { inject: null },
        lastUsedAtUpdater: { inject: null },
      },
    }),
    (zo.NAME = "api"),
    (zo.API_VERSION = "1.128"),
    (zo.UPDATE_VERSION = "1.128_0"),
    (zo.ACTIONS_COUNT_PER_API = 20),
    (Yo.__meta__ = { fields: { rpc: { inject: null } } }),
    (Yo.NAME = "apiCaller"),
    (Yo.DELAY_FOR_DB_REPLICATION = 500),
    (Yo.DELAY_FOR_DB_REPLICATION_FOR_RETRY = 1e3),
    (Yo.API_TRY_COUNT = 3),
    (Wo.__meta__ = {
      fields: {
        dataStore: { inject: null },
        settings: { inject: null },
        wsApi: { inject: null },
      },
    }),
    (Wo.NAME = "apiNote"),
    (Ko.NAME = "appState"),
    (Vo.__meta__ = { fields: { dataStore: { inject: null } } }),
    (Vo.NAME = "dataFactory"),
    (qo.NAME = "dataStore"),
    (qo.STORAGE_KEY_FOR_GENERAL = "dataStore"),
    (qo.STORAGE_KEY_FOR_TALKS = "dataStoreTalks"),
    (qo.STORAGE_KEY_FOR_TALK_STATUSES = "dataStoreTalkStatuses"),
    (qo.STORAGE_KEY_FOR_USERS = "dataStoreUsers"),
    (qo.TYPE_FRIEND = 0),
    (qo.TYPE_ACQUAINSTANCE = 1),
    (qo.TYPE_NONE = 2),
    (Qo.__meta__ = {
      fields: { settings: { inject: null }, dataStore: { inject: null } },
    }),
    (Qo.NAME = "features"),
    (Jo.NAME = "fileService"),
    (Zo.__meta__ = { fields: { settings: { inject: null } } }),
    (ea.__meta__ = {
      fields: {
        settings: { inject: null },
        features: { inject: null },
        dataStore: { inject: null },
      },
    }),
    (ea.NAME = "keywordWatcher"),
    (ea.TALK_KEY_PREFIX = "talk_"),
    (ea.ANNOUNCEMENT_KEY_PREFIX = "announcement_"),
    (ta.TEMP_HALF_SPACE = "&SPACE"),
    (ta.TEMP_FULL_SPACE = "&FULL_PITCH_SPACE"),
    (oa.__meta__ = { fields: { settings: { inject: null } } }),
    (oa.NAME = "lastUsedAtUpdater"),
    (aa.__meta__ = {
      fields: { settings: { inject: null }, dataStore: { inject: null } },
    }),
    (aa.NAME = "limitations"),
    (sa.__meta__ = {
      fields: {
        broadcast: { inject: null },
        lastUsedAtUpdater: { inject: null },
      },
    }),
    (sa.NAME = "rpc"),
    (sa.lastMsgId = 0),
    (la.PING_INTERVAL_ON_CONNECTED = 45e3),
    (da.__meta__ = {
      fields: { dataStore: { inject: null }, api: { inject: null } },
    }),
    (da.NAME = "readStatusUpdater"),
    (fa.__meta__ = {
      fields: { settings: { inject: null }, dataStore: { inject: null } },
    }),
    (fa.NAME = "routing"),
    (ga.__meta__ = { fields: { settings: { inject: null } } }),
    (ga.NAME = "searchService"),
    (ga.LOAD_SIZE = 20),
    (ga.DUMMY_VALUE_FOR_SET = 1),
    (ya.__meta__ = {
      fields: {
        settings: { inject: null },
        dataStore: { inject: null },
        api: { inject: null },
      },
    }),
    (ya.NAME = "session"),
    (Sa.NAME = "settings"),
    (Sa.KEY_ACCESS_TOKEN = "access_token"),
    (Sa.KEY_SELECTED_DOMAIN_ID_H = "selected_domain_id_h"),
    (Sa.KEY_SELECTED_DOMAIN_ID_L = "selected_domain_id_l"),
    (Sa.KEY_SELECTED_STAMP_TAB_ID = "selected_stamp_tab_id"),
    (Sa.KEY_INPUT_TEXT = "input_text"),
    (Sa.KEY_COPY_PROFILE_TO_ALL_DOMAINS = "copy_profile_to_all_domains"),
    (Sa.KEY_SELECTED_DEPARTMENT_IDS = "selected_department_ids"),
    (Sa.KEY_SEARCH_HISTORIES = "search_histories"),
    (Sa.KEY_MULTI_VIEW_MODE = "multi_view_mode"),
    (Sa.KEY_PASSWORD_WARNING_SKIP_UNTIL = "password_warning_skip_until"),
    (Sa.KEY_LAST_USED_EXPIRED_AT = "last_used_expired_at"),
    (Sa.KEY_IDFV = "idfv"),
    (va.__meta__ = { fields: { dataStore: { inject: null } } }),
    (va.NAME = "talksService"),
    (Ta.__meta__ = { fields: { dataStore: { inject: null } } }),
    (Ta.NAME = "userPresences"),
    (Ia.__meta__ = { fields: { apiCaller: { inject: null } } }),
    (Ia.NAME = "wsApi"),
    (p.MAX_CONTENT_BYTES = 1e5),
    (p.NOTE_EVENT_NAMES = ["created", "updated", "deleted"]),
    (Ma.CONTENT_TYPE_XML = "xml"),
    (Ma.CONTENT_TYPE_TEXT = "text"),
    (Da.NAME = "Mediator"),
    (Ra.__meta__ = {
      fields: {
        dataStore: { inject: null },
        api: { inject: null },
        hubotObject: { inject: null },
        messageEvent: { inject: null },
        sendQueue: { inject: null },
      },
    }),
    (Ra.NAME = "commandline"),
    (Ra.DATA_SAVE_SPAN = 5e3),
    (xa.__meta__ = { fields: { dataStore: { inject: null } } }),
    (xa.NAME = "hubotObject"),
    (Ua.__meta__ = {
      fields: { dataStore: { inject: null }, hubotObject: { inject: null } },
    }),
    (Ua.NAME = "messageEvent"),
    (La.minWaitMinute = 1),
    (La.maxWaitMinute = 1440),
    (La.SEND_EVENT_NAME = "send"),
    (Pa.__meta__ = { fields: { messageEvent: { inject: null } } }),
    (Pa.NAME = "sendQueue"),
    (Pa.MIN_SEND_SPAN_MESSAGE = 550),
    (Pa.MIN_SEND_SPAN_ANNOUNCE = 5050),
    (Pa.MAX_TEXT_LENGTH = 1024),
    (Pa.MAX_FIELD_COUNT = 9),
    (Pa.MAX_FIELD_LENGTH = 64),
    (Ga.INVALID_PARAMETER = "INVALID_PARAMETER"),
    (Ga.FORBIDDEN = "FORBIDDEN"),
    (Ga.NOT_FOUND = "NOT_FOUND"),
    (Ga.LOCKED_BY_ANOTHER_USER = "LOCKED_BY_ANOTHER_USER"),
    (Ga.CONFLICT = "CONFLICT"),
    (Ga.TOO_MANY_REQUESTS = "TOO_MANY_REQUESTS"),
    (Ga.CLIENT_UPGRADE_REQUIRED = "CLIENT_UPGRADE_REQUIRED"),
    (Ga.UNKNOWN = "UNKNOWN"),
    (w._mul =
      null != Math.imul
        ? Math.imul
        : function (e, t) {
            return (e * (65535 & t) + (((e * (t >>> 16)) << 16) | 0)) | 0;
          }),
    (Qa.USE_CACHE = !1),
    (Qa.USE_ENUM_INDEX = !1),
    (Qa.BASE64 =
      "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789%:"),
    (ns.count = 0),
    (cs.i64tmp = new E(0, 0)),
    (G.__toStr = {}.toString),
    (Ss.FLOAT_SINGLE_MIN = 1401298464324817e-60),
    (Ss.FLOAT_SINGLE_MAX = 34028234663852886e22),
    (Ss.FLOAT_DOUBLE_MIN = 5e-324),
    (Ss.FLOAT_DOUBLE_MAX = 17976931348623157e292),
    vn.main());
})(
  "undefined" != typeof exports
    ? exports
    : "undefined" != typeof window
      ? window
      : "undefined" != typeof self
        ? self
        : this,
  "undefined" != typeof window
    ? window
    : "undefined" != typeof global
      ? global
      : "undefined" != typeof self
        ? self
        : this,
);
