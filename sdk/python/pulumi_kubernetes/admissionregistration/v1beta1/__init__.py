# coding=utf-8
# *** WARNING: this file was generated by pulumigen. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .MutatingWebhookConfiguration import *
from .MutatingWebhookConfigurationList import *
from .ValidatingWebhookConfiguration import *
from .ValidatingWebhookConfigurationList import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from ... import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfiguration":
                return MutatingWebhookConfiguration(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:admissionregistration.k8s.io/v1beta1:MutatingWebhookConfigurationList":
                return MutatingWebhookConfigurationList(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingWebhookConfiguration":
                return ValidatingWebhookConfiguration(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "kubernetes:admissionregistration.k8s.io/v1beta1:ValidatingWebhookConfigurationList":
                return ValidatingWebhookConfigurationList(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("kubernetes", "admissionregistration.k8s.io/v1beta1", _module_instance)

_register_module()
