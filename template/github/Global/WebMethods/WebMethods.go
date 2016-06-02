package WebMethods

import "strings"
import "github.com/imwithye/gitignore/template"

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
	}
	template.SetTemplate("GitHub/Global/WebMethods", strings.Join(ignore, "\n"))
}
