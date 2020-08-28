// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Networking.V1
{

    [OutputType]
    public sealed class ServiceBackendPort
    {
        /// <summary>
        /// Name is the name of the port on the Service. This is a mutually exclusive setting with "Number".
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with "Name".
        /// </summary>
        public readonly int Number;

        [OutputConstructor]
        private ServiceBackendPort(
            string name,

            int number)
        {
            Name = name;
            Number = number;
        }
    }
}