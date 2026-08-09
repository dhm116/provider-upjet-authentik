// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	application "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/application/application"
	entitlement "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/application/entitlement"
	blueprint "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/authentik/blueprint"
	brand "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/authentik/brand"
	group "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/authentik/group"
	role "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/authentik/role"
	token "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/authentik/token"
	user "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/authentik/user"
	keypair "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/certificate/keypair"
	connectoragent "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/endpoints/connectoragent"
	connectoragentenrollmenttoken "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/endpoints/connectoragentenrollmenttoken"
	deviceaccessgroup "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/endpoints/deviceaccessgroup"
	googlechromeconnector "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/endpoints/googlechromeconnector"
	license "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/enterprise/license"
	rule "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/event/rule"
	transport "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/event/transport"
	flow "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/flow/flow"
	stagebinding "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/flow/stagebinding"
	outpost "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/outpost/outpost"
	providerattachment "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/outpost/providerattachment"
	binding "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/policy/binding"
	dummy "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/policy/dummy"
	eventmatcher "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/policy/eventmatcher"
	expiry "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/policy/expiry"
	expression "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/policy/expression"
	geoip "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/policy/geoip"
	password "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/policy/password"
	reputation "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/policy/reputation"
	uniquepassword "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/policy/uniquepassword"
	mappingnotification "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingnotification"
	mappingprovidergoogleworkspace "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingprovidergoogleworkspace"
	mappingprovidermicrosoftentra "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingprovidermicrosoftentra"
	mappingproviderrac "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingproviderrac"
	mappingproviderradius "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingproviderradius"
	mappingprovidersaml "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingprovidersaml"
	mappingproviderscim "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingproviderscim"
	mappingproviderscope "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingproviderscope"
	mappingsourcekerberos "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingsourcekerberos"
	mappingsourceldap "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingsourceldap"
	mappingsourceoauth "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingsourceoauth"
	mappingsourceplex "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingsourceplex"
	mappingsourcesaml "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingsourcesaml"
	mappingsourcescim "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/property/mappingsourcescim"
	googleworkspace "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/provider/googleworkspace"
	ldap "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/provider/ldap"
	microsoftentra "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/provider/microsoftentra"
	oauth2 "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/provider/oauth2"
	proxy "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/provider/proxy"
	rac "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/provider/rac"
	radius "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/provider/radius"
	saml "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/provider/saml"
	scim "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/provider/scim"
	ssf "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/provider/ssf"
	wsfederation "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/provider/wsfederation"
	providerconfig "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/providerconfig"
	endpoint "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/rac/endpoint"
	initialpermissions "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/rbac/initialpermissions"
	permissionrole "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/rbac/permissionrole"
	permissionuser "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/rbac/permissionuser"
	docker "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/serviceconnection/docker"
	kubernetes "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/serviceconnection/kubernetes"
	kerberos "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/source/kerberos"
	ldapsource "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/source/ldap"
	oauth "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/source/oauth"
	plex "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/source/plex"
	samlsource "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/source/saml"
	scimsource "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/source/scim"
	telegram "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/source/telegram"
	accountlockdown "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/accountlockdown"
	authenticatorduo "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/authenticatorduo"
	authenticatoremail "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/authenticatoremail"
	authenticatorendpointgdtc "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/authenticatorendpointgdtc"
	authenticatorsms "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/authenticatorsms"
	authenticatorstatic "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/authenticatorstatic"
	authenticatortotp "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/authenticatortotp"
	authenticatorvalidate "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/authenticatorvalidate"
	authenticatorwebauthn "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/authenticatorwebauthn"
	captcha "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/captcha"
	consent "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/consent"
	deny "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/deny"
	dummystage "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/dummy"
	email "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/email"
	endpoints "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/endpoints"
	identification "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/identification"
	invitation "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/invitation"
	mutualtls "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/mutualtls"
	passwordstage "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/password"
	prompt "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/prompt"
	promptfield "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/promptfield"
	redirect "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/redirect"
	source "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/source"
	userdelete "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/userdelete"
	userlogin "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/userlogin"
	userlogout "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/userlogout"
	userwrite "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/stage/userwrite"
	settings "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/system/settings"
	schedule "github.com/dhm116/provider-upjet-authentik/internal/controller/namespaced/task/schedule"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		entitlement.Setup,
		blueprint.Setup,
		brand.Setup,
		group.Setup,
		role.Setup,
		token.Setup,
		user.Setup,
		keypair.Setup,
		connectoragent.Setup,
		connectoragentenrollmenttoken.Setup,
		deviceaccessgroup.Setup,
		googlechromeconnector.Setup,
		license.Setup,
		rule.Setup,
		transport.Setup,
		flow.Setup,
		stagebinding.Setup,
		outpost.Setup,
		providerattachment.Setup,
		binding.Setup,
		dummy.Setup,
		eventmatcher.Setup,
		expiry.Setup,
		expression.Setup,
		geoip.Setup,
		password.Setup,
		reputation.Setup,
		uniquepassword.Setup,
		mappingnotification.Setup,
		mappingprovidergoogleworkspace.Setup,
		mappingprovidermicrosoftentra.Setup,
		mappingproviderrac.Setup,
		mappingproviderradius.Setup,
		mappingprovidersaml.Setup,
		mappingproviderscim.Setup,
		mappingproviderscope.Setup,
		mappingsourcekerberos.Setup,
		mappingsourceldap.Setup,
		mappingsourceoauth.Setup,
		mappingsourceplex.Setup,
		mappingsourcesaml.Setup,
		mappingsourcescim.Setup,
		googleworkspace.Setup,
		ldap.Setup,
		microsoftentra.Setup,
		oauth2.Setup,
		proxy.Setup,
		rac.Setup,
		radius.Setup,
		saml.Setup,
		scim.Setup,
		ssf.Setup,
		wsfederation.Setup,
		providerconfig.Setup,
		endpoint.Setup,
		initialpermissions.Setup,
		permissionrole.Setup,
		permissionuser.Setup,
		docker.Setup,
		kubernetes.Setup,
		kerberos.Setup,
		ldapsource.Setup,
		oauth.Setup,
		plex.Setup,
		samlsource.Setup,
		scimsource.Setup,
		telegram.Setup,
		accountlockdown.Setup,
		authenticatorduo.Setup,
		authenticatoremail.Setup,
		authenticatorendpointgdtc.Setup,
		authenticatorsms.Setup,
		authenticatorstatic.Setup,
		authenticatortotp.Setup,
		authenticatorvalidate.Setup,
		authenticatorwebauthn.Setup,
		captcha.Setup,
		consent.Setup,
		deny.Setup,
		dummystage.Setup,
		email.Setup,
		endpoints.Setup,
		identification.Setup,
		invitation.Setup,
		mutualtls.Setup,
		passwordstage.Setup,
		prompt.Setup,
		promptfield.Setup,
		redirect.Setup,
		source.Setup,
		userdelete.Setup,
		userlogin.Setup,
		userlogout.Setup,
		userwrite.Setup,
		settings.Setup,
		schedule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.SetupGated,
		entitlement.SetupGated,
		blueprint.SetupGated,
		brand.SetupGated,
		group.SetupGated,
		role.SetupGated,
		token.SetupGated,
		user.SetupGated,
		keypair.SetupGated,
		connectoragent.SetupGated,
		connectoragentenrollmenttoken.SetupGated,
		deviceaccessgroup.SetupGated,
		googlechromeconnector.SetupGated,
		license.SetupGated,
		rule.SetupGated,
		transport.SetupGated,
		flow.SetupGated,
		stagebinding.SetupGated,
		outpost.SetupGated,
		providerattachment.SetupGated,
		binding.SetupGated,
		dummy.SetupGated,
		eventmatcher.SetupGated,
		expiry.SetupGated,
		expression.SetupGated,
		geoip.SetupGated,
		password.SetupGated,
		reputation.SetupGated,
		uniquepassword.SetupGated,
		mappingnotification.SetupGated,
		mappingprovidergoogleworkspace.SetupGated,
		mappingprovidermicrosoftentra.SetupGated,
		mappingproviderrac.SetupGated,
		mappingproviderradius.SetupGated,
		mappingprovidersaml.SetupGated,
		mappingproviderscim.SetupGated,
		mappingproviderscope.SetupGated,
		mappingsourcekerberos.SetupGated,
		mappingsourceldap.SetupGated,
		mappingsourceoauth.SetupGated,
		mappingsourceplex.SetupGated,
		mappingsourcesaml.SetupGated,
		mappingsourcescim.SetupGated,
		googleworkspace.SetupGated,
		ldap.SetupGated,
		microsoftentra.SetupGated,
		oauth2.SetupGated,
		proxy.SetupGated,
		rac.SetupGated,
		radius.SetupGated,
		saml.SetupGated,
		scim.SetupGated,
		ssf.SetupGated,
		wsfederation.SetupGated,
		providerconfig.SetupGated,
		endpoint.SetupGated,
		initialpermissions.SetupGated,
		permissionrole.SetupGated,
		permissionuser.SetupGated,
		docker.SetupGated,
		kubernetes.SetupGated,
		kerberos.SetupGated,
		ldapsource.SetupGated,
		oauth.SetupGated,
		plex.SetupGated,
		samlsource.SetupGated,
		scimsource.SetupGated,
		telegram.SetupGated,
		accountlockdown.SetupGated,
		authenticatorduo.SetupGated,
		authenticatoremail.SetupGated,
		authenticatorendpointgdtc.SetupGated,
		authenticatorsms.SetupGated,
		authenticatorstatic.SetupGated,
		authenticatortotp.SetupGated,
		authenticatorvalidate.SetupGated,
		authenticatorwebauthn.SetupGated,
		captcha.SetupGated,
		consent.SetupGated,
		deny.SetupGated,
		dummystage.SetupGated,
		email.SetupGated,
		endpoints.SetupGated,
		identification.SetupGated,
		invitation.SetupGated,
		mutualtls.SetupGated,
		passwordstage.SetupGated,
		prompt.SetupGated,
		promptfield.SetupGated,
		redirect.SetupGated,
		source.SetupGated,
		userdelete.SetupGated,
		userlogin.SetupGated,
		userlogout.SetupGated,
		userwrite.SetupGated,
		settings.SetupGated,
		schedule.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		application.SetupWebhookWithManager,
		entitlement.SetupWebhookWithManager,
		blueprint.SetupWebhookWithManager,
		brand.SetupWebhookWithManager,
		group.SetupWebhookWithManager,
		role.SetupWebhookWithManager,
		token.SetupWebhookWithManager,
		user.SetupWebhookWithManager,
		keypair.SetupWebhookWithManager,
		connectoragent.SetupWebhookWithManager,
		connectoragentenrollmenttoken.SetupWebhookWithManager,
		deviceaccessgroup.SetupWebhookWithManager,
		googlechromeconnector.SetupWebhookWithManager,
		license.SetupWebhookWithManager,
		rule.SetupWebhookWithManager,
		transport.SetupWebhookWithManager,
		flow.SetupWebhookWithManager,
		stagebinding.SetupWebhookWithManager,
		outpost.SetupWebhookWithManager,
		providerattachment.SetupWebhookWithManager,
		binding.SetupWebhookWithManager,
		dummy.SetupWebhookWithManager,
		eventmatcher.SetupWebhookWithManager,
		expiry.SetupWebhookWithManager,
		expression.SetupWebhookWithManager,
		geoip.SetupWebhookWithManager,
		password.SetupWebhookWithManager,
		reputation.SetupWebhookWithManager,
		uniquepassword.SetupWebhookWithManager,
		mappingnotification.SetupWebhookWithManager,
		mappingprovidergoogleworkspace.SetupWebhookWithManager,
		mappingprovidermicrosoftentra.SetupWebhookWithManager,
		mappingproviderrac.SetupWebhookWithManager,
		mappingproviderradius.SetupWebhookWithManager,
		mappingprovidersaml.SetupWebhookWithManager,
		mappingproviderscim.SetupWebhookWithManager,
		mappingproviderscope.SetupWebhookWithManager,
		mappingsourcekerberos.SetupWebhookWithManager,
		mappingsourceldap.SetupWebhookWithManager,
		mappingsourceoauth.SetupWebhookWithManager,
		mappingsourceplex.SetupWebhookWithManager,
		mappingsourcesaml.SetupWebhookWithManager,
		mappingsourcescim.SetupWebhookWithManager,
		googleworkspace.SetupWebhookWithManager,
		ldap.SetupWebhookWithManager,
		microsoftentra.SetupWebhookWithManager,
		oauth2.SetupWebhookWithManager,
		proxy.SetupWebhookWithManager,
		rac.SetupWebhookWithManager,
		radius.SetupWebhookWithManager,
		saml.SetupWebhookWithManager,
		scim.SetupWebhookWithManager,
		ssf.SetupWebhookWithManager,
		wsfederation.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
		endpoint.SetupWebhookWithManager,
		initialpermissions.SetupWebhookWithManager,
		permissionrole.SetupWebhookWithManager,
		permissionuser.SetupWebhookWithManager,
		docker.SetupWebhookWithManager,
		kubernetes.SetupWebhookWithManager,
		kerberos.SetupWebhookWithManager,
		ldapsource.SetupWebhookWithManager,
		oauth.SetupWebhookWithManager,
		plex.SetupWebhookWithManager,
		samlsource.SetupWebhookWithManager,
		scimsource.SetupWebhookWithManager,
		telegram.SetupWebhookWithManager,
		accountlockdown.SetupWebhookWithManager,
		authenticatorduo.SetupWebhookWithManager,
		authenticatoremail.SetupWebhookWithManager,
		authenticatorendpointgdtc.SetupWebhookWithManager,
		authenticatorsms.SetupWebhookWithManager,
		authenticatorstatic.SetupWebhookWithManager,
		authenticatortotp.SetupWebhookWithManager,
		authenticatorvalidate.SetupWebhookWithManager,
		authenticatorwebauthn.SetupWebhookWithManager,
		captcha.SetupWebhookWithManager,
		consent.SetupWebhookWithManager,
		deny.SetupWebhookWithManager,
		dummystage.SetupWebhookWithManager,
		email.SetupWebhookWithManager,
		endpoints.SetupWebhookWithManager,
		identification.SetupWebhookWithManager,
		invitation.SetupWebhookWithManager,
		mutualtls.SetupWebhookWithManager,
		passwordstage.SetupWebhookWithManager,
		prompt.SetupWebhookWithManager,
		promptfield.SetupWebhookWithManager,
		redirect.SetupWebhookWithManager,
		source.SetupWebhookWithManager,
		userdelete.SetupWebhookWithManager,
		userlogin.SetupWebhookWithManager,
		userlogout.SetupWebhookWithManager,
		userwrite.SetupWebhookWithManager,
		settings.SetupWebhookWithManager,
		schedule.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}
