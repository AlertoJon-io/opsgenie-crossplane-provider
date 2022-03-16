package user

import "github.com/crossplane/terrajet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("opsgenie_user", func(r *config.Resource) {

		// we need to override the default group that terrajet generated for
		// this resource, which would be "github"
		r.ShortGroup = "user"

		// Identifier for this resource is assigned by the provider. In other
		// words it is not simply the name of the resource.
		r.ExternalName = config.IdentifierFromProvider

		// This resource need the repository in which branch would be created
		// as an input. And by defining it as a reference to Repository
		// object, we can build cross resource referencing. See
		// repositoryRef in the example in the Testing section below.
		r.References["opsgenie_user"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-jet-github/apis/user/v1alpha1.User",
		}
	})
}