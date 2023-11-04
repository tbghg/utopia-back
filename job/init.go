package job

import "utopia-back/database/implement"

func JobInit() {
	centerDal := implement.NewCenterDal()
	videoJobInit(centerDal)
}
