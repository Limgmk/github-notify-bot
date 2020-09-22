package model

func CreateRepo(r *Repository) {
	DB.Save(r)
}

func DeleteRepo(r *Repository) {
	DeleteAllRepoWithSubscriber(r)
	DB.Delete(r)
}

func DeleteAllRepoWithSubscriber(r *Repository) {
	DB.Model(&r).Association("subscribers").Clear()
}

func DeleteRepoWithSubscriber(r *Repository, s *Subscriber) {
	DB.Model(&r).Association("subscribers").Delete(&s)
}

func AddRepoWithSubscriber(r *Repository, s *Subscriber) {
	var sList []*Subscriber
	sList = append(sList, s)
	DB.Model(&r).Association("subscribers").Append(sList)
}

func UpdateRepo(r *Repository) {
	DB.Save(r)
}

func FindRepoByFullName(fullName string) (r *Repository, err error) {
	r = new(Repository)
	if err = DB.Where("full_name=?", fullName).First(r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func FindAllRepos() (rList []*Repository, err error) {
	if err = DB.First(&rList).Error; err != nil {
		return nil, err
	}
	return rList, err
}

func FindReposBySubscriber(s *Subscriber) (rList []*Repository, err error) {
	if err = DB.Model(&s).Association("repositories").Find(&rList).Error; err != nil {
		return nil, err
	}
	return rList, nil
}

//func FindRepoWithSubscriber(r *Repository, s *Subscriber) (r *Repository, s *Subscriber) {
//
//}