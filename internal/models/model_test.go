package models

import (
	"fmt"
	"github.com/ouqiang/gocron/internal/modules/app"
	"github.com/ouqiang/gocron/internal/modules/logger"
	"github.com/ouqiang/gocron/internal/modules/setting"
	"github.com/ouqiang/gocron/internal/modules/utils"
	"testing"
)

func init() {
	fmt.Println("setup")
	app.InitEnv("2.0.0")
	config, err := setting.Read(app.AppConfig)
	if err != nil {
		logger.Fatal("读取应用配置失败", err)
	}
	app.Setting = config

	// 初始化DB
	Db = CreateDb()
}

func TestResetPassword(t *testing.T) {
	//_ = Db.Sync2(new(User))
	user := new(User)
	err := user.Get("admin")
	t.Log(err, user.Id)
	i, err := user.UpdatePassword(user.Id, "Admin@123")
	t.Log(i, err, user.MatchPassword("Admin@123"))
}

func TestCreateTable(t *testing.T) {
	err := Db.CreateTables(Role{})
	//err := Db.Sync2(new(Role))
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAlertTable(t *testing.T) {
	results, err := Db.Query("alter table task add project_id int default 0 not null;")
	t.Log(results, err)
}

func TestProjectSetHosts(t *testing.T) {
	projects := make([]Project, 0)
	_ = Db.Find(&projects)
	p := Project{}
	projects = p.setHostsForProjects(projects)
	t.Log(projects)
}

func TestProcessWorker_GetByProcess(t *testing.T) {
	pw := ProcessWorker{}
	workers, err := pw.GetByProcess(Process{Id: 1})
	t.Log(err)
	t.Log(workers)
}

func TestProcessHost_GetByProcess(t *testing.T) {
	ph := ProcessHost{}
	hosts := ph.GetByProcess(Process{Id: 1})
	//t.Log(err)
	t.Log(hosts)
}

func TestProcessHost_DeleteForProcess(t *testing.T) {
	ph := ProcessHost{}
	ph.DeleteForProcess(Process{Id: 1})
}

func TestGetTasks(t *testing.T) {
	task := Task{}
	tasks, _ := task.List(CommonMap{})
	for _, task := range tasks {
		t.Log(task.Status)
	}
	t.Log(utils.JsonResp.Success("", tasks))
}

func TestTaskHost_GetHostsByTaskId(t *testing.T) {
	th := TaskHost{}
	hosts, err := th.GetHostsByTaskId(1)
	t.Log(hosts, err)
}

func TestGetActionUsers(t *testing.T) {
	_, err := Db.Query(fmt.Sprintf("alter table `%suser` add source varchar(16) default 'system' not null after email;", TablePrefix))
	t.Log(err)

}

func TestProjectHost_GetHostsByProjectId(t *testing.T) {
	ph := ProjectHost{}
	t.Log(ph.GetHostsByProjectId(1))
}

func TestTaskLog_Clear(t *testing.T) {
	l := TaskLog{}

	t.Log(l.Clear(CommonMap{"taskId": "1"}))
}

func TestUser_Get(t *testing.T) {
	user := new(User)
	err := user.Get("ad")

	t.Log(err, user.Id)
}
