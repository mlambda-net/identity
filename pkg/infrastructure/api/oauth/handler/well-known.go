package handler

import (
  "encoding/json"
  "net/http"
)

type well struct {
  Issuer                 string   `json:"issuer"`
  AuthorizationEndpoint  string   `json:"authorization_endpoint"`
  TokenEndpoint          string   `json:"token_endpoint"`
  ScopesSupported        []string `json:"scopes_supported"`
  ResponseTypesSupported []string `json:"response_types_supported"`
  GranTypesSupported [] string `json:"gran_types_supported"`
}

func (h *handler) WellKnow(w http.ResponseWriter, r *http.Request) {
  setupHeaders(w,r)
  w.Header().Set("content-type","application/json")
  //https://docs.akana.com/cm/api_oauth/oauth_discovery/m_oauth_getOpenIdConnectWellknownConfiguration.htm
  _ = json.NewEncoder(w).Encode( &well{
    Issuer: "https://oauth.mitienda.co.cr",
    AuthorizationEndpoint: "https://oauth.mitienda.co.cr/authorize",
    TokenEndpoint: "https://oauth.mitienda.co.cr/token",
    ScopesSupported: []string{"all"},
    ResponseTypesSupported: []string{"code", "token" },
    GranTypesSupported: []string{"authorization_code", "implicit","password", "client_credentials"},
  })
}
