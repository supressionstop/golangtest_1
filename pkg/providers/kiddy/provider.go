package kiddy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"softpro6/internal/usecase"
)

type Kiddy struct {
	httpClient *http.Client

	baseUrl *url.URL
}

func NewKiddy(httpClient *http.Client, baseUrl string) (*Kiddy, error) {
	if baseUrl == "" {
		return nil, errors.New("base url is empty")
	}

	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid base url: %w", err)
	}

	return &Kiddy{
		httpClient: httpClient,
		baseUrl:    u,
	}, nil
}

func (p *Kiddy) FetchLine(ctx context.Context, sportName string) (Line, error) {
	path := "/v1/lines/"

	fullUrl := p.baseUrl.JoinPath(path, sportName)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullUrl.String(), nil)
	if err != nil {
		return Line{}, fmt.Errorf("providers - Kiddy - http.NewRequestWithContext: %w", err)
	}
	resp, err := p.httpClient.Do(req)
	if err != nil {
		return Line{}, fmt.Errorf("providers - Kiddy - httpClient.Do: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err) // todo remove
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return Line{}, fmt.Errorf("providers - Kiddy - resp.StatusCode: %w", err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return Line{}, fmt.Errorf("providers - Kiddy - io.ReadAll - body: %w", err)
	}

	var lines struct {
		Lines Line `json:"lines"`
	}
	err = json.Unmarshal(bodyBytes, &lines)
	if err != nil {
		return Line{}, fmt.Errorf("providers - Kiddy - json.Unmarshal - body: %w", err)
	}

	return lines.Lines, nil
}

func (p *Kiddy) GetLine(ctx context.Context, sportName string) (usecase.Line, error) {
	providerLine, err := p.FetchLine(ctx, sportName)
	if err != nil {
		return nil, err
	}

	return providerLine, err
}
