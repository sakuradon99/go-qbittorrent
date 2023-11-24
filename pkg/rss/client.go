package rss

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/KnutZuidema/go-qbittorrent/pkg"
	"github.com/KnutZuidema/go-qbittorrent/pkg/model"
)

type Client struct {
	BaseUrl string
	Client  *http.Client
	Logger  logrus.FieldLogger
}

func (c Client) AddFolder(folder string) error {
	params := url.Values{}
	params.Add("path", folder)
	if err := pkg.PostWithContentType(c.Client, c.BaseUrl+"/addFolder", strings.NewReader(params.Encode()), "application/x-www-form-urlencoded"); err != nil {
		return err
	}
	return nil
}

func (c Client) AddFeed(link string, folder string) error {
	params := url.Values{}
	params.Add("url", link)
	if folder != "" {
		params.Add("path", folder+"\\"+link)
	}
	if err := pkg.PostWithContentType(c.Client, c.BaseUrl+"/addFeed", strings.NewReader(params.Encode()), "application/x-www-form-urlencoded"); err != nil {
		return err
	}
	return nil
}

func (c Client) RemoveItem(folder string) error {
	params := url.Values{}
	params.Add("path", folder)
	if err := pkg.PostWithContentType(c.Client, c.BaseUrl+"/removeItem", strings.NewReader(params.Encode()), "application/x-www-form-urlencoded"); err != nil {
		return err
	}
	return nil
}

func (c Client) MoveItem(currentFolder, destinationFolder string) error {
	params := url.Values{}
	params.Add("itemPath", currentFolder)
	params.Add("destPath", destinationFolder)
	if err := pkg.PostWithContentType(c.Client, c.BaseUrl+"/moveItem", strings.NewReader(params.Encode()), "application/x-www-form-urlencoded"); err != nil {
		return err
	}
	return nil
}

func (c Client) SetRule(name string, def model.RuleDefinition) error {
	params := url.Values{}
	b, err := json.Marshal(def)
	if err != nil {
		return err
	}
	params.Add("ruleName", name)
	params.Add("ruleDef", string(b))
	if err := pkg.PostWithContentType(c.Client, c.BaseUrl+"/setRule", strings.NewReader(params.Encode()), "application/x-www-form-urlencoded"); err != nil {
		return err
	}
	return nil
}

func (c Client) RenameRule(old, new string) error {
	params := url.Values{}
	params.Add("ruleName", old)
	params.Add("newRuleName", new)
	if err := pkg.PostWithContentType(c.Client, c.BaseUrl+"/renameRule", strings.NewReader(params.Encode()), "application/x-www-form-urlencoded"); err != nil {
		return err
	}
	return nil
}

func (c Client) RemoveRule(name string) error {
	params := url.Values{}
	params.Add("ruleName", name)
	if err := pkg.PostWithContentType(c.Client, c.BaseUrl+"/removeRule", strings.NewReader(params.Encode()), "application/x-www-form-urlencoded"); err != nil {
		return err
	}
	return nil
}

func (c Client) GetRules() (map[string]model.RuleDefinition, error) {
	var res map[string]model.RuleDefinition
	if err := pkg.GetInto(c.Client, &res, c.BaseUrl+"/rules", nil); err != nil {
		return nil, err
	}
	return res, nil
}

func (c Client) GetAllItems() (map[string]interface{}, error) {
	var res map[string]interface{}
	if err := pkg.GetInto(c.Client, &res, c.BaseUrl+"/items", nil); err != nil {
		return nil, err
	}
	return res, nil
}
