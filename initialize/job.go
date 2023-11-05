package initialize

import (
	"utopia-back/database/implement"
	"utopia-back/job"
)

func JobInit() {
	centerDal := implement.NewCenterDal()
	job.VideoJobInit(centerDal)
}
