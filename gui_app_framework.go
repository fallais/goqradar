package goqradar

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

//------------------------------------------------------------------------------
// Structures
//------------------------------------------------------------------------------

// StatusAppInstall is the status of the application installs.
type StatusAppInstall struct {
	ApplicationID int    `json:"application_id"`
	Status        string `json:"status"`
	ErrorMessages string `json:"error_messages,omitempty"`
}

// StatusAppInstallsPaginatedResponse is the paginated response.
type StatusAppInstallsPaginatedResponse struct {
	Total             int                 `json:"total"`
	Min               int                 `json:"min"`
	Max               int                 `json:"max"`
	StatusAppInstalls []*StatusAppInstall `json:"status_app_installs"`
}

// CreatedAppFramework is a QRadar Created application framework.
type CreatedAppFramework struct {
	ApplicationID     int    `json:"application_id"`
	ErrorMessages     string `json:"error_messages"`
	ErrorMessagesJSON []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Source  string `json:"source"`
	} `json:"error_messages_json"`
	Status string `json:"status"`
}

// AuthRequest is an authorisation request for an application install.
type AuthRequest struct {
	Capabilities []string `json:"capabilities"`
}

// AuthRequestResponse is the response of an authorisation request for an application install.
type AuthRequestResponse struct {
	Capabilities []string `json:"capabilities"`
	UserID       int      `json:"user_id"`
}

// AppDefinition is a QRadar application definition
type AppDefinition struct {
	ApplicationDefinitionID int    `json:"application_definition_id"`
	CreatedBy               string `json:"created_by"`
	CreatedOn               int    `json:"created_on"`
	ErrorMessages           string `json:"error_messages"`
	ErrorMessagesJSON       []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Source  string `json:"source"`
	} `json:"error_messages_json"`
	Manifest struct {
		AppID int `json:"app_id"`
		Areas []struct {
			Description          string   `json:"description"`
			ID                   string   `json:"id"`
			NamedService         string   `json:"named_service"`
			RequiredCapabilities []string `json:"required_capabilities"`
			Text                 string   `json:"text"`
			URL                  string   `json:"url"`
		} `json:"areas"`
		Authentication struct {
			Oauth2 struct {
				AuthorizationFlow     string   `json:"authorization_flow"`
				RequestedCapabilities []string `json:"requested_capabilities"`
			} `json:"oauth2"`
		} `json:"authentication"`
		ConfigurationPages []struct {
			Description          string   `json:"description"`
			Icon                 string   `json:"icon"`
			NamedService         string   `json:"named_service"`
			RequiredCapabilities []string `json:"required_capabilities"`
			Text                 string   `json:"text"`
			URL                  string   `json:"url"`
		} `json:"configuration_pages"`
		ConsoleIP     string `json:"console_ip"`
		CustomColumns []struct {
			Label                string   `json:"label"`
			NamedService         string   `json:"named_service"`
			PageID               string   `json:"page_id"`
			RequiredCapabilities []string `json:"required_capabilities"`
			RestEndpoint         string   `json:"rest_endpoint"`
		} `json:"custom_columns"`
		DashboardItems []struct {
			Description          string   `json:"description"`
			RequiredCapabilities []string `json:"required_capabilities"`
			RestMethod           string   `json:"rest_method"`
			Text                 string   `json:"text"`
		} `json:"dashboard_items"`
		Dependencies struct {
			PipDirectory  string `json:"pip_directory"`
			RpmsDirectory string `json:"rpms_directory"`
		} `json:"dependencies"`
		Description          string `json:"description"`
		EnvironmentVariables []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"environment_variables"`
		Fragments []struct {
			AppName              string   `json:"app_name"`
			Location             string   `json:"location"`
			NamedService         string   `json:"named_service"`
			PageID               string   `json:"page_id"`
			RequiredCapabilities []string `json:"required_capabilities"`
			RestEndpoint         string   `json:"rest_endpoint"`
		} `json:"fragments"`
		GuiActions []struct {
			Description          string   `json:"description"`
			Groups               []string `json:"groups"`
			Icon                 string   `json:"icon"`
			ID                   string   `json:"id"`
			Javascript           string   `json:"javascript"`
			NamedService         string   `json:"named_service"`
			RequiredCapabilities []string `json:"required_capabilities"`
			RestMethod           string   `json:"rest_method"`
			Text                 string   `json:"text"`
		} `json:"gui_actions"`
		LoadFlask         string `json:"load_flask"`
		LogLevel          string `json:"log_level"`
		MetadataProviders []struct {
			MetadataType string `json:"metadata_type"`
			RestMethod   string `json:"rest_method"`
		} `json:"metadata_providers"`
		MultitenancySafe string `json:"multitenancy_safe"`
		Name             string `json:"name"`
		PageScripts      []struct {
			AppName      string   `json:"app_name"`
			NamedService string   `json:"named_service"`
			PageID       string   `json:"page_id"`
			Scripts      []string `json:"scripts"`
		} `json:"page_scripts"`
		ResourceBundles []struct {
			Bundle string `json:"bundle"`
			Locale string `json:"locale"`
		} `json:"resource_bundles"`
		Resources struct {
			Memory int `json:"memory"`
		} `json:"resources"`
		RestMethods []struct {
			ArgumentNames        []string `json:"argument_names"`
			Method               string   `json:"method"`
			Name                 string   `json:"name"`
			NamedService         string   `json:"named_service"`
			RequiredCapabilities []string `json:"required_capabilities"`
			URL                  string   `json:"url"`
		} `json:"rest_methods"`
		Services []struct {
			Autorestart string `json:"autorestart"`
			Autostart   string `json:"autostart"`
			Command     string `json:"command"`
			Directory   string `json:"directory"`
			Endpoints   []struct {
				ErrorMimeType string `json:"error_mime_type"`
				HTTPMethod    string `json:"http_method"`
				Name          string `json:"name"`
				Parameters    []struct {
					Definition string `json:"definition"`
					Location   string `json:"location"`
					Name       string `json:"name"`
				} `json:"parameters"`
				Path            string `json:"path"`
				RequestMimeType string `json:"request_mime_type"`
				Response        struct {
					MimeType string `json:"mime_type"`
				} `json:"response"`
			} `json:"endpoints"`
			Environment           string `json:"environment"`
			Exitcodes             string `json:"exitcodes"`
			Name                  string `json:"name"`
			Numprocs              int    `json:"numprocs"`
			Port                  int    `json:"port"`
			Priority              int    `json:"priority"`
			ProcessName           string `json:"process_name"`
			RedirectStderr        string `json:"redirect_stderr"`
			Serverurl             string `json:"serverurl"`
			Startretries          int    `json:"startretries"`
			Startsecs             int    `json:"startsecs"`
			StderrCaptureMaxbytes string `json:"stderr_capture_maxbytes"`
			StderrEventsEnabled   string `json:"stderr_events_enabled"`
			StderrLogfile         string `json:"stderr_logfile"`
			StderrLogfileBackups  int    `json:"stderr_logfile_backups"`
			StderrLogfileMaxbytes string `json:"stderr_logfile_maxbytes"`
			StdoutCaptureMaxbytes string `json:"stdout_capture_maxbytes"`
			StdoutEventsEnabled   string `json:"stdout_events_enabled"`
			StdoutLogfile         string `json:"stdout_logfile"`
			StdoutLogfileBackups  int    `json:"stdout_logfile_backups"`
			StdoutLogfileMaxbyte  string `json:"stdout_logfile_maxbyte"`
			Stopsignal            string `json:"stopsignal"`
			Stopwaitsecs          int    `json:"stopwaitsecs"`
			Umask                 string `json:"umask"`
			User                  string `json:"user"`
			UUID                  string `json:"uuid"`
			Version               string `json:"version"`
		} `json:"services"`
		SingleInstanceOnly string `json:"single_instance_only"`
		UUID               string `json:"uuid"`
		Version            string `json:"version"`
	} `json:"manifest"`
	Status      string `json:"status"`
	UserRoleIds []int  `json:"user_role_ids"`
}

// AppDefinitionsPaginatedResponse is the paginated response.
type AppDefinitionsPaginatedResponse struct {
	Total          int              `json:"total"`
	Min            int              `json:"min"`
	Max            int              `json:"max"`
	AppDefinitions []*AppDefinition `json:"status_app_installs"`
}

// AppDefinitionStatus is a QRadar application definition status.
type AppDefinitionStatus struct {
	ApplicationDefinitionID int    `json:"application_definition_id"`
	ErrorMessages           string `json:"error_messages"`
	Status                  string `json:"status"`
}

// UserRoleID is a Qradar user role id
type UserRoleID struct {
	UserRoles []int `json:"user_roles"`
}

// UserRoleIDsPaginatedResponse is the paginated response.
type UserRoleIDsPaginatedResponse struct {
	Total       int           `json:"total"`
	Min         int           `json:"min"`
	Max         int           `json:"max"`
	UserRoleIDs []*UserRoleID `json:"status_app_installs"`
}

// InstalledApp is a Qradar install application.
type InstalledApp struct {
	ApplicationDefinitionID int `json:"application_definition_id"`
	ApplicationState        struct {
		ApplicationID     string `json:"application_id"`
		ErrorMessages     string `json:"error_messages"`
		ErrorMessagesJSON []struct {
			Code    string `json:"code"`
			Message string `json:"message"`
			Source  string `json:"source"`
		} `json:"error_messages_json"`
		Memory int    `json:"memory"`
		Status string `json:"status"`
	} `json:"application_state"`
	AuthClientUserID int    `json:"auth_client_user_id"`
	InstalledBy      string `json:"installed_by"`
	InstalledOn      int    `json:"installed_on"`
	ManagedHostID    int    `json:"managed_host_id"`
	Manifest         struct {
		AppID int `json:"app_id"`
		Areas []struct {
			Description          string   `json:"description"`
			ID                   string   `json:"id"`
			NamedService         string   `json:"named_service"`
			RequiredCapabilities []string `json:"required_capabilities"`
			Text                 string   `json:"text"`
			URL                  string   `json:"url"`
		} `json:"areas"`
		Authentication struct {
			Oauth2 struct {
				AuthorizationFlow     string   `json:"authorization_flow"`
				RequestedCapabilities []string `json:"requested_capabilities"`
			} `json:"oauth2"`
		} `json:"authentication"`
		ConfigurationPages []struct {
			Description          string   `json:"description"`
			Icon                 string   `json:"icon"`
			NamedService         string   `json:"named_service"`
			RequiredCapabilities []string `json:"required_capabilities"`
			Text                 string   `json:"text"`
			URL                  string   `json:"url"`
		} `json:"configuration_pages"`
		ConsoleIP     string `json:"console_ip"`
		CustomColumns []struct {
			Label                string   `json:"label"`
			NamedService         string   `json:"named_service"`
			PageID               string   `json:"page_id"`
			RequiredCapabilities []string `json:"required_capabilities"`
			RestEndpoint         string   `json:"rest_endpoint"`
		} `json:"custom_columns"`
		DashboardItems []struct {
			Description          string   `json:"description"`
			RequiredCapabilities []string `json:"required_capabilities"`
			RestMethod           string   `json:"rest_method"`
			Text                 string   `json:"text"`
		} `json:"dashboard_items"`
		Dependencies struct {
			PipDirectory  string `json:"pip_directory"`
			RpmsDirectory string `json:"rpms_directory"`
		} `json:"dependencies"`
		Description          string `json:"description"`
		EnvironmentVariables []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"environment_variables"`
		Fragments []struct {
			AppName              string   `json:"app_name"`
			Location             string   `json:"location"`
			NamedService         string   `json:"named_service"`
			PageID               string   `json:"page_id"`
			RequiredCapabilities []string `json:"required_capabilities"`
			RestEndpoint         string   `json:"rest_endpoint"`
		} `json:"fragments"`
		GuiActions []struct {
			Description          string   `json:"description"`
			Groups               []string `json:"groups"`
			Icon                 string   `json:"icon"`
			ID                   string   `json:"id"`
			Javascript           string   `json:"javascript"`
			NamedService         string   `json:"named_service"`
			RequiredCapabilities []string `json:"required_capabilities"`
			RestMethod           string   `json:"rest_method"`
			Text                 string   `json:"text"`
		} `json:"gui_actions"`
		LoadFlask         string `json:"load_flask"`
		LogLevel          string `json:"log_level"`
		MetadataProviders []struct {
			MetadataType string `json:"metadata_type"`
			RestMethod   string `json:"rest_method"`
		} `json:"metadata_providers"`
		MultitenancySafe string `json:"multitenancy_safe"`
		Name             string `json:"name"`
		PageScripts      []struct {
			AppName      string   `json:"app_name"`
			NamedService string   `json:"named_service"`
			PageID       string   `json:"page_id"`
			Scripts      []string `json:"scripts"`
		} `json:"page_scripts"`
		ResourceBundles []struct {
			Bundle string `json:"bundle"`
			Locale string `json:"locale"`
		} `json:"resource_bundles"`
		Resources struct {
			Memory int `json:"memory"`
		} `json:"resources"`
		RestMethods []struct {
			ArgumentNames        []string `json:"argument_names"`
			Method               string   `json:"method"`
			Name                 string   `json:"name"`
			NamedService         string   `json:"named_service"`
			RequiredCapabilities []string `json:"required_capabilities"`
			URL                  string   `json:"url"`
		} `json:"rest_methods"`
		Services []struct {
			Autorestart string `json:"autorestart"`
			Autostart   string `json:"autostart"`
			Command     string `json:"command"`
			Directory   string `json:"directory"`
			Endpoints   []struct {
				ErrorMimeType string `json:"error_mime_type"`
				HTTPMethod    string `json:"http_method"`
				Name          string `json:"name"`
				Parameters    []struct {
					Definition string `json:"definition"`
					Location   string `json:"location"`
					Name       string `json:"name"`
				} `json:"parameters"`
				Path            string `json:"path"`
				RequestMimeType string `json:"request_mime_type"`
				Response        struct {
					MimeType string `json:"mime_type"`
				} `json:"response"`
			} `json:"endpoints"`
			Environment           string `json:"environment"`
			Exitcodes             string `json:"exitcodes"`
			Name                  string `json:"name"`
			Numprocs              int    `json:"numprocs"`
			Port                  int    `json:"port"`
			Priority              int    `json:"priority"`
			ProcessName           string `json:"process_name"`
			RedirectStderr        string `json:"redirect_stderr"`
			Serverurl             string `json:"serverurl"`
			Startretries          int    `json:"startretries"`
			Startsecs             int    `json:"startsecs"`
			StderrCaptureMaxbytes string `json:"stderr_capture_maxbytes"`
			StderrEventsEnabled   string `json:"stderr_events_enabled"`
			StderrLogfile         string `json:"stderr_logfile"`
			StderrLogfileBackups  int    `json:"stderr_logfile_backups"`
			StderrLogfileMaxbytes string `json:"stderr_logfile_maxbytes"`
			StdoutCaptureMaxbytes string `json:"stdout_capture_maxbytes"`
			StdoutEventsEnabled   string `json:"stdout_events_enabled"`
			StdoutLogfile         string `json:"stdout_logfile"`
			StdoutLogfileBackups  int    `json:"stdout_logfile_backups"`
			StdoutLogfileMaxbyte  string `json:"stdout_logfile_maxbyte"`
			Stopsignal            string `json:"stopsignal"`
			Stopwaitsecs          int    `json:"stopwaitsecs"`
			Umask                 string `json:"umask"`
			User                  string `json:"user"`
			UUID                  string `json:"uuid"`
			Version               string `json:"version"`
		} `json:"services"`
		SingleInstanceOnly string `json:"single_instance_only"`
		UUID               string `json:"uuid"`
		Version            string `json:"version"`
	} `json:"manifest"`
	SecurityProfileID int `json:"security_profile_id"`
}

// InstalledAppsPaginatedResponse is the paginated response.
type InstalledAppsPaginatedResponse struct {
	Total         int             `json:"total"`
	Min           int             `json:"min"`
	Max           int             `json:"max"`
	InstalledApps []*InstalledApp `json:"status_app_installs"`
}

// RegisteredService is a QRadar named service registered with the application framework.
type RegisteredService struct {
	Name          string `json:"name"`
	Version       string `json:"version"`
	ApplicationID int    `json:"application_id"`
	UUID          string `json:"uuid"`
	Endpoints     []struct {
		Name       string `json:"name"`
		Path       string `json:"path"`
		HTTPMethod string `json:"http_method"`
		Parameters []struct {
			Location string `json:"location"`
			Name     string `json:"name"`
		} `json:"parameters,omitempty"`
		Response struct {
			MimeType string `json:"mime_type"`
			BodyType struct {
				Type          string `json:"@type"`
				ResourceID    string `json:"resource_id"`
				ResourceName  string `json:"resource_name"`
				ResourceOwner string `json:"resource_owner"`
			} `json:"body_type"`
		} `json:"response"`
		ErrorMimeType   string `json:"error_mime_type"`
		RequestMimeType string `json:"request_mime_type,omitempty"`
		RequestBodyType struct {
			Type          string `json:"@type"`
			ResourceName  string `json:"resource_name"`
			ResourceOwner string `json:"resource_owner"`
		} `json:"request_body_type,omitempty"`
	} `json:"endpoints"`
}

// RegisteredServicesPaginatedResponse is the paginated response.
type RegisteredServicesPaginatedResponse struct {
	Total              int                  `json:"total"`
	Min                int                  `json:"min"`
	Max                int                  `json:"max"`
	RegisteredServices []*RegisteredService `json:"status_app_installs"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListStatusAppInstalls returns the app installs with given fields, filters.
func (endpoint *Endpoint) ListStatusAppInstalls(ctx context.Context, fields, filter string, min, max int) (*StatusAppInstallsPaginatedResponse, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/gui_app_framework/application_creation_task", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Process the Content-Range
	min, max, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing the content-range [%s]: %s", resp.Header.Get("Content-Range"), err)
	}

	// Prepare the response
	response := &StatusAppInstallsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.StatusAppInstalls)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// CreateAppFramework creates a new  application within the application framework. ZIP FILE UPLOAD
func (endpoint *Endpoint) CreateAppFramework(ctx context.Context, filename, fields string) (*CreatedAppFramework, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/application_creation_task"

	// handle zip file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error while opening the file : %s", err)
	}
	defer file.Close()

	buff := new(bytes.Buffer)
	writer := multipart.NewWriter(buff)
	part, err := writer.CreateFormFile(filename, filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), buff)
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/zip")
	if fields != "" {
		req.Header.Set("fields", fields)
	}

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error while reading the request : %s", err)
	}

	// Prepare the response
	var response *CreatedAppFramework

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// GetCreatedAppFramework retrieves an app status by id.
func (endpoint *Endpoint) GetCreatedAppFramework(ctx context.Context, id int, fields string) (*CreatedAppFramework, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/gui_app_framework/application_creation_task/"+strconv.Itoa(id), options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *CreatedAppFramework

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// CancelCreatedAppFramework by id
func (endpoint *Endpoint) CancelCreatedAppFramework(ctx context.Context, id int, fields string) (*CreatedAppFramework, error) {
	status := "CANCELLED"

	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/application_creation_task/"
	reqURL.Path += strconv.Itoa(id)

	// Create the data
	d, err := json.Marshal(status)
	if err != nil {
		return nil, fmt.Errorf("Error while marshalling the values : %s", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), bytes.NewBuffer(d))
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Add optional parameters
	q := req.URL.Query()
	q.Add("fields", fields)
	req.URL.RawQuery = q.Encode()

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *CreatedAppFramework

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// GetAuthRequest retrieves an authorisation request by id.
func (endpoint *Endpoint) GetAuthRequest(ctx context.Context, id int, fields string) (*AuthRequest, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/gui_app_framework/application_creation_task/"+strconv.Itoa(id)+"/auth", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *AuthRequest

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// UpdateAuthRequestResponse by id
func (endpoint *Endpoint) UpdateAuthRequestResponse(ctx context.Context, id int, data map[string]string, fields string) (*AuthRequestResponse, error) {

	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/application_creation_task/" + strconv.Itoa(id) + "/auth"

	// Create the data
	d, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Error while marshalling the values : %s", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), bytes.NewBuffer(d))
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")
	if fields != "" {
		req.Header.Set("fields", fields)
	}

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *AuthRequestResponse

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// ListAppDefinitions returns the application definitions with given fields, filters.
func (endpoint *Endpoint) ListAppDefinitions(ctx context.Context, fields, filter string, min, max int) (*AppDefinitionsPaginatedResponse, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/gui_app_framework/application_definitions", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Process the Content-Range
	min, max, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing the content-range [%s]: %s", resp.Header.Get("Content-Range"), err)
	}

	// Prepare the response
	response := &AppDefinitionsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.AppDefinitions)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// CreateAppDefinition initialises the asynchronous installation of a new application within the application framework. ZIP FILE UPLOAD
func (endpoint *Endpoint) CreateAppDefinition(ctx context.Context, filename, fields string) (*AppDefinitionStatus, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/application_definitions"

	// handle zip file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error while opening the file : %s", err)
	}
	defer file.Close()

	buff := new(bytes.Buffer)
	writer := multipart.NewWriter(buff)
	part, err := writer.CreateFormFile(filename, filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), buff)
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/zip")
	if fields != "" {
		req.Header.Set("fields", fields)
	}

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *AppDefinitionStatus

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// GetAppDefinition retrieves an app definition by id.
func (endpoint *Endpoint) GetAppDefinition(ctx context.Context, id int, fields string) (*AppDefinition, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/gui_app_framework/application_definitions/"+strconv.Itoa(id), options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *AppDefinition

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// CancelAppDefinition by id
func (endpoint *Endpoint) CancelAppDefinition(ctx context.Context, id int, fields string) (*AppDefinitionStatus, error) {
	status := "CANCELLED"

	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/application_definitions/"
	reqURL.Path += strconv.Itoa(id)

	// Create the data
	d, err := json.Marshal(status)
	if err != nil {
		return nil, fmt.Errorf("Error while marshalling the values : %s", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), bytes.NewBuffer(d))
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Add optional parameters
	q := req.URL.Query()
	q.Add("fields", fields)
	req.URL.RawQuery = q.Encode()

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")
	if fields != "" {
		req.Header.Set("fields", fields)
	}

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *AppDefinitionStatus

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// DeleteAppDefinition by ID
func (endpoint *Endpoint) DeleteAppDefinition(ctx context.Context, id int) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/application_definitions/"
	reqURL.Path += strconv.Itoa(id)

	// Create the request
	req, err := http.NewRequest("DELETE", reqURL.String(), nil)
	if err != nil {
		return fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while doing the request : %s", err)
	}

	// Check the status code
	if resp.StatusCode != 204 {
		return fmt.Errorf("Status code is %d : Error while reading the body", resp.StatusCode)
	}

	return nil

}

// UpdateAppDefinition by id ZIP FILE UPLOAD
func (endpoint *Endpoint) UpdateAppDefinition(ctx context.Context, id int, filename, fields string) (*AppDefinitionStatus, error) {

	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/application_definitions/" + strconv.Itoa(id)

	// handle zip file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error while opening the file : %s", err)
	}
	defer file.Close()

	buff := new(bytes.Buffer)
	writer := multipart.NewWriter(buff)
	part, err := writer.CreateFormFile(filename, filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	// Create the request
	req, err := http.NewRequest("PUT", reqURL.String(), buff)
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/zip")
	if fields != "" {
		req.Header.Set("fields", fields)
	}

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *AppDefinitionStatus

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// ListUserRoleIds returns the user role ID assiciated with an application definition with given fields, filters.
func (endpoint *Endpoint) ListUserRoleIds(ctx context.Context, id int, fields string, min, max int) (*UserRoleIDsPaginatedResponse, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/gui_app_framework/application_definitions/"+strconv.Itoa(id)+"/user_role_id", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Process the Content-Range
	min, max, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing the content-range [%s]: %s", resp.Header.Get("Content-Range"), err)
	}

	// Prepare the response
	response := &UserRoleIDsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.UserRoleIDs)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// CreateUserRoleID add a user role to the list associated with an application definition.
func (endpoint *Endpoint) CreateUserRoleID(ctx context.Context, AppID, userRoleID int, fields string) (*UserRoleID, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/application_definitions/" + strconv.Itoa(AppID) + "/user_role_id/" + strconv.Itoa(userRoleID)

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Add optional parameters
	q := req.URL.Query()
	q.Add("fields", fields)
	req.URL.RawQuery = q.Encode()

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *UserRoleID

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// DeleteUserRoles  by ID
func (endpoint *Endpoint) DeleteUserRoles(ctx context.Context, AppID, userRoleID int, fields string) (*UserRoleID, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/application_definitions/" + strconv.Itoa(AppID) + "/user_role_id/" + strconv.Itoa(userRoleID)

	// Create the request
	req, err := http.NewRequest("DELETE", reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Add optional parameters
	q := req.URL.Query()
	q.Add("fields", fields)
	req.URL.RawQuery = q.Encode()

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}

	// Check the status code
	if resp.StatusCode != 404 {
		return nil, fmt.Errorf("Status code is %d : Error while reading the body", resp.StatusCode)
	}

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *UserRoleID

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil

}

// ListInstalledApp returns a list of all installed applications with given fields, filters.
func (endpoint *Endpoint) ListInstalledApp(ctx context.Context, id int, fields, filter string, min, max int) (*InstalledAppsPaginatedResponse, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/gui_app_framework/applications", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Process the Content-Range
	min, max, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing the content-range [%s]: %s", resp.Header.Get("Content-Range"), err)
	}

	// Prepare the response
	response := &InstalledAppsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.InstalledApps)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// CreateApplication initiates the creation of a new application instance within the Application framework.
func (endpoint *Endpoint) CreateApplication(ctx context.Context, appID, securityProfile int, fields string, forceMultitenancySafe bool) (*InstalledApp, error) {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/applications"

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Add optional parameters
	q := req.URL.Query()
	q.Add("application_definition_id", strconv.Itoa(appID))
	q.Add("fields", fields)
	q.Add("force_multitenancy_safe", strconv.FormatBool(forceMultitenancySafe))
	q.Add("security_profile_id", strconv.Itoa(securityProfile))
	req.URL.RawQuery = q.Encode()

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *InstalledApp

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// GetinstalledApp retrieves an installed application by id.
func (endpoint *Endpoint) GetinstalledApp(ctx context.Context, id int, fields string) (*InstalledApp, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/gui_app_framework/applications/"+strconv.Itoa(id), options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *InstalledApp

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// UpdateInstalledApp by id.
func (endpoint *Endpoint) UpdateInstalledApp(ctx context.Context, appID, oauthUserID int, fields, securityProfileID, status string) (*InstalledApp, error) {

	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/applications/" + strconv.Itoa(appID)

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Add optional parameters
	q := req.URL.Query()
	q.Add("fields", fields)
	q.Add("oauth_user_id", strconv.Itoa(oauthUserID))
	q.Add("security_profile_id", securityProfileID)
	q.Add("status", status)
	req.URL.RawQuery = q.Encode()

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *InstalledApp

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// DeleteAppInstance  by ID
func (endpoint *Endpoint) DeleteAppInstance(ctx context.Context, AppID int) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/applications/" + strconv.Itoa(AppID)

	// Create the request
	req, err := http.NewRequest("DELETE", reqURL.String(), nil)
	if err != nil {
		return fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while doing the request : %s", err)
	}

	// Check the status code
	if resp.StatusCode != 204 {
		return fmt.Errorf("Status code is %d : Error while reading the body", resp.StatusCode)
	}

	return nil

}

// UpdateApplication by id ZIP FILE UPLOAD
func (endpoint *Endpoint) UpdateApplication(ctx context.Context, appID int, filename, fields string) (*CreatedAppFramework, error) {

	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/gui_app_framework/applications/" + strconv.Itoa(appID)

	// handle zip file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error while opening the file : %s", err)
	}
	defer file.Close()

	buff := new(bytes.Buffer)
	writer := multipart.NewWriter(buff)
	part, err := writer.CreateFormFile(filename, filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	// Create the request
	req, err := http.NewRequest("PUT", reqURL.String(), buff)
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/zip")
	if fields != "" {
		req.Header.Set("fields", fields)
	}

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *CreatedAppFramework

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// ListRegisteredServices retrieves a list of all named services registered with the Application Framework. with given fields, filters.
func (endpoint *Endpoint) ListRegisteredServices(ctx context.Context, min, max int) (*RegisteredServicesPaginatedResponse, error) {
	// Options
	options := []Option{}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/gui_app_framework/named_services", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Process the Content-Range
	min, max, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing the content-range [%s]: %s", resp.Header.Get("Content-Range"), err)
	}

	// Prepare the response
	response := &RegisteredServicesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.RegisteredServices)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetRegisteredServices retrieves a named service registered with the Application Framework by id.
func (endpoint *Endpoint) GetRegisteredServices(ctx context.Context, uuid int) (*RegisteredService, error) {

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/gui_app_framework/named_services/"+strconv.Itoa(uuid), nil)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *RegisteredService

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}
