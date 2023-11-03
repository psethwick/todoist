package todoist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddItem(t *testing.T) {
	store := Store{}
	store.ConstructItemTree()
	is := Store{
		Items: Items{Item{BaseItem: BaseItem{HaveID: HaveID{ID: "1"}}}},
	}
	is.ConstructItemTree()
	store.ApplyIncrementalSync(&is)
	assert.True(t, store.FindItem("1") != nil)
}

func TestDeleteItem(t *testing.T) {
	s := Store{
		Items: Items{Item{BaseItem: BaseItem{HaveID: HaveID{ID: "1"}}}},
	}
	is := Store{
		Items: Items{Item{IsDeleted: true, BaseItem: BaseItem{HaveID: HaveID{ID: "1"}}}},
	}
	s.ApplyIncrementalSync(&is)
	assert.Empty(t, s.Items)
	assert.Nil(t, s.FindItem("1"))
}

func TestUpdateItem(t *testing.T) {
	s := Store{
		Items: Items{Item{BaseItem: BaseItem{Content: "old", HaveID: HaveID{ID: "1"}}}},
	}
	is := Store{
		Items: Items{Item{BaseItem: BaseItem{Content: "new", HaveID: HaveID{ID: "1"}}}},
	}
	s.ApplyIncrementalSync(&is)
	assert.Equal(t, s.FindItem("1").GetContent(), "new")
}

func TestCompleteItem(t *testing.T) {
	s := Store{
		Items: Items{Item{BaseItem: BaseItem{HaveID: HaveID{ID: "1"}}}},
	}
	is := Store{
		Items: Items{Item{CompletedAt: "2023-10-31T20:00:00", BaseItem: BaseItem{HaveID: HaveID{ID: "1"}}}},
	}
	s.ApplyIncrementalSync(&is)
	assert.Empty(t, s.Items)
	assert.Nil(t, s.FindItem("1"))
}
