package main

type User struct {
	AutoReminder      int         `json:"auto_reminder"`
	AvatarBig         string      `json:"avatar_big"`
	AvatarMedium      string      `json:"avatar_medium"`
	AvatarS640        string      `json:"avatar_s640"`
	AvatarSmall       string      `json:"avatar_small"`
	BusinessAccountID interface{} `json:"business_account_id"`
	CompletedCount    int         `json:"completed_count"`
	CompletedToday    int         `json:"completed_today"`
	DailyGoal         int         `json:"daily_goal"`
	DateFormat        int         `json:"date_format"`
	DefaultReminder   string      `json:"default_reminder"`
	Email             string      `json:"email"`
	Features          struct {
		Beta             int  `json:"beta"`
		GoldTheme        bool `json:"gold_theme"`
		HasPushReminders bool `json:"has_push_reminders"`
		Restriction      int  `json:"restriction"`
	} `json:"features"`
	FullName     string      `json:"full_name"`
	ID           int         `json:"id"`
	ImageID      string      `json:"image_id"`
	InboxProject int         `json:"inbox_project"`
	IsBizAdmin   bool        `json:"is_biz_admin"`
	IsPremium    bool        `json:"is_premium"`
	JoinDate     string      `json:"join_date"`
	Karma        float32     `json:"karma"`
	KarmaTrend   string      `json:"karma_trend"`
	MobileHost   interface{} `json:"mobile_host"`
	MobileNumber interface{} `json:"mobile_number"`
	NextWeek     int         `json:"next_week"`
	PremiumUntil string      `json:"premium_until"`
	SortOrder    int         `json:"sort_order"`
	StartDay     int         `json:"start_day"`
	StartPage    string      `json:"start_page"`
	Theme        int         `json:"theme"`
	TimeFormat   int         `json:"time_format"`
	Token        string      `json:"token"`
	TzInfo       struct {
		GmtString string `json:"gmt_string"`
		Hours     int    `json:"hours"`
		IsDst     int    `json:"is_dst"`
		Minutes   int    `json:"minutes"`
		Timezone  string `json:"timezone"`
	} `json:"tz_info"`
}
