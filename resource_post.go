package bunnynet

import (
	"context"
)

func resourcePostWithResponseString(
	ctx context.Context,
	client *Client,
	path string,
	requestBody any,
) (string, error) {
	var res string

	req, err := client.newPostRequest(path, requestBody)
	if err != nil {
		return "", err
	}

	if err := client.sendRequest(ctx, req, &res); err != nil {
		return "", err
	}

	return res, nil
}

func resourcePostWithResponse[Resp any](
	ctx context.Context,
	client *Client,
	path string,
	requestBody any,
) (*Resp, error) {
	var res Resp

	req, err := client.newPostRequest(path, requestBody)
	if err != nil {
		return nil, err
	}

	if err := client.sendRequest(ctx, req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func resourcePost(
	ctx context.Context,
	client *Client,
	path string,
	requestBody any,
) error {
	req, err := client.newPostRequest(path, requestBody)
	if err != nil {
		return err
	}

	return client.sendRequest(ctx, req, nil)
}
