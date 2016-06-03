package webmethods

import (
	"github.com/imwithye/gitignore/template"
	"strings"
)

func init() {
	ignore := []string{
		"**/IntegrationServer/datastore/",
		"**/IntegrationServer/db/",
		"**/IntegrationServer/DocumentStore/",
		"**/IntegrationServer/lib/",
		"**/IntegrationServer/logs/",
		"**/IntegrationServer/replicate/",
		"**/IntegrationServer/sdk/",
		"**/IntegrationServer/support/",
		"**/IntegrationServer/update/",
		"**/IntegrationServer/userFtpRoot/",
		"**/IntegrationServer/web/",
		"**/IntegrationServer/WmRepository4/",
		"**/IntegrationServer/XAStore/",
		"**/IntegrationServer/packages/Wm*/",
		"",
	}
	template.Add("github/global/webmethods", strings.Join(ignore, "\n"))
}
