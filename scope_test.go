package sentry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetUser(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetUser(User{id: "foo"})
	assert.Equal(User{id: "foo"}, scope.user)
}

func TestSetUserOverrides(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetUser(User{id: "foo"})
	scope.SetUser(User{id: "bar"})
	assert.Equal(User{id: "bar"}, scope.user)
}

func TestSetTag(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetTag("a", "foo")
	assert.Equal(map[string]string{"a": "foo"}, scope.tags)
}

func TestSetTagMerges(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetTag("a", "foo")
	scope.SetTag("b", "bar")
	assert.Equal(map[string]string{"a": "foo", "b": "bar"}, scope.tags)
}

func TestSetTagOverrides(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetTag("a", "foo")
	scope.SetTag("a", "bar")
	assert.Equal(map[string]string{"a": "bar"}, scope.tags)
}

func TestSetTags(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetTags(map[string]string{"a": "foo"})
	assert.Equal(map[string]string{"a": "foo"}, scope.tags)
}
func TestSetTagsMerges(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetTags(map[string]string{"a": "foo"})
	scope.SetTags(map[string]string{"b": "bar", "c": "baz"})
	assert.Equal(map[string]string{"a": "foo", "b": "bar", "c": "baz"}, scope.tags)
}

func TestSetTagsOverrides(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetTags(map[string]string{"a": "foo"})
	scope.SetTags(map[string]string{"a": "bar", "b": "baz"})
	assert.Equal(map[string]string{"a": "bar", "b": "baz"}, scope.tags)
}

func TestSetExtra(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetExtra("a", 1)
	assert.Equal(map[string]interface{}{"a": 1}, scope.extra)
}
func TestSetExtraMerges(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetExtra("a", "foo")
	scope.SetExtra("b", 2)
	assert.Equal(map[string]interface{}{"a": "foo", "b": 2}, scope.extra)
}

func TestSetExtraOverrides(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetExtra("a", "foo")
	scope.SetExtra("a", 2)
	assert.Equal(map[string]interface{}{"a": 2}, scope.extra)
}

func TestSetExtras(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetExtras(map[string]interface{}{"a": 1})
	assert.Equal(map[string]interface{}{"a": 1}, scope.extra)
}
func TestSetExtrasMerges(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetExtras(map[string]interface{}{"a": "foo"})
	scope.SetExtras(map[string]interface{}{"b": 2, "c": 3})
	assert.Equal(map[string]interface{}{"a": "foo", "b": 2, "c": 3}, scope.extra)
}

func TestSetExtrasOverrides(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetExtras(map[string]interface{}{"a": "foo"})
	scope.SetExtras(map[string]interface{}{"a": 2, "b": 3})
	assert.Equal(map[string]interface{}{"a": 2, "b": 3}, scope.extra)
}
func TestSetFingerprint(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetFingerprint([]string{"abcd"})
	assert.Equal([]string{"abcd"}, scope.fingerprint)
}

func TestSetFingerprintOverrides(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetFingerprint([]string{"abc"})
	scope.SetFingerprint([]string{"def"})
	assert.Equal([]string{"def"}, scope.fingerprint)
}

func TestSetLevel(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetLevel(LevelInfo)
	assert.Equal(LevelInfo, scope.level)
}

func TestSetLevelOverrides(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.SetLevel(LevelInfo)
	scope.SetLevel(LevelFatal)
	assert.Equal(LevelFatal, scope.level)
}

func TestAddBreadcrumb(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.AddBreadcrumb(Breadcrumb{message: "test"})
	assert.Equal([]Breadcrumb{{message: "test"}}, scope.breadcrumbs)
}

func TestAddBreadcrumbAppends(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.AddBreadcrumb(Breadcrumb{message: "test1"})
	scope.AddBreadcrumb(Breadcrumb{message: "test2"})
	scope.AddBreadcrumb(Breadcrumb{message: "test3"})
	assert.Equal([]Breadcrumb{{message: "test1"}, {message: "test2"}, {message: "test3"}}, scope.breadcrumbs)
}

func TestAddBreadcrumbDefaultLimit(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	for i := 0; i < 101; i++ {
		scope.AddBreadcrumb(Breadcrumb{message: "test"})
	}
	assert.Len(scope.breadcrumbs, 100)
}

func TestBasicInheritance(t *testing.T) {
	assert := assert.New(t)
	parentScope := NewScope()
	parentScope.SetExtra("a", 1)
	scope := parentScope.Clone()
	assert.Equal(parentScope.extra, scope.extra)
}

func TestParentChangedInheritance(t *testing.T) {
	assert := assert.New(t)
	parentScope := NewScope()
	scope := parentScope.Clone()

	scope.SetTag("foo", "bar")
	scope.SetExtra("foo", "bar")
	scope.SetLevel(LevelDebug)
	scope.SetFingerprint([]string{"foo"})
	scope.AddBreadcrumb(Breadcrumb{message: "foo"})
	scope.SetUser(User{id: "foo"})

	parentScope.SetTag("foo", "baz")
	parentScope.SetExtra("foo", "baz")
	parentScope.SetLevel(LevelFatal)
	parentScope.SetFingerprint([]string{"bar"})
	parentScope.AddBreadcrumb(Breadcrumb{message: "bar"})
	parentScope.SetUser(User{id: "bar"})

	assert.Equal(map[string]string{"foo": "bar"}, scope.tags)
	assert.Equal(map[string]interface{}{"foo": "bar"}, scope.extra)
	assert.Equal(LevelDebug, scope.level)
	assert.Equal([]string{"foo"}, scope.fingerprint)
	assert.Equal([]Breadcrumb{{message: "foo"}}, scope.breadcrumbs)
	assert.Equal(User{id: "foo"}, scope.user)

	assert.Equal(map[string]string{"foo": "baz"}, parentScope.tags)
	assert.Equal(map[string]interface{}{"foo": "baz"}, parentScope.extra)
	assert.Equal(LevelFatal, parentScope.level)
	assert.Equal([]string{"bar"}, parentScope.fingerprint)
	assert.Equal([]Breadcrumb{{message: "bar"}}, parentScope.breadcrumbs)
	assert.Equal(User{id: "bar"}, parentScope.user)
}

func TestChildOverrideInheritance(t *testing.T) {
	assert := assert.New(t)
	parentScope := NewScope()
	parentScope.SetTag("foo", "baz")
	parentScope.SetExtra("foo", "baz")
	parentScope.SetLevel(LevelFatal)
	parentScope.SetFingerprint([]string{"bar"})
	parentScope.AddBreadcrumb(Breadcrumb{message: "bar"})
	parentScope.SetUser(User{id: "bar"})

	scope := parentScope.Clone()
	scope.SetTag("foo", "bar")
	scope.SetExtra("foo", "bar")
	scope.SetLevel(LevelDebug)
	scope.SetFingerprint([]string{"foo"})
	scope.AddBreadcrumb(Breadcrumb{message: "foo"})
	scope.SetUser(User{id: "foo"})

	assert.Equal(map[string]string{"foo": "bar"}, scope.tags)
	assert.Equal(map[string]interface{}{"foo": "bar"}, scope.extra)
	assert.Equal(LevelDebug, scope.level)
	assert.Equal([]string{"foo"}, scope.fingerprint)
	assert.Equal([]Breadcrumb{{message: "bar"}, {message: "foo"}}, scope.breadcrumbs)
	assert.Equal(User{id: "foo"}, scope.user)

	assert.Equal(map[string]string{"foo": "baz"}, parentScope.tags)
	assert.Equal(map[string]interface{}{"foo": "baz"}, parentScope.extra)
	assert.Equal(LevelFatal, parentScope.level)
	assert.Equal([]string{"bar"}, parentScope.fingerprint)
	assert.Equal([]Breadcrumb{{message: "bar"}}, parentScope.breadcrumbs)
	assert.Equal(User{id: "bar"}, parentScope.user)
}

func TestClear(t *testing.T) {
	assert := assert.New(t)
	scope := NewScope()
	scope.AddBreadcrumb(Breadcrumb{message: "test"})
	scope.SetUser(User{id: "1"})
	scope.SetTag("a", "b")
	scope.SetExtra("a", 2)
	scope.SetFingerprint([]string{"abcd"})
	scope.SetLevel(LevelFatal)
	scope.Clear()
	assert.Equal([]Breadcrumb{}, scope.breadcrumbs)
	assert.Equal(User{}, scope.user)
	assert.Equal(map[string]string{}, scope.tags)
	assert.Equal(map[string]interface{}{}, scope.extra)
	assert.Equal([]string{}, scope.fingerprint)
	assert.Equal(LevelInfo, scope.level)
}

func TestApplyToEvent(t *testing.T) {
	assert := assert.New(t)
	scope := &Scope{
		breadcrumbs: []Breadcrumb{{message: "scopeFoo"}},
		user:        User{id: "1337"},
		tags:        map[string]string{"scopeFoo": "scopeBar"},
		extra:       map[string]interface{}{"scopeFoo": "scopeBar"},
		fingerprint: []string{"scopeFoo", "scopeBar"},
		level:       LevelDebug,
	}
	event := &Event{
		breadcrumbs: []Breadcrumb{{message: "eventFoo"}},
		user:        User{id: "42"},
		tags:        map[string]string{"eventFoo": "eventBar"},
		extra:       map[string]interface{}{"eventFoo": "eventBar"},
		fingerprint: []string{"eventFoo", "eventBar"},
		level:       LevelInfo,
	}

	processedEvent := scope.ApplyToEvent(event)

	assert.Len(processedEvent.breadcrumbs, 2, "should merge breadcrumbs")
	assert.Len(processedEvent.tags, 2, "should merge tags")
	assert.Len(processedEvent.extra, 2, "should merge extra")
	assert.Equal(processedEvent.level, scope.level, "should use scope level if its set")
	assert.NotEqual(processedEvent.user, scope.user, "should use event user if one exist")
	assert.NotEqual(processedEvent.fingerprint, scope.fingerprint, "should use event fingerprints if they exist")
}

func TestApplyToEventEmptyScope(t *testing.T) {
	assert := assert.New(t)
	scope := &Scope{}
	event := &Event{
		breadcrumbs: []Breadcrumb{{message: "eventFoo"}},
		user:        User{id: "42"},
		tags:        map[string]string{"eventFoo": "eventBar"},
		extra:       map[string]interface{}{"eventFoo": "eventBar"},
		fingerprint: []string{"eventFoo", "eventBar"},
		level:       LevelInfo,
	}

	processedEvent := scope.ApplyToEvent(event)

	assert.True(true, "Shoudn't blow up")
	assert.Len(processedEvent.breadcrumbs, 1, "should use event breadcrumbs")
	assert.Len(processedEvent.tags, 1, "should use event tags")
	assert.Len(processedEvent.extra, 1, "should use event extra")
	assert.NotEqual(processedEvent.user, scope.user, "should use event user")
	assert.NotEqual(processedEvent.fingerprint, scope.fingerprint, "should use event fingerprint")
	assert.NotEqual(processedEvent.level, scope.level, "should use event level")
}

func TestApplyToEventEmptyEvent(t *testing.T) {
	assert := assert.New(t)
	scope := &Scope{
		breadcrumbs: []Breadcrumb{{message: "scopeFoo"}},
		user:        User{id: "1337"},
		tags:        map[string]string{"scopeFoo": "scopeBar"},
		extra:       map[string]interface{}{"scopeFoo": "scopeBar"},
		fingerprint: []string{"scopeFoo", "scopeBar"},
		level:       LevelDebug,
	}
	event := &Event{}

	processedEvent := scope.ApplyToEvent(event)

	assert.True(true, "Shoudn't blow up")
	assert.Len(processedEvent.breadcrumbs, 1, "should use scope breadcrumbs")
	assert.Len(processedEvent.tags, 1, "should use scope tags")
	assert.Len(processedEvent.extra, 1, "should use scope extra")
	assert.Equal(processedEvent.user, scope.user, "should use scope user")
	assert.Equal(processedEvent.fingerprint, scope.fingerprint, "should use scope fingerprint")
	assert.Equal(processedEvent.level, scope.level, "should use scope level")
}

func TestEventProcessors(t *testing.T) {
	assert := assert.New(t)
	scope := &Scope{
		eventProcessors: []EventProcessor{
			func(event *Event) *Event {
				event.level = LevelFatal
				return event
			},
			func(event *Event) *Event {
				event.fingerprint = []string{"wat"}
				return event
			},
		},
	}

	processedEvent := scope.ApplyToEvent(&Event{})

	assert.NotNil(processedEvent)
	assert.Equal(LevelFatal, processedEvent.level)
	assert.Equal([]string{"wat"}, processedEvent.fingerprint)
}

func TestEventProcessorsCanDropEvent(t *testing.T) {
	assert := assert.New(t)
	scope := &Scope{
		tags: map[string]string{"scopeFoo": "scopeBar"},
		eventProcessors: []EventProcessor{
			func(event *Event) *Event {
				return nil
			},
		},
	}

	processedEvent := scope.ApplyToEvent(&Event{})

	assert.Nil(processedEvent)
}

func TestAddEventProcessor(t *testing.T) {
	assert := assert.New(t)
	scope := &Scope{}

	processedEvent := scope.ApplyToEvent(&Event{})
	assert.NotNil(processedEvent)

	scope.AddEventProcessor(func(event *Event) *Event {
		return nil
	})

	processedEvent = scope.ApplyToEvent(&Event{})
	assert.Nil(processedEvent)
}
