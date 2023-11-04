package todoist

type Store struct {
	CollaboratorStates CollaboratorStates `json:"collaborator_states"`
	Collaborators      Collaborators      `json:"collaborators"`
	DayOrders          interface{}        `json:"day_orders"`
	DayOrdersTimestamp string             `json:"day_orders_timestamp"`
	Filters            Filters            `json:"filters"`
	FullSync           bool               `json:"full_sync"`
	Items              Items              `json:"items"`
	Labels             Labels             `json:"labels"`
	LiveNotifications  []struct {
		CompletedTasks   int     `json:"completed_tasks"`
		Created          string  `json:"created_at"`
		DateReached      string  `json:"date_reached"`
		ID               string  `json:"id"`
		IsDeleted        bool    `json:"is_deleted"`
		KarmaLevel       int     `json:"karma_level"`
		NotificationKey  string  `json:"notification_key"`
		NotificationType string  `json:"notification_type"`
		PromoImg         string  `json:"promo_img"`
		SeqNo            int64   `json:"seq_no"`
		TopProcent       float32 `json:"top_procent"`
	} `json:"live_notifications"`
	LiveNotificationsLastReadID string        `json:"live_notifications_last_read_id"`
	Locations                   []interface{} `json:"locations"`
	Notes                       []struct {
		Content        string      `json:"content"`
		FileAttachment interface{} `json:"file_attachment"`
		ID             string      `json:"id"`
		IsArchived     int         `json:"is_archived"`
		IsDeleted      bool        `json:"is_deleted"`
		ItemID         string      `json:"item_id"`
		Posted         string      `json:"posted_at"`
		PostedUID      string      `json:"posted_uid"`
		ProjectID      *string     `json:"project_id"`
		UidsToNotify   interface{} `json:"uids_to_notify"`
	} `json:"notes"`
	ProjectNotes  []interface{}       `json:"project_notes"`
	Projects      Projects            `json:"projects"`
	Sections      Sections            `json:"sections"`
	Reminders     Reminders           `json:"reminders"`
	SyncToken     string              `json:"sync_token"`
	TempIDMapping struct{}            `json:"temp_id_mapping"`
	User          User                `json:"user"`
	RootItem      *Item               `json:"-"`
	RootProject   *Project            `json:"-"`
	ItemMap       map[string]*Item    `json:"-"`
	ProjectMap    map[string]*Project `json:"-"`
	LabelMap      map[string]*Label   `json:"-"`
	SectionMap    map[string]*Section `json:"-"`
}

type Filter struct {
	HaveID
	Color     string `json:"color"`
	IsDeleted bool   `json:"is_deleted"`
	ItemOrder int    `json:"item_order"`
	Name      string `json:"name"`
	Query     string `json:"query"`
}

func (f Filter) Removable() bool {
	return f.IsDeleted
}

type Filters []Filter

type Reminder struct {
	HaveID
	DateLang     string      `json:"date_lang"`
	Due          *Due        `json:"due"`
	IsDeleted    bool        `json:"is_deleted"`
	ItemID       string      `json:"item_id"`
	MinuteOffset int         `json:"minute_offset"`
	NotifyUID    string      `json:"notify_uid"`
	Service      interface{} `json:"service"`
	Type         string      `json:"type"`
}

func (r Reminder) Removable() bool {
	return r.IsDeleted
}

type Reminders []Reminder

type ID interface {
	GetID() string
}

func (s *Store) FindItem(id string) *Item {
	return s.ItemMap[id]
}

func (s *Store) FindProject(id string) *Project {
	return s.ProjectMap[id]
}

func (s *Store) FindSection(id string) *Section {
	return s.SectionMap[id]
}

func (s *Store) FindLabel(id string) *Label {
	return s.LabelMap[id]
}

type Syncable interface {
	IDCarrier
	Removable() bool
}

func replace[T Syncable](s []T, n T) []T {
	for i, item := range s {
		if item.GetID() == n.GetID() {
			s[i] = n
			return s
		}
	}
	return s
}

func remove[T Syncable](s []T, ID string) []T {
	for i, item := range s {
		if item.GetID() == ID {
			s[i] = s[len(s)-1]
			return s[:len(s)-1]
		}
	}
	return s
}

func syncCollection[T Syncable](target []T, source []T) []T {
	for _, s := range source {
		id := s.GetID()
		if s.Removable() {
			target = remove(target, id)
			break
		}
		found := false
		for _, t := range target {
			if t.GetID() == id {
				target = replace(target, s)
				found = true
				break
			}
		}
		if !found {
			target = append(target, s)
		}
	}
	return target
}

func (target *Store) ApplyIncrementalSync(source *Store) {
	if source.FullSync {
		target = source
		return
	}
	target.Collaborators = syncCollection(target.Collaborators, source.Collaborators)
	// target.CollaboratorStates = syncCollection(target.CollaboratorStates, source.CollaboratorStates)
	// target.DayOrders = syncCollection(target.DayOrders, source.DayOrders)
	target.DayOrdersTimestamp = source.DayOrdersTimestamp
	target.Filters = syncCollection(target.Filters, source.Filters)
	target.FullSync = source.FullSync
	target.Items = syncCollection(target.Items, source.Items)
	target.Labels = syncCollection(target.Labels, source.Labels)
	// target.LiveNotifications = syncCollection(target.LiveNotifications, source.LiveNotifications)
	target.LiveNotificationsLastReadID = source.LiveNotificationsLastReadID
	// target.Locations = syncCollection(target.Locations, source.Locations)
	// target.Notes = syncCollection(target.Notes, source.Notes)
	// target.ProjectNotes = syncCollection(target.ProjectNotes, source.ProjectNotes)
	target.Projects = syncCollection(target.Projects, source.Projects)
	target.Sections = syncCollection(target.Sections, source.Sections)
	target.Reminders = syncCollection(target.Reminders, source.Reminders)
	target.SyncToken = source.SyncToken
	if source.User != (User{}) {
		target.User = source.User
	}
	target.ConstructItemTree()
}

func addToBrotherItem(item *Item, b *Item) {
	i := item
	for {
		if i.BrotherItem == nil {
			i.BrotherItem = b
			return
		}
		i = i.BrotherItem
	}
}

func addToChildItem(item *Item, b *Item) {
	if item.ChildItem == nil {
		item.ChildItem = b
		return
	}
	addToBrotherItem(item.ChildItem, b)
}

func addToBrotherProject(project *Project, b *Project) {
	i := project
	for {
		if i.BrotherProject == nil {
			i.BrotherProject = b
			return
		}
		i = i.BrotherProject
	}
}

func addToChildProject(project *Project, b *Project) {
	if project.ChildProject == nil {
		project.ChildProject = b
		return
	}
	addToBrotherProject(project.ChildProject, b)
}

func (s *Store) ConstructItemTree() {
	s.LabelMap = map[string]*Label{}
	s.ProjectMap = map[string]*Project{}
	s.ItemMap = map[string]*Item{}
	s.SectionMap = map[string]*Section{}

	for i, label := range s.Labels {
		s.LabelMap[label.ID] = &s.Labels[i]
	}

	for i, item := range s.Items {
		s.ItemMap[item.ID] = &s.Items[i]
		s.Items[i].ChildItem = nil
		s.Items[i].BrotherItem = nil
	}

	for i, project := range s.Projects {
		s.ProjectMap[project.ID] = &s.Projects[i]
		s.Projects[i].ChildProject = nil
		s.Projects[i].BrotherProject = nil
	}

	for i, section := range s.Sections {
		s.SectionMap[section.ID] = &s.Sections[i]
	}

	for _, item := range s.Items {
		if item.ParentID == nil {
			s.RootItem = &item
			break
		}
	}

	for _, project := range s.Projects {
		if project.ParentID == nil {
			s.RootProject = &project
			break
		}
	}

	for i := range s.Items {
		if s.Items[i].ID == s.RootItem.ID {
			continue
		}

		if s.Items[i].ParentID == nil {
			addToBrotherItem(s.RootItem, &s.Items[i])
			continue
		}
		id, _ := s.Items[i].GetParentID()
		parent := s.FindItem(id)
		addToChildItem(parent, &s.Items[i])
	}

	for i := range s.Projects {
		if s.Projects[i].ID == s.RootProject.ID {
			continue
		}

		if s.Projects[i].ParentID == nil {
			addToBrotherProject(s.RootProject, &s.Projects[i])
			continue
		}
		id, _ := s.Projects[i].GetParentID()
		parent := s.FindProject(id)
		addToChildProject(parent, &s.Projects[i])
	}
}
