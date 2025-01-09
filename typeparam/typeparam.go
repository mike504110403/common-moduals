package typeparam

import (
	"database/sql"
	"errors"
)

// 基於資料庫的參數管理套件

var cfg = Config{
	FuncGetDB: func() (*sql.DB, error) { return nil, errors.New("func get db is not set") },
}
var mainTypeMap = MainTypeMap{}

type Config struct {
	FuncGetDB func() (*sql.DB, error)
}

type (
	MainType    string
	SubType     string
	SubTypeID   int
	MainTypeMap map[MainType]SubTypeMap
	SubTypeMap  map[SubType]SubTypeID
	TypeParam   struct {
		MainType MainType
		SubType  SubType
	}
)

type (
	TypeGetData struct {
		MainTypeID              int
		MainTypeComparisonValue string
		MainTypeDescription     string
		SubTypeypeID            int
		SubTypeComparisonValue  string
		SubTypeDescription      string
	}
)

func Init(initCfg Config) error {
	cfg = initCfg

	return load()
}

func load() error {
	if db, err := cfg.FuncGetDB(); err != nil {
		return err
	} else {
		typeGetDataList := []TypeGetData{}
		if rows, err := db.Query("CALL Type_Get(?, ?)", nil, nil); err != nil {
			return err
		} else {
			for rows.Next() {
				data := TypeGetData{}
				if err := rows.Scan(
					&data.MainTypeID,
					&data.MainTypeComparisonValue,
					&data.MainTypeDescription,
					&data.SubTypeypeID,
					&data.SubTypeComparisonValue,
					&data.SubTypeDescription,
				); err != nil {
					return err
				} else {
					typeGetDataList = append(typeGetDataList, data)
				}
			}
			// 整理資料
			for _, data := range typeGetDataList {
				if mainTypeMap[MainType(data.MainTypeComparisonValue)] == nil {
					mainTypeMap[MainType(data.MainTypeComparisonValue)] = SubTypeMap{}
				}
				mainTypeMap[MainType(data.MainTypeComparisonValue)][SubType(data.SubTypeComparisonValue)] = SubTypeID(data.SubTypeypeID)
			}

		}
	}

	return nil
}

// MainType.Map : 取得MainType的SubMap
func (c MainType) Map() (SubTypeMap, error) {
	if subTypeMap, isExist := mainTypeMap[c]; !isExist {
		return SubTypeMap{}, errors.New("class not exist")
	} else {
		return subTypeMap, nil
	}
}

// TypeParam.Get : 取得Item的實際參數值
func (p TypeParam) Get() (int, error) {
	if subTypeMap, err := p.MainType.Map(); err != nil {
		return 0, err
	} else {
		if subTypeID, isExist := subTypeMap[p.SubType]; !isExist {
			return 0, errors.New("param not exist")
		} else {
			return int(subTypeID), nil
		}
	}
}

func GetMap() MainTypeMap {
	return mainTypeMap
}

// Find : 找到這個參數值的類別與項目，以供判斷
func Find(id int) TypeParam {
	for mainType, subTypeMap := range mainTypeMap {
		for subType, subTypeID := range subTypeMap {
			if subTypeID == SubTypeID(id) {
				return TypeParam{
					MainType: mainType,
					SubType:  subType,
				}
			}
		}
	}

	return TypeParam{}
}

// GetReversedMap: Map鍵值反轉，吐回Map給多資料對應subtype用
func GetReversedMap(subTypeMap SubTypeMap) map[int]string {
	reversed := make(map[int]string)
	for key, value := range subTypeMap {
		reversed[int(value)] = string(key)
	}
	return reversed
}
