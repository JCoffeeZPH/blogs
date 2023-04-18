package cache

import "blogs/common/constants"

func ZIncrArticleRank(articleName string) {
	redisClient.ZIncrBy(constants.ArticleRankKey, 1, articleName)
}
