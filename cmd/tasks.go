package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

// Task has task data
type Task struct {
	Name    string   `json:"name,omitempty"`
	TaskID  int      `json:"id,omitempty"`
	Descr   string   `json:"description,omitempty"`
	State   string   `json:"state,omitempty"`
	Tags    []string `json:"tags,omitempty"`
	Version string   `json:"version,omitempty"`
}

var todoAppAddress string

func getPath() string {
	return fmt.Sprintf("http://%s/todoapp", todoAppAddress)
}

func init() {
	addr, addrFound := os.LookupEnv("TODOCLI_ADDR")
	port, portFound := os.LookupEnv("TODOCLI_PORT")
	if !addrFound {
		addr = "localhost"
	}
	if !portFound {
		port = "8003"
	}
	todoAppAddress = fmt.Sprintf("%s:%s", addr, port)
}

func handleResponse(resp *http.Response, expectedStatusCode int) error {
	if resp.StatusCode != expectedStatusCode {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println(fmt.Sprintf("response: %s", string(body)))
		}
		return fmt.Errorf("unexpected error code: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}

func updTask(task Task, taskID string) error {
	tasks, err := getTaskWithID(taskID)
	if err != nil {
		return err
	}
	if l := len(tasks); l != 1 {
		return fmt.Errorf("Should find one task for id (%s)(%d)", taskID, l)
	}
	task.Version = tasks[0].Version

	url := fmt.Sprintf("%s/v1/tasks/%s", getPath(), taskID)
	content, err := json.Marshal(&task)
	if err != nil {
		return err
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(content))
	if err != nil {
		return err
	}
	return handleResponse(resp, 200)
}

func delTask(taskID string) error {
	client := &http.Client{}
	u, err := url.Parse(fmt.Sprintf("%s/v1/tasks/%s", getPath(), taskID))
	if err != nil {
		return err
	}
	req := &http.Request{
		Method: "DELETE",
		URL:    u,
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return handleResponse(resp, 200)
}

func importTasks(content []byte) error {
	url := fmt.Sprintf("%s/v1/import/tasks", getPath())
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(content))
	if err != nil {
		return err
	}
	return handleResponse(resp, 201)
}

func addTask(task Task) error {
	content, err := json.Marshal(&task)
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/v1/tasks", getPath())
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(content))
	if err != nil {
		return err
	}
	return handleResponse(resp, 201)
}

func getTaskWithID(taskID string) ([]Task, error) {
	tasks := []Task{}

	resp, err := http.Get(fmt.Sprintf("%s/v1/tasks/%s", getPath(), taskID))
	if err != nil {
		return []Task{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println(fmt.Sprintf("response: %s", string(body)))
		}
		return []Task{}, fmt.Errorf("unexpected error code: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Task{}, err
	}
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		return []Task{}, err
	}
	return tasks, nil
}

func getAllTasks(query string) ([]Task, error) {
	tasks := []Task{}

	var queryPart string
	if query != "" {
		queryPart = "?" + url.PathEscape(query)
	}
	//fmt.Println(fmt.Sprintf("Q=<%s>", queryPart))
	resp, err := http.Get(fmt.Sprintf("%s/v1/tasks%s", getPath(), queryPart))
	if err != nil {
		return []Task{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println(fmt.Sprintf("response: %s", string(body)))
		}
		return []Task{}, fmt.Errorf("unexpected error code: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Task{}, err
	}
	err = json.Unmarshal(body, &tasks)
	if err != nil {
		return []Task{}, err
	}
	return tasks, nil
}

func exportAllTasks(query string) ([]byte, error) {
	var queryPart string
	if query != "" {
		queryPart = "?" + url.PathEscape(query)
	}
	//fmt.Println(fmt.Sprintf("Q=<%s>", queryPart))
	resp, err := http.Get(fmt.Sprintf("%s/v1/tasks%s", getPath(), queryPart))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			fmt.Println(fmt.Sprintf("response: %s", string(body)))
		}
		return []byte{}, fmt.Errorf("unexpected error code: %v", resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
