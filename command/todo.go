package command

import "time"

// Format represents how timestamps show look
var Format = "2006-01-02T15:04"
var todoFormat = "TODO(<user>) <message> <timestamp> <weight>"

// Todoer implements the functionality to view todos
type Todoer interface {
	User() string
	File() string
	Message() string
	Timestamp() time.Time
	Weight() int
}

// Todo represents a todo
type Todo struct {
	user      string
	file      string
	message   string
	timestamp time.Time
	weight    int
}

// NewTodo creates a new reference of a todo
func NewTodo(u, m, ts, f string, w int) (*Todo, error) {
	todo := &Todo{
		user:    u,
		file:    f,
		message: m,
		weight:  w,
	}
	timestamp, err := time.Parse(Format, ts)
	if err != nil {
		return nil, err
	}
	todo.timestamp = timestamp
	return todo, err
}

func (t *Todo) User() string {
	return t.user
}

func (t *Todo) File() string {
	return t.file
}

func (t *Todo) Message() string {
	return t.message
}

func (t *Todo) Timestamp() time.Time {
	return t.timestamp
}

func (t *Todo) Weight() int {
	return t.weight
}

// UserTodos is a slice type made for easier sorting
type UserTodos []Todo

// Len gets the length of the slice
func (u UserTodos) Len() int {
	return len(u)
}

// Less does a comparison of the 2 given arguments
func (u UserTodos) Less(i, j int) bool {
	return u[i].user < u[j].user
}

// Swap switchs the place in the slice for the 2 given arguments
func (u UserTodos) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

// Reverse will reverse the order of elements in the slice
func (u UserTodos) Reverse() {

}

// FileTodos is a slice type made for easier sorting
type FileTodos []Todo

// Len gets the length of the slice
func (f FileTodos) Len() int {
	return len(f)
}

// Less does a comparison of the 2 given arguments
func (f FileTodos) Less(i, j int) bool {
	return f[i].file < f[j].file
}

// Swap switchs the place in the slice for the 2 given arguments
func (f FileTodos) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

// TimestampTodos is a slice type made for easier sorting
type TimestampTodos []Todo

// Len gets the length of the slice
func (t TimestampTodos) Len() int {
	return len(t)
}

// Less does a comparison of the 2 given arguments
func (t TimestampTodos) Less(i, j int) bool {
	return t[i].timestamp.Before(t[j].timestamp)
}

// Swap switchs the place in the slice for the 2 given arguments
func (t TimestampTodos) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// WeightTodos is a slice type made for easier sorting
type WeightTodos []Todo

// Len gets the length of the slice
func (w WeightTodos) Len() int {
	return len(w)
}

// Less does a comparison of the 2 given arguments
func (w WeightTodos) Less(i, j int) bool {
	return w[i].weight < w[j].weight
}

// Swap switchs the place in the slice for the 2 given arguments
func (w WeightTodos) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}
