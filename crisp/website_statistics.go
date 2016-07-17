// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteStatisticsAllData mapping
type WebsiteStatisticsAllData struct {
  Data  *WebsiteStatisticsAll  `json:"data,omitempty"`
}

// WebsiteStatisticsAll mapping
type WebsiteStatisticsAll struct {
  Total    *int  `json:"total,omitempty"`
  Resolve  *int  `json:"resolve,omitempty"`
  Pending  *int  `json:"pending,omitempty"`
  Unread   *int  `json:"unread,omitempty"`
}

// WebsiteStatisticsTotalData mapping
type WebsiteStatisticsTotalData struct {
  Data  *WebsiteStatisticsTotal  `json:"data,omitempty"`
}

// WebsiteStatisticsTotal mapping
type WebsiteStatisticsTotal struct {
  Total    *int  `json:"total,omitempty"`
  Resolve  *int  `json:"resolve,omitempty"`
  Pending  *int  `json:"pending,omitempty"`
  Unread   *int  `json:"unread,omitempty"`
}

// WebsiteStatisticsPendingData mapping
type WebsiteStatisticsPendingData struct {
  Data  *WebsiteStatisticsPending  `json:"data,omitempty"`
}

// WebsiteStatisticsPending mapping
type WebsiteStatisticsPending struct {
  Pending  *int  `json:"pending,omitempty"`
}

// WebsiteStatisticsUnresolvedData mapping
type WebsiteStatisticsUnresolvedData struct {
  Data  *WebsiteStatisticsUnresolved  `json:"data,omitempty"`
}

// WebsiteStatisticsUnresolved mapping
type WebsiteStatisticsUnresolved struct {
  Unresolved  *int  `json:"unresolved,omitempty"`
}

// WebsiteStatisticsResolvedData mapping
type WebsiteStatisticsResolvedData struct {
  Data  *WebsiteStatisticsResolved  `json:"data,omitempty"`
}

// WebsiteStatisticsResolved mapping
type WebsiteStatisticsResolved struct {
  Resolved  *int  `json:"resolved,omitempty"`
}

// WebsiteStatisticsUnreadData mapping
type WebsiteStatisticsUnreadData struct {
  Data  *WebsiteStatisticsUnread  `json:"data,omitempty"`
}

// WebsiteStatisticsUnread mapping
type WebsiteStatisticsUnread struct {
  Unread  *int  `json:"unread,omitempty"`
}


// GetAllStatistics get all statistics for website.
// Reference: https://docs.crisp.im/api/v1/#website-website-statistics-get
func (service *WebsiteService) GetAllStatistics(websiteID string) (*WebsiteStatisticsAll, *Response, error) {
  url := fmt.Sprintf("website/%s/stats", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  statistics := new(WebsiteStatisticsAllData)
  resp, err := service.client.Do(req, statistics)
  if err != nil {
    return nil, resp, err
  }

  return statistics.Data, resp, err
}

// CountTotalNumberOfConversations get total number of conversations for website.
// Reference: https://docs.crisp.im/api/v1/#website-website-statistics-get-1
func (service *WebsiteService) CountTotalNumberOfConversations(websiteID string) (*WebsiteStatisticsTotal, *Response, error) {
  url := fmt.Sprintf("website/%s/stats/total", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  statistics := new(WebsiteStatisticsTotalData)
  resp, err := service.client.Do(req, statistics)
  if err != nil {
    return nil, resp, err
  }

  return statistics.Data, resp, err
}

// CountNumberOfPendingConversations get number of pending conversations for website.
// Reference: https://docs.crisp.im/api/v1/#website-website-statistics-get-2
func (service *WebsiteService) CountNumberOfPendingConversations(websiteID string) (*WebsiteStatisticsPending, *Response, error) {
  url := fmt.Sprintf("website/%s/stats/pending", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  statistics := new(WebsiteStatisticsPendingData)
  resp, err := service.client.Do(req, statistics)
  if err != nil {
    return nil, resp, err
  }

  return statistics.Data, resp, err
}

// CountNumberOfUnresolvedConversations get number of unresolved conversations for website.
// Reference: https://docs.crisp.im/api/v1/#website-website-statistics-get-3
func (service *WebsiteService) CountNumberOfUnresolvedConversations(websiteID string) (*WebsiteStatisticsUnresolved, *Response, error) {
  url := fmt.Sprintf("website/%s/stats/unresolved", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  statistics := new(WebsiteStatisticsUnresolvedData)
  resp, err := service.client.Do(req, statistics)
  if err != nil {
    return nil, resp, err
  }

  return statistics.Data, resp, err
}

// CountNumberOfResolvedConversations get number of resolved conversations for website.
// Reference: https://docs.crisp.im/api/v1/#website-website-statistics-get-4
func (service *WebsiteService) CountNumberOfResolvedConversations(websiteID string) (*WebsiteStatisticsResolved, *Response, error) {
  url := fmt.Sprintf("website/%s/stats/resolved", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  statistics := new(WebsiteStatisticsResolvedData)
  resp, err := service.client.Do(req, statistics)
  if err != nil {
    return nil, resp, err
  }

  return statistics.Data, resp, err
}

// CountNumberOfUnreadMessages get number of unread messages for website.
// Reference: https://docs.crisp.im/api/v1/#website-website-statistics-get-5
func (service *WebsiteService) CountNumberOfUnreadMessages(websiteID string) (*WebsiteStatisticsUnread, *Response, error) {
  url := fmt.Sprintf("website/%s/stats/unread", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  statistics := new(WebsiteStatisticsUnreadData)
  resp, err := service.client.Do(req, statistics)
  if err != nil {
    return nil, resp, err
  }

  return statistics.Data, resp, err
}
