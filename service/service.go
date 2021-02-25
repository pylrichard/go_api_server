package service

import (
	"fmt"
	"sync"

	"go/tiny_http_server/model"
	"go/tiny_http_server/util"
)

// ListUser 获取用户列表
func ListUser(name string, offset, limit int) ([]*model.UserInfo, uint64, error) {
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(name, offset, limit)
	if err != nil {
		return nil, count, err
	}

	ids := []uint64{}
	for _, user := range users {
		ids = append(ids, user.Id)
	}

	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		Infos: make(map[uint64]*model.UserInfo, len(users)),
	}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// 并行查询
	for _, u := range users {
		wg.Add(1)
		go func(u *model.UserModel) {
			defer wg.Done()

			shortID, err := util.GenShortID()
			if err != nil {
				errChan <- err
				return
			}

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.Infos[u.Id] = &model.UserInfo{
				Id:          u.Id,
				UserName:    u.UserName,
				SayHello:    fmt.Sprintf("Hello %s", shortID),
				Password:    u.Password,
				CreatedTime: u.CreatedTime.Format("2021-02-25 15:00:00"),
				UpdatedTime: u.UpdatedTime.Format("2021-02-25 15:00:00"),
			}
		}(u)
	}

	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	for _, id := range ids {
		infos = append(infos, userList.Infos[id])
	}

	return infos, count, nil
}
