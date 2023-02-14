package models

import (
	"fmt"
	"github.com/BeardLeon/tiktok/pkg/setting"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Model base class
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

// Setup gorm db initialize
func Setup() {
	// var (
	//	err                                               error
	//	dbType, dbName, user, password, host, tablePrefix string
	// )
	//
	// sec, err := setting.Cfg.GetSection("database")
	// if err != nil {
	//	log.Fatal(2, "Fail to get section 'database': %v", err)
	// }
	//
	// dbType = sec.Key("TYPE").String()
	// dbName = sec.Key("NAME").String()
	// user = sec.Key("USER").String()
	// password = sec.Key("PASSWORD").String()
	// host = sec.Key("HOST").String()
	// tablePrefix = sec.Key("TABLE_PREFIX").String()
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Println(err)
	}

	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return setting.DatabaseSetting.TablePrefix + defaultTableName
	// }

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", updateTimeStampForDeleteCallback)
}

func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		// scope.FieldByName 通过 scope.Fields() 获取所有字段，判断当前是否包含所需字段
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			// field.IsBlank 可判断该字段的值是否为空
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifyOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	// scope.Get(...) 根据入参获取设置了字面值的参数，
	// 例如本文中是 gorm:update_column ，它会去查找含这个字面值的字段属性
	if _, ok := scope.Get("gorm:update_column"); !ok {
		// scope.SetColumn(...)
		// 假设没有指定 update_column 的字段，我们默认在更新回调设置 ModifiedOn 的值
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// updateTimeStampForDeleteCallback will set `DeleteOn` when deleting
func updateTimeStampForDeleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		// scope.Get("gorm:delete_option") 检查是否手动指定了 delete_option
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		// scope.FieldByName("DeletedOn") 获取我们约定的删除字段，
		// 若存在则 UPDATE 软删除，若不存在则 DELETE 硬删除
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

		// 软删除 UPDATE
		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				// scope.QuotedTableName() 返回引用的表名，
				// 这个方法 GORM 会根据自身逻辑对表名进行一些处理
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),

				// scope.AddToVars 该方法可以添加值作为 SQL 的参数，也可用于防范 SQL 注入
				// func (scope *Scope) AddToVars(value interface{}) string {
				//    _, skipBindVar := scope.InstanceGet("skip_bindvar")
				//
				//    if expr, ok := value.(*expr); ok {
				//        exp := expr.expr
				//        for _, arg := range expr.args {
				//            if skipBindVar {
				//                scope.AddToVars(arg)
				//            } else {
				//                exp = strings.Replace(exp, "?", scope.AddToVars(arg), 1)
				//            }
				//        }
				//        return exp
				//    }
				//
				//    scope.SQLVars = append(scope.SQLVars, value)
				//
				//    if skipBindVar {
				//        return "?"
				//    }
				//    return scope.Dialect().BindVar(len(scope.SQLVars))
				// }
				scope.AddToVars(time.Now().Unix()),

				// scope.CombinedConditionSql() 返回组合好的条件 SQL
				// func (scope *Scope) CombinedConditionSql() string {
				//    joinSQL := scope.joinsSQL()
				//    whereSQL := scope.whereSQL()
				//    if scope.Search.raw {
				//        whereSQL = strings.TrimSuffix(strings.TrimPrefix(whereSQL, "WHERE ("), ")")
				//    }
				//    return joinSQL + whereSQL + scope.groupSQL() +
				//        scope.havingSQL() + scope.orderSQL() + scope.limitAndOffsetSQL()
				// }
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			// 硬删除 DELETE
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
