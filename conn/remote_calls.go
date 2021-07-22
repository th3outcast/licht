package conn

import (
  "fmt"
  "net/http"
  "io"
  "io/ioutil"
  "errors"
)

/* Remote access(get, put, delete) functions */

func remote_get(remote string) (string, error) {
  resp, err := http.Get(remote)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()
  if resp.StatusCode != 200 {
    return "", errors.New(fmt.Sprintf("remote_get: wrong status code %d\n", resp.StatusCode))
  }
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return "", err
  }
  return string(body), nil
}

func remote_put(remote string, length int64, body io.Reader) error {
  req, err := http.NewRequest("PUT", remote, body)
  if err != nil {
    return nil
  }
  req.ContentLength = length
  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return err
  }
  defer resp.Body.Close()
  if resp.StatusCode != 201 && resp.StatusCode != 204 {
    return fmt.Errorf("remote_put: wrong status code %d\n", resp.StatusCode)
  }
  return nil
}

func remote_delete(remote string) error {
  req, err := http.NewRequest("DELETE", remote, nil)
  if err != nil {
    return err
  }
  resp, err := http.DefaultClient.Do(req)
  if err != nil {
    return err
  }
  defer resp.Body.Close()
  if resp.StatusCode != 204 && resp.StatusCode != 404 {
    return fmt.Errorf("remote_delete: wrong status code %d\n", resp.StatusCode)
  }
  return nil
}
