/*
 * @Date: 2022-11-04 13:54:42
 * @LastEditors: Jason Chen
 * @LastEditTime: 2022-11-09 14:16:13
 * @FilePath: /cli/internal/model/error.go
 */

package model

import "errors"

type FeedbackError struct {
	IsRetry bool  `json:"isRetry"`
	Err     error `json:"err"`
}

func (e *FeedbackError) Error() string {
	/* msg := ""
	if e.IsRetry {
		msg = "disable retry"
	} */
	return e.Err.Error()
}

func CreateFeedbackError(err string, isRetry bool) FeedbackError {
	return FeedbackError{
		IsRetry: isRetry,
		Err:     errors.New(err),
	}
}

func CreateFeedbackError2(err string, isRetry bool) error {
	tmpErr := FeedbackError{
		IsRetry: isRetry,
		Err:     errors.New(err),
	}
	return &tmpErr
}
