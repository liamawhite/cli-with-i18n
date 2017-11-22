package shared

import (
	"github.com/liamawhite/cli-with-i18n/actor/actionerror"
	"github.com/liamawhite/cli-with-i18n/actor/pushaction"
	"github.com/liamawhite/cli-with-i18n/actor/sharedaction"
	"github.com/liamawhite/cli-with-i18n/actor/v2action"
	"github.com/liamawhite/cli-with-i18n/api/cloudcontroller/ccerror"
	"github.com/liamawhite/cli-with-i18n/api/uaa"
	"github.com/liamawhite/cli-with-i18n/command/translatableerror"
	"github.com/liamawhite/cli-with-i18n/util/manifest"
)

func HandleError(err error) error {
	switch e := err.(type) {
	case ccerror.APINotFoundError:
		return translatableerror.APINotFoundError(e)
	case ccerror.RequestError:
		return translatableerror.APIRequestError(e)
	case ccerror.SSLValidationHostnameError:
		return translatableerror.SSLCertError(e)
	case ccerror.UnverifiedServerError:
		return translatableerror.InvalidSSLCertError{API: e.URL}

	case ccerror.JobFailedError:
		return translatableerror.JobFailedError(e)
	case ccerror.JobTimeoutError:
		return translatableerror.JobTimeoutError{JobGUID: e.JobGUID}

	case uaa.BadCredentialsError:
		return translatableerror.BadCredentialsError{}
	case uaa.InvalidAuthTokenError:
		return translatableerror.InvalidRefreshTokenError{}

	case sharedaction.NotLoggedInError:
		return translatableerror.NotLoggedInError(e)
	case sharedaction.NoOrganizationTargetedError:
		return translatableerror.NoOrganizationTargetedError(e)
	case sharedaction.NoSpaceTargetedError:
		return translatableerror.NoSpaceTargetedError(e)

	case actionerror.ApplicationNotFoundError:
		return translatableerror.ApplicationNotFoundError{Name: e.Name}
	case v2action.OrganizationNotFoundError:
		return translatableerror.OrganizationNotFoundError{Name: e.Name}
	case v2action.SecurityGroupNotFoundError:
		return translatableerror.SecurityGroupNotFoundError(e)
	case v2action.ServiceInstanceNotFoundError:
		return translatableerror.ServiceInstanceNotFoundError(e)
	case v2action.SpaceNotFoundError:
		return translatableerror.SpaceNotFoundError{Name: e.Name}
	case v2action.StackNotFoundError:
		return translatableerror.StackNotFoundError(e)
	case actionerror.HTTPHealthCheckInvalidError:
		return translatableerror.HTTPHealthCheckInvalidError{}
	case v2action.RouteInDifferentSpaceError:
		return translatableerror.RouteInDifferentSpaceError(e)
	case v2action.FileChangedError:
		return translatableerror.FileChangedError(e)
	case sharedaction.EmptyDirectoryError:
		return translatableerror.EmptyDirectoryError(e)
	case v2action.DomainNotFoundError:
		return translatableerror.DomainNotFoundError(e)
	case actionerror.NoMatchingDomainError:
		return translatableerror.NoMatchingDomainError(e)
	case actionerror.InvalidHTTPRouteSettings:
		return translatableerror.PortNotAllowedWithHTTPDomainError(e)

	case pushaction.AppNotFoundInManifestError:
		return translatableerror.AppNotFoundInManifestError(e)
	case pushaction.CommandLineOptionsWithMultipleAppsError:
		return translatableerror.CommandLineArgsWithMultipleAppsError{}
	case pushaction.NoDomainsFoundError:
		return translatableerror.NoDomainsFoundError{}
	case pushaction.NonexistentAppPathError:
		return translatableerror.FileNotFoundError(e)
	case pushaction.MissingNameError:
		return translatableerror.RequiredNameForPushError{}
	case pushaction.UploadFailedError:
		return translatableerror.UploadFailedError{Err: HandleError(e.Err)}

	case manifest.ManifestCreationError:
		return translatableerror.ManifestCreationError(e)
	}

	return err
}
