package task_queue

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

type Queue struct {
	q []map[string]interface{}
}

type TaskConfig struct {
	Id int
	// 任务队列大小
	Size int
	// 任务队列名称
	Name string
	// 消费者个数
	Consumer int
}

func GetAttribute(name string, config *TaskConfig) (interface{}, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	if config == nil {
		return nil, errors.New("config is nil")
	}
	v, err := getStructField(config, name)
	if err != nil {
		return nil, nil
	}

	return v, nil
}

func getStructField(input interface{}, key string) (value interface{}, err error) {
	rv := reflect.ValueOf(input)
	rt := reflect.TypeOf(input)
	if rt.Kind() != reflect.Struct {
		return value, errors.New("input must be struct")
	}
	for i := 0; i < rt.NumField(); i++ {
		curField := rv.Field(i)
		if rt.Field(i).Name == key {
			switch curField.Kind() {
			case reflect.String, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int, reflect.Float64, reflect.Float32:
				return curField.Interface(), nil
			default:
				return value, errors.New("key must be int float or string")
			}
		}
	}
	return nil, errors.New("key not found")
}

type Task struct {
	config *TaskConfig
	// 队列
	Q *Queue
	m sync.RWMutex
}

func NewTask() *Task {
	return &Task{}
}

var taskList = &TaskList{map[int]*Task{}}

type TaskList struct {
	l map[int]*Task
}

type TaskI interface {
	NewTask(config *TaskConfig) (*Task, error)
	DelTask(id int) error
	Push(taskId int, data map[string]interface{}) error
	Consume(taskId int, handle func()) error
}

// 确保实现
var _ TaskI = (*Task)(nil)

func (t *Task) NewTask(config *TaskConfig) (*Task, error) {
	// create task if not
	id := t.getId()
	taskConfig := &TaskConfig{
		Id:       id,
		Size:     10,
		Name:     config.Name,
		Consumer: config.Consumer,
	}
	taskList.l[id] = &Task{config: taskConfig}
	return nil, nil
}

func (t *Task) getId() int {
	return time.Now().Second()
}

func (t *Task) DelTask(id int) error {
	delete(taskList.l, id)
	return nil
}

func (t *Task) Push(taskId int, data map[string]interface{}) error {
	// task exist
	t.m.Lock()
	taskList.l[taskId].Q.q = append(taskList.l[taskId].Q.q, data)
	t.m.Unlock()
	return nil
}

func (t *Task) Consume(taskId int, handle func()) error {
	handle()
	return nil
}
