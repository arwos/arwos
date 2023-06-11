package agentsrv

import (
	"context"
	"fmt"
	"strings"

	"github.com/deweppro/go-sdk/orm"
)

type AccessDBM struct {
	ID    uint64
	Name  string
	Token string
}

func (v *AgentSrv) selectAccess() ([]AccessDBM, error) {
	var result []AccessDBM
	err := v.db.Pool().QueryContext("select_agents_access", context.TODO(), func(q orm.Querier) {
		q.SQL("SELECT `id`,`name`,`token` FROM `agents_access`;")
		q.Bind(func(bind orm.Scanner) error {
			model := AccessDBM{}
			if err := bind.Scan(&model.ID, &model.Name, &model.Token); err != nil {
				return err
			}
			result = append(result, model)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (v *AgentSrv) selectAccessByToken(token string) (uint64, error) {
	var id uint64
	err := v.db.Pool().QueryContext("select_agents_access", context.TODO(), func(q orm.Querier) {
		q.SQL("SELECT `id` FROM `agents_access` WHERE `token` = ? LIMIT 1;", token)
		q.Bind(func(bind orm.Scanner) error {
			return bind.Scan(&id)
		})
	})
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, fmt.Errorf("not found")
	}
	return id, nil
}

func (v *AgentSrv) insertAccess(name, token string) (uint64, error) {
	var id uint64
	err := v.db.Pool().ExecContext("insert_agents_access", context.TODO(), func(q orm.Executor) {
		q.SQL("INSERT INTO `agents_access` (`name`, `token`) VALUES (?, ?);", name, token)
		q.Bind(func(result orm.Result) error {
			if result.LastInsertId == 0 || result.RowsAffected == 0 {
				return fmt.Errorf("fail add access")
			}
			id = uint64(result.LastInsertId)
			return nil
		})
	})
	return id, err
}

func (v *AgentSrv) deleteAccess(token string) error {
	return v.db.Pool().ExecContext("delete_agents_access", context.TODO(), func(q orm.Executor) {
		q.SQL("DELETE FROM `agents_access` WHERE `token` = ?;", token)
		q.Bind(func(result orm.Result) error {
			if result.RowsAffected == 0 {
				return fmt.Errorf("fail delete access")
			}
			return nil
		})
	})
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type TagDBM struct {
	ID   uint64
	Name string
}

func (v *AgentSrv) selectTags() ([]TagDBM, error) {
	var result []TagDBM
	err := v.db.Pool().QueryContext("select_agents_tags", context.TODO(), func(q orm.Querier) {
		q.SQL("SELECT `id`,`name` FROM `agents_tags`;")
		q.Bind(func(bind orm.Scanner) error {
			tag := TagDBM{}
			if err := bind.Scan(&tag.ID, &tag.Name); err != nil {
				return err
			}
			result = append(result, tag)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (v *AgentSrv) selectTag(tag string) (uint64, error) {
	var id uint64
	err := v.db.Pool().QueryContext("select_agents_tag", context.TODO(), func(q orm.Querier) {
		q.SQL("SELECT `id` FROM `agents_tags` WHERE `name` = ? LIMIT 1;", strings.ToLower(tag))
		q.Bind(func(bind orm.Scanner) error {
			return bind.Scan(&id)
		})
	})
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, fmt.Errorf("not found")
	}
	return id, nil
}

func (v *AgentSrv) insertTag(tag string) (uint64, error) {
	var id uint64
	err := v.db.Pool().ExecContext("insert_agents_tag", context.TODO(), func(q orm.Executor) {
		q.SQL("INSERT INTO `agents_tags` (`name`) VALUES (?);", strings.ToLower(tag))
		q.Bind(func(result orm.Result) error {
			if result.LastInsertId == 0 || result.RowsAffected == 0 {
				return fmt.Errorf("fail add tag")
			}
			id = uint64(result.LastInsertId)
			return nil
		})
	})
	return id, err
}

func (v *AgentSrv) deleteTag(tag string) error {
	return v.db.Pool().ExecContext("delete_agents_tag", context.TODO(), func(q orm.Executor) {
		q.SQL("DELETE FROM `agents_tags` WHERE `name` = ?;", strings.ToLower(tag))
		q.Bind(func(result orm.Result) error {
			if result.RowsAffected == 0 {
				return fmt.Errorf("fail delete tag")
			}
			return nil
		})
	})
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (v *AgentSrv) insertTagLink(aid, tid uint64) error {
	return v.db.Pool().ExecContext("insert_agents_tag_link", context.TODO(), func(q orm.Executor) {
		q.SQL("INSERT INTO `agents_tag_link` (`agent_id`, `tag_id`) VALUES (?,?);", aid, tid)
		q.Bind(func(result orm.Result) error {
			if result.LastInsertId == 0 || result.RowsAffected == 0 {
				return fmt.Errorf("fail add tag link")
			}
			return nil
		})
	})
}

func (v *AgentSrv) deleteTagLink(aid, tid uint64) error {
	return v.db.Pool().ExecContext("delete_agents_tag_link", context.TODO(), func(q orm.Executor) {
		q.SQL("DELETE FROM `agents_tag_link` WHERE `agent_id` = ? AND `tag_id` = ?;", aid, tid)
		q.Bind(func(result orm.Result) error {
			if result.RowsAffected == 0 {
				return fmt.Errorf("fail delete tag link")
			}
			return nil
		})
	})
}

func (v *AgentSrv) selectTagLink(aid uint64) ([]TagDBM, error) {
	var result []TagDBM
	err := v.db.Pool().QueryContext("select_agents_tag_link", context.TODO(), func(q orm.Querier) {
		q.SQL("SELECT t.`id`,t.`name` FROM `agents_tags` AS t "+
			"JOIN `agents_tag_link` AS l ON l.`tag_id` = t.`id` "+
			"WHERE l.`agent_id` = ?;", aid)
		q.Bind(func(bind orm.Scanner) error {
			tag := TagDBM{}
			if err := bind.Scan(&tag.ID, &tag.Name); err != nil {
				return err
			}
			result = append(result, tag)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
