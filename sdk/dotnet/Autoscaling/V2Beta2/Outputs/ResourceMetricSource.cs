// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2
{

    [OutputType]
    public sealed class ResourceMetricSource
    {
        /// <summary>
        /// name is the name of the resource in question.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// target specifies the target value for the given metric
        /// </summary>
        public readonly Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.MetricTarget Target;

        [OutputConstructor]
        private ResourceMetricSource(
            string name,

            Pulumi.Kubernetes.Types.Outputs.Autoscaling.V2Beta2.MetricTarget target)
        {
            Name = name;
            Target = target;
        }
    }
}
