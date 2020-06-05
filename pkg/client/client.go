package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Client is ...
type Client struct {
	space  string
	apiKey string
}

// NewClient is ...
func NewClient(space string, apiKey string) *Client {
	cli := new(Client)
	cli.space = space
	cli.apiKey = apiKey

	return cli
}

// Get - HTTP GET Request.
func (cli *Client) Get(api string, values url.Values) ([]byte, error) {
	format := fmt.Sprintf("%s%s?apiKey=%s", cli.space, api, cli.apiKey)
	var url string
	if values == nil {
		url = format
	} else {
		url = format + "&" + values.Encode()
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//DEBUG
	//fmt.Println(req.URL)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var r ErrorResponse
		if err = json.Unmarshal(body, &r); err != nil {
			return nil, err
		}
		err = fmt.Errorf("Error: StatusCode:%d Code:%d Message: %s MoreInfo:%s",
			res.StatusCode,
			r.Errors[0].Code,
			r.Errors[0].Message,
			r.Errors[0].MoreInfo)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Post - HTTP POST Request.
func (cli *Client) Post(api string, values url.Values) ([]byte, error) {
	uri := fmt.Sprintf("%s%s?apiKey=%s", cli.space, api, cli.apiKey)

	req, err := http.NewRequest(
		"POST",
		uri,
		strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	//DEBUG
	//fmt.Println(req.URL)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var r ErrorResponse
		if err = json.Unmarshal(body, &r); err != nil {
			return nil, err
		}
		err = fmt.Errorf("Error: StatusCode:%d Code:%d Message: %s MoreInfo:%s",
			res.StatusCode,
			r.Errors[0].Code,
			r.Errors[0].Message,
			r.Errors[0].MoreInfo)
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Patch - HTTP PATCH Request
func (cli *Client) Patch(api string, values url.Values) ([]byte, error) {
	uri := fmt.Sprintf("%s%s?apiKey=%s", cli.space, api, cli.apiKey)

	req, err := http.NewRequest(
		"PATCH",
		uri,
		strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}

	//DEBUG
	//fmt.Println(req.URL)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var r ErrorResponse
		if err = json.Unmarshal(body, &r); err != nil {
			return nil, err
		}
		err = fmt.Errorf("Error: StatusCode:%d Code:%d Message: %s MoreInfo:%s",
			res.StatusCode,
			r.Errors[0].Code,
			r.Errors[0].Message,
			r.Errors[0].MoreInfo)
		return nil, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Delete - HTTP DELETE Request
func (cli *Client) Delete(api string, request url.Values) ([]byte, error) {
	uri := fmt.Sprintf("%s%s?apiKey=%s", cli.space, api, cli.apiKey)

	req, err := http.NewRequest(
		"DELETE",
		uri,
		strings.NewReader(request.Encode()))
	if err != nil {
		return nil, err
	}

	//DEBUG
	//fmt.Println(req.URL)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var r ErrorResponse
		if err = json.Unmarshal(body, &r); err != nil {
			return nil, err
		}
		err = fmt.Errorf("Error: StatusCode:%d Code:%d Message: %s MoreInfo:%s",
			res.StatusCode,
			r.Errors[0].Code,
			r.Errors[0].Message,
			r.Errors[0].MoreInfo)
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
