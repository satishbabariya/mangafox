package scheduler

import (
	"mangafox/source/mangadex"
	"mangafox/tasks"
	"path"

	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"
)

func Latest(token string, queue *asynq.Client) error {
	md := mangadex.Initilize()
	items, err := md.Latest(token)
	if err != nil {
		logrus.Errorln(err)
		return err
	}

	for _, item := range items {
		chapter := path.Base(item.Link)
		manga := path.Base(item.MangaLink)
		err = IndexChapter(manga, chapter, queue)
		if err != nil {
			logrus.Errorln(err)
		}
	}

	return nil
}

func IndexChapter(manga string, chapter string, queue *asynq.Client) error {
	payload := map[string]interface{}{"manga_id": manga, "chapter_id": chapter}
	task := asynq.NewTask(string(tasks.IndexMangadexChapter), payload)
	err := queue.Enqueue(task)
	return err
}
