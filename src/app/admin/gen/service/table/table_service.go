// Package table 自动生成模板 TableService
// @description <TODO description class purpose>
// @author
// @File: table
// @version 1.0.0
// @create 2023-08-31 09:09:53
package table

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"go.uber.org/zap"
	"matuto-base/src/app/admin/gen/dao"
	"matuto-base/src/app/admin/gen/model"
	"matuto-base/src/app/admin/gen/service/table/view"
	columm_service "matuto-base/src/app/admin/gen/service/table_column"
	genutils "matuto-base/src/app/admin/gen/utils"
	"matuto-base/src/common"
	"matuto-base/src/common/constants"
	"matuto-base/src/global"
	"matuto-base/src/utils"
	"strings"
	"text/template"
	"time"
)

// Service 结构体
type Service struct {
	tableDao      dao.TableDao
	viewUtils     view.TableViewUtils
	columnDao     dao.TableColumnDao
	columnService columm_service.TableColumnService
}

// Create 创建Table记录
// Author
func (s *Service) Create(tableView *view.TableView) error {
	if err, table := s.viewUtils.View2Data(tableView); err != nil {
		return err
	} else {
		return s.tableDao.Create(global.GormDao, table)
	}
}

// DeleteByIds 批量删除Table记录
// Author
func (s *Service) DeleteByIds(ids []string) (err error) {
	tx := global.GormDao.Begin()
	if err = s.tableDao.DeleteByIds(tx, ids); err != nil {
		tx.Rollback()
		return err
	} else {
		// 删除列信息
		if err := s.columnDao.DeleteByTableIds(tx, ids); err != nil {
			tx.Rollback()
			return err
		} else {
			tx.Commit()
			return nil
		}
	}
}

// Update 更新Table记录
// Author
func (s *Service) Update(tableView *view.TableView) error {
	tx := global.GormDao.Begin()
	// 更新
	options := view.TableViewOptions{
		TreeCode:       tableView.TreeCode,
		TreeParentCode: tableView.TreeParentCode,
		TreeName:       tableView.TreeName,
		ParentMenuId:   tableView.ParentMenuId,
		ParentMenuName: tableView.ParentMenuName,
	}
	if jsonBytes, err := json.Marshal(options); err != nil {
		tx.Rollback()
		return err
	} else {
		tableView.Options = string(jsonBytes)
	}
	if err, table := s.viewUtils.View2Data(tableView); err != nil {
		tx.Rollback()
		return err
	} else {
		if err := s.tableDao.Update(tx, table); err != nil {
			tx.Rollback()
			return err
		}
		// 修改列信息
		for _, columnView := range tableView.ColumnList {
			if err := s.columnService.Update(columnView, tx); err != nil {
				tx.Rollback()
				return err
			}
		}
		tx.Commit()
		return nil
	}
}

// Get 根据id获取Table记录
// Author
func (s *Service) Get(id string) (err error, tableView *view.TableView) {
	err1, table := s.tableDao.Get(id)
	if err1 != nil {
		return err1, nil
	}
	err, tableView = s.viewUtils.Data2View(table)
	// 通过id查询列信息
	if err, tableView.ColumnList = s.columnService.GetColumnListByTableId(id); err != nil {
		global.Logger.Error("GetColumnListByTableId is error ", zap.Error(err))
	}
	// options 转字段
	if tableView != nil && tableView.Options != "" {
		var tableOption view.TableViewOptions
		if err := json.Unmarshal([]byte(tableView.Options), &tableOption); err != nil {
			global.Logger.Error("TableOption Convert is error ", zap.Error(err))
		} else {
			tableView.TreeCode = tableOption.TreeCode
			tableView.TreeParentCode = tableOption.TreeParentCode
			tableView.TreeName = tableOption.TreeName
			tableView.ParentMenuId = tableOption.ParentMenuId
			tableView.ParentMenuName = tableOption.ParentMenuName
		}
	}
	return
}

// Page 分页获取Table记录
// Author
func (s *Service) Page(pageInfo *view.TablePageView) (err error, res *common.PageInfo) {
	if err, res = s.tableDao.Page(pageInfo); err != nil {
		return err, nil
	} else {
		return s.viewUtils.PageData2ViewList(res)
	}
}

// List 获取Table列表
// Author
func (s *Service) List(v *view.TableQueryView) (error, []*view.TableView) {
	if err, dataList := s.tableDao.List(v); err != nil {
		return err, nil
	} else {
		return s.viewUtils.Data2ViewList(dataList)
	}
}

// SelectDbTablePage 获取数据库表列表
func (s *Service) SelectDbTablePage(v *view.TablePageView) (err error, res *common.PageInfo) {
	if err, res = s.tableDao.SelectDbTablePage(v); err != nil {
		return err, nil
	} else {
		return s.viewUtils.PageData2ViewList(res)
	}
}

// ImportTable 导入Table
func (s *Service) ImportTable(tables string, loginUser string) error {
	tableNames := strings.Split(tables, ",")
	if len(tableNames) == 0 {
		return nil
	}
	if err, tables := s.tableDao.SelectDbTableList(tableNames); err != nil {
		return err
	} else {
		return s.ImportGenTable(tables, loginUser)
	}
}

// ImportGenTable 导入GenTable
func (s *Service) ImportGenTable(tables []*model.Table, loginUser string) error {
	if len(tables) == 0 {
		return nil
	}
	tx := global.GormDao.Begin()
	for _, table := range tables {
		table = genutils.InitTable(table, loginUser)
		table.Id = utils.GenUID()
		if err := s.tableDao.Create(tx, table); err != nil {
			tx.Rollback()
			return err
		}
		// 处理列信息
		if err, tableColumns := s.columnDao.SelectDbTableColumns(tx, table.Name); err != nil {
			tx.Rollback()
			return err
		} else {
			for _, column := range tableColumns {
				column = genutils.InitColumnField(column, table)
				column.Id = utils.GenUID()
				if err := s.columnDao.Create(tx, column); err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}
	tx.Commit()
	return nil
}

// ValidateEdit 表单验证
func (s *Service) ValidateEdit(v *view.TableView) error {
	if v.TplCategory == constants.TPL_TREE {
		if v.TreeCode == "" {
			return errors.New("树编码不能为空")
		}
		if v.TreeName == "" {
			return errors.New("树名称不能为空")
		}
		if v.TreeParentCode == "" {
			return errors.New("树父编码不能为空")
		}
	} else if v.TplCategory == constants.TPL_SUB {
		if v.SubTableName == "" {
			return errors.New("子表名称不能为空")
		}
		if v.SubTableFkName == "" {
			return errors.New("子表外键名称不能为空")
		}
	}
	return nil
}

// PreviewCode 预览代码
func (s *Service) PreviewCode(id string) (err error, dataMap map[string]string) {
	if err, tableView := s.SelectGenTableById(id); err != nil {
		return err, nil
	} else {
		// 塞入作者
		tableView.Author = global.Viper.GetString("gen.author")
		tableView.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		// 塞入字典
		tableView.Dicts = strings.Join(genutils.GetDictList(tableView), ",")
		if tableView.TplCategory == constants.TPL_TREE {
			if err, dataMap = s.PreviewTreeCode(tableView); err != nil {
				return err, nil
			}
		} else if tableView.TplCategory == constants.TPL_SUB || tableView.TplCategory == constants.TPL_CRUD || tableView.TplCategory == "" {
			if err, dataMap = s.PreviewSubTable(tableView); err != nil {
				return err, nil
			}
		}
		return nil, dataMap
	}
}

// PreviewTreeCode 预览树编码
func (s *Service) PreviewTreeCode(tableView *view.TableView) (err error, dataMap map[string]string) {
	return nil, nil
}

// SelectGenTableById 根据id获取GenTable包含各种列信息
func (s *Service) SelectGenTableById(id string) (err error, tableView *view.TableView) {
	err, v := s.Get(id)
	if err != nil {
		return err, nil
	} else {
		// 查询主键列
		err, pkColumn := s.columnService.SelectPkColumn(id)
		if err != nil {
			return err, nil
		}
		// 查询列
		err, searchColumn := s.columnService.SelectSearchColumn(id)
		if err != nil {
			return err, nil
		}
		// 新增列
		err, insertColumn := s.columnService.SelectInsertColumn(id)
		if err != nil {
			return err, nil
		}
		// 修改列
		err, editColumn := s.columnService.SelectEditColumn(id)
		if err != nil {
			return err, nil
		}
		// 列表列
		err, listColumn := s.columnService.SelectListColumn(id)
		if err != nil {
			return err, nil
		}
		v.PKColumn = pkColumn
		v.SearchColumn = searchColumn
		v.InsertColumn = insertColumn
		v.EditColumn = editColumn
		v.ListColumn = listColumn
		return nil, v
	}
}

// PreviewSubTable 预览子表
func (s *Service) PreviewSubTable(tableView *view.TableView) (err error, dataMap map[string]string) {
	dataMap = make(map[string]string)
	var templatePath = genutils.GenTemplatePath(tableView.TplCategory)
	for _, path := range templatePath {
		tmpl := template.Must(template.ParseFiles(path))
		var buf bytes.Buffer
		if err = tmpl.Execute(&buf, tableView); err != nil {
			return err, nil
		}
		dataMap[path] = buf.String()
	}
	return
}

// GenCode 生成代码
func (s *Service) GenCode(id string) (error, *bytes.Buffer, string) {
	if err, tableView := s.SelectGenTableById(id); err != nil {
		return err, nil, ""
	} else {
		err, dataMap := s.PreviewCode(id)
		if err != nil {
			return errors.New("生成代码文件失败"), nil, ""
		} else {
			zipName := tableView.BusinessName + ".zip"
			// 创建一个bytes.Buffer来写入zip文件内容
			zipBuffer := new(bytes.Buffer)
			// 使用zip.Writer创建zip文件写入器
			zipWriter := zip.NewWriter(zipBuffer)
			for key, value := range dataMap {
				// 去除 txt 后缀
				newKey := key[:strings.LastIndex(key, ".")]
				// 去除原路径
				newKey = newKey[strings.LastIndex(newKey, "/")+1:]
				// 获取文件类型
				fileType := newKey[strings.LastIndex(newKey, ".")+1:]
				filePath := newKey[:strings.LastIndex(newKey, ".")]
				if strings.Contains(filePath, "convert") {
					filePath = fileType + "/service/convert/" + tableView.BusinessName + "_" + newKey
				} else {
					if fileType == "vue" {
						filePath = fileType + "/" + tableView.BusinessName + "/" + newKey
					} else {
						filePath = fileType + "/" + filePath + "/" + tableView.BusinessName + "." + fileType
					}
				}
				zipFile, err := zipWriter.Create(filePath)
				if err != nil {
					zipWriter.Close()
					return errors.New("创建zip文件失败"), nil, ""
				}
				_, err = zipFile.Write([]byte(value))
				if err != nil {
					zipWriter.Close()
					return errors.New("写入zip文件失败"), nil, ""
				}
			}
			// 关闭zip文件写入器
			zipWriter.Close()
			return nil, zipBuffer, zipName
		}
	}
}
