package model

type Subscriber struct {
	ChatId		int64	`gorm:"chat_id;primary_key"`
	UserName	string	`gorm:"user_name"`
	Repositories	[]Repository	`gorm:"many2many:repository_subscribers"`
}

func CreateSubscriber(s *Subscriber) {
	DB.Save(s)
}

func DeleteSubscriber(s *Subscriber)  {
	DeleteAllSubscriberWithRepo(s)
	DB.Delete(&s)
}

func DeleteAllSubscriberWithRepo(s *Subscriber) {
	DB.Model(&s).Association("repositories").Clear()
}

func UpdateSubscriber(s *Subscriber) {
	DB.Save(s)
}

func FindSubscriberByChatID(id int64) (s *Subscriber, err error) {
	s = new(Subscriber)
	if err = DB.Where("chat_id=?", id).First(s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

func FindAllSubscribers() (sList []*Subscriber, err error) {
	if err = DB.Find(&sList).Error; err != nil {
		return nil, err
	}
	return sList, err
}

func FindSubscribersByRepo(repo *Repository) (sList []*Subscriber, err error) {
	if err = DB.Model(&repo).Association("subscribers").Find(&sList).Error; err != nil {
		return nil, err
	}
	return sList, nil
}
