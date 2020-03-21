package feedback

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Hack-the-Crisis-got-milk/Shops/config"
	"github.com/Hack-the-Crisis-got-milk/Shops/entities"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	logger      *zap.Logger
	httpClient  *http.Client
	apiEndpoint string
}

type request struct {
	path   string
	method string
}

const (
	PING_ACTIVITY                   = "ping"
	GET_FEEDBACK_FOR_SHOPS_ACTIVITY = "getFeedbackForShop"

	PING_PATH                   = "/kristutei"
	GET_FEEDBACK_FOR_SHOPS_PATH = "/getFeedbackForShops"
)

var feedbackRequests = map[string]request{
	PING_ACTIVITY: {
		path:   PING_PATH,
		method: http.MethodGet,
	},
	GET_FEEDBACK_FOR_SHOPS_ACTIVITY: {
		path:   GET_FEEDBACK_FOR_SHOPS_PATH,
		method: http.MethodGet,
	},
}

var ErrFeedbackUnavailable = errors.New("could not reach feedback service")

func NewClient(logger *zap.Logger, cfg *config.AppConfig) (*Client, error) {
	client := &Client{
		logger:      logger,
		httpClient:  &http.Client{},
		apiEndpoint: cfg.FeedbackServiceEndpoint,
	}

	res, err := client.sendHttpRequest(feedbackRequests[PING_ACTIVITY], nil, nil)
	if err != nil {
		logger.Error("could not ping feedback", zap.Error(err))
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		logger.Error("ping to feedback failed", zap.Int("returned status code", res.StatusCode))
		return nil, ErrFeedbackUnavailable
	}

	return client, nil
}

func (c *Client) GetFeedbackForShops(shopIds []string) (map[string][]entities.Feedback, error) {
	start := time.Now()

	res, err := c.sendHttpRequest(feedbackRequests[GET_FEEDBACK_FOR_SHOPS_ACTIVITY], map[string]interface{}{
		"shopIds": shopIds,
	}, nil)
	if err != nil {
		c.logger.Error("could not call GetFeedbackForShops", zap.Error(err))
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.logger.Error("could not read response from GetFeedbackForShops", zap.Error(err))
		return nil, err
	}

	feedbackMap, err := convertGetFeedbackForShopsResponseToFeedbackMap(body)
	if err != nil {
		c.logger.Error("could not parse response from GetFeedbackForShops", zap.String("resposne", string(body)), zap.Error(err))
		return nil, err
	}

	fmt.Println("get feedback for shops time: ", time.Now().Sub(start))

	return feedbackMap, nil
}

func (c *Client) sendHttpRequest(req request, urlParams map[string]interface{}, body io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(req.method, c.apiEndpoint+req.path+c.constructUrlParams(urlParams), body)
	if err != nil {
		return nil, err
	}

	return c.httpClient.Do(request)
}

func (c *Client) constructUrlParams(params map[string]interface{}) string {
	if params == nil {
		return ""
	}

	urlParams := "?"
	for key, value := range params {
		jsonValue, err := json.Marshal(value)
		if err != nil {
			c.logger.Error(fmt.Sprintf("could not marshal %s with value %s to JSON", key, value), zap.Error(err))
		} else {
			urlParams += fmt.Sprintf("%s=%s&", key, url.QueryEscape(string(jsonValue)))
		}
	}

	return urlParams[:len(urlParams)-1]
}

func convertGetFeedbackForShopsResponseToFeedbackMap(response []byte) (map[string][]entities.Feedback, error) {
	var wrapper struct {
		Response map[string][]entities.Feedback `json:"response"`
	}

	err := json.Unmarshal(response, &wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.Response, nil
}
