package database

import "basic/models"

// SystemLevelInit 系统等级初始化
func SystemLevelInit() []*models.Level {
	return []*models.Level{
		{AdminId: 1, Level: 1, Name: "level1Label", Icon: "/assets/level/diamond.png", Money: 120, Days: -1, Status: 10, Desc: "level1Desc"},
		{AdminId: 1, Level: 2, Name: "level2Label", Icon: "/assets/level/diamond.png", Money: 220, Days: -1, Status: 10, Desc: "level2Desc"},
		{AdminId: 1, Level: 3, Name: "level3Label", Icon: "/assets/level/diamond.png", Money: 320, Days: -1, Status: 10, Desc: "level3Desc"},
		{AdminId: 1, Level: 4, Name: "level4Label", Icon: "/assets/level/diamond.png", Money: 580, Days: -1, Status: 10, Desc: "level4Desc"},
		{AdminId: 1, Level: 5, Name: "level5Label", Icon: "/assets/level/diamond.png", Money: 888, Days: -1, Status: 10, Desc: "level5Desc"},
	}
}
