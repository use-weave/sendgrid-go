package main

import (
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"os"
)

///////////////////////////////////////////////////
// Get a user's account information.
// GET /user/account

func Getauser'saccountinformation.() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/account", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve your credit balance
// GET /user/credits

func Retrieveyourcreditbalance() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/credits", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update your account email address
// PUT /user/email

func Updateyouraccountemailaddress() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/email", host, "v3")
  request.Method = "PUT"
  request.RequestBody = []byte(` {
  "email": "example@example.com"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve your account email address
// GET /user/email

func Retrieveyouraccountemailaddress() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/email", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update your password
// PUT /user/password

func Updateyourpassword() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/password", host, "v3")
  request.Method = "PUT"
  request.RequestBody = []byte(` {
  "new_password": "new_password", 
  "old_password": "old_password"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update a user's profile
// PATCH /user/profile

func Updateauser'sprofile() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/profile", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "city": "Orange", 
  "first_name": "Example", 
  "last_name": "User"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Get a user's profile
// GET /user/profile

func Getauser'sprofile() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/profile", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Cancel or pause a scheduled send
// POST /user/scheduled_sends

func Cancelorpauseascheduledsend() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/scheduled_sends", host, "v3")
  request.Method = "POST"
  request.RequestBody = []byte(` {
  "batch_id": "YOUR_BATCH_ID", 
  "status": "pause"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve all scheduled sends
// GET /user/scheduled_sends

func Retrieveallscheduledsends() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/scheduled_sends", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update user scheduled send information
// PATCH /user/scheduled_sends/{batch_id}

func Updateuserscheduledsendinformation() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/scheduled_sends/{batch_id}", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "status": "pause"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve scheduled send
// GET /user/scheduled_sends/{batch_id}

func Retrievescheduledsend() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/scheduled_sends/{batch_id}", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Delete a cancellation or pause of a scheduled send
// DELETE /user/scheduled_sends/{batch_id}

func Deleteacancellationorpauseofascheduledsend() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/scheduled_sends/{batch_id}", host, "v3")
  request.Method = "DELETE"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update Enforced TLS settings
// PATCH /user/settings/enforced_tls

func UpdateEnforcedTLSsettings() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/settings/enforced_tls", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "require_tls": true, 
  "require_valid_cert": false
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve current Enforced TLS settings.
// GET /user/settings/enforced_tls

func RetrievecurrentEnforcedTLSsettings.() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/settings/enforced_tls", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update your username
// PUT /user/username

func Updateyourusername() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/username", host, "v3")
  request.Method = "PUT"
  request.RequestBody = []byte(` {
  "username": "test_username"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve your username
// GET /user/username

func Retrieveyourusername() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/username", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Update Event Notification Settings
// PATCH /user/webhooks/event/settings

func UpdateEventNotificationSettings() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/webhooks/event/settings", host, "v3")
  request.Method = "PATCH"
  request.RequestBody = []byte(` {
  "bounce": true, 
  "click": true, 
  "deferred": true, 
  "delivered": true, 
  "dropped": true, 
  "enabled": true, 
  "group_resubscribe": true, 
  "group_unsubscribe": true, 
  "open": true, 
  "processed": true, 
  "spam_report": true, 
  "unsubscribe": true, 
  "url": "url"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve Event Webhook settings
// GET /user/webhooks/event/settings

func RetrieveEventWebhooksettings() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/webhooks/event/settings", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Test Event Notification Settings 
// POST /user/webhooks/event/test

func TestEventNotificationSettings() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/webhooks/event/test", host, "v3")
  request.Method = "POST"
  request.RequestBody = []byte(` {
  "url": "url"
}`)
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieve Parse Webhook settings
// GET /user/webhooks/parse/settings

func RetrieveParseWebhooksettings() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/webhooks/parse/settings", host, "v3")
  request.Method = "GET"
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

///////////////////////////////////////////////////
// Retrieves Inbound Parse Webhook statistics.
// GET /user/webhooks/parse/stats

func RetrievesInboundParseWebhookstatistics.() void {
  apiKey := "SENDGRID_APIKEY"
  host = "https://api.sendgrid.com"
  request := sendgrid.GetRequest(apiKey, "/user/webhooks/parse/stats", host, "v3")
  request.Method = "GET"
  queryParams := make(map[string]string)
  queryParams["aggregated_by"] = "day"
queryParams["limit"] = "test_string"
queryParams["start_date"] = "2016-01-01"
queryParams["end_date"] = "2016-04-01"
queryParams["offset"] = "test_string"
request.QueryParams = queryParams
  response, err := sendgrid.API(request)
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.ResponseBody)
    fmt.Println(response.ResponseHeaders)
  }
}

