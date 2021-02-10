package main

import (
	"fmt"
	"strconv"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ResultData struct {
	GroupID       string
	TotalData     int
	UserCompanyID int
}

type ResultConvert struct {
	GroupID       []int
	TotalData     int
	UserCompanyID int
}

type UserData struct {
	ID            int
	UserCompanyID int
	UserName      string
	UserEmail     string
}

func main() {
	var result []ResultData
	var allResult []ResultConvert
	var res ResultConvert
	dsn := "root:@tcp(127.0.0.1:3306)/demo_lumen?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	err = db.Raw("SELECT GROUP_CONCAT(id) AS group_id, COUNT(id) AS total_data,user_company_id FROM user_data GROUP BY user_company_id").Scan(&result).Error
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range result {
		res.TotalData = item.TotalData
		res.UserCompanyID = item.UserCompanyID
		splited := strings.Split(item.GroupID, ",")
		for j := 0; j < len(splited); j++ {
			num, _ := strconv.Atoi(splited[j])
			res.GroupID = append(res.GroupID, num)
		}
		allResult = append(allResult, res)
		res.GroupID = nil
	}
	fmt.Println(allResult[0])
	fmt.Println(allResult[1])
	fmt.Println(allResult[2])
}
