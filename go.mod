module github.com/dtcookie/terraform-provider-dynatrace

go 1.15

require (
	github.com/dtcookie/dynatrace/api/config v1.0.0
	github.com/dtcookie/dynatrace/api/config/alertingprofiles v1.0.0
	github.com/dtcookie/dynatrace/api/config/customservices v1.0.6
	github.com/dtcookie/dynatrace/api/config/dashboards v1.0.1
	github.com/dtcookie/dynatrace/api/config/maintenancewindows v1.0.0
	github.com/dtcookie/dynatrace/api/config/managementzones v1.0.1
	github.com/dtcookie/dynatrace/api/config/requestattributes v1.0.1
	github.com/dtcookie/dynatrace/rest v1.0.13
	github.com/dtcookie/dynatrace/terraform v1.0.3
	github.com/dtcookie/opt v1.0.0
	github.com/google/uuid v1.1.2
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.3.0
)
